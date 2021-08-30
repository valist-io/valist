package test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	"github.com/valist-io/registry/internal/core/types"
)

var emptyHash = common.HexToHash("0x0")

type CoreSuite struct {
	suite.Suite
	client types.CoreAPI
}

func (s *CoreSuite) SetClient(client types.CoreAPI) {
	s.client = client
}
