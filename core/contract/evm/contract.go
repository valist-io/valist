package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/valist-io/valist/core/contract/evm/valist"
	"github.com/valist-io/valist/log"
)

//go:generate abigen -abi valist/valist.abi -bin valist/valist.bin -out valist/valist.go -pkg valist -type Valist

var logger = log.New()

type Backend interface {
	bind.DeployBackend
	bind.ContractBackend
}

type Contract struct {
	contract *valist.Valist
	backend  bind.DeployBackend
	accounts *AccountManager
}

// NewContract returns a contract using the given backend to make calls and transactions.
// The accounts manager is used to set the sender and sign transactions.
func NewContract(address common.Address, backend Backend, accounts *AccountManager) (*Contract, error) {
	contract, err := valist.NewValist(address, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{contract, backend, accounts}, nil
}

// CreateTeam creates a new team with the given members.
func (c *Contract) CreateTeam(ctx context.Context, teamName, metaURI string, members []string) error {
	var addresses []common.Address
	for _, hex := range members {
		addresses = append(addresses, common.HexToAddress(hex))
	}
	tx, err := c.contract.CreateTeam(c.txopts(ctx), teamName, metaURI, addresses)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// CreateProject creates a new project. Requires the sender to be a member of the team.
func (c *Contract) CreateProject(ctx context.Context, teamName, projectName, metaURI string, members []string) error {
	var addresses []common.Address
	for _, hex := range members {
		addresses = append(addresses, common.HexToAddress(hex))
	}
	tx, err := c.contract.CreateProject(c.txopts(ctx), teamName, projectName, metaURI, addresses)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// CreateRelease creates a new release. Requires the sender to be a member of the project.
func (c *Contract) CreateRelease(ctx context.Context, teamName, projectName, releaseName, metaURI string) error {
	tx, err := c.contract.CreateRelease(c.txopts(ctx), teamName, projectName, releaseName, metaURI)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// AddTeamMember adds a member to the team. Requires the sender to be a member of the team.
func (c *Contract) AddTeamMember(ctx context.Context, teamName string, address string) error {
	tx, err := c.contract.AddTeamMember(c.txopts(ctx), teamName, common.HexToAddress(address))
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// RemoveTeamMember removes a member from the team. Requires the sender to be a member of the team.
func (c *Contract) RemoveTeamMember(ctx context.Context, teamName string, address string) error {
	tx, err := c.contract.RemoveTeamMember(c.txopts(ctx), teamName, common.HexToAddress(address))
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// AddProjectMemeber adds a member to the project. Requires the sender to be a member of the team.
func (c *Contract) AddProjectMember(ctx context.Context, teamName, projectName string, address string) error {
	tx, err := c.contract.AddProjectMember(c.txopts(ctx), teamName, projectName, common.HexToAddress(address))
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// RemoveProjectMember removes a member from the project. Requires the sender to be a member of the team.
func (c *Contract) RemoveProjectMember(ctx context.Context, teamName, projectName string, address string) error {
	tx, err := c.contract.RemoveProjectMember(c.txopts(ctx), teamName, projectName, common.HexToAddress(address))
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// SetTeamMetaURI sets the team metadata content ID. Requires the sender to be a member of the team.
func (c *Contract) SetTeamMetaURI(ctx context.Context, teamName, metaURI string) error {
	tx, err := c.contract.SetTeamMetaCID(c.txopts(ctx), teamName, metaURI)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// SetProjectMetaURI sets the project metadata content ID. Requires the sender to be a member of the team.
func (c *Contract) SetProjectMetaURI(ctx context.Context, teamName, projectName, metaURI string) error {
	tx, err := c.contract.SetProjectMetaCID(c.txopts(ctx), teamName, projectName, metaURI)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// ApproveRelease approves the release by adding the sender's address to the approvers list.
// The sender's address will be removed from the rejectors list if it exists.
func (c *Contract) ApproveRelease(ctx context.Context, teamName, projectName, releaseName string) error {
	tx, err := c.contract.ApproveRelease(c.txopts(ctx), teamName, projectName, releaseName)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// RejectRelease rejects the release by adding the sender's address to the rejectors list.
// The sender's address will be removed from the approvers list if it exists.
func (c *Contract) RejectRelease(ctx context.Context, teamName, projectName, releaseName string) error {
	tx, err := c.contract.RejectRelease(c.txopts(ctx), teamName, projectName, releaseName)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// GetLatestReleaseName returns the latest release name.
func (c *Contract) GetLatestReleaseName(ctx context.Context, teamName, projectName string) (string, error) {
	return c.contract.GetLatestReleaseName(c.callopts(ctx), teamName, projectName)
}

// GetTeamMetaURI returns the team metadata URI.
func (c *Contract) GetTeamMetaURI(ctx context.Context, teamName string) (string, error) {
	return c.contract.GetTeamMetaCID(c.callopts(ctx), teamName)
}

// GetProjectMetaURI returns the project metadata URI.
func (c *Contract) GetProjectMetaURI(ctx context.Context, teamName, projectName string) (string, error) {
	return c.contract.GetProjectMetaCID(c.callopts(ctx), teamName, projectName)
}

// GetReleaseMetaURI returns the release metadata URI.
func (c *Contract) GetReleaseMetaURI(ctx context.Context, teamName, projectName, releaseName string) (string, error) {
	return c.contract.GetReleaseMetaCID(c.callopts(ctx), teamName, projectName, releaseName)
}

// GetTeamNames returns a paginated list of team names.
func (c *Contract) GetTeamNames(ctx context.Context, page *big.Int, size *big.Int) ([]string, error) {
	return c.contract.GetTeamNames(c.callopts(ctx), page, size)
}

// GetProjectNames returns a paginated list of project names.
func (c *Contract) GetProjectNames(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.contract.GetProjectNames(c.callopts(ctx), teamName, page, size)
}

// GetReleaseNames returns a paginated list of release names.
func (c *Contract) GetReleaseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.contract.GetReleaseNames(c.callopts(ctx), teamName, projectName, page, size)
}

// GetTeamMembers returns a paginated list of team members.
func (c *Contract) GetTeamMembers(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.contract.GetTeamMembers(c.callopts(ctx), teamName, page, size)
	if err != nil {
		return nil, err
	}
	var members []string
	for _, address := range addresses {
		members = append(members, address.Hex())
	}
	return members, nil
}

// GetProjectMembers returns a paginated list of project members.
func (c *Contract) GetProjectMembers(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.contract.GetProjectMembers(c.callopts(ctx), teamName, projectName, page, size)
	if err != nil {
		return nil, err
	}
	var members []string
	for _, address := range addresses {
		members = append(members, address.Hex())
	}
	return members, nil
}

// GetReleaseApprovers returns a paginated list of release approvers.
func (c *Contract) GetReleaseApprovers(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.contract.GetReleaseApprovers(c.callopts(ctx), teamName, projectName, releaseName, page, size)
	if err != nil {
		return nil, err
	}
	var approvers []string
	for _, address := range addresses {
		approvers = append(approvers, address.Hex())
	}
	return approvers, nil
}

// GetReleaseRejectors returns a paginated list of release rejectors.
func (c *Contract) GetReleaseRejectors(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.contract.GetReleaseRejectors(c.callopts(ctx), teamName, projectName, releaseName, page, size)
	if err != nil {
		return nil, err
	}
	var rejectors []string
	for _, address := range addresses {
		rejectors = append(rejectors, address.Hex())
	}
	return rejectors, nil
}

// txopts returns options for executing transactions.
func (c *Contract) txopts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		Context: ctx,
		From: c.accounts.account.Address,
		Signer: c.accounts.signTx,
	}
}

// callopts returns options for executing calls.
func (c *Contract) callopts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{
		Context: ctx,
		From: c.accounts.account.Address,
	}
}

// waitMined waits for a transaction to be mined.
func (c *Contract) waitMined(ctx context.Context, tx *types.Transaction) error {
	if sim, ok := c.backend.(*backends.SimulatedBackend); ok {
		sim.Commit()
	}

	logger.Info("Waiting for transaction: %s", tx.Hash().Hex())
	logger.Info("https://mumbai.polygonscan.com/tx/%s", tx.Hash().Hex())

	_, err := bind.WaitMined(ctx, c.backend, tx)
	return err
}
