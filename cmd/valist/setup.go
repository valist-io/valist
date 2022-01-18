package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/command"
	"github.com/valist-io/valist/core"
	"github.com/valist-io/valist/core/config"
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
	// setup default account from config or override from flags
	client.SetAccount(cfg.Accounts.Default, "")
	if c.IsSet("account") {
		client.SetAccount(c.String("account"), c.String("passphrase"))
	}
	c.Context = context.WithValue(c.Context, command.ClientKey, client)
	c.Context = context.WithValue(c.Context, command.ConfigKey, cfg)
	return nil
}

// setup initializes the config only (no client) before commands are executed
func setupConfig(c *cli.Context) error {
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
	c.Context = context.WithValue(c.Context, command.ConfigKey, cfg)
	return nil
}
