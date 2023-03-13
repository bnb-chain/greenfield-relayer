// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crosschain

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

// CrosschainMetaData contains all meta data concerning the Crosschain contract.
var CrosschainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"AddChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"oracleSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"CrossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnable\",\"type\":\"bool\"}],\"name\":\"EnableOrDisableChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"ParamChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalTypeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"quorum\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"expiredAt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contentHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"ReceivedPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Reopened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Suspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"UnexpectedFailureAssertionInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"UnexpectedRevertInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"UnsupportedPackage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"APP_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUCKET_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUCKET_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCEL_TRANSFER_PROPOSAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMERGENCY_PROPOSAL_EXPIRE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMPTY_CONTENT_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAIL_ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GROUP_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IN_TURN_RELAYER_VALIDITY_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OBJECT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OBJECT_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROXY_ADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REOPEN_PROPOSAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUSPEND_PROPOSAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYN_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNEL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ackRelayFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchSizeForOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attacker\",\"type\":\"address\"}],\"name\":\"cancelTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelHandlerMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelReceiveSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelSendSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"emergencyProposals\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"quorum\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"expiredAt\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"contentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ackRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"encodePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gnfdChainId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorsBitSet\",\"type\":\"uint256\"}],\"name\":\"handlePackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_gnfdChainId\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSuspended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleSequence\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousTxHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"quorumMap\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"registeredContractChannelMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reopen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ackRelayFee\",\"type\":\"uint256\"}],\"name\":\"sendSynPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CrosschainABI is the input ABI used to generate the binding from.
// Deprecated: Use CrosschainMetaData.ABI instead.
var CrosschainABI = CrosschainMetaData.ABI

// Crosschain is an auto generated Go binding around an Ethereum contract.
type Crosschain struct {
	CrosschainCaller     // Read-only binding to the contract
	CrosschainTransactor // Write-only binding to the contract
	CrosschainFilterer   // Log filterer for contract events
}

// CrosschainCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrosschainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrosschainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrosschainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrosschainSession struct {
	Contract     *Crosschain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrosschainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrosschainCallerSession struct {
	Contract *CrosschainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrosschainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrosschainTransactorSession struct {
	Contract     *CrosschainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrosschainRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrosschainRaw struct {
	Contract *Crosschain // Generic contract binding to access the raw methods on
}

// CrosschainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrosschainCallerRaw struct {
	Contract *CrosschainCaller // Generic read-only contract binding to access the raw methods on
}

// CrosschainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrosschainTransactorRaw struct {
	Contract *CrosschainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrosschain creates a new instance of Crosschain, bound to a specific deployed contract.
func NewCrosschain(address common.Address, backend bind.ContractBackend) (*Crosschain, error) {
	contract, err := bindCrosschain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crosschain{CrosschainCaller: CrosschainCaller{contract: contract}, CrosschainTransactor: CrosschainTransactor{contract: contract}, CrosschainFilterer: CrosschainFilterer{contract: contract}}, nil
}

// NewCrosschainCaller creates a new read-only instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainCaller(address common.Address, caller bind.ContractCaller) (*CrosschainCaller, error) {
	contract, err := bindCrosschain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrosschainCaller{contract: contract}, nil
}

// NewCrosschainTransactor creates a new write-only instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainTransactor(address common.Address, transactor bind.ContractTransactor) (*CrosschainTransactor, error) {
	contract, err := bindCrosschain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrosschainTransactor{contract: contract}, nil
}

// NewCrosschainFilterer creates a new log filterer instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainFilterer(address common.Address, filterer bind.ContractFilterer) (*CrosschainFilterer, error) {
	contract, err := bindCrosschain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrosschainFilterer{contract: contract}, nil
}

// bindCrosschain binds a generic wrapper to an already deployed contract.
func bindCrosschain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrosschainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crosschain *CrosschainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crosschain.Contract.CrosschainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crosschain *CrosschainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.Contract.CrosschainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crosschain *CrosschainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crosschain.Contract.CrosschainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crosschain *CrosschainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crosschain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crosschain *CrosschainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crosschain *CrosschainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crosschain.Contract.contract.Transact(opts, method, params...)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainCaller) ACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "ACK_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainSession) ACKPACKAGE() (uint8, error) {
	return _Crosschain.Contract.ACKPACKAGE(&_Crosschain.CallOpts)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) ACKPACKAGE() (uint8, error) {
	return _Crosschain.Contract.ACKPACKAGE(&_Crosschain.CallOpts)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x1124de3a.
//
// Solidity: function APP_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCaller) APPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "APP_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// APPCHANNELID is a free data retrieval call binding the contract method 0x1124de3a.
//
// Solidity: function APP_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainSession) APPCHANNELID() (uint8, error) {
	return _Crosschain.Contract.APPCHANNELID(&_Crosschain.CallOpts)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x1124de3a.
//
// Solidity: function APP_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) APPCHANNELID() (uint8, error) {
	return _Crosschain.Contract.APPCHANNELID(&_Crosschain.CallOpts)
}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCaller) BUCKETCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "BUCKET_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainSession) BUCKETCHANNELID() (uint8, error) {
	return _Crosschain.Contract.BUCKETCHANNELID(&_Crosschain.CallOpts)
}

// BUCKETCHANNELID is a free data retrieval call binding the contract method 0x73f1e3c3.
//
// Solidity: function BUCKET_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) BUCKETCHANNELID() (uint8, error) {
	return _Crosschain.Contract.BUCKETCHANNELID(&_Crosschain.CallOpts)
}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Crosschain *CrosschainCaller) BUCKETHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "BUCKET_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Crosschain *CrosschainSession) BUCKETHUB() (common.Address, error) {
	return _Crosschain.Contract.BUCKETHUB(&_Crosschain.CallOpts)
}

// BUCKETHUB is a free data retrieval call binding the contract method 0x7afffdd2.
//
// Solidity: function BUCKET_HUB() view returns(address)
func (_Crosschain *CrosschainCallerSession) BUCKETHUB() (common.Address, error) {
	return _Crosschain.Contract.BUCKETHUB(&_Crosschain.CallOpts)
}

// CANCELTRANSFERPROPOSAL is a free data retrieval call binding the contract method 0x5692ddd3.
//
// Solidity: function CANCEL_TRANSFER_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainCaller) CANCELTRANSFERPROPOSAL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "CANCEL_TRANSFER_PROPOSAL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELTRANSFERPROPOSAL is a free data retrieval call binding the contract method 0x5692ddd3.
//
// Solidity: function CANCEL_TRANSFER_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainSession) CANCELTRANSFERPROPOSAL() ([32]byte, error) {
	return _Crosschain.Contract.CANCELTRANSFERPROPOSAL(&_Crosschain.CallOpts)
}

// CANCELTRANSFERPROPOSAL is a free data retrieval call binding the contract method 0x5692ddd3.
//
// Solidity: function CANCEL_TRANSFER_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainCallerSession) CANCELTRANSFERPROPOSAL() ([32]byte, error) {
	return _Crosschain.Contract.CANCELTRANSFERPROPOSAL(&_Crosschain.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Crosschain *CrosschainCaller) CROSSCHAIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "CROSS_CHAIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Crosschain *CrosschainSession) CROSSCHAIN() (common.Address, error) {
	return _Crosschain.Contract.CROSSCHAIN(&_Crosschain.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Crosschain *CrosschainCallerSession) CROSSCHAIN() (common.Address, error) {
	return _Crosschain.Contract.CROSSCHAIN(&_Crosschain.CallOpts)
}

