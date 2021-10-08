package estuary

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	host  = "https://pin-proxy-rkl5i.ondigitalocean.app"
	token = "test"
)

func TestWrite(t *testing.T) {
	provider := NewProvider(host, token)

	ctx := context.Background()
	data := []byte("hello")

	_, err := provider.Write(ctx, data)
	require.NoError(t, err, "Failed to write file")
}
