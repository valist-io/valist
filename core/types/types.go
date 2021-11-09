package types

import (
	"context"
	"errors"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
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
	ErrOrgNotExist     = errors.New("Organization does not exist")
	ErrRepoNotExist    = errors.New("Repository does not exist")
	ErrReleaseNotExist = errors.New("Release does not exist")
)

// Organization is a group of users and repositories.
type Organization struct {
	// ID is the unique organization ID.
	ID common.Hash
	// Threshold is the multifactor voting threshold.
	Threshold *big.Int
	// ThresholdDate is when the threshold was updated.
	ThresholdDate *big.Int
	// MetaCID is the path to the metadata file.
	MetaCID string
}

// OrganizationMeta contains info about an organization.
type OrganizationMeta struct {
	// Name is the organization friendly name.
	Name string `json:"name"`
	// Description is a short description of the organization.
	Description string `json:"description"`
	// Homepage is a link to the organization website.
	Homepage string `json:"homepage"`
}

// Release contains info about a published release.
type Release struct {
	// Tag is the unique identifier for the release.
	Tag string
	// ReleaseCID is the path to the release artifacts.
	ReleaseCID string
	// MetaCID is the path to the release metadata file.
	MetaCID string
	// Signers is a list of keys that have signed the release.
	Signers []common.Address
}

// Artifact is file contained in a release.
type Artifact struct {
	// SHA256 is the sha256 of the file.
	SHA256 string `json:"sha256"`
	// Provider is the path to the artifact file.
	Provider string `json:"provider"`
}

// ReleaseMeta contains info about a release.
type ReleaseMeta struct {
	// Name is the full release path.
	Name string `json:"name"`
	// Version is the release version.
	Version string `json:"version"`
	// Readme contains the readme contents.
	Readme string `json:"readme"`
	// License contains the license type.
	License string `json:"license"`
	// Dependencies contains a list of all dependencies.
	Dependencies []string `json:"dependencies"`
	// Artifacts is a mapping of names to artifacts.
	Artifacts map[string]Artifact `json:"artifacts"`
}

// Repository is a set of versioned releases.
type Repository struct {
	// Name is the repository name.
	Name string
	// OrgID is the parent organization ID.
	OrgID common.Hash
	// Threshold is the multifactor voting threshold.
	Threshold *big.Int
	// ThresholdDate is when the threshold was updated.
	ThresholdDate *big.Int
	// MetaCID is the path to the metadata file.
	MetaCID string
}

// RepositoryMeta contains info about a repository.
type RepositoryMeta struct {
	// Name is the repository friendly name.
	Name string `json:"name"`
	// Description is a short description of the repository.
	Description string `json:"description"`
	// ProjectType is used to change how the repository is displayed.
	ProjectType string `json:"projectType"`
	// Homepage is the website for the repository.
	Homepage string `json:"homepage"`
	// Repository is the source code url for the repository.
	Repository string `json:"repository"`
}

// ReleaseTagIterator iterates all repository release tags.
type ReleaseTagIterator interface {
	// Next returns the next tag or io.EOF if finished.
	Next(ctx context.Context) (string, error)
	// ForEach runs the given function for each tag.
	ForEach(ctx context.Context, cb func(string)) error
}

// ReleaseIterator iterates all repository releases.
type ReleaseIterator interface {
	// Next returns the next release or io.EOF if finished.
	Next(ctx context.Context) (*Release, error)
	// ForEach runs the given function for each release.
	ForEach(ctx context.Context, cb func(*Release)) error
}

type ResolvedPath struct {
	Organization *Organization
	OrgID        common.Hash
	OrgName      string
	Repository   *Repository
	RepoName     string
	Release      *Release
	ReleaseTag   string
}
