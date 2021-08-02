package command

import (
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"github.com/valist-io/registry/internal/build"
)

func NewInitCommand() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Generate a new Valist project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "type",
				Usage: "Set project type",
			},
			&cli.BoolFlag{
				Name:  "interactive",
				Value: true,
				Usage: "Enable interactive mode",
			},
		},
		Action: func(c *cli.Context) error {
			// Get interactive flag value
			isInteractive := c.String("interactive")

			// If project type is not set ask for projectType
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

			if isInteractive == "true" {
				build.GenerateFileInteractive(projectType)
			} else {
				build.GenerateValistFile(projectType)
			}

			return nil
		},
	}
}
