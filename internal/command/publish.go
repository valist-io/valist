package command

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/build"
	"github.com/valist-io/valist/internal/command/utils/lifecycle"
	"github.com/valist-io/valist/internal/command/utils/org"
	"github.com/valist-io/valist/internal/command/utils/repo"
	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/prompt"
)

func NewPublishCommand() *cli.Command {
	return &cli.Command{
		Name:  "publish",
		Usage: "Publish a package",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "dryrun",
				Usage: "Build and skip publish",
			},
		},
		Before: lifecycle.SetupClient,
		Action: func(c *cli.Context) error {
			client := c.Context.Value(core.ClientKey).(*client.Client)

			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			var valistFile build.Config
			if err := valistFile.Load(filepath.Join(cwd, "valist.yml")); err != nil {
				return err
			}

			releaseName := fmt.Sprintf("%s/%s/%s",
				strings.ToLower(valistFile.Org),
				strings.ToLower(valistFile.Repo),
				strings.ToLower(valistFile.Tag),
			)

			res, err := client.ResolvePath(c.Context, releaseName)

			switch err {
			case nil:
			case types.ErrOrganizationNotExist:
				answer, err := prompt.Confirm("This organization does not exist, would you like to create it?").Run()
				if err != nil {
					return err
				}

				if strings.ToLower(answer)[0:1] == "y" {
					orgID, err := org.CreateOrg(client, c.Context, valistFile.Org)
					if err != nil {
						return err
					}
					res.OrgID = orgID
				} else {
					return nil
				}

				fmt.Println("Creating repository...")

				err = repo.CreateRepo(client, c.Context, res.OrgID, valistFile.Repo)
				if err != nil {
					return err
				}

			case types.ErrRepositoryNotExist:
				answer, err := prompt.Confirm("This repository does not exist, would you like to create it?").Run()
				if err != nil {
					return err
				}

				if strings.ToLower(answer)[0:1] == "y" {
					err = repo.CreateRepo(client, c.Context, res.OrgID, valistFile.Repo)
					if err != nil {
						return err
					}
				}

			case types.ErrReleaseNotExist:
			default:
				return err
			}

			if res.Release != nil {
				return errors.Errorf("Release %s already exists", res.ReleaseTag)
			}

			_, err = build.Run(cwd, valistFile)
			if err != nil {
				return err
			}

			var readme string
			readmeBytes, err := os.ReadFile("README.md")

			if err != nil {
				fmt.Println("Readme not found")
			} else {
				readme = string(readmeBytes)
			}

			releaseMeta := &types.ReleaseMeta{
				Name:      releaseName,
				Readme:    readme,
				Artifacts: make(map[string]types.Artifact),
			}

			for platform, artifact := range valistFile.Platforms {
				fileData, err := os.ReadFile(filepath.Join(cwd, valistFile.Out, artifact))
				if err != nil {
					return err
				}

				path, err := client.Storage().Write(c.Context, fileData)
				if err != nil {
					return err
				}

				releaseMeta.Artifacts[platform] = types.Artifact{
					SHA256:    fmt.Sprintf("%x", sha256.Sum256(fileData)),
					Providers: []string{path},
				}
			}

			releaseData, err := json.Marshal(releaseMeta)
			if err != nil {
				return err
			}

			releaseCID, err := client.Storage().Write(c.Context, releaseData)
			if err != nil {
				return err
			}

			release := &types.Release{
				Tag:        valistFile.Tag,
				ReleaseCID: releaseCID,
				MetaCID:    "QmRBwMae3Skqzc1GmAKBdcnFFPnHeD585MwYtVZzfh9Tkh", // Deprecation notice
			}

			fmt.Println("Tag:", release.Tag)
			fmt.Println("ReleaseCID:", releaseCID)

			if c.Bool("dryrun") {
				return nil
			}

			vote, err := client.VoteRelease(c.Context, res.OrgID, valistFile.Repo, release)
			if err != nil {
				return err
			}

			if big.NewInt(1).Cmp(vote.Threshold) == -1 && vote.SigCount.Cmp(vote.Threshold) == -1 {
				fmt.Printf("Voted to publish release %s %d/%d\n", release.Tag, vote.SigCount, vote.Threshold)
			} else {
				fmt.Printf("Approved release %s\n", release.Tag)
			}

			return nil
		},
	}
}
