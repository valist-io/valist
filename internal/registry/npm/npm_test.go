package npm

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valist-io/registry/internal/core/mock"
	"github.com/valist-io/registry/internal/core/types"
)

func TestPublish(t *testing.T) {
	ctx := context.Background()

	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create temp dir")

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
	registryPath := "http://localhost:10001/@valist"

	go http.ListenAndServe(registryAddr, NewHandler(client))

	err := exec.Command("npm", "publish", "./testdata", "--registry", registryPath).Run()
	require.NoError(t, err, "Failed to publish npm package")
}