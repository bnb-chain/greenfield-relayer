// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayerhub

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

// RelayerhubMetaData contains all meta data concerning the Relayerhub contract.
var RelayerhubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardToRelayer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BUCKET_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUCKET_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ConfigSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OBJECT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OBJECT_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROXY_ADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_RATIO_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RelayerhubABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayerhubMetaData.ABI instead.
var RelayerhubABI = RelayerhubMetaData.ABI

// Relayerhub is an auto generated Go binding around an Ethereum contract.
type Relayerhub struct {
	RelayerhubCaller     // Read-only binding to the contract
	RelayerhubTransactor // Write-only binding to the contract
	RelayerhubFilterer   // Log filterer for contract events
}

// RelayerhubCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerhubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerhubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerhubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerhubSession struct {
	Contract     *Relayerhub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerhubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerhubCallerSession struct {
	Contract *RelayerhubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RelayerhubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerhubTransactorSession struct {
	Contract     *RelayerhubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RelayerhubRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerhubRaw struct {
	Contract *Relayerhub // Generic contract binding to access the raw methods on
}

// RelayerhubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerhubCallerRaw struct {
	Contract *RelayerhubCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerhubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerhubTransactorRaw struct {
	Contract *RelayerhubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerhub creates a new instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhub(address common.Address, backend bind.ContractBackend) (*Relayerhub, error) {
	contract, err := bindRelayerhub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relayerhub{RelayerhubCaller: RelayerhubCaller{contract: contract}, RelayerhubTransactor: RelayerhubTransactor{contract: contract}, RelayerhubFilterer: RelayerhubFilterer{contract: contract}}, nil
}

// NewRelayerhubCaller creates a new read-only instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubCaller(address common.Address, caller bind.ContractCaller) (*RelayerhubCaller, error) {
	contract, err := bindRelayerhub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerhubCaller{contract: contract}, nil
}

// NewRelayerhubTransactor creates a new write-only instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerhubTransactor, error) {
	contract, err := bindRelayerhub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerhubTransactor{contract: contract}, nil
}

// NewRelayerhubFilterer creates a new log filterer instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerhubFilterer, error) {
	contract, err := bindRelayerhub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerhubFilterer{contract: contract}, nil
}

