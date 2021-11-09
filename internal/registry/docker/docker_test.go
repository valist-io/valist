//go:build experimental
// +build experimental

package docker

import (
	"context"
	"net/http"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valist-io/valist/internal/core/mock"
	"github.com/valist-io/valist/internal/core/types"
)

func TestDockerPush(t *testing.T) {
	ctx := context.Background()

	client, err := mock.NewClient(ctx)
	require.NoError(t, err, "Failed to create mock client")

	orgName := "valist"
	orgMeta := &types.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	repoName := "docker"
	repoMeta := &types.RepositoryMeta{
		Name:        "docker",
		Description: "Valist core sdk.",
		ProjectType: "docker",
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

	registryAddr := "valist.local:5000"
	registryPath := "valist.local:5000/valist/docker"

	go func() {
		err := http.ListenAndServe(registryAddr, NewHandler(client))
		require.NoError(t, err, "Failed to start http server")
	}()

	err = exec.Command("docker", "build", "--tag", registryPath, "./testdata/").Run()
	require.NoError(t, err, "Failed to build docker image")

	err = exec.Command("docker", "push", registryPath).Run()
	require.NoError(t, err, "Failed to push docker image")

	err = exec.Command("docker", "pull", registryPath).Run()
	require.NoError(t, err, "Failed to pull docker image")
}
