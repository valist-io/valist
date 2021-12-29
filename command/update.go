package command

import (
	"context"
	"fmt"

	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/types"
	"github.com/valist-io/valist/prompt"
)

// Update updates organization or repository metadata.
func Update(ctx context.Context, rpath string) error {
	client := ctx.Value(ClientKey).(*client.Client)

	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}

	switch {
	case res.Repository != nil:
		return updateRepository(ctx, res.Repository)
	case res.Organization != nil:
		return updateOrganization(ctx, res.Organization)
	default:
		return fmt.Errorf("invalid path")
	}
}

func updateOrganization(ctx context.Context, org *types.Organization) error {
	client := ctx.Value(ClientKey).(*client.Client)

	logger.Notice("Fetching from distributed storage...")
	meta, err := client.GetOrganizationMeta(ctx, org.MetaCID)
	if err != nil {
		return err
	}

	meta.Name, err = prompt.OrganizationName(meta.Name).Run()
	if err != nil {
		return err
	}

	meta.Description, err = prompt.OrganizationDescription(meta.Description).Run()
	if err != nil {
		return err
	}

	meta.Homepage, err = prompt.OrganizationHomepage(meta.Homepage).Run()
	if err != nil {
		return err
	}

	_, err = client.SetOrganizationMeta(ctx, org.ID, meta)
	if err != nil {
		return err
	}

	logger.Info("Organization updated!")
	return nil
}

func updateRepository(ctx context.Context, repo *types.Repository) error {
	client := ctx.Value(ClientKey).(*client.Client)

	logger.Notice("Fetching from distributed storage...")
	meta, err := client.GetRepositoryMeta(ctx, repo.MetaCID)
	if err != nil {
		return err
	}

	meta.Name, err = prompt.RepositoryName(meta.Name).Run()
	if err != nil {
		return err
	}

	meta.Description, err = prompt.RepositoryDescription(meta.Description).Run()
	if err != nil {
		return err
	}

	meta.Homepage, err = prompt.RepositoryHomepage(meta.Homepage).Run()
	if err != nil {
		return err
	}

	meta.Repository, err = prompt.RepositoryURL(meta.Repository).Run()
	if err != nil {
		return err
	}

	_, err = client.SetRepositoryMeta(ctx, repo.OrgID, repo.Name, meta)
	if err != nil {
		return err
	}

	logger.Info("Repository updated!")
	return nil
}
