package main

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/command"
)

var (
	// accountCommand manages accounts
	accountCommand = cli.Command{
		Name:  "account",
		Usage: "Create, update, or list accounts",
		Subcommands: []*cli.Command{
			&accountCreateCommand,
			&accountDefaultCommand,
			&accountExportCommand,
			&accountImportCommand,
			&accountListCommand,
		},
	}
	// accountCreateCommand creates an account
	accountCreateCommand = cli.Command{
		Name:  "create",
		Usage: "Create an account",
		Before: func(c *cli.Context) error {
			return setupConfig(c)
		},
		Action: func(c *cli.Context) error {
			return command.CreateAccount(c.Context)
		},
	}
	// accountDefaultCommand sets the default account
	accountDefaultCommand = cli.Command{
		Name:      "default",
		Usage:     "Set the default account",
		ArgsUsage: "[address]",
		Before: func(c *cli.Context) error {
			return setupConfig(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.DefaultAccount(c.Context, c.Args().Get(0))
		},
	}
	// accountExportCommand exports an existing account
	accountExportCommand = cli.Command{
		Name:      "export",
		Usage:     "Export an account",
		ArgsUsage: "[address]",
		Before: func(c *cli.Context) error {
			return setupConfig(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.ExportAccount(c.Context, c.Args().Get(0))
		},
	}
	// accountImportCommand imports an account private key
	accountImportCommand = cli.Command{
		Name:  "import",
		Usage: "Import an account",
		Before: func(c *cli.Context) error {
			return setupConfig(c)
		},
		Action: func(c *cli.Context) error {
			return command.ImportAccount(c.Context)
		},
	}
	// accountListCommand prints all account addresses
	accountListCommand = cli.Command{
		Name:  "list",
		Usage: "List all accounts",
		Before: func(c *cli.Context) error {
			return setupConfig(c)
		},
		Action: func(c *cli.Context) error {
			return command.ListAccounts(c.Context)
		},
	}
	// createCommand creates an org or repo
	createCommand = cli.Command{
		Name:      "create",
		Usage:     "Create an organization or repository",
		ArgsUsage: "[name]",
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.Create(c.Context, c.Args().Get(0))
		},
	}
	// daemonCommand starts a valist node
	daemonCommand = cli.Command{
		Name:  "daemon",
		Usage: "Runs a persistent valist node",
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			return command.Daemon(c.Context)
		},
	}
	// getCommand downloads a release file
	getCommand = cli.Command{
		Name:      "get",
		Usage:     "Download a release file",
		ArgsUsage: "[name] [file]",
		Flags: []cli.Flag{
			&outputFlag,
		},
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.Get(c.Context, c.Args().Get(0), c.Args().Get(1), c.String("output"))
		},
	}
	// initCommand initializes a valist.yml
	initCommand = cli.Command{
		Name:  "init",
		Usage: "Initialize a new valist project",
		Flags: []cli.Flag{
			&initWizardFlag,
		},
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.Init(c.Context, c.Args().Get(0), c.Bool("wizard"))
		},
	}
	// installCommand installs a binary artifact
	installCommand = cli.Command{
		Name:      "install",
		Usage:     "Installs a binary artifact",
		ArgsUsage: "[name]",
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.Install(c.Context, c.Args().Get(0))
		},
	}
	// keyCommand manages organization and repository keys
	keyCommand = cli.Command{
		Name:  "key",
		Usage: "Add, revoke, or rotate organization or repository keys",
		Subcommands: []*cli.Command{
			&keyAddCommand,
			&keyRevokeCommand,
			&keyRotateCommand,
		},
	}
	// keyAddCommand adds an organization or repository key
	keyAddCommand = cli.Command{
		Name:      "add",
		Usage:     "Add organization or repository key",
		ArgsUsage: "[path] [address]",
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.KeyAdd(c.Context, c.Args().Get(0), c.Args().Get(1))
		},
	}
	// keyRevokeCommand revokes an organization or repository key
	keyRevokeCommand = cli.Command{
		Name:      "revoke",
		Usage:     "Revoke organization or repository key",
		ArgsUsage: "[path] [address]",
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.KeyRevoke(c.Context, c.Args().Get(0), c.Args().Get(1))
		},
	}
	// keyRotateCommand rotates an organization or repository key
	keyRotateCommand = cli.Command{
		Name:      "rotate",
		Usage:     "Rotate organization or repository key",
		ArgsUsage: "[path] [address]",
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.KeyRotate(c.Context, c.Args().Get(0), c.Args().Get(1))
		},
	}
	// listCommand prints org, repo, or release contents
	listCommand = cli.Command{
		Name:      "list",
		Usage:     "List organization, repository, or release contents",
		ArgsUsage: "[name]",
		Aliases:   []string{"ls"},
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.List(c.Context, c.Args().Get(0))
		},
	}
	// publishCommand creates a new release
	publishCommand = cli.Command{
		Name:  "publish",
		Usage: "Publish a new release",
		Flags: []cli.Flag{
			&publishDryRunFlag,
		},
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			return command.Publish(c.Context, c.Bool("dryrun"))
		},
	}
	// thresholdCommand votes to change signature threshold
	thresholdCommand = cli.Command{
		Name:      "threshold",
		Usage:     "Vote to update organization or repository signature threshold",
		ArgsUsage: "[path] [count]",
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			threshold, err := strconv.ParseInt(c.Args().Get(1), 10, 64)
			if err != nil {
				return fmt.Errorf("invalid threshold: %v", err)
			}

			return command.Threshold(c.Context, c.Args().Get(0), threshold)
		},
	}
	// updateCommand updates organization or repository metadata
	updateCommand = cli.Command{
		Name:      "update",
		Usage:     "Update organization or repository metadata",
		ArgsUsage: "[name]",
		Before: func(c *cli.Context) error {
			return setup(c)
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

			return command.Update(c.Context, c.Args().Get(0))
		},
	}
)
