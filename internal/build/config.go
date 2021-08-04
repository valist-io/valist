package build

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v3"
)

// Define type for Config.
type Config struct {
	Type      string            `yaml:"type"`
	Org       string            `yaml:"org"`
	Repo      string            `yaml:"repo"`
	Tag       string            `yaml:"tag"`
	Meta      string            `yaml:"meta,omitempty"`
	Image     string            `yaml:"image,omitempty"`
	Build     string            `yaml:"build,omitempty"`
	Install   string            `yaml:"install,omitempty"`
	Out       string            `yaml:"out,omitempty"`
	Platforms map[string]string `yaml:"artifacts,omitempty"`
}

var DefaultImages = map[string]string{
	"binary": "gcc:bullseye",
	"node":   "node:buster",
	"npm":    "node:buster",
	"go":     "golang:buster",
	"rust":   "rust:buster",
	"python": "python:buster",
	"docker": "scratch",
	"c++":    "gcc:bullseye",
	"static": "",
}

var DefaultInstalls = map[string]string{
	"binary": "make install",
	"node":   "npm install",
	"npm":    "npm install",
	"go":     "go get ./...",
	"rust":   "cargo install",
	"python": "pip install -r requirements.txt",
	"docker": "",
	"c++":    "make install",
	"static": "",
}

var DefaultBuilds = map[string]string{
	"binary": "make build",
	"node":   "npm run build",
	"npm":    "npm run build",
	"go":     "go build",
	"rust":   "cargo build",
	"python": "python3 -m build",
	"docker": "",
	"c++":    "make build",
	"static": "",
}

func (c Config) Save(path string) error {
	yamlData, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(path, yamlData, 0644)
}

func (c *Config) Load(path string) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlFile, c)
}

func validateLength(value string) error {
	if len(value) > 0 {
		return nil
	}
	return errors.New("Length must be greater than 0")
}

func validateYesNo(value string) error {
	switch strings.ToLower(string(value[0])) {
	case
		"y",
		"n":
		return nil
	}
	return errors.New("Must be y or n")
}

func ValistFileFromWizard() error {
	var out string

	// If project type is not set ask for projectType
	projectPrompt := promptui.Select{
		Label: "Repository type",
		Items: []string{
			"binary", "go", "node", "python", "docker", "static",
		},
	}
	_, projectType, err := projectPrompt.Run()
	if err != nil {
		return err
	}

	orgPrompt := promptui.Prompt{
		Label:    "Valist Organization name or username",
		Validate: validateLength,
	}
	org, err := orgPrompt.Run()
	if err != nil {
		return err
	}

	repoPrompt := promptui.Prompt{
		Label:    "Valist Repository name",
		Validate: validateLength,
	}
	repo, err := repoPrompt.Run()
	if err != nil {
		return err
	}

	tagPrompt := promptui.Prompt{
		Label:   "The latest release tag",
		Default: "0.0.1",
	}
	tag, err := tagPrompt.Run()
	if err != nil {
		return err
	}

	metaPrompt := promptui.Prompt{
		Label: "Path to meta file(README.md)",
	}
	meta, err := metaPrompt.Run()
	if err != nil {
		return err
	}

	defaultInstall := DefaultInstalls[projectType]
	installPrompt := promptui.Prompt{
		Label:   "Command used to install dependencies",
		Default: defaultInstall,
	}
	install, err := installPrompt.Run()
	if err != nil {
		return err
	}

	defaultBuild := DefaultBuilds[projectType]
	buildPrompt := promptui.Prompt{
		Label:   "Command used to build your project",
		Default: defaultBuild,
	}
	buildCommand, err := buildPrompt.Run()
	if err != nil {
		return err
	}

	defaultImage := DefaultImages[projectType]
	imagePrompt := promptui.Prompt{
		Label: fmt.Sprintf("Docker image (if not set, will default to %v)", defaultImage),
	}
	image, err := imagePrompt.Run()
	if err != nil {
		return err
	}

	// If the project type is not node prompt for out path
	if projectType != "node" {
		outPrompt := promptui.Prompt{
			Label: "Build output file/directory",
		}
		out, err = outPrompt.Run()
		if err != nil {
			return err
		}
	}

	// Create valist config with empty artifacts mapping
	var cfg = Config{
		Type:      projectType,
		Org:       org,
		Repo:      repo,
		Tag:       tag,
		Meta:      meta,
		Image:     image,
		Build:     buildCommand,
		Install:   install,
		Out:       out,
		Platforms: map[string]string{},
	}

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
		cfg.Platforms[fmt.Sprintf("%s/%s", osName, arch)] = src
	}

	return cfg.Save("valist.yml")
}

func ValistFileFromTemplate(projectType string, path string) error {
	type TemplateCfg struct {
		RenderMeta      bool
		RenderInstall   bool
		RenderPlatforms bool
		Config
	}

	cfg := TemplateCfg{}
	cfg.Type = projectType
	cfg.RenderMeta = true
	cfg.RenderPlatforms = true

	if projectType != "npm" {
		cfg.Out = "path_to_artifact_or_build_directory"
	}

	if projectType == "npm" {
		cfg.RenderMeta = false
		cfg.RenderPlatforms = false
	}

	if projectType == "static" || projectType == "go" {
		cfg.RenderInstall = false
	}

	cfg.Install = DefaultInstalls[projectType]
	cfg.Image = DefaultImages[projectType]
	cfg.Build = DefaultBuilds[projectType]

	configTemplate, err := template.New("valistConfig").Parse(`# The project type
type: {{.Type}}

# The valist organization
org: 

# The valist repository
repo: 

# The latest release tag
tag: 

# The command used for building the project
build: {{.Build}}

# The command used for installing the project's dependencies
# install: {{.Install}}

# The docker image used for building the project. Will default to {{.Image}} for {{.Type}}.
# image: {{.Image}}
{{if .Out}}
# The project's build/output folder
out: {{.Out}}{{end}}
{{if .RenderMeta}}
# The metadata file for the latest release, typically README.md or RELEASE.md
# meta: README.md{{end}}
{{if .RenderPlatforms}}
# The project's supported os/arch platforms and their corresponding artifacts
# platforms:
  # linux/amd64: path_to_artifact
  # linux/arm64: path_to_artifact
  # darwin/amd64: path_to_artifact
  # windows/amd64: path_to_artifact
{{end}}`)

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write template to valist.yml
	return configTemplate.Execute(f, cfg)
}

// https://pkg.go.dev/github.com/go-playground/validator/v10
// func (c Config) Validate() error {
// 	if c.Type != "go" {
// 		return err
// 	}
// }
