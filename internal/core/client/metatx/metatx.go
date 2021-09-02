// Package metatx defines a Transactor that uses meta transactions to pay gas fees on behalf of a user.
package metatx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/valist-io/gasless"

	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core/client"
	"github.com/valist-io/registry/internal/core/config"
)

const (
	clearPendingKeyBFID       = "a0dfd7b2-fb2b-46da-a662-3cbb87c7b83e" //nolint
	clearPendingReleaseBFID   = "b95d7f2d-6d40-4690-b7df-ec36928aaf77" //nolint
	clearPendingThresholdBFID = "f154fe5a-cd81-4a31-8536-6ea999795f56" //nolint
	createOrganizationBFID    = "7cb293ac-5ed6-4dd8-9956-eb5a9a236403" //nolint
	createRepositoryBFID      = "3b40c07a-d9dd-401a-913b-ef395648ba4d" //nolint
	setOrgMetaBFID            = "1292cba4-8b4e-4828-8989-e2583017cda7" //nolint
	setRepoMetaBFID           = "1857aa6a-b334-4b6a-bf7c-959d5581e8d4" //nolint
	voteKeyBFID               = "82d84700-7a9a-44f5-865d-f34badb00852" //nolint
	voteReleaseBFID           = "c8fc037a-dc5c-4fe3-b2fd-f8c602986d72" //nolint
	voteThresholdBFID         = "f0b640b6-4280-4cf0-afca-0d62046cee09" //nolint
	grantRoleBFID             = "17ec42d7-9f19-407c-8131-3033f7dcc142" //nolint
	initBFID                  = "5336e4c2-fc5c-49bd-b41d-9990dde03982" //nolint
	linkNameToIDBFID          = "8fc893ff-08e1-4cda-9264-62f6467d91a8" //nolint
	overrideNameToIDBFID      = "0455fbcd-4d1e-45ec-b0ce-5eaf73169b3e" //nolint
	renounceRoleBFID          = "08c8a75f-e9d2-4e9d-82e9-8f6c5b2bf8a0" //nolint
	revokeRoleBFID            = "d4040355-b755-4a1a-9f16-0f0462bd56d1" //nolint
)

type Transactor struct {
	eth    *ethclient.Client
	meta   gasless.Transactor
	signer gasless.Signer

	valistBuilder   *gasless.MessageBuilder
	registryBuilder *gasless.MessageBuilder
}

func NewTransactor(meta gasless.Transactor, signer gasless.Signer, eth *ethclient.Client, cfg *config.Config) (client.TransactorAPI, error) {
	valistBuilder, err := gasless.NewMessageBuilder(valist.ValistABI, cfg.Ethereum.Contracts["valist"], eth)
	if err != nil {
		return nil, err
	}

	registryBuilder, err := gasless.NewMessageBuilder(registry.ValistRegistryABI, cfg.Ethereum.Contracts["registry"], eth)
	if err != nil {
		return nil, err
	}

	return &Transactor{eth, meta, signer, valistBuilder, registryBuilder}, nil
}

// TransactOpts returns transaction options for a meta transcation.
func TransactOpts(account accounts.Account, wallet accounts.Wallet, chainID *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   account.Address,
		NoSend: true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != account.Address {
				return nil, bind.ErrNotAuthorized
			}

			return tx, nil
		},
	}
}
