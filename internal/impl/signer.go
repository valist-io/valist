package impl

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (client *Client) Signer(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	if address != client.account.Address {
		return nil, bind.ErrNotAuthorized
	}

	return client.wallet.SignTx(client.account, tx, big.NewInt(1337))
}
