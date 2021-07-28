package impl

import (
	"context"

	"github.com/valist-io/registry/internal/core"
)

func (s *ClientSuite) TestGetRelease() {
	ctx := context.Background()

	_, err := s.client.GetRelease(ctx, emptyHash, "empty", "empty")
	s.Assert().Equal(core.ErrReleaseNotExist, err)
}

func (s *ClientSuite) TestVoteRelease() {
	ctx := context.Background()

	orgMeta := &core.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	repoName := "sdk"
	repoMeta := &core.RepositoryMeta{
		Name:        "sdk",
		Description: "Valist core sdk.",
		ProjectType: "npm",
		Homepage:    "https://valist.io",
		Repository:  "https://github.com/valist-io/valist",
	}

	metaCID, err := s.client.AddFile(ctx, []byte("hello"))
	s.Require().NoError(err, "Failed to add meta file")

	releaseCID, err := s.client.AddFile(ctx, []byte("world"))
	s.Require().NoError(err, "Failed to add release file")

	release := &core.Release{
		Tag:        "v0.0.1",
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
	}

	txc1, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")
	s.client.Commit()

	res1 := <-txc1
	s.Require().NoError(res1.Err, "Failed to create organization")
	orgID := res1.OrgID

	txc2, err := s.client.CreateRepository(ctx, orgID, repoName, repoMeta)
	s.Require().NoError(err, "Failed to create repository")
	s.client.Commit()

	res2 := <-txc2
	s.Require().NoError(res2.Err, "Failed to create repository")

	txc3, err := s.client.VoteRelease(ctx, orgID, repoName, release)
	s.Require().NoError(err, "Failed to vote release")
	s.client.Commit()

	res3 := <-txc3
	s.Require().NoError(res3.Err, "Failed to vote release")

	other, err := s.client.GetRelease(ctx, orgID, repoName, release.Tag)
	s.Require().NoError(err, "Failed to get release")
	s.Assert().Equal(release.ReleaseCID, other.ReleaseCID)
	s.Assert().Equal(release.MetaCID, other.MetaCID)

	latest, err := s.client.GetLatestRelease(ctx, orgID, repoName)
	s.Require().NoError(err, "Failed to get latest release")
	s.Assert().Equal(release.Tag, latest.Tag)
}
