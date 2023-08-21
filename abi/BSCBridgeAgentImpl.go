// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IMessageStructureMessage is an auto generated low-level Go binding around an user-defined struct.
type IMessageStructureMessage struct {
	SwapType string
	Base     string
	Quote    string
	Amount   *big.Int
	Fee      *big.Int
	Exchange *big.Int
	Nonce    *big.Int
	Deadline *big.Int
	Account  common.Address
}

// BSCBridgeAgentImplMetaData contains all meta data concerning the BSCBridgeAgentImpl contract.
var BSCBridgeAgentImplMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"metisTxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"swapFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"metisTxHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"base\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"quote\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapStarted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bnbBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnbWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscBridgeAgentImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"base\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quote\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"internalType\":\"structIMessageStructure.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"metisTxHash\",\"type\":\"bytes32\"}],\"name\":\"fillBNB2TOTMPegin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakePair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeswapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pancakePair\",\"type\":\"address\"}],\"name\":\"setPancakePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slippagePercentage\",\"type\":\"uint256\"}],\"name\":\"setSlippagePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slippagePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"swapDataMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapRouterQuery\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteReserved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wbnb\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BSCBridgeAgentImplABI is the input ABI used to generate the binding from.
// Deprecated: Use BSCBridgeAgentImplMetaData.ABI instead.
var BSCBridgeAgentImplABI = BSCBridgeAgentImplMetaData.ABI

// BSCBridgeAgentImpl is an auto generated Go binding around an Ethereum contract.
type BSCBridgeAgentImpl struct {
	BSCBridgeAgentImplCaller     // Read-only binding to the contract
	BSCBridgeAgentImplTransactor // Write-only binding to the contract
	BSCBridgeAgentImplFilterer   // Log filterer for contract events
}

// BSCBridgeAgentImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type BSCBridgeAgentImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeAgentImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCBridgeAgentImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeAgentImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCBridgeAgentImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCBridgeAgentImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCBridgeAgentImplSession struct {
	Contract     *BSCBridgeAgentImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BSCBridgeAgentImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCBridgeAgentImplCallerSession struct {
	Contract *BSCBridgeAgentImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BSCBridgeAgentImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCBridgeAgentImplTransactorSession struct {
	Contract     *BSCBridgeAgentImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BSCBridgeAgentImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type BSCBridgeAgentImplRaw struct {
	Contract *BSCBridgeAgentImpl // Generic contract binding to access the raw methods on
}

// BSCBridgeAgentImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCBridgeAgentImplCallerRaw struct {
	Contract *BSCBridgeAgentImplCaller // Generic read-only contract binding to access the raw methods on
}

// BSCBridgeAgentImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCBridgeAgentImplTransactorRaw struct {
	Contract *BSCBridgeAgentImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCBridgeAgentImpl creates a new instance of BSCBridgeAgentImpl, bound to a specific deployed contract.
func NewBSCBridgeAgentImpl(address common.Address, backend bind.ContractBackend) (*BSCBridgeAgentImpl, error) {
	contract, err := bindBSCBridgeAgentImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeAgentImpl{BSCBridgeAgentImplCaller: BSCBridgeAgentImplCaller{contract: contract}, BSCBridgeAgentImplTransactor: BSCBridgeAgentImplTransactor{contract: contract}, BSCBridgeAgentImplFilterer: BSCBridgeAgentImplFilterer{contract: contract}}, nil
}

// NewBSCBridgeAgentImplCaller creates a new read-only instance of BSCBridgeAgentImpl, bound to a specific deployed contract.
func NewBSCBridgeAgentImplCaller(address common.Address, caller bind.ContractCaller) (*BSCBridgeAgentImplCaller, error) {
	contract, err := bindBSCBridgeAgentImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeAgentImplCaller{contract: contract}, nil
}

// NewBSCBridgeAgentImplTransactor creates a new write-only instance of BSCBridgeAgentImpl, bound to a specific deployed contract.
func NewBSCBridgeAgentImplTransactor(address common.Address, transactor bind.ContractTransactor) (*BSCBridgeAgentImplTransactor, error) {
	contract, err := bindBSCBridgeAgentImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeAgentImplTransactor{contract: contract}, nil
}

// NewBSCBridgeAgentImplFilterer creates a new log filterer instance of BSCBridgeAgentImpl, bound to a specific deployed contract.
func NewBSCBridgeAgentImplFilterer(address common.Address, filterer bind.ContractFilterer) (*BSCBridgeAgentImplFilterer, error) {
	contract, err := bindBSCBridgeAgentImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeAgentImplFilterer{contract: contract}, nil
}

// bindBSCBridgeAgentImpl binds a generic wrapper to an already deployed contract.
func bindBSCBridgeAgentImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BSCBridgeAgentImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCBridgeAgentImpl.Contract.BSCBridgeAgentImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.BSCBridgeAgentImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.BSCBridgeAgentImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCBridgeAgentImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.contract.Transact(opts, method, params...)
}

