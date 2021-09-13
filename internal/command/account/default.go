package account

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
)

func NewDefaultCommand() *cli.Command {
	return &cli.Command{
		Name:      "default",
		Usage:     "Set the default account",
		ArgsUsage: "[address]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			if !common.IsHexAddress(c.Args().Get(0)) {
				return fmt.Errorf("Invalid account address")
			}

			address := common.HexToAddress(c.Args().Get(0))
			// TODO do we care if the address is not pinned?

			config := c.Context.Value(core.ConfigKey).(*config.Config)
			config.Accounts.Default = address
			return config.Save()
		},
	}
}
