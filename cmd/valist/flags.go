package main

import (
	"github.com/urfave/cli/v2"
)

var (
	// accountFlag sets the account by address
	accountFlag = cli.StringFlag{
		Name:  "account",
		Usage: "Account to transact with",
	}
	// passphraseFlag sets the account passphrase
	passphraseFlag = cli.StringFlag{
		Name:  "passphrase",
		Usage: "Passphrase to unlock account",
	}
	// publishDryRunFlag does all steps but skips publishing
	publishDryRunFlag = cli.BoolFlag{
		Name:  "dryrun",
		Usage: "Skip publish",
	}
)
