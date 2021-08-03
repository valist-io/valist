package build

import (
	"errors"
	"fmt"
	"os"

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
	Artifacts map[string]string `yaml:"artifacts,omitempty"`
}

var DefaultImages = map[string]string{
	"binary": "gcc:bullseye",
	"node":   "node:buster",
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
	"go":     "go get",
	"rust":   "cargo install",
	"python": "pip install -r requirements.txt",
	"docker": "",
	"c++":    "make install",
	"static": "",
}

var DefaultBuilds = map[string]string{
	"binary": "make build",
	"node":   "npm run build",
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
	// Read yaml file from disk
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
	switch value {
	case
		"y",
		"Y",
		"n",
		"N":
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
		Label:    "Organization name or username",
		Validate: validateLength,
	}
	org, err := orgPrompt.Run()
	if err != nil {
		return err
	}

	repoPrompt := promptui.Prompt{
		Label:    "Repository name",
		Validate: validateLength,
	}
	repo, err := repoPrompt.Run()
	if err != nil {
		return err
	}

	tagPrompt := promptui.Prompt{
		Label:   "Your current release tag/version",
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
		Label: fmt.Sprintf("Docker image (if not set, will default to %s)", defaultImage),
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
		Artifacts: map[string]string{},
	}

	artifactsPrompt := promptui.Prompt{
		Label:     "Are you building for multiple architecures? (y,N)",
		IsConfirm: true,
		Validate:  validateYesNo,
	}
	// If there is artifacts set isArtifacts to y
	isArtifacts, err := artifactsPrompt.Run()
	if err != nil {
	}

	// If there are artifacts then prompt for their os, arch, & name/path
	for isArtifacts == "y" || isArtifacts == "Y" {
		osPrompt := promptui.Prompt{
			Label: "Artifact operating system (leave blank to quit)",
		}

		osName, err := osPrompt.Run()
		if err != nil {
			return err
		}

		if osName == "" {
			break
		}

		archPrompt := promptui.Prompt{
			Label: "Artifact architecture ",
		}

		arch, err := archPrompt.Run()
		if err != nil {
			return err
		}

		srcPrompt := promptui.Prompt{
			Label: "Artifact file-path/name",
		}

		src, err := srcPrompt.Run()
		if err != nil {
			return err
		}

		// Set artifact key to os/arch and value to src
		cfg.Artifacts[fmt.Sprintf("%s/%s", osName, arch)] = src
	}

	return cfg.Save("valist.yml")
}

func ValistFileFromTemplate(projectType string) error {
	var outValue string
	if projectType != "npm" {
		outValue = "_"
	}

	yamlData, err := yaml.Marshal(Config{
		Type: projectType,
		Org:  "_",
		Repo: "_",
		Tag:  "_",
		Out:  outValue,
	})
	if err != nil {
		return err
	}
	yamlString := string(yamlData)

	if projectType != "npm" {
		yamlString += "\n# The metadata file for the current release"
		yamlString += "\n# meta: "

		yamlString += "\n\n# The docker image used for building the project"
		yamlString += "\n# image: "
	}

	yamlString += "\n\n# The command used for installing the project's dependencies"
	yamlString += "\n# install: "

	yamlString += "\n\n# The command used for building the project"
	yamlString += "\n# build: "

	yamlData = []byte(yamlString)
	return os.WriteFile("valist.yml", yamlData, 0644)
}

// https://pkg.go.dev/github.com/go-playground/validator/v10
// func (c Config) Validate() error {
// 	if c.Type != "go" {
// 		return err
// 	}
// }
