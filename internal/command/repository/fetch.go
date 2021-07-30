package repository

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/impl"
)

func NewFetchCommand() *cli.Command {
	return &cli.Command{
		Name:  "fetch",
		Usage: "Fetch repository info",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "account",
				Value: "default",
				Usage: "Account to authenticate with",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg, err := config.Load(home)
			if err != nil {
				return err
			}

			var account accounts.Account
			if address, ok := cfg.Accounts[c.String("account")]; ok {
				account.Address = address
			} else {
				account.Address = common.HexToAddress(c.String("account"))
			}

			client, err := impl.NewClient(c.Context, cfg, account)
			if err != nil {
				return err
			}

			orgName := c.Args().Get(0)
			repoName := c.Args().Get(1)

			orgID, err := client.GetOrganizationID(c.Context, orgName)
			if err != nil {
				return err
			}

			repo, err := client.GetRepository(c.Context, orgID, repoName)
			if err != nil {
				return err
			}

			meta, err := client.GetRepositoryMeta(c.Context, repo.MetaCID)
			if err != nil {
				return err
			}

			fmt.Printf("OrgID: %s\n", orgID.String())
			fmt.Printf("Repo: %s/%s\n", orgName, repoName)
			fmt.Printf("Name: %s\n", meta.Name)
			fmt.Printf("Description: %s\n", meta.Description)
			fmt.Printf("Signature Threshold: %d\n", repo.Threshold)

			return nil
		},
	}
}
