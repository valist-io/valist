package organization

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/command/utils/lifecycle"
	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
)

func NewThresholdCommand() *cli.Command {
	return &cli.Command{
		Name:      "threshold",
		Usage:     "Vote for organization threshold",
		ArgsUsage: "[org-name] [threshold]",
		Before:    lifecycle.UnlockAccount,
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			client := c.Context.Value(core.ClientKey).(*client.Client)
			orgName := c.Args().Get(0)

			threshold, err := strconv.ParseInt(c.Args().Get(1), 10, 64)
			if err != nil {
				return err
			}

			orgID, err := client.GetOrganizationID(c.Context, orgName)
			if err != nil {
				return err
			}

			vote, err := client.VoteOrganizationThreshold(c.Context, orgID, big.NewInt(threshold))
			if err != nil {
				return err
			}

			if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
				fmt.Printf("Voted to set threshold %d/%d\n", vote.SigCount, threshold)
			} else {
				fmt.Printf("Approved threshold %d\n", threshold)
			}

			return nil
		},
	}
}
