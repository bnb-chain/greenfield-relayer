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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"AddChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"chainId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"oracleSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"CrossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnable\",\"type\":\"bool\"}],\"name\":\"EnableOrDisableChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"ParamChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalTypeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"quorum\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"expiredAt\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"contentHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"ReceivedPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Reopened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"SuccessChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"Suspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"UnexpectedFailureAssertionInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"UnexpectedRevertInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"UnsupportedPackage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"APP_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCEL_TRANSFER_PROPOSAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMERGENCY_PROPOSAL_EXPIRE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMPTY_CONTENT_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAIL_ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BATCH_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_CANCEL_TRANSFER_QUORUM\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_REOPEN_QUORUM\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_SUSPEND_QUORUM\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INSCRIPTION_LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IN_TURN_RELAYER_VALIDITY_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUT_TURN_RELAYER_BACKOFF_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REOPEN_PROPOSAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUSPEND_PROPOSAL\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYN_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORSET_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchSizeForOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"attacker\",\"type\":\"address\"}],\"name\":\"cancelTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"challenged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelHandlerContractMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelReceiveSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelSendSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelSyncedHeaderMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"emergencyProposals\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"quorum\",\"type\":\"uint16\"},{\"internalType\":\"uint128\",\"name\":\"expiredAt\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"contentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"encodePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validatorSet\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"handlePackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_insChainId\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insChainId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"isRelayRewardFromSystemReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSuspended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleSequence\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousTxHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"quorumMap\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"registeredContractChannelMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reopen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"sendSynPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainCaller) APPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "APP_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainSession) APPCHANNELID() (uint8, error) {
	return _Crosschain.Contract.APPCHANNELID(&_Crosschain.CallOpts)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) APPCHANNELID() (uint8, error) {
	return _Crosschain.Contract.APPCHANNELID(&_Crosschain.CallOpts)
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

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Crosschain *CrosschainCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Crosschain *CrosschainSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.CROSSCHAINCONTRACTADDR(&_Crosschain.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Crosschain *CrosschainCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.CROSSCHAINCONTRACTADDR(&_Crosschain.CallOpts)
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

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() view returns(uint256)
func (_Crosschain *CrosschainCaller) INITBATCHSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "INIT_BATCH_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() view returns(uint256)
func (_Crosschain *CrosschainSession) INITBATCHSIZE() (*big.Int, error) {
	return _Crosschain.Contract.INITBATCHSIZE(&_Crosschain.CallOpts)
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) INITBATCHSIZE() (*big.Int, error) {
	return _Crosschain.Contract.INITBATCHSIZE(&_Crosschain.CallOpts)
}

// INITCANCELTRANSFERQUORUM is a free data retrieval call binding the contract method 0x6a3cb34d.
//
// Solidity: function INIT_CANCEL_TRANSFER_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainCaller) INITCANCELTRANSFERQUORUM(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "INIT_CANCEL_TRANSFER_QUORUM")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITCANCELTRANSFERQUORUM is a free data retrieval call binding the contract method 0x6a3cb34d.
//
// Solidity: function INIT_CANCEL_TRANSFER_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainSession) INITCANCELTRANSFERQUORUM() (uint16, error) {
	return _Crosschain.Contract.INITCANCELTRANSFERQUORUM(&_Crosschain.CallOpts)
}

// INITCANCELTRANSFERQUORUM is a free data retrieval call binding the contract method 0x6a3cb34d.
//
// Solidity: function INIT_CANCEL_TRANSFER_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainCallerSession) INITCANCELTRANSFERQUORUM() (uint16, error) {
	return _Crosschain.Contract.INITCANCELTRANSFERQUORUM(&_Crosschain.CallOpts)
}

// INITREOPENQUORUM is a free data retrieval call binding the contract method 0x6c46aa68.
//
// Solidity: function INIT_REOPEN_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainCaller) INITREOPENQUORUM(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "INIT_REOPEN_QUORUM")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITREOPENQUORUM is a free data retrieval call binding the contract method 0x6c46aa68.
//
// Solidity: function INIT_REOPEN_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainSession) INITREOPENQUORUM() (uint16, error) {
	return _Crosschain.Contract.INITREOPENQUORUM(&_Crosschain.CallOpts)
}

// INITREOPENQUORUM is a free data retrieval call binding the contract method 0x6c46aa68.
//
// Solidity: function INIT_REOPEN_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainCallerSession) INITREOPENQUORUM() (uint16, error) {
	return _Crosschain.Contract.INITREOPENQUORUM(&_Crosschain.CallOpts)
}

