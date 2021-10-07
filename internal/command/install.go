package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/matishsiao/goInfo"
	"github.com/urfave/cli/v2"
	"github.com/valist-io/valist/internal/command/utils/lifecycle"
	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/types"
)

const (
	rootDir = ".valist"
	binDir  = "bin"
)

func getPathInstructions() string {
	shell := os.Getenv("SHELL")
	rcPath := "~/.zshrc"

	if strings.Contains(shell, "bash") {
		rcPath = "~/.bash_profile"
	}

	command := fmt.Sprintf(`
export PATH="$PATH:$HOME/.valist/bin"
echo export PATH=\$PATH:\$HOME/.valist/bin >> %v
`, rcPath)

	return command
}

func printPathInstructions() {
	pathEnv := os.Getenv("PATH")
	if !strings.Contains(pathEnv, path.Join(rootDir, binDir)) {
		fmt.Println()
		fmt.Println("~/.valist/bin not detected in $PATH, please run the following:")
		fmt.Println(getPathInstructions())
	}
}

func NewInstallCommand() *cli.Command {
	return &cli.Command{
		Name:   "install",
		Usage:  "Installs a package or artifact",
		Action: action,
		Before: lifecycle.SetupClient,
	}
}

func action(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowSubcommandHelpAndExit(c, 1)
	}

	client := c.Context.Value(core.ClientKey).(*client.Client)

	hostInfo, err := goInfo.GetInfo()
	if err != nil {
		return nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	res, err := client.ResolvePath(c.Context, c.Args().Get(0))
	if err != nil {
		return err
	}

	if res.ReleaseTag == "" {
		res.Release, err = client.GetLatestRelease(c.Context, res.OrgID, res.RepoName)
		if err != nil {
			return err
		}
		res.ReleaseTag = res.Release.Tag
	}

	meta, err := client.GetRepositoryMeta(c.Context, res.Repository.MetaCID)
	if err != nil {
		return err
	}

	if meta.ProjectType == "npm" {
		return errors.New("For NPM packages please run valist daemon and install using the NPM registry.")
	}

	valistBinDir := path.Join(home, rootDir, binDir)
	installPath := path.Join(valistBinDir, res.RepoName)

	metaPath := res.Release.ReleaseCID + "/meta.json"

	releaseData, err := client.Storage().ReadFile(c.Context, metaPath)
	if err != nil {
		return err
	}

	targetOs := strings.ToLower(hostInfo.OS)
	targetArch := strings.ToLower(hostInfo.Platform)
	if targetArch == "x86_64" {
		targetArch = "amd64"
	}
	targetPlatform := targetOs + "/" + targetArch
	releaseMeta := &types.ReleaseMeta{}
	json.Unmarshal(releaseData, releaseMeta)

	var targetCID string
	if _, ok := releaseMeta.Platforms[targetPlatform]; ok {
		if releaseMeta.Platforms[targetPlatform].StorageProviders == nil {
			return errors.New("Missing storage providers")
		}
		targetCID = releaseMeta.Platforms[targetPlatform].StorageProviders[0]
	} else {
		return errors.New("Target platform not found in release")
	}

	fmt.Println("Installing for target platform: ", targetPlatform)

	targetData, err := client.Storage().ReadFile(c.Context, targetCID)
	if err != nil {
		return err
	}

	if _, err := os.Stat(valistBinDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(valistBinDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if err := os.WriteFile(installPath, targetData, 0744); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Successfully installed %v into ~/%v/%v/%v!", res.RepoName, rootDir, binDir, res.RepoName))

	printPathInstructions()

	return nil
}
