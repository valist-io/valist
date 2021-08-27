package mock

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/stretchr/testify/suite"

	"github.com/valist-io/registry/internal/core/client"
	"github.com/valist-io/registry/internal/core/test"
)

type ClientSuite struct {
	test.CoreSuite
	tmp    string
	client *client.Client
}

func (s *ClientSuite) SetupTest() {
	tmp, err := os.MkdirTemp("", "test")
	s.Require().NoError(err, "Failed to create temp dir")

	signer := keystore.NewKeyStore(tmp, veryLightScryptN, veryLightScryptP)

	test_accounts := []accounts.Account{}

	for i := 0; i < 5; i++ {
		new_account, err := signer.NewAccount(passphrase)
		s.Require().NoError(err, "Failed to create keystore account")

		err = signer.Unlock(new_account, passphrase)
		s.Require().NoError(err, "Failed to unlock keystore account")

		test_accounts = append(test_accounts, new_account)
	}

	client, err := NewClient(signer, test_accounts)
	s.Require().NoError(err, "Failed to create mock client")

	s.client = client
	s.CoreSuite.SetClient(client)
	s.CoreSuite.SetAccounts(signer, test_accounts)
}

func (s *ClientSuite) TearDownTest() {
	os.RemoveAll(s.tmp)
	s.client.Close()
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}
