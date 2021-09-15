package lifecycle

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/prompt"
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
}

func SetupClientWithPassphrase(c *cli.Context) error {
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

	passphrase, err := prompt.AccountPassphrase().RunFlag(c, "passphrase")
	if err != nil {
		return err
	}

	client, err := core.NewClientWithPassphrase(c.Context, cfg, account, passphrase)
	if err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, core.ClientKey, client)
	c.Context = context.WithValue(c.Context, core.ConfigKey, cfg)
	return nil
}

func UnlockAccount(c *cli.Context) error {
	client := c.Context.Value(core.ClientKey).(*client.Client)

	passphrase, err := prompt.AccountPassphrase().RunFlag(c, "passphrase")
	if err != nil {
		return err
	}

	client.Signer().Unlock(client.Account(), passphrase)

	return nil
}
