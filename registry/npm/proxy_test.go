package npm

import (
	"context"
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valist-io/valist/core/mock"
	"github.com/valist-io/valist/database/memory"
)

func TestNpmProxy(t *testing.T) {
	ctx := context.Background()
	database := memory.NewDatabase()

	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create temp dir")
	defer os.RemoveAll(tmp)

	client, err := mock.NewClient(ctx)
	require.NoError(t, err, "Failed to create mock client")

	registryAddr := "localhost:10006"
	registryPath := "http://localhost:10006"

	go func() {
		err := http.ListenAndServe(registryAddr, NewProxy(client, database, registryAddr))
		require.NoError(t, err, "Failed to start http server")
	}()

	cmd := exec.Command("npm", "init", "-y")
	cmd.Dir = tmp

	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "Failed to init npm package: %s", out)

	cmd = exec.Command("npm", "install", "react", "--registry", registryPath)
	cmd.Dir = tmp

	out, err = cmd.CombinedOutput()
	require.NoError(t, err, "Failed to install npm package: %s", out)
}
