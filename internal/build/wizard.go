package build

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func validateLength(value string) error {
	if len(value) > 0 {
		return nil
	}
	return fmt.Errorf("Length must be greater than 0")
}

func validateYesNo(value string) error {
	switch value[0] {
	case 'Y', 'y', 'N', 'n':
		return nil
	default:
		return fmt.Errorf("Must be y or n")
	}
}

// ConfigWizard runs an interactive configurator.
func ConfigWizard() error {
	var cfg Config
	var err error

	// If project type is not set ask for projectType
	projectPrompt := promptui.Select{
		Label: "Repository type",
		Items: []string{
			"binary", "go", "node", "python", "docker", "static",
		},
	}
	_, cfg.Type, err = projectPrompt.Run()
	if err != nil {
		return err
	}

	orgPrompt := promptui.Prompt{
		Label:    "Valist Organization name or username",
		Validate: validateLength,
	}
	cfg.Org, err = orgPrompt.Run()
	if err != nil {
		return err
	}

	repoPrompt := promptui.Prompt{
		Label:    "Valist Repository name",
		Validate: validateLength,
	}
	cfg.Repo, err = repoPrompt.Run()
	if err != nil {
		return err
	}

	tagPrompt := promptui.Prompt{
		Label:   "The latest release tag",
		Default: "0.0.1",
	}
	cfg.Tag, err = tagPrompt.Run()
	if err != nil {
		return err
	}

	metaPrompt := promptui.Prompt{
		Label: "Path to meta file(README.md)",
	}
	cfg.Meta, err = metaPrompt.Run()
	if err != nil {
		return err
	}

	defaultInstall := DefaultInstalls[cfg.Type]
	installPrompt := promptui.Prompt{
		Label:   "Command used to install dependencies",
		Default: defaultInstall,
	}
	cfg.Install, err = installPrompt.Run()
	if err != nil {
		return err
	}

	defaultBuild := DefaultBuilds[cfg.Type]
	buildPrompt := promptui.Prompt{
		Label:   "Command used to build your project",
		Default: defaultBuild,
	}
	cfg.Build, err = buildPrompt.Run()
	if err != nil {
		return err
	}

	defaultImage := DefaultImages[cfg.Type]
	imagePrompt := promptui.Prompt{
		Label: fmt.Sprintf("Docker image (if not set, will default to %v)", defaultImage),
	}
	cfg.Image, err = imagePrompt.Run()
	if err != nil {
		return err
	}

	// If the project type is not node prompt for out path
	if cfg.Type != "node" {
		outPrompt := promptui.Prompt{
			Label: "Build output file/directory",
		}
		cfg.Out, err = outPrompt.Run()
		if err != nil {
			return err
		}
	}

	cfg.Platforms = make(map[string]string)
	platformsPrompt := promptui.Prompt{
		Label:     "Are you building for multiple architecures? (y,N)",
		IsConfirm: true,
		Validate:  validateYesNo,
	}
	// If there is artifacts set isArtifacts to y
	isArtifacts, err := platformsPrompt.Run()
	if err != nil {
		return err
	}

	// If there are artifacts then prompt for their os, arch, & name/path
	for isArtifacts == "y" {
		osPrompt := promptui.Prompt{
			Label: "Platform operating system, e.g. linux, darwin, freebsd, windows (leave blank to quit)",
		}

		osName, err := osPrompt.Run()
		if err != nil {
			return err
		}

		if osName == "" {
			break
		}

		archPrompt := promptui.Prompt{
			Label: "Platform architecture, e.g. amd64, arm64",
		}

		arch, err := archPrompt.Run()
		if err != nil {
			return err
		}

		srcPrompt := promptui.Prompt{
			Label: "Artifact file path",
		}

		src, err := srcPrompt.Run()
		if err != nil {
			return err
		}

		// Set artifact key to os/arch and value to src
		platform := fmt.Sprintf("%s/%s", osName, arch)
		cfg.Platforms[platform] = src
	}

	return cfg.Save("valist.yml")
}
