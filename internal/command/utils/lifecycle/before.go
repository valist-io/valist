package lifecycle

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
)

func SetupClient(c *cli.Context) error {
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

	var opts core.Options
	if c.IsSet("account") {
		opts.Account.Address = common.HexToAddress(c.String("account"))
	} else {
		opts.Account.Address = cfg.Accounts.Default
	}

	if c.IsSet("passphrase") {
		opts.Passphrase = c.String("passphrase")
	}

	client, err := core.NewClient(c.Context, cfg, opts)
	if err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, core.ClientKey, client)
	c.Context = context.WithValue(c.Context, core.ConfigKey, cfg)
	return nil
}
