package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"

	"github.com/valist-io/registry/internal/contract/valist"
)

type RepositoryMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectType string `json:"projectType"`
	Homepage    string `json:"homepage"`
	Repository  string `json:"repository"`
}

type Repository struct {
	OrgID         common.Hash
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       cid.Cid
}

type CreateRepoResult struct {
	Log *valist.ValistRepoCreated
	Err error
}

// GetRepository returns the repository with the given orgID and name.
func (client *Client) GetRepository(ctx context.Context, orgID common.Hash, repoName string) (*Repository, error) {
	callopts := bind.CallOpts{Context: ctx}
	selector := crypto.Keccak256Hash(orgID[:], []byte(repoName))

	repo, err := client.valist.Repos(&callopts, selector)
	if err != nil {
		return nil, err
	}

	if !repo.Exists {
		return nil, ErrRepositoryNotExist
	}

	metaCID, err := cid.Decode(repo.MetaCID)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse organization meta CID: %v", err)
	}

	return &Repository{
		OrgID:         orgID,
		Threshold:     repo.Threshold,
		ThresholdDate: repo.ThresholdDate,
		MetaCID:       metaCID,
	}, nil
}

// GetRepositoryMeta returns the repository meta with the given CID.
func (client *Client) GetRepositoryMeta(ctx context.Context, id cid.Cid) (*RepositoryMeta, error) {
	node, err := client.ipfs.Unixfs().Get(ctx, path.IpfsPath(id))
	if err != nil {
		return nil, err
	}

	file, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("Failed to parse repository meta")
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var meta RepositoryMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

// CreateRepository creates a repository in the organization with the given orgID.
func (client *Client) CreateRepository(ctx context.Context, orgID common.Hash, name string, meta *RepositoryMeta) (<-chan CreateRepoResult, error) {
	data, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	path, err := client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data))
	if err != nil {
		return nil, err
	}

	txopts, err := bind.NewKeyedTransactorWithChainID(client.private, client.chainID)
	if err != nil {
		return nil, err
	}
	txopts.Context = ctx

	tx, err := client.valist.CreateRepository(txopts, orgID, name, path.Cid().String())
	if err != nil {
		return nil, err
	}

	result := make(chan CreateRepoResult, 1)
	go client.createRepository(ctx, tx, result)

	return result, nil
}

func (client *Client) createRepository(ctx context.Context, tx *types.Transaction, result chan<- CreateRepoResult) {
	defer close(result)

	receipt, err := bind.WaitMined(ctx, client.eth, tx)
	if err != nil {
		result <- CreateRepoResult{Err: err}
		return
	}

	if len(receipt.Logs) == 0 {
		result <- CreateRepoResult{Err: fmt.Errorf("Failed to parse log")}
		return
	}

	log, err := client.valist.ParseRepoCreated(*receipt.Logs[0])
	if err != nil {
		result <- CreateRepoResult{Err: err}
		return
	}

	result <- CreateRepoResult{Log: log}
}