// BnbBalance is a free data retrieval call binding the contract method 0xd013cbe2.
//
// Solidity: function bnbBalance() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) BnbBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "bnbBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BnbBalance is a free data retrieval call binding the contract method 0xd013cbe2.
//
// Solidity: function bnbBalance() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) BnbBalance() (*big.Int, error) {
	return _BSCBridgeAgentImpl.Contract.BnbBalance(&_BSCBridgeAgentImpl.CallOpts)
}

// BnbBalance is a free data retrieval call binding the contract method 0xd013cbe2.
//
// Solidity: function bnbBalance() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) BnbBalance() (*big.Int, error) {
	return _BSCBridgeAgentImpl.Contract.BnbBalance(&_BSCBridgeAgentImpl.CallOpts)
}

// BscBridgeAgentImpl is a free data retrieval call binding the contract method 0xd4fde9c9.
//
// Solidity: function bscBridgeAgentImpl() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) BscBridgeAgentImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "bscBridgeAgentImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BscBridgeAgentImpl is a free data retrieval call binding the contract method 0xd4fde9c9.
//
// Solidity: function bscBridgeAgentImpl() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) BscBridgeAgentImpl() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.BscBridgeAgentImpl(&_BSCBridgeAgentImpl.CallOpts)
}

// BscBridgeAgentImpl is a free data retrieval call binding the contract method 0xd4fde9c9.
//
// Solidity: function bscBridgeAgentImpl() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) BscBridgeAgentImpl() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.BscBridgeAgentImpl(&_BSCBridgeAgentImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) Owner() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.Owner(&_BSCBridgeAgentImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) Owner() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.Owner(&_BSCBridgeAgentImpl.CallOpts)
}

// PancakePair is a free data retrieval call binding the contract method 0xb8c9d25c.
//
// Solidity: function pancakePair() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) PancakePair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "pancakePair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakePair is a free data retrieval call binding the contract method 0xb8c9d25c.
//
// Solidity: function pancakePair() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) PancakePair() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.PancakePair(&_BSCBridgeAgentImpl.CallOpts)
}

// PancakePair is a free data retrieval call binding the contract method 0xb8c9d25c.
//
// Solidity: function pancakePair() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) PancakePair() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.PancakePair(&_BSCBridgeAgentImpl.CallOpts)
}

// PancakeswapRouter is a free data retrieval call binding the contract method 0xdb6754ed.
//
// Solidity: function pancakeswapRouter() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) PancakeswapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "pancakeswapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeswapRouter is a free data retrieval call binding the contract method 0xdb6754ed.
//
// Solidity: function pancakeswapRouter() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) PancakeswapRouter() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.PancakeswapRouter(&_BSCBridgeAgentImpl.CallOpts)
}

// PancakeswapRouter is a free data retrieval call binding the contract method 0xdb6754ed.
//
// Solidity: function pancakeswapRouter() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) PancakeswapRouter() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.PancakeswapRouter(&_BSCBridgeAgentImpl.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "proxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) ProxyAdmin() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.ProxyAdmin(&_BSCBridgeAgentImpl.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) ProxyAdmin() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.ProxyAdmin(&_BSCBridgeAgentImpl.CallOpts)
}

// SlippagePercentage is a free data retrieval call binding the contract method 0x4452d81c.
//
// Solidity: function slippagePercentage() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) SlippagePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "slippagePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlippagePercentage is a free data retrieval call binding the contract method 0x4452d81c.
//
// Solidity: function slippagePercentage() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) SlippagePercentage() (*big.Int, error) {
	return _BSCBridgeAgentImpl.Contract.SlippagePercentage(&_BSCBridgeAgentImpl.CallOpts)
}

