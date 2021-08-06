package metatx

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/valist-io/registry/internal/contract"
	"github.com/valist-io/registry/internal/contract/metatx"
	"github.com/valist-io/registry/internal/core"
)

var _ core.CoreTransactorAPI = (*Client)(nil)

var (
	emptyHash    = common.HexToHash("0x0")
	emptyAddress = common.HexToAddress("0x0")
)

type Client struct {
	super core.CoreTransactorAPI

	eth       *ethclient.Client
	forwarder *metatx.Metatx

	chainID *big.Int
	address common.Address

	account accounts.Account
	wallet  accounts.Wallet
}

func NewClient(
	client core.CoreTransactorAPI,
	eth *ethclient.Client,
	address common.Address,
	chainID *big.Int,
	account accounts.Account,
	wallet accounts.Wallet,
) (core.CoreTransactorAPI, error) {
	forwarder, err := contract.NewForwarder(address, eth)
	if err != nil {
		return nil, err
	}

	return &Client{
		super:     client,
		eth:       eth,
		address:   address,
		chainID:   chainID,
		forwarder: forwarder,
		account:   account,
		wallet:    wallet,
	}, nil
}

func (client *Client) CreateOrganizationTx(ctx context.Context, meta *core.OrganizationMeta) (*types.Transaction, error) {
	tx, err := client.super.CreateOrganizationTx(ctx, meta)
	if err != nil {
		return nil, err
	}

	return client.SendMetaTx(tx, createOrganizationBFID)
}
