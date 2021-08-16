package build

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/valist-io/registry/internal/npm"
)

func Run(projectPath, configYml string) ([]string, error) {
	var artifactPaths []string
	var packageName string

	dockerFilePath := filepath.Join(projectPath, "Dockerfile.repro")
	valistFilePath := filepath.Join(projectPath, configYml)

	// Load valist.yml
	var valistFile Config
	if err := valistFile.Load(valistFilePath); err != nil {
		return nil, err
	}

	buildCommand := valistFile.Build
	outPath := valistFile.Out

	// Prevent artifacts from being copied into nested folder if folder already exists
	// if valistFile.Out != filepath.Dir(valistFile.Out) {
	// 	containerPath = fmt.Sprintf("%s/.", filepath.Dir(valistFile.Out))
	// }

	// If projectType is npm, run npm pack and set out to .tgz
	if valistFile.Type == "npm" {
		packageJsonPath := filepath.Join(projectPath, "package.json")
		packageJson, err := npm.ParsePackageJSON(packageJsonPath)
		if err != nil {
			return nil, err
		}

		buildCommand = valistFile.Build + " && npm pack"
		packageName = fmt.Sprintf("%s-%s.tgz", packageJson.Name, packageJson.Version)
		outPath = packageName
	}

	// If image is not set in valist.yml use default image
	if valistFile.Image == "" {
		valistFile.Image = DefaultImages[valistFile.Type]
	}

	if valistFile.Type == "npm" && valistFile.Install == "" {
		valistFile.Install = DefaultInstalls[valistFile.Type]
	}

	// Create dockerConfig used to generate Dockerfile
	dockerConfig := DockerConfig{
		Path:           dockerFilePath,
		BaseImage:      valistFile.Image,
		Source:         "./",
		BuildCommand:   buildCommand,
		InstallCommand: valistFile.Install,
	}

	if err := GenerateDockerfile(dockerConfig); err != nil {
		return nil, err
	}

	// @ TODO Construct image name from (orgName, repoName, tag)
	// Create the build image using the dockerfile
	if err := Create("valist-build", dockerFilePath); err != nil {
		return nil, err
	}

	// Export the build from the build image
	if err := Export("valist-build", projectPath, outPath); err != nil {
		return nil, err
	}

	// If project type is npm return projectPath + packageName
	if valistFile.Type == "npm" {
		return append(artifactPaths, filepath.Join(projectPath, packageName)), nil
	}

	// If platforms are defined in config then use out + artifactPath
	for _, artifact := range valistFile.Platforms {
		artifactPaths = append(artifactPaths, filepath.Join(projectPath, valistFile.Out, artifact))
	}

	// If platforms are not defined but out is defined, use valistFile.Out
	if len(valistFile.Platforms) == 0 && valistFile.Out != "" {
		artifactPaths = append(artifactPaths, filepath.Join(projectPath, valistFile.Out))
	}

	if len(artifactPaths) == 0 {
		return nil, errors.New("Must define either out or platforms in config for this package type")
	}

	return artifactPaths, nil
}