// EMERGENCYPROPOSALEXPIREPERIOD is a free data retrieval call binding the contract method 0xdc404331.
//
// Solidity: function EMERGENCY_PROPOSAL_EXPIRE_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainCaller) EMERGENCYPROPOSALEXPIREPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "EMERGENCY_PROPOSAL_EXPIRE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMERGENCYPROPOSALEXPIREPERIOD is a free data retrieval call binding the contract method 0xdc404331.
//
// Solidity: function EMERGENCY_PROPOSAL_EXPIRE_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainSession) EMERGENCYPROPOSALEXPIREPERIOD() (*big.Int, error) {
	return _Crosschain.Contract.EMERGENCYPROPOSALEXPIREPERIOD(&_Crosschain.CallOpts)
}

// EMERGENCYPROPOSALEXPIREPERIOD is a free data retrieval call binding the contract method 0xdc404331.
//
// Solidity: function EMERGENCY_PROPOSAL_EXPIRE_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) EMERGENCYPROPOSALEXPIREPERIOD() (*big.Int, error) {
	return _Crosschain.Contract.EMERGENCYPROPOSALEXPIREPERIOD(&_Crosschain.CallOpts)
}

// EMPTYCONTENTHASH is a free data retrieval call binding the contract method 0xc780e9de.
//
// Solidity: function EMPTY_CONTENT_HASH() view returns(bytes32)
func (_Crosschain *CrosschainCaller) EMPTYCONTENTHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "EMPTY_CONTENT_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EMPTYCONTENTHASH is a free data retrieval call binding the contract method 0xc780e9de.
//
// Solidity: function EMPTY_CONTENT_HASH() view returns(bytes32)
func (_Crosschain *CrosschainSession) EMPTYCONTENTHASH() ([32]byte, error) {
	return _Crosschain.Contract.EMPTYCONTENTHASH(&_Crosschain.CallOpts)
}

// EMPTYCONTENTHASH is a free data retrieval call binding the contract method 0xc780e9de.
//
// Solidity: function EMPTY_CONTENT_HASH() view returns(bytes32)
func (_Crosschain *CrosschainCallerSession) EMPTYCONTENTHASH() ([32]byte, error) {
	return _Crosschain.Contract.EMPTYCONTENTHASH(&_Crosschain.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainCaller) FAILACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "FAIL_ACK_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainSession) FAILACKPACKAGE() (uint8, error) {
	return _Crosschain.Contract.FAILACKPACKAGE(&_Crosschain.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) FAILACKPACKAGE() (uint8, error) {
	return _Crosschain.Contract.FAILACKPACKAGE(&_Crosschain.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "GOV_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainSession) GOVCHANNELID() (uint8, error) {
	return _Crosschain.Contract.GOVCHANNELID(&_Crosschain.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x81d91480.
//
// Solidity: function GOV_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) GOVCHANNELID() (uint8, error) {
	return _Crosschain.Contract.GOVCHANNELID(&_Crosschain.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Crosschain *CrosschainCaller) GOVHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "GOV_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Crosschain *CrosschainSession) GOVHUB() (common.Address, error) {
	return _Crosschain.Contract.GOVHUB(&_Crosschain.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Crosschain *CrosschainCallerSession) GOVHUB() (common.Address, error) {
	return _Crosschain.Contract.GOVHUB(&_Crosschain.CallOpts)
}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCaller) GROUPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "GROUP_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainSession) GROUPCHANNELID() (uint8, error) {
	return _Crosschain.Contract.GROUPCHANNELID(&_Crosschain.CallOpts)
}

// GROUPCHANNELID is a free data retrieval call binding the contract method 0xe02e86b0.
//
// Solidity: function GROUP_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) GROUPCHANNELID() (uint8, error) {
	return _Crosschain.Contract.GROUPCHANNELID(&_Crosschain.CallOpts)
}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Crosschain *CrosschainCaller) GROUPHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "GROUP_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Crosschain *CrosschainSession) GROUPHUB() (common.Address, error) {
	return _Crosschain.Contract.GROUPHUB(&_Crosschain.CallOpts)
}

// GROUPHUB is a free data retrieval call binding the contract method 0x46934fc8.
//
// Solidity: function GROUP_HUB() view returns(address)
func (_Crosschain *CrosschainCallerSession) GROUPHUB() (common.Address, error) {
	return _Crosschain.Contract.GROUPHUB(&_Crosschain.CallOpts)
}

// INTURNRELAYERVALIDITYPERIOD is a free data retrieval call binding the contract method 0xe036ead6.
//
// Solidity: function IN_TURN_RELAYER_VALIDITY_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainCaller) INTURNRELAYERVALIDITYPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "IN_TURN_RELAYER_VALIDITY_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INTURNRELAYERVALIDITYPERIOD is a free data retrieval call binding the contract method 0xe036ead6.
//
// Solidity: function IN_TURN_RELAYER_VALIDITY_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainSession) INTURNRELAYERVALIDITYPERIOD() (*big.Int, error) {
	return _Crosschain.Contract.INTURNRELAYERVALIDITYPERIOD(&_Crosschain.CallOpts)
}

// INTURNRELAYERVALIDITYPERIOD is a free data retrieval call binding the contract method 0xe036ead6.
//
// Solidity: function IN_TURN_RELAYER_VALIDITY_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) INTURNRELAYERVALIDITYPERIOD() (*big.Int, error) {
	return _Crosschain.Contract.INTURNRELAYERVALIDITYPERIOD(&_Crosschain.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Crosschain *CrosschainCaller) LIGHTCLIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "LIGHT_CLIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Crosschain *CrosschainSession) LIGHTCLIENT() (common.Address, error) {
	return _Crosschain.Contract.LIGHTCLIENT(&_Crosschain.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Crosschain *CrosschainCallerSession) LIGHTCLIENT() (common.Address, error) {
	return _Crosschain.Contract.LIGHTCLIENT(&_Crosschain.CallOpts)
}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCaller) OBJECTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "OBJECT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainSession) OBJECTCHANNELID() (uint8, error) {
	return _Crosschain.Contract.OBJECTCHANNELID(&_Crosschain.CallOpts)
}

// OBJECTCHANNELID is a free data retrieval call binding the contract method 0xeac78b33.
//
// Solidity: function OBJECT_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) OBJECTCHANNELID() (uint8, error) {
	return _Crosschain.Contract.OBJECTCHANNELID(&_Crosschain.CallOpts)
}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Crosschain *CrosschainCaller) OBJECTHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "OBJECT_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Crosschain *CrosschainSession) OBJECTHUB() (common.Address, error) {
	return _Crosschain.Contract.OBJECTHUB(&_Crosschain.CallOpts)
}

// OBJECTHUB is a free data retrieval call binding the contract method 0x7d2e3084.
//
// Solidity: function OBJECT_HUB() view returns(address)
func (_Crosschain *CrosschainCallerSession) OBJECTHUB() (common.Address, error) {
	return _Crosschain.Contract.OBJECTHUB(&_Crosschain.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Crosschain *CrosschainCaller) PROXYADMIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "PROXY_ADMIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Crosschain *CrosschainSession) PROXYADMIN() (common.Address, error) {
	return _Crosschain.Contract.PROXYADMIN(&_Crosschain.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Crosschain *CrosschainCallerSession) PROXYADMIN() (common.Address, error) {
	return _Crosschain.Contract.PROXYADMIN(&_Crosschain.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Crosschain *CrosschainCaller) RELAYERHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "RELAYER_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Crosschain *CrosschainSession) RELAYERHUB() (common.Address, error) {
	return _Crosschain.Contract.RELAYERHUB(&_Crosschain.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Crosschain *CrosschainCallerSession) RELAYERHUB() (common.Address, error) {
	return _Crosschain.Contract.RELAYERHUB(&_Crosschain.CallOpts)
}

// REOPENPROPOSAL is a free data retrieval call binding the contract method 0x6de380bd.
//
// Solidity: function REOPEN_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainCaller) REOPENPROPOSAL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "REOPEN_PROPOSAL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REOPENPROPOSAL is a free data retrieval call binding the contract method 0x6de380bd.
//
// Solidity: function REOPEN_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainSession) REOPENPROPOSAL() ([32]byte, error) {
	return _Crosschain.Contract.REOPENPROPOSAL(&_Crosschain.CallOpts)
}

// REOPENPROPOSAL is a free data retrieval call binding the contract method 0x6de380bd.
//
// Solidity: function REOPEN_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainCallerSession) REOPENPROPOSAL() ([32]byte, error) {
	return _Crosschain.Contract.REOPENPROPOSAL(&_Crosschain.CallOpts)
}