// SlippagePercentage is a free data retrieval call binding the contract method 0x4452d81c.
//
// Solidity: function slippagePercentage() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) SlippagePercentage() (*big.Int, error) {
	return _BSCBridgeAgentImpl.Contract.SlippagePercentage(&_BSCBridgeAgentImpl.CallOpts)
}

// SwapDataMap is a free data retrieval call binding the contract method 0xa7d6dc1b.
//
// Solidity: function swapDataMap(address , bytes32 ) view returns(uint256 nonce, uint256 exchange, string swapType)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) SwapDataMap(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Nonce    *big.Int
	Exchange *big.Int
	SwapType string
}, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "swapDataMap", arg0, arg1)

	outstruct := new(struct {
		Nonce    *big.Int
		Exchange *big.Int
		SwapType string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Exchange = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SwapType = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// SwapDataMap is a free data retrieval call binding the contract method 0xa7d6dc1b.
//
// Solidity: function swapDataMap(address , bytes32 ) view returns(uint256 nonce, uint256 exchange, string swapType)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) SwapDataMap(arg0 common.Address, arg1 [32]byte) (struct {
	Nonce    *big.Int
	Exchange *big.Int
	SwapType string
}, error) {
	return _BSCBridgeAgentImpl.Contract.SwapDataMap(&_BSCBridgeAgentImpl.CallOpts, arg0, arg1)
}

// SwapDataMap is a free data retrieval call binding the contract method 0xa7d6dc1b.
//
// Solidity: function swapDataMap(address , bytes32 ) view returns(uint256 nonce, uint256 exchange, string swapType)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) SwapDataMap(arg0 common.Address, arg1 [32]byte) (struct {
	Nonce    *big.Int
	Exchange *big.Int
	SwapType string
}, error) {
	return _BSCBridgeAgentImpl.Contract.SwapDataMap(&_BSCBridgeAgentImpl.CallOpts, arg0, arg1)
}

// SwapRouterQuery is a free data retrieval call binding the contract method 0xc37dcbd4.
//
// Solidity: function swapRouterQuery(uint256 amount) view returns(uint256 baseReserved, uint256 quoteReserved, uint256 exchange)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) SwapRouterQuery(opts *bind.CallOpts, amount *big.Int) (struct {
	BaseReserved  *big.Int
	QuoteReserved *big.Int
	Exchange      *big.Int
}, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "swapRouterQuery", amount)

	outstruct := new(struct {
		BaseReserved  *big.Int
		QuoteReserved *big.Int
		Exchange      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseReserved = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.QuoteReserved = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exchange = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SwapRouterQuery is a free data retrieval call binding the contract method 0xc37dcbd4.
//
// Solidity: function swapRouterQuery(uint256 amount) view returns(uint256 baseReserved, uint256 quoteReserved, uint256 exchange)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) SwapRouterQuery(amount *big.Int) (struct {
	BaseReserved  *big.Int
	QuoteReserved *big.Int
	Exchange      *big.Int
}, error) {
	return _BSCBridgeAgentImpl.Contract.SwapRouterQuery(&_BSCBridgeAgentImpl.CallOpts, amount)
}

// SwapRouterQuery is a free data retrieval call binding the contract method 0xc37dcbd4.
//
// Solidity: function swapRouterQuery(uint256 amount) view returns(uint256 baseReserved, uint256 quoteReserved, uint256 exchange)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) SwapRouterQuery(amount *big.Int) (struct {
	BaseReserved  *big.Int
	QuoteReserved *big.Int
	Exchange      *big.Int
}, error) {
	return _BSCBridgeAgentImpl.Contract.SwapRouterQuery(&_BSCBridgeAgentImpl.CallOpts, amount)
}

// TotemBalance is a free data retrieval call binding the contract method 0xe2621d59.
//
// Solidity: function totemBalance() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) TotemBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "totemBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotemBalance is a free data retrieval call binding the contract method 0xe2621d59.
//
// Solidity: function totemBalance() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) TotemBalance() (*big.Int, error) {
	return _BSCBridgeAgentImpl.Contract.TotemBalance(&_BSCBridgeAgentImpl.CallOpts)
}

