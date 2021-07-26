package core

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type RegistryAPI interface {
	GetOrganizationID(context.Context, string) (common.Hash, error)
	LinkOrganizationName(context.Context, common.Hash, string) (<-chan LinkOrgNameResult, error)
}

type LinkOrgNameResult struct {
	OrgID common.Hash
	Name  string
	Err   error
}
