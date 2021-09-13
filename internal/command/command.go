package command

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

			// use env variable VALIST_SIGNER as signing key when set
			if os.Getenv("VALIST_SIGNER") != "" {
				// generate cryptographically secure, ephemeral encryption key that will clear when program halts
				randBytes := make([]byte, 32)
				_, err := rand.Read(randBytes)
				if err != nil {
					return err
				}

				passphrase := fmt.Sprintf("%x", randBytes)

				private, err := crypto.HexToECDSA(os.Getenv("VALIST_SIGNER"))
				if err != nil {
					return err
				}

				// when VALIST_SIGNER is set, this encrypted key is stored in an os.TempDir()
				// when program halts, the ephemeral encryption key is cleared, making the keystore file inaccessible
				// this prevents having to clear the keystore file during runtime, which could fail to execute if program is interrupted/crashes
				account, err = cfg.KeyStore().ImportECDSA(private, passphrase)
				if err != nil {
					return err
				}
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
