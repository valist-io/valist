package command

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/prompt"
)

// CreateAccount creates a new account.
func CreateAccount(ctx context.Context) error {
	config := ctx.Value(ConfigKey).(*config.Config)
	kstore := config.OpenKeyStore()

	passphrase, err := prompt.NewAccountPassphrase().Run()
	if err != nil {
		return err
	}

	account, err := kstore.NewAccount(passphrase)
	if err != nil {
		return err
	}

	if config.Accounts.Default == common.HexToAddress("0x0") {
		config.Accounts.Default = account.Address
	}

	config.Accounts.Pinned = append(config.Accounts.Pinned, account.Address)
	return config.Save()
}

// DefaultAccount sets the default account in the config.
func DefaultAccount(ctx context.Context, addr string) error {
	config := ctx.Value(ConfigKey).(*config.Config)
	kstore := config.OpenKeyStore()

	if !common.IsHexAddress(addr) {
		return fmt.Errorf("Invalid account address")
	}

	address := common.HexToAddress(addr)
	if !kstore.HasAddress(address) {
		return fmt.Errorf("Invalid account address")
	}

	config.Accounts.Default = address
	return config.Save()
}

// ExportAccount exports an account in web3 secret json format.
func ExportAccount(ctx context.Context, addr string) error {
	config := ctx.Value(ConfigKey).(*config.Config)
	kstore := config.OpenKeyStore()

	address := common.HexToAddress(addr)
	account := accounts.Account{Address: address}

	passphrase, err := prompt.AccountPassphrase().Run()
	if err != nil {
		return err
	}

	newPassphrase, err := prompt.NewAccountPassphrase().Run()
	if err != nil {
		return err
	}

	json, err := kstore.Export(account, passphrase, newPassphrase)
	if err != nil {
		return err
	}

	fmt.Println(string(json))
	return nil
}

// ImportAccount imports an account private key.
func ImportAccount(ctx context.Context) error {
	config := ctx.Value(ConfigKey).(*config.Config)
	kstore := config.OpenKeyStore()

	privkey, err := prompt.AccountPrivateKey().Run()
	if err != nil {
		return err
	}

	passphrase, err := prompt.AccountPassphrase().Run()
	if err != nil {
		return err
	}

	private, err := crypto.HexToECDSA(privkey)
	if err != nil {
		return err
	}

	account, err := kstore.ImportECDSA(private, passphrase)
	if err != nil {
		return err
	}

	if config.Accounts.Default == common.HexToAddress("0x0") {
		config.Accounts.Default = account.Address
	}

	fmt.Println("Successfully imported", account.Address)
	return config.Save()
}

// ListAccounts prints all account addresses.
func ListAccounts(ctx context.Context) error {
	config := ctx.Value(ConfigKey).(*config.Config)
	kstore := config.OpenKeyStore()

	for _, account := range kstore.Accounts() {
		if config.Accounts.Default == account.Address {
			fmt.Printf("%s (default)\n", account.Address)
		} else {
			fmt.Printf("%s\n", account.Address)
		}
	}

	return nil
}
