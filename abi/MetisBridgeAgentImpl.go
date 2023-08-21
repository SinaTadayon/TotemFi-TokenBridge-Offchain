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

// MetisBridgeAgentImplMetaData contains all meta data concerning the MetisBridgeAgentImpl contract.
var MetisBridgeAgentImplMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"}],\"name\":\"swapClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bscTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"}],\"name\":\"swapFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"base\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"quote\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"claimMTS2TOTMPegIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domain\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bscTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"}],\"name\":\"fillMTS2TOTMPegIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metisBridgeAgentImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mtsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mtsWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"base\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quote\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"internalType\":\"structIMessageStructure.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"pegInMTS2TOTM\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peggedTotemTokenProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"claimStat\",\"type\":\"bool\"}],\"name\":\"setClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"fillStat\",\"type\":\"bool\"}],\"name\":\"setFill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"swapDataMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exchange\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"swapType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isFilled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isClaimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MetisBridgeAgentImplABI is the input ABI used to generate the binding from.
// Deprecated: Use MetisBridgeAgentImplMetaData.ABI instead.
var MetisBridgeAgentImplABI = MetisBridgeAgentImplMetaData.ABI

// MetisBridgeAgentImpl is an auto generated Go binding around an Ethereum contract.
type MetisBridgeAgentImpl struct {
	MetisBridgeAgentImplCaller     // Read-only binding to the contract
	MetisBridgeAgentImplTransactor // Write-only binding to the contract
	MetisBridgeAgentImplFilterer   // Log filterer for contract events
}

// MetisBridgeAgentImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetisBridgeAgentImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetisBridgeAgentImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetisBridgeAgentImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetisBridgeAgentImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetisBridgeAgentImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetisBridgeAgentImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetisBridgeAgentImplSession struct {
	Contract     *MetisBridgeAgentImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MetisBridgeAgentImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetisBridgeAgentImplCallerSession struct {
	Contract *MetisBridgeAgentImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MetisBridgeAgentImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetisBridgeAgentImplTransactorSession struct {
	Contract     *MetisBridgeAgentImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MetisBridgeAgentImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetisBridgeAgentImplRaw struct {
	Contract *MetisBridgeAgentImpl // Generic contract binding to access the raw methods on
}

// MetisBridgeAgentImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetisBridgeAgentImplCallerRaw struct {
	Contract *MetisBridgeAgentImplCaller // Generic read-only contract binding to access the raw methods on
}

// MetisBridgeAgentImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetisBridgeAgentImplTransactorRaw struct {
	Contract *MetisBridgeAgentImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetisBridgeAgentImpl creates a new instance of MetisBridgeAgentImpl, bound to a specific deployed contract.
func NewMetisBridgeAgentImpl(address common.Address, backend bind.ContractBackend) (*MetisBridgeAgentImpl, error) {
	contract, err := bindMetisBridgeAgentImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetisBridgeAgentImpl{MetisBridgeAgentImplCaller: MetisBridgeAgentImplCaller{contract: contract}, MetisBridgeAgentImplTransactor: MetisBridgeAgentImplTransactor{contract: contract}, MetisBridgeAgentImplFilterer: MetisBridgeAgentImplFilterer{contract: contract}}, nil
}

// NewMetisBridgeAgentImplCaller creates a new read-only instance of MetisBridgeAgentImpl, bound to a specific deployed contract.
func NewMetisBridgeAgentImplCaller(address common.Address, caller bind.ContractCaller) (*MetisBridgeAgentImplCaller, error) {
	contract, err := bindMetisBridgeAgentImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetisBridgeAgentImplCaller{contract: contract}, nil
}

// NewMetisBridgeAgentImplTransactor creates a new write-only instance of MetisBridgeAgentImpl, bound to a specific deployed contract.
func NewMetisBridgeAgentImplTransactor(address common.Address, transactor bind.ContractTransactor) (*MetisBridgeAgentImplTransactor, error) {
	contract, err := bindMetisBridgeAgentImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetisBridgeAgentImplTransactor{contract: contract}, nil
}

// NewMetisBridgeAgentImplFilterer creates a new log filterer instance of MetisBridgeAgentImpl, bound to a specific deployed contract.
func NewMetisBridgeAgentImplFilterer(address common.Address, filterer bind.ContractFilterer) (*MetisBridgeAgentImplFilterer, error) {
	contract, err := bindMetisBridgeAgentImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetisBridgeAgentImplFilterer{contract: contract}, nil
}

// bindMetisBridgeAgentImpl binds a generic wrapper to an already deployed contract.
func bindMetisBridgeAgentImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetisBridgeAgentImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetisBridgeAgentImpl.Contract.MetisBridgeAgentImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.MetisBridgeAgentImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.MetisBridgeAgentImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetisBridgeAgentImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.contract.Transact(opts, method, params...)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) Domain(opts *bind.CallOpts) (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "domain")

	outstruct := new(struct {
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _MetisBridgeAgentImpl.Contract.Domain(&_MetisBridgeAgentImpl.CallOpts)
}

// Domain is a free data retrieval call binding the contract method 0xc2fb26a6.
//
// Solidity: function domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _MetisBridgeAgentImpl.Contract.Domain(&_MetisBridgeAgentImpl.CallOpts)
}

