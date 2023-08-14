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

// ILightClientInturnRelayer is an auto generated low-level Go binding around an user-defined struct.
type ILightClientInturnRelayer struct {
	Addr   common.Address
	BlsKey []byte
	Start  *big.Int
	End    *big.Int
}

// GreenfieldlightclientMetaData contains all meta data concerning the Greenfieldlightclient contract.
var GreenfieldlightclientMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"InitConsensusState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"ParamChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"validatorSetChanged\",\"type\":\"bool\"}],\"name\":\"UpdatedConsensusState\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLS_SIGNATURE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUCKET_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUCKET_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHAIN_ID_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_STATE_BASE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_STATE_BYTES_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ConfigSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_VALIDATE_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEIGHT_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSAGE_HASH_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OBJECT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OBJECT_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PACKAGE_VERIFY_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROXY_ADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ADDRESS_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_BLS_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_BYTES_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_PUB_KEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_SET_HASH_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_VOTING_POWER_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsPubKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_blsPubKeys\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"consensusStateBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInturnRelayer\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"internalType\":\"structILightClient.InturnRelayer\",\"name\":\"relayer\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInturnRelayerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInturnRelayerBlsPubKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gnfdHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inTurnRelayerRelayInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inTurnRelayerValidityPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_initConsensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidatorSetHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"submitters\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_lightBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"syncLightBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorSet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"},{\"internalType\":\"int64\",\"name\":\"votingPower\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"relayerBlsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetBitMap\",\"type\":\"uint256\"}],\"name\":\"verifyPackage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"eventTime\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetBitMap\",\"type\":\"uint256\"}],\"name\":\"verifyRelayerAndPackage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) BUCKETCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "BUCKET_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) BUCKETCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.BUCKETCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) BUCKETCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.BUCKETCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) BUCKETHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "BUCKET_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) BUCKETHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.BUCKETHUB(&_Greenfieldlightclient.CallOpts)
}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) BUCKETHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.BUCKETHUB(&_Greenfieldlightclient.CallOpts)
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

// ConfigSlots is a free data retrieval call binding the contract method 0xb76e4aca.
//
// Solidity: function ConfigSlots(uint256 ) view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) ConfigSlots(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "ConfigSlots", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfigSlots is a free data retrieval call binding the contract method 0xb76e4aca.
//
// Solidity: function ConfigSlots(uint256 ) view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) ConfigSlots(arg0 *big.Int) (*big.Int, error) {
	return _Greenfieldlightclient.Contract.ConfigSlots(&_Greenfieldlightclient.CallOpts, arg0)
}

// ConfigSlots is a free data retrieval call binding the contract method 0xb76e4aca.
//
// Solidity: function ConfigSlots(uint256 ) view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) ConfigSlots(arg0 *big.Int) (*big.Int, error) {
	return _Greenfieldlightclient.Contract.ConfigSlots(&_Greenfieldlightclient.CallOpts, arg0)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "GOV_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GOVCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.GOVCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
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

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GROUPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "GROUP_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GROUPCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.GROUPCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GROUPCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.GROUPCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GROUPHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "GROUP_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GROUPHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.GROUPHUB(&_Greenfieldlightclient.CallOpts)
}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GROUPHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.GROUPHUB(&_Greenfieldlightclient.CallOpts)
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

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) OBJECTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "OBJECT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) OBJECTCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.OBJECTCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) OBJECTCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.OBJECTCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) OBJECTHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "OBJECT_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) OBJECTHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.OBJECTHUB(&_Greenfieldlightclient.CallOpts)
}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) OBJECTHUB() (common.Address, error) {
	return _Greenfieldlightclient.Contract.OBJECTHUB(&_Greenfieldlightclient.CallOpts)
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

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "TRANSFER_IN_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.TRANSFERINCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.TRANSFERINCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "TRANSFER_OUT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_Greenfieldlightclient *GreenfieldlightclientSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Greenfieldlightclient.Contract.TRANSFEROUTCHANNELID(&_Greenfieldlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
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

// GetInturnRelayer is a free data retrieval call binding the contract method 0x3baa2219.
//
// Solidity: function getInturnRelayer() view returns((address,bytes,uint256,uint256) relayer)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GetInturnRelayer(opts *bind.CallOpts) (ILightClientInturnRelayer, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "getInturnRelayer")

	if err != nil {
		return *new(ILightClientInturnRelayer), err
	}

	out0 := *abi.ConvertType(out[0], new(ILightClientInturnRelayer)).(*ILightClientInturnRelayer)

	return out0, err

}

// GetInturnRelayer is a free data retrieval call binding the contract method 0x3baa2219.
//
// Solidity: function getInturnRelayer() view returns((address,bytes,uint256,uint256) relayer)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GetInturnRelayer() (ILightClientInturnRelayer, error) {
	return _Greenfieldlightclient.Contract.GetInturnRelayer(&_Greenfieldlightclient.CallOpts)
}

