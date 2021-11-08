package valist

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/valist-io/valist/contract/registry"
	"github.com/valist-io/valist/contract/valist"
	"github.com/valist-io/valist/core/types"
	"github.com/valist-io/valist/signer"
	"github.com/valist-io/valist/storage"
)

// API defines the valist api.
type API interface {
	// GetOrganization returns the organization with the given orgID.
	GetOrganization(ctx context.Context, orgID common.Hash) (*types.Organization, error)
	// GetOrganizationMeta returns the organization metadata from the given path.
	GetOrganizationMeta(ctx context.Context, path string) (*types.OrganizationMeta, error)
	// SetOrganizationMeta updates the metadata of the organization with the given orgID.
	SetOrganizationMeta(ctx context.Context, orgID common.Hash, meta *types.OrganizationMeta) (*valist.ValistMetaUpdate, error)
	// CreateOrganization creates a new organization using the given metadata.
	CreateOrganization(ctx context.Context, meta *types.OrganizationMeta) (*valist.ValistOrgCreated, error)
	// VoteOrganizationAdmin votes to add, revoke, or rotate an organization admin key.
	VoteOrganizationAdmin(ctx context.Context, orgID common.Hash, operation common.Hash, address common.Address) (*valist.ValistVoteKeyEvent, error)
	// VoteOrganizationThreshold votes to set a multifactor voting threshold for the organization.
	VoteOrganizationThreshold(ctx context.Context, orgID common.Hash, threshold *big.Int) (*valist.ValistVoteThresholdEvent, error)
	// GetOrganizationID returns the orgID for the given orgName.
	GetOrganizationID(ctx context.Context, orgName string) (common.Hash, error)
	// LinkOrganizationName links the given orgID to the orgName.
	LinkOrganizationName(ctx context.Context, orgID common.Hash, orgName string) (*registry.ValistRegistryMappingEvent, error)
	// GetRelease returns the release from the repo with the given repoName, orgID, and tag.
	GetRelease(ctx context.Context, orgID common.Hash, repoName string, tag string) (*types.Release, error)
	// GetReleaseMeta returns the release meta from the given path.
	GetReleaseMeta(ctx context.Context, path string) (*types.ReleaseMeta, error)
	// GetLatestRelease returns the most recent release from the repo with the given repoName and orgID.
	GetLatestRelease(ctx context.Context, orgID common.Hash, repoName string) (*types.Release, error)
	// ListReleaseTags returns an iterator for retrieving all release tags from the repo with the given repoName and orgID.
	ListReleaseTags(orgID common.Hash, repoName string) types.ReleaseTagIterator
	// ListReleases returns an iterator for retrieving all releases from the repo with the given repoName and orgID.
	ListReleases(orgID common.Hash, repoName string) types.ReleaseIterator
	// VoteRelease votes to publish a new release to the repo with the given orgID and repoName.
	VoteRelease(ctx context.Context, orgID common.Hash, repoName string, release *types.Release) (*valist.ValistVoteReleaseEvent, error)
	// GetRepository returns the repository with the given orgID and repoName.
	GetRepository(ctx context.Context, orgID common.Hash, repoName string) (*types.Repository, error)
	// GetRepositoryMeta returns the repository meta from the given path.
	GetRepositoryMeta(ctx context.Context, path string) (*types.RepositoryMeta, error)
	// CreateRepository creates a repository in the organization with the given orgID.
	CreateRepository(ctx context.Context, orgID common.Hash, repoName string, meta *types.RepositoryMeta) (*valist.ValistRepoCreated, error)
	// SetRepositoryMeta updates the metadata of the repository with the given orgID and repoName.
	SetRepositoryMeta(ctx context.Context, orgID common.Hash, repoName string, meta *types.RepositoryMeta) (*valist.ValistMetaUpdate, error)
	// VoteRepoDev votes to add, revoke, or rotate a repository dev key.
	VoteRepoDev(ctx context.Context, orgID common.Hash, repoName string, operation common.Hash, address common.Address) (*valist.ValistVoteKeyEvent, error)
	// VoteRepositoryThreshold votes to set a multifactor voting threshold for the repository.
	VoteRepositoryThreshold(ctx context.Context, orgID common.Hash, repoName string, threshold *big.Int) (*valist.ValistVoteThresholdEvent, error)
	// ResolvePath resolves the organization, repository, release, and node from the given path.
	ResolvePath(ctx context.Context, path string) (types.ResolvedPath, error)
	// Signer returns the transaction signer.
	Signer() *signer.Signer
	// Storage returns the storage provider.
	Storage() storage.Provider
}
