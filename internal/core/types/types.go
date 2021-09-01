package types

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cid "github.com/ipfs/go-cid"
	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
)

const (
	ProjectTypeBinary = "binary"
	ProjectTypeNode   = "node"
	ProjectTypeNPM    = "npm"
	ProjectTypeGo     = "go"
	ProjectTypeRust   = "rust"
	ProjectTypePython = "python"
	ProjectTypeDocker = "docker"
	ProjectTypeCPP    = "c++"
	ProjectTypeStatic = "static"
)

var ProjectTypes = []string{
	ProjectTypeBinary,
	ProjectTypeNode,
	ProjectTypeNPM,
	ProjectTypeGo,
	ProjectTypeRust,
	ProjectTypePython,
	ProjectTypeDocker,
	ProjectTypeCPP,
	ProjectTypeStatic,
}

var (
	ErrOrganizationNotExist = errors.New("Organization does not exist")
	ErrRepositoryNotExist   = errors.New("Repository does not exist")
	ErrReleaseNotExist      = errors.New("Release does not exist")
)

// CoreAPI defines the high-level interface for Valist.
type CoreAPI interface {
	OrganizationAPI
	RegistryAPI
	ReleaseAPI
	RepositoryAPI
	StorageAPI
	SwitchAccount(accounts.Account, accounts.Wallet)
	Close()
}

// TransactorAPI defines functions to abstract blockchain transactions.
// TODO: Maybe this can return []*types.Log instead of *types.Transaction and handle waiting and log parsing?
type TransactorAPI interface {
	CreateOrganizationTx(*bind.TransactOpts, cid.Cid) (*types.Transaction, error)
	SetOrganizationMetaTx(*bind.TransactOpts, common.Hash, cid.Cid) (*types.Transaction, error)
	LinkOrganizationNameTx(*bind.TransactOpts, common.Hash, string) (*types.Transaction, error)
	CreateRepositoryTx(*bind.TransactOpts, common.Hash, string, string) (*types.Transaction, error)
	VoteReleaseTx(*bind.TransactOpts, common.Hash, string, *Release) (*types.Transaction, error)
	VoteKeyTx(*bind.TransactOpts, common.Hash, string, common.Hash, common.Address) (*types.Transaction, error)
	SetRepositoryMetaTx(*bind.TransactOpts, common.Hash, string, string) (*types.Transaction, error)
	VoteOrganizationThresholdTx(*bind.TransactOpts, common.Hash, *big.Int) (*types.Transaction, error)
	VoteRepositoryThresholdTx(*bind.TransactOpts, common.Hash, string, *big.Int) (*types.Transaction, error)
}

type OrganizationAPI interface {
	GetOrganization(context.Context, common.Hash) (*Organization, error)
	GetOrganizationMeta(context.Context, cid.Cid) (*OrganizationMeta, error)
	CreateOrganization(context.Context, *OrganizationMeta) (*valist.ValistOrgCreated, error)
	VoteOrganizationAdmin(context.Context, common.Hash, common.Hash, common.Address) (*valist.ValistVoteKeyEvent, error)
	VoteOrganizationThreshold(context.Context, common.Hash, *big.Int) (*valist.ValistVoteThresholdEvent, error)
}

type RegistryAPI interface {
	GetOrganizationID(context.Context, string) (common.Hash, error)
	LinkOrganizationName(context.Context, common.Hash, string) (*registry.ValistRegistryMappingEvent, error)
}

type ReleaseAPI interface {
	GetRelease(context.Context, common.Hash, string, string) (*Release, error)
	GetLatestRelease(context.Context, common.Hash, string) (*Release, error)
	ListReleaseTags(common.Hash, string, *big.Int, *big.Int) ReleaseTagIterator
	ListReleases(common.Hash, string, *big.Int, *big.Int) ReleaseIterator
	VoteRelease(context.Context, common.Hash, string, *Release) (*valist.ValistVoteReleaseEvent, error)
}

type RepositoryAPI interface {
	GetRepository(context.Context, common.Hash, string) (*Repository, error)
	GetRepositoryMeta(context.Context, cid.Cid) (*RepositoryMeta, error)
	CreateRepository(context.Context, common.Hash, string, *RepositoryMeta) (*valist.ValistRepoCreated, error)
	SetRepositoryMeta(context.Context, common.Hash, string, *RepositoryMeta) (*valist.ValistMetaUpdate, error)
	VoteRepoDev(context.Context, common.Hash, string, common.Hash, common.Address) (*valist.ValistVoteKeyEvent, error)
	VoteRepositoryThreshold(context.Context, common.Hash, string, *big.Int) (*valist.ValistVoteThresholdEvent, error)
}

type ReleaseTagIterator interface {
	Next(context.Context) (string, error)
}

type ReleaseIterator interface {
	Next(context.Context) (*Release, error)
	ForEach(context.Context, func(*Release)) error
}

type StorageAPI interface {
	// ReadFile returns the contents of the file with the given CID.
	ReadFile(context.Context, cid.Cid) ([]byte, error)
	// WriteFile writes the given file contents and returns its CID.
	WriteFile(context.Context, []byte) (cid.Cid, error)
	// WriteFilePath writes the contents of the given file path and returns its CID.
	WriteFilePath(context.Context, string) (cid.Cid, error)
	// WriteDirEntries writes the given list of files into a directory and returns its CID.
	WriteDirEntries(context.Context, string, []string) (cid.Cid, error)
}

type Organization struct {
	ID            common.Hash
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       cid.Cid
}

type OrganizationMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

type LinkOrgNameResult struct {
	OrgID common.Hash
	Name  string
	Err   error
}

type Release struct {
	Tag        string
	ReleaseCID cid.Cid
	MetaCID    cid.Cid
	Signers    []common.Address
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
