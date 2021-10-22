package command

import (
	"github.com/urfave/cli/v2"
	"github.com/valist-io/valist/internal/build"
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
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			projectName := c.Args().Get(0)

			cfg := build.Config{
				Name:      projectName,
				Tag:       "0.0.1",
				Artifacts: map[string]string{"linux/amd64": "path_to_bin"},
			}

			return cfg.Save("valist.yml")
		},
	}
}