// MetisBridgeAgentImpl is a free data retrieval call binding the contract method 0x9afad3ed.
//
// Solidity: function metisBridgeAgentImpl() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) MetisBridgeAgentImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "metisBridgeAgentImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetisBridgeAgentImpl is a free data retrieval call binding the contract method 0x9afad3ed.
//
// Solidity: function metisBridgeAgentImpl() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) MetisBridgeAgentImpl() (common.Address, error) {
	return _MetisBridgeAgentImpl.Contract.MetisBridgeAgentImpl(&_MetisBridgeAgentImpl.CallOpts)
}

// MetisBridgeAgentImpl is a free data retrieval call binding the contract method 0x9afad3ed.
//
// Solidity: function metisBridgeAgentImpl() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) MetisBridgeAgentImpl() (common.Address, error) {
	return _MetisBridgeAgentImpl.Contract.MetisBridgeAgentImpl(&_MetisBridgeAgentImpl.CallOpts)
}

// MtsBalance is a free data retrieval call binding the contract method 0x884034ad.
//
// Solidity: function mtsBalance() view returns(uint256)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) MtsBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "mtsBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MtsBalance is a free data retrieval call binding the contract method 0x884034ad.
//
// Solidity: function mtsBalance() view returns(uint256)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) MtsBalance() (*big.Int, error) {
	return _MetisBridgeAgentImpl.Contract.MtsBalance(&_MetisBridgeAgentImpl.CallOpts)
}

// MtsBalance is a free data retrieval call binding the contract method 0x884034ad.
//
// Solidity: function mtsBalance() view returns(uint256)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) MtsBalance() (*big.Int, error) {
	return _MetisBridgeAgentImpl.Contract.MtsBalance(&_MetisBridgeAgentImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) Owner() (common.Address, error) {
	return _MetisBridgeAgentImpl.Contract.Owner(&_MetisBridgeAgentImpl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) Owner() (common.Address, error) {
	return _MetisBridgeAgentImpl.Contract.Owner(&_MetisBridgeAgentImpl.CallOpts)
}

// PeggedTotemTokenProxy is a free data retrieval call binding the contract method 0xa6936d5c.
//
// Solidity: function peggedTotemTokenProxy() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) PeggedTotemTokenProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "peggedTotemTokenProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PeggedTotemTokenProxy is a free data retrieval call binding the contract method 0xa6936d5c.
//
// Solidity: function peggedTotemTokenProxy() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) PeggedTotemTokenProxy() (common.Address, error) {
	return _MetisBridgeAgentImpl.Contract.PeggedTotemTokenProxy(&_MetisBridgeAgentImpl.CallOpts)
}

