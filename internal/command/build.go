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
			var dockerConfig = build.DockerConfig{
				Path:         "Dockerfile",
				BaseImage:    "golang:buster",
				Source:       "./",
				BuildCommand: "make all",
			}

			if err := build.GenerateDockerfile(dockerConfig); err != nil {
				return err
			}

			if err := build.Create("valist-build"); err != nil {
				return err
			}

			return build.Export("valist-build", "dist")
		},
	}
}
