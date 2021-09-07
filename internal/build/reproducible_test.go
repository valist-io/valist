package build

import (
	"os"
	"path/filepath"
	"testing"

	copy "github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateExportBuild(t *testing.T) {
	dockerImageName := "valist-test-create-export"
	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")

	defer os.RemoveAll(tmp)

	// Copy goTestProject from testdata to tmp directory
	err = copy.Copy("testdata/goTestProject", tmp)
	require.NoError(t, err, "Failed to copy test files to tmp directory")

	// Create dockerConfig with Path, Image, Source, & BuildCommand
	dockerFilePath := filepath.Join(tmp, "Dockerfile.repro")
	var dockerConfig = DockerConfig{
		Path:         dockerFilePath,
		BaseImage:    "golang:buster",
		Source:       "./",
		BuildCommand: "go build -o dist/main src/main.go",
	}

	err = GenerateDockerfile(dockerConfig)
	assert.NoError(t, err, "Generate Dockerfile returns with no errors")
	assert.FileExists(t, dockerFilePath, "Dockerfile has been created")

	err = Create(dockerImageName, dockerFilePath)
	assert.NoError(t, err, "Create build returns with no errors")

	containerPath := "dist"
	err = Export(dockerImageName, tmp, containerPath)

	assert.NoError(t, err, "Export build returns with no errors")
	assert.FileExists(t, filepath.Join(tmp, "dist/main"), "Artifact file exists")
	assert.FileExists(t, filepath.Join(tmp, "Dockerfile.repro"), "Dockerfile exists")
	assert.FileExists(t, filepath.Join(tmp, "Dockerfile.repro.dockerignore"), "dockerignore exists")
}
