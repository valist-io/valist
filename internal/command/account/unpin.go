package account

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
)

func NewUnpinCommand() *cli.Command {
	return &cli.Command{
		Name:      "unpin",
		Usage:     "Unpin an account from an external wallet",
		ArgsUsage: "[address]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			if !common.IsHexAddress(c.Args().Get(0)) {
				return fmt.Errorf("Invalid account address")
			}

			config := c.Context.Value(core.ConfigKey).(*config.Config)
			address := common.HexToAddress(c.Args().Get(0))
			// TODO do we care if the address is not pinned?

			var addresses []common.Address
			for _, pinned := range config.Accounts.Pinned {
				if pinned != address {
					addresses = append(addresses, pinned)
				}
			}

			config.Accounts.Pinned = addresses
			return config.Save()
		},
	}
}
