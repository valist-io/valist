package repo

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
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
