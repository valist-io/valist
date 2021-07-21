package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core/path"
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
	Meta          *RepositoryMeta
	MetaCID       cid.Cid
	Tags          []string
}

// GetRepository returns the repository with the given orgID and name.
func (client *Client) GetRepository(ctx context.Context, orgName string, repoName string) (*Repository, error) {
	callopts := bind.CallOpts{Context: ctx}

	orgID, err := client.GetOrganizationID(ctx, orgName)
	if err != nil {
		return nil, err
	}

	selector := crypto.Keccak256Hash(orgID[:], []byte(repoName))
	repo, err := client.valistContract.Repos(&callopts, selector)
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

	meta, err := client.GetRepositoryMeta(ctx, metaCID)
	if err != nil {
		return nil, err
	}

	tags, err := client.valistContract.GetReleaseTags(&callopts, selector, big.NewInt(1), big.NewInt(10))
	if err != nil {
		return nil, err
	}

	return &Repository{
		OrgID:         orgID,
		Threshold:     repo.Threshold,
		ThresholdDate: repo.ThresholdDate,
		Meta:          meta,
		MetaCID:       metaCID,
		Tags:          tags,
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
