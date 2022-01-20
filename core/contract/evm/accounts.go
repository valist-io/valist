package evm

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/valist-io/valist/prompt"
)

const (
	scryptN = keystore.StandardScryptN
	scryptP = keystore.StandardScryptP
)

type AccountManager struct {
	account  accounts.Account
	keystore *keystore.KeyStore
	signer   types.Signer
}

// NewAccountManager returns an account manager
// backed by the keystore at the given path.
func NewAccountManager(path string, chainID *big.Int) *AccountManager {
	return &AccountManager{
		signer:   types.LatestSignerForChainID(chainID),
		keystore: keystore.NewKeyStore(path, scryptN, scryptP),
	}
}

// GetAccount returns the current account.
func (a *AccountManager) GetAccount() string {
	return a.account.Address.Hex()
}

// HasAccount returns true if the account exists.
func (a *AccountManager) HasAccount(address string) bool {
	return a.keystore.HasAddress(common.HexToAddress(address))
}

// SetAccount sets the current account and optionally unlocks it with the passphrase.
func (a *AccountManager) SetAccount(address, passphrase string) (err error) {
	addr := common.HexToAddress(address)
	acct := accounts.Account{Address: addr}

	a.account, err = a.keystore.Find(acct)
	if err != nil {
		return err
	}
	if passphrase == "" {
		return nil
	}
	return a.keystore.TimedUnlock(a.account, passphrase, 0)
}

// CreateAccount creates a new account encrypted with the given password.
func (a *AccountManager) CreateAccount(password string) (string, error) {
	account, err := a.keystore.NewAccount(password)
	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}

// ListAccounts returns a list of all accounts.
func (a *AccountManager) ListAccounts() []string {
	var accounts []string
	for _, acct := range a.keystore.Accounts() {
		accounts = append(accounts, acct.Address.Hex())
	}
	return accounts
}

// ImportAccount imports account data. Format depends on the underlying implementation.
func (a *AccountManager) ImportAccount(data []byte, password, newPassword string) (string, error) {
	acct, err := a.keystore.Import(data, password, newPassword)
	if err != nil {
		return "", err
	}
	return acct.Address.Hex(), nil
}

// ExportAccount exports account data. Format depends on the underlying implementation.
func (a *AccountManager) ExportAccount(address, password, newPassword string) ([]byte, error) {
	addr := common.HexToAddress(address)
	acct := accounts.Account{Address: addr}
	return a.keystore.Export(acct, password, newPassword)
}

func (a *AccountManager) SignTx(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	if address != a.account.Address {
		return nil, fmt.Errorf("not authorized")
	}
	signature, err := a.keystore.SignHash(a.account, a.signer.Hash(tx).Bytes())
	if err == nil {
		return tx.WithSignature(a.signer, signature)
	}
	if err != keystore.ErrLocked {
		return nil, err
	}
SIGN_TX:
	passphrase, err := prompt.AccountPassphrase().Run()
	if err != nil {
		return nil, err
	}
	err = a.keystore.TimedUnlock(a.account, passphrase, 1*time.Minute)
	if err != nil {
		goto SIGN_TX
	}
	signature, err = a.keystore.SignHashWithPassphrase(a.account, passphrase, a.signer.Hash(tx).Bytes())
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(a.signer, signature)
}
