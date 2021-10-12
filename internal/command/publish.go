package command

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

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

			orgID, err := client.GetOrganizationID(c.Context, valistFile.Org)
			if err == types.ErrOrganizationNotExist {
				answer, err := prompt.Confirm("This organization does not exist, would you like to create it?").Run()
				if err != nil {
					return err
				}

				if strings.ToLower(answer) == "y" {
					orgID, err = org.CreateOrg(client, c.Context, valistFile.Org)
					if err != nil {
						return nil
					}
				} else {
					return nil
				}

				err = repo.CreateRepo(client, c.Context, orgID, valistFile.Repo)
				if err != nil {
					return err
				}
			} else if err != nil && err != types.ErrOrganizationNotExist {
				return err
			}

			_, err = build.Run(cwd, valistFile)
			if err != nil {
				return err
			}

			releaseCID, err := client.Storage().WriteFile(c.Context, valistFile.Out)
			if err != nil {
				return err
			}

			metaCID, err := client.Storage().WriteFile(c.Context, valistFile.Meta)
			if err != nil && !os.IsNotExist(err) {
				return err
			}

			release := &types.Release{
				Tag:        valistFile.Tag,
				ReleaseCID: releaseCID,
				MetaCID:    metaCID,
			}

			fmt.Println("Tag:", release.Tag)
			fmt.Println("ReleaseCID:", releaseCID)
			fmt.Println("MetaCID:", metaCID)

			if c.Bool("dryrun") {
				return nil
			}

			vote, err := client.VoteRelease(c.Context, orgID, valistFile.Repo, release)
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
