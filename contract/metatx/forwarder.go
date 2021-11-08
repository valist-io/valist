// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package metatx

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

// ERC20ForwardRequestTypesERC20ForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type ERC20ForwardRequestTypesERC20ForwardRequest struct {
	From          common.Address
	To            common.Address
	Token         common.Address
	TxGas         *big.Int
	TokenGasPrice *big.Int
	BatchId       *big.Int
	BatchNonce    *big.Int
	Deadline      *big.Int
	Data          []byte
}

// MetatxABI is the input ABI used to generate the binding from.
const MetatxABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"domainValue\",\"type\":\"bytes\"}],\"name\":\"DomainRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REQUEST_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"domains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structERC20ForwardRequestTypes.ERC20ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"executeEIP712\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structERC20ForwardRequestTypes.ERC20ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"executePersonalSign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"ret\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"registerDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structERC20ForwardRequestTypes.ERC20ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifyEIP712\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structERC20ForwardRequestTypes.ERC20ForwardRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"verifyPersonalSign\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MetatxBin is the compiled bytecode used for deploying new contracts.
var MetatxBin = "0x60806040523480156200001157600080fd5b5060405162002373380380620023738339818101604052810190620000379190620001bb565b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620000dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4f776e657220416464726573732063616e6e6f7420626520300000000000000081525060200191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600046905080600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200019c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001939062000229565b60405180910390fd5b5050620002aa565b600081519050620001b58162000290565b92915050565b600060208284031215620001ce57600080fd5b6000620001de84828501620001a4565b91505092915050565b6000620001f66019836200024b565b91507f4f776e657220416464726573732063616e6e6f742062652030000000000000006000830152602082019050919050565b600060208201905081810360008301526200024481620001e7565b9050919050565b600082825260208201905092915050565b6000620002698262000270565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200029b816200025c565b8114620002a757600080fd5b50565b6120b980620002ba6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638f32d59b1161008c578063a41a03f211610066578063a41a03f214610202578063c3f28abd1461021e578063c722f1771461023c578063f2fde38b1461026c576100cf565b80638f32d59b146101aa5780639c7b4592146101c85780639e39b73e146101e4576100cf565b806341706c4e146100d45780636e4cb07514610105578063715018a6146101215780638171e6321461012b578063895358031461015c5780638da5cb5b1461018c575b600080fd5b6100ee60048036038101906100e9919061158e565b610288565b6040516100fc929190611b01565b60405180910390f35b61011f600480360381019061011a9190611612565b6103f3565b005b610129610446565b005b61014560048036038101906101409190611612565b610561565b604051610153929190611b01565b60405180910390f35b610176600480360381019061017191906114b4565b6106ca565b6040516101839190611cff565b60405180910390f35b610194610725565b6040516101a19190611acb565b60405180910390f35b6101b261074e565b6040516101bf9190611ae6565b60405180910390f35b6101e260048036038101906101dd9190611519565b6107a5565b005b6101ec6108fb565b6040516101f99190611b31565b60405180910390f35b61021c6004803603810190610217919061158e565b61091e565b005b610226610973565b6040516102339190611c5d565b60405180910390f35b610256600480360381019061025191906114f0565b61098f565b6040516102639190611ae6565b60405180910390f35b6102866004803603810190610281919061148b565b6109af565b005b600060606102db868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610a18565b6102e486610d09565b8560200160208101906102f7919061148b565b73ffffffffffffffffffffffffffffffffffffffff168660600135878061010001906103239190611d1a565b896000016020810190610336919061148b565b60405160200161034893929190611a2d565b6040516020818303038152906040526040516103649190611a57565b60006040518083038160008787f1925050503d80600081146103a2576040519150601f19603f3d011682016040523d82523d6000602084013e6103a7565b606091505b508092508193505050603f8660600135816103be57fe5b045a116103c757fe5b6103ea82826040518060600160405280602d8152602001612057602d9139610d82565b94509492505050565b6104418383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610ddc565b505050565b61044e61074e565b6104a3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526038815260200180611f826038913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600060606105b38585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610ddc565b6105bc85610d09565b8460200160208101906105cf919061148b565b73ffffffffffffffffffffffffffffffffffffffff168560600135868061010001906105fb9190611d1a565b88600001602081019061060e919061148b565b60405160200161062093929190611a2d565b60405160208183030381529060405260405161063c9190611a57565b60006040518083038160008787f1925050503d806000811461067a576040519150601f19603f3d011682016040523d82523d6000602084013e61067f565b606091505b508092508193505050603f85606001358161069657fe5b045a1161069f57fe5b6106c282826040518060600160405280602d8152602001612057602d9139610d82565b935093915050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600083815260200190815260200160002054905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6107ad61074e565b610802576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526038815260200180611f826038913960400191505060405180910390fd5b600046905060006040518060800160405280604f8152602001611eef604f9139805190602001208686604051610839929190611a14565b60405180910390208585604051610851929190611a14565b6040518091039020308560001b604051602001610872959493929190611be8565b6040516020818303038152906040529050600081805190602001209050600180600083815260200190815260200160002060006101000a81548160ff021916908315150217905550807f4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8836040516108ea9190611c3b565b60405180910390a250505050505050565b6040518060c00160405280609d8152602001611fba609d91398051906020012081565b61096d848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610a18565b50505050565b6040518060800160405280604f8152602001611eef604f913981565b60016020528060005260406000206000915054906101000a900460ff1681565b6109b761074e565b610a0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526038815260200180611f826038913960400191505060405180910390fd5b610a1581610fe1565b50565b600046905060008460e001351480610a3757508360e001356014420111155b610a76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6d90611c7f565b60405180910390fd5b6001600084815260200190815260200160002060009054906101000a900460ff16610ad6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610acd90611c9f565b60405180910390fd5b8060025414610b1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1190611cbf565b60405180910390fd5b6000836040518060c00160405280609d8152602001611fba609d913980519060200120866000016020810190610b50919061148b565b876020016020810190610b63919061148b565b886040016020810190610b76919061148b565b89606001358a608001358b60a00135600360008e6000016020810190610b9c919061148b565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008e60a001358152602001908152602001600020548d60e001358e806101000190610c029190611d1a565b604051610c10929190611a14565b6040518091039020604051602001610c319a99989796959493929190611b4c565b60405160208183030381529060405280519060200120604051602001610c58929190611a94565b604051602081830303815290604052805190602001209050846000016020810190610c83919061148b565b73ffffffffffffffffffffffffffffffffffffffff16610cac84836110d890919063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1614610d02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf990611cdf565b60405180910390fd5b5050505050565b60036000826000016020810190610d20919061148b565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008260a0013581526020019081526020016000206000815480929190600101919050555050565b82610dd757600082511115610d9a5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dce9190611c5d565b60405180910390fd5b505050565b60008260e001351480610df657508160e001356014420111155b610e35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2c90611c7f565b60405180910390fd5b6000610f48836000016020810190610e4d919061148b565b846020016020810190610e60919061148b565b856040016020810190610e73919061148b565b866060013587608001358860a00135600360008b6000016020810190610e99919061148b565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b60a001358152602001908152602001600020548a60e001358b806101000190610eff9190611d1a565b604051610f0d929190611a14565b6040518091039020604051602001610f2d99989796959493929190611971565b6040516020818303038152906040528051906020012061136c565b9050826000016020810190610f5d919061148b565b73ffffffffffffffffffffffffffffffffffffffff16610f8683836110d890919063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1614610fdc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fd390611cdf565b60405180910390fd5b505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561101b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006041825114611151576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45434453413a20696e76616c6964207369676e6174757265206c656e6774680081525060200191505060405180910390fd5b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11156111ea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180611f3e6022913960400191505060405180910390fd5b601b8160ff1614806111ff5750601c8160ff16145b611254576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180611f606022913960400191505060405180910390fd5b600060018783868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156112b0573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561135f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f45434453413a20696e76616c6964207369676e6174757265000000000000000081525060200191505060405180910390fd5b8094505050505092915050565b60008160405160200161137f9190611a6e565b604051602081830303815290604052805190602001209050919050565b6000813590506113ab81611ea9565b92915050565b6000813590506113c081611ec0565b92915050565b60008083601f8401126113d857600080fd5b8235905067ffffffffffffffff8111156113f157600080fd5b60208301915083600182028301111561140957600080fd5b9250929050565b60008083601f84011261142257600080fd5b8235905067ffffffffffffffff81111561143b57600080fd5b60208301915083600182028301111561145357600080fd5b9250929050565b6000610120828403121561146d57600080fd5b81905092915050565b60008135905061148581611ed7565b92915050565b60006020828403121561149d57600080fd5b60006114ab8482850161139c565b91505092915050565b600080604083850312156114c757600080fd5b60006114d58582860161139c565b92505060206114e685828601611476565b9150509250929050565b60006020828403121561150257600080fd5b6000611510848285016113b1565b91505092915050565b6000806000806040858703121561152f57600080fd5b600085013567ffffffffffffffff81111561154957600080fd5b61155587828801611410565b9450945050602085013567ffffffffffffffff81111561157457600080fd5b61158087828801611410565b925092505092959194509250565b600080600080606085870312156115a457600080fd5b600085013567ffffffffffffffff8111156115be57600080fd5b6115ca8782880161145a565b94505060206115db878288016113b1565b935050604085013567ffffffffffffffff8111156115f857600080fd5b611604878288016113c6565b925092505092959194509250565b60008060006040848603121561162757600080fd5b600084013567ffffffffffffffff81111561164157600080fd5b61164d8682870161145a565b935050602084013567ffffffffffffffff81111561166a57600080fd5b611676868287016113c6565b92509250509250925092565b61168b81611dbf565b82525050565b6116a261169d82611dbf565b611e53565b82525050565b6116b181611dd1565b82525050565b6116c081611ddd565b82525050565b6116d76116d282611ddd565b611e65565b82525050565b60006116e98385611d98565b93506116f6838584611e11565b82840190509392505050565b600061170d82611d71565b6117178185611d87565b9350611727818560208601611e20565b61173081611e8b565b840191505092915050565b600061174682611d71565b6117508185611d98565b9350611760818560208601611e20565b80840191505092915050565b600061177782611d7c565b6117818185611da3565b9350611791818560208601611e20565b61179a81611e8b565b840191505092915050565b60006117b2600f83611da3565b91507f72657175657374206578706972656400000000000000000000000000000000006000830152602082019050919050565b60006117f2601c83611db4565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b6000611832600283611db4565b91507f19010000000000000000000000000000000000000000000000000000000000006000830152600282019050919050565b6000611872601d83611da3565b91507f756e7265676973746572656420646f6d61696e20736570617261746f720000006000830152602082019050919050565b60006118b2602383611da3565b91507f706f74656e7469616c207265706c61792061747461636b206f6e20746865206660008301527f6f726b00000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611918601283611da3565b91507f7369676e6174757265206d69736d6174636800000000000000000000000000006000830152602082019050919050565b61195481611e07565b82525050565b61196b61196682611e07565b611e81565b82525050565b600061197d828c611691565b60148201915061198d828b611691565b60148201915061199d828a611691565b6014820191506119ad828961195a565b6020820191506119bd828861195a565b6020820191506119cd828761195a565b6020820191506119dd828661195a565b6020820191506119ed828561195a565b6020820191506119fd82846116c6565b6020820191508190509a9950505050505050505050565b6000611a218284866116dd565b91508190509392505050565b6000611a3a8285876116dd565b9150611a468284611691565b601482019150819050949350505050565b6000611a63828461173b565b915081905092915050565b6000611a79826117e5565b9150611a8582846116c6565b60208201915081905092915050565b6000611a9f82611825565b9150611aab82856116c6565b602082019150611abb82846116c6565b6020820191508190509392505050565b6000602082019050611ae06000830184611682565b92915050565b6000602082019050611afb60008301846116a8565b92915050565b6000604082019050611b1660008301856116a8565b8181036020830152611b288184611702565b90509392505050565b6000602082019050611b4660008301846116b7565b92915050565b600061014082019050611b62600083018d6116b7565b611b6f602083018c611682565b611b7c604083018b611682565b611b89606083018a611682565b611b96608083018961194b565b611ba360a083018861194b565b611bb060c083018761194b565b611bbd60e083018661194b565b611bcb61010083018561194b565b611bd96101208301846116b7565b9b9a5050505050505050505050565b600060a082019050611bfd60008301886116b7565b611c0a60208301876116b7565b611c1760408301866116b7565b611c246060830185611682565b611c3160808301846116b7565b9695505050505050565b60006020820190508181036000830152611c558184611702565b905092915050565b60006020820190508181036000830152611c77818461176c565b905092915050565b60006020820190508181036000830152611c98816117a5565b9050919050565b60006020820190508181036000830152611cb881611865565b9050919050565b60006020820190508181036000830152611cd8816118a5565b9050919050565b60006020820190508181036000830152611cf88161190b565b9050919050565b6000602082019050611d14600083018461194b565b92915050565b60008083356001602003843603038112611d3357600080fd5b80840192508235915067ffffffffffffffff821115611d5157600080fd5b602083019250600182023603831315611d6957600080fd5b509250929050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b6000611dca82611de7565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611e3e578082015181840152602081019050611e23565b83811115611e4d576000848401525b50505050565b6000611e5e82611e6f565b9050919050565b6000819050919050565b6000611e7a82611e9c565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b611eb281611dbf565b8114611ebd57600080fd5b50565b611ec981611ddd565b8114611ed457600080fd5b50565b611ee081611e07565b8114611eeb57600080fd5b5056fe454950373132446f6d61696e28737472696e67206e616d652c737472696e672076657273696f6e2c6164647265737320766572696679696e67436f6e74726163742c627974657333322073616c742945434453413a20696e76616c6964207369676e6174757265202773272076616c756545434453413a20696e76616c6964207369676e6174757265202776272076616c75654f6e6c7920636f6e7472616374206f776e657220697320616c6c6f77656420746f20706572666f726d2074686973206f7065726174696f6e4552433230466f72776172645265717565737428616464726573732066726f6d2c6164647265737320746f2c6164647265737320746f6b656e2c75696e743235362074784761732c75696e7432353620746f6b656e47617350726963652c75696e7432353620626174636849642c75696e743235362062617463684e6f6e63652c75696e7432353620646561646c696e652c6279746573206461746129466f727761726465642063616c6c20746f2064657374696e6174696f6e20646964206e6f742073756363656564a2646970667358221220cbeab4f57eb4658c06c16e975fb029ffb5fe612ec4028b0a7ed2aeba4019e01b64736f6c63430007060033"

