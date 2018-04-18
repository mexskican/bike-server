// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// BikeTokenABI is the input ABI used to generate the binding from.
const BikeTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"closeSales\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"founder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"salesOpen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// BikeToken is an auto generated Go binding around an Ethereum contract.
type BikeToken struct {
	BikeTokenCaller     // Read-only binding to the contract
	BikeTokenTransactor // Write-only binding to the contract
	BikeTokenFilterer   // Log filterer for contract events
}

// BikeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BikeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BikeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BikeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BikeTokenSession struct {
	Contract     *BikeToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BikeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BikeTokenCallerSession struct {
	Contract *BikeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BikeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BikeTokenTransactorSession struct {
	Contract     *BikeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BikeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BikeTokenRaw struct {
	Contract *BikeToken // Generic contract binding to access the raw methods on
}

// BikeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BikeTokenCallerRaw struct {
	Contract *BikeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BikeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BikeTokenTransactorRaw struct {
	Contract *BikeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBikeToken creates a new instance of BikeToken, bound to a specific deployed contract.
func NewBikeToken(address common.Address, backend bind.ContractBackend) (*BikeToken, error) {
	contract, err := bindBikeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BikeToken{BikeTokenCaller: BikeTokenCaller{contract: contract}, BikeTokenTransactor: BikeTokenTransactor{contract: contract}, BikeTokenFilterer: BikeTokenFilterer{contract: contract}}, nil
}

// NewBikeTokenCaller creates a new read-only instance of BikeToken, bound to a specific deployed contract.
func NewBikeTokenCaller(address common.Address, caller bind.ContractCaller) (*BikeTokenCaller, error) {
	contract, err := bindBikeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BikeTokenCaller{contract: contract}, nil
}

// NewBikeTokenTransactor creates a new write-only instance of BikeToken, bound to a specific deployed contract.
func NewBikeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BikeTokenTransactor, error) {
	contract, err := bindBikeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BikeTokenTransactor{contract: contract}, nil
}

// NewBikeTokenFilterer creates a new log filterer instance of BikeToken, bound to a specific deployed contract.
func NewBikeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BikeTokenFilterer, error) {
	contract, err := bindBikeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BikeTokenFilterer{contract: contract}, nil
}

// bindBikeToken binds a generic wrapper to an already deployed contract.
func bindBikeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BikeToken *BikeTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BikeToken.Contract.BikeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BikeToken *BikeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeToken.Contract.BikeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BikeToken *BikeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BikeToken.Contract.BikeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BikeToken *BikeTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BikeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BikeToken *BikeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BikeToken *BikeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BikeToken.Contract.contract.Transact(opts, method, params...)
}

// PRICE is a free data retrieval call binding the contract method 0x8d859f3e.
//
// Solidity: function PRICE() constant returns(uint256)
func (_BikeToken *BikeTokenCaller) PRICE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "PRICE")
	return *ret0, err
}

// PRICE is a free data retrieval call binding the contract method 0x8d859f3e.
//
// Solidity: function PRICE() constant returns(uint256)
func (_BikeToken *BikeTokenSession) PRICE() (*big.Int, error) {
	return _BikeToken.Contract.PRICE(&_BikeToken.CallOpts)
}

