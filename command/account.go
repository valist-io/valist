package command

import (
	"context"
	"fmt"
	"os"

	"github.com/valist-io/valist/core"
	"github.com/valist-io/valist/prompt"
)

// CreateAccount creates a new account.
func CreateAccount(ctx context.Context) error {
	config := ctx.Value(ConfigKey).(*core.Config)
	client := ctx.Value(ClientKey).(*core.Client)

	passphrase, err := prompt.NewAccountPassphrase().Run()
	if err != nil {
		return err
	}
	addr, err := client.CreateAccount(passphrase)
	if err != nil {
		return err
	}
	logger.Info("Created account %s", addr)
	if config.GetDefaultAccount() != "" {
		return nil
	}
	config.SetDefaultAccount(addr)
	return config.Save()
}

// DefaultAccount sets the default account in the config.
func DefaultAccount(ctx context.Context, addr string) error {
	client := ctx.Value(ClientKey).(*core.Client)
	config := ctx.Value(ConfigKey).(*core.Config)

	if !client.HasAccount(addr) {
		return fmt.Errorf("account does not exist")
	}

	config.SetDefaultAccount(addr)
	return config.Save()
}

// ExportAccount exports an account in web3 secret json format.
func ExportAccount(ctx context.Context, addr string) error {
	client := ctx.Value(ClientKey).(*core.Client)

	passphrase, err := prompt.AccountPassphrase().Run()
	if err != nil {
		return err
	}
	newPassphrase, err := prompt.NewAccountPassphrase().Run()
	if err != nil {
		return err
	}
	json, err := client.ExportAccount(addr, passphrase, newPassphrase)
	if err != nil {
		return err
	}
	logger.Info(string(json))
	return nil
}

// ImportAccount imports an account private key.
func ImportAccount(ctx context.Context, fpath string) error {
	client := ctx.Value(ClientKey).(*core.Client)
	config := ctx.Value(ConfigKey).(*core.Config)

	data, err := os.ReadFile(fpath)
	if err != nil {
		return err
	}
	passphrase, err := prompt.AccountPassphrase().Run()
	if err != nil {
		return err
	}
	newPassphrase, err := prompt.NewAccountPassphrase().Run()
	if err != nil {
		return err
	}
	address, err := client.ImportAccount(data, passphrase, newPassphrase)
	if err != nil {
		return err
	}
	logger.Info("Successfully imported %s", address)
	if config.GetDefaultAccount() != "" {
		return nil
	}
	config.SetDefaultAccount(address)
	return config.Save()
}

// ListAccounts prints all account addresses.
func ListAccounts(ctx context.Context) error {
	client := ctx.Value(ClientKey).(*core.Client)
	config := ctx.Value(ConfigKey).(*core.Config)

	for _, addr := range client.ListAccounts() {
		if config.GetDefaultAccount() == addr {
			logger.Info("%s (default)", addr)
		} else {
			logger.Info("%s", addr)
		}
	}

	return nil
}
