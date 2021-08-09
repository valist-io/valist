package command

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/command/account"
	"github.com/valist-io/registry/internal/command/organization"
	"github.com/valist-io/registry/internal/command/repository"
	"github.com/valist-io/registry/internal/config"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:        "valist",
		HelpName:    "valist",
		Usage:       "Valist command line interface",
		Description: `Universal package repository.`,
		Before: func(c *cli.Context) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			exists, err := config.Exists(home)
			if err != nil {
				return err
			}

			if exists {
				return nil
			}

			return config.Init(home)
		},
		Commands: []*cli.Command{
			account.NewCommand(),
			organization.NewCommand(),
			repository.NewCommand(),
			NewDaemonCommand(),
			NewInitCommand(),
		},
	}
}
