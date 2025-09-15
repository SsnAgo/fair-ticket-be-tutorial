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

// LotteryResult is an auto generated low-level Go binding around an user-defined struct.
type LotteryResult struct {
	ProjectId   *big.Int
	MagicNumber *big.Int
}

// Participant is an auto generated low-level Go binding around an user-defined struct.
type Participant struct {
	Addr     common.Address
	LuckyNum *big.Int
}

// Project is an auto generated low-level Go binding around an user-defined struct.
type Project struct {
	Id            *big.Int
	Fingerprint   [32]byte
	Owner         common.Address
	TotalSupply   *big.Int
	ProjectStatus uint8
	MerkleRoot    [32]byte
}

// FairTicketMetaData contains all meta data concerning the FairTicket contract.
var FairTicketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_globalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SetMerkleRoot\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createProject\",\"inputs\":[{\"name\":\"_fingerprint\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finishProject\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLotteryResult\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structLotteryResult\",\"components\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"magicNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParticipantInfo\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structParticipant\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"luckyNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProjectInfo\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structProject\",\"components\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fingerprint\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectStatus\",\"type\":\"uint8\",\"internalType\":\"enumProjectStatus\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProjectParticipants\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structParticipant[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"luckyNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProjectParticipantsAmount\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProjectStatus\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumProjectStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lottery\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"participate\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_luckyNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_globalId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_pid2project\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fingerprint\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"projectStatus\",\"type\":\"uint8\",\"internalType\":\"enumProjectStatus\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_projectid_lottery\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"magicNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_projectid_paddr_participant\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"luckyNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_projectid_participants\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"luckyNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startProject\",\"inputs\":[{\"name\":\"_projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMerkleProof\",\"inputs\":[{\"name\":\"_prjectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MagicNumberPublished\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"magicNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProjectCreated\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"fingerprint\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProjectFinished\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProjectStarted\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MerkleProofInvalid\",\"inputs\":[{\"name\":\"projectId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"self\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]},{\"type\":\"error\",\"name\":\"MerkleRootAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OffsetOutOfBounds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyProjectOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ProjectAlreadyStarted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProjectNotFinished\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProjectNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProjectNotInProgress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSupplyZero\",\"inputs\":[]}]",
}

// FairTicketABI is the input ABI used to generate the binding from.
// Deprecated: Use FairTicketMetaData.ABI instead.
var FairTicketABI = FairTicketMetaData.ABI

// FairTicket is an auto generated Go binding around an Ethereum contract.
type FairTicket struct {
	FairTicketCaller     // Read-only binding to the contract
	FairTicketTransactor // Write-only binding to the contract
	FairTicketFilterer   // Log filterer for contract events
}

// FairTicketCaller is an auto generated read-only Go binding around an Ethereum contract.
type FairTicketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairTicketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FairTicketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairTicketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FairTicketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairTicketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FairTicketSession struct {
	Contract     *FairTicket       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FairTicketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FairTicketCallerSession struct {
	Contract *FairTicketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FairTicketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FairTicketTransactorSession struct {
	Contract     *FairTicketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FairTicketRaw is an auto generated low-level Go binding around an Ethereum contract.
type FairTicketRaw struct {
	Contract *FairTicket // Generic contract binding to access the raw methods on
}

// FairTicketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FairTicketCallerRaw struct {
	Contract *FairTicketCaller // Generic read-only contract binding to access the raw methods on
}

// FairTicketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FairTicketTransactorRaw struct {
	Contract *FairTicketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFairTicket creates a new instance of FairTicket, bound to a specific deployed contract.
func NewFairTicket(address common.Address, backend bind.ContractBackend) (*FairTicket, error) {
	contract, err := bindFairTicket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FairTicket{FairTicketCaller: FairTicketCaller{contract: contract}, FairTicketTransactor: FairTicketTransactor{contract: contract}, FairTicketFilterer: FairTicketFilterer{contract: contract}}, nil
}

// NewFairTicketCaller creates a new read-only instance of FairTicket, bound to a specific deployed contract.
func NewFairTicketCaller(address common.Address, caller bind.ContractCaller) (*FairTicketCaller, error) {
	contract, err := bindFairTicket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FairTicketCaller{contract: contract}, nil
}

// NewFairTicketTransactor creates a new write-only instance of FairTicket, bound to a specific deployed contract.
func NewFairTicketTransactor(address common.Address, transactor bind.ContractTransactor) (*FairTicketTransactor, error) {
	contract, err := bindFairTicket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FairTicketTransactor{contract: contract}, nil
}

// NewFairTicketFilterer creates a new log filterer instance of FairTicket, bound to a specific deployed contract.
func NewFairTicketFilterer(address common.Address, filterer bind.ContractFilterer) (*FairTicketFilterer, error) {
	contract, err := bindFairTicket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FairTicketFilterer{contract: contract}, nil
}

// bindFairTicket binds a generic wrapper to an already deployed contract.
func bindFairTicket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FairTicketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FairTicket *FairTicketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FairTicket.Contract.FairTicketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FairTicket *FairTicketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairTicket.Contract.FairTicketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FairTicket *FairTicketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FairTicket.Contract.FairTicketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FairTicket *FairTicketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FairTicket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FairTicket *FairTicketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairTicket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FairTicket *FairTicketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FairTicket.Contract.contract.Transact(opts, method, params...)
}

// GetLotteryResult is a free data retrieval call binding the contract method 0x8772d2dd.
//
// Solidity: function getLotteryResult(uint256 _projectId) view returns((uint256,uint256))
func (_FairTicket *FairTicketCaller) GetLotteryResult(opts *bind.CallOpts, _projectId *big.Int) (LotteryResult, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "getLotteryResult", _projectId)

	if err != nil {
		return *new(LotteryResult), err
	}

	out0 := *abi.ConvertType(out[0], new(LotteryResult)).(*LotteryResult)

	return out0, err

}

