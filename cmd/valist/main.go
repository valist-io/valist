package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	app = cli.NewApp()
	// version set by linker flags
	version = "dev"
	// global flags
	globalFlags = []cli.Flag{
		&accountFlag,
		&passphraseFlag,
	}
)

func init() {
	app.Name = "valist"
	app.HelpName = "valist"
	app.Usage = "Valist command line interface"
	app.Description = `Universal package repository.`
	app.Copyright = "2020-2021 Valist, Inc."
	app.Version = version
	app.Flags = append(app.Flags, globalFlags...)
	app.Commands = []*cli.Command{
		&accountCommand,
		&createCommand,
		&daemonCommand,
		&getCommand,
		&initCommand,
		&installCommand,
		&keyCommand,
		&publishCommand,
		&thresholdCommand,
		&updateCommand,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
