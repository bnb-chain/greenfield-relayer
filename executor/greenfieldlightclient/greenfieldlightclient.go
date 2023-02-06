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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"initConsensusState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"validatorSetChanged\",\"type\":\"bool\"}],\"name\":\"updateConsensusState\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"APP_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_SIGNATURE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_STATE_BASE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_STATE_BYTES_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_VALIDATE_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEIGHT_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSAGE_HASH_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PACKAGE_VERIFY_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROXY_ADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_BLS_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_BYTES_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_PUB_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_SET_HASH_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_VOTING_POWER_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsPubKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_blsPubKeys\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"consensusStateBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"height\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_initConsensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidatorSetHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"submitters\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_lightBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"syncLightBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorSet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"},{\"internalType\":\"int64\",\"name\":\"votingPower\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"relayerBlsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetBitMap\",\"type\":\"uint256\"}],\"name\":\"verifyPackage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// BLSSIGNATURELENGTH is a free data retrieval call binding the contract method 0xdfe5b236.
//
// Solidity: function BLS_SIGNATURE_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) BLSSIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "BLS_SIGNATURE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLSSIGNATURELENGTH is a free data retrieval call binding the contract method 0xdfe5b236.
//
// Solidity: function BLS_SIGNATURE_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) BLSSIGNATURELENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.BLSSIGNATURELENGTH(&_Greenfieldlightclient.CallOpts)
}

// BLSSIGNATURELENGTH is a free data retrieval call binding the contract method 0xdfe5b236.
//
// Solidity: function BLS_SIGNATURE_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) BLSSIGNATURELENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.BLSSIGNATURELENGTH(&_Greenfieldlightclient.CallOpts)
}

// CHAINIDLENGTH is a free data retrieval call binding the contract method 0x72ca1fe3.
//
// Solidity: function CHAIN_ID_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) CHAINIDLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "CHAIN_ID_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINIDLENGTH is a free data retrieval call binding the contract method 0x72ca1fe3.
//
// Solidity: function CHAIN_ID_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) CHAINIDLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.CHAINIDLENGTH(&_Greenfieldlightclient.CallOpts)
}

// CHAINIDLENGTH is a free data retrieval call binding the contract method 0x72ca1fe3.
//
// Solidity: function CHAIN_ID_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) CHAINIDLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.CHAINIDLENGTH(&_Greenfieldlightclient.CallOpts)
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

// CONSENSUSSTATEBASELENGTH is a free data retrieval call binding the contract method 0x7e8982bd.
//
// Solidity: function CONSENSUS_STATE_BASE_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) CONSENSUSSTATEBASELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "CONSENSUS_STATE_BASE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONSENSUSSTATEBASELENGTH is a free data retrieval call binding the contract method 0x7e8982bd.
//
// Solidity: function CONSENSUS_STATE_BASE_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) CONSENSUSSTATEBASELENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.CONSENSUSSTATEBASELENGTH(&_Greenfieldlightclient.CallOpts)
}

// CONSENSUSSTATEBASELENGTH is a free data retrieval call binding the contract method 0x7e8982bd.
//
// Solidity: function CONSENSUS_STATE_BASE_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) CONSENSUSSTATEBASELENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.CONSENSUSSTATEBASELENGTH(&_Greenfieldlightclient.CallOpts)
}

// CONSENSUSSTATEBYTESLENGTH is a free data retrieval call binding the contract method 0xdefd63f1.
//
// Solidity: function CONSENSUS_STATE_BYTES_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) CONSENSUSSTATEBYTESLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "CONSENSUS_STATE_BYTES_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONSENSUSSTATEBYTESLENGTH is a free data retrieval call binding the contract method 0xdefd63f1.
//
// Solidity: function CONSENSUS_STATE_BYTES_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) CONSENSUSSTATEBYTESLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.CONSENSUSSTATEBYTESLENGTH(&_Greenfieldlightclient.CallOpts)
}

