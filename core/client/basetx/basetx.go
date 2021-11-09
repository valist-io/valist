package basetx

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/valist-io/gasless"
	"github.com/valist-io/valist/contract"
	"github.com/valist-io/valist/contract/registry"
	"github.com/valist-io/valist/contract/valist"
)

type Transactor struct {
	valist   *valist.Valist
	registry *registry.ValistRegistry
}

func NewTransactor(eth bind.ContractBackend, valistAddress, registryAddress common.Address) (*Transactor, error) {
	valist, err := contract.NewValist(valistAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize valist contract: %v", err)
	}

	registry, err := contract.NewRegistry(registryAddress, eth)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize registry contract: %v", err)
	}

	return &Transactor{
		valist:   valist,
		registry: registry,
	}, nil
}

func (t *Transactor) CreateOrganizationTx(txopts *gasless.TransactOpts, metaCID string) (*types.Transaction, error) {
	return t.valist.CreateOrganization(&txopts.TransactOpts, metaCID)
}

func (t *Transactor) SetOrganizationMetaTx(txopts *gasless.TransactOpts, orgID common.Hash, metaCID string) (*types.Transaction, error) {
	return t.valist.SetOrgMeta(&txopts.TransactOpts, orgID, metaCID)
}

func (t *Transactor) LinkOrganizationNameTx(txopts *gasless.TransactOpts, orgID common.Hash, name string) (*types.Transaction, error) {
	return t.registry.LinkNameToID(&txopts.TransactOpts, orgID, name)
}

func (t *Transactor) CreateRepositoryTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	return t.valist.CreateRepository(&txopts.TransactOpts, orgID, repoName, repoMeta)
}

func (t *Transactor) VoteKeyTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName string, operation common.Hash, address common.Address) (*types.Transaction, error) {
	return t.valist.VoteKey(&txopts.TransactOpts, orgID, repoName, operation, address)
}

func (t *Transactor) VoteReleaseTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName, tag, releaseCID, metaCID string) (*types.Transaction, error) {
	return t.valist.VoteRelease(&txopts.TransactOpts, orgID, repoName, tag, releaseCID, metaCID)
}

func (t *Transactor) SetRepositoryMetaTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName string, repoMeta string) (*types.Transaction, error) {
	return t.valist.SetRepoMeta(&txopts.TransactOpts, orgID, repoName, repoMeta)
}

func (t *Transactor) VoteRepositoryThresholdTx(txopts *gasless.TransactOpts, orgID common.Hash, repoName string, threshold *big.Int) (*types.Transaction, error) {
	return t.valist.VoteThreshold(&txopts.TransactOpts, orgID, repoName, threshold)
}

func (t *Transactor) VoteOrganizationThresholdTx(txopts *gasless.TransactOpts, orgID common.Hash, threshold *big.Int) (*types.Transaction, error) {
	return t.valist.VoteThreshold(&txopts.TransactOpts, orgID, "", threshold)
}
