package mock

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	coreeth "github.com/ethereum/go-ethereum/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	coremock "github.com/ipfs/go-ipfs/core/mock"

	"github.com/valist-io/registry/internal/contract"
	"github.com/valist-io/registry/internal/core/client"
	"github.com/valist-io/registry/internal/core/client/basetx"
)

var chainID = big.NewInt(1337)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
	passphrase       = "secret"
)

func NewClient(ksLocation string) (*client.Client, error) {
	var onClose []client.Close

	signer := keystore.NewKeyStore(ksLocation, veryLightScryptN, veryLightScryptP)
	account, err := signer.NewAccount(passphrase)
	if err != nil {
		return nil, err
	}

	if err = signer.Unlock(account, passphrase); err != nil {
		return nil, err
	}

	backend := backends.NewSimulatedBackend(coreeth.GenesisAlloc{
		account.Address: {Balance: big.NewInt(9223372036854775807)},
	}, 8000029)
	onClose = append(onClose, backend.Close)

	opts, err := bind.NewKeyStoreTransactorWithChainID(signer, account, chainID)
	if err != nil {
		return nil, err
	}

	forwarderAddress, _, _, err := contract.DeployForwarder(opts, backend, account.Address)
	if err != nil {
		return nil, err
	}

	_, _, valist, err := contract.DeployValist(opts, backend, forwarderAddress)
	if err != nil {
		return nil, err
	}

	_, _, registry, err := contract.DeployRegistry(opts, backend, forwarderAddress)
	if err != nil {
		return nil, err
	}

	// ensure contracts are deployed
	backend.Commit()

	node, err := coremock.NewMockNode()
	if err != nil {
		return nil, err
	}

	ipfs, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return nil, err
	}

	return client.NewClient(&client.Options{
		IPFS:         ipfs,
		Ethereum:     backend,
		ChainID:      chainID,
		Valist:       valist,
		Registry:     registry,
		Account:      account,
		Wallet:       signer.Wallets()[0],
		TransactOpts: basetx.TransactOpts,
		Transactor:   basetx.NewTransactor(valist, registry),
		OnClose:      onClose,
	})
}
