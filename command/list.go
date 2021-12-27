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

	fmt.Println("Fetching from distributed storage...")
	meta, err := client.GetReleaseMeta(ctx, release.ReleaseCID)
	if err != nil {
		return err
	}

	fmt.Printf("Name:     %s\n", meta.Name)
	fmt.Printf("Tag:      %s\n", release.Tag)
	fmt.Printf("Artifacts:  \n")

	for name, artifact := range meta.Artifacts {
		fmt.Printf("- %s\n", name)
		fmt.Printf("  %s\n", artifact.Provider)
		fmt.Printf("  %s\n", artifact.SHA256)
	}

	return nil
}

func listRepository(ctx context.Context, repo *types.Repository) error {
	client := ctx.Value(ClientKey).(*client.Client)

	fmt.Println("Fetching from distributed storage...")
	meta, err := client.GetRepositoryMeta(ctx, repo.MetaCID)
	if err != nil {
		return err
	}

	fmt.Printf("Name:        %s\n", meta.Name)
	fmt.Printf("Description: %s\n", meta.Description)
	fmt.Printf("Homepage:    %s\n", meta.Homepage)
	fmt.Printf("Source:      %s\n", meta.Repository)
	fmt.Printf("Releases:      \n")

	iter := client.ListReleaseTags(repo.OrgID, repo.Name)
	return iter.ForEach(ctx, func(tag string) {
		fmt.Printf("- %s\n", tag)
	})
}

func listOrganization(ctx context.Context, org *types.Organization) error {
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