// bindRelayerhub binds a generic wrapper to an already deployed contract.
func bindRelayerhub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RelayerhubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayerhub *RelayerhubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relayerhub.Contract.RelayerhubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerhub *RelayerhubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.Contract.RelayerhubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerhub *RelayerhubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerhub.Contract.RelayerhubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayerhub *RelayerhubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Relayerhub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerhub *RelayerhubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerhub *RelayerhubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerhub.Contract.contract.Transact(opts, method, params...)
}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCaller) BUCKETCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "BUCKET_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubSession) BUCKETCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.BUCKETCHANNELID(&_Relayerhub.CallOpts)
}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) BUCKETCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.BUCKETCHANNELID(&_Relayerhub.CallOpts)
}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Relayerhub *RelayerhubCaller) BUCKETHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "BUCKET_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Relayerhub *RelayerhubSession) BUCKETHUB() (common.Address, error) {
	return _Relayerhub.Contract.BUCKETHUB(&_Relayerhub.CallOpts)
}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) BUCKETHUB() (common.Address, error) {
	return _Relayerhub.Contract.BUCKETHUB(&_Relayerhub.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Relayerhub *RelayerhubCaller) CROSSCHAIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "CROSS_CHAIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Relayerhub *RelayerhubSession) CROSSCHAIN() (common.Address, error) {
	return _Relayerhub.Contract.CROSSCHAIN(&_Relayerhub.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) CROSSCHAIN() (common.Address, error) {
	return _Relayerhub.Contract.CROSSCHAIN(&_Relayerhub.CallOpts)
}

// ConfigSlots is a free data retrieval call binding the contract method 0xb76e4aca.
//
// Solidity: function ConfigSlots(uint256 ) view returns(uint256)
func (_Relayerhub *RelayerhubCaller) ConfigSlots(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "ConfigSlots", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfigSlots is a free data retrieval call binding the contract method 0xb76e4aca.
//
// Solidity: function ConfigSlots(uint256 ) view returns(uint256)
func (_Relayerhub *RelayerhubSession) ConfigSlots(arg0 *big.Int) (*big.Int, error) {
	return _Relayerhub.Contract.ConfigSlots(&_Relayerhub.CallOpts, arg0)
}

// ConfigSlots is a free data retrieval call binding the contract method 0xb76e4aca.
//
// Solidity: function ConfigSlots(uint256 ) view returns(uint256)
func (_Relayerhub *RelayerhubCallerSession) ConfigSlots(arg0 *big.Int) (*big.Int, error) {
	return _Relayerhub.Contract.ConfigSlots(&_Relayerhub.CallOpts, arg0)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "GOV_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubSession) GOVCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.GOVCHANNELID(&_Relayerhub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) GOVCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.GOVCHANNELID(&_Relayerhub.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Relayerhub *RelayerhubCaller) GOVHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "GOV_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Relayerhub *RelayerhubSession) GOVHUB() (common.Address, error) {
	return _Relayerhub.Contract.GOVHUB(&_Relayerhub.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) GOVHUB() (common.Address, error) {
	return _Relayerhub.Contract.GOVHUB(&_Relayerhub.CallOpts)
}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCaller) GROUPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "GROUP_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubSession) GROUPCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.GROUPCHANNELID(&_Relayerhub.CallOpts)
}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) GROUPCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.GROUPCHANNELID(&_Relayerhub.CallOpts)
}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Relayerhub *RelayerhubCaller) GROUPHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "GROUP_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Relayerhub *RelayerhubSession) GROUPHUB() (common.Address, error) {
	return _Relayerhub.Contract.GROUPHUB(&_Relayerhub.CallOpts)
}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) GROUPHUB() (common.Address, error) {
	return _Relayerhub.Contract.GROUPHUB(&_Relayerhub.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Relayerhub *RelayerhubCaller) LIGHTCLIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "LIGHT_CLIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Relayerhub *RelayerhubSession) LIGHTCLIENT() (common.Address, error) {
	return _Relayerhub.Contract.LIGHTCLIENT(&_Relayerhub.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) LIGHTCLIENT() (common.Address, error) {
	return _Relayerhub.Contract.LIGHTCLIENT(&_Relayerhub.CallOpts)
}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCaller) OBJECTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "OBJECT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubSession) OBJECTCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.OBJECTCHANNELID(&_Relayerhub.CallOpts)
}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) OBJECTCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.OBJECTCHANNELID(&_Relayerhub.CallOpts)
}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Relayerhub *RelayerhubCaller) OBJECTHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "OBJECT_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Relayerhub *RelayerhubSession) OBJECTHUB() (common.Address, error) {
	return _Relayerhub.Contract.OBJECTHUB(&_Relayerhub.CallOpts)
}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) OBJECTHUB() (common.Address, error) {
	return _Relayerhub.Contract.OBJECTHUB(&_Relayerhub.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Relayerhub *RelayerhubCaller) PROXYADMIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "PROXY_ADMIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Relayerhub *RelayerhubSession) PROXYADMIN() (common.Address, error) {
	return _Relayerhub.Contract.PROXYADMIN(&_Relayerhub.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) PROXYADMIN() (common.Address, error) {
	return _Relayerhub.Contract.PROXYADMIN(&_Relayerhub.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Relayerhub *RelayerhubCaller) RELAYERHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "RELAYER_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Relayerhub *RelayerhubSession) RELAYERHUB() (common.Address, error) {
	return _Relayerhub.Contract.RELAYERHUB(&_Relayerhub.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) RELAYERHUB() (common.Address, error) {
	return _Relayerhub.Contract.RELAYERHUB(&_Relayerhub.CallOpts)
}

// REWARDRATIOSCALE is a free data retrieval call binding the contract method 0x132f2adb.
//
// Solidity: function REWARD_RATIO_SCALE() view returns(uint256)
func (_Relayerhub *RelayerhubCaller) REWARDRATIOSCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "REWARD_RATIO_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDRATIOSCALE is a free data retrieval call binding the contract method 0x132f2adb.
//
// Solidity: function REWARD_RATIO_SCALE() view returns(uint256)
func (_Relayerhub *RelayerhubSession) REWARDRATIOSCALE() (*big.Int, error) {
	return _Relayerhub.Contract.REWARDRATIOSCALE(&_Relayerhub.CallOpts)
}

// REWARDRATIOSCALE is a free data retrieval call binding the contract method 0x132f2adb.
//
// Solidity: function REWARD_RATIO_SCALE() view returns(uint256)
func (_Relayerhub *RelayerhubCallerSession) REWARDRATIOSCALE() (*big.Int, error) {
	return _Relayerhub.Contract.REWARDRATIOSCALE(&_Relayerhub.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Relayerhub *RelayerhubCaller) TOKENHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "TOKEN_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Relayerhub *RelayerhubSession) TOKENHUB() (common.Address, error) {
	return _Relayerhub.Contract.TOKENHUB(&_Relayerhub.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Relayerhub *RelayerhubCallerSession) TOKENHUB() (common.Address, error) {
	return _Relayerhub.Contract.TOKENHUB(&_Relayerhub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "TRANSFER_IN_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.TRANSFERINCHANNELID(&_Relayerhub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.TRANSFERINCHANNELID(&_Relayerhub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "TRANSFER_OUT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.TRANSFEROUTCHANNELID(&_Relayerhub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.TRANSFEROUTCHANNELID(&_Relayerhub.CallOpts)
}

// RewardMap is a free data retrieval call binding the contract method 0x83d44339.
//
// Solidity: function rewardMap(address ) view returns(uint256)
func (_Relayerhub *RelayerhubCaller) RewardMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "rewardMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardMap is a free data retrieval call binding the contract method 0x83d44339.
//
// Solidity: function rewardMap(address ) view returns(uint256)
func (_Relayerhub *RelayerhubSession) RewardMap(arg0 common.Address) (*big.Int, error) {
	return _Relayerhub.Contract.RewardMap(&_Relayerhub.CallOpts, arg0)
}

// RewardMap is a free data retrieval call binding the contract method 0x83d44339.
//
// Solidity: function rewardMap(address ) view returns(uint256)
func (_Relayerhub *RelayerhubCallerSession) RewardMap(arg0 common.Address) (*big.Int, error) {
	return _Relayerhub.Contract.RewardMap(&_Relayerhub.CallOpts, arg0)
}

// VersionInfo is a free data retrieval call binding the contract method 0xa21d1354.
//
// Solidity: function versionInfo() pure returns(uint256 version, string name, string description)
func (_Relayerhub *RelayerhubCaller) VersionInfo(opts *bind.CallOpts) (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	var out []interface{}
	err := _Relayerhub.contract.Call(opts, &out, "versionInfo")

	outstruct := new(struct {
		Version     *big.Int
		Name        string
		Description string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// VersionInfo is a free data retrieval call binding the contract method 0xa21d1354.
//
// Solidity: function versionInfo() pure returns(uint256 version, string name, string description)
func (_Relayerhub *RelayerhubSession) VersionInfo() (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	return _Relayerhub.Contract.VersionInfo(&_Relayerhub.CallOpts)
}

// VersionInfo is a free data retrieval call binding the contract method 0xa21d1354.
//
// Solidity: function versionInfo() pure returns(uint256 version, string name, string description)
func (_Relayerhub *RelayerhubCallerSession) VersionInfo() (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	return _Relayerhub.Contract.VersionInfo(&_Relayerhub.CallOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 _reward) returns()
func (_Relayerhub *RelayerhubTransactor) AddReward(opts *bind.TransactOpts, _relayer common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "addReward", _relayer, _reward)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 _reward) returns()
func (_Relayerhub *RelayerhubSession) AddReward(_relayer common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _Relayerhub.Contract.AddReward(&_Relayerhub.TransactOpts, _relayer, _reward)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _relayer, uint256 _reward) returns()
func (_Relayerhub *RelayerhubTransactorSession) AddReward(_relayer common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _Relayerhub.Contract.AddReward(&_Relayerhub.TransactOpts, _relayer, _reward)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _relayer) returns()
func (_Relayerhub *RelayerhubTransactor) ClaimReward(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "claimReward", _relayer)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _relayer) returns()
func (_Relayerhub *RelayerhubSession) ClaimReward(_relayer common.Address) (*types.Transaction, error) {
	return _Relayerhub.Contract.ClaimReward(&_Relayerhub.TransactOpts, _relayer)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address _relayer) returns()
func (_Relayerhub *RelayerhubTransactorSession) ClaimReward(_relayer common.Address) (*types.Transaction, error) {
	return _Relayerhub.Contract.ClaimReward(&_Relayerhub.TransactOpts, _relayer)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Relayerhub *RelayerhubTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Relayerhub *RelayerhubSession) Initialize() (*types.Transaction, error) {
	return _Relayerhub.Contract.Initialize(&_Relayerhub.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Relayerhub *RelayerhubTransactorSession) Initialize() (*types.Transaction, error) {
	return _Relayerhub.Contract.Initialize(&_Relayerhub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Relayerhub *RelayerhubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Relayerhub *RelayerhubSession) Receive() (*types.Transaction, error) {
	return _Relayerhub.Contract.Receive(&_Relayerhub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Relayerhub *RelayerhubTransactorSession) Receive() (*types.Transaction, error) {
	return _Relayerhub.Contract.Receive(&_Relayerhub.TransactOpts)
}

// RelayerhubClaimedRewardIterator is returned from FilterClaimedReward and is used to iterate over the raw logs and unpacked data for ClaimedReward events raised by the Relayerhub contract.
type RelayerhubClaimedRewardIterator struct {
	Event *RelayerhubClaimedReward // Event containing the contract specifics and raw log

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
func (it *RelayerhubClaimedRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubClaimedReward)
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
		it.Event = new(RelayerhubClaimedReward)
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
func (it *RelayerhubClaimedRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubClaimedRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubClaimedReward represents a ClaimedReward event raised by the Relayerhub contract.
type RelayerhubClaimedReward struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedReward is a free log retrieval operation binding the contract event 0xd0813ff03c470dcc7baa9ce36914dc2febdfd276d639deffaac383fd3db42ba3.
//
// Solidity: event ClaimedReward(address relayer, uint256 amount)
func (_Relayerhub *RelayerhubFilterer) FilterClaimedReward(opts *bind.FilterOpts) (*RelayerhubClaimedRewardIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "ClaimedReward")
	if err != nil {
		return nil, err
	}
	return &RelayerhubClaimedRewardIterator{contract: _Relayerhub.contract, event: "ClaimedReward", logs: logs, sub: sub}, nil
}

// WatchClaimedReward is a free log subscription operation binding the contract event 0xd0813ff03c470dcc7baa9ce36914dc2febdfd276d639deffaac383fd3db42ba3.
//
// Solidity: event ClaimedReward(address relayer, uint256 amount)
func (_Relayerhub *RelayerhubFilterer) WatchClaimedReward(opts *bind.WatchOpts, sink chan<- *RelayerhubClaimedReward) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "ClaimedReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubClaimedReward)
				if err := _Relayerhub.contract.UnpackLog(event, "ClaimedReward", log); err != nil {
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

// ParseClaimedReward is a log parse operation binding the contract event 0xd0813ff03c470dcc7baa9ce36914dc2febdfd276d639deffaac383fd3db42ba3.
//
// Solidity: event ClaimedReward(address relayer, uint256 amount)
func (_Relayerhub *RelayerhubFilterer) ParseClaimedReward(log types.Log) (*RelayerhubClaimedReward, error) {
	event := new(RelayerhubClaimedReward)
	if err := _Relayerhub.contract.UnpackLog(event, "ClaimedReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerhubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Relayerhub contract.
type RelayerhubInitializedIterator struct {
	Event *RelayerhubInitialized // Event containing the contract specifics and raw log

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
func (it *RelayerhubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubInitialized)
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
		it.Event = new(RelayerhubInitialized)
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
func (it *RelayerhubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubInitialized represents a Initialized event raised by the Relayerhub contract.
type RelayerhubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Relayerhub *RelayerhubFilterer) FilterInitialized(opts *bind.FilterOpts) (*RelayerhubInitializedIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RelayerhubInitializedIterator{contract: _Relayerhub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Relayerhub *RelayerhubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RelayerhubInitialized) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubInitialized)
				if err := _Relayerhub.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Relayerhub *RelayerhubFilterer) ParseInitialized(log types.Log) (*RelayerhubInitialized, error) {
	event := new(RelayerhubInitialized)
	if err := _Relayerhub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayerhubRewardToRelayerIterator is returned from FilterRewardToRelayer and is used to iterate over the raw logs and unpacked data for RewardToRelayer events raised by the Relayerhub contract.
type RelayerhubRewardToRelayerIterator struct {
	Event *RelayerhubRewardToRelayer // Event containing the contract specifics and raw log

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
func (it *RelayerhubRewardToRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubRewardToRelayer)
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
		it.Event = new(RelayerhubRewardToRelayer)
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
func (it *RelayerhubRewardToRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubRewardToRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubRewardToRelayer represents a RewardToRelayer event raised by the Relayerhub contract.
type RelayerhubRewardToRelayer struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardToRelayer is a free log retrieval operation binding the contract event 0xcc3341048e8fd1ed288bcd99bd6231605849b6301fe5ae9170850a29d9b1c2dd.
//
// Solidity: event RewardToRelayer(address relayer, uint256 amount)
func (_Relayerhub *RelayerhubFilterer) FilterRewardToRelayer(opts *bind.FilterOpts) (*RelayerhubRewardToRelayerIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "RewardToRelayer")
	if err != nil {
		return nil, err
	}
	return &RelayerhubRewardToRelayerIterator{contract: _Relayerhub.contract, event: "RewardToRelayer", logs: logs, sub: sub}, nil
}

// WatchRewardToRelayer is a free log subscription operation binding the contract event 0xcc3341048e8fd1ed288bcd99bd6231605849b6301fe5ae9170850a29d9b1c2dd.
//
// Solidity: event RewardToRelayer(address relayer, uint256 amount)
func (_Relayerhub *RelayerhubFilterer) WatchRewardToRelayer(opts *bind.WatchOpts, sink chan<- *RelayerhubRewardToRelayer) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "RewardToRelayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubRewardToRelayer)
				if err := _Relayerhub.contract.UnpackLog(event, "RewardToRelayer", log); err != nil {
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

// ParseRewardToRelayer is a log parse operation binding the contract event 0xcc3341048e8fd1ed288bcd99bd6231605849b6301fe5ae9170850a29d9b1c2dd.
//
// Solidity: event RewardToRelayer(address relayer, uint256 amount)
func (_Relayerhub *RelayerhubFilterer) ParseRewardToRelayer(log types.Log) (*RelayerhubRewardToRelayer, error) {
	event := new(RelayerhubRewardToRelayer)
	if err := _Relayerhub.contract.UnpackLog(event, "RewardToRelayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
