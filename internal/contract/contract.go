//go:generate abigen --sol ../../contracts/Valist.sol --pkg valist --out ./valist/valist.go
//go:generate abigen --sol ../../contracts/ValistRegistry.sol --pkg registry --out ./valist/registry/registry.go
package contract
