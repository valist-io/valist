package command

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/valist-io/registry/internal/build"
)

func NewBuildCommand() *cli.Command {
	return &cli.Command{
		Name:  "build",
		Usage: "Build the target valist project",
		Action: func(c *cli.Context) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}

			dockerFilePath := filepath.Join(wd, "Dockerfile")
			valistFilePath := filepath.Join(wd, "valist.yml")

			// Load valist.yml
			var valistFile build.Config
			err = valistFile.Load(valistFilePath)
			if err != nil {
				return err
			}

			// Set build command from valist.yml
			buildCommand := valistFile.Build
			// Set outPath to parent folder using filePath.Dir()
			outPath := filepath.Dir(valistFile.Out)

			// If projectType is npm, run npm pack and set out to .tgz
			if valistFile.Type == "npm" {
				packageJsonPath := filepath.Join(wd, "package.json")
				packageJson, err := build.ParsePackageJson(packageJsonPath)
				if err != nil {
					return err
				}

				buildCommand = valistFile.Build + "&& npm run pack"
				outPath = fmt.Sprintf("%s-%s.tgz", packageJson.Name, packageJson.Version)
			}

			// If image is not set in valist.yml use default image
			if valistFile.Image == "" {
				valistFile.Image = build.DefaultImages[valistFile.Type]
			}

			// Create dockerConfig used to generate Dockerfile
			dockerConfig := build.DockerConfig{
				Path:           dockerFilePath,
				BaseImage:      valistFile.Image,
				Source:         "./",
				BuildCommand:   buildCommand,
				InstallCommand: valistFile.Install,
			}

			if err := build.GenerateDockerfile(dockerConfig); err != nil {
				return err
			}

			// Create the build image using the dockerfile
			if err := build.Create("valist-build", dockerFilePath); err != nil {
				return err
			}

			// Export the build from the build image
			return build.Export("valist-build", outPath)
		},
	}
}
