package test

import (
	"context"
)

func (s *CoreSuite) TestStorage() {
	ctx := context.Background()
	data := []byte("hello")

	dataCID, err := s.client.WriteFile(ctx, data)
	s.Require().NoError(err, "Failed to add file")

	expect, err := s.client.ReadFile(ctx, dataCID)
	s.Require().NoError(err, "Failed to get file")
	s.Assert().Equal(data, expect)

	releaseCID1, err := s.client.WriteFilePath(ctx, "testdata/example.txt")
	s.Require().NoError(err)
	s.Assert().True(releaseCID1.Defined())

	releaseCID2, err := s.client.WriteDirEntries(ctx, "base", []string{"testdata/example.txt"})
	s.Require().NoError(err)
	s.Assert().True(releaseCID2.Defined())
}
