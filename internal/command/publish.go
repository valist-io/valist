package command

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/types"
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

	res, err := client.ResolvePath(ctx, valist.Name+"/"+valist.Tag)
	if err == nil {
		return fmt.Errorf("release %s already exists", res.ReleaseTag)
	}

	if err != types.ErrReleaseNotExist {
		return err
	}

	// TODO replace with regex or path matcher
	readme, _ := os.ReadFile("README1.md") //nolint:errcheck

	releaseMeta := &types.ReleaseMeta{
		Name:      fmt.Sprintf("%s@%s", valist.Name, valist.Tag),
		Readme:    string(readme),
		Artifacts: make(map[string]types.Artifact),
	}

	for key, val := range valist.Artifacts {
		fdata, err := os.ReadFile(filepath.Join(cwd, val))
		if err != nil {
			return fmt.Errorf("failed to add %s %s: %v", key, val, err)
		}

		fpath, err := client.Storage().Write(ctx, fdata)
		if err != nil {
			return err
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

	releasePath, err := client.Storage().Write(ctx, releaseData)
	if err != nil {
		return err
	}

	release := &types.Release{
		Tag:        valist.Tag,
		ReleaseCID: releasePath,
		MetaCID:    types.DeprecationNotice,
	}

	fmt.Printf("%s@%s\n", releaseMeta.Name, release.Tag)

	for name, artifact := range releaseMeta.Artifacts {
		fmt.Printf("- %s: %s\n", name, artifact.Provider)
	}

	if dryrun {
		return nil
	}

	vote, err := client.VoteRelease(ctx, res.OrgID, res.RepoName, release)
	if err != nil {
		return err
	}

	if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
		fmt.Printf("Voted to publish release %s %d/%d\n", release.Tag, vote.SigCount, vote.Threshold)
	} else {
		fmt.Printf("Approved release %s\n", release.Tag)
	}

	return nil
}
