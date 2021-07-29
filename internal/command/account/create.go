package account

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/signer"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an account",
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

			address, err := server.New(c.Context)
			if err != nil {
				return err
			}
			fmt.Println("Account created:", address)

			if _, ok := cfg.Accounts["default"]; ok {
				return nil
			}

			fmt.Println("Setting default account")
			cfg.Accounts["default"] = address
			return cfg.Save()
		},
	}
}