// PRICE is a free data retrieval call binding the contract method 0x8d859f3e.
//
// Solidity: function PRICE() constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) PRICE() (*big.Int, error) {
	return _BikeToken.Contract.PRICE(&_BikeToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_BikeToken *BikeTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_BikeToken *BikeTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Allowance(&_BikeToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_BikeToken *BikeTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Allowance(&_BikeToken.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_BikeToken *BikeTokenCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_BikeToken *BikeTokenSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Allowed(&_BikeToken.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Allowed(&_BikeToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BikeToken *BikeTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BikeToken *BikeTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BikeToken.Contract.BalanceOf(&_BikeToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BikeToken *BikeTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BikeToken.Contract.BalanceOf(&_BikeToken.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BikeToken *BikeTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BikeToken *BikeTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Balances(&_BikeToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _BikeToken.Contract.Balances(&_BikeToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BikeToken *BikeTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BikeToken *BikeTokenSession) Decimals() (*big.Int, error) {
	return _BikeToken.Contract.Decimals(&_BikeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) Decimals() (*big.Int, error) {
	return _BikeToken.Contract.Decimals(&_BikeToken.CallOpts)
}

// Founder is a free data retrieval call binding the contract method 0x4d853ee5.
//
// Solidity: function founder() constant returns(address)
func (_BikeToken *BikeTokenCaller) Founder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "founder")
	return *ret0, err
}

// Founder is a free data retrieval call binding the contract method 0x4d853ee5.
//
// Solidity: function founder() constant returns(address)
func (_BikeToken *BikeTokenSession) Founder() (common.Address, error) {
	return _BikeToken.Contract.Founder(&_BikeToken.CallOpts)
}

// Founder is a free data retrieval call binding the contract method 0x4d853ee5.
//
// Solidity: function founder() constant returns(address)
func (_BikeToken *BikeTokenCallerSession) Founder() (common.Address, error) {
	return _BikeToken.Contract.Founder(&_BikeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BikeToken *BikeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BikeToken *BikeTokenSession) Name() (string, error) {
	return _BikeToken.Contract.Name(&_BikeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BikeToken *BikeTokenCallerSession) Name() (string, error) {
	return _BikeToken.Contract.Name(&_BikeToken.CallOpts)
}

// SalesOpen is a free data retrieval call binding the contract method 0x729205e7.
//
// Solidity: function salesOpen() constant returns(bool)
func (_BikeToken *BikeTokenCaller) SalesOpen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "salesOpen")
	return *ret0, err
}

// SalesOpen is a free data retrieval call binding the contract method 0x729205e7.
//
// Solidity: function salesOpen() constant returns(bool)
func (_BikeToken *BikeTokenSession) SalesOpen() (bool, error) {
	return _BikeToken.Contract.SalesOpen(&_BikeToken.CallOpts)
}

// SalesOpen is a free data retrieval call binding the contract method 0x729205e7.
//
// Solidity: function salesOpen() constant returns(bool)
func (_BikeToken *BikeTokenCallerSession) SalesOpen() (bool, error) {
	return _BikeToken.Contract.SalesOpen(&_BikeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BikeToken *BikeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BikeToken *BikeTokenSession) Symbol() (string, error) {
	return _BikeToken.Contract.Symbol(&_BikeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BikeToken *BikeTokenCallerSession) Symbol() (string, error) {
	return _BikeToken.Contract.Symbol(&_BikeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BikeToken *BikeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BikeToken *BikeTokenSession) TotalSupply() (*big.Int, error) {
	return _BikeToken.Contract.TotalSupply(&_BikeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BikeToken *BikeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BikeToken.Contract.TotalSupply(&_BikeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.Approve(&_BikeToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.Approve(&_BikeToken.TransactOpts, _spender, _value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_BikeToken *BikeTokenTransactor) BuyTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "buyTokens")
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_BikeToken *BikeTokenSession) BuyTokens() (*types.Transaction, error) {
	return _BikeToken.Contract.BuyTokens(&_BikeToken.TransactOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd0febe4c.
//
// Solidity: function buyTokens() returns()
func (_BikeToken *BikeTokenTransactorSession) BuyTokens() (*types.Transaction, error) {
	return _BikeToken.Contract.BuyTokens(&_BikeToken.TransactOpts)
}

// CloseSales is a paid mutator transaction binding the contract method 0x15a8f073.
//
// Solidity: function closeSales() returns()
func (_BikeToken *BikeTokenTransactor) CloseSales(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "closeSales")
}

// CloseSales is a paid mutator transaction binding the contract method 0x15a8f073.
//
// Solidity: function closeSales() returns()
func (_BikeToken *BikeTokenSession) CloseSales() (*types.Transaction, error) {
	return _BikeToken.Contract.CloseSales(&_BikeToken.TransactOpts)
}

// CloseSales is a paid mutator transaction binding the contract method 0x15a8f073.
//
// Solidity: function closeSales() returns()
func (_BikeToken *BikeTokenTransactorSession) CloseSales() (*types.Transaction, error) {
	return _BikeToken.Contract.CloseSales(&_BikeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.Transfer(&_BikeToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.Transfer(&_BikeToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.TransferFrom(&_BikeToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_BikeToken *BikeTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BikeToken.Contract.TransferFrom(&_BikeToken.TransactOpts, _from, _to, _value)
}

// BikeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BikeToken contract.
type BikeTokenApprovalIterator struct {
	Event *BikeTokenApproval // Event containing the contract specifics and raw log

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
func (it *BikeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeTokenApproval)
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
		it.Event = new(BikeTokenApproval)
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
func (it *BikeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeTokenApproval represents a Approval event raised by the BikeToken contract.
type BikeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_BikeToken *BikeTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*BikeTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _BikeToken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &BikeTokenApprovalIterator{contract: _BikeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_BikeToken *BikeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BikeTokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _BikeToken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeTokenApproval)
				if err := _BikeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// BikeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BikeToken contract.
type BikeTokenTransferIterator struct {
	Event *BikeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BikeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeTokenTransfer)
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
		it.Event = new(BikeTokenTransfer)
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
func (it *BikeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeTokenTransfer represents a Transfer event raised by the BikeToken contract.
type BikeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_BikeToken *BikeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*BikeTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BikeToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &BikeTokenTransferIterator{contract: _BikeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_BikeToken *BikeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BikeTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _BikeToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeTokenTransfer)
				if err := _BikeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