// CONSENSUSSTATEBYTESLENGTH is a free data retrieval call binding the contract method 0xdefd63f1.
//
// Solidity: function CONSENSUS_STATE_BYTES_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) CONSENSUSSTATEBYTESLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.CONSENSUSSTATEBYTESLENGTH(&_Greenfieldlightclient.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) CROSSCHAIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "CROSS_CHAIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) CROSSCHAIN() (common.Address, error) {
	return _Greenfieldlightclient.Contract.CROSSCHAIN(&_Greenfieldlightclient.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) CROSSCHAIN() (common.Address, error) {
	return _Greenfieldlightclient.Contract.CROSSCHAIN(&_Greenfieldlightclient.CallOpts)
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

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GOVHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "GOV_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GOVHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.GOVHUB(&_Greenfieldlightclient.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GOVHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.GOVHUB(&_Greenfieldlightclient.CallOpts)
}

// HEADERVALIDATECONTRACT is a free data retrieval call binding the contract method 0x9c079c59.
//
// Solidity: function HEADER_VALIDATE_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) HEADERVALIDATECONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "HEADER_VALIDATE_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HEADERVALIDATECONTRACT is a free data retrieval call binding the contract method 0x9c079c59.
//
// Solidity: function HEADER_VALIDATE_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) HEADERVALIDATECONTRACT() (common.Address, error) {
	return _Greenfieldlightclient.Contract.HEADERVALIDATECONTRACT(&_Greenfieldlightclient.CallOpts)
}

// HEADERVALIDATECONTRACT is a free data retrieval call binding the contract method 0x9c079c59.
//
// Solidity: function HEADER_VALIDATE_CONTRACT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) HEADERVALIDATECONTRACT() (common.Address, error) {
	return _Greenfieldlightclient.Contract.HEADERVALIDATECONTRACT(&_Greenfieldlightclient.CallOpts)
}

// HEIGHTLENGTH is a free data retrieval call binding the contract method 0x83e3a498.
//
// Solidity: function HEIGHT_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) HEIGHTLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "HEIGHT_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HEIGHTLENGTH is a free data retrieval call binding the contract method 0x83e3a498.
//
// Solidity: function HEIGHT_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) HEIGHTLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.HEIGHTLENGTH(&_Greenfieldlightclient.CallOpts)
}

// HEIGHTLENGTH is a free data retrieval call binding the contract method 0x83e3a498.
//
// Solidity: function HEIGHT_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) HEIGHTLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.HEIGHTLENGTH(&_Greenfieldlightclient.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) LIGHTCLIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "LIGHT_CLIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) LIGHTCLIENT() (common.Address, error) {
	return _Greenfieldlightclient.Contract.LIGHTCLIENT(&_Greenfieldlightclient.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) LIGHTCLIENT() (common.Address, error) {
	return _Greenfieldlightclient.Contract.LIGHTCLIENT(&_Greenfieldlightclient.CallOpts)
}

// MESSAGEHASHLENGTH is a free data retrieval call binding the contract method 0x4de2b60a.
//
// Solidity: function MESSAGE_HASH_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) MESSAGEHASHLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "MESSAGE_HASH_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MESSAGEHASHLENGTH is a free data retrieval call binding the contract method 0x4de2b60a.
//
// Solidity: function MESSAGE_HASH_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) MESSAGEHASHLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.MESSAGEHASHLENGTH(&_Greenfieldlightclient.CallOpts)
}

// MESSAGEHASHLENGTH is a free data retrieval call binding the contract method 0x4de2b60a.
//
// Solidity: function MESSAGE_HASH_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) MESSAGEHASHLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.MESSAGEHASHLENGTH(&_Greenfieldlightclient.CallOpts)
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

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) PROXYADMIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "PROXY_ADMIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) PROXYADMIN() (common.Address, error) {
	return _Greenfieldlightclient.Contract.PROXYADMIN(&_Greenfieldlightclient.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) PROXYADMIN() (common.Address, error) {
	return _Greenfieldlightclient.Contract.PROXYADMIN(&_Greenfieldlightclient.CallOpts)
}

// RELAYERADDRESSLENGTH is a free data retrieval call binding the contract method 0x7896b4e4.
//
// Solidity: function RELAYER_ADDRESS_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) RELAYERADDRESSLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "RELAYER_ADDRESS_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELAYERADDRESSLENGTH is a free data retrieval call binding the contract method 0x7896b4e4.
//
// Solidity: function RELAYER_ADDRESS_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) RELAYERADDRESSLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.RELAYERADDRESSLENGTH(&_Greenfieldlightclient.CallOpts)
}

