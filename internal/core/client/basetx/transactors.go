package basetx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (t *Transactor) CreateOrganizationTx(txopts *bind.TransactOpts, metaCID string) (*types.Transaction, error) {
	return t.valist.CreateOrganization(txopts, metaCID)
}

func (t *Transactor) LinkOrganizationNameTx(txopts *bind.TransactOpts, orgID common.Hash, name string) (*types.Transaction, error) {
	return t.registry.LinkNameToID(txopts, orgID, name)
}

func (t *Transactor) CreateRepositoryTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	return t.valist.CreateRepository(txopts, orgID, repoName, repoMeta)
}

func (t *Transactor) VoteReleaseTx(txopts *bind.TransactOpts, orgID common.Hash, repoName, tag, releaseCID, metaCID string) (*types.Transaction, error) {
	return t.valist.VoteRelease(txopts, orgID, repoName, tag, releaseCID, metaCID)
}

func (t *Transactor) SetRepositoryMetaTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	return t.valist.SetRepoMeta(txopts, orgID, repoName, repoMeta)
}

func (t *Transactor) VoteRepositoryThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, threshold *big.Int) (*types.Transaction, error) {
	return t.valist.VoteThreshold(txopts, orgID, repoName, threshold)
}

func (t *Transactor) VoteOrganizationThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, threshold *big.Int) (*types.Transaction, error) {
	return t.valist.VoteThreshold(txopts, orgID, "", threshold)
}
