package types

import (
	"context"
	"errors"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"

	"github.com/valist-io/valist/internal/contract/registry"
	"github.com/valist-io/valist/internal/contract/valist"
	"github.com/valist-io/valist/internal/signer"
	"github.com/valist-io/valist/internal/storage"
)

// DeprecationNotice contains a deprecation notice in plain text
const DeprecationNotice = "QmRBwMae3Skqzc1GmAKBdcnFFPnHeD585MwYtVZzfh9Tkh"

const (
	ProjectTypeBinary = "binary"
	ProjectTypeNode   = "node"
	ProjectTypeNPM    = "npm"
	ProjectTypeGit    = "git"
	ProjectTypeRust   = "crate"
	ProjectTypePython = "python"
	ProjectTypeDocker = "docker"
	ProjectTypeStatic = "static"
)

var (
	RegexShortname            = regexp.MustCompile(`^[0-9a-z-_]+$`)
	RegexPath                 = regexp.MustCompile(`^[0-9A-z\-_\/\.]+$`)
	RegexAcceptableCharacters = regexp.MustCompile(`^[0-9A-z-_\\/\. ]*$`)
)

var ProjectTypes = []string{
	ProjectTypeBinary,
	ProjectTypeNode,
	ProjectTypeNPM,
	ProjectTypeGit,
	ProjectTypeRust,
	ProjectTypePython,
	ProjectTypeDocker,
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
	// ResolvePath resolves the organization, repository, release, and node from the given path.
	ResolvePath(context.Context, string) (ResolvedPath, error)
	// Signer returns the transaction signer.
	Signer() *signer.Signer
	// Storage returns the storage provider.
	Storage() storage.Provider
}

type OrganizationAPI interface {
	GetOrganization(context.Context, common.Hash) (*Organization, error)
	GetOrganizationMeta(context.Context, string) (*OrganizationMeta, error)
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
	GetReleaseMeta(context.Context, string) (*ReleaseMeta, error)
	GetLatestRelease(context.Context, common.Hash, string) (*Release, error)
	ListReleaseTags(common.Hash, string) ReleaseTagIterator
	ListReleases(common.Hash, string) ReleaseIterator
	VoteRelease(context.Context, common.Hash, string, *Release) (*valist.ValistVoteReleaseEvent, error)
}

type RepositoryAPI interface {
	GetRepository(context.Context, common.Hash, string) (*Repository, error)
	GetRepositoryMeta(context.Context, string) (*RepositoryMeta, error)
	CreateRepository(context.Context, common.Hash, string, *RepositoryMeta) (*valist.ValistRepoCreated, error)
	SetRepositoryMeta(context.Context, common.Hash, string, *RepositoryMeta) (*valist.ValistMetaUpdate, error)
	VoteRepoDev(context.Context, common.Hash, string, common.Hash, common.Address) (*valist.ValistVoteKeyEvent, error)
	VoteRepositoryThreshold(context.Context, common.Hash, string, *big.Int) (*valist.ValistVoteThresholdEvent, error)
}

type ReleaseTagIterator interface {
	Next(context.Context) (string, error)
	ForEach(context.Context, func(string)) error
}

type ReleaseIterator interface {
	Next(context.Context) (*Release, error)
	ForEach(context.Context, func(*Release)) error
}

type Organization struct {
	ID            common.Hash
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       string
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
	ReleaseCID string
	MetaCID    string
	Signers    []common.Address
}

type Artifact struct {
	SHA256   string `json:"sha256"`
	Provider string `json:"provider"`
}

type ReleaseMeta struct {
	Name      string              `json:"name"`
	Readme    string              `json:"readme"`
	License   string              `json:"license"`
	Artifacts map[string]Artifact `json:"artifacts"`
}

type Repository struct {
	Name          string
	OrgID         common.Hash
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       string
}

type RepositoryMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectType string `json:"projectType"`
	Homepage    string `json:"homepage"`
	Repository  string `json:"repository"`
}

type ResolvedPath struct {
	Organization *Organization
	OrgID        common.Hash
	OrgName      string

	Repository *Repository
	RepoName   string

	Release    *Release
	ReleaseTag string
}
