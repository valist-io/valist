package build

import (
	"errors"
	"fmt"
	"path/filepath"
)

func Run(projectPath, configYml string) ([]string, error) {
	var artifactPaths []string
	var packageName string

	dockerFilePath := filepath.Join(projectPath, "Dockerfile")
	valistFilePath := filepath.Join(projectPath, configYml)

	// Load valist.yml
	var valistFile Config
	err := valistFile.Load(valistFilePath)
	if err != nil {
		return nil, err
	}

	buildCommand := valistFile.Build
	hostPath := filepath.Dir(filepath.Join(projectPath, valistFile.Out))
	containerPath := fmt.Sprintf("valist-build:/opt/build/%s", filepath.Dir(valistFile.Out))

	// Prevent artifacts from being copied into nested folder if folder already exists
	if valistFile.Out != filepath.Dir(valistFile.Out) {
		containerPath = fmt.Sprintf("valist-build:/opt/build/%s/.", filepath.Dir(valistFile.Out))
	}

	fmt.Println("container", containerPath)
	fmt.Println("host", hostPath)

	// If projectType is npm, run npm pack and set out to .tgz
	if valistFile.Type == "npm" {
		packageJsonPath := filepath.Join(projectPath, "package.json")
		packageJson, err := ParsePackageJson(packageJsonPath)
		if err != nil {
			return nil, err
		}

		buildCommand = valistFile.Build + " && npm pack"
		hostPath = projectPath
		packageName = fmt.Sprintf("%s-%s.tgz", packageJson.Name, packageJson.Version)
		containerPath = fmt.Sprintf("valist-build:/opt/build/%s", packageName)
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

	// Create the build image using the dockerfile
	if err := Create("valist-build", projectPath); err != nil {
		return nil, err
	}

	// Export the build from the build image
	if err := Export("valist-build", hostPath, containerPath); err != nil {
		return nil, err
	}

	// If project type is npm return outPath
	if valistFile.Type == "npm" {
		return append(artifactPaths, filepath.Join(projectPath, packageName)), nil
	}

	// If platforms are defined in config then use artifact paths
	if len(valistFile.Platforms) > 0 {
		for _, artifact := range valistFile.Platforms {
			artifactPaths = append(
				artifactPaths,
				filepath.Join(projectPath, valistFile.Out, artifact),
			)
		}
		return artifactPaths, nil
	}

	// If platforms are not defined but out is defined in config then use out path
	if valistFile.Out != "" {
		return append(artifactPaths, filepath.Join(projectPath, valistFile.Out)), nil
	}

	return nil, errors.New("Must define either out or platforms in config for this package type")
}
