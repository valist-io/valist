package command

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/valist-io/valist"
	"github.com/valist-io/valist/core"
	"github.com/valist-io/valist/publish"
)

func Publish(ctx context.Context, dryrun bool) error {
	client := ctx.Value(ClientKey).(*core.Client)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var pub publish.Config
	if err := pub.Load(filepath.Join(cwd, "valist.yml")); err != nil {
		return err
	}
	res, err := client.ResolvePath(ctx, pub.Name+"/"+pub.Tag)
	if err == nil {
		return fmt.Errorf("release %s already exists", res.ReleaseName)
	}
	if err != valist.ErrReleaseNotExist {
		return err
	}

	var artifacts = make(map[string]valist.Artifact)
	for key, val := range pub.Artifacts {
		logger.Info("Adding: %s...", key)
		fpath, err := client.WriteFile(ctx, filepath.Join(cwd, val))
		if err != nil {
			return fmt.Errorf("failed to add %s: %v", key, err)
		}
		// TODO find a way to sha256 directories
		artifacts[key] = valist.Artifact{
			Provider: fpath,
		}
	}

	readme, err := publish.Readme(cwd)
	if err != nil {
		logger.Warn("readme not found")
	}
	dependencies, err := publish.Dependencies(cwd)
	if err != nil {
		logger.Warn("dependencies not found")
	}

	release := &valist.Release{
		Name:         fmt.Sprintf("%s@%s", pub.Name, pub.Tag),
		Readme:       string(readme),
		Version:      pub.Tag,
		Dependencies: dependencies,
		Artifacts:    artifacts,
	}

	releaseData, err := json.Marshal(release)
	if err != nil {
		return err
	}
	releasePath, err := client.WriteBytes(ctx, releaseData)
	if err != nil {
		return err
	}

	logger.Info("%s@%s", release.Name, pub.Tag)
	for name, artifact := range release.Artifacts {
		logger.Info("- %s: %s", name, artifact.Provider)
	}
	if dryrun {
		return nil
	}

	err = client.CreateRelease(ctx, res.TeamName, res.ProjectName, pub.Tag, releasePath)
	if err != nil {
		return err
	}
	logger.Info("Created release %s", pub.Tag)
	return nil
}
