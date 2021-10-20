package build

import (
	"fmt"
	"strings"

	"github.com/valist-io/valist/internal/prompt"
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

	org, err := prompt.OrganizationName("").Run()
	if err != nil {
		return err
	}

	repo, err := prompt.RepositoryName("").Run()
	if err != nil {
		return err
	}

	projectName := fmt.Sprintf("%s/%s",
		strings.ToLower(org),
		strings.ToLower(repo),
	)

	cfg.Name = projectName

	cfg.Tag, err = prompt.ReleaseTag("0.0.1").Run()
	if err != nil {
		return err
	}

	cfg.Artifacts = make(map[string]string)
	// If there are artifacts set isArtifacts to y
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
		cfg.Artifacts[platform] = src
	}

	return cfg.Save("valist.yml")
}