// SUSPENDPROPOSAL is a free data retrieval call binding the contract method 0x63e1394e.
//
// Solidity: function SUSPEND_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainCaller) SUSPENDPROPOSAL(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "SUSPEND_PROPOSAL")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUSPENDPROPOSAL is a free data retrieval call binding the contract method 0x63e1394e.
//
// Solidity: function SUSPEND_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainSession) SUSPENDPROPOSAL() ([32]byte, error) {
	return _Crosschain.Contract.SUSPENDPROPOSAL(&_Crosschain.CallOpts)
}

// SUSPENDPROPOSAL is a free data retrieval call binding the contract method 0x63e1394e.
//
// Solidity: function SUSPEND_PROPOSAL() view returns(bytes32)
func (_Crosschain *CrosschainCallerSession) SUSPENDPROPOSAL() ([32]byte, error) {
	return _Crosschain.Contract.SUSPENDPROPOSAL(&_Crosschain.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainCaller) SYNPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "SYN_PACKAGE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainSession) SYNPACKAGE() (uint8, error) {
	return _Crosschain.Contract.SYNPACKAGE(&_Crosschain.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) SYNPACKAGE() (uint8, error) {
	return _Crosschain.Contract.SYNPACKAGE(&_Crosschain.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Crosschain *CrosschainCaller) TOKENHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "TOKEN_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Crosschain *CrosschainSession) TOKENHUB() (common.Address, error) {
	return _Crosschain.Contract.TOKENHUB(&_Crosschain.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Crosschain *CrosschainCallerSession) TOKENHUB() (common.Address, error) {
	return _Crosschain.Contract.TOKENHUB(&_Crosschain.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "TRANSFER_IN_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFERINCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0xcc12eabc.
//
// Solidity: function TRANSFER_IN_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFERINCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "TRANSFER_OUT_CHANNEL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFEROUTCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0x618d569c.
//
// Solidity: function TRANSFER_OUT_CHANNEL_ID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFEROUTCHANNELID(&_Crosschain.CallOpts)
}

// AckRelayFee is a free data retrieval call binding the contract method 0x6ab31754.
//
// Solidity: function ackRelayFee() view returns(uint256)
func (_Crosschain *CrosschainCaller) AckRelayFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "ackRelayFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AckRelayFee is a free data retrieval call binding the contract method 0x6ab31754.
//
// Solidity: function ackRelayFee() view returns(uint256)
func (_Crosschain *CrosschainSession) AckRelayFee() (*big.Int, error) {
	return _Crosschain.Contract.AckRelayFee(&_Crosschain.CallOpts)
}

// AckRelayFee is a free data retrieval call binding the contract method 0x6ab31754.
//
// Solidity: function ackRelayFee() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) AckRelayFee() (*big.Int, error) {
	return _Crosschain.Contract.AckRelayFee(&_Crosschain.CallOpts)
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() view returns(uint256)
func (_Crosschain *CrosschainCaller) BatchSizeForOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "batchSizeForOracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() view returns(uint256)
func (_Crosschain *CrosschainSession) BatchSizeForOracle() (*big.Int, error) {
	return _Crosschain.Contract.BatchSizeForOracle(&_Crosschain.CallOpts)
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) BatchSizeForOracle() (*big.Int, error) {
	return _Crosschain.Contract.BatchSizeForOracle(&_Crosschain.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Crosschain *CrosschainCaller) ChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Crosschain *CrosschainSession) ChainId() (uint16, error) {
	return _Crosschain.Contract.ChainId(&_Crosschain.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint16)
func (_Crosschain *CrosschainCallerSession) ChainId() (uint16, error) {
	return _Crosschain.Contract.ChainId(&_Crosschain.CallOpts)
}

// ChannelHandlerMap is a free data retrieval call binding the contract method 0xf77d0a34.
//
// Solidity: function channelHandlerMap(uint8 ) view returns(address)
func (_Crosschain *CrosschainCaller) ChannelHandlerMap(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "channelHandlerMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChannelHandlerMap is a free data retrieval call binding the contract method 0xf77d0a34.
//
// Solidity: function channelHandlerMap(uint8 ) view returns(address)
func (_Crosschain *CrosschainSession) ChannelHandlerMap(arg0 uint8) (common.Address, error) {
	return _Crosschain.Contract.ChannelHandlerMap(&_Crosschain.CallOpts, arg0)
}

// ChannelHandlerMap is a free data retrieval call binding the contract method 0xf77d0a34.
//
// Solidity: function channelHandlerMap(uint8 ) view returns(address)
func (_Crosschain *CrosschainCallerSession) ChannelHandlerMap(arg0 uint8) (common.Address, error) {
	return _Crosschain.Contract.ChannelHandlerMap(&_Crosschain.CallOpts, arg0)
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainCaller) ChannelReceiveSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "channelReceiveSequenceMap", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainSession) ChannelReceiveSequenceMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelReceiveSequenceMap(&_Crosschain.CallOpts, arg0)
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainCallerSession) ChannelReceiveSequenceMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelReceiveSequenceMap(&_Crosschain.CallOpts, arg0)
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainCaller) ChannelSendSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "channelSendSequenceMap", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainSession) ChannelSendSequenceMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelSendSequenceMap(&_Crosschain.CallOpts, arg0)
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainCallerSession) ChannelSendSequenceMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelSendSequenceMap(&_Crosschain.CallOpts, arg0)
}

