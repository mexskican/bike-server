// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bike

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

// BikeABI is the input ABI used to generate the binding from.
const BikeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"transferFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hourlyCost\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeTokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_renter\",\"type\":\"address\"}],\"name\":\"getReturnTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bikes\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"available\",\"type\":\"bool\"},{\"name\":\"added\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"rentals\",\"outputs\":[{\"name\":\"bikeId\",\"type\":\"uint256\"},{\"name\":\"renter\",\"type\":\"address\"},{\"name\":\"hoursRent\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"returnTime\",\"type\":\"uint256\"},{\"name\":\"renting\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"removeBike\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"bikeIsAvailable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"name\":\"_hoursRent\",\"type\":\"uint256\"}],\"name\":\"rentBike\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCost\",\"type\":\"uint256\"}],\"name\":\"updateHourlyCost\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"name\":\"_renter\",\"type\":\"address\"}],\"name\":\"returnBike\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_renter\",\"type\":\"address\"}],\"name\":\"getRentalBikeId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"addBike\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"name\":\"_previousRenter\",\"type\":\"address\"}],\"name\":\"transferRentedBike\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_hourlyCost\",\"type\":\"uint256\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"bikeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"}],\"name\":\"bikeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"hoursRent\",\"type\":\"uint256\"}],\"name\":\"bikeRented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_previousRenter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_newRenter\",\"type\":\"address\"}],\"name\":\"bikeTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_bikeId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_onTime\",\"type\":\"bool\"}],\"name\":\"bikeReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_hourlyCost\",\"type\":\"uint256\"}],\"name\":\"newCost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Bike is an auto generated Go binding around an Ethereum contract.
type Bike struct {
	BikeCaller     // Read-only binding to the contract
	BikeTransactor // Write-only binding to the contract
	BikeFilterer   // Log filterer for contract events
}

// BikeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BikeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BikeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BikeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BikeSession struct {
	Contract     *Bike             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BikeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BikeCallerSession struct {
	Contract *BikeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BikeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BikeTransactorSession struct {
	Contract     *BikeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BikeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BikeRaw struct {
	Contract *Bike // Generic contract binding to access the raw methods on
}

// BikeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BikeCallerRaw struct {
	Contract *BikeCaller // Generic read-only contract binding to access the raw methods on
}

// BikeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BikeTransactorRaw struct {
	Contract *BikeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBike creates a new instance of Bike, bound to a specific deployed contract.
func NewBike(address common.Address, backend bind.ContractBackend) (*Bike, error) {
	contract, err := bindBike(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bike{BikeCaller: BikeCaller{contract: contract}, BikeTransactor: BikeTransactor{contract: contract}, BikeFilterer: BikeFilterer{contract: contract}}, nil
}

// NewBikeCaller creates a new read-only instance of Bike, bound to a specific deployed contract.
func NewBikeCaller(address common.Address, caller bind.ContractCaller) (*BikeCaller, error) {
	contract, err := bindBike(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BikeCaller{contract: contract}, nil
}

// NewBikeTransactor creates a new write-only instance of Bike, bound to a specific deployed contract.
func NewBikeTransactor(address common.Address, transactor bind.ContractTransactor) (*BikeTransactor, error) {
	contract, err := bindBike(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BikeTransactor{contract: contract}, nil
}

// NewBikeFilterer creates a new log filterer instance of Bike, bound to a specific deployed contract.
func NewBikeFilterer(address common.Address, filterer bind.ContractFilterer) (*BikeFilterer, error) {
	contract, err := bindBike(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BikeFilterer{contract: contract}, nil
}

// bindBike binds a generic wrapper to an already deployed contract.
func bindBike(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bike *BikeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bike.Contract.BikeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bike *BikeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bike.Contract.BikeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bike *BikeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bike.Contract.BikeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bike *BikeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bike.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bike *BikeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bike.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bike *BikeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bike.Contract.contract.Transact(opts, method, params...)
}

// BikeIsAvailable is a free data retrieval call binding the contract method 0xa708810e.
//
// Solidity: function bikeIsAvailable(_bikeId uint256) constant returns(bool)
func (_Bike *BikeCaller) BikeIsAvailable(opts *bind.CallOpts, _bikeId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "bikeIsAvailable", _bikeId)
	return *ret0, err
}

// BikeIsAvailable is a free data retrieval call binding the contract method 0xa708810e.
//
// Solidity: function bikeIsAvailable(_bikeId uint256) constant returns(bool)
func (_Bike *BikeSession) BikeIsAvailable(_bikeId *big.Int) (bool, error) {
	return _Bike.Contract.BikeIsAvailable(&_Bike.CallOpts, _bikeId)
}

// BikeIsAvailable is a free data retrieval call binding the contract method 0xa708810e.
//
// Solidity: function bikeIsAvailable(_bikeId uint256) constant returns(bool)
func (_Bike *BikeCallerSession) BikeIsAvailable(_bikeId *big.Int) (bool, error) {
	return _Bike.Contract.BikeIsAvailable(&_Bike.CallOpts, _bikeId)
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, available bool, added bool)
func (_Bike *BikeCaller) Bikes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Available bool
	Added     bool
}, error) {
	ret := new(struct {
		Id        *big.Int
		Available bool
		Added     bool
	})
	out := ret
	err := _Bike.contract.Call(opts, out, "bikes", arg0)
	return *ret, err
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, available bool, added bool)
func (_Bike *BikeSession) Bikes(arg0 *big.Int) (struct {
	Id        *big.Int
	Available bool
	Added     bool
}, error) {
	return _Bike.Contract.Bikes(&_Bike.CallOpts, arg0)
}

// Bikes is a free data retrieval call binding the contract method 0x5c526e19.
//
// Solidity: function bikes( uint256) constant returns(id uint256, available bool, added bool)
func (_Bike *BikeCallerSession) Bikes(arg0 *big.Int) (struct {
	Id        *big.Int
	Available bool
	Added     bool
}, error) {
	return _Bike.Contract.Bikes(&_Bike.CallOpts, arg0)
}

// GetRentalBikeId is a free data retrieval call binding the contract method 0xb88055e9.
//
// Solidity: function getRentalBikeId(_renter address) constant returns(uint256)
func (_Bike *BikeCaller) GetRentalBikeId(opts *bind.CallOpts, _renter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "getRentalBikeId", _renter)
	return *ret0, err
}

// GetRentalBikeId is a free data retrieval call binding the contract method 0xb88055e9.
//
// Solidity: function getRentalBikeId(_renter address) constant returns(uint256)
func (_Bike *BikeSession) GetRentalBikeId(_renter common.Address) (*big.Int, error) {
	return _Bike.Contract.GetRentalBikeId(&_Bike.CallOpts, _renter)
}

// GetRentalBikeId is a free data retrieval call binding the contract method 0xb88055e9.
//
// Solidity: function getRentalBikeId(_renter address) constant returns(uint256)
func (_Bike *BikeCallerSession) GetRentalBikeId(_renter common.Address) (*big.Int, error) {
	return _Bike.Contract.GetRentalBikeId(&_Bike.CallOpts, _renter)
}

// GetReturnTime is a free data retrieval call binding the contract method 0x31f3ef00.
//
// Solidity: function getReturnTime(_renter address) constant returns(uint256)
func (_Bike *BikeCaller) GetReturnTime(opts *bind.CallOpts, _renter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "getReturnTime", _renter)
	return *ret0, err
}

// GetReturnTime is a free data retrieval call binding the contract method 0x31f3ef00.
//
// Solidity: function getReturnTime(_renter address) constant returns(uint256)
func (_Bike *BikeSession) GetReturnTime(_renter common.Address) (*big.Int, error) {
	return _Bike.Contract.GetReturnTime(&_Bike.CallOpts, _renter)
}