// PeggedTotemTokenProxy is a free data retrieval call binding the contract method 0xa6936d5c.
//
// Solidity: function peggedTotemTokenProxy() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) PeggedTotemTokenProxy() (common.Address, error) {
	return _MetisBridgeAgentImpl.Contract.PeggedTotemTokenProxy(&_MetisBridgeAgentImpl.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "proxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) ProxyAdmin() (common.Address, error) {
	return _MetisBridgeAgentImpl.Contract.ProxyAdmin(&_MetisBridgeAgentImpl.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) ProxyAdmin() (common.Address, error) {
	return _MetisBridgeAgentImpl.Contract.ProxyAdmin(&_MetisBridgeAgentImpl.CallOpts)
}

// SwapDataMap is a free data retrieval call binding the contract method 0xa7d6dc1b.
//
// Solidity: function swapDataMap(address , bytes32 ) view returns(bytes32 txHash, uint256 nonce, uint256 exchange, string swapType, bool isFilled, bool isClaimed)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) SwapDataMap(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	TxHash    [32]byte
	Nonce     *big.Int
	Exchange  *big.Int
	SwapType  string
	IsFilled  bool
	IsClaimed bool
}, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "swapDataMap", arg0, arg1)

	outstruct := new(struct {
		TxHash    [32]byte
		Nonce     *big.Int
		Exchange  *big.Int
		SwapType  string
		IsFilled  bool
		IsClaimed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exchange = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SwapType = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsFilled = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.IsClaimed = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// SwapDataMap is a free data retrieval call binding the contract method 0xa7d6dc1b.
//
// Solidity: function swapDataMap(address , bytes32 ) view returns(bytes32 txHash, uint256 nonce, uint256 exchange, string swapType, bool isFilled, bool isClaimed)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) SwapDataMap(arg0 common.Address, arg1 [32]byte) (struct {
	TxHash    [32]byte
	Nonce     *big.Int
	Exchange  *big.Int
	SwapType  string
	IsFilled  bool
	IsClaimed bool
}, error) {
	return _MetisBridgeAgentImpl.Contract.SwapDataMap(&_MetisBridgeAgentImpl.CallOpts, arg0, arg1)
}

// SwapDataMap is a free data retrieval call binding the contract method 0xa7d6dc1b.
//
// Solidity: function swapDataMap(address , bytes32 ) view returns(bytes32 txHash, uint256 nonce, uint256 exchange, string swapType, bool isFilled, bool isClaimed)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) SwapDataMap(arg0 common.Address, arg1 [32]byte) (struct {
	TxHash    [32]byte
	Nonce     *big.Int
	Exchange  *big.Int
	SwapType  string
	IsFilled  bool
	IsClaimed bool
}, error) {
	return _MetisBridgeAgentImpl.Contract.SwapDataMap(&_MetisBridgeAgentImpl.CallOpts, arg0, arg1)
}

// SwapNonce is a free data retrieval call binding the contract method 0x6dbf4547.
//
// Solidity: function swapNonce(address ) view returns(uint256)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) SwapNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "swapNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapNonce is a free data retrieval call binding the contract method 0x6dbf4547.
//
// Solidity: function swapNonce(address ) view returns(uint256)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) SwapNonce(arg0 common.Address) (*big.Int, error) {
	return _MetisBridgeAgentImpl.Contract.SwapNonce(&_MetisBridgeAgentImpl.CallOpts, arg0)
}

// SwapNonce is a free data retrieval call binding the contract method 0x6dbf4547.
//
// Solidity: function swapNonce(address ) view returns(uint256)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) SwapNonce(arg0 common.Address) (*big.Int, error) {
	return _MetisBridgeAgentImpl.Contract.SwapNonce(&_MetisBridgeAgentImpl.CallOpts, arg0)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCaller) UpgradeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MetisBridgeAgentImpl.contract.Call(opts, &out, "upgradeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) UpgradeEnabled() (bool, error) {
	return _MetisBridgeAgentImpl.Contract.UpgradeEnabled(&_MetisBridgeAgentImpl.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplCallerSession) UpgradeEnabled() (bool, error) {
	return _MetisBridgeAgentImpl.Contract.UpgradeEnabled(&_MetisBridgeAgentImpl.CallOpts)
}

// ClaimMTS2TOTMPegIn is a paid mutator transaction binding the contract method 0x3ced6d08.
//
// Solidity: function claimMTS2TOTMPegIn(bytes32 dataHash) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactor) ClaimMTS2TOTMPegIn(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.contract.Transact(opts, "claimMTS2TOTMPegIn", dataHash)
}

// ClaimMTS2TOTMPegIn is a paid mutator transaction binding the contract method 0x3ced6d08.
//
// Solidity: function claimMTS2TOTMPegIn(bytes32 dataHash) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) ClaimMTS2TOTMPegIn(dataHash [32]byte) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.ClaimMTS2TOTMPegIn(&_MetisBridgeAgentImpl.TransactOpts, dataHash)
}

