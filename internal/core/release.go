package core

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-cid"
)

type ReleaseAPI interface {
	GetRelease(context.Context, common.Hash, string, string) (*Release, error)
	GetLatestRelease(context.Context, common.Hash, string) (*Release, error)
	VoteRelease(context.Context, common.Hash, string, *Release) (<-chan VoteReleaseResult, error)
	ListReleaseTags(common.Hash, string, *big.Int, *big.Int) ReleaseTagIterator
	ListReleases(common.Hash, string, *big.Int, *big.Int) ReleaseIterator
}

type Release struct {
	Tag        string
	ReleaseCID cid.Cid
	MetaCID    cid.Cid
	Signers    []common.Address
}

type VoteReleaseResult struct {
	OrgID      common.Hash
	RepoName   common.Hash
	Tag        common.Hash
	ReleaseCID cid.Cid
	MetaCID    cid.Cid
	Signer     common.Address
	SigCount   *big.Int
	Threshold  *big.Int
	Err        error
}

type ReleaseTagIterator interface {
	Next(context.Context) (string, error)
}

type ReleaseIterator interface {
	Next(context.Context) (*Release, error)
	ForEach(context.Context, func(*Release)) error
}
