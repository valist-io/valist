package key

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "key",
		Usage: "Manage keys at an organization level",
		Subcommands: []*cli.Command{
			NewAddCommand(),
			NewRevokeCommand(),
			NewRotateCommand(),
		},
	}
}

func NewAddCommand() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "Add a new key to an organization",
		ArgsUsage: "[org-name] [address]",
		Action:    action,
	}
}

func NewRevokeCommand() *cli.Command {
	return &cli.Command{
		Name:      "revoke",
		Usage:     "Remove a key from an organization",
		ArgsUsage: "[org-name] [address]",
		Action:    action,
	}
}

func NewRotateCommand() *cli.Command {
	return &cli.Command{
		Name:      "rotate",
		Usage:     "Replace a key on an organization",
		ArgsUsage: "[org-name] [address]",
		Action:    action,
	}
}

func action(c *cli.Context) error {
	if c.NArg() != 2 {
		cli.ShowSubcommandHelpAndExit(c, 1)
	}

	var operation common.Hash
	switch c.Command.Name {
	case "add":
		fmt.Println("Adding key...")
		operation = client.ADD_KEY
	case "rotate":
		fmt.Println("Rotating key...")
		operation = client.ROTATE_KEY
	case "revoke":
		fmt.Println("Revoking key...")
		operation = client.REVOKE_KEY
	}

	client := c.Context.Value(core.ClientKey).(*client.Client)
	orgID, err := client.GetOrganizationID(c.Context, c.Args().Get(0))
	if err != nil {
		return err
	}

	if !common.IsHexAddress(c.Args().Get(1)) {
		return fmt.Errorf("Invalid address: %s", c.Args().Get(1))
	}

	address := common.HexToAddress(c.Args().Get(1))
	vote, err := client.VoteOrganizationAdmin(c.Context, orgID, operation, address)
	if err != nil {
		return err
	}

	if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
		fmt.Printf("Pending %d/%d votes\n", vote.SigCount, vote.Threshold)
	} else {
		fmt.Println("Approved!")
	}

	return nil
}
