package impl

import (
	"context"
	"math/big"

	"github.com/valist-io/registry/internal/core"
)

func (s *ClientSuite) TestCreateOrganization() {
	ctx := context.Background()

	orgMeta := &core.OrganizationMeta{
		Name:        "Valist, Inc.",
		Description: "Accelerating the transition to web3.",
	}

	txc1, err := s.client.CreateOrganization(ctx, orgMeta)
	s.Require().NoError(err, "Failed to create organization")
	s.client.Commit()

	res1 := <-txc1
	s.Require().NoError(res1.Err, "Failed to create organization")
	orgID := res1.OrgID

	org, err := s.client.GetOrganization(ctx, orgID)
	s.Require().NoError(err, "Failed to get organization")
	s.Assert().Equal(big.NewInt(0).Cmp(org.Threshold), 0)

	meta, err := s.client.GetOrganizationMeta(ctx, org.MetaCID)
	s.Require().NoError(err, "Failed to get organization meta")
	s.Assert().Equal(orgMeta.Name, meta.Name)
	s.Assert().Equal(orgMeta.Description, meta.Description)
}
