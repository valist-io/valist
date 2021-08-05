package account

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/signer"
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

			api, err := signer.NewSignerAPI(cfg)
			if err != nil {
				return err
			}
			server := core.NewUIServerAPI(api)

			accounts, err := server.ListAccounts(c.Context)
			if err != nil {
				return err
			}

			for _, acc := range accounts {
				fmt.Println(acc.Address.String())
			}

			return nil
		},
	}
}
