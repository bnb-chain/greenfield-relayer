// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inscriptionlightclient

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

// InscriptionlightclientMetaData contains all meta data concerning the Inscriptionlightclient contract.
var InscriptionlightclientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"APP_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLS_PUBKEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PACKAGE_VERIFY_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATORSET_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsPubKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRelayers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"relayers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_header\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_blsPubKeys\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"_relayers\",\"type\":\"address[]\"}],\"name\":\"syncTendermintHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pkgKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorSetBitMap\",\"type\":\"uint256\"}],\"name\":\"verifyPackage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InscriptionlightclientABI is the input ABI used to generate the binding from.
// Deprecated: Use InscriptionlightclientMetaData.ABI instead.
var InscriptionlightclientABI = InscriptionlightclientMetaData.ABI

// Inscriptionlightclient is an auto generated Go binding around an Ethereum contract.
type Inscriptionlightclient struct {
	InscriptionlightclientCaller     // Read-only binding to the contract
	InscriptionlightclientTransactor // Write-only binding to the contract
	InscriptionlightclientFilterer   // Log filterer for contract events
}

// InscriptionlightclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type InscriptionlightclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InscriptionlightclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InscriptionlightclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InscriptionlightclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InscriptionlightclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InscriptionlightclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InscriptionlightclientSession struct {
	Contract     *Inscriptionlightclient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InscriptionlightclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InscriptionlightclientCallerSession struct {
	Contract *InscriptionlightclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// InscriptionlightclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InscriptionlightclientTransactorSession struct {
	Contract     *InscriptionlightclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// InscriptionlightclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type InscriptionlightclientRaw struct {
	Contract *Inscriptionlightclient // Generic contract binding to access the raw methods on
}

// InscriptionlightclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InscriptionlightclientCallerRaw struct {
	Contract *InscriptionlightclientCaller // Generic read-only contract binding to access the raw methods on
}

// InscriptionlightclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InscriptionlightclientTransactorRaw struct {
	Contract *InscriptionlightclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInscriptionlightclient creates a new instance of Inscriptionlightclient, bound to a specific deployed contract.
func NewInscriptionlightclient(address common.Address, backend bind.ContractBackend) (*Inscriptionlightclient, error) {
	contract, err := bindInscriptionlightclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inscriptionlightclient{InscriptionlightclientCaller: InscriptionlightclientCaller{contract: contract}, InscriptionlightclientTransactor: InscriptionlightclientTransactor{contract: contract}, InscriptionlightclientFilterer: InscriptionlightclientFilterer{contract: contract}}, nil
}

// NewInscriptionlightclientCaller creates a new read-only instance of Inscriptionlightclient, bound to a specific deployed contract.
func NewInscriptionlightclientCaller(address common.Address, caller bind.ContractCaller) (*InscriptionlightclientCaller, error) {
	contract, err := bindInscriptionlightclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InscriptionlightclientCaller{contract: contract}, nil
}

// NewInscriptionlightclientTransactor creates a new write-only instance of Inscriptionlightclient, bound to a specific deployed contract.
func NewInscriptionlightclientTransactor(address common.Address, transactor bind.ContractTransactor) (*InscriptionlightclientTransactor, error) {
	contract, err := bindInscriptionlightclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InscriptionlightclientTransactor{contract: contract}, nil
}

// NewInscriptionlightclientFilterer creates a new log filterer instance of Inscriptionlightclient, bound to a specific deployed contract.
func NewInscriptionlightclientFilterer(address common.Address, filterer bind.ContractFilterer) (*InscriptionlightclientFilterer, error) {
	contract, err := bindInscriptionlightclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InscriptionlightclientFilterer{contract: contract}, nil
}

// bindInscriptionlightclient binds a generic wrapper to an already deployed contract.
func bindInscriptionlightclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InscriptionlightclientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inscriptionlightclient *InscriptionlightclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inscriptionlightclient.Contract.InscriptionlightclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inscriptionlightclient *InscriptionlightclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inscriptionlightclient.Contract.InscriptionlightclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inscriptionlightclient *InscriptionlightclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inscriptionlightclient.Contract.InscriptionlightclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inscriptionlightclient *InscriptionlightclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inscriptionlightclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inscriptionlightclient *InscriptionlightclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inscriptionlightclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inscriptionlightclient *InscriptionlightclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inscriptionlightclient.Contract.contract.Transact(opts, method, params...)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCaller) APPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "APP_CHANNELID")
	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientSession) APPCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.APPCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) APPCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.APPCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Inscriptionlightclient *InscriptionlightclientCaller) BLSPUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "BLS_PUBKEY_LENGTH")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Inscriptionlightclient *InscriptionlightclientSession) BLSPUBKEYLENGTH() (*big.Int, error) {
	return _Inscriptionlightclient.Contract.BLSPUBKEYLENGTH(&_Inscriptionlightclient.CallOpts)
}

