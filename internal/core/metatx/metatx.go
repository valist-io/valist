package metatx

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/valist-io/gasless"
	"github.com/valist-io/gasless/mexa"

	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/core/basetx"
)

var (
	emptyHash    = common.HexToHash("0x0")
	emptyAddress = common.HexToAddress("0x0")
)

const (
	clearPendingKeyBFID       = "a0dfd7b2-fb2b-46da-a662-3cbb87c7b83e"
	clearPendingReleaseBFID   = "b95d7f2d-6d40-4690-b7df-ec36928aaf77"
	clearPendingThresholdBFID = "f154fe5a-cd81-4a31-8536-6ea999795f56"
	createOrganizationBFID    = "7cb293ac-5ed6-4dd8-9956-eb5a9a236403"
	createRepositoryBFID      = "3b40c07a-d9dd-401a-913b-ef395648ba4d"
	setOrgMetaBFID            = "1292cba4-8b4e-4828-8989-e2583017cda7"
	setRepoMetaBFID           = "1857aa6a-b334-4b6a-bf7c-959d5581e8d4"
	voteKeyBFID               = "82d84700-7a9a-44f5-865d-f34badb00852"
	voteReleaseBFID           = "c8fc037a-dc5c-4fe3-b2fd-f8c602986d72"
	voteThresholdBFID         = "f0b640b6-4280-4cf0-afca-0d62046cee09"
	grantRoleBFID             = "17ec42d7-9f19-407c-8131-3033f7dcc142"
	initBFID                  = "5336e4c2-fc5c-49bd-b41d-9990dde03982"
	linkNameToIDBFID          = "8fc893ff-08e1-4cda-9264-62f6467d91a8"
	overrideNameToIDBFID      = "0455fbcd-4d1e-45ec-b0ce-5eaf73169b3e"
	renounceRoleBFID          = "08c8a75f-e9d2-4e9d-82e9-8f6c5b2bf8a0"
	revokeRoleBFID            = "d4040355-b755-4a1a-9f16-0f0462bd56d1"
)

type Transactor struct {
	eth *ethclient.Client

	base core.TransactorAPI
	meta *mexa.Mexa

	account accounts.Account
	wallet  accounts.Wallet
	signer  gasless.Signer
}

func NewTransactor(
	ctx context.Context,
	eth *ethclient.Client,
	valist *valist.Valist,
	registry *registry.ValistRegistry,
	account accounts.Account,
	wallet accounts.Wallet) (core.TransactorAPI, error) {
	meta, err := mexa.NewMexa(ctx, eth, "qLW9TRUjQ.f77d2f86-c76a-4b9c-b1ee-0453d0ead878", big.NewInt(0)) // public key
	if err != nil {
		return nil, err
	}

	signer := gasless.NewWalletSigner(account, wallet)

	base := basetx.NewTransactor(valist, registry)

	return &Transactor{
		eth:     eth,
		base:    base,
		meta:    meta,
		account: account,
		wallet:  wallet,
		signer:  signer,
	}, nil
}

func setMetaTransactOpts(txopts *bind.TransactOpts) {
	txopts.Signer = func(a common.Address, t *types.Transaction) (*types.Transaction, error) { return t, nil }
	txopts.From = emptyAddress
	txopts.NoSend = true
}

func (t *Transactor) newMessage(ctx context.Context, tx *types.Transaction, apiId string) (gasless.EIP712Message, error) {
	nonce, err := t.meta.Nonce(ctx, t.account.Address)
	if err != nil {
		return nil, err
	}

	return &mexa.Message{
		ApiId:         apiId,
		From:          t.account.Address,
		To:            *tx.To(),
		Token:         emptyAddress,
		TxGas:         tx.Gas(),
		TokenGasPrice: "0",
		BatchId:       big.NewInt(0),
		BatchNonce:    nonce,
		Deadline:      big.NewInt(time.Now().Unix() + 3600),
		Data:          hexutil.Encode(tx.Data()),
	}, nil
}
