package client

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/valist-io/gasless"
	"github.com/valist-io/valist/contract/registry"
	"github.com/valist-io/valist/contract/valist"
	"github.com/valist-io/valist/signer"
	"github.com/valist-io/valist/storage"
)

var (
	emptyHash  = common.HexToHash("0x0")
	ORG_ADMIN  = crypto.Keccak256Hash([]byte("ORG_ADMIN_ROLE"))
	REPO_DEV   = crypto.Keccak256Hash([]byte("REPO_DEV_ROLE"))
	ADD_KEY    = crypto.Keccak256Hash([]byte("ADD_KEY_OPERATION"))
	REVOKE_KEY = crypto.Keccak256Hash([]byte("REVOKE_KEY_OPERATION"))
	ROTATE_KEY = crypto.Keccak256Hash([]byte("ROTATE_KEY_OPERATION"))
)

// TransactorAPI defines functions to abstract blockchain transactions.
// TODO: Maybe this can return []*types.Log instead of *types.Transaction and handle waiting and log parsing?
type TransactorAPI interface {
	CreateOrganizationTx(*gasless.TransactOpts, string) (*types.Transaction, error)
	LinkOrganizationNameTx(*gasless.TransactOpts, common.Hash, string) (*types.Transaction, error)
	CreateRepositoryTx(*gasless.TransactOpts, common.Hash, string, string) (*types.Transaction, error)
	VoteKeyTx(*gasless.TransactOpts, common.Hash, string, common.Hash, common.Address) (*types.Transaction, error)
	VoteReleaseTx(*gasless.TransactOpts, common.Hash, string, string, string, string) (*types.Transaction, error)
	SetOrganizationMetaTx(*gasless.TransactOpts, common.Hash, string) (*types.Transaction, error)
	SetRepositoryMetaTx(*gasless.TransactOpts, common.Hash, string, string) (*types.Transaction, error)
	VoteOrganizationThresholdTx(*gasless.TransactOpts, common.Hash, *big.Int) (*types.Transaction, error)
	VoteRepositoryThresholdTx(*gasless.TransactOpts, common.Hash, string, *big.Int) (*types.Transaction, error)
}

// Options is used to set client options.
type Options struct {
	Storage  storage.Provider
	Ethereum bind.DeployBackend

	Valist   *valist.Valist
	Registry *registry.ValistRegistry

	Signer     *signer.Signer
	Transactor TransactorAPI
}

// Client is a Valist SDK client.
type Client struct {
	eth     bind.DeployBackend
	storage storage.Provider

	valist   *valist.Valist
	registry *registry.ValistRegistry

	signer     *signer.Signer
	transactor TransactorAPI

	orgs map[string]common.Hash
}

// NewClient create a client from the given options.
func NewClient(opts Options) (*Client, error) {
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

	if opts.Transactor == nil {
		return nil, fmt.Errorf("transactor is required")
	}

	if opts.Signer == nil {
		return nil, fmt.Errorf("signer is required")
	}

	return &Client{
		eth:        opts.Ethereum,
		storage:    opts.Storage,
		valist:     opts.Valist,
		registry:   opts.Registry,
		signer:     opts.Signer,
		transactor: opts.Transactor,
		orgs:       make(map[string]common.Hash),
	}, nil
}

func (client *Client) Storage() storage.Provider {
	return client.storage
}

func (client *Client) Signer() *signer.Signer {
	return client.signer
}

func (client *Client) Close() error {
	return nil
}