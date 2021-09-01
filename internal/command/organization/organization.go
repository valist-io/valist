package organization

import (
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:    "organization",
		Aliases: []string{"org"},
		Usage:   "Create, update, or fetch organizations",
		Subcommands: []*cli.Command{
			NewCreateCommand(),
			NewKeyCommand(),
			NewThresholdCommand(),
		},
	}
}
