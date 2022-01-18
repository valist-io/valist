package valist

import (
	"context"
	"errors"
	"math/big"
)

var (
	ErrTeamNotExist    = errors.New("Team does not exist")
	ErrProjectNotExist = errors.New("Project does not exist")
	ErrReleaseNotExist = errors.New("Release does not exist")
)

// API defines the Valist core interface.
type API interface {
	AccountAPI
	ContractAPI
	StorageAPI

	// GetTeam returns the team with the given name.
	GetTeam(ctx context.Context, teamName string) (*Team, error)
	// GetProject returns the project with the given name.
	GetProject(ctx context.Context, teamName, projectName string) (*Project, error)
	// GetRelease returns the release with the given name.
	GetRelease(ctx context.Context, teamName, projectName, releaseName string) (*Release, error)
	// GetLatestRelease returns the latest release.
	GetLatestRelease(ctx context.Context, teamName, projectName string) (*Release, error)
	// ResolvePath resolves the team, project, and release from the given path.
	ResolvePath(ctx context.Context, path string) (ResolvedPath, error)
}

// AccountAPI defines methods for managing accounts.
type AccountAPI interface {
	// GetAccount returns the current account.
	GetAccount() string
	// HasAccount returns true if the account exists.
	HasAccount(address string) bool
	// SetAccount sets the current account with an optional password.
	SetAccount(address, password string) error
	// CreateAccount creates a new account with the given passphrase.
	CreateAccount(passphrase string) (string, error)
	// ListAccounts returns a list of all accounts.
	ListAccounts() []string
	// ImportAccount imports account data. Format depends on the underlying implementation.
	ImportAccount(data []byte, password, newPassword string) (string, error)
	// ExportAccount exports account data. Format depends on the underlying implementation.
	ExportAccount(address, password, newPassword string) ([]byte, error)
}

// StorageAPI defines methods for reading and writing files.
type StorageAPI interface {
	// ListFiles returns the contents of the directory at the given path.
	ListFiles(ctx context.Context, path string) ([]string, error)
	// ReadFile returns the contents of the file at the given path.
	ReadFile(ctx context.Context, path string) ([]byte, error)
	// WriteFile writes the contents of the file at the given path.
	WriteFile(ctx context.Context, path string) (string, error)
	// WriteBytes writes the given contents.
	WriteBytes(ctx context.Context, data []byte) (string, error)
}

// ContractAPI defines methods for interacting with smart contracts.
type ContractAPI interface {
	// CreateTeam creates a new team with the given members.
	CreateTeam(ctx context.Context, teamName, metaURI string, members []string) error
	// CreateProject creates a new project. Requires the sender to be a member of the team.
	CreateProject(ctx context.Context, teamName, projectName, metaURI string, members []string) error
	// CreateRelease creates a new release. Requires the sender to be a member of the project.
	CreateRelease(ctx context.Context, teamName, projectName, releaseName, metaURI string) error
	// AddTeamMember adds a member to the team. Requires the sender to be a member of the team.
	AddTeamMember(ctx context.Context, teamName string, address string) error
	// RemoveTeamMember removes a member from the team. Requires the sender to be a member of the team.
	RemoveTeamMember(ctx context.Context, teamName string, address string) error
	// AddProjectMemeber adds a member to the project. Requires the sender to be a member of the team.
	AddProjectMember(ctx context.Context, teamName, projectName string, address string) error
	// RemoveProjectMember removes a member from the project. Requires the sender to be a member of the team.
	RemoveProjectMember(ctx context.Context, teamName, projectName string, address string) error
	// SetTeamMetaURI sets the team metadata content ID. Requires the sender to be a member of the team.
	SetTeamMetaURI(ctx context.Context, teamName, metaURI string) error
	// SetProjectMetaURI sets the project metadata content ID. Requires the sender to be a member of the team.
	SetProjectMetaURI(ctx context.Context, teamName, projectName, metaURI string) error
	// ApproveRelease approves the release by adding the sender's address to the approvers list.
	// The sender's address will be removed from the rejectors list if it exists.
	ApproveRelease(ctx context.Context, teamName, projectName, releaseName string) error
	// RejectRelease rejects the release by adding the sender's address to the rejectors list.
	// The sender's address will be removed from the approvers list if it exists.
	RejectRelease(ctx context.Context, teamName, projectName, releaseName string) error
	// GetLatestReleaseName returns the latest release name.
	GetLatestReleaseName(ctx context.Context, teamName, projectName string) (string, error)
	// GetTeamMetaURI returns the team metadata URI.
	GetTeamMetaURI(ctx context.Context, teamName string) (string, error)
	// GetProjectMetaURI returns the project metadata URI.
	GetProjectMetaURI(ctx context.Context, teamName, projectName string) (string, error)
	// GetReleaseMetaURI returns the release metadata URI.
	GetReleaseMetaURI(ctx context.Context, teamName, projectName, releaseName string) (string, error)
	// GetTeamNames returns a paginated list of team names.
	GetTeamNames(ctx context.Context, page *big.Int, size *big.Int) ([]string, error)
	// GetProjectNames returns a paginated list of project names.
	GetProjectNames(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseNames returns a paginated list of release names.
	GetReleaseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error)
	// GetTeamMembers returns a paginated list of team members.
	GetTeamMembers(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error)
	// GetProjectMembers returns a paginated list of project members.
	GetProjectMembers(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseApprovers returns a paginated list of release approvers.
	GetReleaseApprovers(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseRejectors returns a paginated list of release rejectors.
	GetReleaseRejectors(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error)
}

type Team struct {
	// Name is the organization friendly name.
	Name string `json:"name"`
	// Description is a short description of the organization.
	Description string `json:"description"`
	// Homepage is a link to the organization website.
	Homepage string `json:"homepage"`
}

type Project struct {
	// Name is the repository friendly name.
	Name string `json:"name"`
	// Description is a short description of the repository.
	Description string `json:"description"`
	// Homepage is the website for the repository.
	Homepage string `json:"homepage"`
	// Repository is the source code url for the repository.
	Repository string `json:"repository"`
}

type Release struct {
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

// Artifact is file contained in a release.
type Artifact struct {
	// SHA256 is the sha256 of the file.
	SHA256 string `json:"sha256"`
	// Provider is the path to the artifact file.
	Provider string `json:"provider"`
}

type ResolvedPath struct {
	Team        *Team    `json:"team"`
	TeamName    string   `json:"team_name"`
	Project     *Project `json:"project"`
	ProjectName string   `json:"project_name"`
	Release     *Release `json:"release"`
	ReleaseName string   `json:"release_name"`
}
