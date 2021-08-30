package ipfs

import (
	"testing"

	"github.com/ipfs/go-ipfs/core/coreapi"
	coremock "github.com/ipfs/go-ipfs/core/mock"
	"github.com/stretchr/testify/suite"

	"github.com/valist-io/registry/internal/storage/test"
)

type StorageSuite struct {
	test.StorageSuite
}

func (s *StorageSuite) SetupTest() {
	node, err := coremock.NewMockNode()
	s.Require().NoError(err, "Failed to create ipfs mock node")

	ipfsapi, err := coreapi.NewCoreAPI(node)
	s.Require().NoError(err, "Failed to create ipfs core api")

	storage := NewStorage(ipfsapi)
	s.StorageSuite.SetStorage(storage)
}

func TestStorageSuite(t *testing.T) {
	suite.Run(t, &StorageSuite{})
}
