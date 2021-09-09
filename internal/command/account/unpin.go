package account

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core/config"
)

func NewUnpinCommand() *cli.Command {
	return &cli.Command{
		Name:  "unpin",
		Usage: "Unpin an account from an external wallet",
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

			if !common.IsHexAddress(c.Args().Get(0)) {
				return fmt.Errorf("Invalid account address")
			}

			address := common.HexToAddress(c.Args().Get(0))
			// TODO do we care if the address is not pinned?

			var addresses []common.Address
			for _, pinned := range cfg.Accounts.Pinned {
				if pinned != address {
					addresses = append(addresses, pinned)
				}
			}

			cfg.Accounts.Pinned = addresses
			return cfg.Save()
		},
	}
}
