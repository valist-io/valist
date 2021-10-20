package signer

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentOverride(t *testing.T) {
	private, err := crypto.GenerateKey()
	require.NoError(t, err, "failed to generate private key")

	address := crypto.PubkeyToAddress(private.PublicKey)
	account := accounts.Account{Address: address}

	key := common.Bytes2Hex(crypto.FromECDSA(private))
	err = os.Setenv(EnvKey, key)
	require.NoError(t, err, "failed to set environment")

	signer, err := NewSigner(big.NewInt(1337))
	require.NoError(t, err, "failed to create signer")

	assert.Equal(t, address, signer.account.Address)
	assert.Len(t, signer.manager.Accounts(), 1)

	wallet, err := signer.manager.Find(account)
	require.NoError(t, err, "failed to find wallet")

	// ensure the wallet is unlocked by signing arbitrary text
	_, err = wallet.SignText(account, []byte("hello"))
	require.NoError(t, err, "failed to sign text")
}
