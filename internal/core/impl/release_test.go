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

	txopts := s.client.NewTransactOpts()

	orgCreatedEvent, err := s.client.CreateOrganization(ctx, txopts, orgMeta)
	s.Require().NoError(err, "Failed to create organization")

	_, err = s.client.CreateRepository(ctx, txopts, orgCreatedEvent.OrgID, repoName, repoMeta)
	s.Require().NoError(err, "Failed to create repository")

	_, err = s.client.VoteRelease(ctx, txopts, orgCreatedEvent.OrgID, repoName, release)
	s.Require().NoError(err, "Failed to vote release")

	released, err := s.client.GetRelease(ctx, orgCreatedEvent.OrgID, repoName, release.Tag)
	s.Require().NoError(err, "Failed to get release")
	s.Assert().Equal(release.ReleaseCID, released.ReleaseCID)
	s.Assert().Equal(release.MetaCID, released.MetaCID)

	latest, err := s.client.GetLatestRelease(ctx, orgCreatedEvent.OrgID, repoName)
	s.Require().NoError(err, "Failed to get latest release")
	s.Assert().Equal(release.Tag, latest.Tag)
}
