package account

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
)

func NewPinCommand() *cli.Command {
	return &cli.Command{
		Name:      "pin",
		Usage:     "Pin an account from an external wallets",
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

			for _, pinned := range config.Accounts.Pinned {
				if pinned == address {
					return fmt.Errorf("Account already pinned")
				}
			}

			if config.Accounts.Default == common.HexToAddress("0x0") {
				config.Accounts.Default = address
			}

			config.Accounts.Pinned = append(config.Accounts.Pinned, address)
			return config.Save()
		},
	}
}
