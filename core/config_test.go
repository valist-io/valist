package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigRoundTrip(t *testing.T) {
	config := NewConfig(t.TempDir())

	err := config.Init()
	require.NoError(t, err, "Failed to initialize config")

	assert.FileExists(t, config.Path(), "File does not exist")

	err = config.Load()
	require.NoError(t, err, "Failed to load config")
}
