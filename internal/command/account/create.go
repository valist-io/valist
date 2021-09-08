package account

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/core/config"
	"github.com/valist-io/registry/internal/prompt"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an account",
		Action: func(c *cli.Context) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg := config.NewConfig(home)
			if err := cfg.Load(); err != nil {
				return err
			}

			passphrase, err := prompt.AccountPassphrase().Run()
			if err != nil {
				return err
			}

			account, err := cfg.KeyStore().NewAccount(passphrase)
			if err != nil {
				return err
			}

			if cfg.Accounts.Default == common.HexToAddress("0x0") {
				cfg.Accounts.Default = account.Address
			}

			cfg.Accounts.Pinned = append(cfg.Accounts.Pinned, account.Address)
			return cfg.Save()
		},
	}
}
