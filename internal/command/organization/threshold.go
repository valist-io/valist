package organization

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/config"
)

func NewThresholdCommand() *cli.Command {
	return &cli.Command{
		Name:  "threshold",
		Usage: "Vote for organization threshold",
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
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
