package core

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/valist-io/valist"
	"github.com/valist-io/valist/core/contract/evm"
	"github.com/valist-io/valist/core/storage/ipfs"
)

type Client struct {
	valist.AccountAPI
	valist.ContractAPI
	valist.StorageAPI
}

// NewClient creates a new Valist client using the given config.
func NewClient(ctx context.Context, config *Config) (valist.API, error) {
	eth, err := ethclient.Dial(config.EthereumRPC)
	if err != nil {
		return nil, err
	}
	storage, err := ipfs.NewStorage(ctx, config.StoragePath(), config.IpfsApiAddress, config.IpfsBootstrapPeers...)
	if err != nil {
		return nil, err
	}
	chainID, err := eth.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	accounts := evm.NewAccountManager(config.KeyStorePath(), chainID)
	address := common.HexToAddress(config.EthereumContracts["valist"])

	contract, err := evm.NewContract(address, eth, accounts)
	if err != nil {
		return nil, err
	}
	return &Client{
		AccountAPI:  accounts,
		ContractAPI: contract,
		StorageAPI:  storage,
	}, nil
}

// GetTeam returns the team with the given name.
func (c *Client) GetTeam(ctx context.Context, teamName string) (*valist.Team, error) {
	path, err := c.GetTeamMetaURI(ctx, teamName)
	if err != nil {
		return nil, err
	}
	data, err := c.ReadFile(ctx, path)
	if err != nil {
		return nil, err
	}
	var team valist.Team
	if err := json.Unmarshal(data, &team); err != nil {
		return nil, err
	}
	return &team, nil
}

// GetProject returns the project with the given name.
func (c *Client) GetProject(ctx context.Context, teamName, projectName string) (*valist.Project, error) {
	path, err := c.GetProjectMetaURI(ctx, teamName, projectName)
	if err != nil {
		return nil, err
	}
	data, err := c.ReadFile(ctx, path)
	if err != nil {
		return nil, err
	}
	var project valist.Project
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// GetRelease returns the release with the given name.
func (c *Client) GetRelease(ctx context.Context, teamName, projectName, releaseName string) (*valist.Release, error) {
	if releaseName == "latest" {
		return c.GetLatestRelease(ctx, teamName, projectName)
	}
	path, err := c.GetReleaseMetaURI(ctx, teamName, projectName, releaseName)
	if err != nil {
		return nil, err
	}
	data, err := c.ReadFile(ctx, path)
	if err != nil {
		return nil, err
	}
	var release valist.Release
	if err := json.Unmarshal(data, &release); err != nil {
		return nil, err
	}
	return &release, nil
}

// GetLatestRelease returns the latest release.
func (c *Client) GetLatestRelease(ctx context.Context, teamName, projectName string) (*valist.Release, error) {
	releaseName, err := c.GetLatestReleaseName(ctx, teamName, projectName)
	if err != nil {
		return nil, err
	}
	path, err := c.GetReleaseMetaURI(ctx, teamName, projectName, releaseName)
	if err != nil {
		return nil, err
	}
	data, err := c.ReadFile(ctx, path)
	if err != nil {
		return nil, err
	}
	var release valist.Release
	if err := json.Unmarshal(data, &release); err != nil {
		return nil, err
	}
	return &release, nil
}

// ResolvePath resolves the team, project, and release from the given path.
func (c *Client) ResolvePath(ctx context.Context, path string) (valist.ResolvedPath, error) {
	var res valist.ResolvedPath
	var err error
	// cleanup path
	clean := strings.TrimLeft(path, "/@")
	parts := strings.Split(clean, "/")
	if len(parts) == 0 || len(parts) > 3 {
		return res, fmt.Errorf("invalid path")
	}
	// resolve team
	res.TeamName = parts[0]
	res.Team, err = c.GetTeam(ctx, res.TeamName)
	if err != nil {
		return res, err
	}
	if len(parts) < 2 {
		return res, nil
	}
	// resolve project
	res.ProjectName = parts[1]
	res.Project, err = c.GetProject(ctx, res.TeamName, res.ProjectName)
	if err != nil {
		return res, err
	}
	if len(parts) < 3 {
		return res, nil
	}
	// resolve release
	res.ReleaseName = parts[2]
	res.Release, err = c.GetRelease(ctx, res.TeamName, res.ProjectName, res.ReleaseName)
	if err != nil {
		return res, err
	}
	return res, nil
}