// RELAYERADDRESSLENGTH is a free data retrieval call binding the contract method 0x7896b4e4.
//
// Solidity: function RELAYER_ADDRESS_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) RELAYERADDRESSLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.RELAYERADDRESSLENGTH(&_Greenfieldlightclient.CallOpts)
}

// RELAYERBLSKEYLENGTH is a free data retrieval call binding the contract method 0x092e7ddb.
//
// Solidity: function RELAYER_BLS_KEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) RELAYERBLSKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "RELAYER_BLS_KEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELAYERBLSKEYLENGTH is a free data retrieval call binding the contract method 0x092e7ddb.
//
// Solidity: function RELAYER_BLS_KEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) RELAYERBLSKEYLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.RELAYERBLSKEYLENGTH(&_Greenfieldlightclient.CallOpts)
}

// RELAYERBLSKEYLENGTH is a free data retrieval call binding the contract method 0x092e7ddb.
//
// Solidity: function RELAYER_BLS_KEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) RELAYERBLSKEYLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.RELAYERBLSKEYLENGTH(&_Greenfieldlightclient.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) RELAYERHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "RELAYER_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) RELAYERHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.RELAYERHUB(&_Greenfieldlightclient.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) RELAYERHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.RELAYERHUB(&_Greenfieldlightclient.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) TOKENHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "TOKEN_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) TOKENHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.TOKENHUB(&_Greenfieldlightclient.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) TOKENHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.TOKENHUB(&_Greenfieldlightclient.CallOpts)
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

// VALIDATORBYTESLENGTH is a free data retrieval call binding the contract method 0x8272ecde.
//
// Solidity: function VALIDATOR_BYTES_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VALIDATORBYTESLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "VALIDATOR_BYTES_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORBYTESLENGTH is a free data retrieval call binding the contract method 0x8272ecde.
//
// Solidity: function VALIDATOR_BYTES_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) VALIDATORBYTESLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.VALIDATORBYTESLENGTH(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORBYTESLENGTH is a free data retrieval call binding the contract method 0x8272ecde.
//
// Solidity: function VALIDATOR_BYTES_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VALIDATORBYTESLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.VALIDATORBYTESLENGTH(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORPUBKEYLENGTH is a free data retrieval call binding the contract method 0x76b7416b.
//
// Solidity: function VALIDATOR_PUB_KEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VALIDATORPUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "VALIDATOR_PUB_KEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORPUBKEYLENGTH is a free data retrieval call binding the contract method 0x76b7416b.
//
// Solidity: function VALIDATOR_PUB_KEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) VALIDATORPUBKEYLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.VALIDATORPUBKEYLENGTH(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORPUBKEYLENGTH is a free data retrieval call binding the contract method 0x76b7416b.
//
// Solidity: function VALIDATOR_PUB_KEY_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VALIDATORPUBKEYLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.VALIDATORPUBKEYLENGTH(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORSETHASHLENGTH is a free data retrieval call binding the contract method 0xd785e50e.
//
// Solidity: function VALIDATOR_SET_HASH_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VALIDATORSETHASHLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "VALIDATOR_SET_HASH_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORSETHASHLENGTH is a free data retrieval call binding the contract method 0xd785e50e.
//
// Solidity: function VALIDATOR_SET_HASH_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) VALIDATORSETHASHLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.VALIDATORSETHASHLENGTH(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORSETHASHLENGTH is a free data retrieval call binding the contract method 0xd785e50e.
//
// Solidity: function VALIDATOR_SET_HASH_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VALIDATORSETHASHLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.VALIDATORSETHASHLENGTH(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORVOTINGPOWERLENGTH is a free data retrieval call binding the contract method 0x66952964.
//
// Solidity: function VALIDATOR_VOTING_POWER_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VALIDATORVOTINGPOWERLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "VALIDATOR_VOTING_POWER_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORVOTINGPOWERLENGTH is a free data retrieval call binding the contract method 0x66952964.
//
// Solidity: function VALIDATOR_VOTING_POWER_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) VALIDATORVOTINGPOWERLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.VALIDATORVOTINGPOWERLENGTH(&_Greenfieldlightclient.CallOpts)
}

