package lifecycle

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/prompt"
)

type contextKey string

const passphraseKey contextKey = "passphrase"

func loadConfig() (*config.Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	if err := config.Initialize(home); err != nil {
		return nil, err
	}

	cfg := config.NewConfig(home)

	if err := cfg.Load(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getAccount(c *cli.Context, cfg *config.Config) (accounts.Account, error) {
	var account accounts.Account
	if c.IsSet("account") {
		account.Address = common.HexToAddress(c.String("account"))
	} else {
		account.Address = cfg.Accounts.Default
	}

	// if no account found, prompt to create one
	if cfg.Accounts.Default == common.HexToAddress("0x0") {
		fmt.Println("No accounts found. Generating new keypair...")
		passphrase, err := prompt.NewAccountPassphrase().Run()
		if err != nil {
			return accounts.Account{}, err
		}

		account, err := cfg.KeyStore().NewAccount(passphrase)
		if err != nil {
			return accounts.Account{}, err
		}

		cfg.Accounts.Default = account.Address
		if err = cfg.Save(); err != nil {
			return accounts.Account{}, err
		}

		c.Context = context.WithValue(c.Context, passphraseKey, passphrase)
	}

	// use env variable VALIST_SIGNER as signing key when set
	if os.Getenv("VALIST_SIGNER") != "" {
		// generate cryptographically secure, ephemeral encryption key that will clear when program halts
		randBytes := make([]byte, 32)
		_, err := rand.Read(randBytes)
		if err != nil {
			return accounts.Account{}, err
		}

		passphrase := fmt.Sprintf("%x", randBytes)

		private, err := crypto.HexToECDSA(os.Getenv("VALIST_SIGNER"))
		if err != nil {
			return accounts.Account{}, err
		}

		// when VALIST_SIGNER is set, this encrypted key is stored in an os.TempDir()
		// when program halts, the ephemeral encryption key is cleared, making the keystore file inaccessible
		// this prevents having to clear the keystore file during runtime, which could fail to execute if program is interrupted/crashes
		// @TODO emulate filesystem in-memory when using VALIST_SIGNER to remove need for this workaround
		account, err = cfg.KeyStore().ImportECDSA(private, passphrase)
		if err != nil {
			return accounts.Account{}, err
		}

		c.Context = context.WithValue(c.Context, passphraseKey, passphrase)
	}

	return account, nil
}

func UnlockAccount(c *cli.Context) error {
	client := c.Context.Value(core.ClientKey).(*client.Client)

	// check for ephemeral encryption key (will be set if using VALIST_SIGNER env var)
	passphrase, _ := c.Context.Value(passphraseKey).(string)

	if passphrase == "" {
		pass, err := prompt.AccountPassphrase().RunFlag(c, "passphrase")
		if err != nil {
			return err
		}
		passphrase = pass
	}

	err := client.Signer().Unlock(client.Account(), passphrase)
	if err != nil {
		return err
	}

	return nil
}

func SetupClient(c *cli.Context) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	account, err := getAccount(c, cfg)
	if err != nil {
		return err
	}

	client, err := core.NewClient(c.Context, cfg, account)
	if err != nil {
		return err
	}

	c.Context = context.WithValue(c.Context, core.ClientKey, client)
	c.Context = context.WithValue(c.Context, core.ConfigKey, cfg)
	return nil
}

func SetupClientWithPassphrase(c *cli.Context) error {
	if err := SetupClient(c); err != nil {
		return err
	}

	return UnlockAccount(c)
}