// TotemBalance is a free data retrieval call binding the contract method 0xe2621d59.
//
// Solidity: function totemBalance() view returns(uint256)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) TotemBalance() (*big.Int, error) {
	return _BSCBridgeAgentImpl.Contract.TotemBalance(&_BSCBridgeAgentImpl.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) TotemToken() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.TotemToken(&_BSCBridgeAgentImpl.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) TotemToken() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.TotemToken(&_BSCBridgeAgentImpl.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) UpgradeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "upgradeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) UpgradeEnabled() (bool, error) {
	return _BSCBridgeAgentImpl.Contract.UpgradeEnabled(&_BSCBridgeAgentImpl.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) UpgradeEnabled() (bool, error) {
	return _BSCBridgeAgentImpl.Contract.UpgradeEnabled(&_BSCBridgeAgentImpl.CallOpts)
}

// Wbnb is a free data retrieval call binding the contract method 0x8d72647e.
//
// Solidity: function wbnb() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCaller) Wbnb(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCBridgeAgentImpl.contract.Call(opts, &out, "wbnb")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wbnb is a free data retrieval call binding the contract method 0x8d72647e.
//
// Solidity: function wbnb() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) Wbnb() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.Wbnb(&_BSCBridgeAgentImpl.CallOpts)
}

// Wbnb is a free data retrieval call binding the contract method 0x8d72647e.
//
// Solidity: function wbnb() view returns(address)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplCallerSession) Wbnb() (common.Address, error) {
	return _BSCBridgeAgentImpl.Contract.Wbnb(&_BSCBridgeAgentImpl.CallOpts)
}

// BnbWithdraw is a paid mutator transaction binding the contract method 0x75bbeee7.
//
// Solidity: function bnbWithdraw() returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactor) BnbWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.contract.Transact(opts, "bnbWithdraw")
}

// BnbWithdraw is a paid mutator transaction binding the contract method 0x75bbeee7.
//
// Solidity: function bnbWithdraw() returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) BnbWithdraw() (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.BnbWithdraw(&_BSCBridgeAgentImpl.TransactOpts)
}

// BnbWithdraw is a paid mutator transaction binding the contract method 0x75bbeee7.
//
// Solidity: function bnbWithdraw() returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactorSession) BnbWithdraw() (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.BnbWithdraw(&_BSCBridgeAgentImpl.TransactOpts)
}

// FillBNB2TOTMPegin is a paid mutator transaction binding the contract method 0x635d9a47.
//
// Solidity: function fillBNB2TOTMPegin((string,string,string,uint256,uint256,uint256,uint256,uint256,address) message, bytes32 metisTxHash) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactor) FillBNB2TOTMPegin(opts *bind.TransactOpts, message IMessageStructureMessage, metisTxHash [32]byte) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.contract.Transact(opts, "fillBNB2TOTMPegin", message, metisTxHash)
}

// FillBNB2TOTMPegin is a paid mutator transaction binding the contract method 0x635d9a47.
//
// Solidity: function fillBNB2TOTMPegin((string,string,string,uint256,uint256,uint256,uint256,uint256,address) message, bytes32 metisTxHash) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) FillBNB2TOTMPegin(message IMessageStructureMessage, metisTxHash [32]byte) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.FillBNB2TOTMPegin(&_BSCBridgeAgentImpl.TransactOpts, message, metisTxHash)
}

// FillBNB2TOTMPegin is a paid mutator transaction binding the contract method 0x635d9a47.
//
// Solidity: function fillBNB2TOTMPegin((string,string,string,uint256,uint256,uint256,uint256,uint256,address) message, bytes32 metisTxHash) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactorSession) FillBNB2TOTMPegin(message IMessageStructureMessage, metisTxHash [32]byte) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.FillBNB2TOTMPegin(&_BSCBridgeAgentImpl.TransactOpts, message, metisTxHash)
}

// SetPancakePair is a paid mutator transaction binding the contract method 0xa5b601be.
//
// Solidity: function setPancakePair(address _pancakePair) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactor) SetPancakePair(opts *bind.TransactOpts, _pancakePair common.Address) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.contract.Transact(opts, "setPancakePair", _pancakePair)
}

// SetPancakePair is a paid mutator transaction binding the contract method 0xa5b601be.
//
// Solidity: function setPancakePair(address _pancakePair) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) SetPancakePair(_pancakePair common.Address) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.SetPancakePair(&_BSCBridgeAgentImpl.TransactOpts, _pancakePair)
}

