package account

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core/config"
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

			cfg := config.NewConfig(home)
			if err := cfg.Load(); err != nil {
				return err
			}

			for _, account := range cfg.Accounts.Pinned {
				if cfg.Accounts.Default == account {
					fmt.Printf("%s (default)\n", account)
				} else {
					fmt.Printf("%s\n", account)
				}
			}

			return nil
		},
	}
}