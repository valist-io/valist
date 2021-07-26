package account

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
)

func NewCreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create an account",
		Action: func(c *cli.Context) error {
			keyStore, err := config.GetKeyStore()
			if err != nil {
				return err
			}

			prompt := promptui.Prompt{
				Label: "Password",
				Mask:  '*',
			}

			password, err := prompt.Run()
			if err != nil {
				return err
			}

			acc, err := keyStore.NewAccount(password)
			if err != nil {
				return err
			}

			fmt.Println(acc.Address)
			return nil
		},
	}
}
