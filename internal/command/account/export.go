package account

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/prompt"
)

func NewExportCommand() *cli.Command {
	return &cli.Command{
		Name:  "export",
		Usage: "Export an account",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg := config.NewConfig(home)
			if err := cfg.Load(); err != nil {
				return err
			}

			address := common.HexToAddress(c.Args().Get(0))
			account := accounts.Account{Address: address}

			passphrase, err := prompt.AccountPassphrase().Run()
			if err != nil {
				return err
			}

			newPassphrase, err := prompt.NewAccountPassphrase().Run()
			if err != nil {
				return err
			}

			json, err := cfg.KeyStore().Export(account, passphrase, newPassphrase)
			if err != nil {
				return err
			}

			fmt.Println(string(json))
			return nil
		},
	}
}
