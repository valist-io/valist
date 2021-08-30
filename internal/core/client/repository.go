package client

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core/types"
)

// GetRepository returns the repository with the given orgID and name.
func (client *Client) GetRepository(ctx context.Context, orgID common.Hash, repoName string) (*types.Repository, error) {
	selector := crypto.Keccak256Hash(orgID[:], []byte(repoName))
	callopts := bind.CallOpts{
		Context: ctx,
		From:    client.account.Address,
	}

	repo, err := client.valist.Repos(&callopts, selector)
	if err != nil {
		return nil, err
	}

	if !repo.Exists {
		return nil, types.ErrRepositoryNotExist
	}

	return &types.Repository{
		OrgID:         orgID,
		Name:          repoName,
		Threshold:     repo.Threshold,
		ThresholdDate: repo.ThresholdDate,
		MetaCID:       repo.MetaCID,
	}, nil
}

// GetRepositoryMeta returns the repository meta with the given CID.
func (client *Client) GetRepositoryMeta(ctx context.Context, p string) (*types.RepositoryMeta, error) {
	data, err := client.storage.ReadFile(ctx, p)
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

	metaCID, err := client.storage.Write(ctx, data)
	if err != nil {
		return nil, err
	}

	txopts := client.transactOpts(client.account, client.wallet, client.chainID)
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

func (client *Client) SetRepositoryMeta(ctx context.Context, orgID common.Hash, name string, meta *types.RepositoryMeta) (*valist.ValistMetaUpdate, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	metaCID, err := client.storage.Write(ctx, data)
	if err != nil {
		return nil, err
	}

	txopts := client.transactOpts(client.account, client.wallet, client.chainID)
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

func (client *Client) VoteRepositoryThreshold(ctx context.Context, orgID common.Hash, name string, threshold *big.Int) (*valist.ValistVoteThresholdEvent, error) {
	txopts := client.transactOpts(client.account, client.wallet, client.chainID)
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
