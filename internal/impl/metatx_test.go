package impl

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/valist-io/registry/internal/contract/metatx"
	"github.com/valist-io/registry/internal/core"
)

func (s *ClientSuite) TestGenerateEIP712Challenge() {
	ctx := context.Background()

	orgMeta := &core.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	data, err := json.Marshal(orgMeta)
	s.Require().NoError(err, "Failed to unmarshal meta")

	metaCID, err := s.client.AddFile(ctx, data)
	s.Require().NoError(err, "Failed to add meta file")

	txopts, err := s.client.transact()
	s.Require().NoError(err, "Failed to get transaction opts")

	txopts.NoSend = true
	tx, err := s.client.valist.CreateOrganization(txopts, metaCID.String())
	s.Require().NoError(err, "Failed to create transaction")

	eip712Challenge, domainSeparator, err := s.client.GenerateEIP712Challenge(txopts, tx)
	s.Require().NoError(err, "Failed to generate EIP712 challenge")

	// req := map[string]{
	// 	"from":          txOpts.From.Hex(),
	// 	"to":            tx.To().Hex(),
	// 	"token":         "0x0000000000000000000000000000000000000000",
	// 	"txGas":         fmt.Sprintf("%d", txOpts.GasLimit),
	// 	"tokenGasPrice": "0",
	// 	"batchId":       "0",
	// 	"batchNonce":    nonce.String(),
	// 	"deadline":      fmt.Sprintf("%d", time.Now().Unix()+3600), // 1 hour timeout
	// 	"data":          tx.Data(),
	// }

	nonce, err := s.client.forwarder.GetNonce(&bind.CallOpts{}, txopts.From, big.NewInt(0))
	s.Require().NoError(err, "Failed to get nonce")

	req := metatx.ERC20ForwardRequestTypesERC20ForwardRequest{
		From:          txopts.From,
		To:            *tx.To(),
		Token:         common.HexToAddress("0x0"),
		TxGas:         big.NewInt(int64(txopts.GasLimit)),
		TokenGasPrice: big.NewInt(0),
		BatchId:       big.NewInt(0),
		BatchNonce:    nonce,
		Deadline:      big.NewInt(time.Now().Add(time.Hour).Unix()),
		Data:          tx.Data(),
	}

	tx2, err := s.client.forwarder.ExecuteEIP712(txopts, req, domainSeparator, tx.RawSignatureValues())
	s.Require().NoError(err, "Failed to execute eip712")
}
