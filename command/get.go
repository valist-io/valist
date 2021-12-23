package command

import (
	"context"
	"fmt"

	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/types"
)

// Get prints organization, repository, or release metadata.
func Get(ctx context.Context, rpath string) error {
	client := ctx.Value(ClientKey).(*client.Client)

	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}

	switch {
	case res.Release != nil:
		return getRelease(ctx, res.Release)
	case res.Repository != nil:
		return getRepository(ctx, res.Repository)
	case res.Organization != nil:
		return getOrganization(ctx, res.Organization)
	default:
		return fmt.Errorf("invalid path")
	}
}

func getRelease(ctx context.Context, release *types.Release) error {
	client := ctx.Value(ClientKey).(*client.Client)

	logger.Info("Fetching from distributed storage...")
	meta, err := client.GetReleaseMeta(ctx, release.ReleaseCID)
	if err != nil {
		return err
	}

	logger.Info("%s@%s", meta.Name, release.Tag)
	for name, artifact := range meta.Artifacts {
		logger.Info("- %s: %s", name, artifact.Provider)
	}

	return nil
}

func getRepository(ctx context.Context, repo *types.Repository) error {
	client := ctx.Value(ClientKey).(*client.Client)

	logger.Info("Fetching from distributed storage...")
	meta, err := client.GetRepositoryMeta(ctx, repo.MetaCID)
	if err != nil {
		return err
	}

	logger.Info("Name:        %s", meta.Name)
	logger.Info("Description: %s", meta.Description)
	logger.Info("Homepage:    %s", meta.Homepage)
	logger.Info("Source code repo:  %s", meta.Repository)

	return nil
}

func getOrganization(ctx context.Context, org *types.Organization) error {
	client := ctx.Value(ClientKey).(*client.Client)

	logger.Info("Fetching from distributed storage...")
	meta, err := client.GetOrganizationMeta(ctx, org.MetaCID)
	if err != nil {
		return err
	}

	logger.Info("Name:        %s", meta.Name)
	logger.Info("Description: %s", meta.Description)
	logger.Info("Homepage:    %s", meta.Homepage)

	return nil
}
