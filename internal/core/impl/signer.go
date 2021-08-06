package impl

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (client *Client) Signer(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	if address != client.account.Address {
		return nil, bind.ErrNotAuthorized
	}

	if client.metaTx {
		return tx, nil
	}

	return client.wallet.SignTx(client.account, tx, client.chainID)
}
