package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-cid"

	"github.com/valist-io/registry/internal/core"
)

// GetRepository returns the repository with the given orgID and name.
func (client *Client) GetRepository(ctx context.Context, orgID common.Hash, repoName string) (*core.Repository, error) {
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
		return nil, core.ErrRepositoryNotExist
	}

	metaCID, err := cid.Decode(repo.MetaCID)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse organization meta CID: %v", err)
	}

	return &core.Repository{
		OrgID:         orgID,
		Threshold:     repo.Threshold,
		ThresholdDate: repo.ThresholdDate,
		MetaCID:       metaCID,
	}, nil
}

// GetRepositoryMeta returns the repository meta with the given CID.
func (client *Client) GetRepositoryMeta(ctx context.Context, id cid.Cid) (*core.RepositoryMeta, error) {
	data, err := client.GetFile(ctx, id)
	if err != nil {
		return nil, err
	}

	var meta core.RepositoryMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

// CreateRepository creates a repository in the organization with the given orgID.
func (client *Client) CreateRepository(ctx context.Context, orgID common.Hash, name string, meta *core.RepositoryMeta) (<-chan core.CreateRepoResult, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	metaCID, err := client.AddFile(ctx, data)
	if err != nil {
		return nil, err
	}

	txopts := bind.TransactOpts{
		Context: ctx,
		From:    client.account.Address,
		Signer:  client.Signer,
	}

	tx, err := client.valist.CreateRepository(&txopts, orgID, name, metaCID.String())
	if err != nil {
		return nil, err
	}

	result := make(chan core.CreateRepoResult, 1)
	go client.createRepository(ctx, tx, result)

	return result, nil
}

func (client *Client) createRepository(ctx context.Context, tx *types.Transaction, result chan<- core.CreateRepoResult) {
	defer close(result)

	receipt, err := bind.WaitMined(ctx, client.eth, tx)
	if err != nil {
		result <- core.CreateRepoResult{Err: err}
		return
	}

	if len(receipt.Logs) == 0 {
		result <- core.CreateRepoResult{Err: fmt.Errorf("Failed to parse log")}
		return
	}

	log, err := client.valist.ParseRepoCreated(*receipt.Logs[0])
	if err != nil {
		result <- core.CreateRepoResult{Err: err}
		return
	}

	metaCID, err := cid.Decode(log.MetaCID)
	if err != nil {
		result <- core.CreateRepoResult{Err: err}
		return
	}

	result <- core.CreateRepoResult{
		OrgID:        log.OrgID,
		RepoNameHash: log.RepoNameHash,
		RepoName:     log.RepoName,
		MetaCIDHash:  log.MetaCIDHash,
		MetaCID:      metaCID,
	}
}
