package mock

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/test"
)

type ClientSuite struct {
	test.CoreSuite
	tmp    string
	client *client.Client
}

func (s *ClientSuite) SetupTest() {
	tmp, err := os.MkdirTemp("", "test")
	s.Require().NoError(err, "Failed to create temp dir")

	kstore, err := NewKeyStore(tmp, 5)
	s.Require().NoError(err, "Failed to create keystore")

	client, err := NewClient(kstore)
	s.Require().NoError(err, "Failed to create mock client")

	s.tmp = tmp
	s.client = client

	s.CoreSuite.SetClient(client)
	s.CoreSuite.SetAccounts(kstore.Accounts())
}

func (s *ClientSuite) TearDownTest() {
	os.RemoveAll(s.tmp)
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}
