package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/valist-io/registry/internal/core/types"
)

func AccountPassphrase() *promptui.Prompt {
	return &promptui.Prompt{
		Label:       "Account passphrase",
		Mask:        '*',
		HideEntered: true,
		Validate:    ValidateMinLength(5),
	}
}

func OrganizationName(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:    "Organization name or username",
		Default:  value,
		Validate: ValidateMinLength(1),
	}
}

func OrganizationDescription(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:    "Organization description",
		Default:  value,
		Validate: ValidateMinLength(1),
	}
}

func OrganizationHomepage(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:   "Organization homepage",
		Default: value,
	}
}

func RepositoryName(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:    "Repository name",
		Default:  value,
		Validate: ValidateMinLength(1),
	}
}

func RepositoryDescription(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:    "Repository description",
		Default:  value,
		Validate: ValidateMinLength(1),
	}
}

func RepositoryProjectType() *promptui.Select {
	return &promptui.Select{
		Label: "Repository project type",
		Items: types.ProjectTypes,
	}
}

func RepositoryHomepage(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:   "Repository homepage",
		Default: value,
	}
}

func RepositoryURL(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:   "Repository url",
		Default: value,
	}
}

func ReleaseTag(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:   "Latest release tag",
		Default: value,
	}
}

func ReleaseMetaPath() *promptui.Prompt {
	return &promptui.Prompt{
		Label: "Path to metadata file (README.md)",
	}
}

func InstallCommand(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:   "Command used to install dependencies",
		Default: value,
	}
}

func BuildCommand(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label:   "Command used to build your project",
		Default: value,
	}
}

func DockerImage(value string) *promptui.Prompt {
	return &promptui.Prompt{
		Label: fmt.Sprintf("Docker image (if not set, will default to %v)", value),
	}
}

func BuildOutPath() *promptui.Prompt {
	return &promptui.Prompt{
		Label: "Build output file/directory",
	}
}

func BuildPlatforms() *promptui.Prompt {
	return &promptui.Prompt{
		Label:     "Are you building for multiple architecures? (y,N)",
		IsConfirm: true,
		Validate:  ValidateYesNo(),
	}
}

func BuildOS() *promptui.Prompt {
	return &promptui.Prompt{
		Label: "Platform operating system, e.g. linux, darwin, freebsd, windows (leave blank to quit)",
	}
}

func BuildArch() *promptui.Prompt {
	return &promptui.Prompt{
		Label: "Platform architecture, e.g. amd64, arm64",
	}
}

func BuildArtifactPath() *promptui.Prompt {
	return &promptui.Prompt{
		Label: "Artifact file path",
	}
}
