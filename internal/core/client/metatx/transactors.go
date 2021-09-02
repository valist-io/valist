package metatx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (t *Transactor) CreateOrganizationTx(txopts *bind.TransactOpts, metaCID string) (*types.Transaction, error) {
	msg, err := t.valistBuilder.Message(txopts.Context, txopts.From, "createOrganization", metaCID)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, createOrganizationBFID)
}

func (t *Transactor) SetOrganizationMetaTx(txopts *bind.TransactOpts, orgID common.Hash, metaCID string) (*types.Transaction, error) {
	msg, err := t.valistBuilder.Message(txopts.Context, txopts.From, "setOrgMeta", orgID, metaCID)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, setOrgMetaBFID)
}

func (t *Transactor) LinkOrganizationNameTx(txopts *bind.TransactOpts, orgID common.Hash, name string) (*types.Transaction, error) {
	msg, err := t.registryBuilder.Message(txopts.Context, txopts.From, "linkNameToID", orgID, name)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, linkNameToIDBFID)
}

func (t *Transactor) CreateRepositoryTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	msg, err := t.valistBuilder.Message(txopts.Context, txopts.From, "createRepository", orgID, repoName, repoMeta)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, createRepositoryBFID)
}

func (t *Transactor) VoteReleaseTx(txopts *bind.TransactOpts, orgID common.Hash, repoName, tag, releaseCID, metaCID string) (*types.Transaction, error) {
	msg, err := t.valistBuilder.Message(txopts.Context, txopts.From, "voteRelease", orgID, repoName, tag, releaseCID, metaCID)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, voteReleaseBFID)
}

func (t *Transactor) VoteKeyTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, operation common.Hash, address common.Address) (*types.Transaction, error) {
	msg, err := t.valistBuilder.Message(txopts.Context, txopts.From, "voteKey", orgID, repoName, operation, address)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, voteKeyBFID)
}

func (t *Transactor) SetRepositoryMetaTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	msg, err := t.valistBuilder.Message(txopts.Context, txopts.From, "setRepoMeta", orgID, repoName, repoMeta)
	if err != nil {
		return nil, err
	}
	return t.meta.Transact(txopts.Context, msg, t.signer, setRepoMetaBFID)
}

func (t *Transactor) VoteRepositoryThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, repoName string, threshold *big.Int) (*types.Transaction, error) {
	msg, err := t.valistBuilder.Message(txopts.Context, txopts.From, "voteThreshold", orgID, repoName, threshold)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, voteThresholdBFID)
}

func (t *Transactor) VoteOrganizationThresholdTx(txopts *bind.TransactOpts, orgID common.Hash, threshold *big.Int) (*types.Transaction, error) {
	msg, err := t.valistBuilder.Message(txopts.Context, txopts.From, "voteThreshold", orgID, "", threshold)
	if err != nil {
		return nil, err
	}

	return t.meta.Transact(txopts.Context, msg, t.signer, voteThresholdBFID)
}