// VALIDATORVOTINGPOWERLENGTH is a free data retrieval call binding the contract method 0x66952964.
//
// Solidity: function VALIDATOR_VOTING_POWER_LENGTH() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VALIDATORVOTINGPOWERLENGTH() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.VALIDATORVOTINGPOWERLENGTH(&_Greenfieldlightclient.CallOpts)
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes _blsPubKeys)
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
// Solidity: function blsPubKeys() view returns(bytes _blsPubKeys)
func (_Greenfieldlightclient *GreenfieldlightclientSession) BlsPubKeys() ([]byte, error) {
	return _Greenfieldlightclient.Contract.BlsPubKeys(&_Greenfieldlightclient.CallOpts)
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes _blsPubKeys)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) BlsPubKeys() ([]byte, error) {
	return _Greenfieldlightclient.Contract.BlsPubKeys(&_Greenfieldlightclient.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(bytes32)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) ChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(bytes32)
func (_Greenfieldlightclient *GreenfieldlightclientSession) ChainID() ([32]byte, error) {
	return _Greenfieldlightclient.Contract.ChainID(&_Greenfieldlightclient.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(bytes32)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) ChainID() ([32]byte, error) {
	return _Greenfieldlightclient.Contract.ChainID(&_Greenfieldlightclient.CallOpts)
}

// ConsensusStateBytes is a free data retrieval call binding the contract method 0x82e2111c.
//
// Solidity: function consensusStateBytes() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) ConsensusStateBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "consensusStateBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConsensusStateBytes is a free data retrieval call binding the contract method 0x82e2111c.
//
// Solidity: function consensusStateBytes() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientSession) ConsensusStateBytes() ([]byte, error) {
	return _Greenfieldlightclient.Contract.ConsensusStateBytes(&_Greenfieldlightclient.CallOpts)
}

// ConsensusStateBytes is a free data retrieval call binding the contract method 0x82e2111c.
//
// Solidity: function consensusStateBytes() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) ConsensusStateBytes() ([]byte, error) {
	return _Greenfieldlightclient.Contract.ConsensusStateBytes(&_Greenfieldlightclient.CallOpts)
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

// Height is a free data retrieval call binding the contract method 0x0ef26743.
//
// Solidity: function height() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) Height(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "height")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Height is a free data retrieval call binding the contract method 0x0ef26743.
//
// Solidity: function height() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientSession) Height() (uint64, error) {
	return _Greenfieldlightclient.Contract.Height(&_Greenfieldlightclient.CallOpts)
}

// Height is a free data retrieval call binding the contract method 0x0ef26743.
//
// Solidity: function height() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) Height() (uint64, error) {
	return _Greenfieldlightclient.Contract.Height(&_Greenfieldlightclient.CallOpts)
}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) InitialHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "initialHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientSession) InitialHeight() (uint64, error) {
	return _Greenfieldlightclient.Contract.InitialHeight(&_Greenfieldlightclient.CallOpts)
}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) InitialHeight() (uint64, error) {
	return _Greenfieldlightclient.Contract.InitialHeight(&_Greenfieldlightclient.CallOpts)
}

// NextValidatorSetHash is a free data retrieval call binding the contract method 0x752d3b89.
//
// Solidity: function nextValidatorSetHash() view returns(bytes32)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) NextValidatorSetHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "nextValidatorSetHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextValidatorSetHash is a free data retrieval call binding the contract method 0x752d3b89.
//
// Solidity: function nextValidatorSetHash() view returns(bytes32)
func (_Greenfieldlightclient *GreenfieldlightclientSession) NextValidatorSetHash() ([32]byte, error) {
	return _Greenfieldlightclient.Contract.NextValidatorSetHash(&_Greenfieldlightclient.CallOpts)
}

// NextValidatorSetHash is a free data retrieval call binding the contract method 0x752d3b89.
//
// Solidity: function nextValidatorSetHash() view returns(bytes32)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) NextValidatorSetHash() ([32]byte, error) {
	return _Greenfieldlightclient.Contract.NextValidatorSetHash(&_Greenfieldlightclient.CallOpts)
}

