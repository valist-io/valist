package repository

import (
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/command/repository/key"
	"github.com/valist-io/valist/internal/command/utils/lifecycle"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:    "repository",
		Aliases: []string{"repo"},
		Usage:   "Create, update, or fetch repositories",
		Before:  lifecycle.SetupClient,
		Subcommands: []*cli.Command{
			NewFetchCommand(),
			NewCreateCommand(),
			NewUpdateCommand(),
			NewThresholdCommand(),
			key.NewCommand(),
		},
	}
}
