package metatx

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"
	"github.com/valist-io/gasless"

	"github.com/valist-io/registry/internal/core"
)

func (t *Transactor) CreateOrganizationTx(ctx context.Context, txopts *bind.TransactOpts, metaCID cid.Cid) (*types.Transaction, error) {
	txopts = gasless.TransactOpts(txopts)

	tx, err := t.base.CreateOrganizationTx(ctx, txopts, metaCID)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(ctx, tx, t.signer, createOrganizationBFID)
}

func (t *Transactor) LinkOrganizationNameTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, name string) (*types.Transaction, error) {
	txopts = gasless.TransactOpts(txopts)

	tx, err := t.base.LinkOrganizationNameTx(ctx, txopts, orgID, name)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(ctx, tx, t.signer, linkNameToIDBFID)
}

func (t *Transactor) CreateRepositoryTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	txopts = gasless.TransactOpts(txopts)

	tx, err := t.base.CreateRepositoryTx(ctx, txopts, orgID, repoName, repoMeta)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(ctx, tx, t.signer, createRepositoryBFID)
}

func (t *Transactor) VoteReleaseTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, repoName string, release *core.Release) (*types.Transaction, error) {
	txopts = gasless.TransactOpts(txopts)

	tx, err := t.base.VoteReleaseTx(ctx, txopts, orgID, repoName, release)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(ctx, tx, t.signer, voteReleaseBFID)
}

func (t *Transactor) SetRepositoryMetaTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	txopts = gasless.TransactOpts(txopts)

	tx, err := t.base.SetRepositoryMetaTx(ctx, txopts, orgID, repoName, repoMeta)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(ctx, tx, t.signer, setRepoMetaBFID)
}

func (t *Transactor) VoteRepositoryThresholdTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, repoName string, threshold *big.Int) (*types.Transaction, error) {
	txopts = gasless.TransactOpts(txopts)

	tx, err := t.base.VoteRepositoryThresholdTx(ctx, txopts, orgID, repoName, threshold)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(ctx, tx, t.signer, voteThresholdBFID)
}