// INITSUSPENDQUORUM is a free data retrieval call binding the contract method 0x719482d5.
//
// Solidity: function INIT_SUSPEND_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainCaller) INITSUSPENDQUORUM(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "INIT_SUSPEND_QUORUM")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITSUSPENDQUORUM is a free data retrieval call binding the contract method 0x719482d5.
//
// Solidity: function INIT_SUSPEND_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainSession) INITSUSPENDQUORUM() (uint16, error) {
	return _Crosschain.Contract.INITSUSPENDQUORUM(&_Crosschain.CallOpts)
}

// INITSUSPENDQUORUM is a free data retrieval call binding the contract method 0x719482d5.
//
// Solidity: function INIT_SUSPEND_QUORUM() view returns(uint16)
func (_Crosschain *CrosschainCallerSession) INITSUSPENDQUORUM() (uint16, error) {
	return _Crosschain.Contract.INITSUSPENDQUORUM(&_Crosschain.CallOpts)
}

// INSCRIPTIONLIGHTCLIENTADDR is a free data retrieval call binding the contract method 0x550ec79c.
//
// Solidity: function INSCRIPTION_LIGHT_CLIENT_ADDR() view returns(address)
func (_Crosschain *CrosschainCaller) INSCRIPTIONLIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "INSCRIPTION_LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INSCRIPTIONLIGHTCLIENTADDR is a free data retrieval call binding the contract method 0x550ec79c.
//
// Solidity: function INSCRIPTION_LIGHT_CLIENT_ADDR() view returns(address)
func (_Crosschain *CrosschainSession) INSCRIPTIONLIGHTCLIENTADDR() (common.Address, error) {
	return _Crosschain.Contract.INSCRIPTIONLIGHTCLIENTADDR(&_Crosschain.CallOpts)
}

// INSCRIPTIONLIGHTCLIENTADDR is a free data retrieval call binding the contract method 0x550ec79c.
//
// Solidity: function INSCRIPTION_LIGHT_CLIENT_ADDR() view returns(address)
func (_Crosschain *CrosschainCallerSession) INSCRIPTIONLIGHTCLIENTADDR() (common.Address, error) {
	return _Crosschain.Contract.INSCRIPTIONLIGHTCLIENTADDR(&_Crosschain.CallOpts)
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

// OUTTURNRELAYERBACKOFFPERIOD is a free data retrieval call binding the contract method 0x34809881.
//
// Solidity: function OUT_TURN_RELAYER_BACKOFF_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainCaller) OUTTURNRELAYERBACKOFFPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "OUT_TURN_RELAYER_BACKOFF_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OUTTURNRELAYERBACKOFFPERIOD is a free data retrieval call binding the contract method 0x34809881.
//
// Solidity: function OUT_TURN_RELAYER_BACKOFF_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainSession) OUTTURNRELAYERBACKOFFPERIOD() (*big.Int, error) {
	return _Crosschain.Contract.OUTTURNRELAYERBACKOFFPERIOD(&_Crosschain.CallOpts)
}

