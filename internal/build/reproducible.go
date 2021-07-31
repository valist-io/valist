package build

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Dockerfile struct {
	Path           string
	BaseImage      string
	Source         string
	BuildCommand   string
	InstallCommand string
}

func GenerateDockerfile(buildConfig Dockerfile) {
	var dockerfile = fmt.Sprintf(
		"FROM %s\nWORKDIR /opt/build\nCOPY %s ./",
		buildConfig.BaseImage, buildConfig.Source,
	)

	if buildConfig.BuildCommand != "" {
		dockerfile += fmt.Sprintf("\nRUN %s", buildConfig.BuildCommand)
	}

	if buildConfig.InstallCommand != "" {
		dockerfile += fmt.Sprintf("\nRUN %s", buildConfig.InstallCommand)
	}

	os.WriteFile(buildConfig.Path, []byte(dockerfile), 0644)
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