// ClaimMTS2TOTMPegIn is a paid mutator transaction binding the contract method 0x3ced6d08.
//
// Solidity: function claimMTS2TOTMPegIn(bytes32 dataHash) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorSession) ClaimMTS2TOTMPegIn(dataHash [32]byte) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.ClaimMTS2TOTMPegIn(&_MetisBridgeAgentImpl.TransactOpts, dataHash)
}

// FillMTS2TOTMPegIn is a paid mutator transaction binding the contract method 0x6dfa08de.
//
// Solidity: function fillMTS2TOTMPegIn(address recipient, bytes32 dataHash, bytes32 bscTxHash, string swapType, uint256 fee, uint256 exchange) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactor) FillMTS2TOTMPegIn(opts *bind.TransactOpts, recipient common.Address, dataHash [32]byte, bscTxHash [32]byte, swapType string, fee *big.Int, exchange *big.Int) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.contract.Transact(opts, "fillMTS2TOTMPegIn", recipient, dataHash, bscTxHash, swapType, fee, exchange)
}

// FillMTS2TOTMPegIn is a paid mutator transaction binding the contract method 0x6dfa08de.
//
// Solidity: function fillMTS2TOTMPegIn(address recipient, bytes32 dataHash, bytes32 bscTxHash, string swapType, uint256 fee, uint256 exchange) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) FillMTS2TOTMPegIn(recipient common.Address, dataHash [32]byte, bscTxHash [32]byte, swapType string, fee *big.Int, exchange *big.Int) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.FillMTS2TOTMPegIn(&_MetisBridgeAgentImpl.TransactOpts, recipient, dataHash, bscTxHash, swapType, fee, exchange)
}

// FillMTS2TOTMPegIn is a paid mutator transaction binding the contract method 0x6dfa08de.
//
// Solidity: function fillMTS2TOTMPegIn(address recipient, bytes32 dataHash, bytes32 bscTxHash, string swapType, uint256 fee, uint256 exchange) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorSession) FillMTS2TOTMPegIn(recipient common.Address, dataHash [32]byte, bscTxHash [32]byte, swapType string, fee *big.Int, exchange *big.Int) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.FillMTS2TOTMPegIn(&_MetisBridgeAgentImpl.TransactOpts, recipient, dataHash, bscTxHash, swapType, fee, exchange)
}

// MtsWithdraw is a paid mutator transaction binding the contract method 0x638ff912.
//
// Solidity: function mtsWithdraw() returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactor) MtsWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.contract.Transact(opts, "mtsWithdraw")
}

// MtsWithdraw is a paid mutator transaction binding the contract method 0x638ff912.
//
// Solidity: function mtsWithdraw() returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) MtsWithdraw() (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.MtsWithdraw(&_MetisBridgeAgentImpl.TransactOpts)
}

// MtsWithdraw is a paid mutator transaction binding the contract method 0x638ff912.
//
// Solidity: function mtsWithdraw() returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorSession) MtsWithdraw() (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.MtsWithdraw(&_MetisBridgeAgentImpl.TransactOpts)
}

// PegInMTS2TOTM is a paid mutator transaction binding the contract method 0x793ce93e.
//
// Solidity: function pegInMTS2TOTM((string,string,string,uint256,uint256,uint256,uint256,uint256,address) message, bytes signature) payable returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactor) PegInMTS2TOTM(opts *bind.TransactOpts, message IMessageStructureMessage, signature []byte) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.contract.Transact(opts, "pegInMTS2TOTM", message, signature)
}

// PegInMTS2TOTM is a paid mutator transaction binding the contract method 0x793ce93e.
//
// Solidity: function pegInMTS2TOTM((string,string,string,uint256,uint256,uint256,uint256,uint256,address) message, bytes signature) payable returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) PegInMTS2TOTM(message IMessageStructureMessage, signature []byte) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.PegInMTS2TOTM(&_MetisBridgeAgentImpl.TransactOpts, message, signature)
}

// PegInMTS2TOTM is a paid mutator transaction binding the contract method 0x793ce93e.
//
// Solidity: function pegInMTS2TOTM((string,string,string,uint256,uint256,uint256,uint256,uint256,address) message, bytes signature) payable returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorSession) PegInMTS2TOTM(message IMessageStructureMessage, signature []byte) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.PegInMTS2TOTM(&_MetisBridgeAgentImpl.TransactOpts, message, signature)
}

