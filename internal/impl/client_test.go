package impl

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	coreeth "github.com/ethereum/go-ethereum/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	coremock "github.com/ipfs/go-ipfs/core/mock"
	"github.com/stretchr/testify/suite"

	"github.com/valist-io/registry/internal/contract"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
	passphrase       = "secret"
)

type ClientSuite struct {
	suite.Suite
	tmp     string
	client  *Client
	backend *backends.SimulatedBackend
}

func (s *ClientSuite) SetupTest() {
	tmp, err := os.MkdirTemp("", "test")
	s.Require().NoError(err, "Failed to make tmp dir")

	signer := keystore.NewKeyStore(tmp, veryLightScryptN, veryLightScryptP)
	account, err := signer.NewAccount(passphrase)
	s.Require().NoError(err, "Failed to create keystore account")

	err = signer.Unlock(account, passphrase)
	s.Require().NoError(err, "Failed to unlock keystore account")

	chainID := big.NewInt(1337)
	backend := backends.NewSimulatedBackend(coreeth.GenesisAlloc{
		account.Address: {Balance: big.NewInt(9223372036854775807)},
	}, 8000029)

	opts, err := bind.NewKeyStoreTransactorWithChainID(signer, account, chainID)
	s.Require().NoError(err, "Failed to create transactor")

	_, _, valist, err := contract.DeployValist(opts, backend, emptyAddress)
	s.Require().NoError(err, "Failed to deploy valist contract")

	_, _, registry, err := contract.DeployRegistry(opts, backend, emptyAddress)
	s.Require().NoError(err, "Failed to deploy registry contract")

	node, err := coremock.NewMockNode()
	s.Require().NoError(err, "Failed to create IPFS mock node")

	ipfs, err := coreapi.NewCoreAPI(node)
	s.Require().NoError(err, "Failed to create IPFS core API")

	// ensure contracts are deployed
	s.backend = backend
	s.backend.Commit()

	s.tmp = tmp
	s.client = &Client{
		eth:      backend,
		ipfs:     ipfs,
		orgs:     make(map[string]common.Hash),
		valist:   valist,
		registry: registry,
		wallet:   signer.Wallets()[0],
		account:  account,
	}
}

func (s *ClientSuite) TearDownTest() {
	s.backend.Close()
	os.RemoveAll(s.tmp)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}
