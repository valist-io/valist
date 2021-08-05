package build

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateBuild(t *testing.T) {
	tmp, err := os.MkdirTemp("", "test")
	require.NoError(t, err, "Failed to create tmp dir")

	ymlFilePath := filepath.Join(tmp, "valist.yml")
	dockerFilePath := filepath.Join(tmp, "Dockerfile")

	var dockerConfig = DockerConfig{
		Path:         ymlFilePath,
		BaseImage:    "golang:buster",
		Source:       "./",
		BuildCommand: "go build -o ./dist/main testdata/main.go",
	}

	err = GenerateDockerfile(dockerConfig)
	assert.NoError(t, err, "Generate Dockerfile returns with no errors")
	assert.FileExists(t, ymlFilePath, "Dockerfile has been created")

	err = Create("valist-build", dockerFilePath)
	assert.NoError(t, err, "Create build returns with no errors")
}

func TestExportBuild(t *testing.T) {
	err := Export("valist-build", "dist")
	assert.NoError(t, err, "Export build returns with no errors")
	assert.FileExists(t, "dist/main", "Artifact has been created")
}
