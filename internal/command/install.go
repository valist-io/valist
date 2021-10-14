package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/valist-io/valist/internal/command/utils/lifecycle"
	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/config"
)

func NewInstallCommand() *cli.Command {
	return &cli.Command{
		Name:   "install",
		Usage:  "Installs a package or artifact",
		Before: lifecycle.SetupClient,
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			client := c.Context.Value(core.ClientKey).(*client.Client)
			config := c.Context.Value(core.ConfigKey).(*config.Config)

			res, err := client.ResolvePath(c.Context, c.Args().Get(0))
			if err != nil {
				return err
			}

			if res.Release == nil {
				return fmt.Errorf("invalid release path: %s", c.Args().Get(0))
			}

			meta, err := client.GetRepositoryMeta(c.Context, res.Repository.MetaCID)
			if err != nil {
				return err
			}

			if meta.ProjectType == "npm" {
				return errors.New("For NPM packages please run valist daemon and install using the NPM registry.")
			}

			releaseMeta, err := client.GetReleaseMeta(c.Context, res.Release.ReleaseCID)
			if err != nil {
				return err
			}

			targetPlatform := runtime.GOOS + "/" + runtime.GOARCH
			artifact, ok := releaseMeta.Artifacts[targetPlatform]
			if !ok {
				return errors.New("Target platform not found in release")
			}

			fmt.Println("Installing for target platform: ", targetPlatform)

			targetData, err := client.Storage().ReadFile(c.Context, artifact.Provider)
			if err != nil {
				return err
			}

			binPath := config.InstallPath()
			exePath := filepath.Join(binPath, res.RepoName)

			if err := os.MkdirAll(binPath, 0744); err != nil {
				return err
			}

			if err := os.WriteFile(exePath, targetData, 0744); err != nil {
				return err
			}

			fmt.Printf("Successfully installed: %s\n", exePath)

			if !strings.Contains(os.Getenv("PATH"), binPath) {
				fmt.Println()
				fmt.Println("Valist bin directory not detected in $PATH, please run the following:")
				fmt.Printf(`    export PATH="$PATH:%s"`, binPath)
			}

			return nil
		},
	}
}
