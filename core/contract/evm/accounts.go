package evm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	scryptN = keystore.StandardScryptN
	scryptP = keystore.StandardScryptP
)

type AccountManager struct {
	account  accounts.Account
	keystore *keystore.KeyStore
	signer   types.Signer
	// cache for passwords so we don't prompt twice
	// this should only be in memory for one command
	passwords map[common.Address]string
}

// NewAccountManager returns an account manager
// backed by the keystore at the given path.
func NewAccountManager(path string, chainID *big.Int) *AccountManager {
	return &AccountManager{
		signer:    types.LatestSignerForChainID(chainID),
		keystore:  keystore.NewKeyStore(path, scryptN, scryptP),
		passwords: make(map[common.Address]string),
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

// SetAccount sets the current account with an optional password.
func (a *AccountManager) SetAccount(address, password string) error {
	addr := common.HexToAddress(address)
	acct, err := a.keystore.Find(accounts.Account{Address: addr})
	if err != nil {
		return err
	}
	a.account = acct
	a.passwords[addr] = password
	return nil
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

func (a *AccountManager) signTx(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	if address != a.account.Address {
		return nil, fmt.Errorf("not authorized")
	}
	signature, err := a.keystore.SignHash(a.account, a.signer.Hash(tx).Bytes())
	if err != nil {
		return nil, err
	}
	// TODO retry with password
	return tx.WithSignature(a.signer, signature)
}
