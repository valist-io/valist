package client

import (
	"context"
	"net/http"

	httpapi "github.com/ipfs/go-ipfs-http-client"
)

func (s *ClientSuite) TestMockStorage() {
	ctx := context.Background()
	data := []byte("hello")
	artifactPaths := [1]string{"testdata/example.txt"}

	dataCID, err := s.client.WriteFile(ctx, data)
	s.Require().NoError(err, "Failed to add file")

	expect, err := s.client.ReadFile(ctx, dataCID)
	s.Require().NoError(err, "Failed to get file")
	s.Assert().Equal(data, expect)

	releaseCID1, err := s.client.WriteFilePath(ctx, "testdata/example.txt")
	s.Require().NoError(err)
	s.Assert().NotNil(releaseCID1)

	releaseCID2, err := s.client.WriteDirEntries(ctx, "base", artifactPaths[:])
	s.Require().NoError(err)
	s.Assert().NotNil(releaseCID2)
}

func (s *ClientSuite) TestRemoteStorage() {
	ctx := context.Background()
	data := []byte("hello")
	artifactPaths := [1]string{"testdata/example.txt"}
	ipfs, err := httpapi.NewURLApiWithClient("https://pin.valist.io", &http.Client{})
	s.Require().NoError(err, "Failed to connect to https://pin.valist.io")
	s.client.ipfs = ipfs

	dataCID, err := s.client.WriteFile(ctx, data)
	s.Require().NoError(err, "Failed to add file")
	s.Assert().NotNil(dataCID)

	expect, err := s.client.ReadFile(ctx, dataCID)
	s.Require().NoError(err, "Failed to get file")
	s.Assert().Equal(data, expect)

	releaseCID1, err := s.client.WriteFilePath(ctx, "testdata/example.txt")
	s.Require().NoError(err)
	s.Assert().NotNil(releaseCID1)

	releaseCID2, err := s.client.WriteDirEntries(ctx, "base", artifactPaths[:])
	s.Require().NoError(err)
	s.Assert().NotNil(releaseCID2)
}
