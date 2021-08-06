package organization

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/core/impl"
	"github.com/valist-io/registry/internal/signer"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an organization",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "account",
				Value: "default",
				Usage: "Account to authenticate with",
			},
		},
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

			listener, _, err := signer.StartIPCEndpoint(cfg)
			if err != nil {
				return err
			}
			defer listener.Close()

			var account accounts.Account
			if address, ok := cfg.Accounts[c.String("account")]; ok {
				account.Address = address
			} else {
				account.Address = common.HexToAddress(c.String("account"))
			}

			client, err := impl.NewClientWithMetaTx(c.Context, cfg, account)
			if err != nil {
				return err
			}

			// TODO prompt
			orgName := c.Args().Get(0)
			orgMeta := core.OrganizationMeta{
				Name:        "test11",
				Description: "test",
			}

			_, err = client.GetOrganizationID(c.Context, orgName)
			if err == nil {
				return fmt.Errorf("Namespace '%v' taken. Please try another orgName/username.", orgName)
			}

			fmt.Println("Creating organization...")
			create, err := client.CreateOrganization(c.Context, &orgMeta)
			if err != nil {
				return err
			}

			fmt.Printf("Linking name '%v' to orgID '%v'...\n", orgName, create.OrgID)
			_, err = client.LinkOrganizationName(c.Context, create.OrgID, orgName)
			if err != nil {
				return err
			}

			fmt.Println("Success!")
			return nil
		},
	}
}
