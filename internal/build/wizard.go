package build

import (
	"fmt"

	"github.com/valist-io/registry/internal/prompt"
)

// ConfigWizard runs an interactive configurator.
func ConfigWizard() error {
	var cfg Config
	var err error

	// If project type is not set ask for projectType
	_, cfg.Type, err = prompt.RepositoryProjectType().Run()
	if err != nil {
		return err
	}

	cfg.Org, err = prompt.OrganizationName("").Run()
	if err != nil {
		return err
	}

	cfg.Repo, err = prompt.RepositoryName("").Run()
	if err != nil {
		return err
	}

	cfg.Tag, err = prompt.ReleaseTag("0.0.1").Run()
	if err != nil {
		return err
	}

	cfg.Meta, err = prompt.ReleaseMetaPath().Run()
	if err != nil {
		return err
	}

	defaultInstall := DefaultInstalls[cfg.Type]
	cfg.Install, err = prompt.InstallCommand(defaultInstall).Run()
	if err != nil {
		return err
	}

	defaultBuild := DefaultBuilds[cfg.Type]
	cfg.Build, err = prompt.BuildCommand(defaultBuild).Run()
	if err != nil {
		return err
	}

	defaultImage := DefaultImages[cfg.Type]
	cfg.Image, err = prompt.DockerImage(defaultImage).Run()
	if err != nil {
		return err
	}

	// If the project type is not node prompt for out path
	if cfg.Type != "node" {
		cfg.Out, err = prompt.BuildOutPath().Run()
		if err != nil {
			return err
		}
	}

	cfg.Platforms = make(map[string]string)
	// If there is artifacts set isArtifacts to y
	isArtifacts, err := prompt.BuildPlatforms().Run()
	if err != nil {
		return err
	}

	// If there are artifacts then prompt for their os, arch, & name/path
	for isArtifacts == "y" {
		osName, err := prompt.BuildOS().Run()
		if err != nil {
			return err
		}

		if osName == "" {
			break
		}

		arch, err := prompt.BuildArch().Run()
		if err != nil {
			return err
		}

		src, err := prompt.BuildArtifactPath().Run()
		if err != nil {
			return err
		}

		// Set artifact key to os/arch and value to src
		platform := fmt.Sprintf("%s/%s", osName, arch)
		cfg.Platforms[platform] = src
	}

	return cfg.Save("valist.yml")
}
