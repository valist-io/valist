package organization

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/prompt"
)

func CreateOrg(client *client.Client, ctx context.Context, OrgName string) ([32]byte, error) {
	emptyID := [32]byte{}
	name, err := prompt.OrganizationName("").Run()
	if err != nil {
		return emptyID, err
	}

	desc, err := prompt.OrganizationDescription("").Run()
	if err != nil {
		return emptyID, err
	}

	orgMeta := types.OrganizationMeta{
		Name:        name,
		Description: desc,
	}

	fmt.Println("Creating organization...")

	create, err := client.CreateOrganization(ctx, &orgMeta)
	if err != nil {
		return emptyID, err
	}

	fmt.Printf("Linking name '%v' to orgID '0x%x'...\n", OrgName, create.OrgID)
	_, err = client.LinkOrganizationName(ctx, create.OrgID, OrgName)
	if err != nil {
		return emptyID, err
	}

	fmt.Printf("Successfully created %v!\n", OrgName)
	fmt.Printf("Your Valist ID: 0x%x\n", create.OrgID)
	return create.OrgID, nil
}

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

			_, err = CreateOrg(client, c.Context, orgName)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
