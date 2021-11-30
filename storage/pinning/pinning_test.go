package pinning

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/valist-io/valist/storage/ipfs"
)

const (
	host = "https://pin.valist.io"
)

func TestWrite(t *testing.T) {
	tmp, err := os.MkdirTemp("", "")
	require.NoError(t, err, "Failed to MkdirTemp")

	ctx := context.Background()
	data := []byte("hello")

	ipfs, err := ipfs.NewProvider(ctx, tmp)
	require.NoError(t, err, "Failed to write file")

	provider := NewProvider(host, ipfs)
	_, err = provider.Write(ctx, data)
	require.NoError(t, err, "Failed to write file")
}
