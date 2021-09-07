package mock

import (
	"os"
	"testing"

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
	s.tmp = tmp

	client, accounts, wallets, err := NewClient(tmp)
	s.Require().NoError(err, "Failed to create mock client")

	s.client = client
	s.CoreSuite.SetClient(client)
	s.CoreSuite.SetAccounts(accounts)
	s.CoreSuite.SetWallets(wallets)
}

func (s *ClientSuite) TearDownTest() {
	os.RemoveAll(s.tmp)
	s.client.Close()
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}
