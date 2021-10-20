package command

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/prompt"
)

// Create creates a new organization or repository.
func Create(ctx context.Context, rpath string) error {
	client := ctx.Value(ClientKey).(*client.Client)

	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}

	switch {
	case res.Repository == nil:
		return createRepository(ctx, res.OrgID, res.RepoName)
	case res.Organization == nil:
		return createOrganization(ctx, res.OrgName)
	default:
		return fmt.Errorf("invalid path")
	}
}

func createOrganization(ctx context.Context, orgName string) error {
	client := ctx.Value(ClientKey).(*client.Client)

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
	create, err := client.CreateOrganization(ctx, &orgMeta)
	if err != nil {
		return err
	}

	fmt.Printf("Linking name '%s' to organization ID '0x%x'...\n", orgName, create.OrgID)
	_, err = client.LinkOrganizationName(ctx, create.OrgID, orgName)
	if err != nil {
		return err
	}

	fmt.Println("Organization created!")
	return nil
}

func createRepository(ctx context.Context, orgID common.Hash, repoName string) error {
	client := ctx.Value(ClientKey).(*client.Client)

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

	_, err = client.CreateRepository(ctx, orgID, repoName, &meta)
	if err != nil {
		return err
	}

	fmt.Println("Repository created!")
	return nil
}
