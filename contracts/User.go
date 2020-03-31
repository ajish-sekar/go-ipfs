// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UserABI is the input ABI used to generate the binding from.
const UserABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_userName\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileAddress\",\"type\":\"address\"}],\"name\":\"addAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileAddress\",\"type\":\"address\"}],\"name\":\"addNewFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccessibleFiles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnedFiles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fileAddress\",\"type\":\"address\"}],\"name\":\"removeAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// GetAccessibleFiles is a free data retrieval call binding the contract method 0x635fd580.
//
// Solidity: function getAccessibleFiles() constant returns(address[])
func (_User *UserCaller) GetAccessibleFiles(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "getAccessibleFiles")
	return *ret0, err
}

// GetAccessibleFiles is a free data retrieval call binding the contract method 0x635fd580.
//
// Solidity: function getAccessibleFiles() constant returns(address[])
func (_User *UserSession) GetAccessibleFiles() ([]common.Address, error) {
	return _User.Contract.GetAccessibleFiles(&_User.CallOpts)
}

// GetAccessibleFiles is a free data retrieval call binding the contract method 0x635fd580.
//
// Solidity: function getAccessibleFiles() constant returns(address[])
func (_User *UserCallerSession) GetAccessibleFiles() ([]common.Address, error) {
	return _User.Contract.GetAccessibleFiles(&_User.CallOpts)
}

// GetOwnedFiles is a free data retrieval call binding the contract method 0x97a0fd3e.
//
// Solidity: function getOwnedFiles() constant returns(address[])
func (_User *UserCaller) GetOwnedFiles(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _User.contract.Call(opts, out, "getOwnedFiles")
	return *ret0, err
}

// GetOwnedFiles is a free data retrieval call binding the contract method 0x97a0fd3e.
//
// Solidity: function getOwnedFiles() constant returns(address[])
func (_User *UserSession) GetOwnedFiles() ([]common.Address, error) {
	return _User.Contract.GetOwnedFiles(&_User.CallOpts)
}

// GetOwnedFiles is a free data retrieval call binding the contract method 0x97a0fd3e.
//
// Solidity: function getOwnedFiles() constant returns(address[])
func (_User *UserCallerSession) GetOwnedFiles() ([]common.Address, error) {
	return _User.Contract.GetOwnedFiles(&_User.CallOpts)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address fileAddress) returns(bool)
func (_User *UserTransactor) AddAccess(opts *bind.TransactOpts, fileAddress common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addAccess", fileAddress)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address fileAddress) returns(bool)
func (_User *UserSession) AddAccess(fileAddress common.Address) (*types.Transaction, error) {
	return _User.Contract.AddAccess(&_User.TransactOpts, fileAddress)
}

// AddAccess is a paid mutator transaction binding the contract method 0xa118f249.
//
// Solidity: function addAccess(address fileAddress) returns(bool)
func (_User *UserTransactorSession) AddAccess(fileAddress common.Address) (*types.Transaction, error) {
	return _User.Contract.AddAccess(&_User.TransactOpts, fileAddress)
}

// AddNewFile is a paid mutator transaction binding the contract method 0xdd47dc89.
//
// Solidity: function addNewFile(address fileAddress) returns(bool)
func (_User *UserTransactor) AddNewFile(opts *bind.TransactOpts, fileAddress common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addNewFile", fileAddress)
}

// AddNewFile is a paid mutator transaction binding the contract method 0xdd47dc89.
//
// Solidity: function addNewFile(address fileAddress) returns(bool)
func (_User *UserSession) AddNewFile(fileAddress common.Address) (*types.Transaction, error) {
	return _User.Contract.AddNewFile(&_User.TransactOpts, fileAddress)
}

// AddNewFile is a paid mutator transaction binding the contract method 0xdd47dc89.
//
// Solidity: function addNewFile(address fileAddress) returns(bool)
func (_User *UserTransactorSession) AddNewFile(fileAddress common.Address) (*types.Transaction, error) {
	return _User.Contract.AddNewFile(&_User.TransactOpts, fileAddress)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address fileAddress) returns(bool)
func (_User *UserTransactor) RemoveAccess(opts *bind.TransactOpts, fileAddress common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "removeAccess", fileAddress)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address fileAddress) returns(bool)
func (_User *UserSession) RemoveAccess(fileAddress common.Address) (*types.Transaction, error) {
	return _User.Contract.RemoveAccess(&_User.TransactOpts, fileAddress)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x8823da6c.
//
// Solidity: function removeAccess(address fileAddress) returns(bool)
func (_User *UserTransactorSession) RemoveAccess(fileAddress common.Address) (*types.Transaction, error) {
	return _User.Contract.RemoveAccess(&_User.TransactOpts, fileAddress)
}
