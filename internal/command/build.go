package command

import (
	"github.com/urfave/cli/v2"
	"github.com/valist-io/registry/internal/build"
)

func NewBuildCommand() *cli.Command {
	return &cli.Command{
		Name:  "build",
		Usage: "Build the target valist project",
		Action: func(c *cli.Context) error {
			var dockerfile = build.Dockerfile{
				Path:         "Dockerfile",
				BaseImage:    "golang:buster",
				Source:       "./",
				BuildCommand: "make all",
			}

			build.GenerateDockerfile(dockerfile)
			build.Create("valist-build")
			build.Export("valist-build", "dist")

			return nil
		},
	}
}
