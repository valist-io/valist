package account

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
)

func NewPinCommand() *cli.Command {
	return &cli.Command{
		Name:  "pin",
		Usage: "Pin an account from an external wallets",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg, err := config.Load(home)
			if err != nil {
				return err
			}

			if !common.IsHexAddress(c.Args().Get(0)) {
				return fmt.Errorf("Invalid account address")
			}

			address := common.HexToAddress(c.Args().Get(0))
			for _, pinned := range cfg.Accounts.Pinned {
				if pinned == address {
					return fmt.Errorf("Account already pinned")
				}
			}

			if cfg.Accounts.Default == common.HexToAddress("0x0") {
				cfg.Accounts.Default = address
			}

			cfg.Accounts.Pinned = append(cfg.Accounts.Pinned, address)
			return cfg.Save()
		},
	}
}
