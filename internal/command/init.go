package command

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"github.com/valist-io/registry/internal/build"
)

func NewInitCommand() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Generate a new valist project",
		Action: func(c *cli.Context) error {
			var out string

			orgPrompt := promptui.Prompt{
				Label: "Organization name or Username",
			}
			org, err := orgPrompt.Run()
			if err != nil {
				return err
			}

			repoPrompt := promptui.Prompt{
				Label: "Repository name",
			}
			repo, err := repoPrompt.Run()
			if err != nil {
				return err
			}

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

			defaultInstall := build.DefaultImages[projectType]
			installPrompt := promptui.Prompt{
				Label:   "Install command",
				Default: defaultInstall,
			}
			install, err := installPrompt.Run()
			if err != nil {
				return err
			}

			defaultBuild := build.DefaultBuilds[projectType]
			buildPrompt := promptui.Prompt{
				Label:   "Build command",
				Default: defaultBuild,
			}
			buildCommand, err := buildPrompt.Run()
			if err != nil {
				return err
			}

			tagPrompt := promptui.Prompt{
				Label:   "Release tag",
				Default: "0.0.1",
			}
			tag, err := tagPrompt.Run()
			if err != nil {
				return err
			}

			metaPrompt := promptui.Prompt{
				Label: "Release meta file",
			}
			meta, err := metaPrompt.Run()
			if err != nil {
				return err
			}

			defaultImage := build.DefaultImages[projectType]
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
					Label: "output file/directory",
				}
				out, err = outPrompt.Run()
				if err != nil {
					return err
				}
			}

			// Create valist config with empty artifacts mapping
			var cfg = build.Config{
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
				Label:     "Does your build have artifacts?",
				IsConfirm: true,
			}
			// If there is artifacts set isArtifacts to y
			isArtifacts, err := artifactsPrompt.Run()
			if err != nil {
			}

			// If there are artifacts prompt for their os, arch, & name/path
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

				namePrompt := promptui.Prompt{
					Label: "Artifact path/name",
				}

				name, err := namePrompt.Run()
				if err != nil {
					return err
				}

				// Set artifact key to os/arch and value to name
				cfg.Artifacts[fmt.Sprintf("%s/%s", osName, arch)] = name
			}

			return cfg.Save("valist.yml")
		},
	}
}
