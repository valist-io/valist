package core

import (
	"context"

	files "github.com/ipfs/go-ipfs-files"
)

func (s *CoreSuite) TestVoteRelease() {
	ctx := context.Background()

	orgMeta := &OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	repoName := "sdk"
	repoMeta := &RepositoryMeta{
		Name:        "sdk",
		Description: "Valist core sdk.",
		ProjectType: "npm",
		Homepage:    "https://valist.io",
		Repository:  "https://github.com/valist-io/valist",
	}

	releaseTag := "v0.0.1"
	releaseMeta := []byte("hello")
	releaseData := []byte("world")

	releaseMetaPath, err := s.client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(releaseMeta))
	s.Require().NoError(err, "Failed to add release meta")
	metaCID := releaseMetaPath.Cid()

	releaseDataPath, err := s.client.ipfs.Unixfs().Add(ctx, files.NewBytesFile(releaseData))
	s.Require().NoError(err, "Failed to add release data")
	releaseCID := releaseDataPath.Cid()

	txc1, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")
	s.backend.Commit()

	res1 := <-txc1
	s.Require().NoError(res1.Err, "Failed to create organization")
	orgID := res1.Log.OrgID

	txc2, err := s.client.CreateRepository(ctx, orgID, repoName, repoMeta)
	s.Require().NoError(err, "Failed to create repository")
	s.backend.Commit()

	res2 := <-txc2
	s.Require().NoError(res2.Err, "Failed to create repository")

	txc3, err := s.client.VoteRelease(ctx, orgID, repoName, releaseTag, releaseCID, metaCID)
	s.Require().NoError(err, "Failed to vote release")
	s.backend.Commit()

	res3 := <-txc3
	s.Require().NoError(res3.Err, "Failed to vote release")

	release, err := s.client.GetRelease(ctx, orgID, repoName, releaseTag)
	s.Require().NoError(err, "Failed to get release")
	s.Assert().Equal(release.ReleaseCID, releaseCID)
	s.Assert().Equal(release.MetaCID, metaCID)

	latest, err := s.client.GetLatestRelease(ctx, orgID, repoName)
	s.Require().NoError(err, "Failed to get latest release")
	s.Assert().Equal(latest.Tag, releaseTag)
}
