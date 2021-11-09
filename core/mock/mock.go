package mock

import (
	"context"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core"

	"github.com/valist-io/valist/contract"
	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/client/basetx"
	"github.com/valist-io/valist/signer"
	"github.com/valist-io/valist/storage/ipfs"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
	passphrase       = "secret"
	numAccounts      = 5
)

var (
	chainID  = big.NewInt(1337)
	gasLimit = uint64(8000000)
	balance  = big.NewInt(9223372036854775807)
)

func NewClient(ctx context.Context) (*client.Client, error) {
	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}

	keystoreDir := filepath.Join(tmp, "keystore")
	storageDir := filepath.Join(tmp, "storage")

	kstore := keystore.NewKeyStore(keystoreDir, veryLightScryptN, veryLightScryptP)
	galloc := make(core.GenesisAlloc)

	for i := 0; i < numAccounts; i++ {
		account, err := kstore.NewAccount(passphrase)
		if err != nil {
			return nil, err
		}

		if err := kstore.Unlock(account, passphrase); err != nil {
			return nil, err
		}

		galloc[account.Address] = core.GenesisAccount{
			Balance: balance,
		}
	}

	signer, err := signer.NewSigner(chainID, kstore)
	if err != nil {
		return nil, err
	}

	// always default to first account
	account := kstore.Accounts()[0]
	signer.SetAccount(account)

	txopts := signer.NewTransactor()
	backend := backends.NewSimulatedBackend(galloc, gasLimit)

	forwarderAddress, _, _, err := contract.DeployForwarder(&txopts.TransactOpts, backend, account.Address)
	if err != nil {
		return nil, err
	}

	valistAddress, _, valist, err := contract.DeployValist(&txopts.TransactOpts, backend, forwarderAddress)
	if err != nil {
		return nil, err
	}

	registryAddress, _, registry, err := contract.DeployRegistry(&txopts.TransactOpts, backend, forwarderAddress)
	if err != nil {
		return nil, err
	}

	// ensure contracts are deployed
	backend.Commit()

	ipfs, err := ipfs.NewProvider(ctx, storageDir)
	if err != nil {
		return nil, err
	}

	transactor, err := basetx.NewTransactor(backend, valistAddress, registryAddress)
	if err != nil {
		return nil, err
	}

	return client.NewClient(client.Options{
		Storage:    ipfs,
		Ethereum:   backend,
		Valist:     valist,
		Registry:   registry,
		Signer:     signer,
		Transactor: transactor,
	})
}