// GetReturnTime is a free data retrieval call binding the contract method 0x31f3ef00.
//
// Solidity: function getReturnTime(_renter address) constant returns(uint256)
func (_Bike *BikeCallerSession) GetReturnTime(_renter common.Address) (*big.Int, error) {
	return _Bike.Contract.GetReturnTime(&_Bike.CallOpts, _renter)
}

// HourlyCost is a free data retrieval call binding the contract method 0x184880f4.
//
// Solidity: function hourlyCost() constant returns(uint256)
func (_Bike *BikeCaller) HourlyCost(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "hourlyCost")
	return *ret0, err
}

// HourlyCost is a free data retrieval call binding the contract method 0x184880f4.
//
// Solidity: function hourlyCost() constant returns(uint256)
func (_Bike *BikeSession) HourlyCost() (*big.Int, error) {
	return _Bike.Contract.HourlyCost(&_Bike.CallOpts)
}

// HourlyCost is a free data retrieval call binding the contract method 0x184880f4.
//
// Solidity: function hourlyCost() constant returns(uint256)
func (_Bike *BikeCallerSession) HourlyCost() (*big.Int, error) {
	return _Bike.Contract.HourlyCost(&_Bike.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bike *BikeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bike *BikeSession) Owner() (common.Address, error) {
	return _Bike.Contract.Owner(&_Bike.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bike *BikeCallerSession) Owner() (common.Address, error) {
	return _Bike.Contract.Owner(&_Bike.CallOpts)
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(bikeId uint256, renter address, hoursRent uint256, price uint256, returnTime uint256, renting bool)
func (_Bike *BikeCaller) Rentals(opts *bind.CallOpts, arg0 common.Address) (struct {
	BikeId     *big.Int
	Renter     common.Address
	HoursRent  *big.Int
	Price      *big.Int
	ReturnTime *big.Int
	Renting    bool
}, error) {
	ret := new(struct {
		BikeId     *big.Int
		Renter     common.Address
		HoursRent  *big.Int
		Price      *big.Int
		ReturnTime *big.Int
		Renting    bool
	})
	out := ret
	err := _Bike.contract.Call(opts, out, "rentals", arg0)
	return *ret, err
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(bikeId uint256, renter address, hoursRent uint256, price uint256, returnTime uint256, renting bool)
func (_Bike *BikeSession) Rentals(arg0 common.Address) (struct {
	BikeId     *big.Int
	Renter     common.Address
	HoursRent  *big.Int
	Price      *big.Int
	ReturnTime *big.Int
	Renting    bool
}, error) {
	return _Bike.Contract.Rentals(&_Bike.CallOpts, arg0)
}

// Rentals is a free data retrieval call binding the contract method 0x6ba85eb4.
//
// Solidity: function rentals( address) constant returns(bikeId uint256, renter address, hoursRent uint256, price uint256, returnTime uint256, renting bool)
func (_Bike *BikeCallerSession) Rentals(arg0 common.Address) (struct {
	BikeId     *big.Int
	Renter     common.Address
	HoursRent  *big.Int
	Price      *big.Int
	ReturnTime *big.Int
	Renting    bool
}, error) {
	return _Bike.Contract.Rentals(&_Bike.CallOpts, arg0)
}

// AddBike is a paid mutator transaction binding the contract method 0xc120f3e5.
//
// Solidity: function addBike(_bikeId uint256) returns(success bool)
func (_Bike *BikeTransactor) AddBike(opts *bind.TransactOpts, _bikeId *big.Int) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "addBike", _bikeId)
}

// AddBike is a paid mutator transaction binding the contract method 0xc120f3e5.
//
// Solidity: function addBike(_bikeId uint256) returns(success bool)
func (_Bike *BikeSession) AddBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.AddBike(&_Bike.TransactOpts, _bikeId)
}

// AddBike is a paid mutator transaction binding the contract method 0xc120f3e5.
//
// Solidity: function addBike(_bikeId uint256) returns(success bool)
func (_Bike *BikeTransactorSession) AddBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.AddBike(&_Bike.TransactOpts, _bikeId)
}