// SetClaim is a paid mutator transaction binding the contract method 0x18b4f71b.
//
// Solidity: function setClaim(address account, bytes32 dataHash, bool claimStat) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactor) SetClaim(opts *bind.TransactOpts, account common.Address, dataHash [32]byte, claimStat bool) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.contract.Transact(opts, "setClaim", account, dataHash, claimStat)
}

// SetClaim is a paid mutator transaction binding the contract method 0x18b4f71b.
//
// Solidity: function setClaim(address account, bytes32 dataHash, bool claimStat) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) SetClaim(account common.Address, dataHash [32]byte, claimStat bool) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.SetClaim(&_MetisBridgeAgentImpl.TransactOpts, account, dataHash, claimStat)
}

// SetClaim is a paid mutator transaction binding the contract method 0x18b4f71b.
//
// Solidity: function setClaim(address account, bytes32 dataHash, bool claimStat) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorSession) SetClaim(account common.Address, dataHash [32]byte, claimStat bool) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.SetClaim(&_MetisBridgeAgentImpl.TransactOpts, account, dataHash, claimStat)
}

// SetFill is a paid mutator transaction binding the contract method 0x64955f0c.
//
// Solidity: function setFill(address account, bytes32 dataHash, bool fillStat) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactor) SetFill(opts *bind.TransactOpts, account common.Address, dataHash [32]byte, fillStat bool) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.contract.Transact(opts, "setFill", account, dataHash, fillStat)
}

// SetFill is a paid mutator transaction binding the contract method 0x64955f0c.
//
// Solidity: function setFill(address account, bytes32 dataHash, bool fillStat) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) SetFill(account common.Address, dataHash [32]byte, fillStat bool) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.SetFill(&_MetisBridgeAgentImpl.TransactOpts, account, dataHash, fillStat)
}

// SetFill is a paid mutator transaction binding the contract method 0x64955f0c.
//
// Solidity: function setFill(address account, bytes32 dataHash, bool fillStat) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorSession) SetFill(account common.Address, dataHash [32]byte, fillStat bool) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.SetFill(&_MetisBridgeAgentImpl.TransactOpts, account, dataHash, fillStat)
}

// SetNonce is a paid mutator transaction binding the contract method 0x1d79f325.
//
// Solidity: function setNonce(address account, uint256 nonce) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactor) SetNonce(opts *bind.TransactOpts, account common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.contract.Transact(opts, "setNonce", account, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0x1d79f325.
//
// Solidity: function setNonce(address account, uint256 nonce) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplSession) SetNonce(account common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.SetNonce(&_MetisBridgeAgentImpl.TransactOpts, account, nonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0x1d79f325.
//
// Solidity: function setNonce(address account, uint256 nonce) returns()
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplTransactorSession) SetNonce(account common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _MetisBridgeAgentImpl.Contract.SetNonce(&_MetisBridgeAgentImpl.TransactOpts, account, nonce)
}

