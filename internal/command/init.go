package command

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/valist-io/registry/internal/build"
)

func NewInitCommand() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Generate a new Valist project",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "wizard",
				Aliases: []string{"i"},
				Usage:   "Enable interactive wizard",
			},
		},
		Action: func(c *cli.Context) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			if c.Bool("wizard") {
				return build.ConfigWizard()
			}

			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			projectType := c.Args().Get(0)
			valistFilePath := filepath.Join(wd, "valist.yml")
			return build.ConfigTemplate(projectType, valistFilePath)
		},
	}
}
