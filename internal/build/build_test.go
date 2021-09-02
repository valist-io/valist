package build

import (
	"os"
	"path/filepath"
	"testing"

	copy "github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunGoBuild(t *testing.T) {
	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")
	defer os.RemoveAll(tmp)

	// Copy goTestProject from testdata to tmp directory
	err = copy.Copy("testdata/goTestProject", tmp)
	require.NoError(t, err, "Failed to copy goTestProject")

	var valistFile Config
	err = valistFile.Load(filepath.Join(tmp, "valist.yml"))
	require.NoError(t, err, "Failed to load config")

	artifactPaths, err := Run(tmp, valistFile)
	assert.NoError(t, err, "build.Run() executes with no errors")

	for _, artifact := range artifactPaths {
		assert.FileExists(t, artifact)
	}
}

func TestRunNpmBuild(t *testing.T) {
	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")
	defer os.RemoveAll(tmp)

	// Copy npmTestProject from testdata to tmp directory
	err = copy.Copy("testdata/npmTestProject", tmp)
	require.NoError(t, err, "Failed to copy npmTestProject")

	var valistFile Config
	err = valistFile.Load(filepath.Join(tmp, "valist.yml"))
	require.NoError(t, err, "Failed to load config")

	artifactPaths, err := Run(tmp, valistFile)
	assert.NoError(t, err, "build.Run() executes with no errors")

	for _, artifact := range artifactPaths {
		assert.FileExists(t, artifact)
	}
}
