package organization

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/impl"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an organization",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "account",
				Value: "default",
				Usage: "account to authorize transaction",
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

			address, ok := cfg.Accounts[c.String("account")]
			if !ok {
				address = common.HexToAddress(c.String("account"))
			}

			account, err := cfg.KeyStore().Find(accounts.Account{Address: address})
			if err != nil {
				return err
			}

			transact := func() (*bind.TransactOpts, error) {
				return bind.NewKeyStoreTransactorWithChainID(cfg.KeyStore(), account, cfg.Ethereum.ChainID)
			}

			client, err := impl.NewClient(c.Context, transact)
			if err != nil {
				return err
			}

			// TODO prompt
			orgName := c.Args().Get(0)
			orgMeta := core.OrganizationMeta{
				Name:        "",
				Description: "",
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