// Submitters is a free data retrieval call binding the contract method 0xda8d08f0.
//
// Solidity: function submitters(uint64 ) view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) Submitters(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "submitters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Submitters is a free data retrieval call binding the contract method 0xda8d08f0.
//
// Solidity: function submitters(uint64 ) view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) Submitters(arg0 uint64) (common.Address, error) {
	return _Greenfieldlightclient.Contract.Submitters(&_Greenfieldlightclient.CallOpts, arg0)
}

// Submitters is a free data retrieval call binding the contract method 0xda8d08f0.
//
// Solidity: function submitters(uint64 ) view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) Submitters(arg0 uint64) (common.Address, error) {
	return _Greenfieldlightclient.Contract.Submitters(&_Greenfieldlightclient.CallOpts, arg0)
}

// ValidatorSet is a free data retrieval call binding the contract method 0xe64808f3.
//
// Solidity: function validatorSet(uint256 ) view returns(bytes32 pubKey, int64 votingPower, address relayerAddress, bytes relayerBlsKey)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) ValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PubKey         [32]byte
	VotingPower    int64
	RelayerAddress common.Address
	RelayerBlsKey  []byte
}, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "validatorSet", arg0)

	outstruct := new(struct {
		PubKey         [32]byte
		VotingPower    int64
		RelayerAddress common.Address
		RelayerBlsKey  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PubKey = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VotingPower = *abi.ConvertType(out[1], new(int64)).(*int64)
	outstruct.RelayerAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.RelayerBlsKey = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ValidatorSet is a free data retrieval call binding the contract method 0xe64808f3.
//
// Solidity: function validatorSet(uint256 ) view returns(bytes32 pubKey, int64 votingPower, address relayerAddress, bytes relayerBlsKey)
func (_Greenfieldlightclient *GreenfieldlightclientSession) ValidatorSet(arg0 *big.Int) (struct {
	PubKey         [32]byte
	VotingPower    int64
	RelayerAddress common.Address
	RelayerBlsKey  []byte
}, error) {
	return _Greenfieldlightclient.Contract.ValidatorSet(&_Greenfieldlightclient.CallOpts, arg0)
}

// ValidatorSet is a free data retrieval call binding the contract method 0xe64808f3.
//
// Solidity: function validatorSet(uint256 ) view returns(bytes32 pubKey, int64 votingPower, address relayerAddress, bytes relayerBlsKey)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) ValidatorSet(arg0 *big.Int) (struct {
	PubKey         [32]byte
	VotingPower    int64
	RelayerAddress common.Address
	RelayerBlsKey  []byte
}, error) {
	return _Greenfieldlightclient.Contract.ValidatorSet(&_Greenfieldlightclient.CallOpts, arg0)
}

// VerifyPackage is a free data retrieval call binding the contract method 0xd527fba1.
//
// Solidity: function verifyPackage(bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VerifyPackage(opts *bind.CallOpts, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) (bool, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "verifyPackage", _payload, _blsSignature, _validatorSetBitMap)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPackage is a free data retrieval call binding the contract method 0xd527fba1.
//
// Solidity: function verifyPackage(bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientSession) VerifyPackage(_payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) (bool, error) {
	return _Greenfieldlightclient.Contract.VerifyPackage(&_Greenfieldlightclient.CallOpts, _payload, _blsSignature, _validatorSetBitMap)
}

// VerifyPackage is a free data retrieval call binding the contract method 0xd527fba1.
//
// Solidity: function verifyPackage(bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VerifyPackage(_payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) (bool, error) {
	return _Greenfieldlightclient.Contract.VerifyPackage(&_Greenfieldlightclient.CallOpts, _payload, _blsSignature, _validatorSetBitMap)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _initConsensusStateBytes) returns()
func (_Greenfieldlightclient *GreenfieldlightclientTransactor) Initialize(opts *bind.TransactOpts, _initConsensusStateBytes []byte) (*types.Transaction, error) {
	return _Greenfieldlightclient.contract.Transact(opts, "initialize", _initConsensusStateBytes)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _initConsensusStateBytes) returns()
