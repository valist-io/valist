package account

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/core/config"
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

			api, _, err := signer.NewSigner(cfg)
			if err != nil {
				return err
			}

			address, err := api.New(c.Context)
			if err != nil {
				return err
			}

			if cfg.Accounts.Default == common.HexToAddress("0x0") {
				cfg.Accounts.Default = address
			}

			cfg.Accounts.Pinned = append(cfg.Accounts.Pinned, address)
			return cfg.Save()
		},
	}
}
