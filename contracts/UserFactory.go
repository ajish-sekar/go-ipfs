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

// UserFactoryABI is the input ABI used to generate the binding from.
const UserFactoryABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_userName\",\"type\":\"string\"}],\"name\":\"createUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddrs\",\"type\":\"address\"}],\"name\":\"getUserContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UserFactory is an auto generated Go binding around an Ethereum contract.
type UserFactory struct {
	UserFactoryCaller     // Read-only binding to the contract
	UserFactoryTransactor // Write-only binding to the contract
	UserFactoryFilterer   // Log filterer for contract events
}

// UserFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserFactorySession struct {
	Contract     *UserFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserFactoryCallerSession struct {
	Contract *UserFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UserFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserFactoryTransactorSession struct {
	Contract     *UserFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UserFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserFactoryRaw struct {
	Contract *UserFactory // Generic contract binding to access the raw methods on
}

// UserFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserFactoryCallerRaw struct {
	Contract *UserFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UserFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserFactoryTransactorRaw struct {
	Contract *UserFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserFactory creates a new instance of UserFactory, bound to a specific deployed contract.
func NewUserFactory(address common.Address, backend bind.ContractBackend) (*UserFactory, error) {
	contract, err := bindUserFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserFactory{UserFactoryCaller: UserFactoryCaller{contract: contract}, UserFactoryTransactor: UserFactoryTransactor{contract: contract}, UserFactoryFilterer: UserFactoryFilterer{contract: contract}}, nil
}

// NewUserFactoryCaller creates a new read-only instance of UserFactory, bound to a specific deployed contract.
func NewUserFactoryCaller(address common.Address, caller bind.ContractCaller) (*UserFactoryCaller, error) {
	contract, err := bindUserFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserFactoryCaller{contract: contract}, nil
}

// NewUserFactoryTransactor creates a new write-only instance of UserFactory, bound to a specific deployed contract.
func NewUserFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UserFactoryTransactor, error) {
	contract, err := bindUserFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserFactoryTransactor{contract: contract}, nil
}

// NewUserFactoryFilterer creates a new log filterer instance of UserFactory, bound to a specific deployed contract.
func NewUserFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFactoryFilterer, error) {
	contract, err := bindUserFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFactoryFilterer{contract: contract}, nil
}

// bindUserFactory binds a generic wrapper to an already deployed contract.
func bindUserFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserFactory *UserFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserFactory.Contract.UserFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserFactory *UserFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserFactory.Contract.UserFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserFactory *UserFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserFactory.Contract.UserFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserFactory *UserFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserFactory *UserFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserFactory *UserFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserFactory.Contract.contract.Transact(opts, method, params...)
}

// GetUserContract is a free data retrieval call binding the contract method 0x7e6c75f2.
//
// Solidity: function getUserContract(address userAddrs) constant returns(address)
func (_UserFactory *UserFactoryCaller) GetUserContract(opts *bind.CallOpts, userAddrs common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserFactory.contract.Call(opts, out, "getUserContract", userAddrs)
	return *ret0, err
}

// GetUserContract is a free data retrieval call binding the contract method 0x7e6c75f2.
//
// Solidity: function getUserContract(address userAddrs) constant returns(address)
func (_UserFactory *UserFactorySession) GetUserContract(userAddrs common.Address) (common.Address, error) {
	return _UserFactory.Contract.GetUserContract(&_UserFactory.CallOpts, userAddrs)
}

// GetUserContract is a free data retrieval call binding the contract method 0x7e6c75f2.
//
// Solidity: function getUserContract(address userAddrs) constant returns(address)
func (_UserFactory *UserFactoryCallerSession) GetUserContract(userAddrs common.Address) (common.Address, error) {
	return _UserFactory.Contract.GetUserContract(&_UserFactory.CallOpts, userAddrs)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(address)
func (_UserFactory *UserFactoryCaller) Users(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserFactory.contract.Call(opts, out, "users", arg0)
	return *ret0, err
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(address)
func (_UserFactory *UserFactorySession) Users(arg0 common.Address) (common.Address, error) {
	return _UserFactory.Contract.Users(&_UserFactory.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) constant returns(address)
func (_UserFactory *UserFactoryCallerSession) Users(arg0 common.Address) (common.Address, error) {
	return _UserFactory.Contract.Users(&_UserFactory.CallOpts, arg0)
}

// CreateUser is a paid mutator transaction binding the contract method 0x507ffba5.
//
// Solidity: function createUser(string _userName) returns(address)
func (_UserFactory *UserFactoryTransactor) CreateUser(opts *bind.TransactOpts, _userName string) (*types.Transaction, error) {
	return _UserFactory.contract.Transact(opts, "createUser", _userName)
}

// CreateUser is a paid mutator transaction binding the contract method 0x507ffba5.
//
// Solidity: function createUser(string _userName) returns(address)
func (_UserFactory *UserFactorySession) CreateUser(_userName string) (*types.Transaction, error) {
	return _UserFactory.Contract.CreateUser(&_UserFactory.TransactOpts, _userName)
}

// CreateUser is a paid mutator transaction binding the contract method 0x507ffba5.
//
// Solidity: function createUser(string _userName) returns(address)
func (_UserFactory *UserFactoryTransactorSession) CreateUser(_userName string) (*types.Transaction, error) {
	return _UserFactory.Contract.CreateUser(&_UserFactory.TransactOpts, _userName)
}