// EmergencyProposals is a free data retrieval call binding the contract method 0x6bacff2c.
//
// Solidity: function emergencyProposals(bytes32 ) view returns(uint16 quorum, uint128 expiredAt, bytes32 contentHash)
func (_Crosschain *CrosschainCaller) EmergencyProposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Quorum      uint16
	ExpiredAt   *big.Int
	ContentHash [32]byte
}, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "emergencyProposals", arg0)

	outstruct := new(struct {
		Quorum      uint16
		ExpiredAt   *big.Int
		ContentHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quorum = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.ExpiredAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ContentHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// EmergencyProposals is a free data retrieval call binding the contract method 0x6bacff2c.
//
// Solidity: function emergencyProposals(bytes32 ) view returns(uint16 quorum, uint128 expiredAt, bytes32 contentHash)
func (_Crosschain *CrosschainSession) EmergencyProposals(arg0 [32]byte) (struct {
	Quorum      uint16
	ExpiredAt   *big.Int
	ContentHash [32]byte
}, error) {
	return _Crosschain.Contract.EmergencyProposals(&_Crosschain.CallOpts, arg0)
}

// EmergencyProposals is a free data retrieval call binding the contract method 0x6bacff2c.
//
// Solidity: function emergencyProposals(bytes32 ) view returns(uint16 quorum, uint128 expiredAt, bytes32 contentHash)
func (_Crosschain *CrosschainCallerSession) EmergencyProposals(arg0 [32]byte) (struct {
	Quorum      uint16
	ExpiredAt   *big.Int
	ContentHash [32]byte
}, error) {
	return _Crosschain.Contract.EmergencyProposals(&_Crosschain.CallOpts, arg0)
}

// EncodePayload is a free data retrieval call binding the contract method 0x9c06ea47.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, uint256 ackRelayFee, bytes msgBytes) view returns(bytes)
func (_Crosschain *CrosschainCaller) EncodePayload(opts *bind.CallOpts, packageType uint8, relayFee *big.Int, ackRelayFee *big.Int, msgBytes []byte) ([]byte, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "encodePayload", packageType, relayFee, ackRelayFee, msgBytes)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePayload is a free data retrieval call binding the contract method 0x9c06ea47.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, uint256 ackRelayFee, bytes msgBytes) view returns(bytes)
func (_Crosschain *CrosschainSession) EncodePayload(packageType uint8, relayFee *big.Int, ackRelayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _Crosschain.Contract.EncodePayload(&_Crosschain.CallOpts, packageType, relayFee, ackRelayFee, msgBytes)
}

// EncodePayload is a free data retrieval call binding the contract method 0x9c06ea47.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, uint256 ackRelayFee, bytes msgBytes) view returns(bytes)
func (_Crosschain *CrosschainCallerSession) EncodePayload(packageType uint8, relayFee *big.Int, ackRelayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _Crosschain.Contract.EncodePayload(&_Crosschain.CallOpts, packageType, relayFee, ackRelayFee, msgBytes)
}

// GnfdChainId is a free data retrieval call binding the contract method 0x96b1ec6e.
//
// Solidity: function gnfdChainId() view returns(uint16)
func (_Crosschain *CrosschainCaller) GnfdChainId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "gnfdChainId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GnfdChainId is a free data retrieval call binding the contract method 0x96b1ec6e.
//
// Solidity: function gnfdChainId() view returns(uint16)
func (_Crosschain *CrosschainSession) GnfdChainId() (uint16, error) {
	return _Crosschain.Contract.GnfdChainId(&_Crosschain.CallOpts)
}

// GnfdChainId is a free data retrieval call binding the contract method 0x96b1ec6e.
//
// Solidity: function gnfdChainId() view returns(uint16)
func (_Crosschain *CrosschainCallerSession) GnfdChainId() (uint16, error) {
	return _Crosschain.Contract.GnfdChainId(&_Crosschain.CallOpts)
}

// IsSuspended is a free data retrieval call binding the contract method 0x1d130935.
//
// Solidity: function isSuspended() view returns(bool)
func (_Crosschain *CrosschainCaller) IsSuspended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "isSuspended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSuspended is a free data retrieval call binding the contract method 0x1d130935.
//
// Solidity: function isSuspended() view returns(bool)
func (_Crosschain *CrosschainSession) IsSuspended() (bool, error) {
	return _Crosschain.Contract.IsSuspended(&_Crosschain.CallOpts)
}

// IsSuspended is a free data retrieval call binding the contract method 0x1d130935.
//
// Solidity: function isSuspended() view returns(bool)
func (_Crosschain *CrosschainCallerSession) IsSuspended() (bool, error) {
	return _Crosschain.Contract.IsSuspended(&_Crosschain.CallOpts)
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() view returns(int64)
func (_Crosschain *CrosschainCaller) OracleSequence(opts *bind.CallOpts) (int64, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "oracleSequence")

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() view returns(int64)
func (_Crosschain *CrosschainSession) OracleSequence() (int64, error) {
	return _Crosschain.Contract.OracleSequence(&_Crosschain.CallOpts)
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() view returns(int64)
func (_Crosschain *CrosschainCallerSession) OracleSequence() (int64, error) {
	return _Crosschain.Contract.OracleSequence(&_Crosschain.CallOpts)
}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() view returns(uint256)
func (_Crosschain *CrosschainCaller) PreviousTxHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "previousTxHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() view returns(uint256)
func (_Crosschain *CrosschainSession) PreviousTxHeight() (*big.Int, error) {
	return _Crosschain.Contract.PreviousTxHeight(&_Crosschain.CallOpts)
}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) PreviousTxHeight() (*big.Int, error) {
	return _Crosschain.Contract.PreviousTxHeight(&_Crosschain.CallOpts)
}

// QuorumMap is a free data retrieval call binding the contract method 0x299b533d.
//
// Solidity: function quorumMap(bytes32 ) view returns(uint16)
func (_Crosschain *CrosschainCaller) QuorumMap(opts *bind.CallOpts, arg0 [32]byte) (uint16, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "quorumMap", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// QuorumMap is a free data retrieval call binding the contract method 0x299b533d.
//
// Solidity: function quorumMap(bytes32 ) view returns(uint16)
func (_Crosschain *CrosschainSession) QuorumMap(arg0 [32]byte) (uint16, error) {
	return _Crosschain.Contract.QuorumMap(&_Crosschain.CallOpts, arg0)
}

// QuorumMap is a free data retrieval call binding the contract method 0x299b533d.
//
// Solidity: function quorumMap(bytes32 ) view returns(uint16)
func (_Crosschain *CrosschainCallerSession) QuorumMap(arg0 [32]byte) (uint16, error) {
	return _Crosschain.Contract.QuorumMap(&_Crosschain.CallOpts, arg0)
}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) view returns(bool)
func (_Crosschain *CrosschainCaller) RegisteredContractChannelMap(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (bool, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "registeredContractChannelMap", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) view returns(bool)
func (_Crosschain *CrosschainSession) RegisteredContractChannelMap(arg0 common.Address, arg1 uint8) (bool, error) {
	return _Crosschain.Contract.RegisteredContractChannelMap(&_Crosschain.CallOpts, arg0, arg1)
}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) view returns(bool)
func (_Crosschain *CrosschainCallerSession) RegisteredContractChannelMap(arg0 common.Address, arg1 uint8) (bool, error) {
	return _Crosschain.Contract.RegisteredContractChannelMap(&_Crosschain.CallOpts, arg0, arg1)
}

// RelayFee is a free data retrieval call binding the contract method 0x71d30863.
//
// Solidity: function relayFee() view returns(uint256)
func (_Crosschain *CrosschainCaller) RelayFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "relayFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayFee is a free data retrieval call binding the contract method 0x71d30863.
//
// Solidity: function relayFee() view returns(uint256)
func (_Crosschain *CrosschainSession) RelayFee() (*big.Int, error) {
	return _Crosschain.Contract.RelayFee(&_Crosschain.CallOpts)
}

// RelayFee is a free data retrieval call binding the contract method 0x71d30863.
//
// Solidity: function relayFee() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) RelayFee() (*big.Int, error) {
	return _Crosschain.Contract.RelayFee(&_Crosschain.CallOpts)
}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() view returns(uint256)
func (_Crosschain *CrosschainCaller) TxCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "txCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() view returns(uint256)
func (_Crosschain *CrosschainSession) TxCounter() (*big.Int, error) {
	return _Crosschain.Contract.TxCounter(&_Crosschain.CallOpts)
}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) TxCounter() (*big.Int, error) {
	return _Crosschain.Contract.TxCounter(&_Crosschain.CallOpts)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xa7c1e629.
//
// Solidity: function cancelTransfer(address attacker) returns()
func (_Crosschain *CrosschainTransactor) CancelTransfer(opts *bind.TransactOpts, attacker common.Address) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "cancelTransfer", attacker)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xa7c1e629.
//
// Solidity: function cancelTransfer(address attacker) returns()
func (_Crosschain *CrosschainSession) CancelTransfer(attacker common.Address) (*types.Transaction, error) {
	return _Crosschain.Contract.CancelTransfer(&_Crosschain.TransactOpts, attacker)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0xa7c1e629.
//
// Solidity: function cancelTransfer(address attacker) returns()
func (_Crosschain *CrosschainTransactorSession) CancelTransfer(attacker common.Address) (*types.Transaction, error) {
	return _Crosschain.Contract.CancelTransfer(&_Crosschain.TransactOpts, attacker)
}

// HandlePackage is a paid mutator transaction binding the contract method 0xc9978d24.
//
// Solidity: function handlePackage(bytes _payload, bytes _blsSignature, uint256 _validatorsBitSet) returns()
func (_Crosschain *CrosschainTransactor) HandlePackage(opts *bind.TransactOpts, _payload []byte, _blsSignature []byte, _validatorsBitSet *big.Int) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "handlePackage", _payload, _blsSignature, _validatorsBitSet)
}

