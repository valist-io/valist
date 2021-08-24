package basetx

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core/types"
)

type Transactor struct {
	valist   *valist.Valist
	registry *registry.ValistRegistry
}

func NewTransactor(valist *valist.Valist, registry *registry.ValistRegistry) types.TransactorAPI {
	return &Transactor{valist, registry}
}

// TransactOpts returns default transaction options.
func TransactOpts(account accounts.Account, wallet accounts.Wallet, chainID *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: account.Address,
		Signer: func(address common.Address, tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
			if address != account.Address {
				return nil, bind.ErrNotAuthorized
			}

			return wallet.SignTx(account, tx, chainID)
		},
	}
}