// DeployMetatx deploys a new Ethereum contract, binding an instance of Metatx to it.
func DeployMetatx(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *Metatx, error) {
	parsed, err := abi.JSON(strings.NewReader(MetatxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MetatxBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Metatx{MetatxCaller: MetatxCaller{contract: contract}, MetatxTransactor: MetatxTransactor{contract: contract}, MetatxFilterer: MetatxFilterer{contract: contract}}, nil
}

// Metatx is an auto generated Go binding around an Ethereum contract.
type Metatx struct {
	MetatxCaller     // Read-only binding to the contract
	MetatxTransactor // Write-only binding to the contract
	MetatxFilterer   // Log filterer for contract events
}

// MetatxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetatxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetatxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetatxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetatxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetatxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetatxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetatxSession struct {
	Contract     *Metatx           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetatxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetatxCallerSession struct {
	Contract *MetatxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MetatxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetatxTransactorSession struct {
	Contract     *MetatxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetatxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetatxRaw struct {
	Contract *Metatx // Generic contract binding to access the raw methods on
}

// MetatxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetatxCallerRaw struct {
	Contract *MetatxCaller // Generic read-only contract binding to access the raw methods on
}

// MetatxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetatxTransactorRaw struct {
	Contract *MetatxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetatx creates a new instance of Metatx, bound to a specific deployed contract.
func NewMetatx(address common.Address, backend bind.ContractBackend) (*Metatx, error) {
	contract, err := bindMetatx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Metatx{MetatxCaller: MetatxCaller{contract: contract}, MetatxTransactor: MetatxTransactor{contract: contract}, MetatxFilterer: MetatxFilterer{contract: contract}}, nil
}

// NewMetatxCaller creates a new read-only instance of Metatx, bound to a specific deployed contract.
func NewMetatxCaller(address common.Address, caller bind.ContractCaller) (*MetatxCaller, error) {
	contract, err := bindMetatx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetatxCaller{contract: contract}, nil
}

// NewMetatxTransactor creates a new write-only instance of Metatx, bound to a specific deployed contract.
func NewMetatxTransactor(address common.Address, transactor bind.ContractTransactor) (*MetatxTransactor, error) {
	contract, err := bindMetatx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetatxTransactor{contract: contract}, nil
}

// NewMetatxFilterer creates a new log filterer instance of Metatx, bound to a specific deployed contract.
func NewMetatxFilterer(address common.Address, filterer bind.ContractFilterer) (*MetatxFilterer, error) {
	contract, err := bindMetatx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetatxFilterer{contract: contract}, nil
}

// bindMetatx binds a generic wrapper to an already deployed contract.
func bindMetatx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetatxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Metatx *MetatxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Metatx.Contract.MetatxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Metatx *MetatxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metatx.Contract.MetatxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Metatx *MetatxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Metatx.Contract.MetatxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Metatx *MetatxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Metatx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Metatx *MetatxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metatx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Metatx *MetatxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Metatx.Contract.contract.Transact(opts, method, params...)
}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Metatx *MetatxCaller) EIP712DOMAINTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Metatx.contract.Call(opts, &out, "EIP712_DOMAIN_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Metatx *MetatxSession) EIP712DOMAINTYPE() (string, error) {
	return _Metatx.Contract.EIP712DOMAINTYPE(&_Metatx.CallOpts)
}

// EIP712DOMAINTYPE is a free data retrieval call binding the contract method 0xc3f28abd.
//
// Solidity: function EIP712_DOMAIN_TYPE() view returns(string)
func (_Metatx *MetatxCallerSession) EIP712DOMAINTYPE() (string, error) {
	return _Metatx.Contract.EIP712DOMAINTYPE(&_Metatx.CallOpts)
}

// REQUESTTYPEHASH is a free data retrieval call binding the contract method 0x9e39b73e.
//
// Solidity: function REQUEST_TYPEHASH() view returns(bytes32)
func (_Metatx *MetatxCaller) REQUESTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Metatx.contract.Call(opts, &out, "REQUEST_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REQUESTTYPEHASH is a free data retrieval call binding the contract method 0x9e39b73e.
//
// Solidity: function REQUEST_TYPEHASH() view returns(bytes32)
func (_Metatx *MetatxSession) REQUESTTYPEHASH() ([32]byte, error) {
	return _Metatx.Contract.REQUESTTYPEHASH(&_Metatx.CallOpts)
}

// REQUESTTYPEHASH is a free data retrieval call binding the contract method 0x9e39b73e.
//
// Solidity: function REQUEST_TYPEHASH() view returns(bytes32)
func (_Metatx *MetatxCallerSession) REQUESTTYPEHASH() ([32]byte, error) {
	return _Metatx.Contract.REQUESTTYPEHASH(&_Metatx.CallOpts)
}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Metatx *MetatxCaller) Domains(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Metatx.contract.Call(opts, &out, "domains", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Metatx *MetatxSession) Domains(arg0 [32]byte) (bool, error) {
	return _Metatx.Contract.Domains(&_Metatx.CallOpts, arg0)
}

// Domains is a free data retrieval call binding the contract method 0xc722f177.
//
// Solidity: function domains(bytes32 ) view returns(bool)
func (_Metatx *MetatxCallerSession) Domains(arg0 [32]byte) (bool, error) {
	return _Metatx.Contract.Domains(&_Metatx.CallOpts, arg0)
}

// GetNonce is a free data retrieval call binding the contract method 0x89535803.
//
// Solidity: function getNonce(address from, uint256 batchId) view returns(uint256)
func (_Metatx *MetatxCaller) GetNonce(opts *bind.CallOpts, from common.Address, batchId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Metatx.contract.Call(opts, &out, "getNonce", from, batchId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x89535803.
//
// Solidity: function getNonce(address from, uint256 batchId) view returns(uint256)
func (_Metatx *MetatxSession) GetNonce(from common.Address, batchId *big.Int) (*big.Int, error) {
	return _Metatx.Contract.GetNonce(&_Metatx.CallOpts, from, batchId)
}

// GetNonce is a free data retrieval call binding the contract method 0x89535803.
//
// Solidity: function getNonce(address from, uint256 batchId) view returns(uint256)
func (_Metatx *MetatxCallerSession) GetNonce(from common.Address, batchId *big.Int) (*big.Int, error) {
	return _Metatx.Contract.GetNonce(&_Metatx.CallOpts, from, batchId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Metatx *MetatxCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Metatx.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Metatx *MetatxSession) IsOwner() (bool, error) {
	return _Metatx.Contract.IsOwner(&_Metatx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Metatx *MetatxCallerSession) IsOwner() (bool, error) {
	return _Metatx.Contract.IsOwner(&_Metatx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Metatx *MetatxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Metatx.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Metatx *MetatxSession) Owner() (common.Address, error) {
	return _Metatx.Contract.Owner(&_Metatx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Metatx *MetatxCallerSession) Owner() (common.Address, error) {
	return _Metatx.Contract.Owner(&_Metatx.CallOpts)
}

// VerifyEIP712 is a free data retrieval call binding the contract method 0xa41a03f2.
//
// Solidity: function verifyEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) view returns()
func (_Metatx *MetatxCaller) VerifyEIP712(opts *bind.CallOpts, req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) error {
	var out []interface{}
	err := _Metatx.contract.Call(opts, &out, "verifyEIP712", req, domainSeparator, sig)

	if err != nil {
		return err
	}

	return err

}

// VerifyEIP712 is a free data retrieval call binding the contract method 0xa41a03f2.
//
// Solidity: function verifyEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) view returns()
func (_Metatx *MetatxSession) VerifyEIP712(req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) error {
	return _Metatx.Contract.VerifyEIP712(&_Metatx.CallOpts, req, domainSeparator, sig)
}

// VerifyEIP712 is a free data retrieval call binding the contract method 0xa41a03f2.
//
// Solidity: function verifyEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) view returns()
func (_Metatx *MetatxCallerSession) VerifyEIP712(req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) error {
	return _Metatx.Contract.VerifyEIP712(&_Metatx.CallOpts, req, domainSeparator, sig)
}

// VerifyPersonalSign is a free data retrieval call binding the contract method 0x6e4cb075.
//
// Solidity: function verifyPersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) view returns()
func (_Metatx *MetatxCaller) VerifyPersonalSign(opts *bind.CallOpts, req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) error {
	var out []interface{}
	err := _Metatx.contract.Call(opts, &out, "verifyPersonalSign", req, sig)

	if err != nil {
		return err
	}

	return err

}

// VerifyPersonalSign is a free data retrieval call binding the contract method 0x6e4cb075.
//
// Solidity: function verifyPersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) view returns()
func (_Metatx *MetatxSession) VerifyPersonalSign(req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) error {
	return _Metatx.Contract.VerifyPersonalSign(&_Metatx.CallOpts, req, sig)
}

// VerifyPersonalSign is a free data retrieval call binding the contract method 0x6e4cb075.
//
// Solidity: function verifyPersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) view returns()
func (_Metatx *MetatxCallerSession) VerifyPersonalSign(req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) error {
	return _Metatx.Contract.VerifyPersonalSign(&_Metatx.CallOpts, req, sig)
}

// ExecuteEIP712 is a paid mutator transaction binding the contract method 0x41706c4e.
//
// Solidity: function executeEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) returns(bool success, bytes ret)
func (_Metatx *MetatxTransactor) ExecuteEIP712(opts *bind.TransactOpts, req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) (*types.Transaction, error) {
	return _Metatx.contract.Transact(opts, "executeEIP712", req, domainSeparator, sig)
}

// ExecuteEIP712 is a paid mutator transaction binding the contract method 0x41706c4e.
//
// Solidity: function executeEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) returns(bool success, bytes ret)
func (_Metatx *MetatxSession) ExecuteEIP712(req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) (*types.Transaction, error) {
	return _Metatx.Contract.ExecuteEIP712(&_Metatx.TransactOpts, req, domainSeparator, sig)
}

// ExecuteEIP712 is a paid mutator transaction binding the contract method 0x41706c4e.
//
// Solidity: function executeEIP712((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes32 domainSeparator, bytes sig) returns(bool success, bytes ret)
func (_Metatx *MetatxTransactorSession) ExecuteEIP712(req ERC20ForwardRequestTypesERC20ForwardRequest, domainSeparator [32]byte, sig []byte) (*types.Transaction, error) {
	return _Metatx.Contract.ExecuteEIP712(&_Metatx.TransactOpts, req, domainSeparator, sig)
}

// ExecutePersonalSign is a paid mutator transaction binding the contract method 0x8171e632.
//
// Solidity: function executePersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) returns(bool success, bytes ret)
func (_Metatx *MetatxTransactor) ExecutePersonalSign(opts *bind.TransactOpts, req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) (*types.Transaction, error) {
	return _Metatx.contract.Transact(opts, "executePersonalSign", req, sig)
}

// ExecutePersonalSign is a paid mutator transaction binding the contract method 0x8171e632.
//
// Solidity: function executePersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) returns(bool success, bytes ret)
func (_Metatx *MetatxSession) ExecutePersonalSign(req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) (*types.Transaction, error) {
	return _Metatx.Contract.ExecutePersonalSign(&_Metatx.TransactOpts, req, sig)
}

// ExecutePersonalSign is a paid mutator transaction binding the contract method 0x8171e632.
//
// Solidity: function executePersonalSign((address,address,address,uint256,uint256,uint256,uint256,uint256,bytes) req, bytes sig) returns(bool success, bytes ret)
func (_Metatx *MetatxTransactorSession) ExecutePersonalSign(req ERC20ForwardRequestTypesERC20ForwardRequest, sig []byte) (*types.Transaction, error) {
	return _Metatx.Contract.ExecutePersonalSign(&_Metatx.TransactOpts, req, sig)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Metatx *MetatxTransactor) RegisterDomainSeparator(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	return _Metatx.contract.Transact(opts, "registerDomainSeparator", name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Metatx *MetatxSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Metatx.Contract.RegisterDomainSeparator(&_Metatx.TransactOpts, name, version)
}

// RegisterDomainSeparator is a paid mutator transaction binding the contract method 0x9c7b4592.
//
// Solidity: function registerDomainSeparator(string name, string version) returns()
func (_Metatx *MetatxTransactorSession) RegisterDomainSeparator(name string, version string) (*types.Transaction, error) {
	return _Metatx.Contract.RegisterDomainSeparator(&_Metatx.TransactOpts, name, version)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Metatx *MetatxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Metatx.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Metatx *MetatxSession) RenounceOwnership() (*types.Transaction, error) {
	return _Metatx.Contract.RenounceOwnership(&_Metatx.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Metatx *MetatxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Metatx.Contract.RenounceOwnership(&_Metatx.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Metatx *MetatxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Metatx.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Metatx *MetatxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Metatx.Contract.TransferOwnership(&_Metatx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Metatx *MetatxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Metatx.Contract.TransferOwnership(&_Metatx.TransactOpts, newOwner)
}

// MetatxDomainRegisteredIterator is returned from FilterDomainRegistered and is used to iterate over the raw logs and unpacked data for DomainRegistered events raised by the Metatx contract.
type MetatxDomainRegisteredIterator struct {
	Event *MetatxDomainRegistered // Event containing the contract specifics and raw log

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
func (it *MetatxDomainRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetatxDomainRegistered)
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
		it.Event = new(MetatxDomainRegistered)
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
func (it *MetatxDomainRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetatxDomainRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetatxDomainRegistered represents a DomainRegistered event raised by the Metatx contract.
type MetatxDomainRegistered struct {
	DomainSeparator [32]byte
	DomainValue     []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDomainRegistered is a free log retrieval operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Metatx *MetatxFilterer) FilterDomainRegistered(opts *bind.FilterOpts, domainSeparator [][32]byte) (*MetatxDomainRegisteredIterator, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Metatx.contract.FilterLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return &MetatxDomainRegisteredIterator{contract: _Metatx.contract, event: "DomainRegistered", logs: logs, sub: sub}, nil
}

// WatchDomainRegistered is a free log subscription operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Metatx *MetatxFilterer) WatchDomainRegistered(opts *bind.WatchOpts, sink chan<- *MetatxDomainRegistered, domainSeparator [][32]byte) (event.Subscription, error) {

	var domainSeparatorRule []interface{}
	for _, domainSeparatorItem := range domainSeparator {
		domainSeparatorRule = append(domainSeparatorRule, domainSeparatorItem)
	}

	logs, sub, err := _Metatx.contract.WatchLogs(opts, "DomainRegistered", domainSeparatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetatxDomainRegistered)
				if err := _Metatx.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
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

// ParseDomainRegistered is a log parse operation binding the contract event 0x4bc68689cbe89a4a6333a3ab0a70093874da3e5bfb71e93102027f3f073687d8.
//
// Solidity: event DomainRegistered(bytes32 indexed domainSeparator, bytes domainValue)
func (_Metatx *MetatxFilterer) ParseDomainRegistered(log types.Log) (*MetatxDomainRegistered, error) {
	event := new(MetatxDomainRegistered)
	if err := _Metatx.contract.UnpackLog(event, "DomainRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetatxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Metatx contract.
type MetatxOwnershipTransferredIterator struct {
	Event *MetatxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MetatxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetatxOwnershipTransferred)
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
		it.Event = new(MetatxOwnershipTransferred)
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
func (it *MetatxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetatxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetatxOwnershipTransferred represents a OwnershipTransferred event raised by the Metatx contract.
type MetatxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Metatx *MetatxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MetatxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Metatx.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MetatxOwnershipTransferredIterator{contract: _Metatx.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Metatx *MetatxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MetatxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Metatx.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetatxOwnershipTransferred)
				if err := _Metatx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Metatx *MetatxFilterer) ParseOwnershipTransferred(log types.Log) (*MetatxOwnershipTransferred, error) {
	event := new(MetatxOwnershipTransferred)
	if err := _Metatx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
