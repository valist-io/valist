package signer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core"

	"github.com/valist-io/gasless"
	"github.com/valist-io/valist/internal/prompt"
)

type Signer struct {
	chainID *big.Int
	manager *accounts.Manager
	cache   map[common.Address]string
}

func NewSigner(chainID *big.Int, backends ...accounts.Backend) *Signer {
	return &Signer{
		chainID: chainID,
		manager: accounts.NewManager(&accounts.Config{}, backends...),
		cache:   make(map[common.Address]string),
	}
}

func (s *Signer) find(account accounts.Account) (gasless.Wallet, error) {
	wallet, err := s.manager.Find(account)
	if err != nil {
		return gasless.Wallet{}, err
	}

	return gasless.NewWallet(wallet), nil
}

func (s *Signer) passphrase(account accounts.Account) (string, error) {
	if passphrase, ok := s.cache[account.Address]; ok {
		return passphrase, nil
	}

	return prompt.AccountPassphrase().Run()
}

// SignTx signs the given transaction using the given account's wallet.
func (s *Signer) SignTx(account accounts.Account, tx *types.Transaction) (*types.Transaction, error) {
	wallet, err := s.find(account)
	if err != nil {
		return nil, err
	}

	out, err := wallet.SignTx(account, tx, s.chainID)
	switch v := err.(type) {
	case *accounts.AuthNeededError:
		fmt.Println(v.Needed)
	default:
		return out, err
	}

	passphrase, err := s.passphrase(account)
	if err != nil {
		return nil, err
	}

	out, err = wallet.SignTxWithPassphrase(account, passphrase, tx, s.chainID)
	if err != nil {
		return nil, err
	}

	s.cache[account.Address] = passphrase
	return out, nil
}

// SignTypedData signs the given typedData using the given account's wallet.
func (s *Signer) SignTypedData(account accounts.Account, typedData core.TypedData) ([]byte, error) {
	wallet, err := s.find(account)
	if err != nil {
		return nil, err
	}

	out, err := wallet.SignTypedData(account, typedData)
	switch v := err.(type) {
	case *accounts.AuthNeededError:
		fmt.Println(v.Needed)
	default:
		return out, err
	}

	passphrase, err := s.passphrase(account)
	if err != nil {
		return nil, err
	}

	out, err = wallet.SignTypedDataWithPassphrase(account, passphrase, typedData)
	if err != nil {
		return nil, err
	}

	s.cache[account.Address] = passphrase
	return out, nil
}

// NewTransactor returns TransactOpts for meta or regular transactions.
func (s *Signer) NewTransactor(account accounts.Account) *gasless.TransactOpts {
	return &gasless.TransactOpts{
		TransactOpts: bind.TransactOpts{
			Context: context.Background(),
			From:    account.Address,
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				if address != account.Address {
					return nil, bind.ErrNotAuthorized
				}

				return s.SignTx(account, tx)
			},
		},
		MetaSigner: func(address common.Address, typedData core.TypedData) ([]byte, error) {
			if address != account.Address {
				return nil, bind.ErrNotAuthorized
			}

			return s.SignTypedData(account, typedData)
		},
	}
}
