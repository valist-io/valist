package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDockerfile(t *testing.T) {

	var dockerConfig = DockerConfig{
		Path:         "Dockerfile",
		BaseImage:    "golang:buster",
		Source:       "./",
		BuildCommand: "go build -o ./dist/main testdata/main.go",
	}

	GenerateDockerfile(dockerConfig)
	assert.FileExists(t, "Dockerfile", "Dockerfile has been created")
}

func TestCreateBuild(t *testing.T) {
	err := Create("valist-build")
	assert.NoError(t, err, "Valist build returns with no errors")
}

func TestExportBuild(t *testing.T) {
	Export("valist-build", "dist")
	assert.FileExists(t, "dist/main", "Artifact has been created")
}
