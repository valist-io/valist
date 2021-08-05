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

	if dockerConfig.BuildCommand != "" {
		dockerfile += fmt.Sprintf("\nRUN %s", dockerConfig.BuildCommand)
	}

	if dockerConfig.InstallCommand != "" {
		dockerfile += fmt.Sprintf("\nRUN %s", dockerConfig.InstallCommand)
	}

	return os.WriteFile(dockerConfig.Path, []byte(dockerfile), 0644)
}

func Create(imageTag string, path string) error {
	if path != "" {
		path = "."
	}

	cmd := exec.Command("docker", "build", path, "-t", imageTag, "--progress=plain")
	cmd.Env = append(os.Environ(), "DOCKER_BUILDKIT=1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func Export(image string, out string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	hostPath := filepath.Join(cwd, out)
	// TODO fix for windows path
	containerPath := fmt.Sprintf("valist-build:/opt/build/%s/.", out)

	createCmd := exec.Command("docker", "create", "--name=valist-build", "valist-build")
	createCmd.Stdout = os.Stdout
	createCmd.Stderr = os.Stderr
	if err := createCmd.Run(); err != nil {
		return err
	}

	defer func() error {
		rmCmd := exec.Command("docker", "rm", "valist-build")
		rmCmd.Stdout = os.Stdout
		rmCmd.Stderr = os.Stderr
		return rmCmd.Run()
	}()

	cpCmd := exec.Command("docker", "cp", containerPath, hostPath)
	cpCmd.Stdout = os.Stdout
	cpCmd.Stderr = os.Stderr
	return cpCmd.Run()
}
