package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDockerfile(t *testing.T) {

	var dockerfile = Dockerfile{
		Path:         "Dockerfile",
		BaseImage:    "golang:buster",
		Source:       "./",
		BuildCommand: "go build -o ./dist/main testdata/main.go",
	}

	GenerateDockerfile(dockerfile)

	assert.FileExists(t, "Dockerfile", "Dockerfile has been created")
}

func TestCreateBuild(t *testing.T) {
	Create("valist-build")
}

func TestExportBuild(t *testing.T) {
	Export("valist-build", "dist")
}
