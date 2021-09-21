package repository

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
		Usage:     "Update repository metadata",
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

			name, err := prompt.RepositoryName(meta.Name).Run()
			if err != nil {
				return err
			}

			desc, err := prompt.RepositoryDescription(meta.Description).Run()
			if err != nil {
				return err
			}

			homepage, err := prompt.RepositoryHomepage(meta.Homepage).Run()
			if err != nil {
				return err
			}

			url, err := prompt.RepositoryURL(meta.Repository).Run()
			if err != nil {
				return err
			}

			meta.Name = name
			meta.Description = desc
			meta.Homepage = homepage
			meta.Repository = url

			_, err = client.SetRepositoryMeta(c.Context, res.OrgID, res.RepoName, meta)
			if err != nil {
				return err
			}

			fmt.Println("Repository updated!")
			return nil
		},
	}
}
