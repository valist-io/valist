package core

import (
	"context"
	"io"
	"net/http"
	"testing"

	files "github.com/ipfs/go-ipfs-files"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoteIPFS(t *testing.T) {
	ctx := context.Background()
	data := []byte("hello")

	ipfs, err := httpapi.NewURLApiWithClient("https://pin.valist.io", &http.Client{})
	require.NoError(t, err, "Failed to connect to https://pin.valist.io")

	path, err := ipfs.Unixfs().Add(ctx, files.NewBytesFile(data), options.Unixfs.Pin(true))
	require.NoError(t, err, "Failed to add file")

	node, err := ipfs.Unixfs().Get(ctx, path)
	require.NoError(t, err, "Failed to get file")

	file, ok := node.(files.File)
	require.True(t, ok, "Failed to get file")

	expect, err := io.ReadAll(file)
	require.NoError(t, err, "Failed to read file")
	assert.Equal(t, data, expect)
}
