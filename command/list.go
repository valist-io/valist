package command

import (
	"context"
	"fmt"

	"github.com/valist-io/valist/core"
)

// List prints organization, repository, or release contents.
func List(ctx context.Context, rpath string) error {
	client := ctx.Value(ClientKey).(*core.Client)

	logger.Notice("Fetching from distributed storage...")
	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}

	switch {
	case res.Release != nil:
		logger.Info("Name:      %s", res.Release.Name)
		logger.Info("Tag:       %s", res.ReleaseName)
		logger.Info("Artifacts:")
		// print all artifacts
		for name, artifact := range res.Release.Artifacts {
			logger.Info("- %s", name)
			logger.Info("  %s", artifact.Provider)
			logger.Info("  %s", artifact.SHA256)
		}
	case res.Project != nil:
		logger.Info("Name:        %s", res.Project.Name)
		logger.Info("Description: %s", res.Project.Description)
		logger.Info("Homepage:    %s", res.Project.Homepage)
		logger.Info("Source:      %s", res.Project.Repository)
		// TODO print project members
		// TODO print project releases
	case res.Team != nil:
		logger.Info("Name:        %s", res.Team.Name)
		logger.Info("Description: %s", res.Team.Description)
		logger.Info("Homepage:    %s", res.Team.Homepage)
		// TODO print team members
	default:
		return fmt.Errorf("invalid path")
	}

	return nil
}
