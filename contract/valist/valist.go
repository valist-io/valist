// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package valist

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC2771ContextABI is the input ABI used to generate the binding from.
const ERC2771ContextABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC2771ContextFuncSigs maps the 4-byte function signature to its string representation.
var ERC2771ContextFuncSigs = map[string]string{
	"572b6c05": "isTrustedForwarder(address)",
}

// ERC2771Context is an auto generated Go binding around an Ethereum contract.
type ERC2771Context struct {
	ERC2771ContextCaller     // Read-only binding to the contract
	ERC2771ContextTransactor // Write-only binding to the contract
	ERC2771ContextFilterer   // Log filterer for contract events
}

// ERC2771ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC2771ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC2771ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC2771ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC2771ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC2771ContextSession struct {
	Contract     *ERC2771Context   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC2771ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC2771ContextCallerSession struct {
	Contract *ERC2771ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC2771ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC2771ContextTransactorSession struct {
	Contract     *ERC2771ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC2771ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC2771ContextRaw struct {
	Contract *ERC2771Context // Generic contract binding to access the raw methods on
}

// ERC2771ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC2771ContextCallerRaw struct {
	Contract *ERC2771ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ERC2771ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC2771ContextTransactorRaw struct {
	Contract *ERC2771ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC2771Context creates a new instance of ERC2771Context, bound to a specific deployed contract.
func NewERC2771Context(address common.Address, backend bind.ContractBackend) (*ERC2771Context, error) {
	contract, err := bindERC2771Context(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC2771Context{ERC2771ContextCaller: ERC2771ContextCaller{contract: contract}, ERC2771ContextTransactor: ERC2771ContextTransactor{contract: contract}, ERC2771ContextFilterer: ERC2771ContextFilterer{contract: contract}}, nil
}

// NewERC2771ContextCaller creates a new read-only instance of ERC2771Context, bound to a specific deployed contract.
func NewERC2771ContextCaller(address common.Address, caller bind.ContractCaller) (*ERC2771ContextCaller, error) {
	contract, err := bindERC2771Context(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC2771ContextCaller{contract: contract}, nil
}

// NewERC2771ContextTransactor creates a new write-only instance of ERC2771Context, bound to a specific deployed contract.
func NewERC2771ContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC2771ContextTransactor, error) {
	contract, err := bindERC2771Context(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC2771ContextTransactor{contract: contract}, nil
}

// NewERC2771ContextFilterer creates a new log filterer instance of ERC2771Context, bound to a specific deployed contract.
func NewERC2771ContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC2771ContextFilterer, error) {
	contract, err := bindERC2771Context(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC2771ContextFilterer{contract: contract}, nil
}

// bindERC2771Context binds a generic wrapper to an already deployed contract.
func bindERC2771Context(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC2771ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC2771Context *ERC2771ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC2771Context.Contract.ERC2771ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC2771Context *ERC2771ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Context.Contract.ERC2771ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC2771Context *ERC2771ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC2771Context.Contract.ERC2771ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC2771Context *ERC2771ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC2771Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC2771Context *ERC2771ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC2771Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC2771Context *ERC2771ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC2771Context.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ERC2771Context *ERC2771ContextCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _ERC2771Context.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ERC2771Context *ERC2771ContextSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _ERC2771Context.Contract.IsTrustedForwarder(&_ERC2771Context.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_ERC2771Context *ERC2771ContextCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _ERC2771Context.Contract.IsTrustedForwarder(&_ERC2771Context.CallOpts, forwarder)
}

// EnumerableSetABI is the input ABI used to generate the binding from.
const EnumerableSetABI = "[]"

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
var EnumerableSetBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b4518151963fbd9760c952ed2b303e68e8115cf7dbf39b36ca60fd1ebfb36c8664736f6c63430008060033"

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// ValistABI is the input ABI used to generate the binding from.
const ValistABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"metaTxForwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaCID\",\"type\":\"string\"}],\"name\":\"MetaUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_metaCIDHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaCID\",\"type\":\"string\"}],\"name\":\"OrgCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_repoNameHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_metaCIDHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaCID\",\"type\":\"string\"}],\"name\":\"RepoCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_operation\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_key\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sigCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"VoteKeyEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_releaseCID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaCID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sigCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"VoteReleaseEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_pendingThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sigCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"VoteThresholdEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_operation\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestIndex\",\"type\":\"uint256\"}],\"name\":\"clearPendingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_releaseCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metaCID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_requestIndex\",\"type\":\"uint256\"}],\"name\":\"clearPendingRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestIndex\",\"type\":\"uint256\"}],\"name\":\"clearPendingThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgMeta\",\"type\":\"string\"}],\"name\":\"createOrganization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_repoMeta\",\"type\":\"string\"}],\"name\":\"createRepository\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"}],\"name\":\"getLatestRelease\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_selector\",\"type\":\"bytes32\"}],\"name\":\"getPendingReleaseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_selector\",\"type\":\"bytes32\"}],\"name\":\"getPendingVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_selector\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_resultsPerPage\",\"type\":\"uint256\"}],\"name\":\"getReleaseTags\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_resultsPerPage\",\"type\":\"uint256\"}],\"name\":\"getRepoNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_selector\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_selector\",\"type\":\"bytes32\"}],\"name\":\"getRoleRequestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_selector\",\"type\":\"bytes32\"}],\"name\":\"getThresholdRequestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isOrgAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isRepoDev\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orgCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orgIDs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"thresholdDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metaCID\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingReleaseRequests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"releaseCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metaCID\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingRoleRequests\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingThresholdRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"releases\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"releaseCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metaCID\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"repos\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"thresholdDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metaCID\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"roleModifiedTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_metaCID\",\"type\":\"string\"}],\"name\":\"setOrgMeta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metaCID\",\"type\":\"string\"}],\"name\":\"setRepoMeta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_operation\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_key\",\"type\":\"address\"}],\"name\":\"voteKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_releaseCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metaCID\",\"type\":\"string\"}],\"name\":\"voteRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_repoName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"voteThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValistFuncSigs maps the 4-byte function signature to its string representation.
var ValistFuncSigs = map[string]string{
	"40fd48a7": "clearPendingKey(bytes32,string,bytes32,address,uint256)",
	"c6c2a0be": "clearPendingRelease(bytes32,string,string,string,string,uint256)",
	"b93d1685": "clearPendingThreshold(bytes32,string,uint256,uint256)",
	"6427acca": "createOrganization(string)",
	"e59dbfa4": "createRepository(bytes32,string,string)",
	"5cd20d29": "getLatestRelease(bytes32,string)",
	"924a5443": "getPendingReleaseCount(bytes32)",
	"d565e860": "getPendingVotes(bytes32)",
	"f3439d56": "getReleaseTags(bytes32,uint256,uint256)",
	"837dc4d0": "getRepoNames(bytes32,uint256,uint256)",
	"a3246ad3": "getRoleMembers(bytes32)",
	"3e1bf64b": "getRoleRequestCount(bytes32)",
	"3b0659fd": "getThresholdRequestCount(bytes32)",
	"d96162a1": "isOrgAdmin(bytes32,address)",
	"1fd522a6": "isRepoDev(bytes32,string,address)",
	"572b6c05": "isTrustedForwarder(address)",
	"d8106690": "orgCount()",
	"8775692f": "orgIDs(uint256)",
	"a3e84beb": "orgs(bytes32)",
	"a940abcb": "pendingReleaseRequests(bytes32,uint256)",
	"dfec1bb0": "pendingRoleRequests(bytes32,uint256)",
	"c372be80": "pendingThresholdRequests(bytes32,uint256)",
	"f491a84c": "releases(bytes32)",
	"02b2583a": "repos(bytes32)",
	"8bf67370": "roleModifiedTimestamps(bytes32)",
	"e253f5e1": "setOrgMeta(bytes32,string)",
	"dedc3391": "setRepoMeta(bytes32,string,string)",
	"486ff0cd": "versionRecipient()",
	"f26f3c56": "voteKey(bytes32,string,bytes32,address)",
	"328d3ddf": "voteRelease(bytes32,string,string,string,string)",
	"f735b352": "voteThreshold(bytes32,string,uint256)",
}

// ValistBin is the compiled bytecode used for deploying new contracts.
var ValistBin = "0x60e0604052600560a0819052640322e322e360dc1b60c09081526200002891600091906200006f565b503480156200003657600080fd5b50604051620049b1380380620049b1833981016040819052620000599162000115565b60601b6001600160601b03191660805262000184565b8280546200007d9062000147565b90600052602060002090601f016020900481019282620000a15760008555620000ec565b82601f10620000bc57805160ff1916838001178555620000ec565b82800160010185558215620000ec579182015b82811115620000ec578251825591602001919060010190620000cf565b50620000fa929150620000fe565b5090565b5b80821115620000fa5760008155600101620000ff565b6000602082840312156200012857600080fd5b81516001600160a01b03811681146200014057600080fd5b9392505050565b600181811c908216806200015c57607f821691505b602082108114156200017e57634e487b7160e01b600052602260045260246000fd5b50919050565b60805160601c614807620001aa600039600081816102d401526137c901526148076000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c8063a3e84beb1161010f578063dedc3391116100a2578063f26f3c5611610071578063f26f3c56146104eb578063f3439d56146104fe578063f491a84c14610511578063f735b3521461053257600080fd5b8063dedc339114610487578063dfec1bb01461049a578063e253f5e1146104c5578063e59dbfa4146104d857600080fd5b8063c6c2a0be116100de578063c6c2a0be14610437578063d565e8601461044a578063d81066901461046b578063d96162a11461047457600080fd5b8063a3e84beb146103cd578063a940abcb146103ef578063b93d168514610411578063c372be801461042457600080fd5b8063572b6c05116101875780638775692f116101565780638775692f1461035a5780638bf673701461036d578063924a54431461038d578063a3246ad3146103ad57600080fd5b8063572b6c05146102c45780635cd20d29146103045780636427acca14610327578063837dc4d01461033a57600080fd5b80633b0659fd116101c35780633b0659fd1461024e5780633e1bf64b1461027c57806340fd48a71461029c578063486ff0cd146102af57600080fd5b806302b2583a146101ea5780631fd522a614610216578063328d3ddf14610239575b600080fd5b6101fd6101f8366004613c20565b610545565b60405161020d949392919061431d565b60405180910390f35b610229610224366004613cab565b6105fc565b604051901515815260200161020d565b61024c610247366004613e35565b610703565b005b61026e61025c366004613c20565b6000908152600a602052604090205490565b60405190815260200161020d565b61026e61028a366004613c20565b60009081526009602052604090205490565b61024c6102aa366004613d61565b610e36565b6102b76111e7565b60405161020d919061446a565b6102296102d2366004613c05565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b610317610312366004613c65565b611275565b60405161020d9493929190614530565b61024c61033536600461409c565b61157d565b61034d610348366004614070565b6117af565b60405161020d91906142bb565b61026e610368366004613c20565b611950565b61026e61037b366004613c20565b60076020526000908152604090205481565b61026e61039b366004613c20565b60009081526008602052604090205490565b6103c06103bb366004613c20565b611971565b60405161020d91906142a8565b6103e06103db366004613c20565b611a49565b60405161020d9392919061460f565b6104026103fd36600461404e565b611af4565b60405161020d939291906144f7565b61024c61041f366004613ff8565b611c50565b61026e61043236600461404e565b611eb4565b61024c610445366004613eeb565b611ee5565b61045d610458366004613c20565b6121e0565b60405161020d9291906145ee565b61026e60015481565b610229610482366004613c39565b6122d2565b61024c610495366004613dc9565b61233a565b6104ad6104a836600461404e565b61241b565b6040516001600160a01b03909116815260200161020d565b61024c6104d3366004613c65565b612453565b61024c6104e6366004613dc9565b612518565b61024c6104f9366004613d01565b61270e565b61034d61050c366004614070565b612f7f565b61052461051f366004613c20565b613116565b60405161020d92919061447d565b61024c610540366004613fa9565b613242565b600460205260009081526040902080546001820154600283015460038401805460ff909416949293919291610579906146c3565b80601f01602080910402602001604051908101604052809291908181526020018280546105a5906146c3565b80156105f25780601f106105c7576101008083540402835291602001916105f2565b820191906000526020600020905b8154815290600101906020018083116105d557829003601f168201915b5050505050905084565b8151600090156106ef5760046000858560405160200161061d929190614140565b60408051601f198184030181529181528151602092830120835290820192909252016000205460ff166106815760405162461bcd60e51b81526020600482015260076024820152664e6f207265706f60c81b60448201526064015b60405180910390fd5b600084846000805160206147928339815191526040516020016106a6939291906141a5565b60408051601f1981840301815291815281516020928301206000818152600690935291209091506106d790846137a3565b806106e757506106e785846122d2565b9150506106fc565b6106f984836122d2565b90505b9392505050565b848461071282826102246137c5565b61072e5760405162461bcd60e51b8152600401610678906145ad565b60008551116107685760405162461bcd60e51b81526020600482015260066024820152654e6f2074616760d01b6044820152606401610678565b60008451116107a95760405162461bcd60e51b815260206004820152600d60248201526c139bc81c995b19585cd950d251609a1b6044820152606401610678565b60008351116107e75760405162461bcd60e51b815260206004820152600a602482015269139bc81b595d1850d25160b21b6044820152606401610678565b60008787876040516020016107fe93929190614344565b60408051601f198184030181529181528151602092830120600081815260059093529120805491925090610831906146c3565b15905061086b5760405162461bcd60e51b8152602060048201526008602482015267151859c81d5cd95960c21b6044820152606401610678565b6000888888888860405160200161088695949392919061436f565b604051602081830303815290604052805190602001209050600089896040516020016108b3929190614140565b6040516020818303038152906040528051906020012090506001600460008381526020019081526020016000206001015411610a365760008381526005602090815260409091208851610908928a0190613a0c565b506000838152600560209081526040909120875161092e92600190920191890190613a0c565b5060008381526005602052604090206002016109486137c5565b8154600180820184556000938452602080852090920180546001600160a01b0319166001600160a01b03949094169390931790925583835260048082526040842001805492830181558352918290208a516109ab939190920191908b0190613a0c565b50876040516109ba919061421a565b6040518091039020896040516109d0919061421a565b60405180910390208b7f38b17387282322f8d6de03e5d4b3ff512d3bb4d3db8cc2be4611ecfb8a126c6b8a8a610a046137c5565b600088815260046020526040908190206001908101549151610a2995949392906144ab565b60405180910390a4610e2a565b6000828152600b6020526040902054610b1857610a564262093a8061462e565b6000838152600b6020526040902055610a87610a706137c5565b6000848152600b602052604090206001019061380a565b50600081815260086020908152604080832081516060810183528c81528084018c90529182018a905280546001810182559084529282902081518051929460030290910192610adb92849290910190613a0c565b506020828101518051610af49260018501920190613a0c565b5060408201518051610b10916002840191602090910190613a0c565b505050610d8b565b6000828152600b6020526040902054421115610b465760405162461bcd60e51b8152600401610678906145cd565b610b68610b516137c5565b6000848152600b60205260409020600101906137a3565b15610ba25760405162461bcd60e51b815260206004820152600a602482015269155cd95c881d9bdd195960b21b6044820152606401610678565b610bad610a706137c5565b5060005b6000838152600b60205260409020610bcb9060010161381f565b811015610c45576000838152600b60205260409020610bf5908c908c906102249060010185613829565b610c35576000838152600b60205260409020610c2f90610c189060010183613829565b6000858152600b6020526040902060010190613835565b50600090505b610c3e816146fe565b9050610bb1565b506000818152600460209081526040808320600190810154868552600b90935292209091610c73910161381f565b10610d8b5760008381526005602090815260409091208851610c97928a0190613a0c565b5060008381526005602090815260409091208751610cbd92600190920191890190613a0c565b5060005b6000838152600b60205260409020610cdb9060010161381f565b811015610d53576000848152600560209081526040808320868452600b909252909120600290910190610d119060010183613829565b81546001810183556000928352602090922090910180546001600160a01b0319166001600160a01b03909216919091179055610d4c816146fe565b9050610cc1565b50600081815260046020818152604083209091018054600181018255908352918190208a51610d899391909101918b0190613a0c565b505b87604051610d99919061421a565b604051809103902089604051610daf919061421a565b60405180910390208b7f38b17387282322f8d6de03e5d4b3ff512d3bb4d3db8cc2be4611ecfb8a126c6b8a8a610de36137c5565b6000898152600b60205260409020610dfd9060010161381f565b60008981526004602052604090819020600101549051610e219594939291906144ab565b60405180910390a45b50505050505050505050565b8484610e4582826102246137c5565b610e615760405162461bcd60e51b8152600401610678906145ad565b600080516020614772833981519152851480610e9c57507fb85a6b22a26ce426daea6357c60a74cb0b4d36234cd7c96170cd6a64102786ff85145b610ed55760405162461bcd60e51b815260206004820152600a6024820152690496e76616c6964206f760b41b6044820152606401610678565b6000806000898988604051602001610eef93929190614166565b604051602081830303815290604052805190602001209050600089511115610f805789896000805160206147928339815191528a8a604051602001610f389594939291906141d2565b6040516020818303038152906040528051906020012092508989604051602001610f63929190614140565b604051602081830303815290604052805190602001209150610fda565b60408051602081018c90526000805160206147b28339815191529181019190915260608082018a905288901b6001600160601b03191660808201526094016040516020818303038152906040528051906020012092508991505b6000838152600b60205260409020544210158061102157506000838152600b602052604090205461100f9062093a8090614665565b60008281526007602052604090205410155b61103d5760405162461bcd60e51b815260040161067890614588565b600082815260096020526040902080546001600160a01b03891691908890811061106957611069614745565b6000918252602090912001546001600160a01b0316146110b75760405162461bcd60e51b815260206004820152600960248201526857726f6e67206b657960b81b6044820152606401610678565b6000838152600b602052604081206110d19060010161381f565b11156110fc576000838152600b602052604081206110f691610c189160010190613829565b506110b7565b6000838152600b6020908152604080832083905584835260099091529020805461112890600190614665565b8154811061113857611138614745565b60009182526020808320909101548483526009909152604090912080546001600160a01b03909216918890811061117157611171614745565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394909416939093179092558381526009909152604090208054806111b9576111b961472f565b600082815260209020810160001990810180546001600160a01b031916905501905550505050505050505050565b600080546111f4906146c3565b80601f0160208091040260200160405190810160405280929190818152602001828054611220906146c3565b801561126d5780601f106112425761010080835404028352916020019161126d565b820191906000526020600020905b81548152906001019060200180831161125057829003601f168201915b505050505081565b6060806060806000600460008888604051602001611294929190614140565b6040516020818303038152906040528051906020012081526020019081526020016000209050600060056000898985600401600187600401805490506112da9190614665565b815481106112ea576112ea614745565b90600052602060002001604051602001611306939291906143ce565b604051602081830303815290604052805190602001208152602001908152602001600020905081600401600183600401805490506113449190614665565b8154811061135457611354614745565b90600052602060002001816000018260010183600201838054611376906146c3565b80601f01602080910402602001604051908101604052809291908181526020018280546113a2906146c3565b80156113ef5780601f106113c4576101008083540402835291602001916113ef565b820191906000526020600020905b8154815290600101906020018083116113d257829003601f168201915b50505050509350828054611402906146c3565b80601f016020809104026020016040519081016040528092919081815260200182805461142e906146c3565b801561147b5780601f106114505761010080835404028352916020019161147b565b820191906000526020600020905b81548152906001019060200180831161145e57829003601f168201915b5050505050925081805461148e906146c3565b80601f01602080910402602001604051908101604052809291908181526020018280546114ba906146c3565b80156115075780601f106114dc57610100808354040283529160200191611507565b820191906000526020600020905b8154815290600101906020018083116114ea57829003601f168201915b505050505091508080548060200260200160405190810160405280929190818152602001828054801561156357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611545575b505050505090509550955095509550505092959194509250565b60008151116115bb5760405162461bcd60e51b815260206004820152600a6024820152694e6f206f72674d65746160b01b6044820152606401610678565b60006001600081546115cc906146fe565b9182905550604080516020810192909252469082015260600160408051601f1981840301815291815281516020928301206000818152600384529190912084519193506116229260029091019190850190613a0c565b50600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018190556116c36116636137c5565b60066000846000805160206147b2833981519152604051602001611691929190918252602082015260400190565b60405160208183030381529060405280519060200120815260200190815260200160002061380a90919063ffffffff16565b50816040516116d2919061421a565b6040518091039020817fc5eb86c0b2c1ce6abdc8dea996a5aa6cf196b33ee7a2c140ce4f04f2fbb3baab8460405161170a919061446a565b60405180910390a361171a6137c5565b6001600160a01b03167fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470827fa8cea71b77054741d93ac504e0ee90fa3e815d68104468d65f9eb36924a8d59061176e6137c5565b604080516001600160a01b03909216825260008051602061477283398151915260208301526000908201819052606082015260800160405180910390a45050565b60606000826117be8582614646565b6117c89190614665565b905060006117d68486614646565b600087815260036020819052604090912001549091508111156118085750600085815260036020819052604090912001545b6000846001600160401b038111156118225761182261475b565b60405190808252806020026020018201604052801561185557816020015b60608152602001906001900390816118405790505b5090505b818310156119465760008781526003602081905260409091200180548490811061188557611885614745565b90600052602060002001805461189a906146c3565b80601f01602080910402602001604051908101604052809291908181526020018280546118c6906146c3565b80156119135780601f106118e857610100808354040283529160200191611913565b820191906000526020600020905b8154815290600101906020018083116118f657829003601f168201915b505050505081848151811061192a5761192a614745565b60200260200101819052508261193f906146fe565b9250611859565b9695505050505050565b6002818154811061196057600080fd5b600091825260209091200154905081565b60008181526006602052604081206060919061198c9061381f565b6001600160401b038111156119a3576119a361475b565b6040519080825280602002602001820160405280156119cc578160200160208202803683370190505b50905060005b60008481526006602052604090206119e99061381f565b811015611a42576000848152600660205260409020611a089082613829565b828281518110611a1a57611a1a614745565b6001600160a01b0390921660209283029190910190910152611a3b816146fe565b90506119d2565b5092915050565b60036020526000908152604090208054600182015460028301805492939192611a71906146c3565b80601f0160208091040260200160405190810160405280929190818152602001828054611a9d906146c3565b8015611aea5780601f10611abf57610100808354040283529160200191611aea565b820191906000526020600020905b815481529060010190602001808311611acd57829003601f168201915b5050505050905083565b60086020528160005260406000208181548110611b1057600080fd5b906000526020600020906003020160009150915050806000018054611b34906146c3565b80601f0160208091040260200160405190810160405280929190818152602001828054611b60906146c3565b8015611bad5780601f10611b8257610100808354040283529160200191611bad565b820191906000526020600020905b815481529060010190602001808311611b9057829003601f168201915b505050505090806001018054611bc2906146c3565b80601f0160208091040260200160405190810160405280929190818152602001828054611bee906146c3565b8015611c3b5780601f10611c1057610100808354040283529160200191611c3b565b820191906000526020600020905b815481529060010190602001808311611c1e57829003601f168201915b505050505090806002018054611a71906146c3565b8383611c5f82826102246137c5565b611c7b5760405162461bcd60e51b8152600401610678906145ad565b60008086511190506000878787604051602001611c9a939291906141a5565b60405160208183030381529060405280519060200120905060008215611cea578888604051602001611ccd929190614140565b604051602081830303815290604052805190602001209050611ced565b50875b6000828152600b602052604090205442101580611d3f5750828015611d3f57506000828152600b6020526040902054611d2a9062093a8090614665565b60008281526004602052604090206002015410155b80611d80575082158015611d8057506000828152600b6020526040902054611d6b9062093a8090614665565b60008a81526003602052604090206001015410155b611d9c5760405162461bcd60e51b815260040161067890614588565b6000828152600b60205260408120611db69060010161381f565b1115611df8576000828152600b60205260408120611df291611ddb9160010190613829565b6000848152600b6020526040902060010190613835565b50611d9c565b6000828152600b60209081526040808320839055838352600a90915290208054611e2490600190614665565b81548110611e3457611e34614745565b9060005260206000200154600a60008381526020019081526020016000208781548110611e6357611e63614745565b9060005260206000200181905550600a6000828152602001908152602001600020805480611e9357611e9361472f565b60019003818190600052602060002001600090559055505050505050505050565b600a6020528160005260406000208181548110611ed057600080fd5b90600052602060002001600091509150505481565b8585611ef482826102246137c5565b611f105760405162461bcd60e51b8152600401610678906145ad565b6000888888604051602001611f2793929190614344565b60405160208183030381529060405280519060200120905060008989898989604051602001611f5a95949392919061436f565b60408051601f1981840301815291815281516020928301206000818152600b90935291205490915042101580611fa8575060008281526005602052604081208054611fa4906146c3565b9050115b611fc45760405162461bcd60e51b815260040161067890614588565b60008a8a604051602001611fd9929190614140565b60405160208183030381529060405280519060200120905088604051602001612002919061421a565b60408051601f19818403018152918152815160209283012060008481526008909352912080548890811061203857612038614745565b90600052602060002090600302016000016040516020016120599190614236565b60405160208183030381529060405280519060200120146120a85760405162461bcd60e51b815260206004820152600960248201526857726f6e672074616760b81b6044820152606401610678565b600081815260086020526040902080546120c490600190614665565b815481106120d4576120d4614745565b906000526020600020906003020160086000838152602001908152602001600020878154811061210657612106614745565b90600052602060002090600302016000820181600001908054612128906146c3565b612133929190613a90565b506001820181600101908054612148906146c3565b612153929190613a90565b506002820181600201908054612168906146c3565b612173929190613a90565b50505060008181526008602052604090208054806121935761219361472f565b600082815260208120600019909201916003830201906121b38282613b0b565b6121c1600183016000613b0b565b6121cf600283016000613b0b565b505090555050505050505050505050565b6000818152600b6020526040812060609082906121ff9060010161381f565b6001600160401b038111156122165761221661475b565b60405190808252806020026020018201604052801561223f578160200160208202803683370190505b50905060005b6000858152600b6020526040902061225f9060010161381f565b8110156122bb576000858152600b602052604090206122819060010182613829565b82828151811061229357612293614745565b6001600160a01b03909216602092830291909101909101526122b4816146fe565b9050612245565b506000938452600b60205260409093205493915050565b600080836000805160206147b28339815191526040516020016122ff929190918252602082015260400190565b60408051601f19818403018152918152815160209283012060008181526006909352912090915061233090846137a3565b9150505b92915050565b82612347816104826137c5565b6123635760405162461bcd60e51b8152600401610678906145ad565b8160046000868660405160200161237b929190614140565b60405160208183030381529060405280519060200120815260200190815260200160002060030190805190602001906123b5929190613a0c565b506123be6137c5565b6001600160a01b0316836040516123d5919061421a565b6040518091039020857f082dfc0cda6cc25674875e2dcd882c68a8b2c80bf48e58c6a3dc2384c69859558560405161240d919061446a565b60405180910390a450505050565b6009602052816000526040600020818154811061243757600080fd5b6000918252602090912001546001600160a01b03169150829050565b81612460816104826137c5565b61247c5760405162461bcd60e51b8152600401610678906145ad565b600083815260036020908152604090912083516124a192600290920191850190613a0c565b506124aa6137c5565b6040516001600160a01b0391909116907fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709085907f082dfc0cda6cc25674875e2dcd882c68a8b2c80bf48e58c6a3dc2384c69859559061250b90879061446a565b60405180910390a4505050565b82612525816104826137c5565b6125415760405162461bcd60e51b8152600401610678906145ad565b60008484604051602001612556929190614140565b60408051601f1981840301815291815281516020928301206000818152600490935291205490915060ff16156125bc5760405162461bcd60e51b815260206004820152600b60248201526a5265706f2065786973747360a81b6044820152606401610678565b60008451116125fb5760405162461bcd60e51b815260206004820152600b60248201526a4e6f207265706f4e616d6560a81b6044820152606401610678565b600083511161263a5760405162461bcd60e51b815260206004820152600b60248201526a4e6f207265706f4d65746160a81b6044820152606401610678565b60008581526003602081815260408320909101805460018101825590835291819020865161266f939190910191870190613a0c565b506000818152600460209081526040909120805460ff1916600117815584516126a092600390920191860190613a0c565b50826040516126af919061421a565b6040518091039020846040516126c5919061421a565b6040518091039020867f50b56e7c402d556bc61b6fc9bba647b83c4590f2fca5a6d463450e78f3a2d44c87876040516126ff92919061447d565b60405180910390a45050505050565b838361271d82826102246137c5565b6127395760405162461bcd60e51b8152600401610678906145ad565b60008051602061477283398151915284148061277457507fb85a6b22a26ce426daea6357c60a74cb0b4d36234cd7c96170cd6a64102786ff84145b8061279e57507fd7381cbf120651c7bcd9cff77aa8449385520247d98372358854435c90a30af484145b6127d75760405162461bcd60e51b815260206004820152600a6024820152690496e76616c6964206f760b41b6044820152606401610678565b60008086511190506000806000808a8a896040516020016127fa93929190614166565b604051602081830303815290604052805190602001209050600085156128d9578b8b6000805160206147928339815191528c8c6040516020016128419594939291906141d2565b6040516020818303038152906040528051906020012094508b8b60008051602061479283398151915260405160200161287c939291906141a5565b6040516020818303038152906040528051906020012093508b8b6040516020016128a7929190614140565b60408051601f19818403018152918152815160209283012060008181526004909352912060010154909350905061297b565b60408051602081018e90526000805160206147b28339815191529181019190915260608082018c90528a901b6001600160601b031916608082015260940160408051601f1981840301815282825280516020918201209083018f90526000805160206147b283398151915291830191909152955060600160408051601f19818403018152918152815160209283012060008f8152600390935291205490945090505b7fb85a6b22a26ce426daea6357c60a74cb0b4d36234cd7c96170cd6a64102786ff8a14156129f55760008481526006602052604090206129bb908a6137a3565b6129f05760405162461bcd60e51b81526020600482015260066024820152654e6f204b657960d01b6044820152606401610678565b612a47565b6000848152600660205260409020612a0d908a6137a3565b15612a475760405162461bcd60e51b815260206004820152600a6024820152694b65792065786973747360b01b6044820152606401610678565b7fd7381cbf120651c7bcd9cff77aa8449385520247d98372358854435c90a30af48a1415612ae857612a8e612a7a6137c5565b6000868152600660205260409020906137a3565b612aaa5760405162461bcd60e51b8152600401610678906145ad565b612ac9612ab56137c5565b600086815260066020526040902090613835565b506000848152600660205260409020612ae2908a61380a565b50612edc565b60018111612b36576000805160206147728339815191528a1415612b1e576000848152600660205260409020612ae2908a61380a565b6000848152600660205260409020612ae2908a613835565b6000858152600b6020526040902054612be357612b564262093a8061462e565b6000868152600b60205260409020558515612ba75760008381526009602090815260408220805460018101825590835291200180546001600160a01b0319166001600160a01b038b16179055612c46565b60008c81526009602090815260408220805460018101825590835291200180546001600160a01b0319166001600160a01b038b16179055612c46565b6000858152600b602052604090205442111580612c2a57506000858152600b6020526040902054612c189062093a8090614665565b60008381526007602052604090205411155b612c465760405162461bcd60e51b8152600401610678906145cd565b612c68612c516137c5565b6000878152600b602052604090206001019061380a565b5060005b6000868152600b60205260409020612c869060010161381f565b811015612d3d57868015612cba57506000868152600b60205260409020612cb8908e908e906102249060010185613829565b155b80612cec575086158015612cec57506000868152600b60205260409020612cea908e906104829060010184613829565b155b15612d2d576000868152600b60205260409020612d2790612d109060010183613829565b6000888152600b6020526040902060010190613835565b50600090505b612d36816146fe565b9050612c6c565b506000858152600b602052604090208190612d5a9060010161381f565b10612edc576000805160206147728339815191528a1415612d93576000848152600660205260409020612d8d908a61380a565b50612eca565b6000848152600660205260409020612dab908a613835565b506000612e08600660008f6000805160206147b2833981519152604051602001612ddf929190918252602082015260400190565b60405160208183030381529060405280519060200120815260200190815260200160002061381f565b9050868015612e45575060008581526006602052604090208290600190612e2e9061381f565b612e38908461462e565b612e429190614665565b11155b15612e71576000848152600460205260408120600101805491612e67836146ac565b9190505550612ec8565b86158015612ea3575060008581526006602052604090208290600190612e969061381f565b612ea09190614665565b11155b15612ec85760008d8152600360205260408120805491612ec2836146ac565b91905055505b505b60008281526007602052604090204290555b886001600160a01b03168b604051612ef4919061421a565b60405180910390208d7fa8cea71b77054741d93ac504e0ee90fa3e815d68104468d65f9eb36924a8d590612f266137c5565b60008a8152600b602052604090208f90612f429060010161381f565b604080516001600160a01b0390941684526020840192909252908201526060810186905260800160405180910390a4505050505050505050505050565b6060600082612f8e8582614646565b612f989190614665565b90506000612fa68486614646565b60008781526004602081905260409091200154909150811115612fd85750600085815260046020819052604090912001545b6000846001600160401b03811115612ff257612ff261475b565b60405190808252806020026020018201604052801561302557816020015b60608152602001906001900390816130105790505b5090505b818310156119465760008781526004602081905260409091200180548490811061305557613055614745565b90600052602060002001805461306a906146c3565b80601f0160208091040260200160405190810160405280929190818152602001828054613096906146c3565b80156130e35780601f106130b8576101008083540402835291602001916130e3565b820191906000526020600020905b8154815290600101906020018083116130c657829003601f168201915b50505050508184815181106130fa576130fa614745565b60200260200101819052508261310f906146fe565b9250613029565b600560205260009081526040902080548190613131906146c3565b80601f016020809104026020016040519081016040528092919081815260200182805461315d906146c3565b80156131aa5780601f1061317f576101008083540402835291602001916131aa565b820191906000526020600020905b81548152906001019060200180831161318d57829003601f168201915b5050505050908060010180546131bf906146c3565b80601f01602080910402602001604051908101604052809291908181526020018280546131eb906146c3565b80156132385780601f1061320d57610100808354040283529160200191613238565b820191906000526020600020905b81548152906001019060200180831161321b57829003601f168201915b5050505050905082565b828261325182826102246137c5565b61326d5760405162461bcd60e51b8152600401610678906145ad565b60008085511190506000868660405160200161328a929190614140565b60405160208183030381529060405280519060200120905060008083156132c657505060008181526004602052604090206001015481906132da565b505060008781526003602052604090205487905b8087141561331a5760405162461bcd60e51b815260206004820152600d60248201526c151a1c995cda1bdb19081cd95d609a1b6044820152606401610678565b6000898989604051602001613331939291906141a5565b60405160208183030381529060405280519060200120905061336b6133546137c5565b6000838152600b60205260409020600101906137a3565b156133a55760405162461bcd60e51b815260206004820152600a602482015269155cd95c881d9bdd195960b21b6044820152606401610678565b848015613422575060016133e0600660008d6000805160206147b2833981519152604051602001612ddf929190918252602082015260400190565b61340a600660008e8e600080516020614792833981519152604051602001612ddf939291906141a5565b613414919061462e565b61341e9190614665565b8811155b8061347257508415801561347257506001613464600660008d6000805160206147b2833981519152604051602001612ddf929190918252602082015260400190565b61346e9190614665565b8811155b6134b35760405162461bcd60e51b81526020600482015260126024820152714e6f7420656e6f756768206d656d6265727360701b6044820152606401610678565b6000818152600b6020526040902054613509576134d34262093a8061462e565b6000828152600b6020908152604080832093909355858252600a81529181208054600181018255908252919020018890556135b8565b6000818152600b60205260409020544211158061355b575084801561355b57506000818152600b60205260409020546135469062093a8090614665565b60008581526004602052604090206002015411155b8061359c57508415801561359c57506000818152600b60205260409020546135879062093a8090614665565b60008b81526003602052604090206001015411155b6135b85760405162461bcd60e51b8152600401610678906145cd565b6135da6135c36137c5565b6000838152600b602052604090206001019061380a565b5060005b6000828152600b602052604090206135f89060010161381f565b8110156136985785801561362c57506000828152600b6020526040902061362a908c908c906102249060010185613829565b155b8061365e57508515801561365e57506000828152600b6020526040902061365c908c906104829060010184613829565b155b15613688576000828152600b6020526040902061368290611ddb9060010183613829565b50600090505b613691816146fe565b90506135de565b506000818152600b6020526040902082906136b59060010161381f565b101580156136dd57506000818152600b6020526040902088906136da9060010161381f565b10155b156137225784156137095760008481526004602052604090206001810189905542600290910155613722565b60008a8152600360205260409020888155426001909101555b8789604051613731919061421a565b60405180910390208b7f2daad345bbc18b1fa8a7f542403081df42c546dd841a44507d53e174e761ce896137636137c5565b6000868152600b6020526040902061377d9060010161381f565b604080516001600160a01b03909316835260208301919091528101879052606001610e21565b6001600160a01b038116600090815260018301602052604081205415156106fc565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316331415613805575060131936013560601c90565b503390565b60006106fc836001600160a01b03841661384a565b6000612334825490565b60006106fc8383613899565b60006106fc836001600160a01b03841661391f565b600081815260018301602052604081205461389157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155612334565b506000612334565b815460009082106138f75760405162461bcd60e51b815260206004820152602260248201527f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e604482015261647360f01b6064820152608401610678565b82600001828154811061390c5761390c614745565b9060005260206000200154905092915050565b60008181526001830160205260408120548015613a02576000613943600183614665565b855490915060009061395790600190614665565b9050600086600001828154811061397057613970614745565b906000526020600020015490508087600001848154811061399357613993614745565b6000918252602080832090910192909255828152600189019091526040902084905586548790806139c6576139c661472f565b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050612334565b6000915050612334565b828054613a18906146c3565b90600052602060002090601f016020900481019282613a3a5760008555613a80565b82601f10613a5357805160ff1916838001178555613a80565b82800160010185558215613a80579182015b82811115613a80578251825591602001919060010190613a65565b50613a8c929150613b48565b5090565b828054613a9c906146c3565b90600052602060002090601f016020900481019282613abe5760008555613a80565b82601f10613acf5780548555613a80565b82800160010185558215613a8057600052602060002091601f016020900482015b82811115613a80578254825591600101919060010190613af0565b508054613b17906146c3565b6000825580601f10613b27575050565b601f016020900490600052602060002090810190613b459190613b48565b50565b5b80821115613a8c5760008155600101613b49565b80356001600160a01b0381168114613b7457600080fd5b919050565b600082601f830112613b8a57600080fd5b81356001600160401b0380821115613ba457613ba461475b565b604051601f8301601f19908116603f01168101908282118183101715613bcc57613bcc61475b565b81604052838152866020858801011115613be557600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215613c1757600080fd5b6106fc82613b5d565b600060208284031215613c3257600080fd5b5035919050565b60008060408385031215613c4c57600080fd5b82359150613c5c60208401613b5d565b90509250929050565b60008060408385031215613c7857600080fd5b8235915060208301356001600160401b03811115613c9557600080fd5b613ca185828601613b79565b9150509250929050565b600080600060608486031215613cc057600080fd5b8335925060208401356001600160401b03811115613cdd57600080fd5b613ce986828701613b79565b925050613cf860408501613b5d565b90509250925092565b60008060008060808587031215613d1757600080fd5b8435935060208501356001600160401b03811115613d3457600080fd5b613d4087828801613b79565b93505060408501359150613d5660608601613b5d565b905092959194509250565b600080600080600060a08688031215613d7957600080fd5b8535945060208601356001600160401b03811115613d9657600080fd5b613da288828901613b79565b94505060408601359250613db860608701613b5d565b949793965091946080013592915050565b600080600060608486031215613dde57600080fd5b8335925060208401356001600160401b0380821115613dfc57600080fd5b613e0887838801613b79565b93506040860135915080821115613e1e57600080fd5b50613e2b86828701613b79565b9150509250925092565b600080600080600060a08688031215613e4d57600080fd5b8535945060208601356001600160401b0380821115613e6b57600080fd5b613e7789838a01613b79565b95506040880135915080821115613e8d57600080fd5b613e9989838a01613b79565b94506060880135915080821115613eaf57600080fd5b613ebb89838a01613b79565b93506080880135915080821115613ed157600080fd5b50613ede88828901613b79565b9150509295509295909350565b60008060008060008060c08789031215613f0457600080fd5b8635955060208701356001600160401b0380821115613f2257600080fd5b613f2e8a838b01613b79565b96506040890135915080821115613f4457600080fd5b613f508a838b01613b79565b95506060890135915080821115613f6657600080fd5b613f728a838b01613b79565b94506080890135915080821115613f8857600080fd5b50613f9589828a01613b79565b92505060a087013590509295509295509295565b600080600060608486031215613fbe57600080fd5b8335925060208401356001600160401b03811115613fdb57600080fd5b613fe786828701613b79565b925050604084013590509250925092565b6000806000806080858703121561400e57600080fd5b8435935060208501356001600160401b0381111561402b57600080fd5b61403787828801613b79565b949794965050505060408301359260600135919050565b6000806040838503121561406157600080fd5b50508035926020909101359150565b60008060006060848603121561408557600080fd5b505081359360208301359350604090920135919050565b6000602082840312156140ae57600080fd5b81356001600160401b038111156140c457600080fd5b61233084828501613b79565b600081518084526020808501945080840160005b838110156141095781516001600160a01b0316875295820195908201906001016140e4565b509495945050505050565b6000815180845261412c81602086016020860161467c565b601f01601f19169290920160200192915050565b8281526000825161415881602085016020870161467c565b919091016020019392505050565b8381526000835161417e81602085016020880161467c565b60609390931b6001600160601b031916602092909301918201929092526034019392505050565b838152600083516141bd81602085016020880161467c565b60209201918201929092526040019392505050565b858152600085516141ea816020850160208a0161467c565b60209201918201949094526040810192909252606090811b6001600160601b031916908201526074019392505050565b6000825161422c81846020870161467c565b9190910192915050565b6000808354614244816146c3565b6001828116801561425c576001811461426d5761429c565b60ff1984168752828701945061429c565b8760005260208060002060005b858110156142935781548a82015290840190820161427a565b50505082870194505b50929695505050505050565b6020815260006106fc60208301846140d0565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561431057603f198886030184526142fe858351614114565b945092850192908501906001016142e2565b5092979650505050505050565b84151581528360208201528260408201526080606082015260006119466080830184614114565b83815260606020820152600061435d6060830185614114565b82810360408401526119468185614114565b85815260a06020820152600061438860a0830187614114565b828103604084015261439a8187614114565b905082810360608401526143ae8186614114565b905082810360808401526143c28185614114565b98975050505050505050565b838152600060206060818401526143e86060840186614114565b8381036040850152600085546143fd816146c3565b80845260018281168015614418576001811461442c5761445a565b60ff1984168688015260408601945061445a565b896000528660002060005b848110156144525781548882018a0152908301908801614437565b870188019550505b50929a9950505050505050505050565b6020815260006106fc6020830184614114565b6040815260006144906040830185614114565b82810360208401526144a28185614114565b95945050505050565b60a0815260006144be60a0830188614114565b82810360208401526144d08188614114565b6001600160a01b039690961660408401525050606081019290925260809091015292915050565b60608152600061450a6060830186614114565b828103602084015261451c8186614114565b905082810360408401526119468185614114565b6080815260006145436080830187614114565b82810360208401526145558187614114565b905082810360408401526145698186614114565b9050828103606084015261457d81856140d0565b979650505050505050565b6020808252600b908201526a139bdd08195e1c1a5c995960aa1b604082015260600190565b60208082526006908201526511195b9a595960d21b604082015260600190565b602080825260079082015266115e1c1a5c995960ca1b604082015260600190565b82815260406020820152600061460760408301846140d0565b949350505050565b8381528260208201526060604082015260006144a26060830184614114565b6000821982111561464157614641614719565b500190565b600081600019048311821515161561466057614660614719565b500290565b60008282101561467757614677614719565b500390565b60005b8381101561469757818101518382015260200161467f565b838111156146a6576000848401525b50505050565b6000816146bb576146bb614719565b506000190190565b600181811c908216806146d757607f821691505b602082108114156146f857634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561471257614712614719565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfedb7bfe957d82520d83f2a439f82c0ab65bbe62c5da8b7b7aa22924c027ef518d069bf569f27d389f2c70410107860b2e82ff561283b097a89e897daa5e34b1b6123b642491709420c2370bb98c4e7de2b1bc05c5f9fd95ac4111e12683553c62a2646970667358221220cb268f26ce5ca82d1969c92fcc776c7700c671b2f112e7aa1ea22fc8aa602d1164736f6c63430008060033"

// DeployValist deploys a new Ethereum contract, binding an instance of Valist to it.
func DeployValist(auth *bind.TransactOpts, backend bind.ContractBackend, metaTxForwarder common.Address) (common.Address, *types.Transaction, *Valist, error) {
	parsed, err := abi.JSON(strings.NewReader(ValistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValistBin), backend, metaTxForwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Valist{ValistCaller: ValistCaller{contract: contract}, ValistTransactor: ValistTransactor{contract: contract}, ValistFilterer: ValistFilterer{contract: contract}}, nil
}

// Valist is an auto generated Go binding around an Ethereum contract.
type Valist struct {
	ValistCaller     // Read-only binding to the contract
	ValistTransactor // Write-only binding to the contract
	ValistFilterer   // Log filterer for contract events
}

// ValistCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValistSession struct {
	Contract     *Valist           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValistCallerSession struct {
	Contract *ValistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValistTransactorSession struct {
	Contract     *ValistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValistRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValistRaw struct {
	Contract *Valist // Generic contract binding to access the raw methods on
}

// ValistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValistCallerRaw struct {
	Contract *ValistCaller // Generic read-only contract binding to access the raw methods on
}

// ValistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValistTransactorRaw struct {
	Contract *ValistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValist creates a new instance of Valist, bound to a specific deployed contract.
func NewValist(address common.Address, backend bind.ContractBackend) (*Valist, error) {
	contract, err := bindValist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Valist{ValistCaller: ValistCaller{contract: contract}, ValistTransactor: ValistTransactor{contract: contract}, ValistFilterer: ValistFilterer{contract: contract}}, nil
}

// NewValistCaller creates a new read-only instance of Valist, bound to a specific deployed contract.
func NewValistCaller(address common.Address, caller bind.ContractCaller) (*ValistCaller, error) {
	contract, err := bindValist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValistCaller{contract: contract}, nil
}

// NewValistTransactor creates a new write-only instance of Valist, bound to a specific deployed contract.
func NewValistTransactor(address common.Address, transactor bind.ContractTransactor) (*ValistTransactor, error) {
	contract, err := bindValist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValistTransactor{contract: contract}, nil
}

// NewValistFilterer creates a new log filterer instance of Valist, bound to a specific deployed contract.
func NewValistFilterer(address common.Address, filterer bind.ContractFilterer) (*ValistFilterer, error) {
	contract, err := bindValist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValistFilterer{contract: contract}, nil
}

// bindValist binds a generic wrapper to an already deployed contract.
func bindValist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Valist *ValistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Valist.Contract.ValistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Valist *ValistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Valist.Contract.ValistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Valist *ValistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Valist.Contract.ValistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Valist *ValistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Valist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Valist *ValistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Valist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Valist *ValistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Valist.Contract.contract.Transact(opts, method, params...)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0x5cd20d29.
//
// Solidity: function getLatestRelease(bytes32 _orgID, string _repoName) view returns(string, string, string, address[])
func (_Valist *ValistCaller) GetLatestRelease(opts *bind.CallOpts, _orgID [32]byte, _repoName string) (string, string, string, []common.Address, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "getLatestRelease", _orgID, _repoName)

	if err != nil {
		return *new(string), *new(string), *new(string), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return out0, out1, out2, out3, err

}

// GetLatestRelease is a free data retrieval call binding the contract method 0x5cd20d29.
//
// Solidity: function getLatestRelease(bytes32 _orgID, string _repoName) view returns(string, string, string, address[])
func (_Valist *ValistSession) GetLatestRelease(_orgID [32]byte, _repoName string) (string, string, string, []common.Address, error) {
	return _Valist.Contract.GetLatestRelease(&_Valist.CallOpts, _orgID, _repoName)
}

// GetLatestRelease is a free data retrieval call binding the contract method 0x5cd20d29.
//
// Solidity: function getLatestRelease(bytes32 _orgID, string _repoName) view returns(string, string, string, address[])
func (_Valist *ValistCallerSession) GetLatestRelease(_orgID [32]byte, _repoName string) (string, string, string, []common.Address, error) {
	return _Valist.Contract.GetLatestRelease(&_Valist.CallOpts, _orgID, _repoName)
}

// GetPendingReleaseCount is a free data retrieval call binding the contract method 0x924a5443.
//
// Solidity: function getPendingReleaseCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistCaller) GetPendingReleaseCount(opts *bind.CallOpts, _selector [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "getPendingReleaseCount", _selector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingReleaseCount is a free data retrieval call binding the contract method 0x924a5443.
//
// Solidity: function getPendingReleaseCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistSession) GetPendingReleaseCount(_selector [32]byte) (*big.Int, error) {
	return _Valist.Contract.GetPendingReleaseCount(&_Valist.CallOpts, _selector)
}

// GetPendingReleaseCount is a free data retrieval call binding the contract method 0x924a5443.
//
// Solidity: function getPendingReleaseCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistCallerSession) GetPendingReleaseCount(_selector [32]byte) (*big.Int, error) {
	return _Valist.Contract.GetPendingReleaseCount(&_Valist.CallOpts, _selector)
}

// GetPendingVotes is a free data retrieval call binding the contract method 0xd565e860.
//
// Solidity: function getPendingVotes(bytes32 _selector) view returns(uint256, address[])
func (_Valist *ValistCaller) GetPendingVotes(opts *bind.CallOpts, _selector [32]byte) (*big.Int, []common.Address, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "getPendingVotes", _selector)

	if err != nil {
		return *new(*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetPendingVotes is a free data retrieval call binding the contract method 0xd565e860.
//
// Solidity: function getPendingVotes(bytes32 _selector) view returns(uint256, address[])
func (_Valist *ValistSession) GetPendingVotes(_selector [32]byte) (*big.Int, []common.Address, error) {
	return _Valist.Contract.GetPendingVotes(&_Valist.CallOpts, _selector)
}

// GetPendingVotes is a free data retrieval call binding the contract method 0xd565e860.
//
// Solidity: function getPendingVotes(bytes32 _selector) view returns(uint256, address[])
func (_Valist *ValistCallerSession) GetPendingVotes(_selector [32]byte) (*big.Int, []common.Address, error) {
	return _Valist.Contract.GetPendingVotes(&_Valist.CallOpts, _selector)
}

// GetReleaseTags is a free data retrieval call binding the contract method 0xf3439d56.
//
// Solidity: function getReleaseTags(bytes32 _selector, uint256 _page, uint256 _resultsPerPage) view returns(string[])
func (_Valist *ValistCaller) GetReleaseTags(opts *bind.CallOpts, _selector [32]byte, _page *big.Int, _resultsPerPage *big.Int) ([]string, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "getReleaseTags", _selector, _page, _resultsPerPage)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetReleaseTags is a free data retrieval call binding the contract method 0xf3439d56.
//
// Solidity: function getReleaseTags(bytes32 _selector, uint256 _page, uint256 _resultsPerPage) view returns(string[])
func (_Valist *ValistSession) GetReleaseTags(_selector [32]byte, _page *big.Int, _resultsPerPage *big.Int) ([]string, error) {
	return _Valist.Contract.GetReleaseTags(&_Valist.CallOpts, _selector, _page, _resultsPerPage)
}

// GetReleaseTags is a free data retrieval call binding the contract method 0xf3439d56.
//
// Solidity: function getReleaseTags(bytes32 _selector, uint256 _page, uint256 _resultsPerPage) view returns(string[])
func (_Valist *ValistCallerSession) GetReleaseTags(_selector [32]byte, _page *big.Int, _resultsPerPage *big.Int) ([]string, error) {
	return _Valist.Contract.GetReleaseTags(&_Valist.CallOpts, _selector, _page, _resultsPerPage)
}

// GetRepoNames is a free data retrieval call binding the contract method 0x837dc4d0.
//
// Solidity: function getRepoNames(bytes32 _orgID, uint256 _page, uint256 _resultsPerPage) view returns(string[])
func (_Valist *ValistCaller) GetRepoNames(opts *bind.CallOpts, _orgID [32]byte, _page *big.Int, _resultsPerPage *big.Int) ([]string, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "getRepoNames", _orgID, _page, _resultsPerPage)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetRepoNames is a free data retrieval call binding the contract method 0x837dc4d0.
//
// Solidity: function getRepoNames(bytes32 _orgID, uint256 _page, uint256 _resultsPerPage) view returns(string[])
func (_Valist *ValistSession) GetRepoNames(_orgID [32]byte, _page *big.Int, _resultsPerPage *big.Int) ([]string, error) {
	return _Valist.Contract.GetRepoNames(&_Valist.CallOpts, _orgID, _page, _resultsPerPage)
}

// GetRepoNames is a free data retrieval call binding the contract method 0x837dc4d0.
//
// Solidity: function getRepoNames(bytes32 _orgID, uint256 _page, uint256 _resultsPerPage) view returns(string[])
func (_Valist *ValistCallerSession) GetRepoNames(_orgID [32]byte, _page *big.Int, _resultsPerPage *big.Int) ([]string, error) {
	return _Valist.Contract.GetRepoNames(&_Valist.CallOpts, _orgID, _page, _resultsPerPage)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 _selector) view returns(address[])
func (_Valist *ValistCaller) GetRoleMembers(opts *bind.CallOpts, _selector [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "getRoleMembers", _selector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 _selector) view returns(address[])
func (_Valist *ValistSession) GetRoleMembers(_selector [32]byte) ([]common.Address, error) {
	return _Valist.Contract.GetRoleMembers(&_Valist.CallOpts, _selector)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 _selector) view returns(address[])
func (_Valist *ValistCallerSession) GetRoleMembers(_selector [32]byte) ([]common.Address, error) {
	return _Valist.Contract.GetRoleMembers(&_Valist.CallOpts, _selector)
}

// GetRoleRequestCount is a free data retrieval call binding the contract method 0x3e1bf64b.
//
// Solidity: function getRoleRequestCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistCaller) GetRoleRequestCount(opts *bind.CallOpts, _selector [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "getRoleRequestCount", _selector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleRequestCount is a free data retrieval call binding the contract method 0x3e1bf64b.
//
// Solidity: function getRoleRequestCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistSession) GetRoleRequestCount(_selector [32]byte) (*big.Int, error) {
	return _Valist.Contract.GetRoleRequestCount(&_Valist.CallOpts, _selector)
}

// GetRoleRequestCount is a free data retrieval call binding the contract method 0x3e1bf64b.
//
// Solidity: function getRoleRequestCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistCallerSession) GetRoleRequestCount(_selector [32]byte) (*big.Int, error) {
	return _Valist.Contract.GetRoleRequestCount(&_Valist.CallOpts, _selector)
}

// GetThresholdRequestCount is a free data retrieval call binding the contract method 0x3b0659fd.
//
// Solidity: function getThresholdRequestCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistCaller) GetThresholdRequestCount(opts *bind.CallOpts, _selector [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "getThresholdRequestCount", _selector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThresholdRequestCount is a free data retrieval call binding the contract method 0x3b0659fd.
//
// Solidity: function getThresholdRequestCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistSession) GetThresholdRequestCount(_selector [32]byte) (*big.Int, error) {
	return _Valist.Contract.GetThresholdRequestCount(&_Valist.CallOpts, _selector)
}

// GetThresholdRequestCount is a free data retrieval call binding the contract method 0x3b0659fd.
//
// Solidity: function getThresholdRequestCount(bytes32 _selector) view returns(uint256)
func (_Valist *ValistCallerSession) GetThresholdRequestCount(_selector [32]byte) (*big.Int, error) {
	return _Valist.Contract.GetThresholdRequestCount(&_Valist.CallOpts, _selector)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0xd96162a1.
//
// Solidity: function isOrgAdmin(bytes32 _orgID, address _address) view returns(bool)
func (_Valist *ValistCaller) IsOrgAdmin(opts *bind.CallOpts, _orgID [32]byte, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "isOrgAdmin", _orgID, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrgAdmin is a free data retrieval call binding the contract method 0xd96162a1.
//
// Solidity: function isOrgAdmin(bytes32 _orgID, address _address) view returns(bool)
func (_Valist *ValistSession) IsOrgAdmin(_orgID [32]byte, _address common.Address) (bool, error) {
	return _Valist.Contract.IsOrgAdmin(&_Valist.CallOpts, _orgID, _address)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0xd96162a1.
//
// Solidity: function isOrgAdmin(bytes32 _orgID, address _address) view returns(bool)
func (_Valist *ValistCallerSession) IsOrgAdmin(_orgID [32]byte, _address common.Address) (bool, error) {
	return _Valist.Contract.IsOrgAdmin(&_Valist.CallOpts, _orgID, _address)
}

// IsRepoDev is a free data retrieval call binding the contract method 0x1fd522a6.
//
// Solidity: function isRepoDev(bytes32 _orgID, string _repoName, address _address) view returns(bool)
func (_Valist *ValistCaller) IsRepoDev(opts *bind.CallOpts, _orgID [32]byte, _repoName string, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "isRepoDev", _orgID, _repoName, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRepoDev is a free data retrieval call binding the contract method 0x1fd522a6.
//
// Solidity: function isRepoDev(bytes32 _orgID, string _repoName, address _address) view returns(bool)
func (_Valist *ValistSession) IsRepoDev(_orgID [32]byte, _repoName string, _address common.Address) (bool, error) {
	return _Valist.Contract.IsRepoDev(&_Valist.CallOpts, _orgID, _repoName, _address)
}

// IsRepoDev is a free data retrieval call binding the contract method 0x1fd522a6.
//
// Solidity: function isRepoDev(bytes32 _orgID, string _repoName, address _address) view returns(bool)
func (_Valist *ValistCallerSession) IsRepoDev(_orgID [32]byte, _repoName string, _address common.Address) (bool, error) {
	return _Valist.Contract.IsRepoDev(&_Valist.CallOpts, _orgID, _repoName, _address)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Valist *ValistCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Valist *ValistSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Valist.Contract.IsTrustedForwarder(&_Valist.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Valist *ValistCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Valist.Contract.IsTrustedForwarder(&_Valist.CallOpts, forwarder)
}

// OrgCount is a free data retrieval call binding the contract method 0xd8106690.
//
// Solidity: function orgCount() view returns(uint256)
func (_Valist *ValistCaller) OrgCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "orgCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrgCount is a free data retrieval call binding the contract method 0xd8106690.
//
// Solidity: function orgCount() view returns(uint256)
func (_Valist *ValistSession) OrgCount() (*big.Int, error) {
	return _Valist.Contract.OrgCount(&_Valist.CallOpts)
}

// OrgCount is a free data retrieval call binding the contract method 0xd8106690.
//
// Solidity: function orgCount() view returns(uint256)
func (_Valist *ValistCallerSession) OrgCount() (*big.Int, error) {
	return _Valist.Contract.OrgCount(&_Valist.CallOpts)
}

// OrgIDs is a free data retrieval call binding the contract method 0x8775692f.
//
// Solidity: function orgIDs(uint256 ) view returns(bytes32)
func (_Valist *ValistCaller) OrgIDs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "orgIDs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OrgIDs is a free data retrieval call binding the contract method 0x8775692f.
//
// Solidity: function orgIDs(uint256 ) view returns(bytes32)
func (_Valist *ValistSession) OrgIDs(arg0 *big.Int) ([32]byte, error) {
	return _Valist.Contract.OrgIDs(&_Valist.CallOpts, arg0)
}

// OrgIDs is a free data retrieval call binding the contract method 0x8775692f.
//
// Solidity: function orgIDs(uint256 ) view returns(bytes32)
func (_Valist *ValistCallerSession) OrgIDs(arg0 *big.Int) ([32]byte, error) {
	return _Valist.Contract.OrgIDs(&_Valist.CallOpts, arg0)
}

// Orgs is a free data retrieval call binding the contract method 0xa3e84beb.
//
// Solidity: function orgs(bytes32 ) view returns(uint256 threshold, uint256 thresholdDate, string metaCID)
func (_Valist *ValistCaller) Orgs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       string
}, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "orgs", arg0)

	outstruct := new(struct {
		Threshold     *big.Int
		ThresholdDate *big.Int
		MetaCID       string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Threshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ThresholdDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MetaCID = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Orgs is a free data retrieval call binding the contract method 0xa3e84beb.
//
// Solidity: function orgs(bytes32 ) view returns(uint256 threshold, uint256 thresholdDate, string metaCID)
func (_Valist *ValistSession) Orgs(arg0 [32]byte) (struct {
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       string
}, error) {
	return _Valist.Contract.Orgs(&_Valist.CallOpts, arg0)
}

// Orgs is a free data retrieval call binding the contract method 0xa3e84beb.
//
// Solidity: function orgs(bytes32 ) view returns(uint256 threshold, uint256 thresholdDate, string metaCID)
func (_Valist *ValistCallerSession) Orgs(arg0 [32]byte) (struct {
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       string
}, error) {
	return _Valist.Contract.Orgs(&_Valist.CallOpts, arg0)
}

// PendingReleaseRequests is a free data retrieval call binding the contract method 0xa940abcb.
//
// Solidity: function pendingReleaseRequests(bytes32 , uint256 ) view returns(string tag, string releaseCID, string metaCID)
func (_Valist *ValistCaller) PendingReleaseRequests(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (struct {
	Tag        string
	ReleaseCID string
	MetaCID    string
}, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "pendingReleaseRequests", arg0, arg1)

	outstruct := new(struct {
		Tag        string
		ReleaseCID string
		MetaCID    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tag = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ReleaseCID = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.MetaCID = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// PendingReleaseRequests is a free data retrieval call binding the contract method 0xa940abcb.
//
// Solidity: function pendingReleaseRequests(bytes32 , uint256 ) view returns(string tag, string releaseCID, string metaCID)
func (_Valist *ValistSession) PendingReleaseRequests(arg0 [32]byte, arg1 *big.Int) (struct {
	Tag        string
	ReleaseCID string
	MetaCID    string
}, error) {
	return _Valist.Contract.PendingReleaseRequests(&_Valist.CallOpts, arg0, arg1)
}

// PendingReleaseRequests is a free data retrieval call binding the contract method 0xa940abcb.
//
// Solidity: function pendingReleaseRequests(bytes32 , uint256 ) view returns(string tag, string releaseCID, string metaCID)
func (_Valist *ValistCallerSession) PendingReleaseRequests(arg0 [32]byte, arg1 *big.Int) (struct {
	Tag        string
	ReleaseCID string
	MetaCID    string
}, error) {
	return _Valist.Contract.PendingReleaseRequests(&_Valist.CallOpts, arg0, arg1)
}

// PendingRoleRequests is a free data retrieval call binding the contract method 0xdfec1bb0.
//
// Solidity: function pendingRoleRequests(bytes32 , uint256 ) view returns(address)
func (_Valist *ValistCaller) PendingRoleRequests(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "pendingRoleRequests", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingRoleRequests is a free data retrieval call binding the contract method 0xdfec1bb0.
//
// Solidity: function pendingRoleRequests(bytes32 , uint256 ) view returns(address)
func (_Valist *ValistSession) PendingRoleRequests(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Valist.Contract.PendingRoleRequests(&_Valist.CallOpts, arg0, arg1)
}

// PendingRoleRequests is a free data retrieval call binding the contract method 0xdfec1bb0.
//
// Solidity: function pendingRoleRequests(bytes32 , uint256 ) view returns(address)
func (_Valist *ValistCallerSession) PendingRoleRequests(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Valist.Contract.PendingRoleRequests(&_Valist.CallOpts, arg0, arg1)
}

// PendingThresholdRequests is a free data retrieval call binding the contract method 0xc372be80.
//
// Solidity: function pendingThresholdRequests(bytes32 , uint256 ) view returns(uint256)
func (_Valist *ValistCaller) PendingThresholdRequests(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "pendingThresholdRequests", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingThresholdRequests is a free data retrieval call binding the contract method 0xc372be80.
//
// Solidity: function pendingThresholdRequests(bytes32 , uint256 ) view returns(uint256)
func (_Valist *ValistSession) PendingThresholdRequests(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _Valist.Contract.PendingThresholdRequests(&_Valist.CallOpts, arg0, arg1)
}

// PendingThresholdRequests is a free data retrieval call binding the contract method 0xc372be80.
//
// Solidity: function pendingThresholdRequests(bytes32 , uint256 ) view returns(uint256)
func (_Valist *ValistCallerSession) PendingThresholdRequests(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _Valist.Contract.PendingThresholdRequests(&_Valist.CallOpts, arg0, arg1)
}

// Releases is a free data retrieval call binding the contract method 0xf491a84c.
//
// Solidity: function releases(bytes32 ) view returns(string releaseCID, string metaCID)
func (_Valist *ValistCaller) Releases(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ReleaseCID string
	MetaCID    string
}, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "releases", arg0)

	outstruct := new(struct {
		ReleaseCID string
		MetaCID    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReleaseCID = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.MetaCID = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Releases is a free data retrieval call binding the contract method 0xf491a84c.
//
// Solidity: function releases(bytes32 ) view returns(string releaseCID, string metaCID)
func (_Valist *ValistSession) Releases(arg0 [32]byte) (struct {
	ReleaseCID string
	MetaCID    string
}, error) {
	return _Valist.Contract.Releases(&_Valist.CallOpts, arg0)
}

// Releases is a free data retrieval call binding the contract method 0xf491a84c.
//
// Solidity: function releases(bytes32 ) view returns(string releaseCID, string metaCID)
func (_Valist *ValistCallerSession) Releases(arg0 [32]byte) (struct {
	ReleaseCID string
	MetaCID    string
}, error) {
	return _Valist.Contract.Releases(&_Valist.CallOpts, arg0)
}

// Repos is a free data retrieval call binding the contract method 0x02b2583a.
//
// Solidity: function repos(bytes32 ) view returns(bool exists, uint256 threshold, uint256 thresholdDate, string metaCID)
func (_Valist *ValistCaller) Repos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Exists        bool
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       string
}, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "repos", arg0)

	outstruct := new(struct {
		Exists        bool
		Threshold     *big.Int
		ThresholdDate *big.Int
		MetaCID       string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Threshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ThresholdDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MetaCID = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Repos is a free data retrieval call binding the contract method 0x02b2583a.
//
// Solidity: function repos(bytes32 ) view returns(bool exists, uint256 threshold, uint256 thresholdDate, string metaCID)
func (_Valist *ValistSession) Repos(arg0 [32]byte) (struct {
	Exists        bool
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       string
}, error) {
	return _Valist.Contract.Repos(&_Valist.CallOpts, arg0)
}

// Repos is a free data retrieval call binding the contract method 0x02b2583a.
//
// Solidity: function repos(bytes32 ) view returns(bool exists, uint256 threshold, uint256 thresholdDate, string metaCID)
func (_Valist *ValistCallerSession) Repos(arg0 [32]byte) (struct {
	Exists        bool
	Threshold     *big.Int
	ThresholdDate *big.Int
	MetaCID       string
}, error) {
	return _Valist.Contract.Repos(&_Valist.CallOpts, arg0)
}

// RoleModifiedTimestamps is a free data retrieval call binding the contract method 0x8bf67370.
//
// Solidity: function roleModifiedTimestamps(bytes32 ) view returns(uint256)
func (_Valist *ValistCaller) RoleModifiedTimestamps(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "roleModifiedTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoleModifiedTimestamps is a free data retrieval call binding the contract method 0x8bf67370.
//
// Solidity: function roleModifiedTimestamps(bytes32 ) view returns(uint256)
func (_Valist *ValistSession) RoleModifiedTimestamps(arg0 [32]byte) (*big.Int, error) {
	return _Valist.Contract.RoleModifiedTimestamps(&_Valist.CallOpts, arg0)
}

// RoleModifiedTimestamps is a free data retrieval call binding the contract method 0x8bf67370.
//
// Solidity: function roleModifiedTimestamps(bytes32 ) view returns(uint256)
func (_Valist *ValistCallerSession) RoleModifiedTimestamps(arg0 [32]byte) (*big.Int, error) {
	return _Valist.Contract.RoleModifiedTimestamps(&_Valist.CallOpts, arg0)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Valist *ValistCaller) VersionRecipient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Valist.contract.Call(opts, &out, "versionRecipient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Valist *ValistSession) VersionRecipient() (string, error) {
	return _Valist.Contract.VersionRecipient(&_Valist.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Valist *ValistCallerSession) VersionRecipient() (string, error) {
	return _Valist.Contract.VersionRecipient(&_Valist.CallOpts)
}

// ClearPendingKey is a paid mutator transaction binding the contract method 0x40fd48a7.
//
// Solidity: function clearPendingKey(bytes32 _orgID, string _repoName, bytes32 _operation, address _key, uint256 _requestIndex) returns()
func (_Valist *ValistTransactor) ClearPendingKey(opts *bind.TransactOpts, _orgID [32]byte, _repoName string, _operation [32]byte, _key common.Address, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "clearPendingKey", _orgID, _repoName, _operation, _key, _requestIndex)
}

// ClearPendingKey is a paid mutator transaction binding the contract method 0x40fd48a7.
//
// Solidity: function clearPendingKey(bytes32 _orgID, string _repoName, bytes32 _operation, address _key, uint256 _requestIndex) returns()
func (_Valist *ValistSession) ClearPendingKey(_orgID [32]byte, _repoName string, _operation [32]byte, _key common.Address, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.Contract.ClearPendingKey(&_Valist.TransactOpts, _orgID, _repoName, _operation, _key, _requestIndex)
}

// ClearPendingKey is a paid mutator transaction binding the contract method 0x40fd48a7.
//
// Solidity: function clearPendingKey(bytes32 _orgID, string _repoName, bytes32 _operation, address _key, uint256 _requestIndex) returns()
func (_Valist *ValistTransactorSession) ClearPendingKey(_orgID [32]byte, _repoName string, _operation [32]byte, _key common.Address, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.Contract.ClearPendingKey(&_Valist.TransactOpts, _orgID, _repoName, _operation, _key, _requestIndex)
}

// ClearPendingRelease is a paid mutator transaction binding the contract method 0xc6c2a0be.
//
// Solidity: function clearPendingRelease(bytes32 _orgID, string _repoName, string _tag, string _releaseCID, string _metaCID, uint256 _requestIndex) returns()
func (_Valist *ValistTransactor) ClearPendingRelease(opts *bind.TransactOpts, _orgID [32]byte, _repoName string, _tag string, _releaseCID string, _metaCID string, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "clearPendingRelease", _orgID, _repoName, _tag, _releaseCID, _metaCID, _requestIndex)
}

// ClearPendingRelease is a paid mutator transaction binding the contract method 0xc6c2a0be.
//
// Solidity: function clearPendingRelease(bytes32 _orgID, string _repoName, string _tag, string _releaseCID, string _metaCID, uint256 _requestIndex) returns()
func (_Valist *ValistSession) ClearPendingRelease(_orgID [32]byte, _repoName string, _tag string, _releaseCID string, _metaCID string, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.Contract.ClearPendingRelease(&_Valist.TransactOpts, _orgID, _repoName, _tag, _releaseCID, _metaCID, _requestIndex)
}

// ClearPendingRelease is a paid mutator transaction binding the contract method 0xc6c2a0be.
//
// Solidity: function clearPendingRelease(bytes32 _orgID, string _repoName, string _tag, string _releaseCID, string _metaCID, uint256 _requestIndex) returns()
func (_Valist *ValistTransactorSession) ClearPendingRelease(_orgID [32]byte, _repoName string, _tag string, _releaseCID string, _metaCID string, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.Contract.ClearPendingRelease(&_Valist.TransactOpts, _orgID, _repoName, _tag, _releaseCID, _metaCID, _requestIndex)
}

// ClearPendingThreshold is a paid mutator transaction binding the contract method 0xb93d1685.
//
// Solidity: function clearPendingThreshold(bytes32 _orgID, string _repoName, uint256 _threshold, uint256 _requestIndex) returns()
func (_Valist *ValistTransactor) ClearPendingThreshold(opts *bind.TransactOpts, _orgID [32]byte, _repoName string, _threshold *big.Int, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "clearPendingThreshold", _orgID, _repoName, _threshold, _requestIndex)
}

// ClearPendingThreshold is a paid mutator transaction binding the contract method 0xb93d1685.
//
// Solidity: function clearPendingThreshold(bytes32 _orgID, string _repoName, uint256 _threshold, uint256 _requestIndex) returns()
func (_Valist *ValistSession) ClearPendingThreshold(_orgID [32]byte, _repoName string, _threshold *big.Int, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.Contract.ClearPendingThreshold(&_Valist.TransactOpts, _orgID, _repoName, _threshold, _requestIndex)
}

// ClearPendingThreshold is a paid mutator transaction binding the contract method 0xb93d1685.
//
// Solidity: function clearPendingThreshold(bytes32 _orgID, string _repoName, uint256 _threshold, uint256 _requestIndex) returns()
func (_Valist *ValistTransactorSession) ClearPendingThreshold(_orgID [32]byte, _repoName string, _threshold *big.Int, _requestIndex *big.Int) (*types.Transaction, error) {
	return _Valist.Contract.ClearPendingThreshold(&_Valist.TransactOpts, _orgID, _repoName, _threshold, _requestIndex)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0x6427acca.
//
// Solidity: function createOrganization(string _orgMeta) returns()
func (_Valist *ValistTransactor) CreateOrganization(opts *bind.TransactOpts, _orgMeta string) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "createOrganization", _orgMeta)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0x6427acca.
//
// Solidity: function createOrganization(string _orgMeta) returns()
func (_Valist *ValistSession) CreateOrganization(_orgMeta string) (*types.Transaction, error) {
	return _Valist.Contract.CreateOrganization(&_Valist.TransactOpts, _orgMeta)
}

// CreateOrganization is a paid mutator transaction binding the contract method 0x6427acca.
//
// Solidity: function createOrganization(string _orgMeta) returns()
func (_Valist *ValistTransactorSession) CreateOrganization(_orgMeta string) (*types.Transaction, error) {
	return _Valist.Contract.CreateOrganization(&_Valist.TransactOpts, _orgMeta)
}

// CreateRepository is a paid mutator transaction binding the contract method 0xe59dbfa4.
//
// Solidity: function createRepository(bytes32 _orgID, string _repoName, string _repoMeta) returns()
func (_Valist *ValistTransactor) CreateRepository(opts *bind.TransactOpts, _orgID [32]byte, _repoName string, _repoMeta string) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "createRepository", _orgID, _repoName, _repoMeta)
}

// CreateRepository is a paid mutator transaction binding the contract method 0xe59dbfa4.
//
// Solidity: function createRepository(bytes32 _orgID, string _repoName, string _repoMeta) returns()
func (_Valist *ValistSession) CreateRepository(_orgID [32]byte, _repoName string, _repoMeta string) (*types.Transaction, error) {
	return _Valist.Contract.CreateRepository(&_Valist.TransactOpts, _orgID, _repoName, _repoMeta)
}

// CreateRepository is a paid mutator transaction binding the contract method 0xe59dbfa4.
//
// Solidity: function createRepository(bytes32 _orgID, string _repoName, string _repoMeta) returns()
func (_Valist *ValistTransactorSession) CreateRepository(_orgID [32]byte, _repoName string, _repoMeta string) (*types.Transaction, error) {
	return _Valist.Contract.CreateRepository(&_Valist.TransactOpts, _orgID, _repoName, _repoMeta)
}

// SetOrgMeta is a paid mutator transaction binding the contract method 0xe253f5e1.
//
// Solidity: function setOrgMeta(bytes32 _orgID, string _metaCID) returns()
func (_Valist *ValistTransactor) SetOrgMeta(opts *bind.TransactOpts, _orgID [32]byte, _metaCID string) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "setOrgMeta", _orgID, _metaCID)
}

// SetOrgMeta is a paid mutator transaction binding the contract method 0xe253f5e1.
//
// Solidity: function setOrgMeta(bytes32 _orgID, string _metaCID) returns()
func (_Valist *ValistSession) SetOrgMeta(_orgID [32]byte, _metaCID string) (*types.Transaction, error) {
	return _Valist.Contract.SetOrgMeta(&_Valist.TransactOpts, _orgID, _metaCID)
}

// SetOrgMeta is a paid mutator transaction binding the contract method 0xe253f5e1.
//
// Solidity: function setOrgMeta(bytes32 _orgID, string _metaCID) returns()
func (_Valist *ValistTransactorSession) SetOrgMeta(_orgID [32]byte, _metaCID string) (*types.Transaction, error) {
	return _Valist.Contract.SetOrgMeta(&_Valist.TransactOpts, _orgID, _metaCID)
}

// SetRepoMeta is a paid mutator transaction binding the contract method 0xdedc3391.
//
// Solidity: function setRepoMeta(bytes32 _orgID, string _repoName, string _metaCID) returns()
func (_Valist *ValistTransactor) SetRepoMeta(opts *bind.TransactOpts, _orgID [32]byte, _repoName string, _metaCID string) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "setRepoMeta", _orgID, _repoName, _metaCID)
}

// SetRepoMeta is a paid mutator transaction binding the contract method 0xdedc3391.
//
// Solidity: function setRepoMeta(bytes32 _orgID, string _repoName, string _metaCID) returns()
func (_Valist *ValistSession) SetRepoMeta(_orgID [32]byte, _repoName string, _metaCID string) (*types.Transaction, error) {
	return _Valist.Contract.SetRepoMeta(&_Valist.TransactOpts, _orgID, _repoName, _metaCID)
}

// SetRepoMeta is a paid mutator transaction binding the contract method 0xdedc3391.
//
// Solidity: function setRepoMeta(bytes32 _orgID, string _repoName, string _metaCID) returns()
func (_Valist *ValistTransactorSession) SetRepoMeta(_orgID [32]byte, _repoName string, _metaCID string) (*types.Transaction, error) {
	return _Valist.Contract.SetRepoMeta(&_Valist.TransactOpts, _orgID, _repoName, _metaCID)
}

// VoteKey is a paid mutator transaction binding the contract method 0xf26f3c56.
//
// Solidity: function voteKey(bytes32 _orgID, string _repoName, bytes32 _operation, address _key) returns()
func (_Valist *ValistTransactor) VoteKey(opts *bind.TransactOpts, _orgID [32]byte, _repoName string, _operation [32]byte, _key common.Address) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "voteKey", _orgID, _repoName, _operation, _key)
}

// VoteKey is a paid mutator transaction binding the contract method 0xf26f3c56.
//
// Solidity: function voteKey(bytes32 _orgID, string _repoName, bytes32 _operation, address _key) returns()
func (_Valist *ValistSession) VoteKey(_orgID [32]byte, _repoName string, _operation [32]byte, _key common.Address) (*types.Transaction, error) {
	return _Valist.Contract.VoteKey(&_Valist.TransactOpts, _orgID, _repoName, _operation, _key)
}

// VoteKey is a paid mutator transaction binding the contract method 0xf26f3c56.
//
// Solidity: function voteKey(bytes32 _orgID, string _repoName, bytes32 _operation, address _key) returns()
func (_Valist *ValistTransactorSession) VoteKey(_orgID [32]byte, _repoName string, _operation [32]byte, _key common.Address) (*types.Transaction, error) {
	return _Valist.Contract.VoteKey(&_Valist.TransactOpts, _orgID, _repoName, _operation, _key)
}

// VoteRelease is a paid mutator transaction binding the contract method 0x328d3ddf.
//
// Solidity: function voteRelease(bytes32 _orgID, string _repoName, string _tag, string _releaseCID, string _metaCID) returns()
func (_Valist *ValistTransactor) VoteRelease(opts *bind.TransactOpts, _orgID [32]byte, _repoName string, _tag string, _releaseCID string, _metaCID string) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "voteRelease", _orgID, _repoName, _tag, _releaseCID, _metaCID)
}

// VoteRelease is a paid mutator transaction binding the contract method 0x328d3ddf.
//
// Solidity: function voteRelease(bytes32 _orgID, string _repoName, string _tag, string _releaseCID, string _metaCID) returns()
func (_Valist *ValistSession) VoteRelease(_orgID [32]byte, _repoName string, _tag string, _releaseCID string, _metaCID string) (*types.Transaction, error) {
	return _Valist.Contract.VoteRelease(&_Valist.TransactOpts, _orgID, _repoName, _tag, _releaseCID, _metaCID)
}

// VoteRelease is a paid mutator transaction binding the contract method 0x328d3ddf.
//
// Solidity: function voteRelease(bytes32 _orgID, string _repoName, string _tag, string _releaseCID, string _metaCID) returns()
func (_Valist *ValistTransactorSession) VoteRelease(_orgID [32]byte, _repoName string, _tag string, _releaseCID string, _metaCID string) (*types.Transaction, error) {
	return _Valist.Contract.VoteRelease(&_Valist.TransactOpts, _orgID, _repoName, _tag, _releaseCID, _metaCID)
}

// VoteThreshold is a paid mutator transaction binding the contract method 0xf735b352.
//
// Solidity: function voteThreshold(bytes32 _orgID, string _repoName, uint256 _threshold) returns()
func (_Valist *ValistTransactor) VoteThreshold(opts *bind.TransactOpts, _orgID [32]byte, _repoName string, _threshold *big.Int) (*types.Transaction, error) {
	return _Valist.contract.Transact(opts, "voteThreshold", _orgID, _repoName, _threshold)
}

// VoteThreshold is a paid mutator transaction binding the contract method 0xf735b352.
//
// Solidity: function voteThreshold(bytes32 _orgID, string _repoName, uint256 _threshold) returns()
func (_Valist *ValistSession) VoteThreshold(_orgID [32]byte, _repoName string, _threshold *big.Int) (*types.Transaction, error) {
	return _Valist.Contract.VoteThreshold(&_Valist.TransactOpts, _orgID, _repoName, _threshold)
}

// VoteThreshold is a paid mutator transaction binding the contract method 0xf735b352.
//
// Solidity: function voteThreshold(bytes32 _orgID, string _repoName, uint256 _threshold) returns()
func (_Valist *ValistTransactorSession) VoteThreshold(_orgID [32]byte, _repoName string, _threshold *big.Int) (*types.Transaction, error) {
	return _Valist.Contract.VoteThreshold(&_Valist.TransactOpts, _orgID, _repoName, _threshold)
}

// ValistMetaUpdateIterator is returned from FilterMetaUpdate and is used to iterate over the raw logs and unpacked data for MetaUpdate events raised by the Valist contract.
type ValistMetaUpdateIterator struct {
	Event *ValistMetaUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValistMetaUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValistMetaUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValistMetaUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValistMetaUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValistMetaUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValistMetaUpdate represents a MetaUpdate event raised by the Valist contract.
type ValistMetaUpdate struct {
	OrgID    [32]byte
	RepoName common.Hash
	Signer   common.Address
	MetaCID  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMetaUpdate is a free log retrieval operation binding the contract event 0x082dfc0cda6cc25674875e2dcd882c68a8b2c80bf48e58c6a3dc2384c6985955.
//
// Solidity: event MetaUpdate(bytes32 indexed _orgID, string indexed _repoName, address indexed _signer, string _metaCID)
func (_Valist *ValistFilterer) FilterMetaUpdate(opts *bind.FilterOpts, _orgID [][32]byte, _repoName []string, _signer []common.Address) (*ValistMetaUpdateIterator, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameRule []interface{}
	for _, _repoNameItem := range _repoName {
		_repoNameRule = append(_repoNameRule, _repoNameItem)
	}
	var _signerRule []interface{}
	for _, _signerItem := range _signer {
		_signerRule = append(_signerRule, _signerItem)
	}

	logs, sub, err := _Valist.contract.FilterLogs(opts, "MetaUpdate", _orgIDRule, _repoNameRule, _signerRule)
	if err != nil {
		return nil, err
	}
	return &ValistMetaUpdateIterator{contract: _Valist.contract, event: "MetaUpdate", logs: logs, sub: sub}, nil
}

// WatchMetaUpdate is a free log subscription operation binding the contract event 0x082dfc0cda6cc25674875e2dcd882c68a8b2c80bf48e58c6a3dc2384c6985955.
//
// Solidity: event MetaUpdate(bytes32 indexed _orgID, string indexed _repoName, address indexed _signer, string _metaCID)
func (_Valist *ValistFilterer) WatchMetaUpdate(opts *bind.WatchOpts, sink chan<- *ValistMetaUpdate, _orgID [][32]byte, _repoName []string, _signer []common.Address) (event.Subscription, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameRule []interface{}
	for _, _repoNameItem := range _repoName {
		_repoNameRule = append(_repoNameRule, _repoNameItem)
	}
	var _signerRule []interface{}
	for _, _signerItem := range _signer {
		_signerRule = append(_signerRule, _signerItem)
	}

	logs, sub, err := _Valist.contract.WatchLogs(opts, "MetaUpdate", _orgIDRule, _repoNameRule, _signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValistMetaUpdate)
				if err := _Valist.contract.UnpackLog(event, "MetaUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetaUpdate is a log parse operation binding the contract event 0x082dfc0cda6cc25674875e2dcd882c68a8b2c80bf48e58c6a3dc2384c6985955.
//
// Solidity: event MetaUpdate(bytes32 indexed _orgID, string indexed _repoName, address indexed _signer, string _metaCID)
func (_Valist *ValistFilterer) ParseMetaUpdate(log types.Log) (*ValistMetaUpdate, error) {
	event := new(ValistMetaUpdate)
	if err := _Valist.contract.UnpackLog(event, "MetaUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValistOrgCreatedIterator is returned from FilterOrgCreated and is used to iterate over the raw logs and unpacked data for OrgCreated events raised by the Valist contract.
type ValistOrgCreatedIterator struct {
	Event *ValistOrgCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValistOrgCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValistOrgCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValistOrgCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValistOrgCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValistOrgCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValistOrgCreated represents a OrgCreated event raised by the Valist contract.
type ValistOrgCreated struct {
	OrgID       [32]byte
	MetaCIDHash common.Hash
	MetaCID     string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrgCreated is a free log retrieval operation binding the contract event 0xc5eb86c0b2c1ce6abdc8dea996a5aa6cf196b33ee7a2c140ce4f04f2fbb3baab.
//
// Solidity: event OrgCreated(bytes32 indexed _orgID, string indexed _metaCIDHash, string _metaCID)
func (_Valist *ValistFilterer) FilterOrgCreated(opts *bind.FilterOpts, _orgID [][32]byte, _metaCIDHash []string) (*ValistOrgCreatedIterator, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _metaCIDHashRule []interface{}
	for _, _metaCIDHashItem := range _metaCIDHash {
		_metaCIDHashRule = append(_metaCIDHashRule, _metaCIDHashItem)
	}

	logs, sub, err := _Valist.contract.FilterLogs(opts, "OrgCreated", _orgIDRule, _metaCIDHashRule)
	if err != nil {
		return nil, err
	}
	return &ValistOrgCreatedIterator{contract: _Valist.contract, event: "OrgCreated", logs: logs, sub: sub}, nil
}

// WatchOrgCreated is a free log subscription operation binding the contract event 0xc5eb86c0b2c1ce6abdc8dea996a5aa6cf196b33ee7a2c140ce4f04f2fbb3baab.
//
// Solidity: event OrgCreated(bytes32 indexed _orgID, string indexed _metaCIDHash, string _metaCID)
func (_Valist *ValistFilterer) WatchOrgCreated(opts *bind.WatchOpts, sink chan<- *ValistOrgCreated, _orgID [][32]byte, _metaCIDHash []string) (event.Subscription, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _metaCIDHashRule []interface{}
	for _, _metaCIDHashItem := range _metaCIDHash {
		_metaCIDHashRule = append(_metaCIDHashRule, _metaCIDHashItem)
	}

	logs, sub, err := _Valist.contract.WatchLogs(opts, "OrgCreated", _orgIDRule, _metaCIDHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValistOrgCreated)
				if err := _Valist.contract.UnpackLog(event, "OrgCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrgCreated is a log parse operation binding the contract event 0xc5eb86c0b2c1ce6abdc8dea996a5aa6cf196b33ee7a2c140ce4f04f2fbb3baab.
//
// Solidity: event OrgCreated(bytes32 indexed _orgID, string indexed _metaCIDHash, string _metaCID)
func (_Valist *ValistFilterer) ParseOrgCreated(log types.Log) (*ValistOrgCreated, error) {
	event := new(ValistOrgCreated)
	if err := _Valist.contract.UnpackLog(event, "OrgCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValistRepoCreatedIterator is returned from FilterRepoCreated and is used to iterate over the raw logs and unpacked data for RepoCreated events raised by the Valist contract.
type ValistRepoCreatedIterator struct {
	Event *ValistRepoCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValistRepoCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValistRepoCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValistRepoCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValistRepoCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValistRepoCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValistRepoCreated represents a RepoCreated event raised by the Valist contract.
type ValistRepoCreated struct {
	OrgID        [32]byte
	RepoNameHash common.Hash
	RepoName     string
	MetaCIDHash  common.Hash
	MetaCID      string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRepoCreated is a free log retrieval operation binding the contract event 0x50b56e7c402d556bc61b6fc9bba647b83c4590f2fca5a6d463450e78f3a2d44c.
//
// Solidity: event RepoCreated(bytes32 indexed _orgID, string indexed _repoNameHash, string _repoName, string indexed _metaCIDHash, string _metaCID)
func (_Valist *ValistFilterer) FilterRepoCreated(opts *bind.FilterOpts, _orgID [][32]byte, _repoNameHash []string, _metaCIDHash []string) (*ValistRepoCreatedIterator, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameHashRule []interface{}
	for _, _repoNameHashItem := range _repoNameHash {
		_repoNameHashRule = append(_repoNameHashRule, _repoNameHashItem)
	}

	var _metaCIDHashRule []interface{}
	for _, _metaCIDHashItem := range _metaCIDHash {
		_metaCIDHashRule = append(_metaCIDHashRule, _metaCIDHashItem)
	}

	logs, sub, err := _Valist.contract.FilterLogs(opts, "RepoCreated", _orgIDRule, _repoNameHashRule, _metaCIDHashRule)
	if err != nil {
		return nil, err
	}
	return &ValistRepoCreatedIterator{contract: _Valist.contract, event: "RepoCreated", logs: logs, sub: sub}, nil
}

// WatchRepoCreated is a free log subscription operation binding the contract event 0x50b56e7c402d556bc61b6fc9bba647b83c4590f2fca5a6d463450e78f3a2d44c.
//
// Solidity: event RepoCreated(bytes32 indexed _orgID, string indexed _repoNameHash, string _repoName, string indexed _metaCIDHash, string _metaCID)
func (_Valist *ValistFilterer) WatchRepoCreated(opts *bind.WatchOpts, sink chan<- *ValistRepoCreated, _orgID [][32]byte, _repoNameHash []string, _metaCIDHash []string) (event.Subscription, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameHashRule []interface{}
	for _, _repoNameHashItem := range _repoNameHash {
		_repoNameHashRule = append(_repoNameHashRule, _repoNameHashItem)
	}

	var _metaCIDHashRule []interface{}
	for _, _metaCIDHashItem := range _metaCIDHash {
		_metaCIDHashRule = append(_metaCIDHashRule, _metaCIDHashItem)
	}

	logs, sub, err := _Valist.contract.WatchLogs(opts, "RepoCreated", _orgIDRule, _repoNameHashRule, _metaCIDHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValistRepoCreated)
				if err := _Valist.contract.UnpackLog(event, "RepoCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepoCreated is a log parse operation binding the contract event 0x50b56e7c402d556bc61b6fc9bba647b83c4590f2fca5a6d463450e78f3a2d44c.
//
// Solidity: event RepoCreated(bytes32 indexed _orgID, string indexed _repoNameHash, string _repoName, string indexed _metaCIDHash, string _metaCID)
func (_Valist *ValistFilterer) ParseRepoCreated(log types.Log) (*ValistRepoCreated, error) {
	event := new(ValistRepoCreated)
	if err := _Valist.contract.UnpackLog(event, "RepoCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValistVoteKeyEventIterator is returned from FilterVoteKeyEvent and is used to iterate over the raw logs and unpacked data for VoteKeyEvent events raised by the Valist contract.
type ValistVoteKeyEventIterator struct {
	Event *ValistVoteKeyEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValistVoteKeyEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValistVoteKeyEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValistVoteKeyEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValistVoteKeyEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValistVoteKeyEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValistVoteKeyEvent represents a VoteKeyEvent event raised by the Valist contract.
type ValistVoteKeyEvent struct {
	OrgID     [32]byte
	RepoName  common.Hash
	Signer    common.Address
	Operation [32]byte
	Key       common.Address
	SigCount  *big.Int
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteKeyEvent is a free log retrieval operation binding the contract event 0xa8cea71b77054741d93ac504e0ee90fa3e815d68104468d65f9eb36924a8d590.
//
// Solidity: event VoteKeyEvent(bytes32 indexed _orgID, string indexed _repoName, address _signer, bytes32 _operation, address indexed _key, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) FilterVoteKeyEvent(opts *bind.FilterOpts, _orgID [][32]byte, _repoName []string, _key []common.Address) (*ValistVoteKeyEventIterator, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameRule []interface{}
	for _, _repoNameItem := range _repoName {
		_repoNameRule = append(_repoNameRule, _repoNameItem)
	}

	var _keyRule []interface{}
	for _, _keyItem := range _key {
		_keyRule = append(_keyRule, _keyItem)
	}

	logs, sub, err := _Valist.contract.FilterLogs(opts, "VoteKeyEvent", _orgIDRule, _repoNameRule, _keyRule)
	if err != nil {
		return nil, err
	}
	return &ValistVoteKeyEventIterator{contract: _Valist.contract, event: "VoteKeyEvent", logs: logs, sub: sub}, nil
}

// WatchVoteKeyEvent is a free log subscription operation binding the contract event 0xa8cea71b77054741d93ac504e0ee90fa3e815d68104468d65f9eb36924a8d590.
//
// Solidity: event VoteKeyEvent(bytes32 indexed _orgID, string indexed _repoName, address _signer, bytes32 _operation, address indexed _key, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) WatchVoteKeyEvent(opts *bind.WatchOpts, sink chan<- *ValistVoteKeyEvent, _orgID [][32]byte, _repoName []string, _key []common.Address) (event.Subscription, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameRule []interface{}
	for _, _repoNameItem := range _repoName {
		_repoNameRule = append(_repoNameRule, _repoNameItem)
	}

	var _keyRule []interface{}
	for _, _keyItem := range _key {
		_keyRule = append(_keyRule, _keyItem)
	}

	logs, sub, err := _Valist.contract.WatchLogs(opts, "VoteKeyEvent", _orgIDRule, _repoNameRule, _keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValistVoteKeyEvent)
				if err := _Valist.contract.UnpackLog(event, "VoteKeyEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteKeyEvent is a log parse operation binding the contract event 0xa8cea71b77054741d93ac504e0ee90fa3e815d68104468d65f9eb36924a8d590.
//
// Solidity: event VoteKeyEvent(bytes32 indexed _orgID, string indexed _repoName, address _signer, bytes32 _operation, address indexed _key, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) ParseVoteKeyEvent(log types.Log) (*ValistVoteKeyEvent, error) {
	event := new(ValistVoteKeyEvent)
	if err := _Valist.contract.UnpackLog(event, "VoteKeyEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValistVoteReleaseEventIterator is returned from FilterVoteReleaseEvent and is used to iterate over the raw logs and unpacked data for VoteReleaseEvent events raised by the Valist contract.
type ValistVoteReleaseEventIterator struct {
	Event *ValistVoteReleaseEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValistVoteReleaseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValistVoteReleaseEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValistVoteReleaseEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValistVoteReleaseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValistVoteReleaseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValistVoteReleaseEvent represents a VoteReleaseEvent event raised by the Valist contract.
type ValistVoteReleaseEvent struct {
	OrgID      [32]byte
	RepoName   common.Hash
	Tag        common.Hash
	ReleaseCID string
	MetaCID    string
	Signer     common.Address
	SigCount   *big.Int
	Threshold  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteReleaseEvent is a free log retrieval operation binding the contract event 0x38b17387282322f8d6de03e5d4b3ff512d3bb4d3db8cc2be4611ecfb8a126c6b.
//
// Solidity: event VoteReleaseEvent(bytes32 indexed _orgID, string indexed _repoName, string indexed _tag, string _releaseCID, string _metaCID, address _signer, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) FilterVoteReleaseEvent(opts *bind.FilterOpts, _orgID [][32]byte, _repoName []string, _tag []string) (*ValistVoteReleaseEventIterator, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameRule []interface{}
	for _, _repoNameItem := range _repoName {
		_repoNameRule = append(_repoNameRule, _repoNameItem)
	}
	var _tagRule []interface{}
	for _, _tagItem := range _tag {
		_tagRule = append(_tagRule, _tagItem)
	}

	logs, sub, err := _Valist.contract.FilterLogs(opts, "VoteReleaseEvent", _orgIDRule, _repoNameRule, _tagRule)
	if err != nil {
		return nil, err
	}
	return &ValistVoteReleaseEventIterator{contract: _Valist.contract, event: "VoteReleaseEvent", logs: logs, sub: sub}, nil
}

// WatchVoteReleaseEvent is a free log subscription operation binding the contract event 0x38b17387282322f8d6de03e5d4b3ff512d3bb4d3db8cc2be4611ecfb8a126c6b.
//
// Solidity: event VoteReleaseEvent(bytes32 indexed _orgID, string indexed _repoName, string indexed _tag, string _releaseCID, string _metaCID, address _signer, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) WatchVoteReleaseEvent(opts *bind.WatchOpts, sink chan<- *ValistVoteReleaseEvent, _orgID [][32]byte, _repoName []string, _tag []string) (event.Subscription, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameRule []interface{}
	for _, _repoNameItem := range _repoName {
		_repoNameRule = append(_repoNameRule, _repoNameItem)
	}
	var _tagRule []interface{}
	for _, _tagItem := range _tag {
		_tagRule = append(_tagRule, _tagItem)
	}

	logs, sub, err := _Valist.contract.WatchLogs(opts, "VoteReleaseEvent", _orgIDRule, _repoNameRule, _tagRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValistVoteReleaseEvent)
				if err := _Valist.contract.UnpackLog(event, "VoteReleaseEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteReleaseEvent is a log parse operation binding the contract event 0x38b17387282322f8d6de03e5d4b3ff512d3bb4d3db8cc2be4611ecfb8a126c6b.
//
// Solidity: event VoteReleaseEvent(bytes32 indexed _orgID, string indexed _repoName, string indexed _tag, string _releaseCID, string _metaCID, address _signer, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) ParseVoteReleaseEvent(log types.Log) (*ValistVoteReleaseEvent, error) {
	event := new(ValistVoteReleaseEvent)
	if err := _Valist.contract.UnpackLog(event, "VoteReleaseEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValistVoteThresholdEventIterator is returned from FilterVoteThresholdEvent and is used to iterate over the raw logs and unpacked data for VoteThresholdEvent events raised by the Valist contract.
type ValistVoteThresholdEventIterator struct {
	Event *ValistVoteThresholdEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValistVoteThresholdEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValistVoteThresholdEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValistVoteThresholdEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValistVoteThresholdEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValistVoteThresholdEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValistVoteThresholdEvent represents a VoteThresholdEvent event raised by the Valist contract.
type ValistVoteThresholdEvent struct {
	OrgID            [32]byte
	RepoName         common.Hash
	Signer           common.Address
	PendingThreshold *big.Int
	SigCount         *big.Int
	Threshold        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVoteThresholdEvent is a free log retrieval operation binding the contract event 0x2daad345bbc18b1fa8a7f542403081df42c546dd841a44507d53e174e761ce89.
//
// Solidity: event VoteThresholdEvent(bytes32 indexed _orgID, string indexed _repoName, address _signer, uint256 indexed _pendingThreshold, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) FilterVoteThresholdEvent(opts *bind.FilterOpts, _orgID [][32]byte, _repoName []string, _pendingThreshold []*big.Int) (*ValistVoteThresholdEventIterator, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameRule []interface{}
	for _, _repoNameItem := range _repoName {
		_repoNameRule = append(_repoNameRule, _repoNameItem)
	}

	var _pendingThresholdRule []interface{}
	for _, _pendingThresholdItem := range _pendingThreshold {
		_pendingThresholdRule = append(_pendingThresholdRule, _pendingThresholdItem)
	}

	logs, sub, err := _Valist.contract.FilterLogs(opts, "VoteThresholdEvent", _orgIDRule, _repoNameRule, _pendingThresholdRule)
	if err != nil {
		return nil, err
	}
	return &ValistVoteThresholdEventIterator{contract: _Valist.contract, event: "VoteThresholdEvent", logs: logs, sub: sub}, nil
}

// WatchVoteThresholdEvent is a free log subscription operation binding the contract event 0x2daad345bbc18b1fa8a7f542403081df42c546dd841a44507d53e174e761ce89.
//
// Solidity: event VoteThresholdEvent(bytes32 indexed _orgID, string indexed _repoName, address _signer, uint256 indexed _pendingThreshold, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) WatchVoteThresholdEvent(opts *bind.WatchOpts, sink chan<- *ValistVoteThresholdEvent, _orgID [][32]byte, _repoName []string, _pendingThreshold []*big.Int) (event.Subscription, error) {

	var _orgIDRule []interface{}
	for _, _orgIDItem := range _orgID {
		_orgIDRule = append(_orgIDRule, _orgIDItem)
	}
	var _repoNameRule []interface{}
	for _, _repoNameItem := range _repoName {
		_repoNameRule = append(_repoNameRule, _repoNameItem)
	}

	var _pendingThresholdRule []interface{}
	for _, _pendingThresholdItem := range _pendingThreshold {
		_pendingThresholdRule = append(_pendingThresholdRule, _pendingThresholdItem)
	}

	logs, sub, err := _Valist.contract.WatchLogs(opts, "VoteThresholdEvent", _orgIDRule, _repoNameRule, _pendingThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValistVoteThresholdEvent)
				if err := _Valist.contract.UnpackLog(event, "VoteThresholdEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteThresholdEvent is a log parse operation binding the contract event 0x2daad345bbc18b1fa8a7f542403081df42c546dd841a44507d53e174e761ce89.
//
// Solidity: event VoteThresholdEvent(bytes32 indexed _orgID, string indexed _repoName, address _signer, uint256 indexed _pendingThreshold, uint256 _sigCount, uint256 _threshold)
func (_Valist *ValistFilterer) ParseVoteThresholdEvent(log types.Log) (*ValistVoteThresholdEvent, error) {
	event := new(ValistVoteThresholdEvent)
	if err := _Valist.contract.UnpackLog(event, "VoteThresholdEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
