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

	var fullConfigObject = Config{
		Type:    "go",
		Org:     "test",
		Repo:    "binary",
		Tag:     "0.0.2",
		Meta:    "README.md",
		Build:   "make all",
		Install: "go mod tidy",
		Out:     "dist",
		Artifacts: map[string]string{
			"linux/amd64":  "bin/linux/hello-world",
			"darwin/amd64": "bin/darwin/hello-world",
		},
	}

	cfgPath := filepath.Join(tmp, "valist.yml")
	err = fullConfigObject.Save(cfgPath)
	require.NoError(t, err, "Failed to create tmp dir")
	assert.FileExists(t, cfgPath, "Valist file has been created")

	var fullConfigFile Config
	err = fullConfigFile.Load(cfgPath)
	require.NoError(t, err, "Failed to load config")
	assert.Equal(t, fullConfigObject, fullConfigFile)
}
