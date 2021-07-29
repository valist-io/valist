package impl

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	coreeth "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-ipfs/core/coreapi"
	coremock "github.com/ipfs/go-ipfs/core/mock"

	"github.com/valist-io/registry/internal/contract"
	"github.com/valist-io/registry/internal/core"
)

var _ core.CoreAPI = (*MockClient)(nil)

// MockClient is an in memory API client.
type MockClient struct {
	core.CoreAPI
	backend *backends.SimulatedBackend
}

// NewMockClient returns a new in memory API client.
func NewMockClient() (*MockClient, error) {
	private, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	chainID := big.NewInt(1337)
	address := crypto.PubkeyToAddress(private.PublicKey)
	backend := backends.NewSimulatedBackend(coreeth.GenesisAlloc{
		address: {Balance: big.NewInt(9223372036854775807)},
	}, 8000029)

	opts, err := bind.NewKeyedTransactorWithChainID(private, chainID)
	if err != nil {
		return nil, err
	}

	_, _, valist, err := contract.DeployValist(opts, backend, emptyAddress)
	if err != nil {
		return nil, err
	}

	_, _, registry, err := contract.DeployRegistry(opts, backend, emptyAddress)
	if err != nil {
		return nil, err
	}

	node, err := coremock.NewMockNode()
	if err != nil {
		return nil, err
	}

	ipfs, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return nil, err
	}

	// ensure contracts are deployed
	backend.Commit()

	return &MockClient{
		backend: backend,
		CoreAPI: &Client{
			eth:      backend,
			ipfs:     ipfs,
			orgs:     make(map[string]common.Hash),
			valist:   valist,
			registry: registry,
		},
	}, nil
}

func (c *MockClient) Commit() {
	c.backend.Commit()
}

func (c *MockClient) Close() {
	c.backend.Close()
}
