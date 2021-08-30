package client

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/storage"
)

var emptyHash = common.HexToHash("0x0")

// Close is a callback invoked when the client is closed.
type Close func() error

// TransactOpts is a function that returns transaction options for an Ethereum transaction.
type TransactOpts func(account accounts.Account, wallet accounts.Wallet, chainID *big.Int) *bind.TransactOpts

// TransactorAPI defines functions to abstract blockchain transactions.
// TODO: Maybe this can return []*types.Log instead of *types.Transaction and handle waiting and log parsing?
type TransactorAPI interface {
	CreateOrganizationTx(*bind.TransactOpts, string) (*types.Transaction, error)
	LinkOrganizationNameTx(*bind.TransactOpts, common.Hash, string) (*types.Transaction, error)
	CreateRepositoryTx(*bind.TransactOpts, common.Hash, string, string) (*types.Transaction, error)
	VoteReleaseTx(*bind.TransactOpts, common.Hash, string, string, string, string) (*types.Transaction, error)
	SetRepositoryMetaTx(*bind.TransactOpts, common.Hash, string, string) (*types.Transaction, error)
	VoteOrganizationThresholdTx(*bind.TransactOpts, common.Hash, *big.Int) (*types.Transaction, error)
	VoteRepositoryThresholdTx(*bind.TransactOpts, common.Hash, string, *big.Int) (*types.Transaction, error)
}

// Options is used to set client options.
type Options struct {
	Storage  storage.Storage
	Ethereum bind.DeployBackend
	ChainID  *big.Int

	Valist   *valist.Valist
	Registry *registry.ValistRegistry

	Account accounts.Account
	Wallet  accounts.Wallet

	TransactOpts TransactOpts
	Transactor   TransactorAPI

	OnClose []Close
}

// Client is a Valist SDK client.
type Client struct {
	eth     bind.DeployBackend
	storage storage.Storage
	chainID *big.Int

	valist   *valist.Valist
	registry *registry.ValistRegistry

	wallet  accounts.Wallet
	account accounts.Account

	transactor   TransactorAPI
	transactOpts TransactOpts

	onClose []Close
	orgs    map[string]common.Hash
}

// NewClient create a client from the given options.
func NewClient(opts *Options) (*Client, error) {
	if opts.Ethereum == nil {
		return nil, fmt.Errorf("ethereum client is required")
	}

	if opts.Storage == nil {
		return nil, fmt.Errorf("storage is required")
	}

	if opts.Valist == nil {
		return nil, fmt.Errorf("valist contract is required")
	}

	if opts.Registry == nil {
		return nil, fmt.Errorf("registry contract is required")
	}

	if opts.TransactOpts == nil {
		return nil, fmt.Errorf("transact opts is required")
	}

	if opts.Transactor == nil {
		return nil, fmt.Errorf("transactor is required")
	}

	return &Client{
		eth:          opts.Ethereum,
		storage:      opts.Storage,
		chainID:      opts.ChainID,
		valist:       opts.Valist,
		registry:     opts.Registry,
		wallet:       opts.Wallet,
		account:      opts.Account,
		transactor:   opts.Transactor,
		transactOpts: opts.TransactOpts,
		onClose:      opts.OnClose,
		orgs:         make(map[string]common.Hash),
	}, nil
}

// Close releases all client resources.
func (client *Client) Close() {
	for _, close := range client.onClose {
		close() // nolint:errcheck
	}
}

func (client *Client) Storage() storage.Storage {
	return client.storage
}
