package organization

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/prompt"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:      "create",
		Usage:     "Create an organization",
		ArgsUsage: "[org-name]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			client := c.Context.Value(core.ClientKey).(*client.Client)
			orgName := c.Args().Get(0)

			_, err := client.GetOrganizationID(c.Context, orgName)
			if err == nil {
				return fmt.Errorf("Namespace '%v' taken. Please try another orgName/username.", orgName)
			}

			if err != nil {
				return err
			}

			name, err := prompt.OrganizationName("").Run()
			if err != nil {
				return err
			}

			desc, err := prompt.OrganizationDescription("").Run()
			if err != nil {
				return err
			}

			orgMeta := types.OrganizationMeta{
				Name:        name,
				Description: desc,
			}

			fmt.Println("Creating organization...")

			create, err := client.CreateOrganization(c.Context, &orgMeta)
			if err != nil {
				return err
			}

			fmt.Printf("Linking name '%v' to orgID 0x'%x'...\n", orgName, create.OrgID)
			_, err = client.LinkOrganizationName(c.Context, create.OrgID, orgName)
			if err != nil {
				return err
			}

			fmt.Printf("Successfully created %v!\n", orgName)
			fmt.Printf("Your Valist ID: 0x%x\n", create.OrgID)

			return nil
		},
	}
}
