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

// FileFactoryABI is the input ABI used to generate the binding from.
const FileFactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"file\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FileCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fileHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_iv\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_salt\",\"type\":\"string\"}],\"name\":\"createFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FileFactory is an auto generated Go binding around an Ethereum contract.
type FileFactory struct {
	FileFactoryCaller     // Read-only binding to the contract
	FileFactoryTransactor // Write-only binding to the contract
	FileFactoryFilterer   // Log filterer for contract events
}

// FileFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileFactorySession struct {
	Contract     *FileFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileFactoryCallerSession struct {
	Contract *FileFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FileFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileFactoryTransactorSession struct {
	Contract     *FileFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FileFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileFactoryRaw struct {
	Contract *FileFactory // Generic contract binding to access the raw methods on
}

// FileFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileFactoryCallerRaw struct {
	Contract *FileFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FileFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileFactoryTransactorRaw struct {
	Contract *FileFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileFactory creates a new instance of FileFactory, bound to a specific deployed contract.
func NewFileFactory(address common.Address, backend bind.ContractBackend) (*FileFactory, error) {
	contract, err := bindFileFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileFactory{FileFactoryCaller: FileFactoryCaller{contract: contract}, FileFactoryTransactor: FileFactoryTransactor{contract: contract}, FileFactoryFilterer: FileFactoryFilterer{contract: contract}}, nil
}

// NewFileFactoryCaller creates a new read-only instance of FileFactory, bound to a specific deployed contract.
func NewFileFactoryCaller(address common.Address, caller bind.ContractCaller) (*FileFactoryCaller, error) {
	contract, err := bindFileFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileFactoryCaller{contract: contract}, nil
}

// NewFileFactoryTransactor creates a new write-only instance of FileFactory, bound to a specific deployed contract.
func NewFileFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FileFactoryTransactor, error) {
	contract, err := bindFileFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileFactoryTransactor{contract: contract}, nil
}

// NewFileFactoryFilterer creates a new log filterer instance of FileFactory, bound to a specific deployed contract.
func NewFileFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FileFactoryFilterer, error) {
	contract, err := bindFileFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileFactoryFilterer{contract: contract}, nil
}

// bindFileFactory binds a generic wrapper to an already deployed contract.
func bindFileFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileFactory *FileFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileFactory.Contract.FileFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileFactory *FileFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileFactory.Contract.FileFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileFactory *FileFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileFactory.Contract.FileFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileFactory *FileFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileFactory *FileFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileFactory *FileFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateFile is a paid mutator transaction binding the contract method 0xd590e7ef.
//
// Solidity: function createFile(string _fileHash, string _fileName, string _iv, string _salt) returns()
func (_FileFactory *FileFactoryTransactor) CreateFile(opts *bind.TransactOpts, _fileHash string, _fileName string, _iv string, _salt string) (*types.Transaction, error) {
	return _FileFactory.contract.Transact(opts, "createFile", _fileHash, _fileName, _iv, _salt)
}

// CreateFile is a paid mutator transaction binding the contract method 0xd590e7ef.
//
// Solidity: function createFile(string _fileHash, string _fileName, string _iv, string _salt) returns()
func (_FileFactory *FileFactorySession) CreateFile(_fileHash string, _fileName string, _iv string, _salt string) (*types.Transaction, error) {
	return _FileFactory.Contract.CreateFile(&_FileFactory.TransactOpts, _fileHash, _fileName, _iv, _salt)
}

// CreateFile is a paid mutator transaction binding the contract method 0xd590e7ef.
//
// Solidity: function createFile(string _fileHash, string _fileName, string _iv, string _salt) returns()
func (_FileFactory *FileFactoryTransactorSession) CreateFile(_fileHash string, _fileName string, _iv string, _salt string) (*types.Transaction, error) {
	return _FileFactory.Contract.CreateFile(&_FileFactory.TransactOpts, _fileHash, _fileName, _iv, _salt)
}

// FileFactoryFileCreatedIterator is returned from FilterFileCreated and is used to iterate over the raw logs and unpacked data for FileCreated events raised by the FileFactory contract.
type FileFactoryFileCreatedIterator struct {
	Event *FileFactoryFileCreated // Event containing the contract specifics and raw log

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
func (it *FileFactoryFileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileFactoryFileCreated)
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
		it.Event = new(FileFactoryFileCreated)
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
func (it *FileFactoryFileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileFactoryFileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileFactoryFileCreated represents a FileCreated event raised by the FileFactory contract.
type FileFactoryFileCreated struct {
	File  common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFileCreated is a free log retrieval operation binding the contract event 0xa1dd78172f8da12ccfa5ee29ea0cbc0b880b96476c8155a666449af0aeb0796c.
//
// Solidity: event FileCreated(address file, address owner)
func (_FileFactory *FileFactoryFilterer) FilterFileCreated(opts *bind.FilterOpts) (*FileFactoryFileCreatedIterator, error) {

	logs, sub, err := _FileFactory.contract.FilterLogs(opts, "FileCreated")
	if err != nil {
		return nil, err
	}
	return &FileFactoryFileCreatedIterator{contract: _FileFactory.contract, event: "FileCreated", logs: logs, sub: sub}, nil
}

// WatchFileCreated is a free log subscription operation binding the contract event 0xa1dd78172f8da12ccfa5ee29ea0cbc0b880b96476c8155a666449af0aeb0796c.
//
// Solidity: event FileCreated(address file, address owner)
func (_FileFactory *FileFactoryFilterer) WatchFileCreated(opts *bind.WatchOpts, sink chan<- *FileFactoryFileCreated) (event.Subscription, error) {

	logs, sub, err := _FileFactory.contract.WatchLogs(opts, "FileCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileFactoryFileCreated)
				if err := _FileFactory.contract.UnpackLog(event, "FileCreated", log); err != nil {
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

// ParseFileCreated is a log parse operation binding the contract event 0xa1dd78172f8da12ccfa5ee29ea0cbc0b880b96476c8155a666449af0aeb0796c.
//
// Solidity: event FileCreated(address file, address owner)
func (_FileFactory *FileFactoryFilterer) ParseFileCreated(log types.Log) (*FileFactoryFileCreated, error) {
	event := new(FileFactoryFileCreated)
	if err := _FileFactory.contract.UnpackLog(event, "FileCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
