package npm

import (
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:    "npm",
		Usage:   "Node package manager",
		Subcommands: []*cli.Command{
			NewPublishCommand(),
		},
	}
}
