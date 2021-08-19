package client

import (
	"context"
)

func (s *ClientSuite) TestStorage() {
	ctx := context.Background()
	data := []byte("hello")

	dataCID, err := s.client.WriteFile(ctx, data)
	s.Require().NoError(err, "Failed to add file")

	expect, err := s.client.ReadFile(ctx, dataCID)
	s.Require().NoError(err, "Failed to get file")
	s.Assert().Equal(data, expect)
}
