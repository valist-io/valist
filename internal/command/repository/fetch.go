package repository

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/impl"
)

func NewFetchCommand() *cli.Command {
	return &cli.Command{
		Name:  "fetch",
		Usage: "Fetch repository info",
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			orgName := c.Args().Get(0)
			repoName := c.Args().Get(1)

			client, err := impl.NewClient(c.Context, nil)
			if err != nil {
				return err
			}

			orgID, err := client.GetOrganizationID(c.Context, orgName)
			if err != nil {
				return err
			}

			repo, err := client.GetRepository(c.Context, orgID, repoName)
			if err != nil {
				return err
			}

			meta, err := client.GetRepositoryMeta(c.Context, repo.MetaCID)
			if err != nil {
				return err
			}

			fmt.Printf("OrgID: %s\n", orgID.String())
			fmt.Printf("Repo: %s/%s\n", orgName, repoName)
			fmt.Printf("Name: %s\n", meta.Name)
			fmt.Printf("Description: %s\n", meta.Description)
			fmt.Printf("Signature Threshold: %d\n", repo.Threshold)

			return nil
		},
	}
}
