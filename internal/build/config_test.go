package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvalidConfig(t *testing.T) {
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
		Platforms: map[string]string{
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

func TestLoadSaveValistConfig(t *testing.T) {
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
		Platforms: map[string]string{
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

func TestValistFileFromTemplate(t *testing.T) {
	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")
	defer os.RemoveAll(tmp)

	GoCfgPath := filepath.Join(tmp, "go.valist.yml")
	NpmCfgPath := filepath.Join(tmp, "npm.valist.yml")

	err = ConfigTemplate("go", GoCfgPath)
	require.NoError(t, err)

	err = ConfigTemplate("npm", NpmCfgPath)
	require.NoError(t, err)

	assert.FileExists(t, GoCfgPath, "Valist file for go project has been created")
	assert.FileExists(t, NpmCfgPath, "Valist file for npm project has been created")

	var GoConfigObject = Config{
		Org:   "test",
		Repo:  "test",
		Tag:   "test",
		Type:  "go",
		Build: "go build",
		Out:   "path_to_artifact_or_build_directory",
	}

	var NpmConfigObject = Config{
		Org:   "test",
		Repo:  "test",
		Tag:   "test",
		Type:  "npm",
		Build: "npm run build",
		// Out:   "should error",
	}

	var GoConfigFile Config
	err = GoConfigFile.Load(GoCfgPath)

	require.NoError(t, err, "Failed to load config")
	assert.Equal(t, GoConfigObject, GoConfigFile)

	var NpmConfigFile Config
	err = NpmConfigFile.Load(NpmCfgPath)
	require.NoError(t, err, "Failed to load config")
	assert.Equal(t, NpmConfigObject, NpmConfigFile)
}
