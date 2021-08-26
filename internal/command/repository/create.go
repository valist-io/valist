package repository

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/core/config"
	"github.com/valist-io/registry/internal/core/types"
	"github.com/valist-io/registry/internal/prompt"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create a repository",
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
			if c.IsSet("account") {
				account.Address = common.HexToAddress(c.String("account"))
			} else {
				account.Address = cfg.Accounts.Default
			}


			client, err := core.NewClient(c.Context, cfg, account)
			if err != nil {
				return err
			}
			defer client.Close()

			orgName := c.Args().Get(0)
			repoName := c.Args().Get(1)

			orgID, err := client.GetOrganizationID(c.Context, orgName)
			if err != nil {
				return err
			}

			name, err := prompt.RepositoryName("").Run()
			if err != nil {
				return err
			}

			desc, err := prompt.RepositoryDescription("").Run()
			if err != nil {
				return err
			}

			_, projectType, err := prompt.RepositoryProjectType().Run()
			if err != nil {
				return err
			}

			homepage, err := prompt.RepositoryHomepage("").Run()
			if err != nil {
				return err
			}

			url, err := prompt.RepositoryURL("").Run()
			if err != nil {
				return err
			}

			meta := types.RepositoryMeta{
				Name:        name,
				Description: desc,
				ProjectType: projectType,
				Homepage:    homepage,
				Repository:  url,
			}

			_, err = client.CreateRepository(c.Context, orgID, repoName, &meta)
			if err != nil {
				return err
			}

			fmt.Println("Repository created!")
			return nil
		},
	}
}
