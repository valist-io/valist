package signer

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"

	"github.com/valist-io/gasless"
	"github.com/valist-io/valist/prompt"
)

const EnvKey = "VALIST_SIGNER"

// Signer signs transactions.
type Signer struct {
	account accounts.Account
	chainID *big.Int
	manager *accounts.Manager
	cache   map[common.Address]string
}

// NewSigner returns a signer that uses the given backends to sign transactions.
// If VALIST_SIGNER env is set, an ephemeral keystore is created and the default account is set.
func NewSigner(chainID *big.Int, backends ...accounts.Backend) (*Signer, error) {
	signer := &Signer{
		chainID: chainID,
		manager: accounts.NewManager(&accounts.Config{}, backends...),
		cache:   make(map[common.Address]string),
	}

	privatekey := os.Getenv(EnvKey)
	if privatekey == "" {
		return signer, nil
	}

	private, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		return nil, err
	}

	randBytes := make([]byte, 32)
	if _, err := rand.Read(randBytes); err != nil {
		return nil, err
	}

	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}

	kstore := keystore.NewKeyStore(tmp, keystore.StandardScryptN, keystore.StandardScryptP)
	signer.manager.AddBackend(kstore)

	passphrase := fmt.Sprintf("%x", randBytes)
	account, err := kstore.ImportECDSA(private, passphrase)
	if err != nil {
		return nil, err
	}

	if err := kstore.Unlock(account, passphrase); err != nil {
		return nil, err
	}

	signer.account = account
	return signer, nil
}

// List returns a list of all signer accounts.
func (s *Signer) List() []accounts.Account {
	var list []accounts.Account
	for _, wallet := range s.manager.Wallets() {
		list = append(list, wallet.Accounts()...)
	}
	return list
}

// Account returns the default signer account.
func (s *Signer) Account() accounts.Account {
	return s.account
}

// SetAccount sets the default signer account.
func (s *Signer) SetAccount(account accounts.Account) {
	s.account = account
}

// SetAccountWithPassphrase sets the default signer account and passphrase to unlock the account.
func (s *Signer) SetAccountWithPassphrase(account accounts.Account, passphrase string) {
	s.account = account
	s.cache[account.Address] = passphrase
}

// SignTx signs the given transaction using the given account's wallet.
func (s *Signer) SignTx(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	if address != s.account.Address {
		return nil, bind.ErrNotAuthorized
	}

	wallet, err := s.wallet(s.account)
	if err != nil {
		return nil, err
	}

	out, err := wallet.SignTx(s.account, tx, s.chainID)
	if _, ok := err.(*accounts.AuthNeededError); !ok {
		return out, err
	}

	passphrase, err := s.passphrase(s.account)
	if err != nil {
		return nil, err
	}

	out, err = wallet.SignTxWithPassphrase(s.account, passphrase, tx, s.chainID)
	if err != nil {
		return nil, err
	}

	// remember the passphrase only after successful transaction
	s.cache[s.account.Address] = passphrase
	return out, nil
}

// SignTypedData signs the given typedData using the given account's wallet.
func (s *Signer) SignTypedData(address common.Address, typedData core.TypedData) ([]byte, error) {
	if address != s.account.Address {
		return nil, bind.ErrNotAuthorized
	}

	wallet, err := s.wallet(s.account)
	if err != nil {
		return nil, err
	}

	out, err := wallet.SignTypedData(s.account, typedData)
	if _, ok := err.(*accounts.AuthNeededError); !ok {
		return out, err
	}

	passphrase, err := s.passphrase(s.account)
	if err != nil {
		return nil, err
	}

	out, err = wallet.SignTypedDataWithPassphrase(s.account, passphrase, typedData)
	if err != nil {
		return nil, err
	}

	// remember the passphrase only after successful transaction
	s.cache[s.account.Address] = passphrase
	return out, nil
}

// NewTransactor returns TransactOpts for meta or regular transactions.
func (s *Signer) NewTransactor() *gasless.TransactOpts {
	return &gasless.TransactOpts{
		MetaSigner: s.SignTypedData,
		TransactOpts: bind.TransactOpts{
			Context: context.Background(),
			From:    s.account.Address,
			Signer:  s.SignTx,
		},
	}
}

func (s *Signer) wallet(account accounts.Account) (gasless.Wallet, error) {
	wallet, err := s.manager.Find(account)
	if err != nil {
		return gasless.Wallet{}, err
	}

	return gasless.NewWallet(wallet), nil
}

func (s *Signer) passphrase(account accounts.Account) (string, error) {
	passphrase, ok := s.cache[account.Address]
	if ok && passphrase != "" {
		return passphrase, nil
	}

	return prompt.AccountPassphrase().Run()
}
