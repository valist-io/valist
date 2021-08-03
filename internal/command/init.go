package command

import (
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
			// Get interactive flag value
			isInteractive := c.Bool("wizard")

			if isInteractive {
				return build.ValistFileFromWizard()
			} else {
				if c.NArg() != 1 {
					cli.ShowSubcommandHelpAndExit(c, 1)
				}

				projectType := c.Args().Get(0)
				return build.ValistFileFromTemplate(projectType)
			}
		},
	}
}