// BLSPUBKEYLENGTH is a free data retrieval call binding the contract method 0x0c20fe41.
//
// Solidity: function BLS_PUBKEY_LENGTH() view returns(uint256)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) BLSPUBKEYLENGTH() (*big.Int, error) {
	return _Inscriptionlightclient.Contract.BLSPUBKEYLENGTH(&_Inscriptionlightclient.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Inscriptionlightclient *InscriptionlightclientCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "CODE_OK")
	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Inscriptionlightclient *InscriptionlightclientSession) CODEOK() (uint32, error) {
	return _Inscriptionlightclient.Contract.CODEOK(&_Inscriptionlightclient.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) CODEOK() (uint32, error) {
	return _Inscriptionlightclient.Contract.CODEOK(&_Inscriptionlightclient.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Inscriptionlightclient *InscriptionlightclientCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "ERROR_FAIL_DECODE")
	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Inscriptionlightclient *InscriptionlightclientSession) ERRORFAILDECODE() (uint32, error) {
	return _Inscriptionlightclient.Contract.ERRORFAILDECODE(&_Inscriptionlightclient.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Inscriptionlightclient.Contract.ERRORFAILDECODE(&_Inscriptionlightclient.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "GOV_CHANNELID")
	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientSession) GOVCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.GOVCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) GOVCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.GOVCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientCaller) LIGHTCLIENTCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "LIGHT_CLIENT_CONTRACT")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientSession) LIGHTCLIENTCONTRACT() (common.Address, error) {
	return _Inscriptionlightclient.Contract.LIGHTCLIENTCONTRACT(&_Inscriptionlightclient.CallOpts)
}

// LIGHTCLIENTCONTRACT is a free data retrieval call binding the contract method 0xdf8193da.
//
// Solidity: function LIGHT_CLIENT_CONTRACT() view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) LIGHTCLIENTCONTRACT() (common.Address, error) {
	return _Inscriptionlightclient.Contract.LIGHTCLIENTCONTRACT(&_Inscriptionlightclient.CallOpts)
}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientCaller) PACKAGEVERIFYCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "PACKAGE_VERIFY_CONTRACT")
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientSession) PACKAGEVERIFYCONTRACT() (common.Address, error) {
	return _Inscriptionlightclient.Contract.PACKAGEVERIFYCONTRACT(&_Inscriptionlightclient.CallOpts)
}

// PACKAGEVERIFYCONTRACT is a free data retrieval call binding the contract method 0x86a780c5.
//
// Solidity: function PACKAGE_VERIFY_CONTRACT() view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) PACKAGEVERIFYCONTRACT() (common.Address, error) {
	return _Inscriptionlightclient.Contract.PACKAGEVERIFYCONTRACT(&_Inscriptionlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")
	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.TRANSFERINCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.TRANSFERINCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")
	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.TRANSFEROUTCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.TRANSFEROUTCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCaller) VALIDATORSETCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "VALIDATORSET_CHANNELID")
	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientSession) VALIDATORSETCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.VALIDATORSETCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// VALIDATORSETCHANNELID is a free data retrieval call binding the contract method 0x99770281.
//
// Solidity: function VALIDATORSET_CHANNELID() view returns(uint8)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) VALIDATORSETCHANNELID() (uint8, error) {
	return _Inscriptionlightclient.Contract.VALIDATORSETCHANNELID(&_Inscriptionlightclient.CallOpts)
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Inscriptionlightclient *InscriptionlightclientCaller) BlsPubKeys(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "blsPubKeys")
	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Inscriptionlightclient *InscriptionlightclientSession) BlsPubKeys() ([]byte, error) {
	return _Inscriptionlightclient.Contract.BlsPubKeys(&_Inscriptionlightclient.CallOpts)
}

