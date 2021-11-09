package main

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/command"
	"github.com/valist-io/valist/core"
	"github.com/valist-io/valist/core/config"
	"github.com/valist-io/valist/signer"
)

// setup initializes the client before commands are executed
func setup(c *cli.Context) error {
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

	client, err := core.NewClient(c.Context, cfg)
	if err != nil {
		return err
	}

	var account accounts.Account
	if os.Getenv(signer.EnvKey) == "" {
		if c.IsSet("account") {
			account.Address = common.HexToAddress(c.String("account"))
		} else {
			account.Address = cfg.Accounts.Default
		}

		if c.IsSet("passphrase") {
			client.Signer().SetAccountWithPassphrase(account, c.String("passphrase"))
		} else {
			client.Signer().SetAccount(account)
		}
	}

	c.Context = context.WithValue(c.Context, command.ClientKey, client)
	c.Context = context.WithValue(c.Context, command.ConfigKey, cfg)
	return nil
}
