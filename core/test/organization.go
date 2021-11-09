package test

import (
	"context"
	"math/big"

	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/types"
)

func (s *CoreSuite) TestCreateOrganization() {
	ctx := context.Background()

	orgMeta := &types.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	orgCreatedEvent, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")

	org, err := s.client.GetOrganization(ctx, orgCreatedEvent.OrgID)
	s.Require().NoError(err, "Failed to get organization")
	s.Assert().Equal(big.NewInt(0).Cmp(org.Threshold), 0)

	meta, err := s.client.GetOrganizationMeta(ctx, org.MetaCID)
	s.Require().NoError(err, "Failed to get organization meta")
	s.Assert().Equal(orgMeta.Name, meta.Name)
	s.Assert().Equal(orgMeta.Description, meta.Description)
}

func (s *CoreSuite) TestVoteOrganizationThreshold() {
	ctx := context.Background()

	orgMeta := &types.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	orgCreatedEvent, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")

	org, err := s.client.GetOrganization(ctx, orgCreatedEvent.OrgID)
	s.Require().NoError(err, "Failed to get organization")
	s.Assert().Equal(big.NewInt(0).Cmp(org.Threshold), 0)

	_, err = s.client.VoteOrganizationThreshold(ctx, orgCreatedEvent.OrgID, big.NewInt(2))
	s.Require().Error(err, "Should not be able to vote for threshold without enough members")

	_, err = s.client.VoteOrganizationAdmin(ctx, orgCreatedEvent.OrgID, client.ADD_KEY, s.accounts[1].Address)
	s.Require().NoError(err, "Failed to add second org admin key")

	_, err = s.client.VoteOrganizationAdmin(ctx, orgCreatedEvent.OrgID, client.ADD_KEY, s.accounts[2].Address)
	s.Require().NoError(err, "Failed to add third org admin key")

	_, err = s.client.VoteOrganizationThreshold(ctx, orgCreatedEvent.OrgID, big.NewInt(2))
	s.Require().NoError(err, "Failed to vote for organization threshold")

	_, err = s.client.VoteOrganizationThreshold(ctx, orgCreatedEvent.OrgID, big.NewInt(2))
	s.Require().Error(err, "Should not be able to vote for threshold again")

	s.client.Signer().SetAccount(s.accounts[1])

	_, err = s.client.VoteOrganizationThreshold(ctx, orgCreatedEvent.OrgID, big.NewInt(2))
	s.Require().NoError(err, "Failed to vote for organization threshold")

	org, err = s.client.GetOrganization(ctx, orgCreatedEvent.OrgID)
	s.Require().NoError(err, "Failed to get organization")
	s.Require().Equal(big.NewInt(2).Cmp(org.Threshold), 0)
}
