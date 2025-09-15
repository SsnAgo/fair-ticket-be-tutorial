// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SimpleMockMetaData contains all meta data concerning the SimpleMock contract.
var SimpleMockMetaData = &bind.MetaData{
	ABI: "[]",
}

// SimpleMockABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleMockMetaData.ABI instead.
var SimpleMockABI = SimpleMockMetaData.ABI

// SimpleMock is an auto generated Go binding around an Ethereum contract.
type SimpleMock struct {
	SimpleMockCaller     // Read-only binding to the contract
	SimpleMockTransactor // Write-only binding to the contract
	SimpleMockFilterer   // Log filterer for contract events
}

// SimpleMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleMockSession struct {
	Contract     *SimpleMock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleMockCallerSession struct {
	Contract *SimpleMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimpleMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleMockTransactorSession struct {
	Contract     *SimpleMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleMockRaw struct {
	Contract *SimpleMock // Generic contract binding to access the raw methods on
}

// SimpleMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleMockCallerRaw struct {
	Contract *SimpleMockCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleMockTransactorRaw struct {
	Contract *SimpleMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleMock creates a new instance of SimpleMock, bound to a specific deployed contract.
func NewSimpleMock(address common.Address, backend bind.ContractBackend) (*SimpleMock, error) {
	contract, err := bindSimpleMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleMock{SimpleMockCaller: SimpleMockCaller{contract: contract}, SimpleMockTransactor: SimpleMockTransactor{contract: contract}, SimpleMockFilterer: SimpleMockFilterer{contract: contract}}, nil
}

// NewSimpleMockCaller creates a new read-only instance of SimpleMock, bound to a specific deployed contract.
func NewSimpleMockCaller(address common.Address, caller bind.ContractCaller) (*SimpleMockCaller, error) {
	contract, err := bindSimpleMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMockCaller{contract: contract}, nil
}

// NewSimpleMockTransactor creates a new write-only instance of SimpleMock, bound to a specific deployed contract.
func NewSimpleMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleMockTransactor, error) {
	contract, err := bindSimpleMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleMockTransactor{contract: contract}, nil
}

// NewSimpleMockFilterer creates a new log filterer instance of SimpleMock, bound to a specific deployed contract.
func NewSimpleMockFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleMockFilterer, error) {
	contract, err := bindSimpleMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleMockFilterer{contract: contract}, nil
}

// bindSimpleMock binds a generic wrapper to an already deployed contract.
func bindSimpleMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleMock *SimpleMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleMock.Contract.SimpleMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleMock *SimpleMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMock.Contract.SimpleMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleMock *SimpleMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMock.Contract.SimpleMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleMock *SimpleMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleMock *SimpleMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleMock *SimpleMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleMock.Contract.contract.Transact(opts, method, params...)
}