// RemoveBike is a paid mutator transaction binding the contract method 0x9bdd6b6e.
//
// Solidity: function removeBike(_bikeId uint256) returns(success bool)
func (_Bike *BikeTransactor) RemoveBike(opts *bind.TransactOpts, _bikeId *big.Int) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "removeBike", _bikeId)
}

// RemoveBike is a paid mutator transaction binding the contract method 0x9bdd6b6e.
//
// Solidity: function removeBike(_bikeId uint256) returns(success bool)
func (_Bike *BikeSession) RemoveBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.RemoveBike(&_Bike.TransactOpts, _bikeId)
}

// RemoveBike is a paid mutator transaction binding the contract method 0x9bdd6b6e.
//
// Solidity: function removeBike(_bikeId uint256) returns(success bool)
func (_Bike *BikeTransactorSession) RemoveBike(_bikeId *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.RemoveBike(&_Bike.TransactOpts, _bikeId)
}

// RentBike is a paid mutator transaction binding the contract method 0xa786de18.
//
// Solidity: function rentBike(_bikeId uint256, _hoursRent uint256) returns(success bool)
func (_Bike *BikeTransactor) RentBike(opts *bind.TransactOpts, _bikeId *big.Int, _hoursRent *big.Int) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "rentBike", _bikeId, _hoursRent)
}

// RentBike is a paid mutator transaction binding the contract method 0xa786de18.
//
// Solidity: function rentBike(_bikeId uint256, _hoursRent uint256) returns(success bool)
func (_Bike *BikeSession) RentBike(_bikeId *big.Int, _hoursRent *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.RentBike(&_Bike.TransactOpts, _bikeId, _hoursRent)
}

// RentBike is a paid mutator transaction binding the contract method 0xa786de18.
//
// Solidity: function rentBike(_bikeId uint256, _hoursRent uint256) returns(success bool)
func (_Bike *BikeTransactorSession) RentBike(_bikeId *big.Int, _hoursRent *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.RentBike(&_Bike.TransactOpts, _bikeId, _hoursRent)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa9e6e3f5.
//
// Solidity: function returnBike(_bikeId uint256, _renter address) returns(bool)
func (_Bike *BikeTransactor) ReturnBike(opts *bind.TransactOpts, _bikeId *big.Int, _renter common.Address) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "returnBike", _bikeId, _renter)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa9e6e3f5.
//
// Solidity: function returnBike(_bikeId uint256, _renter address) returns(bool)
func (_Bike *BikeSession) ReturnBike(_bikeId *big.Int, _renter common.Address) (*types.Transaction, error) {
	return _Bike.Contract.ReturnBike(&_Bike.TransactOpts, _bikeId, _renter)
}

// ReturnBike is a paid mutator transaction binding the contract method 0xa9e6e3f5.
//
// Solidity: function returnBike(_bikeId uint256, _renter address) returns(bool)
func (_Bike *BikeTransactorSession) ReturnBike(_bikeId *big.Int, _renter common.Address) (*types.Transaction, error) {
	return _Bike.Contract.ReturnBike(&_Bike.TransactOpts, _bikeId, _renter)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_bikeTokenAddress address) returns(bool)
