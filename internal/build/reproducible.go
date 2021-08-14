package build

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type DockerConfig struct {
	Path           string
	BaseImage      string
	Source         string
	BuildCommand   string
	InstallCommand string
}

func GenerateDockerfile(dockerConfig DockerConfig) error {
	var dockerfile = fmt.Sprintf(
		"FROM %s\nWORKDIR /opt/build\nCOPY %s ./",
		dockerConfig.BaseImage, dockerConfig.Source,
	)

	if dockerConfig.InstallCommand != "" {
		dockerfile += fmt.Sprintf("\nRUN %s", dockerConfig.InstallCommand)
	}

	if dockerConfig.BuildCommand != "" {
		dockerfile += fmt.Sprintf("\nRUN %s", dockerConfig.BuildCommand)
	}

	return os.WriteFile(dockerConfig.Path, []byte(dockerfile), 0644)
}

func Create(imageTag string, dockerFilePath string) error {
	hostPath := filepath.Dir(dockerFilePath)
	ignoreFilePaths := [2]string{
		filepath.Join(hostPath, ".gitignore"),
		filepath.Join(hostPath, ".dockerignore"),
	}
	var fileToWrite []byte
	for _, file := range ignoreFilePaths {
		fileBytes, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		fileToWrite = append(fileToWrite, fileBytes...)
		fileToWrite = append(fileToWrite, byte(10)) // add new line, equiv to \n
	}
	if err := os.WriteFile(dockerFilePath+".dockerignore", fileToWrite, 0644); err != nil {
		return err
	}

	cmd := exec.Command("docker", "build", "-f", dockerFilePath, "-t", imageTag, "--progress=plain", hostPath)
	cmd.Env = append(os.Environ(), "DOCKER_BUILDKIT=1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Export(image string, hostPath string, out string) error {
	cpCmd := exec.Command("bash", "-c", fmt.Sprintf("docker run -v %s:/opt/out -i %s cp -R %s /opt/out", hostPath, image, out))
	cpCmd.Stdout = os.Stdout
	cpCmd.Stderr = os.Stderr
	if err := cpCmd.Run(); err != nil {
		return err
	}

	return nil
}
