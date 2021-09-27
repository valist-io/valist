package organization

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/prompt"
)

func NewUpdateCommand() *cli.Command {
	return &cli.Command{
		Name:      "update",
		Usage:     "Update organization metadata",
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

			name, err := prompt.OrganizationName(meta.Name).Run()
			if err != nil {
				return err
			}

			desc, err := prompt.OrganizationDescription(meta.Description).Run()
			if err != nil {
				return err
			}

			homepage, err := prompt.OrganizationHomepage(meta.Homepage).Run()
			if err != nil {
				return err
			}

			meta.Name = name
			meta.Description = desc
			meta.Homepage = homepage

			_, err = client.SetOrganizationMeta(c.Context, orgID, meta)
			if err != nil {
				return err
			}

			fmt.Println("Organization updated!")
			return nil
		},
	}
}
