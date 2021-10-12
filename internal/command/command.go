package command

import (
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/command/account"
	"github.com/valist-io/valist/internal/command/organization"
	"github.com/valist-io/valist/internal/command/repository"
	"github.com/valist-io/valist/internal/command/utils/flags"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:        "valist",
		HelpName:    "valist",
		Usage:       "Valist command line interface",
		Description: `Universal package repository.`,
		Flags: []cli.Flag{
			flags.Account(),
			flags.AccountPassphrase(),
			flags.Mock(),
		},
		Commands: []*cli.Command{
			account.NewCommand(),
			organization.NewCommand(),
			repository.NewCommand(),
			NewDaemonCommand(),
			NewBuildCommand(),
			NewInitCommand(),
			NewPublishCommand(),
			NewInstallCommand(),
		},
	}
}
