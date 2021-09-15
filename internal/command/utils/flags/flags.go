package flags

import (
	"github.com/urfave/cli/v2"
)

func Account() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "account",
		Usage: "Account to transact with",
	}
}

func AccountPassphrase() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "passphrase",
		Usage: "Passphrase to unlock account",
	}
}

func AccountPrivateKey() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "key",
		Usage: "Hex-encoded ECDSA private key",
	}
}