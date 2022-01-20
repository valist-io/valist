package core

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/valist-io/valist/core/contract/evm"
	"github.com/valist-io/valist/core/contract/evm/valist"
	"github.com/valist-io/valist/core/storage/ipfs"
)

type ClientSuite struct {
    suite.Suite
    client *Client
}

func (s *ClientSuite) SetupTest() {
	s.client = setupEVM(s.T())
}

func (s *ClientSuite) TestGetTeam() {
	ctx := context.Background()
    _, err := s.client.GetTeam(ctx, "team")
    require.NoError(s.T(), err, "failed to get team")
}

func (s *ClientSuite) TestGetProject() {
	ctx := context.Background()
    _, err := s.client.GetProject(ctx, "team", "proj")
    require.NoError(s.T(), err, "failed to get project")
}

func (s *ClientSuite) TestGetRelease() {
	ctx := context.Background()
    _, err := s.client.GetRelease(ctx, "team", "proj", "rel")
    require.NoError(s.T(), err, "failed to get release")
}

func (s *ClientSuite) TestGetLatestRelease() {
	ctx := context.Background()
    _, err := s.client.GetLatestRelease(ctx, "team", "proj")
    require.NoError(s.T(), err, "failed to get latest release")
}

func (s *ClientSuite) TestResolvePath() {
	ctx := context.Background()
    _, err := s.client.ResolvePath(ctx, "team/proj/rel")
    require.NoError(s.T(), err, "failed to resolve path")
}

func TestEvmContracts(t *testing.T) {
    suite.Run(t, &ClientSuite{})
}

func setupEVM(t *testing.T) *Client {
	storage, err := ipfs.NewStorage(context.Background(), t.TempDir(), "")
	require.NoError(t, err, "failed to create storage")

	accounts := evm.NewAccountManager(t.TempDir(), big.NewInt(1337))
	genalloc := make(core.GenesisAlloc)

	// create some test accounts
	var addresses []common.Address
	for i := 0; i < 5; i++ {
		account, err := accounts.CreateAccount("testing")
		require.NoError(t, err, "failed to create account")

		address := common.HexToAddress(account)
		addresses = append(addresses, address)
		
		// allocate some funds to each account
		genalloc[address] = core.GenesisAccount{
			Balance: big.NewInt(9223372036854775807),
		}
	}

	// set the current account and unlock it
	err = accounts.SetAccount(addresses[0].Hex(), "testing")
	require.NoError(t, err, "failed to set account")

	// create a simulated backend with our accounts
	backend := backends.NewSimulatedBackend(genalloc, uint64(8000000))
	txopts := &bind.TransactOpts{
		From:   addresses[0],
		Signer: accounts.SignTx,
	}

	// deploy the contract
	address, _, _, err := valist.DeployValist(txopts, backend, common.HexToAddress("0x0"))
	require.NoError(t, err, "failed to deploy valist contract")

	// ensure transactions are executed
	backend.Commit()

	contract, err := evm.NewContract(address, backend, accounts)
	require.NoError(t, err, "failed to create contract")

	return &Client{
		AccountAPI:  accounts,
		ContractAPI: contract,
		StorageAPI:  storage,
	}
}
