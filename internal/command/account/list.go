package account

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
)

func NewListCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all accounts",
		Action: func(c *cli.Context) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg, err := config.Load(home)
			if err != nil {
				return err
			}

			for _, acc := range cfg.KeyStore().Accounts() {
				fmt.Println(acc.Address)
			}

			return nil
		},
	}
}
