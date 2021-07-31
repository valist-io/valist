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

func GenerateDockerfile(dockerConfig DockerConfig) {
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

	os.WriteFile(dockerConfig.Path, []byte(dockerfile), 0644)
}

func Create(imageTag string) error {
	cmd := exec.Command("docker", "build", ".", "-t", imageTag)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "DOCKER_BUILDKIT=1")

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", stdoutStderr)
	return nil
}

func Export(image string, out string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	hostPath := filepath.Join(cwd, out)
	containerPath := fmt.Sprintf("valist-build:/opt/build/%s", out)

	err = exec.Command("docker", "create", "--name=valist-build", "valist-build").Run()
	if err != nil {
		return err
	}

	err = exec.Command("docker", "cp", containerPath, hostPath).Run()
	if err != nil {
		return err
	}

	err = exec.Command("docker", "rm", "valist-build").Run()
	if err != nil {
		return err
	}
	return nil
}