// GetInturnRelayer is a free data retrieval call binding the contract method 0x3baa2219.
//
// Solidity: function getInturnRelayer() view returns((address,bytes,uint256,uint256) relayer)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GetInturnRelayer() (ILightClientInturnRelayer, error) {
	return _Greenfieldlightclient.Contract.GetInturnRelayer(&_Greenfieldlightclient.CallOpts)
}

// GetInturnRelayerAddress is a free data retrieval call binding the contract method 0xd5b0f519.
//
// Solidity: function getInturnRelayerAddress() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GetInturnRelayerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "getInturnRelayerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInturnRelayerAddress is a free data retrieval call binding the contract method 0xd5b0f519.
//
// Solidity: function getInturnRelayerAddress() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GetInturnRelayerAddress() (common.Address, error) {
	return _Greenfieldlightclient.Contract.GetInturnRelayerAddress(&_Greenfieldlightclient.CallOpts)
}

// GetInturnRelayerAddress is a free data retrieval call binding the contract method 0xd5b0f519.
//
// Solidity: function getInturnRelayerAddress() view returns(address)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GetInturnRelayerAddress() (common.Address, error) {
	return _Greenfieldlightclient.Contract.GetInturnRelayerAddress(&_Greenfieldlightclient.CallOpts)
}

// GetInturnRelayerBlsPubKey is a free data retrieval call binding the contract method 0x56bbb6e7.
//
// Solidity: function getInturnRelayerBlsPubKey() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GetInturnRelayerBlsPubKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "getInturnRelayerBlsPubKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInturnRelayerBlsPubKey is a free data retrieval call binding the contract method 0x56bbb6e7.
//
// Solidity: function getInturnRelayerBlsPubKey() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GetInturnRelayerBlsPubKey() ([]byte, error) {
	return _Greenfieldlightclient.Contract.GetInturnRelayerBlsPubKey(&_Greenfieldlightclient.CallOpts)
}

// GetInturnRelayerBlsPubKey is a free data retrieval call binding the contract method 0x56bbb6e7.
//
// Solidity: function getInturnRelayerBlsPubKey() view returns(bytes)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GetInturnRelayerBlsPubKey() ([]byte, error) {
	return _Greenfieldlightclient.Contract.GetInturnRelayerBlsPubKey(&_Greenfieldlightclient.CallOpts)
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

// GnfdHeight is a free data retrieval call binding the contract method 0x64118b9e.
//
// Solidity: function gnfdHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) GnfdHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "gnfdHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GnfdHeight is a free data retrieval call binding the contract method 0x64118b9e.
//
// Solidity: function gnfdHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientSession) GnfdHeight() (uint64, error) {
	return _Greenfieldlightclient.Contract.GnfdHeight(&_Greenfieldlightclient.CallOpts)
}

