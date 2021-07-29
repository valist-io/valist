package impl

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	signer "github.com/ethereum/go-ethereum/signer/core"
)

func (client *Client) GenerateEIP712Challenge(txOpts *bind.TransactOpts, tx *types.Transaction) ([]byte, []byte, error) {

	forwarderAddress := "0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b"
	salt := fmt.Sprintf("0x%s", common.Bytes2Hex(common.LeftPadBytes(chainID.Bytes(), 32)))
	nonce, err := client.forwarder.GetNonce(&bind.CallOpts{}, txOpts.From, big.NewInt(0))

	if err != nil {
		return nil, nil, err
	}

	signerData := signer.TypedData{
		Types: signer.Types{
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
		},
		PrimaryType: "ERC20ForwardRequest",
		Domain: signer.TypedDataDomain{
			Name:              "Biconomy Forwarder",
			Version:           "1",
			Salt:              salt,
			VerifyingContract: forwarderAddress,
		},
		Message: signer.TypedDataMessage{
			"from":          txOpts.From.Hex(),
			"to":            tx.To().Hex(),
			"token":         "0x0000000000000000000000000000000000000000",
			"txGas":         fmt.Sprintf("%d", txOpts.GasLimit),
			"tokenGasPrice": "0",
			"batchId":       "0",
			"batchNonce":    nonce.String(),
			"deadline":      fmt.Sprintf("%d", time.Now().Unix()+3600), // 1 hour timeout
			"data":          tx.Data(),
		},
	}

	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		return nil, nil, err
	}

	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		return nil, nil, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	// challengeHash := crypto.Keccak256Hash(rawData)
	return rawData, domainSeparator, nil
}
