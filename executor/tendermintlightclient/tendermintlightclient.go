// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tendermintlightclient

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

// TendermintlightclientMetaData contains all meta data concerning the Tendermintlightclient contract.
var TendermintlightclientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"APP_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBKEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INSCRIPTION_LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PACKAGE_VERIFY_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORSET_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsPubKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"height\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relayers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_header\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_blsPubKeys\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"_relayers\",\"type\":\"address[]\"}],\"name\":\"syncTendermintHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pkgKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetBitMap\",\"type\":\"uint256\"}],\"name\":\"verifyPackage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TendermintlightclientABI is the input ABI used to generate the binding from.
// Deprecated: Use TendermintlightclientMetaData.ABI instead.
var TendermintlightclientABI = TendermintlightclientMetaData.ABI

// Tendermintlightclient is an auto generated Go binding around an Ethereum contract.
type Tendermintlightclient struct {
	TendermintlightclientCaller     // Read-only binding to the contract
	TendermintlightclientTransactor // Write-only binding to the contract
	TendermintlightclientFilterer   // Log filterer for contract events
}

// TendermintlightclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TendermintlightclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintlightclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TendermintlightclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintlightclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TendermintlightclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintlightclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TendermintlightclientSession struct {
	Contract     *Tendermintlightclient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TendermintlightclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TendermintlightclientCallerSession struct {
	Contract *TendermintlightclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TendermintlightclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TendermintlightclientTransactorSession struct {
	Contract     *TendermintlightclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TendermintlightclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TendermintlightclientRaw struct {
	Contract *Tendermintlightclient // Generic contract binding to access the raw methods on
}

// TendermintlightclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TendermintlightclientCallerRaw struct {
	Contract *TendermintlightclientCaller // Generic read-only contract binding to access the raw methods on
}

// TendermintlightclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TendermintlightclientTransactorRaw struct {
	Contract *TendermintlightclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTendermintlightclient creates a new instance of Tendermintlightclient, bound to a specific deployed contract.
func NewTendermintlightclient(address common.Address, backend bind.ContractBackend) (*Tendermintlightclient, error) {
	contract, err := bindTendermintlightclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tendermintlightclient{TendermintlightclientCaller: TendermintlightclientCaller{contract: contract}, TendermintlightclientTransactor: TendermintlightclientTransactor{contract: contract}, TendermintlightclientFilterer: TendermintlightclientFilterer{contract: contract}}, nil
}

// NewTendermintlightclientCaller creates a new read-only instance of Tendermintlightclient, bound to a specific deployed contract.
func NewTendermintlightclientCaller(address common.Address, caller bind.ContractCaller) (*TendermintlightclientCaller, error) {
	contract, err := bindTendermintlightclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientCaller{contract: contract}, nil
}

// NewTendermintlightclientTransactor creates a new write-only instance of Tendermintlightclient, bound to a specific deployed contract.
func NewTendermintlightclientTransactor(address common.Address, transactor bind.ContractTransactor) (*TendermintlightclientTransactor, error) {
	contract, err := bindTendermintlightclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientTransactor{contract: contract}, nil
}

// NewTendermintlightclientFilterer creates a new log filterer instance of Tendermintlightclient, bound to a specific deployed contract.
func NewTendermintlightclientFilterer(address common.Address, filterer bind.ContractFilterer) (*TendermintlightclientFilterer, error) {
	contract, err := bindTendermintlightclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientFilterer{contract: contract}, nil
}

// bindTendermintlightclient binds a generic wrapper to an already deployed contract.
func bindTendermintlightclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TendermintlightclientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tendermintlightclient *TendermintlightclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tendermintlightclient.Contract.TendermintlightclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tendermintlightclient *TendermintlightclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.TendermintlightclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tendermintlightclient *TendermintlightclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.TendermintlightclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tendermintlightclient *TendermintlightclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tendermintlightclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tendermintlightclient *TendermintlightclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tendermintlightclient *TendermintlightclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.contract.Transact(opts, method, params...)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) APPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "APP_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) APPCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.APPCHANNELID(&_Tendermintlightclient.CallOpts)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) APPCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.APPCHANNELID(&_Tendermintlightclient.CallOpts)
}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientCaller) BLSPUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "BLS_PUBKEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientSession) BLSPUBKEYLENGTH() (*big.Int, error) {
	return _Tendermintlightclient.Contract.BLSPUBKEYLENGTH(&_Tendermintlightclient.CallOpts)
}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientCallerSession) BLSPUBKEYLENGTH() (*big.Int, error) {
	return _Tendermintlightclient.Contract.BLSPUBKEYLENGTH(&_Tendermintlightclient.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.CROSSCHAINCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.CROSSCHAINCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// INSCRIPTIONLIGHTCLIENTADDR is a free data retrieval call binding the contract method 0x550ec79c.
//
// Solidity: function INSCRIPTION_LIGHT_CLIENT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) INSCRIPTIONLIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "INSCRIPTION_LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INSCRIPTIONLIGHTCLIENTADDR is a free data retrieval call binding the contract method 0x550ec79c.
//
// Solidity: function INSCRIPTION_LIGHT_CLIENT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) INSCRIPTIONLIGHTCLIENTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.INSCRIPTIONLIGHTCLIENTADDR(&_Tendermintlightclient.CallOpts)
}

// INSCRIPTIONLIGHTCLIENTADDR is a free data retrieval call binding the contract method 0x550ec79c.
//
// Solidity: function INSCRIPTION_LIGHT_CLIENT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) INSCRIPTIONLIGHTCLIENTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.INSCRIPTIONLIGHTCLIENTADDR(&_Tendermintlightclient.CallOpts)
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) LIGHTCLIENTCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "LIGHT_CLIENT_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) LIGHTCLIENTCONTRACT() (common.Address, error) {
	return _Tendermintlightclient.Contract.LIGHTCLIENTCONTRACT(&_Tendermintlightclient.CallOpts)
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) LIGHTCLIENTCONTRACT() (common.Address, error) {
	return _Tendermintlightclient.Contract.LIGHTCLIENTCONTRACT(&_Tendermintlightclient.CallOpts)
}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) PACKAGEVERIFYCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "PACKAGE_VERIFY_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) PACKAGEVERIFYCONTRACT() (common.Address, error) {
	return _Tendermintlightclient.Contract.PACKAGEVERIFYCONTRACT(&_Tendermintlightclient.CallOpts)
}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) PACKAGEVERIFYCONTRACT() (common.Address, error) {
	return _Tendermintlightclient.Contract.PACKAGEVERIFYCONTRACT(&_Tendermintlightclient.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) TOKENHUBADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.TOKENHUBADDR(&_Tendermintlightclient.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.TOKENHUBADDR(&_Tendermintlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.TRANSFERINCHANNELID(&_Tendermintlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.TRANSFERINCHANNELID(&_Tendermintlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.TRANSFEROUTCHANNELID(&_Tendermintlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.TRANSFEROUTCHANNELID(&_Tendermintlightclient.CallOpts)
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) VALIDATORSETCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "VALIDATORSET_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) VALIDATORSETCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.VALIDATORSETCHANNELID(&_Tendermintlightclient.CallOpts)
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) VALIDATORSETCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.VALIDATORSETCHANNELID(&_Tendermintlightclient.CallOpts)
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Tendermintlightclient *TendermintlightclientCaller) BlsPubKeys(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "blsPubKeys")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Tendermintlightclient *TendermintlightclientSession) BlsPubKeys() ([]byte, error) {
	return _Tendermintlightclient.Contract.BlsPubKeys(&_Tendermintlightclient.CallOpts)
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Tendermintlightclient *TendermintlightclientCallerSession) BlsPubKeys() ([]byte, error) {
	return _Tendermintlightclient.Contract.BlsPubKeys(&_Tendermintlightclient.CallOpts)
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Tendermintlightclient *TendermintlightclientCaller) GetRelayers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "getRelayers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Tendermintlightclient *TendermintlightclientSession) GetRelayers() ([]common.Address, error) {
	return _Tendermintlightclient.Contract.GetRelayers(&_Tendermintlightclient.CallOpts)
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Tendermintlightclient *TendermintlightclientCallerSession) GetRelayers() ([]common.Address, error) {
	return _Tendermintlightclient.Contract.GetRelayers(&_Tendermintlightclient.CallOpts)
}

// Height is a free data retrieval call binding the contract method 0x0ef26743.
//
// Solidity: function height() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientCaller) Height(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "height")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Height is a free data retrieval call binding the contract method 0x0ef26743.
//
// Solidity: function height() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientSession) Height() (uint64, error) {
	return _Tendermintlightclient.Contract.Height(&_Tendermintlightclient.CallOpts)
}

