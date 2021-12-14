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

	fmt.Println("Fetching from distributed storage...")
	meta, err := client.GetReleaseMeta(ctx, release.ReleaseCID)
	if err != nil {
		return err
	}

	fmt.Printf("%s@%s\n", meta.Name, release.Tag)

	for name, artifact := range meta.Artifacts {
		fmt.Printf("- %s: %s\n", name, artifact.Provider)
	}

	return nil
}

func getRepository(ctx context.Context, repo *types.Repository) error {
	client := ctx.Value(ClientKey).(*client.Client)

	fmt.Println("Fetching from distributed storage...")
	meta, err := client.GetRepositoryMeta(ctx, repo.MetaCID)
	if err != nil {
		return err
	}

	fmt.Printf("Name:        %s\n", meta.Name)
	fmt.Printf("Description: %s\n", meta.Description)
	fmt.Printf("Homepage:    %s\n", meta.Homepage)
	fmt.Printf("Source code repo:  %s\n", meta.Repository)

	return nil
}

func getOrganization(ctx context.Context, org *types.Organization) error {
	client := ctx.Value(ClientKey).(*client.Client)

	fmt.Println("Fetching from distributed storage...")
	meta, err := client.GetOrganizationMeta(ctx, org.MetaCID)
	if err != nil {
		return err
	}

	fmt.Printf("Name:        %s\n", meta.Name)
	fmt.Printf("Description: %s\n", meta.Description)
	fmt.Printf("Homepage:    %s\n", meta.Homepage)

	return nil
}
