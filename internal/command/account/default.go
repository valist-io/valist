package account

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/core/config"
)

func NewDefaultCommand() *cli.Command {
	return &cli.Command{
		Name:  "default",
		Usage: "Set the default account",
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

			cfg.Accounts.Default = address
			return cfg.Save()
		},
	}
}
