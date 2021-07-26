package impl

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	client *MockClient
}

func (s *ClientSuite) SetupTest() {
	client, err := NewMockClient()
	s.Require().NoError(err, "Failed to create mock client")
	s.client = client
}

func (s *ClientSuite) TearDownTest() {
	s.client.Close()
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}