// Height is a free data retrieval call binding the contract method 0x0ef26743.
//
// Solidity: function height() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientCallerSession) Height() (uint64, error) {
	return _Tendermintlightclient.Contract.Height(&_Tendermintlightclient.CallOpts)
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) Relayers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "relayers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) Relayers(arg0 *big.Int) (common.Address, error) {
	return _Tendermintlightclient.Contract.Relayers(&_Tendermintlightclient.CallOpts, arg0)
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) Relayers(arg0 *big.Int) (common.Address, error) {
	return _Tendermintlightclient.Contract.Relayers(&_Tendermintlightclient.CallOpts, arg0)
}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Tendermintlightclient *TendermintlightclientCaller) VerifyPackage(opts *bind.CallOpts, _pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "verifyPackage", _pkgKey, _payload, _blsSignature, _validatorSetBitMap)

	if err != nil {
		return err
	}

	return err

}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Tendermintlightclient *TendermintlightclientSession) VerifyPackage(_pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	return _Tendermintlightclient.Contract.VerifyPackage(&_Tendermintlightclient.CallOpts, _pkgKey, _payload, _blsSignature, _validatorSetBitMap)
}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Tendermintlightclient *TendermintlightclientCallerSession) VerifyPackage(_pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	return _Tendermintlightclient.Contract.VerifyPackage(&_Tendermintlightclient.CallOpts, _pkgKey, _payload, _blsSignature, _validatorSetBitMap)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Tendermintlightclient *TendermintlightclientTransactor) SyncTendermintHeader(opts *bind.TransactOpts, _header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Tendermintlightclient.contract.Transact(opts, "syncTendermintHeader", _header, _height, _blsPubKeys, _relayers)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Tendermintlightclient *TendermintlightclientSession) SyncTendermintHeader(_header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.SyncTendermintHeader(&_Tendermintlightclient.TransactOpts, _header, _height, _blsPubKeys, _relayers)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Tendermintlightclient *TendermintlightclientTransactorSession) SyncTendermintHeader(_header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.SyncTendermintHeader(&_Tendermintlightclient.TransactOpts, _header, _height, _blsPubKeys, _relayers)
}
