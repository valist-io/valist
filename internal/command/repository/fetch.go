package repository

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
)

func NewFetchCommand() *cli.Command {
	return &cli.Command{
		Name:      "fetch",
		Usage:     "Fetch repository info",
		Aliases:   []string{"get"},
		ArgsUsage: "[repo-path]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			client := c.Context.Value(core.ClientKey).(*client.Client)
			res, err := client.ResolvePath(c.Context, c.Args().Get(0))
			if err != nil {
				return err
			}

			meta, err := client.GetRepositoryMeta(c.Context, res.Repository.MetaCID)
			if err != nil {
				return err
			}

			fmt.Printf("OrgID: %s\n", res.Organization.ID.String())
			fmt.Printf("Name: %s\n", meta.Name)
			fmt.Printf("Description: %s\n", meta.Description)
			fmt.Printf("Signature Threshold: %d\n", res.Repository.Threshold)

			return nil
		},
	}
}
