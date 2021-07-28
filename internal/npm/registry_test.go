package npm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/impl"
)

func TestGetScopedPackage(t *testing.T) {
	ctx := context.Background()

	client, err := impl.NewMockClient()
	require.NoError(t, err, "Failed to create mock client")

	registry := NewRegistry(client)

	orgName := "valist"
	orgMeta := &core.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	repoName := "sdk"
	repoMeta := &core.RepositoryMeta{
		Name:        "sdk",
		Description: "Valist core sdk.",
		ProjectType: "npm",
		Homepage:    "https://valist.io",
		Repository:  "https://github.com/valist-io/valist",
	}

	metaCID, err := client.AddFile(ctx, []byte("{}"))
	require.NoError(t, err, "Failed to add meta file")

	releaseCID, err := client.AddFile(ctx, []byte("hello"))
	require.NoError(t, err, "Failed to add release file")

	release := &core.Release{
		Tag:        "v0.0.1",
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
	}

	txc1, err := client.CreateOrganization(ctx, orgMeta)
	require.NoError(t, err, "Failed to create organization")
	client.Commit()

	res1 := <-txc1
	require.NoError(t, res1.Err, "Failed to create organization")
	orgID := res1.OrgID

	txc2, err := client.LinkOrganizationName(ctx, orgID, orgName)
	require.NoError(t, err, "Failed to link organization name")
	client.Commit()

	res2 := <-txc2
	require.NoError(t, res2.Err, "Failed to link organization name")

	txc3, err := client.CreateRepository(ctx, orgID, repoName, repoMeta)
	require.NoError(t, err, "Failed to create repository")
	client.Commit()

	res3 := <-txc3
	require.NoError(t, res3.Err, "Failed to create repository")

	txc4, err := client.VoteRelease(ctx, orgID, repoName, release)
	require.NoError(t, err, "Failed to vote release")
	client.Commit()

	res4 := <-txc4
	require.NoError(t, res4.Err, "Failed to vote release")

	pack, err := registry.GetScopedPackage(ctx, orgName, repoName)
	require.NoError(t, err, "Failed to get package")
	assert.Equal(t, "@valist/sdk", pack.ID)
}
