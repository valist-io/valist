package build

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type Dockerfile struct {
	path           string
	baseImage      string
	source         string
	buildCommand   string
	installCommand string
}

func generateDockerfile(buildConfig Dockerfile) {
	var dockerfile = fmt.Sprintf(
		"FROM %s\nWORKDIR /opt/build\nCOPY %s ./",
		buildConfig.baseImage, buildConfig.source,
	)

	if buildConfig.buildCommand != "" {
		dockerfile += fmt.Sprintf("\nRUN %s", buildConfig.buildCommand)
	}

	if buildConfig.installCommand != "" {
		dockerfile += fmt.Sprintf("\nRUN %s", buildConfig.installCommand)
	}

	os.WriteFile(buildConfig.path, []byte(dockerfile), 0644)
}

func CreateBuild(imageTag string) error {
	cmd := exec.Command(fmt.Sprintf("DOCKER_BUILDKIT=1 docker build -t %s .", imageTag))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Start()

	buf := bufio.NewReader(stdout)
	for {
		line, _, _ := buf.ReadLine()
		fmt.Println(string(line))
	}
}
