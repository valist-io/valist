package organization

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
		Usage: "Update organization metadata",
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

			client, err := core.NewClient(c.Context, cfg, account, c.String("passphrase"))
			if err != nil {
				return err
			}

			orgName := c.Args().Get(0)

			orgID, err := client.GetOrganizationID(c.Context, orgName)
			if err != nil {
				return err
			}

			org, err := client.GetOrganization(c.Context, orgID)
			if err != nil {
				return err
			}

			meta, err := client.GetOrganizationMeta(c.Context, org.MetaCID)
			if err != nil {
				return err
			}

			name, err := prompt.OrganizationName(meta.Name).Run()
			if err != nil {
				return err
			}

			desc, err := prompt.OrganizationDescription(meta.Description).Run()
			if err != nil {
				return err
			}

			homepage, err := prompt.OrganizationHomepage(meta.Homepage).Run()
			if err != nil {
				return err
			}

			meta.Name = name
			meta.Description = desc
			meta.Homepage = homepage

			_, err = client.SetOrganizationMeta(c.Context, orgID, meta)
			if err != nil {
				return err
			}

			fmt.Println("Organization updated!")
			return nil
		},
	}
}
