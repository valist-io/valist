package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/prompt"
)

func CreateRepo(client *client.Client, context context.Context, OrgID common.Hash, RepoName string) error {
	name, err := prompt.RepositoryName("").Run()
	if err != nil {
		return err
	}

	desc, err := prompt.RepositoryDescription("").Run()
	if err != nil {
		return err
	}

	_, projectType, err := prompt.RepositoryProjectType().Run()
	if err != nil {
		return err
	}

	homepage, err := prompt.RepositoryHomepage("").Run()
	if err != nil {
		return err
	}

	url, err := prompt.RepositoryURL("").Run()
	if err != nil {
		return err
	}

	meta := types.RepositoryMeta{
		Name:        name,
		Description: desc,
		ProjectType: projectType,
		Homepage:    homepage,
		Repository:  url,
	}

	_, err = client.CreateRepository(context, OrgID, RepoName, &meta)
	if err != nil {
		return err
	}

	fmt.Println("Repository created!")
	return nil
}

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:      "create",
		Usage:     "Create a repository",
		ArgsUsage: "[repo-path]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			client := c.Context.Value(core.ClientKey).(*client.Client)
			res, err := client.ResolvePath(c.Context, c.Args().Get(0))
			if err != nil && !errors.Is(err, types.ErrRepositoryNotExist) {
				return err
			}

			err = CreateRepo(client, c.Context, res.OrgID, res.RepoName)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
