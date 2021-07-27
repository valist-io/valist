package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateValistConfig(t *testing.T) {
	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")
	defer os.RemoveAll(tmp)

	var config = Config{
		Type:    "go",
		Org:     "test",
		Repo:    "binary",
		Tag:     "0.0.2",
		Meta:    "README.md",
		Build:   "make all",
		Install: "go mod tidy",
		Out:     "dist",
		Artifacts: map[string]string{
			"linux/amd64":  "bin/lin000x/hello-world",
			"darwin/amd64": "bin/macz/hello-world",
		},
	}

	cfgPath := filepath.Join(tmp, "valist.yml")
	err = config.Save(cfgPath)
	require.NoError(t, err, "Failed to create tmp dir")
	assert.FileExists(t, cfgPath, "Valist file has been created")

	var other Config
	err = other.Load(cfgPath)
	require.NoError(t, err, "Failed to load config")
	assert.Equal(t, config, other)

	// var other2 Config
	// err = other2.Load("testdata/2.valist.yml")
	// require.NoError(t, err, "Failed to load config")
	// assert.Equal(t, config, other2)
}
