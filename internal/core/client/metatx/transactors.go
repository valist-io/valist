package metatx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/registry/internal/core/types"
)

func (t *Transactor) CreateOrganizationTx(txopts *bind.TransactOpts, metaCID cid.Cid) (*ethtypes.Transaction, error) {
	tx, err := t.base.CreateOrganizationTx(txopts, metaCID)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, tx, t.signer, createOrganizationBFID)
}

func (t *Transactor) LinkOrganizationNameTx(txopts *bind.TransactOpts, orgID common.Hash, name string) (*ethtypes.Transaction, error) {
	tx, err := t.base.LinkOrganizationNameTx(txopts, orgID, name)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, tx, t.signer, linkNameToIDBFID)
}

func (t *Transactor) CreateRepositoryTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*ethtypes.Transaction, error) {
	tx, err := t.base.CreateRepositoryTx(txopts, orgID, repoName, repoMeta)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, tx, t.signer, createRepositoryBFID)
}

func (t *Transactor) VoteReleaseTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, release *types.Release) (*ethtypes.Transaction, error) {
	tx, err := t.base.VoteReleaseTx(txopts, orgID, repoName, release)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, tx, t.signer, voteReleaseBFID)
}

func (t *Transactor) VoteKeyTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, operation common.Hash, address common.Address) (*ethtypes.Transaction, error) {
	tx, err := t.base.VoteKeyTx(txopts, orgID, repoName, operation, address)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, tx, t.signer, voteKeyBFID)
}

func (t *Transactor) SetRepositoryMetaTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*ethtypes.Transaction, error) {
	tx, err := t.base.SetRepositoryMetaTx(txopts, orgID, repoName, repoMeta)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, tx, t.signer, setRepoMetaBFID)
}

func (t *Transactor) VoteRepositoryThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, threshold *big.Int) (*ethtypes.Transaction, error) {
	tx, err := t.base.VoteRepositoryThresholdTx(txopts, orgID, repoName, threshold)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, tx, t.signer, voteThresholdBFID)
}

func (t *Transactor) VoteOrganizationThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, threshold *big.Int) (*ethtypes.Transaction, error) {
	tx, err := t.base.VoteOrganizationThresholdTx(txopts, orgID, threshold)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, tx, t.signer, voteThresholdBFID)

}
