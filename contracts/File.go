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

// FileABI is the input ABI used to generate the binding from.
const FileABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_iv\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_salt\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIv\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSalt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"hasAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// File is an auto generated Go binding around an Ethereum contract.
type File struct {
	FileCaller     // Read-only binding to the contract
	FileTransactor // Write-only binding to the contract
	FileFilterer   // Log filterer for contract events
}

// FileCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileSession struct {
	Contract     *File             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileCallerSession struct {
	Contract *FileCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileTransactorSession struct {
	Contract     *FileTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileRaw struct {
	Contract *File // Generic contract binding to access the raw methods on
}

// FileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileCallerRaw struct {
	Contract *FileCaller // Generic read-only contract binding to access the raw methods on
}

// FileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileTransactorRaw struct {
	Contract *FileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFile creates a new instance of File, bound to a specific deployed contract.
func NewFile(address common.Address, backend bind.ContractBackend) (*File, error) {
	contract, err := bindFile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &File{FileCaller: FileCaller{contract: contract}, FileTransactor: FileTransactor{contract: contract}, FileFilterer: FileFilterer{contract: contract}}, nil
}

// NewFileCaller creates a new read-only instance of File, bound to a specific deployed contract.
func NewFileCaller(address common.Address, caller bind.ContractCaller) (*FileCaller, error) {
	contract, err := bindFile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileCaller{contract: contract}, nil
}

// NewFileTransactor creates a new write-only instance of File, bound to a specific deployed contract.
func NewFileTransactor(address common.Address, transactor bind.ContractTransactor) (*FileTransactor, error) {
	contract, err := bindFile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileTransactor{contract: contract}, nil
}

// NewFileFilterer creates a new log filterer instance of File, bound to a specific deployed contract.
func NewFileFilterer(address common.Address, filterer bind.ContractFilterer) (*FileFilterer, error) {
	contract, err := bindFile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileFilterer{contract: contract}, nil
}

// bindFile binds a generic wrapper to an already deployed contract.
func bindFile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_File *FileRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _File.Contract.FileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_File *FileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _File.Contract.FileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_File *FileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _File.Contract.FileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_File *FileCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _File.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_File *FileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _File.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_File *FileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _File.Contract.contract.Transact(opts, method, params...)
}

// GetFileHash is a free data retrieval call binding the contract method 0x8493f71f.
//
// Solidity: function getFileHash() constant returns(string)
func (_File *FileCaller) GetFileHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "getFileHash")
	return *ret0, err
}

// GetFileHash is a free data retrieval call binding the contract method 0x8493f71f.
//
// Solidity: function getFileHash() constant returns(string)
func (_File *FileSession) GetFileHash() (string, error) {
	return _File.Contract.GetFileHash(&_File.CallOpts)
}

// GetFileHash is a free data retrieval call binding the contract method 0x8493f71f.
//
// Solidity: function getFileHash() constant returns(string)
func (_File *FileCallerSession) GetFileHash() (string, error) {
	return _File.Contract.GetFileHash(&_File.CallOpts)
}

// GetFileName is a free data retrieval call binding the contract method 0xfcf6c23c.
//
// Solidity: function getFileName() constant returns(string)
func (_File *FileCaller) GetFileName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "getFileName")
	return *ret0, err
}

// GetFileName is a free data retrieval call binding the contract method 0xfcf6c23c.
//
// Solidity: function getFileName() constant returns(string)
func (_File *FileSession) GetFileName() (string, error) {
	return _File.Contract.GetFileName(&_File.CallOpts)
}

// GetFileName is a free data retrieval call binding the contract method 0xfcf6c23c.
//
// Solidity: function getFileName() constant returns(string)
func (_File *FileCallerSession) GetFileName() (string, error) {
	return _File.Contract.GetFileName(&_File.CallOpts)
}

// GetIv is a free data retrieval call binding the contract method 0xdbfe04af.
//
// Solidity: function getIv() constant returns(string)
func (_File *FileCaller) GetIv(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "getIv")
	return *ret0, err
}

// GetIv is a free data retrieval call binding the contract method 0xdbfe04af.
//
// Solidity: function getIv() constant returns(string)
func (_File *FileSession) GetIv() (string, error) {
	return _File.Contract.GetIv(&_File.CallOpts)
}

// GetIv is a free data retrieval call binding the contract method 0xdbfe04af.
//
// Solidity: function getIv() constant returns(string)
func (_File *FileCallerSession) GetIv() (string, error) {
	return _File.Contract.GetIv(&_File.CallOpts)
}

// GetSalt is a free data retrieval call binding the contract method 0x13a9589c.
//
// Solidity: function getSalt() constant returns(string)
func (_File *FileCaller) GetSalt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "getSalt")
	return *ret0, err
}

// GetSalt is a free data retrieval call binding the contract method 0x13a9589c.
//
// Solidity: function getSalt() constant returns(string)
func (_File *FileSession) GetSalt() (string, error) {
	return _File.Contract.GetSalt(&_File.CallOpts)
}

// GetSalt is a free data retrieval call binding the contract method 0x13a9589c.
//
// Solidity: function getSalt() constant returns(string)
func (_File *FileCallerSession) GetSalt() (string, error) {
	return _File.Contract.GetSalt(&_File.CallOpts)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address user) constant returns(bool)
func (_File *FileCaller) HasAccess(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "hasAccess", user)
	return *ret0, err
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address user) constant returns(bool)
func (_File *FileSession) HasAccess(user common.Address) (bool, error) {
	return _File.Contract.HasAccess(&_File.CallOpts, user)
}

// HasAccess is a free data retrieval call binding the contract method 0x95a078e8.
//
// Solidity: function hasAccess(address user) constant returns(bool)
func (_File *FileCallerSession) HasAccess(user common.Address) (bool, error) {
	return _File.Contract.HasAccess(&_File.CallOpts, user)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address user) constant returns(bool)
func (_File *FileCaller) IsOwner(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _File.contract.Call(opts, out, "isOwner", user)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address user) constant returns(bool)
func (_File *FileSession) IsOwner(user common.Address) (bool, error) {
	return _File.Contract.IsOwner(&_File.CallOpts, user)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address user) constant returns(bool)
func (_File *FileCallerSession) IsOwner(user common.Address) (bool, error) {
	return _File.Contract.IsOwner(&_File.CallOpts, user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_File *FileTransactor) AddUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _File.contract.Transact(opts, "addUser", user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_File *FileSession) AddUser(user common.Address) (*types.Transaction, error) {
	return _File.Contract.AddUser(&_File.TransactOpts, user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_File *FileTransactorSession) AddUser(user common.Address) (*types.Transaction, error) {
	return _File.Contract.AddUser(&_File.TransactOpts, user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address user) returns()
func (_File *FileTransactor) RemoveUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _File.contract.Transact(opts, "removeUser", user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address user) returns()
func (_File *FileSession) RemoveUser(user common.Address) (*types.Transaction, error) {
	return _File.Contract.RemoveUser(&_File.TransactOpts, user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address user) returns()
func (_File *FileTransactorSession) RemoveUser(user common.Address) (*types.Transaction, error) {
	return _File.Contract.RemoveUser(&_File.TransactOpts, user)
}
