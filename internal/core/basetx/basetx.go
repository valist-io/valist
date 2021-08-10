package basetx

import (
	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
	"github.com/valist-io/registry/internal/core"
)

type Transactor struct {
	valist   *valist.Valist
	registry *registry.ValistRegistry
}

func NewTransactor(valist *valist.Valist, registry *registry.ValistRegistry) core.TransactorAPI {
	return &Transactor{valist, registry}
}
