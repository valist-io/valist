package core

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-cid"
)

type RepositoryAPI interface {
	GetRepository(context.Context, common.Hash, string) (*Repository, error)
	GetRepositoryMeta(context.Context, cid.Cid) (*RepositoryMeta, error)
	CreateRepository(context.Context, common.Hash, string, *RepositoryMeta) (<-chan CreateRepoResult, error)
}

type Repository struct {
	OrgID         common.Hash
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       cid.Cid
}

type RepositoryMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectType string `json:"projectType"`
	Homepage    string `json:"homepage"`
	Repository  string `json:"repository"`
}

type CreateRepoResult struct {
	OrgID        common.Hash
	RepoNameHash common.Hash
	RepoName     string
	MetaCIDHash  common.Hash
	MetaCID      cid.Cid
	Err          error
}
