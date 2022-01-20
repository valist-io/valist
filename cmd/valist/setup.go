package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/command"
	"github.com/valist-io/valist/core"
)

// setup initializes the client before commands are executed
func setup(c *cli.Context) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	config := core.NewConfig(home)
	if err := config.Init(); err != nil {
		return err
	}
	if err := config.Load(); err != nil {
		return err
	}
	client, err := core.NewClient(c.Context, config)
	if err != nil {
		return err
	}
	if err := setupAccount(c, client, config); err != nil {
		logger.Warn("Failed to set default account")
	}

	c.Context = context.WithValue(c.Context, command.ClientKey, client)
	c.Context = context.WithValue(c.Context, command.ConfigKey, config)
}

// setupAccount sets the default account and optionally passphrase.
func setupAccount(c *cli.Context, client *core.Client, config *core.Config) error {
	acct := config.GetDefaultAccount()
	if c.IsSet("account") {
		acct = c.String("account")
	}
	if c.IsSet("passphrase") {
		return client.SetAccountWithPassphrase(acct, c.String("passphrase"))
	}
	return client.SetAccount(acct)
}