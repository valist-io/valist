package build

import (
	"fmt"
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
	dockerFilePath := filepath.Join(tmp, "Dockerfile")
	var dockerConfig = DockerConfig{
		Path:         filepath.Join(tmp, "Dockerfile"),
		BaseImage:    "golang:buster",
		Source:       "./",
		BuildCommand: "go build -o dist/main src/main.go",
	}

	err = GenerateDockerfile(dockerConfig)
	assert.NoError(t, err, "Generate Dockerfile returns with no errors")
	assert.FileExists(t, dockerFilePath, "Dockerfile has been created")

	err = Create(dockerImageName, tmp)
	assert.NoError(t, err, "Create build returns with no errors")

	containerPath := fmt.Sprintf("%s:/opt/build/%s", dockerImageName, "dist")
	hostPath := filepath.Join(tmp, "dist")

	err = Export(dockerImageName, hostPath, containerPath)
	assert.NoError(t, err, "Export build returns with no errors")
	assert.FileExists(t, filepath.Join(tmp, "dist/main"), "Artifact file exists")
}
