package build

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/valist-io/valist/internal/registry/npm"
)

func Run(projectPath string, valistFile Config) ([]string, error) {
	var artifactPaths []string
	if valistFile.Org == "" || valistFile.Repo == "" || valistFile.Tag == "" {
		return nil, errors.New("Org, Repo & Tag required in valist config")
	}

	buildImageName := fmt.Sprintf("%s-%s-%s",
		strings.ToLower(valistFile.Org),
		strings.ToLower(valistFile.Repo),
		strings.ToLower(valistFile.Tag),
	)

	// If projectType is npm, run npm pack and set out to .tgz
	if valistFile.Type == "npm" {
		packageJsonPath := filepath.Join(projectPath, "package.json")
		packageJson, err := npm.ParsePackageJSON(packageJsonPath)
		if err != nil {
			return nil, err
		}

		valistFile.Build = valistFile.Build + " && npm pack"
		valistFile.Out = fmt.Sprintf("%s-%s.tgz", packageJson.Name, packageJson.Version)
	}

	// If image is not set in valist.yml use default image
	if valistFile.Image == "" {
		valistFile.Image = DefaultImages[valistFile.Type]
	}

	if valistFile.Type == "npm" && valistFile.Install == "" {
		valistFile.Install = DefaultInstalls[valistFile.Type]
	}

	// Create dockerConfig used to generate Dockerfile
	dockerFilePath := filepath.Join(projectPath, "Dockerfile.repro")
	dockerConfig := DockerConfig{
		Path:           dockerFilePath,
		BaseImage:      valistFile.Image,
		Source:         "./",
		BuildCommand:   valistFile.Build,
		InstallCommand: valistFile.Install,
	}

	if err := GenerateDockerfile(dockerConfig); err != nil {
		return nil, err
	}

	// Create the build image using the dockerfile
	if err := Create(buildImageName, dockerFilePath); err != nil {
		return nil, err
	}

	// Export the build from the build image
	if err := Export(buildImageName, projectPath, valistFile.Out); err != nil {
		return nil, err
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
