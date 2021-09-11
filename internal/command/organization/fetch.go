package organization

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
)

func NewFetchCommand() *cli.Command {
	return &cli.Command{
		Name:      "fetch",
		Usage:     "Fetch organization info",
		Aliases:   []string{"get"},
		ArgsUsage: "[org-name]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			client := c.Context.Value(core.ClientKey).(*client.Client)
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
