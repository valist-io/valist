package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/valist-io/valist/internal/build"
)

func NewBuildCommand() *cli.Command {
	return &cli.Command{
		Name:  "build",
		Usage: "Build the target valist project",
		Action: func(c *cli.Context) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			var valistFile build.Config
			if err := valistFile.Load(filepath.Join(cwd, "valist.yml")); err != nil {
				return err
			}

			artifactPaths, err := build.Run(cwd, valistFile)
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