// HandlePackage is a paid mutator transaction binding the contract method 0xc9978d24.
//
// Solidity: function handlePackage(bytes _payload, bytes _blsSignature, uint256 _validatorsBitSet) returns()
func (_Crosschain *CrosschainSession) HandlePackage(_payload []byte, _blsSignature []byte, _validatorsBitSet *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.HandlePackage(&_Crosschain.TransactOpts, _payload, _blsSignature, _validatorsBitSet)
}

// HandlePackage is a paid mutator transaction binding the contract method 0xc9978d24.
//
// Solidity: function handlePackage(bytes _payload, bytes _blsSignature, uint256 _validatorsBitSet) returns()
func (_Crosschain *CrosschainTransactorSession) HandlePackage(_payload []byte, _blsSignature []byte, _validatorsBitSet *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.HandlePackage(&_Crosschain.TransactOpts, _payload, _blsSignature, _validatorsBitSet)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _gnfdChainId) returns()
func (_Crosschain *CrosschainTransactor) Initialize(opts *bind.TransactOpts, _gnfdChainId uint16) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "initialize", _gnfdChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _gnfdChainId) returns()
func (_Crosschain *CrosschainSession) Initialize(_gnfdChainId uint16) (*types.Transaction, error) {
	return _Crosschain.Contract.Initialize(&_Crosschain.TransactOpts, _gnfdChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _gnfdChainId) returns()
func (_Crosschain *CrosschainTransactorSession) Initialize(_gnfdChainId uint16) (*types.Transaction, error) {
	return _Crosschain.Contract.Initialize(&_Crosschain.TransactOpts, _gnfdChainId)
}

// Reopen is a paid mutator transaction binding the contract method 0xccc108d7.
//
// Solidity: function reopen() returns()
func (_Crosschain *CrosschainTransactor) Reopen(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "reopen")
}

// Reopen is a paid mutator transaction binding the contract method 0xccc108d7.
//
// Solidity: function reopen() returns()
func (_Crosschain *CrosschainSession) Reopen() (*types.Transaction, error) {
	return _Crosschain.Contract.Reopen(&_Crosschain.TransactOpts)
}

// Reopen is a paid mutator transaction binding the contract method 0xccc108d7.
//
// Solidity: function reopen() returns()
func (_Crosschain *CrosschainTransactorSession) Reopen() (*types.Transaction, error) {
	return _Crosschain.Contract.Reopen(&_Crosschain.TransactOpts)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0x8f884dda.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee, uint256 ackRelayFee) returns()
func (_Crosschain *CrosschainTransactor) SendSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte, relayFee *big.Int, ackRelayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "sendSynPackage", channelId, msgBytes, relayFee, ackRelayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0x8f884dda.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee, uint256 ackRelayFee) returns()
func (_Crosschain *CrosschainSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int, ackRelayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.SendSynPackage(&_Crosschain.TransactOpts, channelId, msgBytes, relayFee, ackRelayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0x8f884dda.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee, uint256 ackRelayFee) returns()
func (_Crosschain *CrosschainTransactorSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int, ackRelayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.SendSynPackage(&_Crosschain.TransactOpts, channelId, msgBytes, relayFee, ackRelayFee)
}

// Suspend is a paid mutator transaction binding the contract method 0xe6400bbe.
//
// Solidity: function suspend() returns()
func (_Crosschain *CrosschainTransactor) Suspend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "suspend")
}

// Suspend is a paid mutator transaction binding the contract method 0xe6400bbe.
//
// Solidity: function suspend() returns()
func (_Crosschain *CrosschainSession) Suspend() (*types.Transaction, error) {
	return _Crosschain.Contract.Suspend(&_Crosschain.TransactOpts)
}

// Suspend is a paid mutator transaction binding the contract method 0xe6400bbe.
//
// Solidity: function suspend() returns()
func (_Crosschain *CrosschainTransactorSession) Suspend() (*types.Transaction, error) {
	return _Crosschain.Contract.Suspend(&_Crosschain.TransactOpts)
}

