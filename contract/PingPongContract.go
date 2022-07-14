// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Ping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"Pong\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pinger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"pong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Pinger is a free data retrieval call binding the contract method 0xb97399f5.
//
// Solidity: function pinger() view returns(address)
func (_Contract *ContractCaller) Pinger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pinger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pinger is a free data retrieval call binding the contract method 0xb97399f5.
//
// Solidity: function pinger() view returns(address)
func (_Contract *ContractSession) Pinger() (common.Address, error) {
	return _Contract.Contract.Pinger(&_Contract.CallOpts)
}

// Pinger is a free data retrieval call binding the contract method 0xb97399f5.
//
// Solidity: function pinger() view returns(address)
func (_Contract *ContractCallerSession) Pinger() (common.Address, error) {
	return _Contract.Contract.Pinger(&_Contract.CallOpts)
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_Contract *ContractTransactor) Ping(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "ping")
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_Contract *ContractSession) Ping() (*types.Transaction, error) {
	return _Contract.Contract.Ping(&_Contract.TransactOpts)
}

// Ping is a paid mutator transaction binding the contract method 0x5c36b186.
//
// Solidity: function ping() returns()
func (_Contract *ContractTransactorSession) Ping() (*types.Transaction, error) {
	return _Contract.Contract.Ping(&_Contract.TransactOpts)
}

// Pong is a paid mutator transaction binding the contract method 0x05ba79a2.
//
// Solidity: function pong(bytes32 _txHash) returns()
func (_Contract *ContractTransactor) Pong(opts *bind.TransactOpts, _txHash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pong", _txHash)
}

// Pong is a paid mutator transaction binding the contract method 0x05ba79a2.
//
// Solidity: function pong(bytes32 _txHash) returns()
func (_Contract *ContractSession) Pong(_txHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Pong(&_Contract.TransactOpts, _txHash)
}

// Pong is a paid mutator transaction binding the contract method 0x05ba79a2.
//
// Solidity: function pong(bytes32 _txHash) returns()
func (_Contract *ContractTransactorSession) Pong(_txHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Pong(&_Contract.TransactOpts, _txHash)
}

// ContractPingIterator is returned from FilterPing and is used to iterate over the raw logs and unpacked data for Ping events raised by the Contract contract.
type ContractPingIterator struct {
	Event *ContractPing // Event containing the contract specifics and raw log

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
func (it *ContractPingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPing)
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
		it.Event = new(ContractPing)
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
func (it *ContractPingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPing represents a Ping event raised by the Contract contract.
type ContractPing struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPing is a free log retrieval operation binding the contract event 0xca6e822df923f741dfe968d15d80a18abd25bd1e748bcb9ad81fea5bbb7386af.
//
// Solidity: event Ping()
func (_Contract *ContractFilterer) FilterPing(opts *bind.FilterOpts) (*ContractPingIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return &ContractPingIterator{contract: _Contract.contract, event: "Ping", logs: logs, sub: sub}, nil
}

// WatchPing is a free log subscription operation binding the contract event 0xca6e822df923f741dfe968d15d80a18abd25bd1e748bcb9ad81fea5bbb7386af.
//
// Solidity: event Ping()
func (_Contract *ContractFilterer) WatchPing(opts *bind.WatchOpts, sink chan<- *ContractPing) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Ping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPing)
				if err := _Contract.contract.UnpackLog(event, "Ping", log); err != nil {
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

// ParsePing is a log parse operation binding the contract event 0xca6e822df923f741dfe968d15d80a18abd25bd1e748bcb9ad81fea5bbb7386af.
//
// Solidity: event Ping()
func (_Contract *ContractFilterer) ParsePing(log types.Log) (*ContractPing, error) {
	event := new(ContractPing)
	if err := _Contract.contract.UnpackLog(event, "Ping", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPongIterator is returned from FilterPong and is used to iterate over the raw logs and unpacked data for Pong events raised by the Contract contract.
type ContractPongIterator struct {
	Event *ContractPong // Event containing the contract specifics and raw log

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
func (it *ContractPongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPong)
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
		it.Event = new(ContractPong)
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
func (it *ContractPongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPong represents a Pong event raised by the Contract contract.
type ContractPong struct {
	TxHash [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPong is a free log retrieval operation binding the contract event 0x67050610046771547cf1d6e467b904ccfc523370eebc895dad1d9a73349b9804.
//
// Solidity: event Pong(bytes32 txHash)
func (_Contract *ContractFilterer) FilterPong(opts *bind.FilterOpts) (*ContractPongIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return &ContractPongIterator{contract: _Contract.contract, event: "Pong", logs: logs, sub: sub}, nil
}

// WatchPong is a free log subscription operation binding the contract event 0x67050610046771547cf1d6e467b904ccfc523370eebc895dad1d9a73349b9804.
//
// Solidity: event Pong(bytes32 txHash)
func (_Contract *ContractFilterer) WatchPong(opts *bind.WatchOpts, sink chan<- *ContractPong) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPong)
				if err := _Contract.contract.UnpackLog(event, "Pong", log); err != nil {
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

// ParsePong is a log parse operation binding the contract event 0x67050610046771547cf1d6e467b904ccfc523370eebc895dad1d9a73349b9804.
//
// Solidity: event Pong(bytes32 txHash)
func (_Contract *ContractFilterer) ParsePong(log types.Log) (*ContractPong, error) {
	event := new(ContractPong)
	if err := _Contract.contract.UnpackLog(event, "Pong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
