package repository

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/prompt"
)

func NewUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update repository metadata",
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
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

			client, err := core.NewClient(c.Context, cfg, account, c.String("passphrase"))
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

			name, err := prompt.RepositoryName(meta.Name).Run()
			if err != nil {
				return err
			}

			desc, err := prompt.RepositoryDescription(meta.Description).Run()
			if err != nil {
				return err
			}

			homepage, err := prompt.RepositoryHomepage(meta.Homepage).Run()
			if err != nil {
				return err
			}

			url, err := prompt.RepositoryURL(meta.Repository).Run()
			if err != nil {
				return err
			}

			meta.Name = name
			meta.Description = desc
			meta.Homepage = homepage
			meta.Repository = url

			_, err = client.SetRepositoryMeta(c.Context, orgID, repoName, meta)
			if err != nil {
				return err
			}

			fmt.Println("Repository updated!")
			return nil
		},
	}
}