func (_Greenfieldlightclient *GreenfieldlightclientSession) Initialize(_initConsensusStateBytes []byte) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.Initialize(&_Greenfieldlightclient.TransactOpts, _initConsensusStateBytes)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _initConsensusStateBytes) returns()
func (_Greenfieldlightclient *GreenfieldlightclientTransactorSession) Initialize(_initConsensusStateBytes []byte) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.Initialize(&_Greenfieldlightclient.TransactOpts, _initConsensusStateBytes)
}

// SyncLightBlock is a paid mutator transaction binding the contract method 0xff08c81e.
//
// Solidity: function syncLightBlock(bytes _lightBlock, uint64 _height) returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientTransactor) SyncLightBlock(opts *bind.TransactOpts, _lightBlock []byte, _height uint64) (*types.Transaction, error) {
	return _Greenfieldlightclient.contract.Transact(opts, "syncLightBlock", _lightBlock, _height)
}

// SyncLightBlock is a paid mutator transaction binding the contract method 0xff08c81e.
//
// Solidity: function syncLightBlock(bytes _lightBlock, uint64 _height) returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientSession) SyncLightBlock(_lightBlock []byte, _height uint64) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.SyncLightBlock(&_Greenfieldlightclient.TransactOpts, _lightBlock, _height)
}

// SyncLightBlock is a paid mutator transaction binding the contract method 0xff08c81e.
//
// Solidity: function syncLightBlock(bytes _lightBlock, uint64 _height) returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientTransactorSession) SyncLightBlock(_lightBlock []byte, _height uint64) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.SyncLightBlock(&_Greenfieldlightclient.TransactOpts, _lightBlock, _height)
}

// GreenfieldlightclientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Greenfieldlightclient contract.
type GreenfieldlightclientInitializedIterator struct {
	Event *GreenfieldlightclientInitialized // Event containing the contract specifics and raw log

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
func (it *GreenfieldlightclientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GreenfieldlightclientInitialized)
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
		it.Event = new(GreenfieldlightclientInitialized)
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
func (it *GreenfieldlightclientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GreenfieldlightclientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GreenfieldlightclientInitialized represents a Initialized event raised by the Greenfieldlightclient contract.
type GreenfieldlightclientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) FilterInitialized(opts *bind.FilterOpts) (*GreenfieldlightclientInitializedIterator, error) {

	logs, sub, err := _Greenfieldlightclient.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientInitializedIterator{contract: _Greenfieldlightclient.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GreenfieldlightclientInitialized) (event.Subscription, error) {

	logs, sub, err := _Greenfieldlightclient.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GreenfieldlightclientInitialized)
				if err := _Greenfieldlightclient.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) ParseInitialized(log types.Log) (*GreenfieldlightclientInitialized, error) {
	event := new(GreenfieldlightclientInitialized)
	if err := _Greenfieldlightclient.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GreenfieldlightclientInitConsensusStateIterator is returned from FilterInitConsensusState and is used to iterate over the raw logs and unpacked data for InitConsensusState events raised by the Greenfieldlightclient contract.
type GreenfieldlightclientInitConsensusStateIterator struct {
	Event *GreenfieldlightclientInitConsensusState // Event containing the contract specifics and raw log

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
func (it *GreenfieldlightclientInitConsensusStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GreenfieldlightclientInitConsensusState)
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
		it.Event = new(GreenfieldlightclientInitConsensusState)
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
func (it *GreenfieldlightclientInitConsensusStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GreenfieldlightclientInitConsensusStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GreenfieldlightclientInitConsensusState represents a InitConsensusState event raised by the Greenfieldlightclient contract.
type GreenfieldlightclientInitConsensusState struct {
	Height uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInitConsensusState is a free log retrieval operation binding the contract event 0x37d76474dadef267cd6fb443be36fea42881c1f0b5163a6cc00fa67480e3aac4.
//
// Solidity: event initConsensusState(uint64 height)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) FilterInitConsensusState(opts *bind.FilterOpts) (*GreenfieldlightclientInitConsensusStateIterator, error) {

	logs, sub, err := _Greenfieldlightclient.contract.FilterLogs(opts, "initConsensusState")
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientInitConsensusStateIterator{contract: _Greenfieldlightclient.contract, event: "initConsensusState", logs: logs, sub: sub}, nil
}

// WatchInitConsensusState is a free log subscription operation binding the contract event 0x37d76474dadef267cd6fb443be36fea42881c1f0b5163a6cc00fa67480e3aac4.
//
// Solidity: event initConsensusState(uint64 height)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) WatchInitConsensusState(opts *bind.WatchOpts, sink chan<- *GreenfieldlightclientInitConsensusState) (event.Subscription, error) {

	logs, sub, err := _Greenfieldlightclient.contract.WatchLogs(opts, "initConsensusState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GreenfieldlightclientInitConsensusState)
				if err := _Greenfieldlightclient.contract.UnpackLog(event, "initConsensusState", log); err != nil {
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

// ParseInitConsensusState is a log parse operation binding the contract event 0x37d76474dadef267cd6fb443be36fea42881c1f0b5163a6cc00fa67480e3aac4.
//
// Solidity: event initConsensusState(uint64 height)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) ParseInitConsensusState(log types.Log) (*GreenfieldlightclientInitConsensusState, error) {
	event := new(GreenfieldlightclientInitConsensusState)
	if err := _Greenfieldlightclient.contract.UnpackLog(event, "initConsensusState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GreenfieldlightclientUpdateConsensusStateIterator is returned from FilterUpdateConsensusState and is used to iterate over the raw logs and unpacked data for UpdateConsensusState events raised by the Greenfieldlightclient contract.
type GreenfieldlightclientUpdateConsensusStateIterator struct {
	Event *GreenfieldlightclientUpdateConsensusState // Event containing the contract specifics and raw log

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
func (it *GreenfieldlightclientUpdateConsensusStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GreenfieldlightclientUpdateConsensusState)
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
		it.Event = new(GreenfieldlightclientUpdateConsensusState)
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
func (it *GreenfieldlightclientUpdateConsensusStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GreenfieldlightclientUpdateConsensusStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GreenfieldlightclientUpdateConsensusState represents a UpdateConsensusState event raised by the Greenfieldlightclient contract.
type GreenfieldlightclientUpdateConsensusState struct {
	Height              uint64
	ValidatorSetChanged bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateConsensusState is a free log retrieval operation binding the contract event 0x691118119129ba12b1904302d5fe91091cdcbd2be353a35219a4bcadabe4dcfd.
//
// Solidity: event updateConsensusState(uint64 height, bool validatorSetChanged)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) FilterUpdateConsensusState(opts *bind.FilterOpts) (*GreenfieldlightclientUpdateConsensusStateIterator, error) {

	logs, sub, err := _Greenfieldlightclient.contract.FilterLogs(opts, "updateConsensusState")
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientUpdateConsensusStateIterator{contract: _Greenfieldlightclient.contract, event: "updateConsensusState", logs: logs, sub: sub}, nil
}

// WatchUpdateConsensusState is a free log subscription operation binding the contract event 0x691118119129ba12b1904302d5fe91091cdcbd2be353a35219a4bcadabe4dcfd.
//
// Solidity: event updateConsensusState(uint64 height, bool validatorSetChanged)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) WatchUpdateConsensusState(opts *bind.WatchOpts, sink chan<- *GreenfieldlightclientUpdateConsensusState) (event.Subscription, error) {

	logs, sub, err := _Greenfieldlightclient.contract.WatchLogs(opts, "updateConsensusState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GreenfieldlightclientUpdateConsensusState)
				if err := _Greenfieldlightclient.contract.UnpackLog(event, "updateConsensusState", log); err != nil {
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

// ParseUpdateConsensusState is a log parse operation binding the contract event 0x691118119129ba12b1904302d5fe91091cdcbd2be353a35219a4bcadabe4dcfd.
//
// Solidity: event updateConsensusState(uint64 height, bool validatorSetChanged)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) ParseUpdateConsensusState(log types.Log) (*GreenfieldlightclientUpdateConsensusState, error) {
	event := new(GreenfieldlightclientUpdateConsensusState)
	if err := _Greenfieldlightclient.contract.UnpackLog(event, "updateConsensusState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
