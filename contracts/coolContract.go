// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CoolContractMetaData contains all meta data concerning the CoolContract contract.
var CoolContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610208806100206000396000f3fe6080604052600436106100295760003560e01c8063429a1bb01461002e578063ed21248c14610059575b600080fd5b34801561003a57600080fd5b50610043610063565b60405161005091906100d6565b60405180910390f35b61006161006c565b005b60008054905090565b3460008082825461007d9190610120565b925050819055507f20a3fe276908af20d1fa2f8f48225ac6a1997eb2d38695c580adeb4124a1a59b336040516100b391906101b7565b60405180910390a1565b6000819050919050565b6100d0816100bd565b82525050565b60006020820190506100eb60008301846100c7565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061012b826100bd565b9150610136836100bd565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561016b5761016a6100f1565b5b828201905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101a182610176565b9050919050565b6101b181610196565b82525050565b60006020820190506101cc60008301846101a8565b9291505056fea264697066735822122098d45f8207162f2ca19d8e6cd9a7602a81fe1962d77a5104d9a4af72b13436c564736f6c634300080e0033",
}

// CoolContractABI is the input ABI used to generate the binding from.
// Deprecated: Use CoolContractMetaData.ABI instead.
var CoolContractABI = CoolContractMetaData.ABI

// CoolContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CoolContractMetaData.Bin instead.
var CoolContractBin = CoolContractMetaData.Bin

// DeployCoolContract deploys a new Ethereum contract, binding an instance of CoolContract to it.
func DeployCoolContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CoolContract, error) {
	parsed, err := CoolContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CoolContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoolContract{CoolContractCaller: CoolContractCaller{contract: contract}, CoolContractTransactor: CoolContractTransactor{contract: contract}, CoolContractFilterer: CoolContractFilterer{contract: contract}}, nil
}

// CoolContract is an auto generated Go binding around an Ethereum contract.
type CoolContract struct {
	CoolContractCaller     // Read-only binding to the contract
	CoolContractTransactor // Write-only binding to the contract
	CoolContractFilterer   // Log filterer for contract events
}

// CoolContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoolContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoolContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoolContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoolContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoolContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoolContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoolContractSession struct {
	Contract     *CoolContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoolContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoolContractCallerSession struct {
	Contract *CoolContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CoolContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoolContractTransactorSession struct {
	Contract     *CoolContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CoolContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoolContractRaw struct {
	Contract *CoolContract // Generic contract binding to access the raw methods on
}

// CoolContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoolContractCallerRaw struct {
	Contract *CoolContractCaller // Generic read-only contract binding to access the raw methods on
}

// CoolContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoolContractTransactorRaw struct {
	Contract *CoolContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoolContract creates a new instance of CoolContract, bound to a specific deployed contract.
func NewCoolContract(address common.Address, backend bind.ContractBackend) (*CoolContract, error) {
	contract, err := bindCoolContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoolContract{CoolContractCaller: CoolContractCaller{contract: contract}, CoolContractTransactor: CoolContractTransactor{contract: contract}, CoolContractFilterer: CoolContractFilterer{contract: contract}}, nil
}

// NewCoolContractCaller creates a new read-only instance of CoolContract, bound to a specific deployed contract.
func NewCoolContractCaller(address common.Address, caller bind.ContractCaller) (*CoolContractCaller, error) {
	contract, err := bindCoolContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoolContractCaller{contract: contract}, nil
}

// NewCoolContractTransactor creates a new write-only instance of CoolContract, bound to a specific deployed contract.
func NewCoolContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CoolContractTransactor, error) {
	contract, err := bindCoolContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoolContractTransactor{contract: contract}, nil
}

// NewCoolContractFilterer creates a new log filterer instance of CoolContract, bound to a specific deployed contract.
func NewCoolContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CoolContractFilterer, error) {
	contract, err := bindCoolContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoolContractFilterer{contract: contract}, nil
}

// bindCoolContract binds a generic wrapper to an already deployed contract.
func bindCoolContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoolContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoolContract *CoolContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoolContract.Contract.CoolContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoolContract *CoolContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoolContract.Contract.CoolContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoolContract *CoolContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoolContract.Contract.CoolContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoolContract *CoolContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoolContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoolContract *CoolContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoolContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoolContract *CoolContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoolContract.Contract.contract.Transact(opts, method, params...)
}

// SeeBalance is a free data retrieval call binding the contract method 0x429a1bb0.
//
// Solidity: function seeBalance() view returns(uint256)
func (_CoolContract *CoolContractCaller) SeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CoolContract.contract.Call(opts, &out, "seeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SeeBalance is a free data retrieval call binding the contract method 0x429a1bb0.
//
// Solidity: function seeBalance() view returns(uint256)
func (_CoolContract *CoolContractSession) SeeBalance() (*big.Int, error) {
	return _CoolContract.Contract.SeeBalance(&_CoolContract.CallOpts)
}

// SeeBalance is a free data retrieval call binding the contract method 0x429a1bb0.
//
// Solidity: function seeBalance() view returns(uint256)
func (_CoolContract *CoolContractCallerSession) SeeBalance() (*big.Int, error) {
	return _CoolContract.Contract.SeeBalance(&_CoolContract.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_CoolContract *CoolContractTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoolContract.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_CoolContract *CoolContractSession) Deposit() (*types.Transaction, error) {
	return _CoolContract.Contract.Deposit(&_CoolContract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_CoolContract *CoolContractTransactorSession) Deposit() (*types.Transaction, error) {
	return _CoolContract.Contract.Deposit(&_CoolContract.TransactOpts)
}

// CoolContractDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the CoolContract contract.
type CoolContractDepositedIterator struct {
	Event *CoolContractDeposited // Event containing the contract specifics and raw log

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
func (it *CoolContractDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoolContractDeposited)
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
		it.Event = new(CoolContractDeposited)
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
func (it *CoolContractDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoolContractDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoolContractDeposited represents a Deposited event raised by the CoolContract contract.
type CoolContractDeposited struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x20a3fe276908af20d1fa2f8f48225ac6a1997eb2d38695c580adeb4124a1a59b.
//
// Solidity: event Deposited(address addr)
func (_CoolContract *CoolContractFilterer) FilterDeposited(opts *bind.FilterOpts) (*CoolContractDepositedIterator, error) {

	logs, sub, err := _CoolContract.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &CoolContractDepositedIterator{contract: _CoolContract.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x20a3fe276908af20d1fa2f8f48225ac6a1997eb2d38695c580adeb4124a1a59b.
//
// Solidity: event Deposited(address addr)
func (_CoolContract *CoolContractFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *CoolContractDeposited) (event.Subscription, error) {

	logs, sub, err := _CoolContract.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoolContractDeposited)
				if err := _CoolContract.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x20a3fe276908af20d1fa2f8f48225ac6a1997eb2d38695c580adeb4124a1a59b.
//
// Solidity: event Deposited(address addr)
func (_CoolContract *CoolContractFilterer) ParseDeposited(log types.Log) (*CoolContractDeposited, error) {
	event := new(CoolContractDeposited)
	if err := _CoolContract.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
