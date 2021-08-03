package organization

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/impl"
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
			if c.IsSet("account") {
				account.Address = common.HexToAddress(c.String("account"))
			} else {
				account.Address = cfg.Accounts.Default
			}

			client, err := impl.NewClient(c.Context, cfg, account)
			if err != nil {
				return err
			}

			// TODO prompt
			orgName := c.Args().Get(0)
			orgMeta := core.OrganizationMeta{
				Name:        "test",
				Description: "test",
			}

			fmt.Println("Creating organization...")
			createTx, err := client.CreateOrganization(c.Context, &orgMeta)
			if err != nil {
				return err
			}

			createRes := <-createTx
			if createRes.Err != nil {
				return createRes.Err
			}

			fmt.Println("Linking organization name...")
			linkTx, err := client.LinkOrganizationName(c.Context, createRes.OrgID, orgName)
			if err != nil {
				return err
			}

			linkRes := <-linkTx
			if linkRes.Err != nil {
				return linkRes.Err
			}

			fmt.Println("Success!")
			return nil
		},
	}
}
