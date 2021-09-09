//go:generate abigen --sol ../../contracts/Valist.sol --pkg valist --out ./valist/valist.go
//go:generate abigen --sol ../../contracts/ValistRegistry.sol --pkg registry --out ./registry/registry.go
//go:generate abigen --abi ../../contracts/BiconomyForwarderABI.json --bin ../../contracts/BiconomyForwarder.bin --pkg metatx --out ./metatx/forwarder.go
package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/valist-io/valist/internal/contract/metatx"
	"github.com/valist-io/valist/internal/contract/registry"
	"github.com/valist-io/valist/internal/contract/valist"
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

func NewForwarder(address common.Address, backend bind.ContractBackend) (*metatx.Metatx, error) {
	return metatx.NewMetatx(address, backend)
}

func DeployForwarder(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *types.Transaction, *metatx.Metatx, error) {
	return metatx.DeployMetatx(auth, backend, owner)
}

func NewValistABI() (abi.ABI, error) {
	return abi.JSON(strings.NewReader(valist.ValistABI))
}