// SetPancakePair is a paid mutator transaction binding the contract method 0xa5b601be.
//
// Solidity: function setPancakePair(address _pancakePair) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactorSession) SetPancakePair(_pancakePair common.Address) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.SetPancakePair(&_BSCBridgeAgentImpl.TransactOpts, _pancakePair)
}

// SetSlippagePercentage is a paid mutator transaction binding the contract method 0x03baf066.
//
// Solidity: function setSlippagePercentage(uint256 _slippagePercentage) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactor) SetSlippagePercentage(opts *bind.TransactOpts, _slippagePercentage *big.Int) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.contract.Transact(opts, "setSlippagePercentage", _slippagePercentage)
}

// SetSlippagePercentage is a paid mutator transaction binding the contract method 0x03baf066.
//
// Solidity: function setSlippagePercentage(uint256 _slippagePercentage) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplSession) SetSlippagePercentage(_slippagePercentage *big.Int) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.SetSlippagePercentage(&_BSCBridgeAgentImpl.TransactOpts, _slippagePercentage)
}

// SetSlippagePercentage is a paid mutator transaction binding the contract method 0x03baf066.
//
// Solidity: function setSlippagePercentage(uint256 _slippagePercentage) returns()
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplTransactorSession) SetSlippagePercentage(_slippagePercentage *big.Int) (*types.Transaction, error) {
	return _BSCBridgeAgentImpl.Contract.SetSlippagePercentage(&_BSCBridgeAgentImpl.TransactOpts, _slippagePercentage)
}

