package metatx

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"
	"github.com/valist-io/gasless"
)

func (t *Transactor) CreateOrganizationTx(ctx context.Context, txopts *bind.TransactOpts, metaCID cid.Cid) (*types.Transaction, error) {
	setMetaTransactOpts(txopts)

	tx, err := t.base.CreateOrganizationTx(ctx, txopts, metaCID)
	if err != nil {
		return nil, err
	}

	message, err := t.newMessage(ctx, tx, createOrganizationBFID)
	if err != nil {
		return nil, err
	}

	return gasless.SendTransaction(ctx, t.meta, message, t.account, t.wallet)
}

func (t *Transactor) LinkOrganizationNameTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, name string) (*types.Transaction, error) {
	setMetaTransactOpts(txopts)

	tx, err := t.base.LinkOrganizationNameTx(ctx, txopts, orgID, name)
	if err != nil {
		return nil, err
	}

	message, err := t.newMessage(ctx, tx, linkNameToIDBFID)
	if err != nil {
		return nil, err
	}

	return gasless.SendTransaction(ctx, t.meta, message, t.account, t.wallet)
}
