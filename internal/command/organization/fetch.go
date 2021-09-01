package organization

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
		Usage:   "Fetch organization info",
		Aliases: []string{"get"},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
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

			fmt.Printf("OrgID: %s\n", orgID.String())
			fmt.Printf("Name: %s\n", meta.Name)
			fmt.Printf("Homepage: %s\n", meta.Homepage)
			fmt.Printf("Description: %s\n", meta.Description)
			fmt.Printf("Signature Threshold: %d\n", org.Threshold)

			return nil
		},
	}
}