// GetLotteryResult is a free data retrieval call binding the contract method 0x8772d2dd.
//
// Solidity: function getLotteryResult(uint256 _projectId) view returns((uint256,uint256))
func (_FairTicket *FairTicketSession) GetLotteryResult(_projectId *big.Int) (LotteryResult, error) {
	return _FairTicket.Contract.GetLotteryResult(&_FairTicket.CallOpts, _projectId)
}

// GetLotteryResult is a free data retrieval call binding the contract method 0x8772d2dd.
//
// Solidity: function getLotteryResult(uint256 _projectId) view returns((uint256,uint256))
func (_FairTicket *FairTicketCallerSession) GetLotteryResult(_projectId *big.Int) (LotteryResult, error) {
	return _FairTicket.Contract.GetLotteryResult(&_FairTicket.CallOpts, _projectId)
}

// GetParticipantInfo is a free data retrieval call binding the contract method 0x99abff4b.
//
// Solidity: function getParticipantInfo(uint256 _projectId, address _addr) view returns((address,uint256))
func (_FairTicket *FairTicketCaller) GetParticipantInfo(opts *bind.CallOpts, _projectId *big.Int, _addr common.Address) (Participant, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "getParticipantInfo", _projectId, _addr)

	if err != nil {
		return *new(Participant), err
	}

	out0 := *abi.ConvertType(out[0], new(Participant)).(*Participant)

	return out0, err

}

// GetParticipantInfo is a free data retrieval call binding the contract method 0x99abff4b.
//
// Solidity: function getParticipantInfo(uint256 _projectId, address _addr) view returns((address,uint256))
func (_FairTicket *FairTicketSession) GetParticipantInfo(_projectId *big.Int, _addr common.Address) (Participant, error) {
	return _FairTicket.Contract.GetParticipantInfo(&_FairTicket.CallOpts, _projectId, _addr)
}

// GetParticipantInfo is a free data retrieval call binding the contract method 0x99abff4b.
//
// Solidity: function getParticipantInfo(uint256 _projectId, address _addr) view returns((address,uint256))
func (_FairTicket *FairTicketCallerSession) GetParticipantInfo(_projectId *big.Int, _addr common.Address) (Participant, error) {
	return _FairTicket.Contract.GetParticipantInfo(&_FairTicket.CallOpts, _projectId, _addr)
}

// GetProjectInfo is a free data retrieval call binding the contract method 0xfabf5968.
//
// Solidity: function getProjectInfo(uint256 _projectId) view returns((uint256,bytes32,address,uint256,uint8,bytes32))
func (_FairTicket *FairTicketCaller) GetProjectInfo(opts *bind.CallOpts, _projectId *big.Int) (Project, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "getProjectInfo", _projectId)

	if err != nil {
		return *new(Project), err
	}

	out0 := *abi.ConvertType(out[0], new(Project)).(*Project)

	return out0, err

}

// GetProjectInfo is a free data retrieval call binding the contract method 0xfabf5968.
//
// Solidity: function getProjectInfo(uint256 _projectId) view returns((uint256,bytes32,address,uint256,uint8,bytes32))
func (_FairTicket *FairTicketSession) GetProjectInfo(_projectId *big.Int) (Project, error) {
	return _FairTicket.Contract.GetProjectInfo(&_FairTicket.CallOpts, _projectId)
}

// GetProjectInfo is a free data retrieval call binding the contract method 0xfabf5968.
//
// Solidity: function getProjectInfo(uint256 _projectId) view returns((uint256,bytes32,address,uint256,uint8,bytes32))
func (_FairTicket *FairTicketCallerSession) GetProjectInfo(_projectId *big.Int) (Project, error) {
	return _FairTicket.Contract.GetProjectInfo(&_FairTicket.CallOpts, _projectId)
}

// GetProjectParticipants is a free data retrieval call binding the contract method 0x862e9f0f.
//
// Solidity: function getProjectParticipants(uint256 _projectId, uint256 _offset, uint256 _limit) view returns((address,uint256)[])
func (_FairTicket *FairTicketCaller) GetProjectParticipants(opts *bind.CallOpts, _projectId *big.Int, _offset *big.Int, _limit *big.Int) ([]Participant, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "getProjectParticipants", _projectId, _offset, _limit)

	if err != nil {
		return *new([]Participant), err
	}

	out0 := *abi.ConvertType(out[0], new([]Participant)).(*[]Participant)

	return out0, err

}

