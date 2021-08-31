package organization

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/core/client"
	"github.com/valist-io/registry/internal/core/config"
)

func voteOrganizationAdmin(c *cli.Context, operation common.Hash) (*valist.ValistVoteKeyEvent, error) {
	if c.NArg() != 2 {
		cli.ShowSubcommandHelpAndExit(c, 1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cfg, err := config.Load(home)
	if err != nil {
		return nil, err
	}

	var account accounts.Account
	if c.IsSet("account") {
		account.Address = common.HexToAddress(c.String("account"))
	} else {
		account.Address = cfg.Accounts.Default
	}

	valist, err := core.NewClient(c.Context, cfg, account)
	if err != nil {
		return nil, err
	}
	defer valist.Close()

	// @TODO: Add extra validation here
	orgName := c.Args().Get(0)

	orgID, err := valist.GetOrganizationID(c.Context, orgName)
	if err != nil {
		return nil, err
	}

	if !common.IsHexAddress(c.Args().Get(1)) {
		return nil, fmt.Errorf("Invalid address: %s", c.Args().Get(1))
	}

	address := common.HexToAddress(c.Args().Get(1))

	return valist.VoteOrganizationAdmin(c.Context, orgID, operation, address)
}

func NewKeyCommand() *cli.Command {
	return &cli.Command{
		Name:  "key",
		Usage: "Manage keys at an organization level",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Add a new key to an organization",
				Action: func(c *cli.Context) error {
					fmt.Println("Adding key to organization...")
					vote, err := voteOrganizationAdmin(c, client.ADD_KEY)
					if err != nil {
						return err
					}

					if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
						fmt.Printf("Voted to add key, %d/%d\n votes", vote.SigCount, vote.Threshold)
					} else {
						fmt.Printf("Key successfully approved!")
					}
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "Remove a key from an organization",
				Action: func(c *cli.Context) error {
					fmt.Println("Removing key to organization...")
					vote, err := voteOrganizationAdmin(c, client.REVOKE_KEY)
					if err != nil {
						return err
					}

					if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
						fmt.Printf("Voted to remove key, %d/%d\n votes", vote.SigCount, vote.Threshold)
					} else {
						fmt.Printf("Key successfully revoked!")
					}
					return nil
				},
			},
			{
				Name:  "rotate",
				Usage: "Replace a key on an organization",
				Action: func(c *cli.Context) error {
					fmt.Println("Rotating key on organization...")
					_, err := voteOrganizationAdmin(c, client.ROTATE_KEY)
					if err != nil {
						return err
					}
					fmt.Printf("Key successfully rotated!")
					return nil
				},
			},
		},
	}
}
