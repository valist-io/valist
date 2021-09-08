package repository

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/core/config"
)

func NewFetchCommand() *cli.Command {
	return &cli.Command{
		Name:    "fetch",
		Usage:   "Fetch repository info",
		Aliases: []string{"get"},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg := config.NewConfig(home)
			if err := cfg.Load(); err != nil {
				return err
			}

			var account accounts.Account
			if c.IsSet("account") {
				account.Address = common.HexToAddress(c.String("account"))
			} else {
				account.Address = cfg.Accounts.Default
			}

			client, err := core.NewClient(c.Context, cfg, account)
			if err != nil {
				return err
			}

			res, err := client.ResolvePath(c.Context, c.Args().Get(0))
			if err != nil {
				return err
			}

			meta, err := client.GetRepositoryMeta(c.Context, res.Repository.MetaCID)
			if err != nil {
				return err
			}

			fmt.Printf("OrgID: %s\n", res.Organization.ID.String())
			fmt.Printf("Name: %s\n", meta.Name)
			fmt.Printf("Description: %s\n", meta.Description)
			fmt.Printf("Signature Threshold: %d\n", res.Repository.Threshold)

			return nil
		},
	}
}
