package impl

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/valist-io/registry/internal/core"
)

func (s *ClientSuite) TestGetOrganizationID() {
	ctx := context.Background()

	_, err := s.client.GetOrganizationID(ctx, "empty")
	s.Assert().Equal(core.ErrOrganizationNotExist, err)

	orgName := "valist"
	orgID := common.HexToHash("0xDEADBEEF")

	txc1, err := s.client.LinkOrganizationName(ctx, orgID, orgName)
	s.Require().NoError(err, "Failed to link organization name")
	s.client.Commit()

	res1 := <-txc1
	s.Require().NoError(res1.Err, "Failed to link organization name")

	id, err := s.client.GetOrganizationID(ctx, orgName)
	s.Require().NoError(err, "Failed to get organization id")
	s.Assert().Equal(orgID[:], id.Bytes())
}