// OUTTURNRELAYERBACKOFFPERIOD is a free data retrieval call binding the contract method 0x34809881.
//
// Solidity: function OUT_TURN_RELAYER_BACKOFF_PERIOD() view returns(uint256)
func (_Crosschain *CrosschainCallerSession) OUTTURNRELAYERBACKOFFPERIOD() (*big.Int, error) {
	return _Crosschain.Contract.OUTTURNRELAYERBACKOFFPERIOD(&_Crosschain.CallOpts)
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

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Crosschain *CrosschainCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Crosschain *CrosschainSession) TOKENHUBADDR() (common.Address, error) {
	return _Crosschain.Contract.TOKENHUBADDR(&_Crosschain.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Crosschain *CrosschainCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Crosschain.Contract.TOKENHUBADDR(&_Crosschain.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFERINCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFERINCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFEROUTCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFEROUTCHANNELID(&_Crosschain.CallOpts)
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainCaller) VALIDATORSETCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "VALIDATORSET_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainSession) VALIDATORSETCHANNELID() (uint8, error) {
	return _Crosschain.Contract.VALIDATORSETCHANNELID(&_Crosschain.CallOpts)
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Crosschain *CrosschainCallerSession) VALIDATORSETCHANNELID() (uint8, error) {
	return _Crosschain.Contract.VALIDATORSETCHANNELID(&_Crosschain.CallOpts)
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
// Solidity: function chainId() view returns(uint32)
func (_Crosschain *CrosschainCaller) ChainId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint32)
func (_Crosschain *CrosschainSession) ChainId() (uint32, error) {
	return _Crosschain.Contract.ChainId(&_Crosschain.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint32)
func (_Crosschain *CrosschainCallerSession) ChainId() (uint32, error) {
	return _Crosschain.Contract.ChainId(&_Crosschain.CallOpts)
}

// Challenged is a free data retrieval call binding the contract method 0x2af6f399.
//
// Solidity: function challenged(bytes32 ) view returns(bool)
func (_Crosschain *CrosschainCaller) Challenged(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "challenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Challenged is a free data retrieval call binding the contract method 0x2af6f399.
//
// Solidity: function challenged(bytes32 ) view returns(bool)
func (_Crosschain *CrosschainSession) Challenged(arg0 [32]byte) (bool, error) {
	return _Crosschain.Contract.Challenged(&_Crosschain.CallOpts, arg0)
}

// Challenged is a free data retrieval call binding the contract method 0x2af6f399.
//
// Solidity: function challenged(bytes32 ) view returns(bool)
func (_Crosschain *CrosschainCallerSession) Challenged(arg0 [32]byte) (bool, error) {
	return _Crosschain.Contract.Challenged(&_Crosschain.CallOpts, arg0)
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) view returns(address)
func (_Crosschain *CrosschainCaller) ChannelHandlerContractMap(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "channelHandlerContractMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) view returns(address)
func (_Crosschain *CrosschainSession) ChannelHandlerContractMap(arg0 uint8) (common.Address, error) {
	return _Crosschain.Contract.ChannelHandlerContractMap(&_Crosschain.CallOpts, arg0)
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) view returns(address)
func (_Crosschain *CrosschainCallerSession) ChannelHandlerContractMap(arg0 uint8) (common.Address, error) {
	return _Crosschain.Contract.ChannelHandlerContractMap(&_Crosschain.CallOpts, arg0)
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

// ChannelSyncedHeaderMap is a free data retrieval call binding the contract method 0x3a648b15.
//
// Solidity: function channelSyncedHeaderMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainCaller) ChannelSyncedHeaderMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "channelSyncedHeaderMap", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChannelSyncedHeaderMap is a free data retrieval call binding the contract method 0x3a648b15.
//
// Solidity: function channelSyncedHeaderMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainSession) ChannelSyncedHeaderMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelSyncedHeaderMap(&_Crosschain.CallOpts, arg0)
}

// ChannelSyncedHeaderMap is a free data retrieval call binding the contract method 0x3a648b15.
//
// Solidity: function channelSyncedHeaderMap(uint8 ) view returns(uint64)
func (_Crosschain *CrosschainCallerSession) ChannelSyncedHeaderMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelSyncedHeaderMap(&_Crosschain.CallOpts, arg0)
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

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) view returns(bytes)
func (_Crosschain *CrosschainCaller) EncodePayload(opts *bind.CallOpts, packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "encodePayload", packageType, relayFee, msgBytes)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) view returns(bytes)
func (_Crosschain *CrosschainSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _Crosschain.Contract.EncodePayload(&_Crosschain.CallOpts, packageType, relayFee, msgBytes)
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) view returns(bytes)
func (_Crosschain *CrosschainCallerSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _Crosschain.Contract.EncodePayload(&_Crosschain.CallOpts, packageType, relayFee, msgBytes)
}

// InsChainId is a free data retrieval call binding the contract method 0x70b6dca6.
//
// Solidity: function insChainId() view returns(uint32)
func (_Crosschain *CrosschainCaller) InsChainId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "insChainId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// InsChainId is a free data retrieval call binding the contract method 0x70b6dca6.
//
// Solidity: function insChainId() view returns(uint32)
func (_Crosschain *CrosschainSession) InsChainId() (uint32, error) {
	return _Crosschain.Contract.InsChainId(&_Crosschain.CallOpts)
}

// InsChainId is a free data retrieval call binding the contract method 0x70b6dca6.
//
// Solidity: function insChainId() view returns(uint32)
func (_Crosschain *CrosschainCallerSession) InsChainId() (uint32, error) {
	return _Crosschain.Contract.InsChainId(&_Crosschain.CallOpts)
}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) view returns(bool)
func (_Crosschain *CrosschainCaller) IsRelayRewardFromSystemReward(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "isRelayRewardFromSystemReward", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) view returns(bool)
func (_Crosschain *CrosschainSession) IsRelayRewardFromSystemReward(arg0 uint8) (bool, error) {
	return _Crosschain.Contract.IsRelayRewardFromSystemReward(&_Crosschain.CallOpts, arg0)
}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) view returns(bool)
func (_Crosschain *CrosschainCallerSession) IsRelayRewardFromSystemReward(arg0 uint8) (bool, error) {
	return _Crosschain.Contract.IsRelayRewardFromSystemReward(&_Crosschain.CallOpts, arg0)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crosschain *CrosschainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crosschain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crosschain *CrosschainSession) Owner() (common.Address, error) {
	return _Crosschain.Contract.Owner(&_Crosschain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crosschain *CrosschainCallerSession) Owner() (common.Address, error) {
	return _Crosschain.Contract.Owner(&_Crosschain.CallOpts)
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

// CancelTransfer is a paid mutator transaction binding the contract method 0x5f832177.
//
// Solidity: function cancelTransfer(address tokenAddr, address attacker) returns()
func (_Crosschain *CrosschainTransactor) CancelTransfer(opts *bind.TransactOpts, tokenAddr common.Address, attacker common.Address) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "cancelTransfer", tokenAddr, attacker)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0x5f832177.
//
// Solidity: function cancelTransfer(address tokenAddr, address attacker) returns()
func (_Crosschain *CrosschainSession) CancelTransfer(tokenAddr common.Address, attacker common.Address) (*types.Transaction, error) {
	return _Crosschain.Contract.CancelTransfer(&_Crosschain.TransactOpts, tokenAddr, attacker)
}

// CancelTransfer is a paid mutator transaction binding the contract method 0x5f832177.
//
// Solidity: function cancelTransfer(address tokenAddr, address attacker) returns()
func (_Crosschain *CrosschainTransactorSession) CancelTransfer(tokenAddr common.Address, attacker common.Address) (*types.Transaction, error) {
	return _Crosschain.Contract.CancelTransfer(&_Crosschain.TransactOpts, tokenAddr, attacker)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x7bffa585.
//
// Solidity: function handlePackage(bytes payload, bytes blsSignature, uint256 validatorSet, uint64 packageSequence, uint8 channelId) returns()
func (_Crosschain *CrosschainTransactor) HandlePackage(opts *bind.TransactOpts, payload []byte, blsSignature []byte, validatorSet *big.Int, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "handlePackage", payload, blsSignature, validatorSet, packageSequence, channelId)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x7bffa585.
//
// Solidity: function handlePackage(bytes payload, bytes blsSignature, uint256 validatorSet, uint64 packageSequence, uint8 channelId) returns()
func (_Crosschain *CrosschainSession) HandlePackage(payload []byte, blsSignature []byte, validatorSet *big.Int, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Crosschain.Contract.HandlePackage(&_Crosschain.TransactOpts, payload, blsSignature, validatorSet, packageSequence, channelId)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x7bffa585.
//
// Solidity: function handlePackage(bytes payload, bytes blsSignature, uint256 validatorSet, uint64 packageSequence, uint8 channelId) returns()
func (_Crosschain *CrosschainTransactorSession) HandlePackage(payload []byte, blsSignature []byte, validatorSet *big.Int, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Crosschain.Contract.HandlePackage(&_Crosschain.TransactOpts, payload, blsSignature, validatorSet, packageSequence, channelId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8d8965bd.
//
// Solidity: function initialize(uint32 _insChainId) returns()
func (_Crosschain *CrosschainTransactor) Initialize(opts *bind.TransactOpts, _insChainId uint32) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "initialize", _insChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8d8965bd.
//
// Solidity: function initialize(uint32 _insChainId) returns()
func (_Crosschain *CrosschainSession) Initialize(_insChainId uint32) (*types.Transaction, error) {
	return _Crosschain.Contract.Initialize(&_Crosschain.TransactOpts, _insChainId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8d8965bd.
//
// Solidity: function initialize(uint32 _insChainId) returns()
func (_Crosschain *CrosschainTransactorSession) Initialize(_insChainId uint32) (*types.Transaction, error) {
	return _Crosschain.Contract.Initialize(&_Crosschain.TransactOpts, _insChainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crosschain *CrosschainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crosschain *CrosschainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crosschain.Contract.RenounceOwnership(&_Crosschain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crosschain *CrosschainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crosschain.Contract.RenounceOwnership(&_Crosschain.TransactOpts)
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

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_Crosschain *CrosschainTransactor) SendSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "sendSynPackage", channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_Crosschain *CrosschainSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.SendSynPackage(&_Crosschain.TransactOpts, channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_Crosschain *CrosschainTransactorSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.SendSynPackage(&_Crosschain.TransactOpts, channelId, msgBytes, relayFee)
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

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crosschain *CrosschainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crosschain *CrosschainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crosschain.Contract.TransferOwnership(&_Crosschain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crosschain *CrosschainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crosschain.Contract.TransferOwnership(&_Crosschain.TransactOpts, newOwner)
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
	ChainId         uint32
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCrossChainPackage is a free log retrieval operation binding the contract event 0x1cd5706f63b9e9177bd01b759fa8bcbd7452d79f49564943ed85a52f57bef179.
//
// Solidity: event CrossChainPackage(uint32 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
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

// WatchCrossChainPackage is a free log subscription operation binding the contract event 0x1cd5706f63b9e9177bd01b759fa8bcbd7452d79f49564943ed85a52f57bef179.
//
// Solidity: event CrossChainPackage(uint32 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
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

// ParseCrossChainPackage is a log parse operation binding the contract event 0x1cd5706f63b9e9177bd01b759fa8bcbd7452d79f49564943ed85a52f57bef179.
//
// Solidity: event CrossChainPackage(uint32 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
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

// CrosschainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crosschain contract.
type CrosschainOwnershipTransferredIterator struct {
	Event *CrosschainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrosschainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainOwnershipTransferred)
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
		it.Event = new(CrosschainOwnershipTransferred)
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
func (it *CrosschainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainOwnershipTransferred represents a OwnershipTransferred event raised by the Crosschain contract.
type CrosschainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crosschain *CrosschainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrosschainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainOwnershipTransferredIterator{contract: _Crosschain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crosschain *CrosschainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrosschainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainOwnershipTransferred)
				if err := _Crosschain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Crosschain *CrosschainFilterer) ParseOwnershipTransferred(log types.Log) (*CrosschainOwnershipTransferred, error) {
	event := new(CrosschainOwnershipTransferred)
	if err := _Crosschain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// CrosschainSuccessChallengeIterator is returned from FilterSuccessChallenge and is used to iterate over the raw logs and unpacked data for SuccessChallenge events raised by the Crosschain contract.
type CrosschainSuccessChallengeIterator struct {
	Event *CrosschainSuccessChallenge // Event containing the contract specifics and raw log

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
func (it *CrosschainSuccessChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainSuccessChallenge)
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
		it.Event = new(CrosschainSuccessChallenge)
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
func (it *CrosschainSuccessChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainSuccessChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainSuccessChallenge represents a SuccessChallenge event raised by the Crosschain contract.
type CrosschainSuccessChallenge struct {
	Challenger      common.Address
	PackageSequence uint64
	ChannelId       uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSuccessChallenge is a free log retrieval operation binding the contract event 0x039eb91179ffd7d3b6e97f8ea106e748e827f910b872375dbc9c14a362319c3c.
//
// Solidity: event SuccessChallenge(address indexed challenger, uint64 packageSequence, uint8 channelId)
func (_Crosschain *CrosschainFilterer) FilterSuccessChallenge(opts *bind.FilterOpts, challenger []common.Address) (*CrosschainSuccessChallengeIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "SuccessChallenge", challengerRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainSuccessChallengeIterator{contract: _Crosschain.contract, event: "SuccessChallenge", logs: logs, sub: sub}, nil
}

// WatchSuccessChallenge is a free log subscription operation binding the contract event 0x039eb91179ffd7d3b6e97f8ea106e748e827f910b872375dbc9c14a362319c3c.
//
// Solidity: event SuccessChallenge(address indexed challenger, uint64 packageSequence, uint8 channelId)
func (_Crosschain *CrosschainFilterer) WatchSuccessChallenge(opts *bind.WatchOpts, sink chan<- *CrosschainSuccessChallenge, challenger []common.Address) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "SuccessChallenge", challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainSuccessChallenge)
				if err := _Crosschain.contract.UnpackLog(event, "SuccessChallenge", log); err != nil {
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

// ParseSuccessChallenge is a log parse operation binding the contract event 0x039eb91179ffd7d3b6e97f8ea106e748e827f910b872375dbc9c14a362319c3c.
//
// Solidity: event SuccessChallenge(address indexed challenger, uint64 packageSequence, uint8 channelId)
func (_Crosschain *CrosschainFilterer) ParseSuccessChallenge(log types.Log) (*CrosschainSuccessChallenge, error) {
	event := new(CrosschainSuccessChallenge)
	if err := _Crosschain.contract.UnpackLog(event, "SuccessChallenge", log); err != nil {
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
