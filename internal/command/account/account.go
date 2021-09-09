package account

import (
	"github.com/urfave/cli/v2"
)

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "account",
		Usage: "Create, update, or list accounts",
		Subcommands: []*cli.Command{
			NewCreateCommand(),
			NewListCommand(),
			NewDefaultCommand(),
			NewPinCommand(),
			NewUnpinCommand(),
		},
	}
}
