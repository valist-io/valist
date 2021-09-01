package basetx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/registry/internal/core/types"
)

func (t *Transactor) CreateOrganizationTx(txopts *bind.TransactOpts, metaCID cid.Cid) (*ethtypes.Transaction, error) {
	return t.valist.CreateOrganization(txopts, metaCID.String())
}

func (t *Transactor) SetOrganizationMetaTx(txopts *bind.TransactOpts, orgID common.Hash, metaCID cid.Cid) (*ethtypes.Transaction, error) {
	return t.valist.SetOrgMeta(txopts, orgID, metaCID.String())
}

func (t *Transactor) LinkOrganizationNameTx(txopts *bind.TransactOpts, orgID common.Hash, name string) (*ethtypes.Transaction, error) {
	return t.registry.LinkNameToID(txopts, orgID, name)
}

func (t *Transactor) CreateRepositoryTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*ethtypes.Transaction, error) {
	return t.valist.CreateRepository(txopts, orgID, repoName, repoMeta)
}

func (t *Transactor) VoteKeyTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, operation common.Hash, address common.Address) (*ethtypes.Transaction, error) {
	return t.valist.VoteKey(txopts, orgID, repoName, operation, address)
}

func (t *Transactor) VoteReleaseTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, release *types.Release) (*ethtypes.Transaction, error) {
	return t.valist.VoteRelease(txopts, orgID, repoName, release.Tag, release.ReleaseCID.String(), release.MetaCID.String())
}

func (t *Transactor) SetRepositoryMetaTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*ethtypes.Transaction, error) {
	return t.valist.SetRepoMeta(txopts, orgID, repoName, repoMeta)
}

func (t *Transactor) VoteRepositoryThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, threshold *big.Int) (*ethtypes.Transaction, error) {
	return t.valist.VoteThreshold(txopts, orgID, repoName, threshold)
}

func (t *Transactor) VoteOrganizationThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, threshold *big.Int) (*ethtypes.Transaction, error) {
	return t.valist.VoteThreshold(txopts, orgID, "", threshold)
}
