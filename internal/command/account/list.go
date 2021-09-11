package account

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
)

func NewListCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all accounts",
		Action: func(c *cli.Context) error {
			config := c.Context.Value(core.ConfigKey).(*config.Config)

			for _, account := range config.Accounts.Pinned {
				if config.Accounts.Default == account {
					fmt.Printf("%s (default)\n", account)
				} else {
					fmt.Printf("%s\n", account)
				}
			}

			return nil
		},
	}
}