// CrosschainAddChannelIterator is returned from FilterAddChannel and is used to iterate over the raw logs and unpacked data for AddChannel events raised by the Crosschain contract.
type CrosschainAddChannelIterator struct {
	Event *CrosschainAddChannel // Event containing the contract specifics and raw log

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
func (it *CrosschainAddChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainAddChannel)
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
		it.Event = new(CrosschainAddChannel)
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
func (it *CrosschainAddChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainAddChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainAddChannel represents a AddChannel event raised by the Crosschain contract.
type CrosschainAddChannel struct {
	ChannelId    uint8
	ContractAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddChannel is a free log retrieval operation binding the contract event 0x50509989df6f2738ba0458b9a89f29591b60973aa0556c6ce0db9be78f5d5688.
//
// Solidity: event AddChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Crosschain *CrosschainFilterer) FilterAddChannel(opts *bind.FilterOpts, channelId []uint8, contractAddr []common.Address) (*CrosschainAddChannelIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "AddChannel", channelIdRule, contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainAddChannelIterator{contract: _Crosschain.contract, event: "AddChannel", logs: logs, sub: sub}, nil
}

// WatchAddChannel is a free log subscription operation binding the contract event 0x50509989df6f2738ba0458b9a89f29591b60973aa0556c6ce0db9be78f5d5688.
//
// Solidity: event AddChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Crosschain *CrosschainFilterer) WatchAddChannel(opts *bind.WatchOpts, sink chan<- *CrosschainAddChannel, channelId []uint8, contractAddr []common.Address) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "AddChannel", channelIdRule, contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainAddChannel)
				if err := _Crosschain.contract.UnpackLog(event, "AddChannel", log); err != nil {
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

// ParseAddChannel is a log parse operation binding the contract event 0x50509989df6f2738ba0458b9a89f29591b60973aa0556c6ce0db9be78f5d5688.
//
// Solidity: event AddChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Crosschain *CrosschainFilterer) ParseAddChannel(log types.Log) (*CrosschainAddChannel, error) {
	event := new(CrosschainAddChannel)
	if err := _Crosschain.contract.UnpackLog(event, "AddChannel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainCrossChainPackageIterator is returned from FilterCrossChainPackage and is used to iterate over the raw logs and unpacked data for CrossChainPackage events raised by the Crosschain contract.
type CrosschainCrossChainPackageIterator struct {
	Event *CrosschainCrossChainPackage // Event containing the contract specifics and raw log

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
func (it *CrosschainCrossChainPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainCrossChainPackage)
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
		it.Event = new(CrosschainCrossChainPackage)
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
func (it *CrosschainCrossChainPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainCrossChainPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainCrossChainPackage represents a CrossChainPackage event raised by the Crosschain contract.
type CrosschainCrossChainPackage struct {
	SrcChainId      uint32
	DstChainId      uint32
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCrossChainPackage is a free log retrieval operation binding the contract event 0x64998dc5a229e7324e622192f111c691edccc3534bbea4b2bd90fbaec936845a.
//
// Solidity: event CrossChainPackage(uint32 srcChainId, uint32 dstChainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) FilterCrossChainPackage(opts *bind.FilterOpts, oracleSequence []uint64, packageSequence []uint64, channelId []uint8) (*CrosschainCrossChainPackageIterator, error) {

	var oracleSequenceRule []interface{}
	for _, oracleSequenceItem := range oracleSequence {
		oracleSequenceRule = append(oracleSequenceRule, oracleSequenceItem)
	}
	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "CrossChainPackage", oracleSequenceRule, packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainCrossChainPackageIterator{contract: _Crosschain.contract, event: "CrossChainPackage", logs: logs, sub: sub}, nil
}

// WatchCrossChainPackage is a free log subscription operation binding the contract event 0x64998dc5a229e7324e622192f111c691edccc3534bbea4b2bd90fbaec936845a.
//
// Solidity: event CrossChainPackage(uint32 srcChainId, uint32 dstChainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) WatchCrossChainPackage(opts *bind.WatchOpts, sink chan<- *CrosschainCrossChainPackage, oracleSequence []uint64, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var oracleSequenceRule []interface{}
	for _, oracleSequenceItem := range oracleSequence {
		oracleSequenceRule = append(oracleSequenceRule, oracleSequenceItem)
	}
	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "CrossChainPackage", oracleSequenceRule, packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainCrossChainPackage)
				if err := _Crosschain.contract.UnpackLog(event, "CrossChainPackage", log); err != nil {
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

// ParseCrossChainPackage is a log parse operation binding the contract event 0x64998dc5a229e7324e622192f111c691edccc3534bbea4b2bd90fbaec936845a.
//
// Solidity: event CrossChainPackage(uint32 srcChainId, uint32 dstChainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) ParseCrossChainPackage(log types.Log) (*CrosschainCrossChainPackage, error) {
	event := new(CrosschainCrossChainPackage)
	if err := _Crosschain.contract.UnpackLog(event, "CrossChainPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainEnableOrDisableChannelIterator is returned from FilterEnableOrDisableChannel and is used to iterate over the raw logs and unpacked data for EnableOrDisableChannel events raised by the Crosschain contract.
type CrosschainEnableOrDisableChannelIterator struct {
	Event *CrosschainEnableOrDisableChannel // Event containing the contract specifics and raw log

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
func (it *CrosschainEnableOrDisableChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainEnableOrDisableChannel)
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
		it.Event = new(CrosschainEnableOrDisableChannel)
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
func (it *CrosschainEnableOrDisableChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainEnableOrDisableChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainEnableOrDisableChannel represents a EnableOrDisableChannel event raised by the Crosschain contract.
type CrosschainEnableOrDisableChannel struct {
	ChannelId uint8
	IsEnable  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEnableOrDisableChannel is a free log retrieval operation binding the contract event 0x7aa83b30cf2563aa45de6f2f23245bf55fa4f2a8425f37a0cbf027c8281bdc82.
//
// Solidity: event EnableOrDisableChannel(uint8 indexed channelId, bool isEnable)
func (_Crosschain *CrosschainFilterer) FilterEnableOrDisableChannel(opts *bind.FilterOpts, channelId []uint8) (*CrosschainEnableOrDisableChannelIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "EnableOrDisableChannel", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainEnableOrDisableChannelIterator{contract: _Crosschain.contract, event: "EnableOrDisableChannel", logs: logs, sub: sub}, nil
}

// WatchEnableOrDisableChannel is a free log subscription operation binding the contract event 0x7aa83b30cf2563aa45de6f2f23245bf55fa4f2a8425f37a0cbf027c8281bdc82.
//
// Solidity: event EnableOrDisableChannel(uint8 indexed channelId, bool isEnable)
func (_Crosschain *CrosschainFilterer) WatchEnableOrDisableChannel(opts *bind.WatchOpts, sink chan<- *CrosschainEnableOrDisableChannel, channelId []uint8) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "EnableOrDisableChannel", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainEnableOrDisableChannel)
				if err := _Crosschain.contract.UnpackLog(event, "EnableOrDisableChannel", log); err != nil {
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

// ParseEnableOrDisableChannel is a log parse operation binding the contract event 0x7aa83b30cf2563aa45de6f2f23245bf55fa4f2a8425f37a0cbf027c8281bdc82.
//
// Solidity: event EnableOrDisableChannel(uint8 indexed channelId, bool isEnable)
func (_Crosschain *CrosschainFilterer) ParseEnableOrDisableChannel(log types.Log) (*CrosschainEnableOrDisableChannel, error) {
	event := new(CrosschainEnableOrDisableChannel)
	if err := _Crosschain.contract.UnpackLog(event, "EnableOrDisableChannel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Crosschain contract.
type CrosschainInitializedIterator struct {
	Event *CrosschainInitialized // Event containing the contract specifics and raw log

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
func (it *CrosschainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainInitialized)
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
		it.Event = new(CrosschainInitialized)
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
func (it *CrosschainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainInitialized represents a Initialized event raised by the Crosschain contract.
type CrosschainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Crosschain *CrosschainFilterer) FilterInitialized(opts *bind.FilterOpts) (*CrosschainInitializedIterator, error) {

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CrosschainInitializedIterator{contract: _Crosschain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Crosschain *CrosschainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CrosschainInitialized) (event.Subscription, error) {

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainInitialized)
				if err := _Crosschain.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Crosschain *CrosschainFilterer) ParseInitialized(log types.Log) (*CrosschainInitialized, error) {
	event := new(CrosschainInitialized)
	if err := _Crosschain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Crosschain contract.
type CrosschainParamChangeIterator struct {
	Event *CrosschainParamChange // Event containing the contract specifics and raw log

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
func (it *CrosschainParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainParamChange)
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
		it.Event = new(CrosschainParamChange)
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
func (it *CrosschainParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainParamChange represents a ParamChange event raised by the Crosschain contract.
type CrosschainParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Crosschain *CrosschainFilterer) FilterParamChange(opts *bind.FilterOpts) (*CrosschainParamChangeIterator, error) {

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return &CrosschainParamChangeIterator{contract: _Crosschain.contract, event: "ParamChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Crosschain *CrosschainFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *CrosschainParamChange) (event.Subscription, error) {

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainParamChange)
				if err := _Crosschain.contract.UnpackLog(event, "ParamChange", log); err != nil {
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
func (_Crosschain *CrosschainFilterer) ParseParamChange(log types.Log) (*CrosschainParamChange, error) {
	event := new(CrosschainParamChange)
	if err := _Crosschain.contract.UnpackLog(event, "ParamChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainProposalSubmittedIterator is returned from FilterProposalSubmitted and is used to iterate over the raw logs and unpacked data for ProposalSubmitted events raised by the Crosschain contract.
type CrosschainProposalSubmittedIterator struct {
	Event *CrosschainProposalSubmitted // Event containing the contract specifics and raw log

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
func (it *CrosschainProposalSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainProposalSubmitted)
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
		it.Event = new(CrosschainProposalSubmitted)
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
func (it *CrosschainProposalSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainProposalSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainProposalSubmitted represents a ProposalSubmitted event raised by the Crosschain contract.
type CrosschainProposalSubmitted struct {
	ProposalTypeHash [32]byte
	Proposer         common.Address
	Quorum           *big.Int
	ExpiredAt        *big.Int
	ContentHash      [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProposalSubmitted is a free log retrieval operation binding the contract event 0x9e109f0e55ef32e99e4880be2ec357f1ddb3469c79d0747ef4762da6e89fabe5.
//
// Solidity: event ProposalSubmitted(bytes32 indexed proposalTypeHash, address indexed proposer, uint128 quorum, uint128 expiredAt, bytes32 contentHash)
func (_Crosschain *CrosschainFilterer) FilterProposalSubmitted(opts *bind.FilterOpts, proposalTypeHash [][32]byte, proposer []common.Address) (*CrosschainProposalSubmittedIterator, error) {

	var proposalTypeHashRule []interface{}
	for _, proposalTypeHashItem := range proposalTypeHash {
		proposalTypeHashRule = append(proposalTypeHashRule, proposalTypeHashItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "ProposalSubmitted", proposalTypeHashRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainProposalSubmittedIterator{contract: _Crosschain.contract, event: "ProposalSubmitted", logs: logs, sub: sub}, nil
}

// WatchProposalSubmitted is a free log subscription operation binding the contract event 0x9e109f0e55ef32e99e4880be2ec357f1ddb3469c79d0747ef4762da6e89fabe5.
//
// Solidity: event ProposalSubmitted(bytes32 indexed proposalTypeHash, address indexed proposer, uint128 quorum, uint128 expiredAt, bytes32 contentHash)
func (_Crosschain *CrosschainFilterer) WatchProposalSubmitted(opts *bind.WatchOpts, sink chan<- *CrosschainProposalSubmitted, proposalTypeHash [][32]byte, proposer []common.Address) (event.Subscription, error) {

	var proposalTypeHashRule []interface{}
	for _, proposalTypeHashItem := range proposalTypeHash {
		proposalTypeHashRule = append(proposalTypeHashRule, proposalTypeHashItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "ProposalSubmitted", proposalTypeHashRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainProposalSubmitted)
				if err := _Crosschain.contract.UnpackLog(event, "ProposalSubmitted", log); err != nil {
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

// ParseProposalSubmitted is a log parse operation binding the contract event 0x9e109f0e55ef32e99e4880be2ec357f1ddb3469c79d0747ef4762da6e89fabe5.
//
// Solidity: event ProposalSubmitted(bytes32 indexed proposalTypeHash, address indexed proposer, uint128 quorum, uint128 expiredAt, bytes32 contentHash)
func (_Crosschain *CrosschainFilterer) ParseProposalSubmitted(log types.Log) (*CrosschainProposalSubmitted, error) {
	event := new(CrosschainProposalSubmitted)
	if err := _Crosschain.contract.UnpackLog(event, "ProposalSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainReceivedPackageIterator is returned from FilterReceivedPackage and is used to iterate over the raw logs and unpacked data for ReceivedPackage events raised by the Crosschain contract.
type CrosschainReceivedPackageIterator struct {
	Event *CrosschainReceivedPackage // Event containing the contract specifics and raw log

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
func (it *CrosschainReceivedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainReceivedPackage)
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
		it.Event = new(CrosschainReceivedPackage)
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
func (it *CrosschainReceivedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainReceivedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainReceivedPackage represents a ReceivedPackage event raised by the Crosschain contract.
type CrosschainReceivedPackage struct {
	PackageType     uint8
	PackageSequence uint64
	ChannelId       uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReceivedPackage is a free log retrieval operation binding the contract event 0x48484b8ae53584e6447d0535a274159337a74351c4adf243a6bf94b4c7a16c2e.
//
// Solidity: event ReceivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_Crosschain *CrosschainFilterer) FilterReceivedPackage(opts *bind.FilterOpts, packageSequence []uint64, channelId []uint8) (*CrosschainReceivedPackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "ReceivedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainReceivedPackageIterator{contract: _Crosschain.contract, event: "ReceivedPackage", logs: logs, sub: sub}, nil
}

// WatchReceivedPackage is a free log subscription operation binding the contract event 0x48484b8ae53584e6447d0535a274159337a74351c4adf243a6bf94b4c7a16c2e.
//
// Solidity: event ReceivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_Crosschain *CrosschainFilterer) WatchReceivedPackage(opts *bind.WatchOpts, sink chan<- *CrosschainReceivedPackage, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "ReceivedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainReceivedPackage)
				if err := _Crosschain.contract.UnpackLog(event, "ReceivedPackage", log); err != nil {
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

// ParseReceivedPackage is a log parse operation binding the contract event 0x48484b8ae53584e6447d0535a274159337a74351c4adf243a6bf94b4c7a16c2e.
//
// Solidity: event ReceivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_Crosschain *CrosschainFilterer) ParseReceivedPackage(log types.Log) (*CrosschainReceivedPackage, error) {
	event := new(CrosschainReceivedPackage)
	if err := _Crosschain.contract.UnpackLog(event, "ReceivedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainReopenedIterator is returned from FilterReopened and is used to iterate over the raw logs and unpacked data for Reopened events raised by the Crosschain contract.
type CrosschainReopenedIterator struct {
	Event *CrosschainReopened // Event containing the contract specifics and raw log

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
func (it *CrosschainReopenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainReopened)
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
		it.Event = new(CrosschainReopened)
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
func (it *CrosschainReopenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainReopenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainReopened represents a Reopened event raised by the Crosschain contract.
type CrosschainReopened struct {
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReopened is a free log retrieval operation binding the contract event 0x899fe8c37dc61708a3aaa99c4bf143346c1d1da69af79be9e8920c0a6785b752.
//
// Solidity: event Reopened(address indexed executor)
func (_Crosschain *CrosschainFilterer) FilterReopened(opts *bind.FilterOpts, executor []common.Address) (*CrosschainReopenedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "Reopened", executorRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainReopenedIterator{contract: _Crosschain.contract, event: "Reopened", logs: logs, sub: sub}, nil
}

// WatchReopened is a free log subscription operation binding the contract event 0x899fe8c37dc61708a3aaa99c4bf143346c1d1da69af79be9e8920c0a6785b752.
//
// Solidity: event Reopened(address indexed executor)
func (_Crosschain *CrosschainFilterer) WatchReopened(opts *bind.WatchOpts, sink chan<- *CrosschainReopened, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "Reopened", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainReopened)
				if err := _Crosschain.contract.UnpackLog(event, "Reopened", log); err != nil {
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

// ParseReopened is a log parse operation binding the contract event 0x899fe8c37dc61708a3aaa99c4bf143346c1d1da69af79be9e8920c0a6785b752.
//
// Solidity: event Reopened(address indexed executor)
func (_Crosschain *CrosschainFilterer) ParseReopened(log types.Log) (*CrosschainReopened, error) {
	event := new(CrosschainReopened)
	if err := _Crosschain.contract.UnpackLog(event, "Reopened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainSuspendedIterator is returned from FilterSuspended and is used to iterate over the raw logs and unpacked data for Suspended events raised by the Crosschain contract.
type CrosschainSuspendedIterator struct {
	Event *CrosschainSuspended // Event containing the contract specifics and raw log

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
func (it *CrosschainSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainSuspended)
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
		it.Event = new(CrosschainSuspended)
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
func (it *CrosschainSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainSuspended represents a Suspended event raised by the Crosschain contract.
type CrosschainSuspended struct {
	Executor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSuspended is a free log retrieval operation binding the contract event 0x6f123d3d54c84a7960a573b31c221dcd86e13fd849c5adb0c6ca851468cc1ae4.
//
// Solidity: event Suspended(address indexed executor)
func (_Crosschain *CrosschainFilterer) FilterSuspended(opts *bind.FilterOpts, executor []common.Address) (*CrosschainSuspendedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "Suspended", executorRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainSuspendedIterator{contract: _Crosschain.contract, event: "Suspended", logs: logs, sub: sub}, nil
}

// WatchSuspended is a free log subscription operation binding the contract event 0x6f123d3d54c84a7960a573b31c221dcd86e13fd849c5adb0c6ca851468cc1ae4.
//
// Solidity: event Suspended(address indexed executor)
func (_Crosschain *CrosschainFilterer) WatchSuspended(opts *bind.WatchOpts, sink chan<- *CrosschainSuspended, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "Suspended", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainSuspended)
				if err := _Crosschain.contract.UnpackLog(event, "Suspended", log); err != nil {
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

// ParseSuspended is a log parse operation binding the contract event 0x6f123d3d54c84a7960a573b31c221dcd86e13fd849c5adb0c6ca851468cc1ae4.
//
// Solidity: event Suspended(address indexed executor)
func (_Crosschain *CrosschainFilterer) ParseSuspended(log types.Log) (*CrosschainSuspended, error) {
	event := new(CrosschainSuspended)
	if err := _Crosschain.contract.UnpackLog(event, "Suspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainUnexpectedFailureAssertionInPackageHandlerIterator is returned from FilterUnexpectedFailureAssertionInPackageHandler and is used to iterate over the raw logs and unpacked data for UnexpectedFailureAssertionInPackageHandler events raised by the Crosschain contract.
type CrosschainUnexpectedFailureAssertionInPackageHandlerIterator struct {
	Event *CrosschainUnexpectedFailureAssertionInPackageHandler // Event containing the contract specifics and raw log

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
func (it *CrosschainUnexpectedFailureAssertionInPackageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainUnexpectedFailureAssertionInPackageHandler)
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
		it.Event = new(CrosschainUnexpectedFailureAssertionInPackageHandler)
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
func (it *CrosschainUnexpectedFailureAssertionInPackageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainUnexpectedFailureAssertionInPackageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainUnexpectedFailureAssertionInPackageHandler represents a UnexpectedFailureAssertionInPackageHandler event raised by the Crosschain contract.
type CrosschainUnexpectedFailureAssertionInPackageHandler struct {
	ContractAddr common.Address
	LowLevelData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedFailureAssertionInPackageHandler is a free log retrieval operation binding the contract event 0xfb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68.
//
// Solidity: event UnexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Crosschain *CrosschainFilterer) FilterUnexpectedFailureAssertionInPackageHandler(opts *bind.FilterOpts, contractAddr []common.Address) (*CrosschainUnexpectedFailureAssertionInPackageHandlerIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "UnexpectedFailureAssertionInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainUnexpectedFailureAssertionInPackageHandlerIterator{contract: _Crosschain.contract, event: "UnexpectedFailureAssertionInPackageHandler", logs: logs, sub: sub}, nil
}

// WatchUnexpectedFailureAssertionInPackageHandler is a free log subscription operation binding the contract event 0xfb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68.
//
// Solidity: event UnexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Crosschain *CrosschainFilterer) WatchUnexpectedFailureAssertionInPackageHandler(opts *bind.WatchOpts, sink chan<- *CrosschainUnexpectedFailureAssertionInPackageHandler, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "UnexpectedFailureAssertionInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainUnexpectedFailureAssertionInPackageHandler)
				if err := _Crosschain.contract.UnpackLog(event, "UnexpectedFailureAssertionInPackageHandler", log); err != nil {
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

// ParseUnexpectedFailureAssertionInPackageHandler is a log parse operation binding the contract event 0xfb111707dba108ae503813ed49a361e67f4720ac104ff7f5e9692e32039a1a68.
//
// Solidity: event UnexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Crosschain *CrosschainFilterer) ParseUnexpectedFailureAssertionInPackageHandler(log types.Log) (*CrosschainUnexpectedFailureAssertionInPackageHandler, error) {
	event := new(CrosschainUnexpectedFailureAssertionInPackageHandler)
	if err := _Crosschain.contract.UnpackLog(event, "UnexpectedFailureAssertionInPackageHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainUnexpectedRevertInPackageHandlerIterator is returned from FilterUnexpectedRevertInPackageHandler and is used to iterate over the raw logs and unpacked data for UnexpectedRevertInPackageHandler events raised by the Crosschain contract.
type CrosschainUnexpectedRevertInPackageHandlerIterator struct {
	Event *CrosschainUnexpectedRevertInPackageHandler // Event containing the contract specifics and raw log

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
func (it *CrosschainUnexpectedRevertInPackageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainUnexpectedRevertInPackageHandler)
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
		it.Event = new(CrosschainUnexpectedRevertInPackageHandler)
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
func (it *CrosschainUnexpectedRevertInPackageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainUnexpectedRevertInPackageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainUnexpectedRevertInPackageHandler represents a UnexpectedRevertInPackageHandler event raised by the Crosschain contract.
type CrosschainUnexpectedRevertInPackageHandler struct {
	ContractAddr common.Address
	Reason       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedRevertInPackageHandler is a free log retrieval operation binding the contract event 0xad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b136.
//
// Solidity: event UnexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Crosschain *CrosschainFilterer) FilterUnexpectedRevertInPackageHandler(opts *bind.FilterOpts, contractAddr []common.Address) (*CrosschainUnexpectedRevertInPackageHandlerIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "UnexpectedRevertInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainUnexpectedRevertInPackageHandlerIterator{contract: _Crosschain.contract, event: "UnexpectedRevertInPackageHandler", logs: logs, sub: sub}, nil
}

// WatchUnexpectedRevertInPackageHandler is a free log subscription operation binding the contract event 0xad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b136.
//
// Solidity: event UnexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Crosschain *CrosschainFilterer) WatchUnexpectedRevertInPackageHandler(opts *bind.WatchOpts, sink chan<- *CrosschainUnexpectedRevertInPackageHandler, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "UnexpectedRevertInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainUnexpectedRevertInPackageHandler)
				if err := _Crosschain.contract.UnpackLog(event, "UnexpectedRevertInPackageHandler", log); err != nil {
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

// ParseUnexpectedRevertInPackageHandler is a log parse operation binding the contract event 0xad48df68de34d2557f1ab71adc0bf7d0a1e23c433408e7b17ff51d588e28b136.
//
// Solidity: event UnexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Crosschain *CrosschainFilterer) ParseUnexpectedRevertInPackageHandler(log types.Log) (*CrosschainUnexpectedRevertInPackageHandler, error) {
	event := new(CrosschainUnexpectedRevertInPackageHandler)
	if err := _Crosschain.contract.UnpackLog(event, "UnexpectedRevertInPackageHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrosschainUnsupportedPackageIterator is returned from FilterUnsupportedPackage and is used to iterate over the raw logs and unpacked data for UnsupportedPackage events raised by the Crosschain contract.
type CrosschainUnsupportedPackageIterator struct {
	Event *CrosschainUnsupportedPackage // Event containing the contract specifics and raw log

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
func (it *CrosschainUnsupportedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainUnsupportedPackage)
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
		it.Event = new(CrosschainUnsupportedPackage)
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
func (it *CrosschainUnsupportedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainUnsupportedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainUnsupportedPackage represents a UnsupportedPackage event raised by the Crosschain contract.
type CrosschainUnsupportedPackage struct {
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnsupportedPackage is a free log retrieval operation binding the contract event 0xdee9845a11ea343955bc9858cfa6d6fbcce9c8c1cc4ca22a12997e88b1726b97.
//
// Solidity: event UnsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) FilterUnsupportedPackage(opts *bind.FilterOpts, packageSequence []uint64, channelId []uint8) (*CrosschainUnsupportedPackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "UnsupportedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainUnsupportedPackageIterator{contract: _Crosschain.contract, event: "UnsupportedPackage", logs: logs, sub: sub}, nil
}

// WatchUnsupportedPackage is a free log subscription operation binding the contract event 0xdee9845a11ea343955bc9858cfa6d6fbcce9c8c1cc4ca22a12997e88b1726b97.
//
// Solidity: event UnsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) WatchUnsupportedPackage(opts *bind.WatchOpts, sink chan<- *CrosschainUnsupportedPackage, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "UnsupportedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainUnsupportedPackage)
				if err := _Crosschain.contract.UnpackLog(event, "UnsupportedPackage", log); err != nil {
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

// ParseUnsupportedPackage is a log parse operation binding the contract event 0xdee9845a11ea343955bc9858cfa6d6fbcce9c8c1cc4ca22a12997e88b1726b97.
//
// Solidity: event UnsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) ParseUnsupportedPackage(log types.Log) (*CrosschainUnsupportedPackage, error) {
	event := new(CrosschainUnsupportedPackage)
	if err := _Crosschain.contract.UnpackLog(event, "UnsupportedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
