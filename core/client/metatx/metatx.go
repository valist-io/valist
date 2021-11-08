// Package metatx defines a Transactor that uses meta transactions to pay gas fees on behalf of a user.
package metatx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/valist-io/gasless"
	"github.com/valist-io/gasless/mexa"

	"github.com/valist-io/valist/contract/registry"
	"github.com/valist-io/valist/contract/valist"
	"github.com/valist-io/valist/core/client"
)

type Transactor struct {
	valist   gasless.Transactor
	registry gasless.Transactor
}

func NewTransactor(eth *ethclient.Client, valistAddress, registryAddress common.Address, biconomyApiKey string) (client.TransactorAPI, error) {
	valist, err := mexa.NewTransactor(eth, valistAddress, valist.ValistABI, biconomyApiKey)
	if err != nil {
		return nil, err
	}

	registry, err := mexa.NewTransactor(eth, registryAddress, registry.ValistRegistryABI, biconomyApiKey)
	if err != nil {
		return nil, err
	}

	return &Transactor{
		valist:   valist,
		registry: registry,
	}, nil
}

func (t *Transactor) CreateOrganizationTx(txopts *gasless.TransactOpts, metaCID string) (*types.Transaction, error) {
	return t.valist.Transact(txopts, "createOrganization", metaCID)
}

func (t *Transactor) SetOrganizationMetaTx(txopts *gasless.TransactOpts, orgID common.Hash, metaCID string) (*types.Transaction, error) {
	return t.valist.Transact(txopts, "setOrgMeta", orgID, metaCID)
}

func (t *Transactor) LinkOrganizationNameTx(txopts *gasless.TransactOpts, orgID common.Hash, name string) (*types.Transaction, error) {
	return t.registry.Transact(txopts, "linkNameToID", orgID, name)
}

func (t *Transactor) CreateRepositoryTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	return t.valist.Transact(txopts, "createRepository", orgID, repoName, repoMeta)
}

func (t *Transactor) VoteReleaseTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName, tag, releaseCID, metaCID string) (*types.Transaction, error) {
	return t.valist.Transact(txopts, "voteRelease", orgID, repoName, tag, releaseCID, metaCID)
}

func (t *Transactor) VoteKeyTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName string, operation common.Hash, address common.Address) (*types.Transaction, error) {
	return t.valist.Transact(txopts, "voteKey", orgID, repoName, operation, address)
}

func (t *Transactor) SetRepositoryMetaTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	return t.valist.Transact(txopts, "setRepoMeta", orgID, repoName, repoMeta)
}

func (t *Transactor) VoteRepositoryThresholdTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName string, threshold *big.Int) (*types.Transaction, error) {
	return t.valist.Transact(txopts, "voteThreshold", orgID, repoName, threshold)
}

func (t *Transactor) VoteOrganizationThresholdTx(txopts *gasless.TransactOpts, orgID common.Hash, threshold *big.Int) (*types.Transaction, error) {
	return t.valist.Transact(txopts, "voteThreshold", orgID, "", threshold)
}