// MetisBridgeAgentImplSwapClaimedIterator is returned from FilterSwapClaimed and is used to iterate over the raw logs and unpacked data for SwapClaimed events raised by the MetisBridgeAgentImpl contract.
type MetisBridgeAgentImplSwapClaimedIterator struct {
	Event *MetisBridgeAgentImplSwapClaimed // Event containing the contract specifics and raw log

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
func (it *MetisBridgeAgentImplSwapClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetisBridgeAgentImplSwapClaimed)
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
		it.Event = new(MetisBridgeAgentImplSwapClaimed)
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
func (it *MetisBridgeAgentImplSwapClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetisBridgeAgentImplSwapClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetisBridgeAgentImplSwapClaimed represents a SwapClaimed event raised by the MetisBridgeAgentImpl contract.
type MetisBridgeAgentImplSwapClaimed struct {
	Recipient common.Address
	DataHash  [32]byte
	SwapType  string
	Exchange  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapClaimed is a free log retrieval operation binding the contract event 0xc4da37a5927e9eef1c8037e3357e69b90e1b899efe166d7692425c19b4ff25a3.
//
// Solidity: event swapClaimed(address indexed recipient, bytes32 indexed dataHash, string swapType, uint256 exchange)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) FilterSwapClaimed(opts *bind.FilterOpts, recipient []common.Address, dataHash [][32]byte) (*MetisBridgeAgentImplSwapClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _MetisBridgeAgentImpl.contract.FilterLogs(opts, "swapClaimed", recipientRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return &MetisBridgeAgentImplSwapClaimedIterator{contract: _MetisBridgeAgentImpl.contract, event: "swapClaimed", logs: logs, sub: sub}, nil
}

// WatchSwapClaimed is a free log subscription operation binding the contract event 0xc4da37a5927e9eef1c8037e3357e69b90e1b899efe166d7692425c19b4ff25a3.
//
// Solidity: event swapClaimed(address indexed recipient, bytes32 indexed dataHash, string swapType, uint256 exchange)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) WatchSwapClaimed(opts *bind.WatchOpts, sink chan<- *MetisBridgeAgentImplSwapClaimed, recipient []common.Address, dataHash [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _MetisBridgeAgentImpl.contract.WatchLogs(opts, "swapClaimed", recipientRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetisBridgeAgentImplSwapClaimed)
				if err := _MetisBridgeAgentImpl.contract.UnpackLog(event, "swapClaimed", log); err != nil {
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

// ParseSwapClaimed is a log parse operation binding the contract event 0xc4da37a5927e9eef1c8037e3357e69b90e1b899efe166d7692425c19b4ff25a3.
//
// Solidity: event swapClaimed(address indexed recipient, bytes32 indexed dataHash, string swapType, uint256 exchange)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) ParseSwapClaimed(log types.Log) (*MetisBridgeAgentImplSwapClaimed, error) {
	event := new(MetisBridgeAgentImplSwapClaimed)
	if err := _MetisBridgeAgentImpl.contract.UnpackLog(event, "swapClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetisBridgeAgentImplSwapFilledIterator is returned from FilterSwapFilled and is used to iterate over the raw logs and unpacked data for SwapFilled events raised by the MetisBridgeAgentImpl contract.
type MetisBridgeAgentImplSwapFilledIterator struct {
	Event *MetisBridgeAgentImplSwapFilled // Event containing the contract specifics and raw log

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
func (it *MetisBridgeAgentImplSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetisBridgeAgentImplSwapFilled)
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
		it.Event = new(MetisBridgeAgentImplSwapFilled)
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
func (it *MetisBridgeAgentImplSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetisBridgeAgentImplSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetisBridgeAgentImplSwapFilled represents a SwapFilled event raised by the MetisBridgeAgentImpl contract.
type MetisBridgeAgentImplSwapFilled struct {
	Recipient common.Address
	BscTxHash [32]byte
	DataHash  [32]byte
	SwapType  string
	Fee       *big.Int
	Exchange  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapFilled is a free log retrieval operation binding the contract event 0x763c79ae59cfda0e593247e48708dc5d5f033caae09822b9ec6ab8e5ebd5dfb1.
//
// Solidity: event swapFilled(address indexed recipient, bytes32 indexed bscTxHash, bytes32 indexed dataHash, string swapType, uint256 fee, uint256 exchange)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) FilterSwapFilled(opts *bind.FilterOpts, recipient []common.Address, bscTxHash [][32]byte, dataHash [][32]byte) (*MetisBridgeAgentImplSwapFilledIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var bscTxHashRule []interface{}
	for _, bscTxHashItem := range bscTxHash {
		bscTxHashRule = append(bscTxHashRule, bscTxHashItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _MetisBridgeAgentImpl.contract.FilterLogs(opts, "swapFilled", recipientRule, bscTxHashRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return &MetisBridgeAgentImplSwapFilledIterator{contract: _MetisBridgeAgentImpl.contract, event: "swapFilled", logs: logs, sub: sub}, nil
}

// WatchSwapFilled is a free log subscription operation binding the contract event 0x763c79ae59cfda0e593247e48708dc5d5f033caae09822b9ec6ab8e5ebd5dfb1.
//
// Solidity: event swapFilled(address indexed recipient, bytes32 indexed bscTxHash, bytes32 indexed dataHash, string swapType, uint256 fee, uint256 exchange)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) WatchSwapFilled(opts *bind.WatchOpts, sink chan<- *MetisBridgeAgentImplSwapFilled, recipient []common.Address, bscTxHash [][32]byte, dataHash [][32]byte) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var bscTxHashRule []interface{}
	for _, bscTxHashItem := range bscTxHash {
		bscTxHashRule = append(bscTxHashRule, bscTxHashItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _MetisBridgeAgentImpl.contract.WatchLogs(opts, "swapFilled", recipientRule, bscTxHashRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetisBridgeAgentImplSwapFilled)
				if err := _MetisBridgeAgentImpl.contract.UnpackLog(event, "swapFilled", log); err != nil {
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

// ParseSwapFilled is a log parse operation binding the contract event 0x763c79ae59cfda0e593247e48708dc5d5f033caae09822b9ec6ab8e5ebd5dfb1.
//
// Solidity: event swapFilled(address indexed recipient, bytes32 indexed bscTxHash, bytes32 indexed dataHash, string swapType, uint256 fee, uint256 exchange)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) ParseSwapFilled(log types.Log) (*MetisBridgeAgentImplSwapFilled, error) {
	event := new(MetisBridgeAgentImplSwapFilled)
	if err := _MetisBridgeAgentImpl.contract.UnpackLog(event, "swapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetisBridgeAgentImplSwapStartedIterator is returned from FilterSwapStarted and is used to iterate over the raw logs and unpacked data for SwapStarted events raised by the MetisBridgeAgentImpl contract.
type MetisBridgeAgentImplSwapStartedIterator struct {
	Event *MetisBridgeAgentImplSwapStarted // Event containing the contract specifics and raw log

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
func (it *MetisBridgeAgentImplSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetisBridgeAgentImplSwapStarted)
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
		it.Event = new(MetisBridgeAgentImplSwapStarted)
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
func (it *MetisBridgeAgentImplSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetisBridgeAgentImplSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetisBridgeAgentImplSwapStarted represents a SwapStarted event raised by the MetisBridgeAgentImpl contract.
type MetisBridgeAgentImplSwapStarted struct {
	Spender  common.Address
	DataHash [32]byte
	SwapType string
	Base     string
	Quote    string
	Amount   *big.Int
	Fee      *big.Int
	Exchange *big.Int
	Nonce    *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSwapStarted is a free log retrieval operation binding the contract event 0xddbc9ff29a9285afc6faa85774fd182b7d957a0096a14716b88d020717071b1c.
//
// Solidity: event swapStarted(address indexed spender, bytes32 indexed dataHash, string swapType, string base, string quote, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce, uint256 deadline)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) FilterSwapStarted(opts *bind.FilterOpts, spender []common.Address, dataHash [][32]byte) (*MetisBridgeAgentImplSwapStartedIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _MetisBridgeAgentImpl.contract.FilterLogs(opts, "swapStarted", spenderRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return &MetisBridgeAgentImplSwapStartedIterator{contract: _MetisBridgeAgentImpl.contract, event: "swapStarted", logs: logs, sub: sub}, nil
}

// WatchSwapStarted is a free log subscription operation binding the contract event 0xddbc9ff29a9285afc6faa85774fd182b7d957a0096a14716b88d020717071b1c.
//
// Solidity: event swapStarted(address indexed spender, bytes32 indexed dataHash, string swapType, string base, string quote, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce, uint256 deadline)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) WatchSwapStarted(opts *bind.WatchOpts, sink chan<- *MetisBridgeAgentImplSwapStarted, spender []common.Address, dataHash [][32]byte) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var dataHashRule []interface{}
	for _, dataHashItem := range dataHash {
		dataHashRule = append(dataHashRule, dataHashItem)
	}

	logs, sub, err := _MetisBridgeAgentImpl.contract.WatchLogs(opts, "swapStarted", spenderRule, dataHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetisBridgeAgentImplSwapStarted)
				if err := _MetisBridgeAgentImpl.contract.UnpackLog(event, "swapStarted", log); err != nil {
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
// Solidity: event swapStarted(address indexed spender, bytes32 indexed dataHash, string swapType, string base, string quote, uint256 amount, uint256 fee, uint256 exchange, uint256 nonce, uint256 deadline)
func (_MetisBridgeAgentImpl *MetisBridgeAgentImplFilterer) ParseSwapStarted(log types.Log) (*MetisBridgeAgentImplSwapStarted, error) {
	event := new(MetisBridgeAgentImplSwapStarted)
	if err := _MetisBridgeAgentImpl.contract.UnpackLog(event, "swapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
