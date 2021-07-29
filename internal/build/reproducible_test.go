package build

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDockerfile(t *testing.T) {

	var dockerfile = Dockerfile{
		path:         "Dockerfile",
		baseImage:    "golang:buster",
		source:       "./",
		buildCommand: "go build -o ./dist/main testdata/main.go",
	}

	generateDockerfile(dockerfile)

	assert.FileExists(t, "Dockerfile", "Dockerfile has been created")
}

func TestCreateBuild(t *testing.T) {

	CreateBuild("valist-build")
}
