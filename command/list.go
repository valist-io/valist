package command

import (
	"context"
	"fmt"

	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/types"
)

// List prints organization, repository, or release contents.
func List(ctx context.Context, rpath string) error {
	client := ctx.Value(ClientKey).(*client.Client)

	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}

	switch {
	case res.Release != nil:
		return listRelease(ctx, res.Release)
	case res.Repository != nil:
		return listRepository(ctx, res.Repository)
	case res.Organization != nil:
		return listOrganization(ctx, res.Organization)
	default:
		return fmt.Errorf("invalid path")
	}
}

func listRelease(ctx context.Context, release *types.Release) error {
	client := ctx.Value(ClientKey).(*client.Client)

	logger.Notice("Fetching from distributed storage...")
	meta, err := client.GetReleaseMeta(ctx, release.ReleaseCID)
	if err != nil {
		return err
	}

	logger.Info("Name:      %s", meta.Name)
	logger.Info("Tag:       %s", release.Tag)
	logger.Info("Artifacts:")

	for name, artifact := range meta.Artifacts {
		logger.Info("- %s", name)
		logger.Info("  %s", artifact.Provider)
		logger.Info("  %s", artifact.SHA256)
	}

	return nil
}

func listRepository(ctx context.Context, repo *types.Repository) error {
	client := ctx.Value(ClientKey).(*client.Client)

	logger.Notice("Fetching from distributed storage...")
	meta, err := client.GetRepositoryMeta(ctx, repo.MetaCID)
	if err != nil {
		return err
	}

	members, err := client.GetRepositoryMembers(ctx, repo.OrgID, repo.Name)
	if err != nil {
		return err
	}

	logger.Info("Name:        %s", meta.Name)
	logger.Info("Description: %s", meta.Description)
	logger.Info("Homepage:    %s", meta.Homepage)
	logger.Info("Source:      %s", meta.Repository)

	logger.Info("Members:")
	for _, address := range members {
		logger.Info("- %s", address)
	}

	logger.Info("Releases:")
	return client.ListReleaseTags(repo.OrgID, repo.Name).ForEach(ctx, func(tag string) {
		logger.Info("- %s", tag)
	})
}

func listOrganization(ctx context.Context, org *types.Organization) error {
	client := ctx.Value(ClientKey).(*client.Client)

	logger.Notice("Fetching from distributed storage...")
	meta, err := client.GetOrganizationMeta(ctx, org.MetaCID)
	if err != nil {
		return err
	}

	members, err := client.GetOrganizationMembers(ctx, org.ID)
	if err != nil {
		return err
	}

	logger.Info("Name:        %s", meta.Name)
	logger.Info("Description: %s", meta.Description)
	logger.Info("Homepage:    %s", meta.Homepage)

	logger.Info("Members:")
	for _, address := range members {
		logger.Info("- %s", address)
	}

	return nil
}
