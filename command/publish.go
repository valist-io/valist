package command

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/valist-io/valist"
	"github.com/valist-io/valist/core"
	"golang.org/x/mod/modfile"
)

func Publish(ctx context.Context, dryrun bool) error {
	client := ctx.Value(ClientKey).(*core.Client)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var build Config
	if err := build.Load(filepath.Join(cwd, "valist.yml")); err != nil {
		return err
	}
	res, err := client.ResolvePath(ctx, build.Name+"/"+build.Tag)
	if err == nil {
		return fmt.Errorf("release %s already exists", res.ReleaseName)
	}
	if err != valist.ErrReleaseNotExist {
		return err
	}

	var dependencies []string
	if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
		goModData, err := os.ReadFile(filepath.Join(cwd, "go.mod"))
		if err != nil {
			return err
		}
		modFile, err := modfile.Parse("go.mod", goModData, nil)
		if err != nil {
			return err
		}
		for _, url := range modFile.Require {
			dependencies = append(dependencies, url.Mod.String())
		}
	}

	// TODO replace with regex or path matcher
	readme, err := os.ReadFile("README.md")
	if err != nil {
		logger.Warn("readme not found")
	}

	release := &valist.Release{
		Name:         fmt.Sprintf("%s@%s", build.Name, build.Tag),
		Readme:       string(readme),
		Version:      build.Tag,
		Dependencies: dependencies,
		Artifacts:    make(map[string]valist.Artifact),
	}

	// TODO run file uploads in parallel and print progress
	for key, val := range build.Artifacts {
		logger.Info("Adding: %s...", key)

		fdata, err := os.ReadFile(filepath.Join(cwd, val))
		if err != nil {
			return fmt.Errorf("failed to add %s: %v", key, err)
		}

		fpath, err := client.WriteBytes(ctx, fdata)
		if err != nil {
			return fmt.Errorf("failed to add %s: %v", key, err)
		}

		release.Artifacts[key] = valist.Artifact{
			SHA256:   fmt.Sprintf("%x", sha256.Sum256(fdata)),
			Provider: fpath,
		}
	}

	releaseData, err := json.Marshal(release)
	if err != nil {
		return err
	}
	releasePath, err := client.WriteBytes(ctx, releaseData)
	if err != nil {
		return err
	}

	logger.Info("%s@%s", release.Name, build.Tag)
	for name, artifact := range release.Artifacts {
		logger.Info("- %s: %s", name, artifact.Provider)
	}
	if dryrun {
		return nil
	}

	err = client.CreateRelease(ctx, res.TeamName, res.ProjectName, build.Tag, releasePath)
	if err != nil {
		return err
	}
	logger.Info("Created release %s", build.Tag)
	return nil
}
