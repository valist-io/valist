package command

import (
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/command/account"
	"github.com/valist-io/registry/internal/command/repository"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:        "valist",
		HelpName:    "valist",
		Usage:       "Valist command line interface",
		Description: `Universal package repository.`,
		Commands: []*cli.Command{
			account.NewCommand(),
			repository.NewCommand(),
			NewDaemonCommand(),
		},
	}
}
