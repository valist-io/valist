package basetx

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/registry/internal/core"
)

func (t *Transactor) CreateOrganizationTx(ctx context.Context, txopts *bind.TransactOpts, metaCID cid.Cid) (*types.Transaction, error) {
	return t.valist.CreateOrganization(txopts, metaCID.String())
}

func (t *Transactor) LinkOrganizationNameTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, name string) (*types.Transaction, error) {
	return t.registry.LinkNameToID(txopts, orgID, name)
}

func (t *Transactor) CreateRepositoryTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	return t.valist.CreateRepository(txopts, orgID, repoName, repoMeta)
}

func (t *Transactor) VoteReleaseTx(ctx context.Context, txopts *bind.TransactOpts, orgID common.Hash, repoName string, release *core.Release) (*types.Transaction, error) {
	return t.valist.VoteRelease(txopts, orgID, repoName, release.Tag, release.ReleaseCID.String(), release.MetaCID.String())
}
