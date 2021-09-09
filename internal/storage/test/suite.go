package test

import (
	"context"

	"github.com/stretchr/testify/suite"

	"github.com/valist-io/valist/internal/storage"
)

type StorageSuite struct {
	suite.Suite
	storage storage.Storage
}

func (s *StorageSuite) SetStorage(storage storage.Storage) {
	s.storage = storage
}

func (s *StorageSuite) TestReadWrite() {
	ctx := context.Background()
	data := []byte("hello")

	p, err := s.storage.Write(ctx, data)
	s.Require().NoError(err, "Failed to add file")

	expect, err := s.storage.ReadFile(ctx, p)
	s.Require().NoError(err, "Failed to get file")
	s.Assert().Equal(data, expect)
}
