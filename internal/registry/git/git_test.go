package git

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

func TestGitPush(t *testing.T) {
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

	registryAddr := "localhost:10002"
	registryPath := "http://localhost:10002/valist/sdk"

	go func() {
		err := http.ListenAndServe(registryAddr, NewHandler(client))
		require.NoError(t, err, "Failed to start http server")
	}()

	err = exec.Command("git", "push", "--all", registryPath).Run()
	require.NoError(t, err, "Failed to push git repo")
}
