//go:generate abigen --sol ../../contracts/Valist.sol --pkg valist --out ./valist/valist.go
//go:generate abigen --sol ../../contracts/ValistRegistry.sol --pkg registry --out ./registry/registry.go
package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/valist-io/registry/internal/contract/registry"
	"github.com/valist-io/registry/internal/contract/valist"
)

func NewValist(address common.Address, backend bind.ContractBackend) (*valist.Valist, error) {
	return valist.NewValist(address, backend)
}

func DeployValist(auth *bind.TransactOpts, backend bind.ContractBackend, metaTxForwarder common.Address) (common.Address, *types.Transaction, *valist.Valist, error) {
	return valist.DeployValist(auth, backend, metaTxForwarder)
}

func NewRegistry(address common.Address, backend bind.ContractBackend) (*registry.ValistRegistry, error) {
	return registry.NewValistRegistry(address, backend)
}

func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, metaTxForwarder common.Address) (common.Address, *types.Transaction, *registry.ValistRegistry, error) {
	return registry.DeployValistRegistry(auth, backend, metaTxForwarder)
}
