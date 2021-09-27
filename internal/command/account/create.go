package account

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/prompt"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an account",
		Action: func(c *cli.Context) error {
			config := c.Context.Value(core.ConfigKey).(*config.Config)

			passphrase, err := prompt.NewAccountPassphrase().Run()
			if err != nil {
				return err
			}

			account, err := config.KeyStore().NewAccount(passphrase)
			if err != nil {
				return err
			}

			if config.Accounts.Default == common.HexToAddress("0x0") {
				config.Accounts.Default = account.Address
			}

			config.Accounts.Pinned = append(config.Accounts.Pinned, account.Address)
			return config.Save()
		},
	}
}
