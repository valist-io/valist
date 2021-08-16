package command

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/valist-io/registry/internal/build"
)

func NewBuildCommand() *cli.Command {
	return &cli.Command{
		Name:  "build",
		Usage: "Build the target valist project",
		Action: func(c *cli.Context) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			artifactPaths, err := build.Run(wd, "valist.yml")
			if err != nil {
				return err
			}
			for _, artifact := range artifactPaths {
				fmt.Println("Project artifact created @", artifact)
			}
			return nil
		},
	}
}
