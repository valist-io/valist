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

	c.Context = context.WithValue(c.Context, command.ClientKey, client)
	c.Context = context.WithValue(c.Context, command.ConfigKey, config)

	var acct = config.GetDefaultAccount()
	var pass = c.String("passphrase")
	// setup default account from config or override from flags
	if c.IsSet("account") {
		acct = c.String("account")
	}
	if err = client.SetAccount(acct, pass); err != nil {
		logger.Warn("Default account not found. Create one with the following command:")
		logger.Warn("valist account create")
	}
	return nil
}