// BlsPubKeys is a free data retrieval call binding the contract method 0xbb9dae06.
//
// Solidity: function blsPubKeys() view returns(bytes)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) BlsPubKeys() ([]byte, error) {
	return _Inscriptionlightclient.Contract.BlsPubKeys(&_Inscriptionlightclient.CallOpts)
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Inscriptionlightclient *InscriptionlightclientCaller) GetRelayers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "getRelayers")
	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Inscriptionlightclient *InscriptionlightclientSession) GetRelayers() ([]common.Address, error) {
	return _Inscriptionlightclient.Contract.GetRelayers(&_Inscriptionlightclient.CallOpts)
}

// GetRelayers is a free data retrieval call binding the contract method 0x179ff4b2.
//
// Solidity: function getRelayers() view returns(address[])
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) GetRelayers() ([]common.Address, error) {
	return _Inscriptionlightclient.Contract.GetRelayers(&_Inscriptionlightclient.CallOpts)
}

// InsHeight is a free data retrieval call binding the contract method 0xc7234309.
//
// Solidity: function insHeight() view returns(uint64)
func (_Inscriptionlightclient *InscriptionlightclientCaller) InsHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "insHeight")
	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err
}

// InsHeight is a free data retrieval call binding the contract method 0xc7234309.
//
// Solidity: function insHeight() view returns(uint64)
func (_Inscriptionlightclient *InscriptionlightclientSession) InsHeight() (uint64, error) {
	return _Inscriptionlightclient.Contract.InsHeight(&_Inscriptionlightclient.CallOpts)
}

// InsHeight is a free data retrieval call binding the contract method 0xc7234309.
//
// Solidity: function insHeight() view returns(uint64)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) InsHeight() (uint64, error) {
	return _Inscriptionlightclient.Contract.InsHeight(&_Inscriptionlightclient.CallOpts)
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientCaller) Relayers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "relayers", arg0)
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientSession) Relayers(arg0 *big.Int) (common.Address, error) {
	return _Inscriptionlightclient.Contract.Relayers(&_Inscriptionlightclient.CallOpts, arg0)
}

// Relayers is a free data retrieval call binding the contract method 0x9a48e7f9.
//
// Solidity: function relayers(uint256 ) view returns(address)
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) Relayers(arg0 *big.Int) (common.Address, error) {
	return _Inscriptionlightclient.Contract.Relayers(&_Inscriptionlightclient.CallOpts, arg0)
}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Inscriptionlightclient *InscriptionlightclientCaller) VerifyPackage(opts *bind.CallOpts, _pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	var out []interface{}
	err := _Inscriptionlightclient.contract.Call(opts, &out, "verifyPackage", _pkgKey, _payload, _blsSignature, _validatorSetBitMap)
	if err != nil {
		return err
	}

	return err
}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Inscriptionlightclient *InscriptionlightclientSession) VerifyPackage(_pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	return _Inscriptionlightclient.Contract.VerifyPackage(&_Inscriptionlightclient.CallOpts, _pkgKey, _payload, _blsSignature, _validatorSetBitMap)
}

// VerifyPackage is a free data retrieval call binding the contract method 0xfcdbebc5.
//
// Solidity: function verifyPackage(bytes _pkgKey, bytes _payload, bytes _blsSignature, uint256 _validatorSetBitMap) view returns()
func (_Inscriptionlightclient *InscriptionlightclientCallerSession) VerifyPackage(_pkgKey []byte, _payload []byte, _blsSignature []byte, _validatorSetBitMap *big.Int) error {
	return _Inscriptionlightclient.Contract.VerifyPackage(&_Inscriptionlightclient.CallOpts, _pkgKey, _payload, _blsSignature, _validatorSetBitMap)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Inscriptionlightclient *InscriptionlightclientTransactor) SyncTendermintHeader(opts *bind.TransactOpts, _header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Inscriptionlightclient.contract.Transact(opts, "syncTendermintHeader", _header, _height, _blsPubKeys, _relayers)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Inscriptionlightclient *InscriptionlightclientSession) SyncTendermintHeader(_header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Inscriptionlightclient.Contract.SyncTendermintHeader(&_Inscriptionlightclient.TransactOpts, _header, _height, _blsPubKeys, _relayers)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0x09f1b923.
//
// Solidity: function syncTendermintHeader(bytes _header, uint64 _height, bytes _blsPubKeys, address[] _relayers) returns()
func (_Inscriptionlightclient *InscriptionlightclientTransactorSession) SyncTendermintHeader(_header []byte, _height uint64, _blsPubKeys []byte, _relayers []common.Address) (*types.Transaction, error) {
	return _Inscriptionlightclient.Contract.SyncTendermintHeader(&_Inscriptionlightclient.TransactOpts, _header, _height, _blsPubKeys, _relayers)
}
