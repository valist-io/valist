package command

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/command/account"
	"github.com/valist-io/valist/internal/command/organization"
	"github.com/valist-io/valist/internal/command/repository"
	"github.com/valist-io/valist/internal/core/config"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:        "valist",
		HelpName:    "valist",
		Usage:       "Valist command line interface",
		Description: `Universal package repository.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "account",
				Usage: "Account to transact with",
			},
		},
		Commands: []*cli.Command{
			account.NewCommand(),
			organization.NewCommand(),
			repository.NewCommand(),
			NewDaemonCommand(),
			NewBuildCommand(),
			NewInitCommand(),
			NewPublishCommand(),
		},
		Before: func(c *cli.Context) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			return config.Initialize(home)
		},
	}
}