func (_Bike *BikeTransactor) SetTokenAddress(opts *bind.TransactOpts, _bikeTokenAddress common.Address) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "setTokenAddress", _bikeTokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_bikeTokenAddress address) returns(bool)
func (_Bike *BikeSession) SetTokenAddress(_bikeTokenAddress common.Address) (*types.Transaction, error) {
	return _Bike.Contract.SetTokenAddress(&_Bike.TransactOpts, _bikeTokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(_bikeTokenAddress address) returns(bool)
func (_Bike *BikeTransactorSession) SetTokenAddress(_bikeTokenAddress common.Address) (*types.Transaction, error) {
	return _Bike.Contract.SetTokenAddress(&_Bike.TransactOpts, _bikeTokenAddress)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x133ae30b.
//
// Solidity: function transferFunds(_balance uint256) returns(bool)
func (_Bike *BikeTransactor) TransferFunds(opts *bind.TransactOpts, _balance *big.Int) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "transferFunds", _balance)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x133ae30b.
//
// Solidity: function transferFunds(_balance uint256) returns(bool)
func (_Bike *BikeSession) TransferFunds(_balance *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.TransferFunds(&_Bike.TransactOpts, _balance)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x133ae30b.
//
// Solidity: function transferFunds(_balance uint256) returns(bool)
func (_Bike *BikeTransactorSession) TransferFunds(_balance *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.TransferFunds(&_Bike.TransactOpts, _balance)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bike *BikeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bike *BikeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bike.Contract.TransferOwnership(&_Bike.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Bike *BikeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bike.Contract.TransferOwnership(&_Bike.TransactOpts, newOwner)
}

// TransferRentedBike is a paid mutator transaction binding the contract method 0xe97607c6.
//
// Solidity: function transferRentedBike(_bikeId uint256, _previousRenter address) returns(success bool)
func (_Bike *BikeTransactor) TransferRentedBike(opts *bind.TransactOpts, _bikeId *big.Int, _previousRenter common.Address) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "transferRentedBike", _bikeId, _previousRenter)
}

// TransferRentedBike is a paid mutator transaction binding the contract method 0xe97607c6.
//
// Solidity: function transferRentedBike(_bikeId uint256, _previousRenter address) returns(success bool)
func (_Bike *BikeSession) TransferRentedBike(_bikeId *big.Int, _previousRenter common.Address) (*types.Transaction, error) {
	return _Bike.Contract.TransferRentedBike(&_Bike.TransactOpts, _bikeId, _previousRenter)
}

// TransferRentedBike is a paid mutator transaction binding the contract method 0xe97607c6.
//
// Solidity: function transferRentedBike(_bikeId uint256, _previousRenter address) returns(success bool)
func (_Bike *BikeTransactorSession) TransferRentedBike(_bikeId *big.Int, _previousRenter common.Address) (*types.Transaction, error) {
	return _Bike.Contract.TransferRentedBike(&_Bike.TransactOpts, _bikeId, _previousRenter)
}

// UpdateHourlyCost is a paid mutator transaction binding the contract method 0xa9983bcb.
//
// Solidity: function updateHourlyCost(_newCost uint256) returns(success bool)
func (_Bike *BikeTransactor) UpdateHourlyCost(opts *bind.TransactOpts, _newCost *big.Int) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "updateHourlyCost", _newCost)
}

// UpdateHourlyCost is a paid mutator transaction binding the contract method 0xa9983bcb.
//
// Solidity: function updateHourlyCost(_newCost uint256) returns(success bool)
func (_Bike *BikeSession) UpdateHourlyCost(_newCost *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.UpdateHourlyCost(&_Bike.TransactOpts, _newCost)
}

// UpdateHourlyCost is a paid mutator transaction binding the contract method 0xa9983bcb.
//
// Solidity: function updateHourlyCost(_newCost uint256) returns(success bool)
func (_Bike *BikeTransactorSession) UpdateHourlyCost(_newCost *big.Int) (*types.Transaction, error) {
	return _Bike.Contract.UpdateHourlyCost(&_Bike.TransactOpts, _newCost)
}

// BikeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bike contract.
type BikeOwnershipTransferredIterator struct {
	Event *BikeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BikeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeOwnershipTransferred)
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
		it.Event = new(BikeOwnershipTransferred)
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
func (it *BikeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeOwnershipTransferred represents a OwnershipTransferred event raised by the Bike contract.
type BikeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Bike *BikeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BikeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bike.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BikeOwnershipTransferredIterator{contract: _Bike.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Bike *BikeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BikeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bike.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeOwnershipTransferred)
				if err := _Bike.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BikeBikeAddedIterator is returned from FilterBikeAdded and is used to iterate over the raw logs and unpacked data for BikeAdded events raised by the Bike contract.
type BikeBikeAddedIterator struct {
	Event *BikeBikeAdded // Event containing the contract specifics and raw log

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
func (it *BikeBikeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeBikeAdded)
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
		it.Event = new(BikeBikeAdded)
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
func (it *BikeBikeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeBikeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeBikeAdded represents a BikeAdded event raised by the Bike contract.
type BikeBikeAdded struct {
	BikeId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBikeAdded is a free log retrieval operation binding the contract event 0xc258a4d499ba0ac7814b8b97b3a19c3f1fb816eefefc370a942bf91b7543c838.
//
// Solidity: event bikeAdded(_bikeId uint256)
func (_Bike *BikeFilterer) FilterBikeAdded(opts *bind.FilterOpts) (*BikeBikeAddedIterator, error) {

	logs, sub, err := _Bike.contract.FilterLogs(opts, "bikeAdded")
	if err != nil {
		return nil, err
	}
	return &BikeBikeAddedIterator{contract: _Bike.contract, event: "bikeAdded", logs: logs, sub: sub}, nil
}

// WatchBikeAdded is a free log subscription operation binding the contract event 0xc258a4d499ba0ac7814b8b97b3a19c3f1fb816eefefc370a942bf91b7543c838.
//
// Solidity: event bikeAdded(_bikeId uint256)
func (_Bike *BikeFilterer) WatchBikeAdded(opts *bind.WatchOpts, sink chan<- *BikeBikeAdded) (event.Subscription, error) {

	logs, sub, err := _Bike.contract.WatchLogs(opts, "bikeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeBikeAdded)
				if err := _Bike.contract.UnpackLog(event, "bikeAdded", log); err != nil {
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

// BikeBikeRemovedIterator is returned from FilterBikeRemoved and is used to iterate over the raw logs and unpacked data for BikeRemoved events raised by the Bike contract.
type BikeBikeRemovedIterator struct {
	Event *BikeBikeRemoved // Event containing the contract specifics and raw log

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
func (it *BikeBikeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeBikeRemoved)
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
		it.Event = new(BikeBikeRemoved)
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
func (it *BikeBikeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeBikeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeBikeRemoved represents a BikeRemoved event raised by the Bike contract.
type BikeBikeRemoved struct {
	BikeId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBikeRemoved is a free log retrieval operation binding the contract event 0x822f820b092d98c84c17f3cab57b3857ea83d4b9a8edc7bb26e4010d288f6627.
//
// Solidity: event bikeRemoved(_bikeId uint256)
func (_Bike *BikeFilterer) FilterBikeRemoved(opts *bind.FilterOpts) (*BikeBikeRemovedIterator, error) {

	logs, sub, err := _Bike.contract.FilterLogs(opts, "bikeRemoved")
	if err != nil {
		return nil, err
	}
	return &BikeBikeRemovedIterator{contract: _Bike.contract, event: "bikeRemoved", logs: logs, sub: sub}, nil
}

// WatchBikeRemoved is a free log subscription operation binding the contract event 0x822f820b092d98c84c17f3cab57b3857ea83d4b9a8edc7bb26e4010d288f6627.
//
// Solidity: event bikeRemoved(_bikeId uint256)
func (_Bike *BikeFilterer) WatchBikeRemoved(opts *bind.WatchOpts, sink chan<- *BikeBikeRemoved) (event.Subscription, error) {

	logs, sub, err := _Bike.contract.WatchLogs(opts, "bikeRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeBikeRemoved)
				if err := _Bike.contract.UnpackLog(event, "bikeRemoved", log); err != nil {
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

// BikeBikeRentedIterator is returned from FilterBikeRented and is used to iterate over the raw logs and unpacked data for BikeRented events raised by the Bike contract.
type BikeBikeRentedIterator struct {
	Event *BikeBikeRented // Event containing the contract specifics and raw log

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
func (it *BikeBikeRentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeBikeRented)
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
		it.Event = new(BikeBikeRented)
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
func (it *BikeBikeRentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeBikeRentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeBikeRented represents a BikeRented event raised by the Bike contract.
type BikeBikeRented struct {
	BikeId    *big.Int
	Renter    common.Address
	HoursRent *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBikeRented is a free log retrieval operation binding the contract event 0xa31ebe8cd208580e0861dc1098d788eb14f15cff38c0ee834f7cb81ec9a51467.
//
// Solidity: event bikeRented(_bikeId uint256, _renter address, hoursRent uint256)
func (_Bike *BikeFilterer) FilterBikeRented(opts *bind.FilterOpts) (*BikeBikeRentedIterator, error) {

	logs, sub, err := _Bike.contract.FilterLogs(opts, "bikeRented")
	if err != nil {
		return nil, err
	}
	return &BikeBikeRentedIterator{contract: _Bike.contract, event: "bikeRented", logs: logs, sub: sub}, nil
}

// WatchBikeRented is a free log subscription operation binding the contract event 0xa31ebe8cd208580e0861dc1098d788eb14f15cff38c0ee834f7cb81ec9a51467.
//
// Solidity: event bikeRented(_bikeId uint256, _renter address, hoursRent uint256)
func (_Bike *BikeFilterer) WatchBikeRented(opts *bind.WatchOpts, sink chan<- *BikeBikeRented) (event.Subscription, error) {

	logs, sub, err := _Bike.contract.WatchLogs(opts, "bikeRented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeBikeRented)
				if err := _Bike.contract.UnpackLog(event, "bikeRented", log); err != nil {
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

// BikeBikeReturnedIterator is returned from FilterBikeReturned and is used to iterate over the raw logs and unpacked data for BikeReturned events raised by the Bike contract.
type BikeBikeReturnedIterator struct {
	Event *BikeBikeReturned // Event containing the contract specifics and raw log

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
func (it *BikeBikeReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeBikeReturned)
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
		it.Event = new(BikeBikeReturned)
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
func (it *BikeBikeReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeBikeReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeBikeReturned represents a BikeReturned event raised by the Bike contract.
type BikeBikeReturned struct {
	BikeId *big.Int
	Renter common.Address
	OnTime bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBikeReturned is a free log retrieval operation binding the contract event 0x6cdb62cd573c17558f2e746f8dba25f431ab46e3a6e078aac047c92f2a1f92d9.
//
// Solidity: event bikeReturned(_bikeId uint256, _renter address, _onTime bool)
func (_Bike *BikeFilterer) FilterBikeReturned(opts *bind.FilterOpts) (*BikeBikeReturnedIterator, error) {

	logs, sub, err := _Bike.contract.FilterLogs(opts, "bikeReturned")
	if err != nil {
		return nil, err
	}
	return &BikeBikeReturnedIterator{contract: _Bike.contract, event: "bikeReturned", logs: logs, sub: sub}, nil
}

// WatchBikeReturned is a free log subscription operation binding the contract event 0x6cdb62cd573c17558f2e746f8dba25f431ab46e3a6e078aac047c92f2a1f92d9.
//
// Solidity: event bikeReturned(_bikeId uint256, _renter address, _onTime bool)
func (_Bike *BikeFilterer) WatchBikeReturned(opts *bind.WatchOpts, sink chan<- *BikeBikeReturned) (event.Subscription, error) {

	logs, sub, err := _Bike.contract.WatchLogs(opts, "bikeReturned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeBikeReturned)
				if err := _Bike.contract.UnpackLog(event, "bikeReturned", log); err != nil {
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

// BikeBikeTransferredIterator is returned from FilterBikeTransferred and is used to iterate over the raw logs and unpacked data for BikeTransferred events raised by the Bike contract.
type BikeBikeTransferredIterator struct {
	Event *BikeBikeTransferred // Event containing the contract specifics and raw log

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
func (it *BikeBikeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeBikeTransferred)
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
		it.Event = new(BikeBikeTransferred)
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
func (it *BikeBikeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeBikeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeBikeTransferred represents a BikeTransferred event raised by the Bike contract.
type BikeBikeTransferred struct {
	BikeId         *big.Int
	PreviousRenter common.Address
	NewRenter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBikeTransferred is a free log retrieval operation binding the contract event 0x4a5e53238d46903ca21c2bb619203a173b1635c625ff4f7c103053c104dac557.
//
// Solidity: event bikeTransferred(_bikeId uint256, _previousRenter address, _newRenter address)
func (_Bike *BikeFilterer) FilterBikeTransferred(opts *bind.FilterOpts) (*BikeBikeTransferredIterator, error) {

	logs, sub, err := _Bike.contract.FilterLogs(opts, "bikeTransferred")
	if err != nil {
		return nil, err
	}
	return &BikeBikeTransferredIterator{contract: _Bike.contract, event: "bikeTransferred", logs: logs, sub: sub}, nil
}

// WatchBikeTransferred is a free log subscription operation binding the contract event 0x4a5e53238d46903ca21c2bb619203a173b1635c625ff4f7c103053c104dac557.
//
// Solidity: event bikeTransferred(_bikeId uint256, _previousRenter address, _newRenter address)
func (_Bike *BikeFilterer) WatchBikeTransferred(opts *bind.WatchOpts, sink chan<- *BikeBikeTransferred) (event.Subscription, error) {

	logs, sub, err := _Bike.contract.WatchLogs(opts, "bikeTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeBikeTransferred)
				if err := _Bike.contract.UnpackLog(event, "bikeTransferred", log); err != nil {
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

// BikeNewCostIterator is returned from FilterNewCost and is used to iterate over the raw logs and unpacked data for NewCost events raised by the Bike contract.
type BikeNewCostIterator struct {
	Event *BikeNewCost // Event containing the contract specifics and raw log

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
func (it *BikeNewCostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeNewCost)
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
		it.Event = new(BikeNewCost)
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
func (it *BikeNewCostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeNewCostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeNewCost represents a NewCost event raised by the Bike contract.
type BikeNewCost struct {
	HourlyCost *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewCost is a free log retrieval operation binding the contract event 0x400839ad27de876fb868336736140fa3adfe86e36fb08dd14b269372ff2bbeff.
//
// Solidity: event newCost(_hourlyCost uint256)
func (_Bike *BikeFilterer) FilterNewCost(opts *bind.FilterOpts) (*BikeNewCostIterator, error) {

	logs, sub, err := _Bike.contract.FilterLogs(opts, "newCost")
	if err != nil {
		return nil, err
	}
	return &BikeNewCostIterator{contract: _Bike.contract, event: "newCost", logs: logs, sub: sub}, nil
}

// WatchNewCost is a free log subscription operation binding the contract event 0x400839ad27de876fb868336736140fa3adfe86e36fb08dd14b269372ff2bbeff.
//
// Solidity: event newCost(_hourlyCost uint256)
func (_Bike *BikeFilterer) WatchNewCost(opts *bind.WatchOpts, sink chan<- *BikeNewCost) (event.Subscription, error) {

	logs, sub, err := _Bike.contract.WatchLogs(opts, "newCost")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeNewCost)
				if err := _Bike.contract.UnpackLog(event, "newCost", log); err != nil {
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
