package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	signer "github.com/ethereum/go-ethereum/signer/core"
)

type biconomyEIP712Message struct {
	From          string   `json:"from"`
	To            string   `json:"to"`
	Token         string   `json:"token"`
	TxGas         uint64   `json:"txGas"`
	TokenGasPrice string   `json:"tokenGasPrice"`
	BatchId       uint     `json:"batchId"`
	BatchNonce    *big.Int `json:"batchNonce"`
	Deadline      string   `json:"deadline"`
	Data          string   `json:"data"`
}
type biconomyRequest struct {
	To            string        `json:"to"`
	ApiId         string        `json:"apiId"`
	Params        []interface{} `json:"params"`
	From          string        `json:"from"`
	SignatureType string        `json:"signatureType"`
}

type biconomyResponse struct {
	TxHash string `json:"txHash"`
}

// maps function names to public Biconomy API IDs
var biconomyFunctionIDMap = map[string]string{
	"clearPendingKey":       "a0dfd7b2-fb2b-46da-a662-3cbb87c7b83e",
	"clearPendingRelease":   "b95d7f2d-6d40-4690-b7df-ec36928aaf77",
	"clearPendingThreshold": "f154fe5a-cd81-4a31-8536-6ea999795f56",
	"createOrganization":    "7cb293ac-5ed6-4dd8-9956-eb5a9a236403",
	"createRepository":      "3b40c07a-d9dd-401a-913b-ef395648ba4d",
	"setOrgMeta":            "1292cba4-8b4e-4828-8989-e2583017cda7",
	"setRepoMeta":           "1857aa6a-b334-4b6a-bf7c-959d5581e8d4",
	"voteKey":               "82d84700-7a9a-44f5-865d-f34badb00852",
	"voteRelease":           "c8fc037a-dc5c-4fe3-b2fd-f8c602986d72",
	"voteThreshold":         "f0b640b6-4280-4cf0-afca-0d62046cee09",
	"grantRole":             "17ec42d7-9f19-407c-8131-3033f7dcc142",
	"init":                  "5336e4c2-fc5c-49bd-b41d-9990dde03982",
	"linkNameToID":          "8fc893ff-08e1-4cda-9264-62f6467d91a8",
	"overrideNameToID":      "0455fbcd-4d1e-45ec-b0ce-5eaf73169b3e",
	"renounceRole":          "08c8a75f-e9d2-4e9d-82e9-8f6c5b2bf8a0",
	"revokeRole":            "d4040355-b755-4a1a-9f16-0f0462bd56d1",
}

var biconomyForwarderAddressMap = map[string]string{
	"80001": "0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b", // Polygon mumbai
	"137":   "0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8", // Polygon mainnet
	"1":     "0x84a0856b038eaAd1cC7E297cF34A7e72685A8693", // Ethereum mainnet
}

const eip712PrimaryType = "ERC20ForwardRequest"

var eip712Types = signer.Types{
	"EIP712Domain": []signer.Type{
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "verifyingContract", Type: "address"},
		{Name: "salt", Type: "bytes32"},
	},
	"ERC20ForwardRequest": []signer.Type{
		{Name: "from", Type: "address"},
		{Name: "to", Type: "address"},
		{Name: "token", Type: "address"},
		{Name: "txGas", Type: "uint256"},
		{Name: "tokenGasPrice", Type: "uint256"},
		{Name: "batchId", Type: "uint256"},
		{Name: "batchNonce", Type: "uint256"},
		{Name: "deadline", Type: "uint256"},
		{Name: "data", Type: "bytes"},
	},
}

func getFunctionName(function interface{}) string {
	// get full path to function like 'github.com/valist-io/registry/internal/contract/valist.(*ValistTransactor).CreateOrganization-fm'
	functionName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	// get last section of function name minus '.' prefix
	functionName = filepath.Ext(functionName)[1:]
	// get plain function name minus '-fm' suffix
	functionName = strings.Split(functionName, "-")[0]
	// convert string to rune to manipulate individual chars
	fn := []rune(functionName)
	// lowercase first character of rune
	fn[0] = unicode.ToLower(fn[0])
	return string(fn)
}

