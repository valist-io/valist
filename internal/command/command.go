package command

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/command/account"
	"github.com/valist-io/valist/internal/command/organization"
	"github.com/valist-io/valist/internal/command/repository"
	"github.com/valist-io/valist/internal/core"
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

			if err := config.Initialize(home); err != nil {
				return err
			}

			cfg := config.NewConfig(home)
			if err := cfg.Load(); err != nil {
				return err
			}

			var account accounts.Account
			if c.IsSet("account") {
				account.Address = common.HexToAddress(c.String("account"))
			} else {
				account.Address = cfg.Accounts.Default
			}

			client, err := core.NewClient(c.Context, cfg, account)
			if err != nil {
				return err
			}

			c.Context = context.WithValue(c.Context, core.ClientKey, client)
			c.Context = context.WithValue(c.Context, core.ConfigKey, cfg)
			return nil
		},
	}
}
