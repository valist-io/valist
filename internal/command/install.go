package command

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/valist-io/valist/internal/command/utils/lifecycle"
	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
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
		npmrcContent := "@" + res.OrgName + ":registry=http://localhost:9000/api/npm" + "\n"

		file, err := ioutil.ReadFile(".npmrc")
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
		fileText := string(file)
		stringSet := strings.Contains(fileText, npmrcContent)

		if !stringSet {
			fileText = fileText + "\n" + npmrcContent
			if err := os.WriteFile(".npmrc", []byte(fileText), 0644); err != nil {
				return err
			}
		}

		npmPackageName := "@" + res.OrgName + "/" + res.RepoName + "@" + res.ReleaseTag
		cmd := exec.Command("npm", "i", npmPackageName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
		return nil
	}

	valistBinDir := path.Join(home, rootDir, binDir)
	installPath := path.Join(valistBinDir, res.RepoName)

	data, err := client.Storage().ReadFile(c.Context, res.Release.ReleaseCID)
	if err != nil {
		return err
	}

	if _, err := os.Stat(valistBinDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(valistBinDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if err := os.WriteFile(installPath, data, 0744); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Successfully installed %v into ~/%v/%v/%v!", res.RepoName, rootDir, binDir, res.RepoName))

	printPathInstructions()

	return nil
}