// GetProjectParticipants is a free data retrieval call binding the contract method 0x862e9f0f.
//
// Solidity: function getProjectParticipants(uint256 _projectId, uint256 _offset, uint256 _limit) view returns((address,uint256)[])
func (_FairTicket *FairTicketSession) GetProjectParticipants(_projectId *big.Int, _offset *big.Int, _limit *big.Int) ([]Participant, error) {
	return _FairTicket.Contract.GetProjectParticipants(&_FairTicket.CallOpts, _projectId, _offset, _limit)
}

// GetProjectParticipants is a free data retrieval call binding the contract method 0x862e9f0f.
//
// Solidity: function getProjectParticipants(uint256 _projectId, uint256 _offset, uint256 _limit) view returns((address,uint256)[])
func (_FairTicket *FairTicketCallerSession) GetProjectParticipants(_projectId *big.Int, _offset *big.Int, _limit *big.Int) ([]Participant, error) {
	return _FairTicket.Contract.GetProjectParticipants(&_FairTicket.CallOpts, _projectId, _offset, _limit)
}

// GetProjectParticipantsAmount is a free data retrieval call binding the contract method 0xe10f37c0.
//
// Solidity: function getProjectParticipantsAmount(uint256 _projectId) view returns(uint256)
func (_FairTicket *FairTicketCaller) GetProjectParticipantsAmount(opts *bind.CallOpts, _projectId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "getProjectParticipantsAmount", _projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProjectParticipantsAmount is a free data retrieval call binding the contract method 0xe10f37c0.
//
// Solidity: function getProjectParticipantsAmount(uint256 _projectId) view returns(uint256)
func (_FairTicket *FairTicketSession) GetProjectParticipantsAmount(_projectId *big.Int) (*big.Int, error) {
	return _FairTicket.Contract.GetProjectParticipantsAmount(&_FairTicket.CallOpts, _projectId)
}

// GetProjectParticipantsAmount is a free data retrieval call binding the contract method 0xe10f37c0.
//
// Solidity: function getProjectParticipantsAmount(uint256 _projectId) view returns(uint256)
func (_FairTicket *FairTicketCallerSession) GetProjectParticipantsAmount(_projectId *big.Int) (*big.Int, error) {
	return _FairTicket.Contract.GetProjectParticipantsAmount(&_FairTicket.CallOpts, _projectId)
}

// GetProjectStatus is a free data retrieval call binding the contract method 0x9675c009.
//
// Solidity: function getProjectStatus(uint256 _projectId) view returns(uint8)
func (_FairTicket *FairTicketCaller) GetProjectStatus(opts *bind.CallOpts, _projectId *big.Int) (uint8, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "getProjectStatus", _projectId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProjectStatus is a free data retrieval call binding the contract method 0x9675c009.
//
// Solidity: function getProjectStatus(uint256 _projectId) view returns(uint8)
func (_FairTicket *FairTicketSession) GetProjectStatus(_projectId *big.Int) (uint8, error) {
	return _FairTicket.Contract.GetProjectStatus(&_FairTicket.CallOpts, _projectId)
}

// GetProjectStatus is a free data retrieval call binding the contract method 0x9675c009.
//
// Solidity: function getProjectStatus(uint256 _projectId) view returns(uint8)
func (_FairTicket *FairTicketCallerSession) GetProjectStatus(_projectId *big.Int) (uint8, error) {
	return _FairTicket.Contract.GetProjectStatus(&_FairTicket.CallOpts, _projectId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairTicket *FairTicketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairTicket *FairTicketSession) Owner() (common.Address, error) {
	return _FairTicket.Contract.Owner(&_FairTicket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FairTicket *FairTicketCallerSession) Owner() (common.Address, error) {
	return _FairTicket.Contract.Owner(&_FairTicket.CallOpts)
}

// SGlobalId is a free data retrieval call binding the contract method 0xac7db523.
//
// Solidity: function s_globalId() view returns(uint256)
func (_FairTicket *FairTicketCaller) SGlobalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "s_globalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SGlobalId is a free data retrieval call binding the contract method 0xac7db523.
//
// Solidity: function s_globalId() view returns(uint256)
func (_FairTicket *FairTicketSession) SGlobalId() (*big.Int, error) {
	return _FairTicket.Contract.SGlobalId(&_FairTicket.CallOpts)
}

// SGlobalId is a free data retrieval call binding the contract method 0xac7db523.
//
// Solidity: function s_globalId() view returns(uint256)
func (_FairTicket *FairTicketCallerSession) SGlobalId() (*big.Int, error) {
	return _FairTicket.Contract.SGlobalId(&_FairTicket.CallOpts)
}

// SPid2project is a free data retrieval call binding the contract method 0xc1c5cb2b.
//
// Solidity: function s_pid2project(uint256 ) view returns(uint256 id, bytes32 fingerprint, address owner, uint256 totalSupply, uint8 projectStatus, bytes32 merkleRoot)
func (_FairTicket *FairTicketCaller) SPid2project(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id            *big.Int
	Fingerprint   [32]byte
	Owner         common.Address
	TotalSupply   *big.Int
	ProjectStatus uint8
	MerkleRoot    [32]byte
}, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "s_pid2project", arg0)

	outstruct := new(struct {
		Id            *big.Int
		Fingerprint   [32]byte
		Owner         common.Address
		TotalSupply   *big.Int
		ProjectStatus uint8
		MerkleRoot    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fingerprint = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TotalSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ProjectStatus = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.MerkleRoot = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SPid2project is a free data retrieval call binding the contract method 0xc1c5cb2b.
//
// Solidity: function s_pid2project(uint256 ) view returns(uint256 id, bytes32 fingerprint, address owner, uint256 totalSupply, uint8 projectStatus, bytes32 merkleRoot)
func (_FairTicket *FairTicketSession) SPid2project(arg0 *big.Int) (struct {
	Id            *big.Int
	Fingerprint   [32]byte
	Owner         common.Address
	TotalSupply   *big.Int
	ProjectStatus uint8
	MerkleRoot    [32]byte
}, error) {
	return _FairTicket.Contract.SPid2project(&_FairTicket.CallOpts, arg0)
}

// SPid2project is a free data retrieval call binding the contract method 0xc1c5cb2b.
//
// Solidity: function s_pid2project(uint256 ) view returns(uint256 id, bytes32 fingerprint, address owner, uint256 totalSupply, uint8 projectStatus, bytes32 merkleRoot)
func (_FairTicket *FairTicketCallerSession) SPid2project(arg0 *big.Int) (struct {
	Id            *big.Int
	Fingerprint   [32]byte
	Owner         common.Address
	TotalSupply   *big.Int
	ProjectStatus uint8
	MerkleRoot    [32]byte
}, error) {
	return _FairTicket.Contract.SPid2project(&_FairTicket.CallOpts, arg0)
}

// SProjectidLottery is a free data retrieval call binding the contract method 0x8105ddb7.
//
// Solidity: function s_projectid_lottery(uint256 ) view returns(uint256 projectId, uint256 magicNumber)
func (_FairTicket *FairTicketCaller) SProjectidLottery(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProjectId   *big.Int
	MagicNumber *big.Int
}, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "s_projectid_lottery", arg0)

	outstruct := new(struct {
		ProjectId   *big.Int
		MagicNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProjectId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MagicNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SProjectidLottery is a free data retrieval call binding the contract method 0x8105ddb7.
//
// Solidity: function s_projectid_lottery(uint256 ) view returns(uint256 projectId, uint256 magicNumber)
func (_FairTicket *FairTicketSession) SProjectidLottery(arg0 *big.Int) (struct {
	ProjectId   *big.Int
	MagicNumber *big.Int
}, error) {
	return _FairTicket.Contract.SProjectidLottery(&_FairTicket.CallOpts, arg0)
}

// SProjectidLottery is a free data retrieval call binding the contract method 0x8105ddb7.
//
// Solidity: function s_projectid_lottery(uint256 ) view returns(uint256 projectId, uint256 magicNumber)
func (_FairTicket *FairTicketCallerSession) SProjectidLottery(arg0 *big.Int) (struct {
	ProjectId   *big.Int
	MagicNumber *big.Int
}, error) {
	return _FairTicket.Contract.SProjectidLottery(&_FairTicket.CallOpts, arg0)
}

// SProjectidPaddrParticipant is a free data retrieval call binding the contract method 0xc3cff05b.
//
// Solidity: function s_projectid_paddr_participant(uint256 , address ) view returns(address addr, uint256 luckyNum)
func (_FairTicket *FairTicketCaller) SProjectidPaddrParticipant(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Addr     common.Address
	LuckyNum *big.Int
}, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "s_projectid_paddr_participant", arg0, arg1)

	outstruct := new(struct {
		Addr     common.Address
		LuckyNum *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LuckyNum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SProjectidPaddrParticipant is a free data retrieval call binding the contract method 0xc3cff05b.
//
// Solidity: function s_projectid_paddr_participant(uint256 , address ) view returns(address addr, uint256 luckyNum)
func (_FairTicket *FairTicketSession) SProjectidPaddrParticipant(arg0 *big.Int, arg1 common.Address) (struct {
	Addr     common.Address
	LuckyNum *big.Int
}, error) {
	return _FairTicket.Contract.SProjectidPaddrParticipant(&_FairTicket.CallOpts, arg0, arg1)
}

// SProjectidPaddrParticipant is a free data retrieval call binding the contract method 0xc3cff05b.
//
// Solidity: function s_projectid_paddr_participant(uint256 , address ) view returns(address addr, uint256 luckyNum)
func (_FairTicket *FairTicketCallerSession) SProjectidPaddrParticipant(arg0 *big.Int, arg1 common.Address) (struct {
	Addr     common.Address
	LuckyNum *big.Int
}, error) {
	return _FairTicket.Contract.SProjectidPaddrParticipant(&_FairTicket.CallOpts, arg0, arg1)
}

// SProjectidParticipants is a free data retrieval call binding the contract method 0x5c402ecf.
//
// Solidity: function s_projectid_participants(uint256 , uint256 ) view returns(address addr, uint256 luckyNum)
func (_FairTicket *FairTicketCaller) SProjectidParticipants(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Addr     common.Address
	LuckyNum *big.Int
}, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "s_projectid_participants", arg0, arg1)

	outstruct := new(struct {
		Addr     common.Address
		LuckyNum *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LuckyNum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SProjectidParticipants is a free data retrieval call binding the contract method 0x5c402ecf.
//
// Solidity: function s_projectid_participants(uint256 , uint256 ) view returns(address addr, uint256 luckyNum)
func (_FairTicket *FairTicketSession) SProjectidParticipants(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr     common.Address
	LuckyNum *big.Int
}, error) {
	return _FairTicket.Contract.SProjectidParticipants(&_FairTicket.CallOpts, arg0, arg1)
}

// SProjectidParticipants is a free data retrieval call binding the contract method 0x5c402ecf.
//
// Solidity: function s_projectid_participants(uint256 , uint256 ) view returns(address addr, uint256 luckyNum)
func (_FairTicket *FairTicketCallerSession) SProjectidParticipants(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr     common.Address
	LuckyNum *big.Int
}, error) {
	return _FairTicket.Contract.SProjectidParticipants(&_FairTicket.CallOpts, arg0, arg1)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xd4adf169.
//
// Solidity: function verifyMerkleProof(uint256 _prjectId, bytes32[] proof) view returns(bool)
func (_FairTicket *FairTicketCaller) VerifyMerkleProof(opts *bind.CallOpts, _prjectId *big.Int, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _FairTicket.contract.Call(opts, &out, "verifyMerkleProof", _prjectId, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xd4adf169.
//
// Solidity: function verifyMerkleProof(uint256 _prjectId, bytes32[] proof) view returns(bool)
func (_FairTicket *FairTicketSession) VerifyMerkleProof(_prjectId *big.Int, proof [][32]byte) (bool, error) {
	return _FairTicket.Contract.VerifyMerkleProof(&_FairTicket.CallOpts, _prjectId, proof)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xd4adf169.
//
// Solidity: function verifyMerkleProof(uint256 _prjectId, bytes32[] proof) view returns(bool)
func (_FairTicket *FairTicketCallerSession) VerifyMerkleProof(_prjectId *big.Int, proof [][32]byte) (bool, error) {
	return _FairTicket.Contract.VerifyMerkleProof(&_FairTicket.CallOpts, _prjectId, proof)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x06300242.
//
// Solidity: function SetMerkleRoot(uint256 _projectId, bytes32 _merkleRoot) returns()
func (_FairTicket *FairTicketTransactor) SetMerkleRoot(opts *bind.TransactOpts, _projectId *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _FairTicket.contract.Transact(opts, "SetMerkleRoot", _projectId, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x06300242.
//
// Solidity: function SetMerkleRoot(uint256 _projectId, bytes32 _merkleRoot) returns()
func (_FairTicket *FairTicketSession) SetMerkleRoot(_projectId *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _FairTicket.Contract.SetMerkleRoot(&_FairTicket.TransactOpts, _projectId, _merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x06300242.
//
// Solidity: function SetMerkleRoot(uint256 _projectId, bytes32 _merkleRoot) returns()
func (_FairTicket *FairTicketTransactorSession) SetMerkleRoot(_projectId *big.Int, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _FairTicket.Contract.SetMerkleRoot(&_FairTicket.TransactOpts, _projectId, _merkleRoot)
}

// CreateProject is a paid mutator transaction binding the contract method 0xfcc01b28.
//
// Solidity: function createProject(bytes32 _fingerprint, address _owner, uint256 _totalSupply) returns()
func (_FairTicket *FairTicketTransactor) CreateProject(opts *bind.TransactOpts, _fingerprint [32]byte, _owner common.Address, _totalSupply *big.Int) (*types.Transaction, error) {
	return _FairTicket.contract.Transact(opts, "createProject", _fingerprint, _owner, _totalSupply)
}

// CreateProject is a paid mutator transaction binding the contract method 0xfcc01b28.
//
// Solidity: function createProject(bytes32 _fingerprint, address _owner, uint256 _totalSupply) returns()
func (_FairTicket *FairTicketSession) CreateProject(_fingerprint [32]byte, _owner common.Address, _totalSupply *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.CreateProject(&_FairTicket.TransactOpts, _fingerprint, _owner, _totalSupply)
}

// CreateProject is a paid mutator transaction binding the contract method 0xfcc01b28.
//
// Solidity: function createProject(bytes32 _fingerprint, address _owner, uint256 _totalSupply) returns()
func (_FairTicket *FairTicketTransactorSession) CreateProject(_fingerprint [32]byte, _owner common.Address, _totalSupply *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.CreateProject(&_FairTicket.TransactOpts, _fingerprint, _owner, _totalSupply)
}

// FinishProject is a paid mutator transaction binding the contract method 0xe4f23e61.
//
// Solidity: function finishProject(uint256 _projectId) returns()
func (_FairTicket *FairTicketTransactor) FinishProject(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.contract.Transact(opts, "finishProject", _projectId)
}

// FinishProject is a paid mutator transaction binding the contract method 0xe4f23e61.
//
// Solidity: function finishProject(uint256 _projectId) returns()
func (_FairTicket *FairTicketSession) FinishProject(_projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.FinishProject(&_FairTicket.TransactOpts, _projectId)
}

// FinishProject is a paid mutator transaction binding the contract method 0xe4f23e61.
//
// Solidity: function finishProject(uint256 _projectId) returns()
func (_FairTicket *FairTicketTransactorSession) FinishProject(_projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.FinishProject(&_FairTicket.TransactOpts, _projectId)
}

// Lottery is a paid mutator transaction binding the contract method 0xa57d1560.
//
// Solidity: function lottery(uint256 _projectId) returns()
func (_FairTicket *FairTicketTransactor) Lottery(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.contract.Transact(opts, "lottery", _projectId)
}

// Lottery is a paid mutator transaction binding the contract method 0xa57d1560.
//
// Solidity: function lottery(uint256 _projectId) returns()
func (_FairTicket *FairTicketSession) Lottery(_projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.Lottery(&_FairTicket.TransactOpts, _projectId)
}

// Lottery is a paid mutator transaction binding the contract method 0xa57d1560.
//
// Solidity: function lottery(uint256 _projectId) returns()
func (_FairTicket *FairTicketTransactorSession) Lottery(_projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.Lottery(&_FairTicket.TransactOpts, _projectId)
}

// Participate is a paid mutator transaction binding the contract method 0x6b7fd1a9.
//
// Solidity: function participate(uint256 _projectId, address _addr, uint256 _luckyNum) returns()
func (_FairTicket *FairTicketTransactor) Participate(opts *bind.TransactOpts, _projectId *big.Int, _addr common.Address, _luckyNum *big.Int) (*types.Transaction, error) {
	return _FairTicket.contract.Transact(opts, "participate", _projectId, _addr, _luckyNum)
}

// Participate is a paid mutator transaction binding the contract method 0x6b7fd1a9.
//
// Solidity: function participate(uint256 _projectId, address _addr, uint256 _luckyNum) returns()
func (_FairTicket *FairTicketSession) Participate(_projectId *big.Int, _addr common.Address, _luckyNum *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.Participate(&_FairTicket.TransactOpts, _projectId, _addr, _luckyNum)
}

// Participate is a paid mutator transaction binding the contract method 0x6b7fd1a9.
//
// Solidity: function participate(uint256 _projectId, address _addr, uint256 _luckyNum) returns()
func (_FairTicket *FairTicketTransactorSession) Participate(_projectId *big.Int, _addr common.Address, _luckyNum *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.Participate(&_FairTicket.TransactOpts, _projectId, _addr, _luckyNum)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairTicket *FairTicketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FairTicket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairTicket *FairTicketSession) RenounceOwnership() (*types.Transaction, error) {
	return _FairTicket.Contract.RenounceOwnership(&_FairTicket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FairTicket *FairTicketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FairTicket.Contract.RenounceOwnership(&_FairTicket.TransactOpts)
}

// StartProject is a paid mutator transaction binding the contract method 0xf62bf2de.
//
// Solidity: function startProject(uint256 _projectId) returns()
func (_FairTicket *FairTicketTransactor) StartProject(opts *bind.TransactOpts, _projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.contract.Transact(opts, "startProject", _projectId)
}

// StartProject is a paid mutator transaction binding the contract method 0xf62bf2de.
//
// Solidity: function startProject(uint256 _projectId) returns()
func (_FairTicket *FairTicketSession) StartProject(_projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.StartProject(&_FairTicket.TransactOpts, _projectId)
}

// StartProject is a paid mutator transaction binding the contract method 0xf62bf2de.
//
// Solidity: function startProject(uint256 _projectId) returns()
func (_FairTicket *FairTicketTransactorSession) StartProject(_projectId *big.Int) (*types.Transaction, error) {
	return _FairTicket.Contract.StartProject(&_FairTicket.TransactOpts, _projectId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairTicket *FairTicketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FairTicket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairTicket *FairTicketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FairTicket.Contract.TransferOwnership(&_FairTicket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FairTicket *FairTicketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FairTicket.Contract.TransferOwnership(&_FairTicket.TransactOpts, newOwner)
}

// FairTicketMagicNumberPublishedIterator is returned from FilterMagicNumberPublished and is used to iterate over the raw logs and unpacked data for MagicNumberPublished events raised by the FairTicket contract.
type FairTicketMagicNumberPublishedIterator struct {
	Event *FairTicketMagicNumberPublished // Event containing the contract specifics and raw log

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
func (it *FairTicketMagicNumberPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTicketMagicNumberPublished)
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
		it.Event = new(FairTicketMagicNumberPublished)
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
func (it *FairTicketMagicNumberPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTicketMagicNumberPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTicketMagicNumberPublished represents a MagicNumberPublished event raised by the FairTicket contract.
type FairTicketMagicNumberPublished struct {
	ProjectId   *big.Int
	MagicNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMagicNumberPublished is a free log retrieval operation binding the contract event 0x866549e6c822e157d2b1fbce0bb661c62d3925e7ee20a0b70d93b81ac6f6fc0c.
//
// Solidity: event MagicNumberPublished(uint256 indexed projectId, uint256 magicNumber)
func (_FairTicket *FairTicketFilterer) FilterMagicNumberPublished(opts *bind.FilterOpts, projectId []*big.Int) (*FairTicketMagicNumberPublishedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _FairTicket.contract.FilterLogs(opts, "MagicNumberPublished", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &FairTicketMagicNumberPublishedIterator{contract: _FairTicket.contract, event: "MagicNumberPublished", logs: logs, sub: sub}, nil
}

// WatchMagicNumberPublished is a free log subscription operation binding the contract event 0x866549e6c822e157d2b1fbce0bb661c62d3925e7ee20a0b70d93b81ac6f6fc0c.
//
// Solidity: event MagicNumberPublished(uint256 indexed projectId, uint256 magicNumber)
func (_FairTicket *FairTicketFilterer) WatchMagicNumberPublished(opts *bind.WatchOpts, sink chan<- *FairTicketMagicNumberPublished, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _FairTicket.contract.WatchLogs(opts, "MagicNumberPublished", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTicketMagicNumberPublished)
				if err := _FairTicket.contract.UnpackLog(event, "MagicNumberPublished", log); err != nil {
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

// ParseMagicNumberPublished is a log parse operation binding the contract event 0x866549e6c822e157d2b1fbce0bb661c62d3925e7ee20a0b70d93b81ac6f6fc0c.
//
// Solidity: event MagicNumberPublished(uint256 indexed projectId, uint256 magicNumber)
func (_FairTicket *FairTicketFilterer) ParseMagicNumberPublished(log types.Log) (*FairTicketMagicNumberPublished, error) {
	event := new(FairTicketMagicNumberPublished)
	if err := _FairTicket.contract.UnpackLog(event, "MagicNumberPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTicketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FairTicket contract.
type FairTicketOwnershipTransferredIterator struct {
	Event *FairTicketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FairTicketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTicketOwnershipTransferred)
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
		it.Event = new(FairTicketOwnershipTransferred)
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
func (it *FairTicketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTicketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTicketOwnershipTransferred represents a OwnershipTransferred event raised by the FairTicket contract.
type FairTicketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FairTicket *FairTicketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FairTicketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FairTicket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FairTicketOwnershipTransferredIterator{contract: _FairTicket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FairTicket *FairTicketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FairTicketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FairTicket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTicketOwnershipTransferred)
				if err := _FairTicket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FairTicket *FairTicketFilterer) ParseOwnershipTransferred(log types.Log) (*FairTicketOwnershipTransferred, error) {
	event := new(FairTicketOwnershipTransferred)
	if err := _FairTicket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTicketProjectCreatedIterator is returned from FilterProjectCreated and is used to iterate over the raw logs and unpacked data for ProjectCreated events raised by the FairTicket contract.
type FairTicketProjectCreatedIterator struct {
	Event *FairTicketProjectCreated // Event containing the contract specifics and raw log

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
func (it *FairTicketProjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTicketProjectCreated)
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
		it.Event = new(FairTicketProjectCreated)
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
func (it *FairTicketProjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTicketProjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTicketProjectCreated represents a ProjectCreated event raised by the FairTicket contract.
type FairTicketProjectCreated struct {
	ProjectId   *big.Int
	Fingerprint [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProjectCreated is a free log retrieval operation binding the contract event 0x5323249ea93983a70e1acd0f3f061840514c3c76e3420fc547a57ed69cd6fd1d.
//
// Solidity: event ProjectCreated(uint256 indexed projectId, bytes32 indexed fingerprint)
func (_FairTicket *FairTicketFilterer) FilterProjectCreated(opts *bind.FilterOpts, projectId []*big.Int, fingerprint [][32]byte) (*FairTicketProjectCreatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var fingerprintRule []interface{}
	for _, fingerprintItem := range fingerprint {
		fingerprintRule = append(fingerprintRule, fingerprintItem)
	}

	logs, sub, err := _FairTicket.contract.FilterLogs(opts, "ProjectCreated", projectIdRule, fingerprintRule)
	if err != nil {
		return nil, err
	}
	return &FairTicketProjectCreatedIterator{contract: _FairTicket.contract, event: "ProjectCreated", logs: logs, sub: sub}, nil
}

// WatchProjectCreated is a free log subscription operation binding the contract event 0x5323249ea93983a70e1acd0f3f061840514c3c76e3420fc547a57ed69cd6fd1d.
//
// Solidity: event ProjectCreated(uint256 indexed projectId, bytes32 indexed fingerprint)
func (_FairTicket *FairTicketFilterer) WatchProjectCreated(opts *bind.WatchOpts, sink chan<- *FairTicketProjectCreated, projectId []*big.Int, fingerprint [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var fingerprintRule []interface{}
	for _, fingerprintItem := range fingerprint {
		fingerprintRule = append(fingerprintRule, fingerprintItem)
	}

	logs, sub, err := _FairTicket.contract.WatchLogs(opts, "ProjectCreated", projectIdRule, fingerprintRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTicketProjectCreated)
				if err := _FairTicket.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
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

// ParseProjectCreated is a log parse operation binding the contract event 0x5323249ea93983a70e1acd0f3f061840514c3c76e3420fc547a57ed69cd6fd1d.
//
// Solidity: event ProjectCreated(uint256 indexed projectId, bytes32 indexed fingerprint)
func (_FairTicket *FairTicketFilterer) ParseProjectCreated(log types.Log) (*FairTicketProjectCreated, error) {
	event := new(FairTicketProjectCreated)
	if err := _FairTicket.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTicketProjectFinishedIterator is returned from FilterProjectFinished and is used to iterate over the raw logs and unpacked data for ProjectFinished events raised by the FairTicket contract.
type FairTicketProjectFinishedIterator struct {
	Event *FairTicketProjectFinished // Event containing the contract specifics and raw log

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
func (it *FairTicketProjectFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTicketProjectFinished)
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
		it.Event = new(FairTicketProjectFinished)
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
func (it *FairTicketProjectFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTicketProjectFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTicketProjectFinished represents a ProjectFinished event raised by the FairTicket contract.
type FairTicketProjectFinished struct {
	ProjectId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectFinished is a free log retrieval operation binding the contract event 0x5032973ebecb355d37badc06a2e7e34401c2b21eaade274aa03b3730fa18ff8a.
//
// Solidity: event ProjectFinished(uint256 indexed projectId)
func (_FairTicket *FairTicketFilterer) FilterProjectFinished(opts *bind.FilterOpts, projectId []*big.Int) (*FairTicketProjectFinishedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _FairTicket.contract.FilterLogs(opts, "ProjectFinished", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &FairTicketProjectFinishedIterator{contract: _FairTicket.contract, event: "ProjectFinished", logs: logs, sub: sub}, nil
}

// WatchProjectFinished is a free log subscription operation binding the contract event 0x5032973ebecb355d37badc06a2e7e34401c2b21eaade274aa03b3730fa18ff8a.
//
// Solidity: event ProjectFinished(uint256 indexed projectId)
func (_FairTicket *FairTicketFilterer) WatchProjectFinished(opts *bind.WatchOpts, sink chan<- *FairTicketProjectFinished, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _FairTicket.contract.WatchLogs(opts, "ProjectFinished", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTicketProjectFinished)
				if err := _FairTicket.contract.UnpackLog(event, "ProjectFinished", log); err != nil {
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

// ParseProjectFinished is a log parse operation binding the contract event 0x5032973ebecb355d37badc06a2e7e34401c2b21eaade274aa03b3730fa18ff8a.
//
// Solidity: event ProjectFinished(uint256 indexed projectId)
func (_FairTicket *FairTicketFilterer) ParseProjectFinished(log types.Log) (*FairTicketProjectFinished, error) {
	event := new(FairTicketProjectFinished)
	if err := _FairTicket.contract.UnpackLog(event, "ProjectFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FairTicketProjectStartedIterator is returned from FilterProjectStarted and is used to iterate over the raw logs and unpacked data for ProjectStarted events raised by the FairTicket contract.
type FairTicketProjectStartedIterator struct {
	Event *FairTicketProjectStarted // Event containing the contract specifics and raw log

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
func (it *FairTicketProjectStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FairTicketProjectStarted)
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
		it.Event = new(FairTicketProjectStarted)
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
func (it *FairTicketProjectStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FairTicketProjectStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FairTicketProjectStarted represents a ProjectStarted event raised by the FairTicket contract.
type FairTicketProjectStarted struct {
	ProjectId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectStarted is a free log retrieval operation binding the contract event 0xc6103cc0479d202218c64bbfa4df30c34a1f97e64c85a5a441103a0766241e66.
//
// Solidity: event ProjectStarted(uint256 indexed projectId)
func (_FairTicket *FairTicketFilterer) FilterProjectStarted(opts *bind.FilterOpts, projectId []*big.Int) (*FairTicketProjectStartedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _FairTicket.contract.FilterLogs(opts, "ProjectStarted", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &FairTicketProjectStartedIterator{contract: _FairTicket.contract, event: "ProjectStarted", logs: logs, sub: sub}, nil
}

// WatchProjectStarted is a free log subscription operation binding the contract event 0xc6103cc0479d202218c64bbfa4df30c34a1f97e64c85a5a441103a0766241e66.
//
// Solidity: event ProjectStarted(uint256 indexed projectId)
func (_FairTicket *FairTicketFilterer) WatchProjectStarted(opts *bind.WatchOpts, sink chan<- *FairTicketProjectStarted, projectId []*big.Int) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _FairTicket.contract.WatchLogs(opts, "ProjectStarted", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FairTicketProjectStarted)
				if err := _FairTicket.contract.UnpackLog(event, "ProjectStarted", log); err != nil {
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

// ParseProjectStarted is a log parse operation binding the contract event 0xc6103cc0479d202218c64bbfa4df30c34a1f97e64c85a5a441103a0766241e66.
//
// Solidity: event ProjectStarted(uint256 indexed projectId)
func (_FairTicket *FairTicketFilterer) ParseProjectStarted(log types.Log) (*FairTicketProjectStarted, error) {
	event := new(FairTicketProjectStarted)
	if err := _FairTicket.contract.UnpackLog(event, "ProjectStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