// GnfdHeight is a free data retrieval call binding the contract method 0x64118b9e.
//
// Solidity: function gnfdHeight() view returns(uint64)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) GnfdHeight() (uint64, error) {
	return _Greenfieldlightclient.Contract.GnfdHeight(&_Greenfieldlightclient.CallOpts)
}

// InTurnRelayerRelayInterval is a free data retrieval call binding the contract method 0x3f8e2a84.
//
// Solidity: function inTurnRelayerRelayInterval() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) InTurnRelayerRelayInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "inTurnRelayerRelayInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InTurnRelayerRelayInterval is a free data retrieval call binding the contract method 0x3f8e2a84.
//
// Solidity: function inTurnRelayerRelayInterval() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) InTurnRelayerRelayInterval() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.InTurnRelayerRelayInterval(&_Greenfieldlightclient.CallOpts)
}

// InTurnRelayerRelayInterval is a free data retrieval call binding the contract method 0x3f8e2a84.
//
// Solidity: function inTurnRelayerRelayInterval() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) InTurnRelayerRelayInterval() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.InTurnRelayerRelayInterval(&_Greenfieldlightclient.CallOpts)
}

// InTurnRelayerValidityPeriod is a free data retrieval call binding the contract method 0xf4efa5a7.
//
// Solidity: function inTurnRelayerValidityPeriod() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) InTurnRelayerValidityPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "inTurnRelayerValidityPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InTurnRelayerValidityPeriod is a free data retrieval call binding the contract method 0xf4efa5a7.
//
// Solidity: function inTurnRelayerValidityPeriod() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientSession) InTurnRelayerValidityPeriod() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.InTurnRelayerValidityPeriod(&_Greenfieldlightclient.CallOpts)
}

// InTurnRelayerValidityPeriod is a free data retrieval call binding the contract method 0xf4efa5a7.
//
// Solidity: function inTurnRelayerValidityPeriod() view returns(uint256)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) InTurnRelayerValidityPeriod() (*big.Int, error) {
	return _Greenfieldlightclient.Contract.InTurnRelayerValidityPeriod(&_Greenfieldlightclient.CallOpts)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address ) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) IsRelayer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "isRelayer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address ) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientSession) IsRelayer(arg0 common.Address) (bool, error) {
	return _Greenfieldlightclient.Contract.IsRelayer(&_Greenfieldlightclient.CallOpts, arg0)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address ) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) IsRelayer(arg0 common.Address) (bool, error) {
	return _Greenfieldlightclient.Contract.IsRelayer(&_Greenfieldlightclient.CallOpts, arg0)
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

// VerifyRelayerAndPackage is a free data retrieval call binding the contract method 0xf2acebd4.
//
// Solidity: function verifyRelayerAndPackage(uint64 eventTime, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VerifyRelayerAndPackage(opts *bind.CallOpts, eventTime uint64, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) (bool, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "verifyRelayerAndPackage", eventTime, _payload, _blsSignature, _validatorSetBitMap)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRelayerAndPackage is a free data retrieval call binding the contract method 0xf2acebd4.
//
// Solidity: function verifyRelayerAndPackage(uint64 eventTime, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientSession) VerifyRelayerAndPackage(eventTime uint64, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) (bool, error) {
	return _Greenfieldlightclient.Contract.VerifyRelayerAndPackage(&_Greenfieldlightclient.CallOpts, eventTime, _payload, _blsSignature, _validatorSetBitMap)
}

// VerifyRelayerAndPackage is a free data retrieval call binding the contract method 0xf2acebd4.
//
// Solidity: function verifyRelayerAndPackage(uint64 eventTime, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns(bool)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VerifyRelayerAndPackage(eventTime uint64, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) (bool, error) {
	return _Greenfieldlightclient.Contract.VerifyRelayerAndPackage(&_Greenfieldlightclient.CallOpts, eventTime, _payload, _blsSignature, _validatorSetBitMap)
}

