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

	var fullConfigObjectCorrect = Config{
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

	err = fullConfigObjectCorrect.Validate()
	require.NoError(t, err, "Should create correct config")

	fullConfigObject := fullConfigObjectCorrect
	fullConfigObject.Type = "; myshellcommand"

	err = fullConfigObject.Validate()
	require.Error(t, err, "Should fail with invalid project type")

	fullConfigObject = fullConfigObjectCorrect
	fullConfigObject.Org = "; myshellcommand"

	err = fullConfigObject.Validate()
	require.Error(t, err, "Should fail with invalid org")

	fullConfigObject = fullConfigObjectCorrect
	fullConfigObject.Platforms = make(map[string]string)

	fullConfigObject.Platforms["inValiD@@@333"] = "/bin/linux/hello-world"

	err = fullConfigObject.Validate()
	require.Error(t, err, "Should fail with invalid platform key")

	fullConfigObject = fullConfigObjectCorrect
	fullConfigObject.Platforms = make(map[string]string)

	fullConfigObject.Platforms["linux/amd64"] = "; myshellcommand"

	err = fullConfigObject.Validate()
	require.Error(t, err, "Should fail with invalid platform path")
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
		Org:   "acme-co",
		Repo:  "go-example",
		Tag:   "1.0.0",
		Type:  "go",
		Build: "go build",
		Out:   "path_to_artifact_or_build_directory",
	}

	var NpmConfigObject = Config{
		Org:   "acme-co",
		Repo:  "npm-example",
		Tag:   "1.0.0",
		Type:  "npm",
		Build: "npm run build",
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
