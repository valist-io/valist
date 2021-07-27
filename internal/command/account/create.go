package account

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an account",
		Action: func(c *cli.Context) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg, err := config.Load(home)
			if err != nil {
				return err
			}

			prompt := promptui.Prompt{
				Label:       "Password",
				Mask:        '*',
				HideEntered: true,
			}

			password, err := prompt.Run()
			if err != nil {
				return err
			}

			acc, err := cfg.KeyStore().NewAccount(password)
			if err != nil {
				return err
			}

			fmt.Println(acc.Address)
			return nil
		},
	}
}
