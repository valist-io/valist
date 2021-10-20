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
		Name: "acme-co/go-example",
		Tag:  "0.0.2",
		Artifacts: map[string]string{
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
	fullConfigObject.Name = "; myshellcommand"

	err = fullConfigObject.Validate()
	require.Error(t, err, "Should fail with invalid org")

	fullConfigObject = fullConfigObjectCorrect
	fullConfigObject.Artifacts = make(map[string]string)

	fullConfigObject.Artifacts["inValiD@@@333"] = "/bin/linux/hello-world"

	err = fullConfigObject.Validate()
	require.Error(t, err, "Should fail with invalid platform key")

	fullConfigObject = fullConfigObjectCorrect
	fullConfigObject.Artifacts = make(map[string]string)

	fullConfigObject.Artifacts["linux/amd64"] = "; myshellcommand"

	err = fullConfigObject.Validate()
	require.Error(t, err, "Should fail with invalid platform path")
}

func TestLoadSaveValistConfig(t *testing.T) {
	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")
	defer os.RemoveAll(tmp)

	var fullConfigObject = Config{
		Type: "go",
		Name: "test",
		Tag:  "0.0.2",
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
