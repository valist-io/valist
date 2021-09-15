package organization

import (
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/command/organization/key"
	"github.com/valist-io/valist/internal/command/utils/lifecycle"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:    "organization",
		Aliases: []string{"org"},
		Usage:   "Create, update, or fetch organizations",
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
