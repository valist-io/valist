package mock

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	coremock "github.com/ipfs/go-ipfs/core/mock"

	"github.com/valist-io/registry/internal/contract"
	"github.com/valist-io/registry/internal/core/client"
	"github.com/valist-io/registry/internal/core/client/basetx"
	"github.com/valist-io/registry/internal/storage/ipfs"
)

var chainID = big.NewInt(1337)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
	passphrase       = "secret"
	testAccounts     = 5
)

func NewClient(ksLocation string) (*client.Client, []accounts.Account, []accounts.Wallet, error) {
	var onClose []client.Close

	signer := keystore.NewKeyStore(ksLocation, veryLightScryptN, veryLightScryptP)
	alloc := make(core.GenesisAlloc)

	var accounts []accounts.Account
	for i := 0; i < testAccounts; i++ {
		account, err := signer.NewAccount(passphrase)
		if err != nil {
			return nil, nil, nil, err
		}

		err = signer.Unlock(account, passphrase)
		if err != nil {
			return nil, nil, nil, err
		}

		accounts = append(accounts, account)
		alloc[account.Address] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}
	}

	backend := backends.NewSimulatedBackend(alloc, 8000029)
	onClose = append(onClose, backend.Close)

	opts, err := bind.NewKeyStoreTransactorWithChainID(signer, accounts[0], chainID)
	if err != nil {
		return nil, nil, nil, err
	}

	forwarderAddress, _, _, err := contract.DeployForwarder(opts, backend, accounts[0].Address)
	if err != nil {
		return nil, nil, nil, err
	}

	_, _, valist, err := contract.DeployValist(opts, backend, forwarderAddress)
	if err != nil {
		return nil, nil, nil, err
	}

	_, _, registry, err := contract.DeployRegistry(opts, backend, forwarderAddress)
	if err != nil {
		return nil, nil, nil, err
	}

	// ensure contracts are deployed
	backend.Commit()

	node, err := coremock.NewMockNode()
	if err != nil {
		return nil, nil, nil, err
	}

	ipfsapi, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return nil, nil, nil, err
	}

	wallets := signer.Wallets()
	client, err := client.NewClient(&client.Options{
		Storage:      ipfs.NewStorage(ipfsapi),
		Ethereum:     backend,
		ChainID:      chainID,
		Valist:       valist,
		Registry:     registry,
		Account:      accounts[0],
		Wallet:       wallets[0],
		TransactOpts: basetx.TransactOpts,
		Transactor:   basetx.NewTransactor(valist, registry),
		OnClose:      onClose,
	})

	if err != nil {
		return nil, nil, nil, err
	}

	return client, accounts, wallets, nil
}
