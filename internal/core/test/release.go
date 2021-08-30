package test

import (
	"context"

	"github.com/valist-io/registry/internal/core/types"
)

func (s *CoreSuite) TestGetRelease() {
	ctx := context.Background()

	_, err := s.client.GetRelease(ctx, emptyHash, "empty", "empty")
	s.Assert().Equal(types.ErrReleaseNotExist, err)
}

func (s *CoreSuite) TestVoteRelease() {
	ctx := context.Background()

	orgMeta := &types.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	repoName := "sdk"
	repoMeta := &types.RepositoryMeta{
		Name:        "sdk",
		Description: "Valist core sdk.",
		ProjectType: "npm",
		Homepage:    "https://valist.io",
		Repository:  "https://github.com/valist-io/valist",
	}

	metaCID, err := s.client.Storage().Write(ctx, []byte("hello"))
	s.Require().NoError(err, "Failed to add meta file")

	releaseCID, err := s.client.Storage().Write(ctx, []byte("world"))
	s.Require().NoError(err, "Failed to add release file")

	release := &types.Release{
		Tag:        "v0.0.1",
		ReleaseCID: releaseCID,
		MetaCID:    metaCID,
	}

	orgCreatedEvent, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")

	_, err = s.client.CreateRepository(ctx, orgCreatedEvent.OrgID, repoName, repoMeta)
	s.Require().NoError(err, "Failed to create repository")

	_, err = s.client.VoteRelease(ctx, orgCreatedEvent.OrgID, repoName, release)
	s.Require().NoError(err, "Failed to vote release")

	released, err := s.client.GetRelease(ctx, orgCreatedEvent.OrgID, repoName, release.Tag)
	s.Require().NoError(err, "Failed to get release")
	s.Assert().Equal(release.ReleaseCID, released.ReleaseCID)
	s.Assert().Equal(release.MetaCID, released.MetaCID)

	latest, err := s.client.GetLatestRelease(ctx, orgCreatedEvent.OrgID, repoName)
	s.Require().NoError(err, "Failed to get latest release")
	s.Assert().Equal(release.Tag, latest.Tag)
}
