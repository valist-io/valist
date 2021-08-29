package npm

import (
	"context"
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valist-io/registry/internal/core/mock"
	"github.com/valist-io/registry/internal/core/types"
)

func TestNpmPublish(t *testing.T) {
	ctx := context.Background()

	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create temp dir")
	defer os.RemoveAll(tmp)

	client, err := mock.NewClient(tmp)
	require.NoError(t, err, "Failed to create mock client")

	orgName := "valist"
	orgMeta := &types.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	repoName := "sdk"
	repoMeta := &types.RepositoryMeta{
		Name:        "sdk",
		Description: "Valist core sdk.",
		ProjectType: "npm",
		Homepage:    "https://valist.io",
		Repository:  "https://github.com/valist-io/valist",
	}

	createEvent, err := client.CreateOrganization(ctx, orgMeta)
	require.NoError(t, err, "Failed to create organization")
	orgID := createEvent.OrgID

	_, err = client.LinkOrganizationName(ctx, orgID, orgName)
	require.NoError(t, err, "Failed to link organization name")

	_, err = client.CreateRepository(ctx, orgID, repoName, repoMeta)
	require.NoError(t, err, "Failed to create repository")

	registryAddr := "localhost:10001"
	registryPath := "http://localhost:10001"

	go http.ListenAndServe(registryAddr, NewHandler(client)) //nolint:errcheck

	err = exec.Command("npm", "publish", "./testdata/v0", "--registry", registryPath).Run()
	require.NoError(t, err, "Failed to publish npm package")

	err = exec.Command("npm", "view", "@valist/sdk", "--registry", registryPath).Run()
	require.NoError(t, err, "Failed to view npm package")

	err = exec.Command("npm", "publish", "./testdata/v1", "--registry", registryPath).Run()
	require.NoError(t, err, "Failed to publish npm package")
}
