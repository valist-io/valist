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
	"github.com/valist-io/valist/prompt"
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

	err = setupTelemetry(cfg)
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

func setupTelemetry(cfg *config.Config) error {
	if cfg.Telemetry != config.TelemetryNone {
		return nil
	}

	option, err := prompt.StatsOptIn().Run()
	if err == prompt.ErrNonInteractive {
		return nil
	}

	if err != nil {
		return err
	}

	switch option[0] {
	case 'n', 'N':
		cfg.Telemetry = config.TelemetryDeny
	default:
		cfg.Telemetry = config.TelemetryAllow
	}

	return cfg.Save()
}
