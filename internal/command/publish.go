package command

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cid "github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/build"
	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/core"
	"github.com/valist-io/registry/internal/core/client"
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
		Action: func(c *cli.Context) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg, err := config.Load(home)
			if err != nil {
				return err
			}

			var account accounts.Account
			if c.IsSet("account") {
				account.Address = common.HexToAddress(c.String("account"))
			} else {
				account.Address = cfg.Accounts.Default
			}

			client, err := client.NewClient(c.Context, cfg, account)
			if err != nil {
				return err
			}
			defer client.Close()

			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			var valistFile build.Config
			if err := valistFile.Load(filepath.Join(cwd, "valist.yml")); err != nil {
				return err
			}

			orgID, err := client.GetOrganizationID(c.Context, valistFile.Org)
			if err != nil {
				return err
			}

			artifactPaths, err := build.Run(cwd, valistFile)
			if err != nil {
				return err
			}

			var releaseCID cid.Cid
			if len(artifactPaths) == 0 {
				releaseCID, err = client.WriteFilePath(c.Context, artifactPaths[0])
			} else {
				releaseCID, err = client.WriteDirEntries(c.Context, cwd, artifactPaths)
			}

			if err != nil {
				return err
			}

			metaCID, err := client.WriteFilePath(c.Context, valistFile.Meta)
			if err != nil && !os.IsNotExist(err) {
				return err
			}

			release := &core.Release{
				Tag:        valistFile.Tag,
				ReleaseCID: releaseCID,
				MetaCID:    metaCID,
			}

			fmt.Println("Tag:", release.Tag)
			fmt.Println("ReleaseCID:", releaseCID.String())

			if metaCID.Defined() {
				fmt.Println("MetaCID:", metaCID.String())
			}

			if c.Bool("dryrun") {
				return nil
			}

			vote, err := client.VoteRelease(c.Context, &bind.TransactOpts{}, orgID, valistFile.Repo, release)
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
