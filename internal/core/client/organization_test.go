package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/valist-io/registry/internal/core"
)

func (s *ClientSuite) TestCreateOrganization() {
	ctx := context.Background()

	orgMeta := &core.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	txopts := s.client.TransactOpts()
	orgCreatedEvent, err := s.client.CreateOrganization(ctx, txopts, orgMeta)
	s.Require().NoError(err, "Failed to create organization")

	org, err := s.client.GetOrganization(ctx, orgCreatedEvent.OrgID)
	s.Require().NoError(err, "Failed to get organization")
	s.Assert().Equal(big.NewInt(0).Cmp(org.Threshold), 0)

	meta, err := s.client.GetOrganizationMeta(ctx, org.MetaCID)
	s.Require().NoError(err, "Failed to get organization meta")
	s.Assert().Equal(orgMeta.Name, meta.Name)
	s.Assert().Equal(orgMeta.Description, meta.Description)
}

func (s *ClientSuite) TestVoteOrganizationThreshold() {
	ctx := context.Background()

	orgMeta := &core.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	txopts := s.client.TransactOpts()
	orgCreatedEvent, err := s.client.CreateOrganization(ctx, txopts, orgMeta)
	s.Require().NoError(err, "Failed to create organization")

	org, err := s.client.GetOrganization(ctx, orgCreatedEvent.OrgID)
	s.Require().NoError(err, "Failed to get organization")
	s.Assert().Equal(big.NewInt(0).Cmp(org.Threshold), 0)

	voteThresholdEvent, err := s.client.VoteOrganizationThreshold(ctx, txopts, orgCreatedEvent.OrgID, big.NewInt(2))
	s.Require().NoError(err, "Failed to set organization threshold")
	fmt.Println(voteThresholdEvent)
}
