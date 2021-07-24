package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-ipfs/core/coreapi"
	coremock "github.com/ipfs/go-ipfs/core/mock"
	"github.com/stretchr/testify/suite"

	"github.com/valist-io/registry/internal/contract"
)

var chainID = big.NewInt(1337)

type CoreSuite struct {
	suite.Suite
	client  *Client
	address common.Address
	backend *backends.SimulatedBackend
}

func (s *CoreSuite) SetupTest() {
	private, err := crypto.GenerateKey()
	s.Require().NoError(err, "Failed to generate private key")

	s.address = crypto.PubkeyToAddress(private.PublicKey)
	s.backend = backends.NewSimulatedBackend(core.GenesisAlloc{
		s.address: {Balance: big.NewInt(9223372036854775807)},
	}, 8000029)

	opts, err := bind.NewKeyedTransactorWithChainID(private, chainID)
	s.Require().NoError(err, "Failed to create transactor")

	_, _, valist, err := contract.DeployValist(opts, s.backend, emptyAddress)
	s.Require().NoError(err, "Failed to deploy valist contract")

	_, _, registry, err := contract.DeployRegistry(opts, s.backend, emptyAddress)
	s.Require().NoError(err, "Failed to deploy valist registry contract")

	node, err := coremock.NewMockNode()
	s.Require().NoError(err, "Failed to create IPFS node")

	ipfs, err := coreapi.NewCoreAPI(node)
	s.Require().NoError(err, "Failed to create IPFS coreapi")

	s.client = NewClient(s.backend, ipfs, valist, registry, private, chainID)
}

func (s *CoreSuite) TearDownTest() {
	s.backend.Close()
}

func TestCoreSuite(t *testing.T) {
	suite.Run(t, &CoreSuite{})
}
