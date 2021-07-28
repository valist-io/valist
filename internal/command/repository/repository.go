package repository

import (
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:    "repository",
		Aliases: []string{"repo"},
		Usage:   "Create, update, or fetch repositories",
		Subcommands: []*cli.Command{
			NewFetchCommand(),
		},
	}
}
