package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	dir, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")
	defer os.RemoveAll(dir)

	exists, err := Exists(dir)
	require.NoError(t, err, "Failed to check if config exists")
	assert.False(t, exists)

	err = Init(dir)
	require.NoError(t, err, "Failed to init config")

	_, err = Load(dir)
	require.NoError(t, err, "Failed to load config")
}
