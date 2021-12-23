package command

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/types"
	"golang.org/x/mod/modfile"
)

func Publish(ctx context.Context, dryrun bool) error {
	client := ctx.Value(ClientKey).(*client.Client)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var valist Config
	if err := valist.Load(filepath.Join(cwd, "valist.yml")); err != nil {
		return err
	}

	// create will do nothing if org and repo already exist
	if err := Create(ctx, valist.Name); err != nil {
		return err
	}

	res, err := client.ResolvePath(ctx, valist.Name+"/"+valist.Tag)
	if err == nil {
		return fmt.Errorf("release %s already exists", res.ReleaseTag)
	}

	if err != types.ErrReleaseNotExist {
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

	releaseMeta := &types.ReleaseMeta{
		Name:         fmt.Sprintf("%s@%s", valist.Name, valist.Tag),
		Readme:       string(readme),
		Version:      valist.Tag,
		Dependencies: dependencies,
		Artifacts:    make(map[string]types.Artifact),
	}

	// TODO run file uploads in parallel and print progress
	for key, val := range valist.Artifacts {
		logger.Info("Adding: %s...", key)

		fdata, err := os.ReadFile(filepath.Join(cwd, val))
		if err != nil {
			return fmt.Errorf("failed to add %s: %v", key, err)
		}

		fpath, err := client.WriteFile(ctx, fdata)
		if err != nil {
			return fmt.Errorf("failed to add %s: %v", key, err)
		}

		releaseMeta.Artifacts[key] = types.Artifact{
			SHA256:   fmt.Sprintf("%x", sha256.Sum256(fdata)),
			Provider: fpath,
		}
	}

	releaseData, err := json.Marshal(releaseMeta)
	if err != nil {
		return err
	}

	releasePath, err := client.WriteFile(ctx, releaseData)
	if err != nil {
		return err
	}

	release := &types.Release{
		Tag:        valist.Tag,
		ReleaseCID: releasePath,
		MetaCID:    types.DeprecationNotice,
	}

	logger.Info("%s@%s", releaseMeta.Name, release.Tag)
	for name, artifact := range releaseMeta.Artifacts {
		logger.Info("- %s: %s", name, artifact.Provider)
	}

	if dryrun {
		return nil
	}

	_, err = client.VoteRelease(ctx, res.OrgID, res.RepoName, release)
	if err != nil {
		return err
	}

	logger.Info("Approved release %s", release.Tag)
	return nil
}
