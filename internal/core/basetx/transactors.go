package basetx

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"
)

func (t *Transactor) CreateOrganizationTx(ctx context.Context, txopts *bind.TransactOpts, metaCID cid.Cid) (*types.Transaction, error) {
	return t.valist.CreateOrganization(txopts, metaCID.String())
}

func (t *Transactor) LinkOrganizationNameTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, name string) (*types.Transaction, error) {
	return t.registry.LinkNameToID(txopts, orgID, name)
}

func (t *Transactor) CreateRepositoryTx(ctx context.Context, txopts *bind.TransactOpts, orgID [32]byte, repoName string, repoMeta string) (*types.Transaction, error) {
	return t.valist.CreateRepository(txopts, orgID, repoName, repoMeta)
}

func (t *Transactor) VoteReleaseTx(
	ctx context.Context,
	txopts *bind.TransactOpts,
	orgID [32]byte,
	repoName string,
	tag string,
	releaseCID string,
	metaCID string,
) (*types.Transaction, error) {
	return t.valist.VoteRelease(txopts, orgID, repoName, tag, releaseCID, metaCID)
}