// Modifies TransactionOpts to prepare tx object meta
// Sets From to the 0x0 address to enable gas estimation with an empty balance
// Sets Signer to EmptySigner to generate tx.Data() without signing
// Sets NoSend to true to keep tx local and passable to SendMetaTx
func PrepareMetaTx(client *Client, txopts *bind.TransactOpts) {
	if client.metaTx {
		txopts.From = common.HexToAddress("0x0000000000000000000000000000000000000000")
		txopts.Signer = client.EmptySigner
		txopts.NoSend = true
	}
}

func SendMetaTx(client *Client, tx *types.Transaction, function interface{}) (*types.Transaction, error) {

	if !client.metaTx {
		return tx, nil
	}

	forwarderAddress, ok := biconomyForwarderAddressMap[client.chainID.String()]
	if !ok {
		return nil, fmt.Errorf("MetaTransactions not available on this network")
	}

	opts := bind.CallOpts{
		From: client.account.Address,
	}

	nonce, err := client.forwarder.GetNonce(&opts, opts.From, big.NewInt(0))
	if err != nil {
		return nil, err
	}

	eip712Request := signer.TypedData{
		Types:       eip712Types,
		PrimaryType: eip712PrimaryType,
		Domain: signer.TypedDataDomain{
			Name:              "Biconomy Forwarder",
			Version:           "1",
			Salt:              hexutil.Encode(common.LeftPadBytes(client.chainID.Bytes(), 32)),
			VerifyingContract: forwarderAddress,
		},
		Message: signer.TypedDataMessage{
			"from":          opts.From.Hex(),
			"to":            tx.To().Hex(),
			"token":         "0x0000000000000000000000000000000000000000",
			"txGas":         tx.Gas(),
			"tokenGasPrice": "0",
			"batchId":       0,
			"batchNonce":    nonce,
			"deadline":      time.Now().Unix() + 3600, // 1 hour timeout
			"data":          hexutil.Encode(tx.Data()),
		},
	}

	data, err := json.Marshal(eip712Request)
	if err != nil {
		return nil, err
	}

	eip712Message := biconomyEIP712Message{
		From:          opts.From.Hex(),
		To:            tx.To().Hex(),
		Token:         "0x0000000000000000000000000000000000000000",
		TxGas:         tx.Gas(),
		TokenGasPrice: "0",
		BatchId:       0,
		BatchNonce:    nonce,
		Deadline:      fmt.Sprint(time.Now().Unix() + 3600), // 1 hour timeout
		Data:          hexutil.Encode(tx.Data()),
	}

	signature, err := client.wallet.SignData(client.account, signer.DataTyped.Mime, data)
	if err != nil {
		return nil, err
	}

	domainSeparator, err := eip712Request.HashStruct("EIP712Domain", eip712Request.Domain.Map())
	if err != nil {
		return nil, err
	}

	functionName := getFunctionName(function)

	functionID, ok := biconomyFunctionIDMap[functionName]
	if !ok {
		return nil, fmt.Errorf("Function name '%v' does not have a registered functionID", functionName)
	}

	reqBody, err := json.Marshal(biconomyRequest{
		To:            tx.To().Hex(),
		ApiId:         functionID,
		Params:        []interface{}{eip712Message, hexutil.Encode(domainSeparator), hexutil.Encode(signature)},
		From:          opts.From.Hex(),
		SignatureType: "EIP712_SIGN",
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.biconomy.io/api/v2/meta-tx/native", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("x-api-key", "qLW9TRUjQ.f77d2f86-c76a-4b9c-b1ee-0453d0ead878") // public api key
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var txResp biconomyResponse
	err = json.NewDecoder(resp.Body).Decode(&txResp)
	// @TODO: Return biconomy response in JSON on error to help debugging
	if err != nil {
		return nil, err
	}

	if txResp.TxHash == "" {
		return nil, fmt.Errorf("Could not parse Biconomy response")
	}

	metaTx, _, err := client.eth.TransactionByHash(context.Background(), common.HexToHash(txResp.TxHash))
	if err != nil {
		return nil, err
	}

	return metaTx, err
}
