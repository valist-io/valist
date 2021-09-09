package organization

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/prompt"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an organization",
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

			orgName := c.Args().Get(0)
			_, err = client.GetOrganizationID(c.Context, orgName)
			if err == nil {
				return fmt.Errorf("Namespace '%v' taken. Please try another orgName/username.", orgName)
			}

			name, err := prompt.OrganizationName("").Run()
			if err != nil {
				return err
			}

			desc, err := prompt.OrganizationDescription("").Run()
			if err != nil {
				return err
			}

			orgMeta := types.OrganizationMeta{
				Name:        name,
				Description: desc,
			}

			fmt.Println("Creating organization...")

			create, err := client.CreateOrganization(c.Context, &orgMeta)
			if err != nil {
				return err
			}

			fmt.Printf("Linking name '%v' to orgID 0x'%x'...\n", orgName, create.OrgID)
			_, err = client.LinkOrganizationName(c.Context, create.OrgID, orgName)
			if err != nil {
				return err
			}

			fmt.Printf("Successfully created %v!\n", orgName)
			fmt.Printf("Your Valist ID: 0x%x\n", create.OrgID)

			return nil
		},
	}
}
