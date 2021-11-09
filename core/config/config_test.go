package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	dir, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")
	defer os.RemoveAll(dir)

	err = Initialize(dir)
	require.NoError(t, err, "Failed to initialize config")

	err = NewConfig(dir).Load()
	require.NoError(t, err, "Failed to load config")
}
