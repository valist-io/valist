package account

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
)

func NewListCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all accounts",
		Action: func(c *cli.Context) error {
			keyStore, err := config.GetKeyStore()
			if err != nil {
				return err
			}

			for _, acc := range keyStore.Accounts() {
				fmt.Println(acc.Address)
			}

			return nil
		},
	}
}
