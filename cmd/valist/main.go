package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/log"
)

var logger = log.New()

var (
	app = cli.NewApp()
	// version set by linker flags
	Version = "dev"
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
	app.Description = `Web3-native software distribution`
	app.Copyright = "2020-2021 Valist, Inc."
	app.Version = Version
	app.Flags = append(app.Flags, globalFlags...)
	app.EnableBashCompletion = true
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
		logger.Error("%v", err)
	}
}
