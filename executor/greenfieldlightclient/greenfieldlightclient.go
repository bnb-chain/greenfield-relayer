// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package greenfieldlightclient

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
	_ = abi.ConvertType
)

// GreenfieldlightclientMetaData contains all meta data concerning the Greenfieldlightclient contract.
var GreenfieldlightclientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"APP_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBKEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PACKAGE_VERIFY_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORSET_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsPubKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relayers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_header\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_blsPubKeys\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"_relayers\",\"type\":\"address[]\"}],\"name\":\"syncTendermintHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pkgKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetBitMap\",\"type\":\"uint256\"}],\"name\":\"verifyPackage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GreenfieldlightclientABI is the input ABI used to generate the binding from.
// Deprecated: Use GreenfieldlightclientMetaData.ABI instead.
var GreenfieldlightclientABI = GreenfieldlightclientMetaData.ABI

// Greenfieldlightclient is an auto generated Go binding around an Ethereum contract.
type Greenfieldlightclient struct {
	GreenfieldlightclientCaller     // Read-only binding to the contract
	GreenfieldlightclientTransactor // Write-only binding to the contract
	GreenfieldlightclientFilterer   // Log filterer for contract events
}

// GreenfieldlightclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type GreenfieldlightclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreenfieldlightclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GreenfieldlightclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreenfieldlightclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GreenfieldlightclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreenfieldlightclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GreenfieldlightclientSession struct {
	Contract     *Greenfieldlightclient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GreenfieldlightclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GreenfieldlightclientCallerSession struct {
	Contract *GreenfieldlightclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GreenfieldlightclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GreenfieldlightclientTransactorSession struct {
	Contract     *GreenfieldlightclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GreenfieldlightclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type GreenfieldlightclientRaw struct {
	Contract *Greenfieldlightclient // Generic contract binding to access the raw methods on
}

// GreenfieldlightclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GreenfieldlightclientCallerRaw struct {
	Contract *GreenfieldlightclientCaller // Generic read-only contract binding to access the raw methods on
}

// GreenfieldlightclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GreenfieldlightclientTransactorRaw struct {
	Contract *GreenfieldlightclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGreenfieldlightclient creates a new instance of Greenfieldlightclient, bound to a specific deployed contract.
func NewGreenfieldlightclient(address common.Address, backend bind.ContractBackend) (*Greenfieldlightclient, error) {
	contract, err := bindGreenfieldlightclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Greenfieldlightclient{GreenfieldlightclientCaller: GreenfieldlightclientCaller{contract: contract}, GreenfieldlightclientTransactor: GreenfieldlightclientTransactor{contract: contract}, GreenfieldlightclientFilterer: GreenfieldlightclientFilterer{contract: contract}}, nil
}

// NewGreenfieldlightclientCaller creates a new read-only instance of Greenfieldlightclient, bound to a specific deployed contract.
func NewGreenfieldlightclientCaller(address common.Address, caller bind.ContractCaller) (*GreenfieldlightclientCaller, error) {
	contract, err := bindGreenfieldlightclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientCaller{contract: contract}, nil
}

// NewGreenfieldlightclientTransactor creates a new write-only instance of Greenfieldlightclient, bound to a specific deployed contract.
func NewGreenfieldlightclientTransactor(address common.Address, transactor bind.ContractTransactor) (*GreenfieldlightclientTransactor, error) {
	contract, err := bindGreenfieldlightclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientTransactor{contract: contract}, nil
}

// NewGreenfieldlightclientFilterer creates a new log filterer instance of Greenfieldlightclient, bound to a specific deployed contract.
func NewGreenfieldlightclientFilterer(address common.Address, filterer bind.ContractFilterer) (*GreenfieldlightclientFilterer, error) {
	contract, err := bindGreenfieldlightclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientFilterer{contract: contract}, nil
}

// bindGreenfieldlightclient binds a generic wrapper to an already deployed contract.
func bindGreenfieldlightclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GreenfieldlightclientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Greenfieldlightclient *GreenfieldlightclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Greenfieldlightclient.Contract.GreenfieldlightclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Greenfieldlightclient *GreenfieldlightclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.GreenfieldlightclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Greenfieldlightclient *GreenfieldlightclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.GreenfieldlightclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Greenfieldlightclient *GreenfieldlightclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Greenfieldlightclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Greenfieldlightclient *GreenfieldlightclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Greenfieldlightclient *GreenfieldlightclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.contract.Transact(opts, method, params...)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) APPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "APP_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) APPCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.APPCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) APPCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.APPCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) BLSPUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "BLS_PUBKEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) BLSPUBKEYLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.BLSPUBKEYLENGTH(&_Greenfieldlightclient.CallOpts)
}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) BLSPUBKEYLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.BLSPUBKEYLENGTH(&_Greenfieldlightclient.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Greenfieldlightclient *GreenfieldlightclientSession) CODEOK() (uint32, error) {
	return _Greenfieldlightclient.Contract.CODEOK(&_Greenfieldlightclient.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) CODEOK() (uint32, error) {
	return _Greenfieldlightclient.Contract.CODEOK(&_Greenfieldlightclient.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Greenfieldlightclient *GreenfieldlightclientSession) ERRORFAILDECODE() (uint32, error) {
	return _Greenfieldlightclient.Contract.ERRORFAILDECODE(&_Greenfieldlightclient.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Greenfieldlightclient.Contract.ERRORFAILDECODE(&_Greenfieldlightclient.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GOVCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.GOVCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GOVCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.GOVCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) LIGHTCLIENTCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "LIGHT_CLIENT_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) LIGHTCLIENTCONTRACT() (common.Address, error) {
	return _Greenfieldlightclient.Contract.LIGHTCLIENTCONTRACT(&_Greenfieldlightclient.CallOpts)
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) LIGHTCLIENTCONTRACT() (common.Address, error) {
	return _Greenfieldlightclient.Contract.LIGHTCLIENTCONTRACT(&_Greenfieldlightclient.CallOpts)
}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) PACKAGEVERIFYCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "PACKAGE_VERIFY_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) PACKAGEVERIFYCONTRACT() (common.Address, error) {
	return _Greenfieldlightclient.Contract.PACKAGEVERIFYCONTRACT(&_Greenfieldlightclient.CallOpts)
}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) PACKAGEVERIFYCONTRACT() (common.Address, error) {
	return _Greenfieldlightclient.Contract.PACKAGEVERIFYCONTRACT(&_Greenfieldlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.TRANSFERINCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.TRANSFERINCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.TRANSFEROUTCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.TRANSFEROUTCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VALIDATORSETCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "VALIDATORSET_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) VALIDATORSETCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.VALIDATORSETCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VALIDATORSETCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.VALIDATORSETCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) BlsPubKeys(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "blsPubKeys")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientSession) BlsPubKeys() ([]byte, error) {
	return _Greenfieldlightclient.Contract.BlsPubKeys(&_Greenfieldlightclient.CallOpts)
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) BlsPubKeys() ([]byte, error) {
	return _Greenfieldlightclient.Contract.BlsPubKeys(&_Greenfieldlightclient.CallOpts)
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GetRelayers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "getRelayers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Greenfieldlightclient *GreenfieldlightclientSession) GetRelayers() ([]common.Address, error) {
	return _Greenfieldlightclient.Contract.GetRelayers(&_Greenfieldlightclient.CallOpts)
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GetRelayers() ([]common.Address, error) {
	return _Greenfieldlightclient.Contract.GetRelayers(&_Greenfieldlightclient.CallOpts)
}

// InsHeight is a free data retrieval call binding the contract method 0xc7234309.
//
// Solidity: function insHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) InsHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "insHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InsHeight is a free data retrieval call binding the contract method 0xc7234309.
//
// Solidity: function insHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientSession) InsHeight() (uint64, error) {
	return _Greenfieldlightclient.Contract.InsHeight(&_Greenfieldlightclient.CallOpts)
}

// InsHeight is a free data retrieval call binding the contract method 0xc7234309.
//
// Solidity: function insHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) InsHeight() (uint64, error) {
	return _Greenfieldlightclient.Contract.InsHeight(&_Greenfieldlightclient.CallOpts)
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) Relayers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "relayers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) Relayers(arg0 *big.Int) (common.Address, error) {
	return _Greenfieldlightclient.Contract.Relayers(&_Greenfieldlightclient.CallOpts, arg0)
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) Relayers(arg0 *big.Int) (common.Address, error) {
	return _Greenfieldlightclient.Contract.Relayers(&_Greenfieldlightclient.CallOpts, arg0)
}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VerifyPackage(opts *bind.CallOpts, _pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "verifyPackage", _pkgKey, _payload, _blsSignature, _validatorSetBitMap)

	if err != nil {
		return err
	}

	return err

}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Greenfieldlightclient *GreenfieldlightclientSession) VerifyPackage(_pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	return _Greenfieldlightclient.Contract.VerifyPackage(&_Greenfieldlightclient.CallOpts, _pkgKey, _payload, _blsSignature, _validatorSetBitMap)
}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VerifyPackage(_pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	return _Greenfieldlightclient.Contract.VerifyPackage(&_Greenfieldlightclient.CallOpts, _pkgKey, _payload, _blsSignature, _validatorSetBitMap)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Greenfieldlightclient *GreenfieldlightclientTransactor) SyncTendermintHeader(opts *bind.TransactOpts, _header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Greenfieldlightclient.contract.Transact(opts, "syncTendermintHeader", _header, _height, _blsPubKeys, _relayers)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Greenfieldlightclient *GreenfieldlightclientSession) SyncTendermintHeader(_header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.SyncTendermintHeader(&_Greenfieldlightclient.TransactOpts, _header, _height, _blsPubKeys, _relayers)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Greenfieldlightclient *GreenfieldlightclientTransactorSession) SyncTendermintHeader(_header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.SyncTendermintHeader(&_Greenfieldlightclient.TransactOpts, _header, _height, _blsPubKeys, _relayers)
}
