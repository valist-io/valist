package client

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	coreiface "github.com/ipfs/interface-go-ipfs-core"

	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core/types"
)

var (
	emptyHash  = common.HexToHash("0x0")
	ORG_ADMIN  = crypto.Keccak256Hash([]byte("ORG_ADMIN_ROLE"))
	REPO_DEV   = crypto.Keccak256Hash([]byte("REPO_DEV_ROLE"))
	ADD_KEY    = crypto.Keccak256Hash([]byte("ADD_KEY_OPERATION"))
	REVOKE_KEY = crypto.Keccak256Hash([]byte("REVOKE_KEY_OPERATION"))
	ROTATE_KEY = crypto.Keccak256Hash([]byte("ROTATE_KEY_OPERATION"))
)

// Close is a callback invoked when the client is closed.
type Close func() error

// TransactOpts is a function that returns transaction options for an Ethereum transaction.
type TransactOpts func(account accounts.Account, wallet accounts.Wallet, chainID *big.Int) *bind.TransactOpts

// Options is used to set client options.
type Options struct {
	IPFS     coreiface.CoreAPI
	Ethereum bind.DeployBackend
	ChainID  *big.Int

	Valist   *valist.Valist
	Registry *registry.ValistRegistry

	Account accounts.Account
	Wallet  accounts.Wallet

	TransactOpts TransactOpts
	Transactor   types.TransactorAPI

	OnClose []Close
}

// Client is a Valist SDK client.
type Client struct {
	eth     bind.DeployBackend
	ipfs    coreiface.CoreAPI
	chainID *big.Int

	valist   *valist.Valist
	registry *registry.ValistRegistry

	wallet  accounts.Wallet
	account accounts.Account

	transactor   types.TransactorAPI
	transactOpts TransactOpts

	onClose []Close
	orgs    map[string]common.Hash
}

// NewClient create a client from the given options.
func NewClient(opts *Options) (*Client, error) {
	if opts.Ethereum == nil {
		return nil, fmt.Errorf("ethereum client is required")
	}

	if opts.IPFS == nil {
		return nil, fmt.Errorf("ipfs client is required")
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
		ipfs:         opts.IPFS,
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

func (client *Client) SwitchAccount(account accounts.Account, wallet accounts.Wallet) {
	client.account = account
	client.wallet = wallet
}
