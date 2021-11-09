package ipfs

import (
	"context"
	"testing"

	"github.com/ipfs/go-ipfs/core/coreapi"
	coremock "github.com/ipfs/go-ipfs/core/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadWrite(t *testing.T) {
	node, err := coremock.NewMockNode()
	require.NoError(t, err, "Failed to create ipfs mock node")

	ipfsapi, err := coreapi.NewCoreAPI(node)
	require.NoError(t, err, "Failed to create ipfs core api")

	provider := &Provider{ipfsapi}
	ctx := context.Background()
	data := []byte("hello")

	p, err := provider.Write(ctx, data)
	require.NoError(t, err, "Failed to write file")

	expect, err := provider.ReadFile(ctx, p)
	require.NoError(t, err, "Failed to get file")
	assert.Equal(t, data, expect)
}
