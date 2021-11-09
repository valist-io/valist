package mock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/valist-io/valist/core/test"
)

type ClientSuite struct {
	test.CoreSuite
}

func (s *ClientSuite) SetupTest() {
	ctx := context.Background()

	client, err := NewClient(ctx)
	s.CoreSuite.Require().NoError(err, "Failed to create mock client")

	s.CoreSuite.SetClient(client)
	s.CoreSuite.SetAccounts(client.Signer().List())
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}
