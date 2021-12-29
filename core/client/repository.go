package client

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/valist-io/valist/contract/valist"
	"github.com/valist-io/valist/core/types"
)

// GetRepository returns the repository with the given orgID and repoName.
func (client *Client) GetRepository(ctx context.Context, orgID common.Hash, repoName string) (*types.Repository, error) {
	selector := crypto.Keccak256Hash(orgID[:], []byte(repoName))
	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.signer.Account().Address,
	}

	repo, err := client.valist.Repos(&callopts, selector)
	if err != nil {
		return nil, err
	}

	if !repo.Exists {
		return nil, types.ErrRepoNotExist
	}

	return &types.Repository{
		OrgID:         orgID,
		Name:          repoName,
		Threshold:     repo.Threshold,
		ThresholdDate: repo.ThresholdDate,
		MetaCID:       repo.MetaCID,
	}, nil
}

// GetRepositoryMeta returns the repository meta from the given path.
func (client *Client) GetRepositoryMeta(ctx context.Context, p string) (*types.RepositoryMeta, error) {
	data, err := client.ReadFile(ctx, p)
	if err != nil {
		return nil, err
	}

	var meta types.RepositoryMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

// CreateRepository creates a repository in the organization with the given orgID.
func (client *Client) CreateRepository(ctx context.Context, orgID common.Hash, name string, meta *types.RepositoryMeta) (*valist.ValistRepoCreated, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	metaCID, err := client.WriteFile(ctx, data)
	if err != nil {
		return nil, err
	}

	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.CreateRepositoryTx(txopts, orgID, name, metaCID)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseRepoCreated(*logs[0])
}

// VoteRepoDev votes to add, revoke, or rotate a repository dev key.
func (client *Client) VoteRepoDev(ctx context.Context, orgID common.Hash, repoName string, operation common.Hash, address common.Address) (*valist.ValistVoteKeyEvent, error) {
	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.VoteKeyTx(txopts, orgID, repoName, operation, address)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseVoteKeyEvent(*logs[0])
}

// SetRepositoryMeta updates the metadata of the repository with the given orgID and repoName.
func (client *Client) SetRepositoryMeta(ctx context.Context, orgID common.Hash, name string, meta *types.RepositoryMeta) (*valist.ValistMetaUpdate, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	metaCID, err := client.WriteFile(ctx, data)
	if err != nil {
		return nil, err
	}

	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.SetRepositoryMetaTx(txopts, orgID, name, metaCID)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseMetaUpdate(*logs[0])
}

// VoteRepositoryThreshold votes to set a multifactor voting threshold for the repository.
func (client *Client) VoteRepositoryThreshold(ctx context.Context, orgID common.Hash, name string, threshold *big.Int) (*valist.ValistVoteThresholdEvent, error) {
	txopts := client.signer.NewTransactor()
	txopts.Context = ctx

	tx, err := client.transactor.VoteRepositoryThresholdTx(txopts, orgID, name, threshold)
	if err != nil {
		return nil, err
	}

	logs, err := waitMined(ctx, client.eth, tx)
	if err != nil {
		return nil, err
	}

	return client.valist.ParseVoteThresholdEvent(*logs[0])
}

// GetRepositoryMembers returns a list of all keys currently active in the organization.
func (client *Client) GetRepositoryMembers(ctx context.Context, orgID common.Hash, repoName string) ([]common.Address, error) {
	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.signer.Account().Address,
	}

	packed, err := packedEncoding(orgID, repoName, REPO_DEV)
	if err != nil {
		return nil, err
	}
	selector := crypto.Keccak256Hash(packed)

	return client.valist.GetRoleMembers(&callopts, selector)
}