// BSCBridgeAgentImplSwapFilledIterator is returned from FilterSwapFilled and is used to iterate over the raw logs and unpacked data for SwapFilled events raised by the BSCBridgeAgentImpl contract.
type BSCBridgeAgentImplSwapFilledIterator struct {
	Event *BSCBridgeAgentImplSwapFilled // Event containing the contract specifics and raw log

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
func (it *BSCBridgeAgentImplSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeAgentImplSwapFilled)
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
		it.Event = new(BSCBridgeAgentImplSwapFilled)
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
func (it *BSCBridgeAgentImplSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeAgentImplSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeAgentImplSwapFilled represents a SwapFilled event raised by the BSCBridgeAgentImpl contract.
type BSCBridgeAgentImplSwapFilled struct {
	Recipient   common.Address
	MetisTxHash [32]byte
	SwapType    string
	Amount      *big.Int
	Fee         *big.Int
	Exchange    *big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapFilled is a free log retrieval operation binding the contract event 0x061eaa4b2045ee0ba76d9754efebbb41b3b16198923b4c19afcc434b6136d3b0.
//
// Solidity: event swapFilled(address indexed recipient, bytes32 indexed metisTxHash, string swapType, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplFilterer) FilterSwapFilled(opts *bind.FilterOpts, recipient []common.Address, metisTxHash [][32]byte) (*BSCBridgeAgentImplSwapFilledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var metisTxHashRule []interface{}
	for _, metisTxHashItem := range metisTxHash {
		metisTxHashRule = append(metisTxHashRule, metisTxHashItem)
	}

	logs, sub, err := _BSCBridgeAgentImpl.contract.FilterLogs(opts, "swapFilled", recipientRule, metisTxHashRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeAgentImplSwapFilledIterator{contract: _BSCBridgeAgentImpl.contract, event: "swapFilled", logs: logs, sub: sub}, nil
}

// WatchSwapFilled is a free log subscription operation binding the contract event 0x061eaa4b2045ee0ba76d9754efebbb41b3b16198923b4c19afcc434b6136d3b0.
//
// Solidity: event swapFilled(address indexed recipient, bytes32 indexed metisTxHash, string swapType, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplFilterer) WatchSwapFilled(opts *bind.WatchOpts, sink chan<- *BSCBridgeAgentImplSwapFilled, recipient []common.Address, metisTxHash [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var metisTxHashRule []interface{}
	for _, metisTxHashItem := range metisTxHash {
		metisTxHashRule = append(metisTxHashRule, metisTxHashItem)
	}

	logs, sub, err := _BSCBridgeAgentImpl.contract.WatchLogs(opts, "swapFilled", recipientRule, metisTxHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeAgentImplSwapFilled)
				if err := _BSCBridgeAgentImpl.contract.UnpackLog(event, "swapFilled", log); err != nil {
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

// ParseSwapFilled is a log parse operation binding the contract event 0x061eaa4b2045ee0ba76d9754efebbb41b3b16198923b4c19afcc434b6136d3b0.
//
// Solidity: event swapFilled(address indexed recipient, bytes32 indexed metisTxHash, string swapType, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplFilterer) ParseSwapFilled(log types.Log) (*BSCBridgeAgentImplSwapFilled, error) {
	event := new(BSCBridgeAgentImplSwapFilled)
	if err := _BSCBridgeAgentImpl.contract.UnpackLog(event, "swapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCBridgeAgentImplSwapStartedIterator is returned from FilterSwapStarted and is used to iterate over the raw logs and unpacked data for SwapStarted events raised by the BSCBridgeAgentImpl contract.
type BSCBridgeAgentImplSwapStartedIterator struct {
	Event *BSCBridgeAgentImplSwapStarted // Event containing the contract specifics and raw log

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
func (it *BSCBridgeAgentImplSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCBridgeAgentImplSwapStarted)
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
		it.Event = new(BSCBridgeAgentImplSwapStarted)
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
func (it *BSCBridgeAgentImplSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCBridgeAgentImplSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCBridgeAgentImplSwapStarted represents a SwapStarted event raised by the BSCBridgeAgentImpl contract.
type BSCBridgeAgentImplSwapStarted struct {
	Spender     common.Address
	MetisTxHash [32]byte
	SwapType    string
	Base        string
	Quote       string
	Amount      *big.Int
	Fee         *big.Int
	Exchange    *big.Int
	Nonce       *big.Int
	Deadline    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapStarted is a free log retrieval operation binding the contract event 0xddbc9ff29a9285afc6faa85774fd182b7d957a0096a14716b88d020717071b1c.
//
// Solidity: event swapStarted(address indexed spender, bytes32 indexed metisTxHash, string swapType, string base, string quote, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce, uint256 deadline)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplFilterer) FilterSwapStarted(opts *bind.FilterOpts, spender []common.Address, metisTxHash [][32]byte) (*BSCBridgeAgentImplSwapStartedIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var metisTxHashRule []interface{}
	for _, metisTxHashItem := range metisTxHash {
		metisTxHashRule = append(metisTxHashRule, metisTxHashItem)
	}

	logs, sub, err := _BSCBridgeAgentImpl.contract.FilterLogs(opts, "swapStarted", spenderRule, metisTxHashRule)
	if err != nil {
		return nil, err
	}
	return &BSCBridgeAgentImplSwapStartedIterator{contract: _BSCBridgeAgentImpl.contract, event: "swapStarted", logs: logs, sub: sub}, nil
}

// WatchSwapStarted is a free log subscription operation binding the contract event 0xddbc9ff29a9285afc6faa85774fd182b7d957a0096a14716b88d020717071b1c.
//
// Solidity: event swapStarted(address indexed spender, bytes32 indexed metisTxHash, string swapType, string base, string quote, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce, uint256 deadline)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplFilterer) WatchSwapStarted(opts *bind.WatchOpts, sink chan<- *BSCBridgeAgentImplSwapStarted, spender []common.Address, metisTxHash [][32]byte) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var metisTxHashRule []interface{}
	for _, metisTxHashItem := range metisTxHash {
		metisTxHashRule = append(metisTxHashRule, metisTxHashItem)
	}

	logs, sub, err := _BSCBridgeAgentImpl.contract.WatchLogs(opts, "swapStarted", spenderRule, metisTxHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCBridgeAgentImplSwapStarted)
				if err := _BSCBridgeAgentImpl.contract.UnpackLog(event, "swapStarted", log); err != nil {
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

// ParseSwapStarted is a log parse operation binding the contract event 0xddbc9ff29a9285afc6faa85774fd182b7d957a0096a14716b88d020717071b1c.
//
// Solidity: event swapStarted(address indexed spender, bytes32 indexed metisTxHash, string swapType, string base, string quote, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce, uint256 deadline)
func (_BSCBridgeAgentImpl *BSCBridgeAgentImplFilterer) ParseSwapStarted(log types.Log) (*BSCBridgeAgentImplSwapStarted, error) {
	event := new(BSCBridgeAgentImplSwapStarted)
	if err := _BSCBridgeAgentImpl.contract.UnpackLog(event, "swapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