// VersionInfo is a free data retrieval call binding the contract method 0xa21d1354.
//
// Solidity: function versionInfo() pure returns(uint256 version, string name, string description)
func (_Greenfieldlightclient *GreenfieldlightclientCaller) VersionInfo(opts *bind.CallOpts) (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	var out []interface{}
	err := _Greenfieldlightclient.contract.Call(opts, &out, "versionInfo")

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
func (_Greenfieldlightclient *GreenfieldlightclientSession) VersionInfo() (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	return _Greenfieldlightclient.Contract.VersionInfo(&_Greenfieldlightclient.CallOpts)
}

// VersionInfo is a free data retrieval call binding the contract method 0xa21d1354.
//
// Solidity: function versionInfo() pure returns(uint256 version, string name, string description)
func (_Greenfieldlightclient *GreenfieldlightclientCallerSession) VersionInfo() (struct {
	Version     *big.Int
	Name        string
	Description string
}, error) {
	return _Greenfieldlightclient.Contract.VersionInfo(&_Greenfieldlightclient.CallOpts)
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

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Greenfieldlightclient *GreenfieldlightclientTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Greenfieldlightclient.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Greenfieldlightclient *GreenfieldlightclientSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.UpdateParam(&_Greenfieldlightclient.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Greenfieldlightclient *GreenfieldlightclientTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Greenfieldlightclient.Contract.UpdateParam(&_Greenfieldlightclient.TransactOpts, key, value)
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

// FilterInitConsensusState is a free log retrieval operation binding the contract event 0x4742d6829e4cda17f0acee2327b4c31e69243f15f2e5493e784b7e584c89fbf0.
//
// Solidity: event InitConsensusState(uint64 height)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) FilterInitConsensusState(opts *bind.FilterOpts) (*GreenfieldlightclientInitConsensusStateIterator, error) {

	logs, sub, err := _Greenfieldlightclient.contract.FilterLogs(opts, "InitConsensusState")
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientInitConsensusStateIterator{contract: _Greenfieldlightclient.contract, event: "InitConsensusState", logs: logs, sub: sub}, nil
}

// WatchInitConsensusState is a free log subscription operation binding the contract event 0x4742d6829e4cda17f0acee2327b4c31e69243f15f2e5493e784b7e584c89fbf0.
//
// Solidity: event InitConsensusState(uint64 height)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) WatchInitConsensusState(opts *bind.WatchOpts, sink chan<- *GreenfieldlightclientInitConsensusState) (event.Subscription, error) {

	logs, sub, err := _Greenfieldlightclient.contract.WatchLogs(opts, "InitConsensusState")
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
				if err := _Greenfieldlightclient.contract.UnpackLog(event, "InitConsensusState", log); err != nil {
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

// ParseInitConsensusState is a log parse operation binding the contract event 0x4742d6829e4cda17f0acee2327b4c31e69243f15f2e5493e784b7e584c89fbf0.
//
// Solidity: event InitConsensusState(uint64 height)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) ParseInitConsensusState(log types.Log) (*GreenfieldlightclientInitConsensusState, error) {
	event := new(GreenfieldlightclientInitConsensusState)
	if err := _Greenfieldlightclient.contract.UnpackLog(event, "InitConsensusState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// GreenfieldlightclientParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Greenfieldlightclient contract.
type GreenfieldlightclientParamChangeIterator struct {
	Event *GreenfieldlightclientParamChange // Event containing the contract specifics and raw log

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
func (it *GreenfieldlightclientParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GreenfieldlightclientParamChange)
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
		it.Event = new(GreenfieldlightclientParamChange)
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
func (it *GreenfieldlightclientParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GreenfieldlightclientParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GreenfieldlightclientParamChange represents a ParamChange event raised by the Greenfieldlightclient contract.
type GreenfieldlightclientParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) FilterParamChange(opts *bind.FilterOpts) (*GreenfieldlightclientParamChangeIterator, error) {

	logs, sub, err := _Greenfieldlightclient.contract.FilterLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientParamChangeIterator{contract: _Greenfieldlightclient.contract, event: "ParamChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *GreenfieldlightclientParamChange) (event.Subscription, error) {

	logs, sub, err := _Greenfieldlightclient.contract.WatchLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GreenfieldlightclientParamChange)
				if err := _Greenfieldlightclient.contract.UnpackLog(event, "ParamChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) ParseParamChange(log types.Log) (*GreenfieldlightclientParamChange, error) {
	event := new(GreenfieldlightclientParamChange)
	if err := _Greenfieldlightclient.contract.UnpackLog(event, "ParamChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GreenfieldlightclientUpdatedConsensusStateIterator is returned from FilterUpdatedConsensusState and is used to iterate over the raw logs and unpacked data for UpdatedConsensusState events raised by the Greenfieldlightclient contract.
type GreenfieldlightclientUpdatedConsensusStateIterator struct {
	Event *GreenfieldlightclientUpdatedConsensusState // Event containing the contract specifics and raw log

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
func (it *GreenfieldlightclientUpdatedConsensusStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GreenfieldlightclientUpdatedConsensusState)
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
		it.Event = new(GreenfieldlightclientUpdatedConsensusState)
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
func (it *GreenfieldlightclientUpdatedConsensusStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GreenfieldlightclientUpdatedConsensusStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GreenfieldlightclientUpdatedConsensusState represents a UpdatedConsensusState event raised by the Greenfieldlightclient contract.
type GreenfieldlightclientUpdatedConsensusState struct {
	Height              uint64
	ValidatorSetChanged bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdatedConsensusState is a free log retrieval operation binding the contract event 0xa44b844260c407b2179345e959e21dad09b1630c624779683c7cbbdd8b6ecd7b.
//
// Solidity: event UpdatedConsensusState(uint64 height, bool validatorSetChanged)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) FilterUpdatedConsensusState(opts *bind.FilterOpts) (*GreenfieldlightclientUpdatedConsensusStateIterator, error) {

	logs, sub, err := _Greenfieldlightclient.contract.FilterLogs(opts, "UpdatedConsensusState")
	if err != nil {
		return nil, err
	}
	return &GreenfieldlightclientUpdatedConsensusStateIterator{contract: _Greenfieldlightclient.contract, event: "UpdatedConsensusState", logs: logs, sub: sub}, nil
}

// WatchUpdatedConsensusState is a free log subscription operation binding the contract event 0xa44b844260c407b2179345e959e21dad09b1630c624779683c7cbbdd8b6ecd7b.
//
// Solidity: event UpdatedConsensusState(uint64 height, bool validatorSetChanged)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) WatchUpdatedConsensusState(opts *bind.WatchOpts, sink chan<- *GreenfieldlightclientUpdatedConsensusState) (event.Subscription, error) {

	logs, sub, err := _Greenfieldlightclient.contract.WatchLogs(opts, "UpdatedConsensusState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GreenfieldlightclientUpdatedConsensusState)
				if err := _Greenfieldlightclient.contract.UnpackLog(event, "UpdatedConsensusState", log); err != nil {
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

// ParseUpdatedConsensusState is a log parse operation binding the contract event 0xa44b844260c407b2179345e959e21dad09b1630c624779683c7cbbdd8b6ecd7b.
//
// Solidity: event UpdatedConsensusState(uint64 height, bool validatorSetChanged)
func (_Greenfieldlightclient *GreenfieldlightclientFilterer) ParseUpdatedConsensusState(log types.Log) (*GreenfieldlightclientUpdatedConsensusState, error) {
	event := new(GreenfieldlightclientUpdatedConsensusState)
	if err := _Greenfieldlightclient.contract.UnpackLog(event, "UpdatedConsensusState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
