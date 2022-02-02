// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gennetsepio

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

// GennetsepioMetaData contains all meta data concerning the Gennetsepio contract.
var GennetsepioMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domainAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"siteURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"siteType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"siteTag\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"siteSafety\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"ReviewCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerOrApproved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ReviewDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerOrApproved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldInfoHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newInfoHash\",\"type\":\"string\"}],\"name\":\"ReviewUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETSEPIO_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETSEPIO_MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETSEPIO_VOTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Reviews\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteTag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteSafety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"infoHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteTag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteSafety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"createReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteTag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteSafety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"delegateReviewCreation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"deleteReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"readMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newInfoHash\",\"type\":\"string\"}],\"name\":\"updateReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GennetsepioABI is the input ABI used to generate the binding from.
// Deprecated: Use GennetsepioMetaData.ABI instead.
var GennetsepioABI = GennetsepioMetaData.ABI

// Gennetsepio is an auto generated Go binding around an Ethereum contract.
type Gennetsepio struct {
	GennetsepioCaller     // Read-only binding to the contract
	GennetsepioTransactor // Write-only binding to the contract
	GennetsepioFilterer   // Log filterer for contract events
}

// GennetsepioCaller is an auto generated read-only Go binding around an Ethereum contract.
type GennetsepioCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GennetsepioTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GennetsepioTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GennetsepioFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GennetsepioFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GennetsepioSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GennetsepioSession struct {
	Contract     *Gennetsepio      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GennetsepioCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GennetsepioCallerSession struct {
	Contract *GennetsepioCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GennetsepioTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GennetsepioTransactorSession struct {
	Contract     *GennetsepioTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GennetsepioRaw is an auto generated low-level Go binding around an Ethereum contract.
type GennetsepioRaw struct {
	Contract *Gennetsepio // Generic contract binding to access the raw methods on
}

// GennetsepioCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GennetsepioCallerRaw struct {
	Contract *GennetsepioCaller // Generic read-only contract binding to access the raw methods on
}

// GennetsepioTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GennetsepioTransactorRaw struct {
	Contract *GennetsepioTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGennetsepio creates a new instance of Gennetsepio, bound to a specific deployed contract.
func NewGennetsepio(address common.Address, backend bind.ContractBackend) (*Gennetsepio, error) {
	contract, err := bindGennetsepio(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gennetsepio{GennetsepioCaller: GennetsepioCaller{contract: contract}, GennetsepioTransactor: GennetsepioTransactor{contract: contract}, GennetsepioFilterer: GennetsepioFilterer{contract: contract}}, nil
}

// NewGennetsepioCaller creates a new read-only instance of Gennetsepio, bound to a specific deployed contract.
func NewGennetsepioCaller(address common.Address, caller bind.ContractCaller) (*GennetsepioCaller, error) {
	contract, err := bindGennetsepio(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GennetsepioCaller{contract: contract}, nil
}

// NewGennetsepioTransactor creates a new write-only instance of Gennetsepio, bound to a specific deployed contract.
func NewGennetsepioTransactor(address common.Address, transactor bind.ContractTransactor) (*GennetsepioTransactor, error) {
	contract, err := bindGennetsepio(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GennetsepioTransactor{contract: contract}, nil
}

// NewGennetsepioFilterer creates a new log filterer instance of Gennetsepio, bound to a specific deployed contract.
func NewGennetsepioFilterer(address common.Address, filterer bind.ContractFilterer) (*GennetsepioFilterer, error) {
	contract, err := bindGennetsepio(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GennetsepioFilterer{contract: contract}, nil
}

// bindGennetsepio binds a generic wrapper to an already deployed contract.
func bindGennetsepio(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GennetsepioABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gennetsepio *GennetsepioRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gennetsepio.Contract.GennetsepioCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gennetsepio *GennetsepioRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gennetsepio.Contract.GennetsepioTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gennetsepio *GennetsepioRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gennetsepio.Contract.GennetsepioTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gennetsepio *GennetsepioCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gennetsepio.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gennetsepio *GennetsepioTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gennetsepio.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gennetsepio *GennetsepioTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gennetsepio.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Gennetsepio.Contract.DEFAULTADMINROLE(&_Gennetsepio.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Gennetsepio.Contract.DEFAULTADMINROLE(&_Gennetsepio.CallOpts)
}

// NETSEPIOADMINROLE is a free data retrieval call binding the contract method 0x2297017f.
//
// Solidity: function NETSEPIO_ADMIN_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioCaller) NETSEPIOADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "NETSEPIO_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NETSEPIOADMINROLE is a free data retrieval call binding the contract method 0x2297017f.
//
// Solidity: function NETSEPIO_ADMIN_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioSession) NETSEPIOADMINROLE() ([32]byte, error) {
	return _Gennetsepio.Contract.NETSEPIOADMINROLE(&_Gennetsepio.CallOpts)
}

// NETSEPIOADMINROLE is a free data retrieval call binding the contract method 0x2297017f.
//
// Solidity: function NETSEPIO_ADMIN_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioCallerSession) NETSEPIOADMINROLE() ([32]byte, error) {
	return _Gennetsepio.Contract.NETSEPIOADMINROLE(&_Gennetsepio.CallOpts)
}

// NETSEPIOMODERATORROLE is a free data retrieval call binding the contract method 0x6b1f3b5e.
//
// Solidity: function NETSEPIO_MODERATOR_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioCaller) NETSEPIOMODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "NETSEPIO_MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NETSEPIOMODERATORROLE is a free data retrieval call binding the contract method 0x6b1f3b5e.
//
// Solidity: function NETSEPIO_MODERATOR_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioSession) NETSEPIOMODERATORROLE() ([32]byte, error) {
	return _Gennetsepio.Contract.NETSEPIOMODERATORROLE(&_Gennetsepio.CallOpts)
}

// NETSEPIOMODERATORROLE is a free data retrieval call binding the contract method 0x6b1f3b5e.
//
// Solidity: function NETSEPIO_MODERATOR_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioCallerSession) NETSEPIOMODERATORROLE() ([32]byte, error) {
	return _Gennetsepio.Contract.NETSEPIOMODERATORROLE(&_Gennetsepio.CallOpts)
}

// NETSEPIOVOTERROLE is a free data retrieval call binding the contract method 0x30879cb7.
//
// Solidity: function NETSEPIO_VOTER_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioCaller) NETSEPIOVOTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "NETSEPIO_VOTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NETSEPIOVOTERROLE is a free data retrieval call binding the contract method 0x30879cb7.
//
// Solidity: function NETSEPIO_VOTER_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioSession) NETSEPIOVOTERROLE() ([32]byte, error) {
	return _Gennetsepio.Contract.NETSEPIOVOTERROLE(&_Gennetsepio.CallOpts)
}

// NETSEPIOVOTERROLE is a free data retrieval call binding the contract method 0x30879cb7.
//
// Solidity: function NETSEPIO_VOTER_ROLE() view returns(bytes32)
func (_Gennetsepio *GennetsepioCallerSession) NETSEPIOVOTERROLE() ([32]byte, error) {
	return _Gennetsepio.Contract.NETSEPIOVOTERROLE(&_Gennetsepio.CallOpts)
}

// Reviews is a free data retrieval call binding the contract method 0x113ea69f.
//
// Solidity: function Reviews(uint256 ) view returns(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string infoHash)
func (_Gennetsepio *GennetsepioCaller) Reviews(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Category      string
	DomainAddress string
	SiteURL       string
	SiteType      string
	SiteTag       string
	SiteSafety    string
	InfoHash      string
}, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "Reviews", arg0)

	outstruct := new(struct {
		Category      string
		DomainAddress string
		SiteURL       string
		SiteType      string
		SiteTag       string
		SiteSafety    string
		InfoHash      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Category = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DomainAddress = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.SiteURL = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.SiteType = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.SiteTag = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.SiteSafety = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.InfoHash = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// Reviews is a free data retrieval call binding the contract method 0x113ea69f.
//
// Solidity: function Reviews(uint256 ) view returns(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string infoHash)
func (_Gennetsepio *GennetsepioSession) Reviews(arg0 *big.Int) (struct {
	Category      string
	DomainAddress string
	SiteURL       string
	SiteType      string
	SiteTag       string
	SiteSafety    string
	InfoHash      string
}, error) {
	return _Gennetsepio.Contract.Reviews(&_Gennetsepio.CallOpts, arg0)
}

// Reviews is a free data retrieval call binding the contract method 0x113ea69f.
//
// Solidity: function Reviews(uint256 ) view returns(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string infoHash)
func (_Gennetsepio *GennetsepioCallerSession) Reviews(arg0 *big.Int) (struct {
	Category      string
	DomainAddress string
	SiteURL       string
	SiteType      string
	SiteTag       string
	SiteSafety    string
	InfoHash      string
}, error) {
	return _Gennetsepio.Contract.Reviews(&_Gennetsepio.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Gennetsepio *GennetsepioCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Gennetsepio *GennetsepioSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Gennetsepio.Contract.BalanceOf(&_Gennetsepio.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Gennetsepio *GennetsepioCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Gennetsepio.Contract.BalanceOf(&_Gennetsepio.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Gennetsepio *GennetsepioCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Gennetsepio *GennetsepioSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Gennetsepio.Contract.GetApproved(&_Gennetsepio.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Gennetsepio *GennetsepioCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Gennetsepio.Contract.GetApproved(&_Gennetsepio.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Gennetsepio *GennetsepioCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Gennetsepio *GennetsepioSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Gennetsepio.Contract.GetRoleAdmin(&_Gennetsepio.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Gennetsepio *GennetsepioCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Gennetsepio.Contract.GetRoleAdmin(&_Gennetsepio.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Gennetsepio *GennetsepioCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Gennetsepio *GennetsepioSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Gennetsepio.Contract.GetRoleMember(&_Gennetsepio.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Gennetsepio *GennetsepioCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Gennetsepio.Contract.GetRoleMember(&_Gennetsepio.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Gennetsepio *GennetsepioCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Gennetsepio *GennetsepioSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Gennetsepio.Contract.GetRoleMemberCount(&_Gennetsepio.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Gennetsepio *GennetsepioCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Gennetsepio.Contract.GetRoleMemberCount(&_Gennetsepio.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Gennetsepio *GennetsepioCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Gennetsepio *GennetsepioSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Gennetsepio.Contract.HasRole(&_Gennetsepio.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Gennetsepio *GennetsepioCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Gennetsepio.Contract.HasRole(&_Gennetsepio.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Gennetsepio *GennetsepioCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Gennetsepio *GennetsepioSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Gennetsepio.Contract.IsApprovedForAll(&_Gennetsepio.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Gennetsepio *GennetsepioCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Gennetsepio.Contract.IsApprovedForAll(&_Gennetsepio.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gennetsepio *GennetsepioCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gennetsepio *GennetsepioSession) Name() (string, error) {
	return _Gennetsepio.Contract.Name(&_Gennetsepio.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gennetsepio *GennetsepioCallerSession) Name() (string, error) {
	return _Gennetsepio.Contract.Name(&_Gennetsepio.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Gennetsepio *GennetsepioCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Gennetsepio *GennetsepioSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Gennetsepio.Contract.OwnerOf(&_Gennetsepio.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Gennetsepio *GennetsepioCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Gennetsepio.Contract.OwnerOf(&_Gennetsepio.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gennetsepio *GennetsepioCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gennetsepio *GennetsepioSession) Paused() (bool, error) {
	return _Gennetsepio.Contract.Paused(&_Gennetsepio.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gennetsepio *GennetsepioCallerSession) Paused() (bool, error) {
	return _Gennetsepio.Contract.Paused(&_Gennetsepio.CallOpts)
}

// ReadMetadata is a free data retrieval call binding the contract method 0x48960a4f.
//
// Solidity: function readMetadata(uint256 tokenId) view returns(string)
func (_Gennetsepio *GennetsepioCaller) ReadMetadata(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "readMetadata", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadMetadata is a free data retrieval call binding the contract method 0x48960a4f.
//
// Solidity: function readMetadata(uint256 tokenId) view returns(string)
func (_Gennetsepio *GennetsepioSession) ReadMetadata(tokenId *big.Int) (string, error) {
	return _Gennetsepio.Contract.ReadMetadata(&_Gennetsepio.CallOpts, tokenId)
}

// ReadMetadata is a free data retrieval call binding the contract method 0x48960a4f.
//
// Solidity: function readMetadata(uint256 tokenId) view returns(string)
func (_Gennetsepio *GennetsepioCallerSession) ReadMetadata(tokenId *big.Int) (string, error) {
	return _Gennetsepio.Contract.ReadMetadata(&_Gennetsepio.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Gennetsepio *GennetsepioCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Gennetsepio *GennetsepioSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Gennetsepio.Contract.SupportsInterface(&_Gennetsepio.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Gennetsepio *GennetsepioCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Gennetsepio.Contract.SupportsInterface(&_Gennetsepio.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gennetsepio *GennetsepioCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gennetsepio *GennetsepioSession) Symbol() (string, error) {
	return _Gennetsepio.Contract.Symbol(&_Gennetsepio.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gennetsepio *GennetsepioCallerSession) Symbol() (string, error) {
	return _Gennetsepio.Contract.Symbol(&_Gennetsepio.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Gennetsepio *GennetsepioCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Gennetsepio *GennetsepioSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Gennetsepio.Contract.TokenByIndex(&_Gennetsepio.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Gennetsepio *GennetsepioCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Gennetsepio.Contract.TokenByIndex(&_Gennetsepio.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Gennetsepio *GennetsepioCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Gennetsepio *GennetsepioSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Gennetsepio.Contract.TokenOfOwnerByIndex(&_Gennetsepio.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Gennetsepio *GennetsepioCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Gennetsepio.Contract.TokenOfOwnerByIndex(&_Gennetsepio.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Gennetsepio *GennetsepioCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Gennetsepio *GennetsepioSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Gennetsepio.Contract.TokenURI(&_Gennetsepio.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Gennetsepio *GennetsepioCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Gennetsepio.Contract.TokenURI(&_Gennetsepio.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gennetsepio *GennetsepioCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gennetsepio.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gennetsepio *GennetsepioSession) TotalSupply() (*big.Int, error) {
	return _Gennetsepio.Contract.TotalSupply(&_Gennetsepio.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gennetsepio *GennetsepioCallerSession) TotalSupply() (*big.Int, error) {
	return _Gennetsepio.Contract.TotalSupply(&_Gennetsepio.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.Contract.Approve(&_Gennetsepio.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.Contract.Approve(&_Gennetsepio.TransactOpts, to, tokenId)
}

// CreateReview is a paid mutator transaction binding the contract method 0x6297fac1.
//
// Solidity: function createReview(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI) returns()
func (_Gennetsepio *GennetsepioTransactor) CreateReview(opts *bind.TransactOpts, category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "createReview", category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI)
}

// CreateReview is a paid mutator transaction binding the contract method 0x6297fac1.
//
// Solidity: function createReview(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI) returns()
func (_Gennetsepio *GennetsepioSession) CreateReview(category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string) (*types.Transaction, error) {
	return _Gennetsepio.Contract.CreateReview(&_Gennetsepio.TransactOpts, category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI)
}

// CreateReview is a paid mutator transaction binding the contract method 0x6297fac1.
//
// Solidity: function createReview(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI) returns()
func (_Gennetsepio *GennetsepioTransactorSession) CreateReview(category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string) (*types.Transaction, error) {
	return _Gennetsepio.Contract.CreateReview(&_Gennetsepio.TransactOpts, category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI)
}

// DelegateReviewCreation is a paid mutator transaction binding the contract method 0xb1c58bba.
//
// Solidity: function delegateReviewCreation(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI, address voter) returns()
func (_Gennetsepio *GennetsepioTransactor) DelegateReviewCreation(opts *bind.TransactOpts, category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string, voter common.Address) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "delegateReviewCreation", category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI, voter)
}

// DelegateReviewCreation is a paid mutator transaction binding the contract method 0xb1c58bba.
//
// Solidity: function delegateReviewCreation(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI, address voter) returns()
func (_Gennetsepio *GennetsepioSession) DelegateReviewCreation(category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string, voter common.Address) (*types.Transaction, error) {
	return _Gennetsepio.Contract.DelegateReviewCreation(&_Gennetsepio.TransactOpts, category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI, voter)
}

// DelegateReviewCreation is a paid mutator transaction binding the contract method 0xb1c58bba.
//
// Solidity: function delegateReviewCreation(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI, address voter) returns()
func (_Gennetsepio *GennetsepioTransactorSession) DelegateReviewCreation(category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string, voter common.Address) (*types.Transaction, error) {
	return _Gennetsepio.Contract.DelegateReviewCreation(&_Gennetsepio.TransactOpts, category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI, voter)
}

// DeleteReview is a paid mutator transaction binding the contract method 0xd71e4890.
//
// Solidity: function deleteReview(uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioTransactor) DeleteReview(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "deleteReview", tokenId)
}

// DeleteReview is a paid mutator transaction binding the contract method 0xd71e4890.
//
// Solidity: function deleteReview(uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioSession) DeleteReview(tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.Contract.DeleteReview(&_Gennetsepio.TransactOpts, tokenId)
}

// DeleteReview is a paid mutator transaction binding the contract method 0xd71e4890.
//
// Solidity: function deleteReview(uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioTransactorSession) DeleteReview(tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.Contract.DeleteReview(&_Gennetsepio.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.Contract.GrantRole(&_Gennetsepio.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.Contract.GrantRole(&_Gennetsepio.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gennetsepio *GennetsepioTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gennetsepio *GennetsepioSession) Pause() (*types.Transaction, error) {
	return _Gennetsepio.Contract.Pause(&_Gennetsepio.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gennetsepio *GennetsepioTransactorSession) Pause() (*types.Transaction, error) {
	return _Gennetsepio.Contract.Pause(&_Gennetsepio.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.Contract.RenounceRole(&_Gennetsepio.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.Contract.RenounceRole(&_Gennetsepio.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.Contract.RevokeRole(&_Gennetsepio.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Gennetsepio *GennetsepioTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gennetsepio.Contract.RevokeRole(&_Gennetsepio.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.Contract.SafeTransferFrom(&_Gennetsepio.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.Contract.SafeTransferFrom(&_Gennetsepio.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Gennetsepio *GennetsepioTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Gennetsepio *GennetsepioSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Gennetsepio.Contract.SafeTransferFrom0(&_Gennetsepio.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Gennetsepio *GennetsepioTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Gennetsepio.Contract.SafeTransferFrom0(&_Gennetsepio.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Gennetsepio *GennetsepioTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Gennetsepio *GennetsepioSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Gennetsepio.Contract.SetApprovalForAll(&_Gennetsepio.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Gennetsepio *GennetsepioTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Gennetsepio.Contract.SetApprovalForAll(&_Gennetsepio.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.Contract.TransferFrom(&_Gennetsepio.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Gennetsepio *GennetsepioTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Gennetsepio.Contract.TransferFrom(&_Gennetsepio.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gennetsepio *GennetsepioTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gennetsepio *GennetsepioSession) Unpause() (*types.Transaction, error) {
	return _Gennetsepio.Contract.Unpause(&_Gennetsepio.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gennetsepio *GennetsepioTransactorSession) Unpause() (*types.Transaction, error) {
	return _Gennetsepio.Contract.Unpause(&_Gennetsepio.TransactOpts)
}

// UpdateReview is a paid mutator transaction binding the contract method 0x66b58c03.
//
// Solidity: function updateReview(uint256 tokenId, string newInfoHash) returns()
func (_Gennetsepio *GennetsepioTransactor) UpdateReview(opts *bind.TransactOpts, tokenId *big.Int, newInfoHash string) (*types.Transaction, error) {
	return _Gennetsepio.contract.Transact(opts, "updateReview", tokenId, newInfoHash)
}

// UpdateReview is a paid mutator transaction binding the contract method 0x66b58c03.
//
// Solidity: function updateReview(uint256 tokenId, string newInfoHash) returns()
func (_Gennetsepio *GennetsepioSession) UpdateReview(tokenId *big.Int, newInfoHash string) (*types.Transaction, error) {
	return _Gennetsepio.Contract.UpdateReview(&_Gennetsepio.TransactOpts, tokenId, newInfoHash)
}

// UpdateReview is a paid mutator transaction binding the contract method 0x66b58c03.
//
// Solidity: function updateReview(uint256 tokenId, string newInfoHash) returns()
func (_Gennetsepio *GennetsepioTransactorSession) UpdateReview(tokenId *big.Int, newInfoHash string) (*types.Transaction, error) {
	return _Gennetsepio.Contract.UpdateReview(&_Gennetsepio.TransactOpts, tokenId, newInfoHash)
}

// GennetsepioApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Gennetsepio contract.
type GennetsepioApprovalIterator struct {
	Event *GennetsepioApproval // Event containing the contract specifics and raw log

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
func (it *GennetsepioApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioApproval)
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
		it.Event = new(GennetsepioApproval)
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
func (it *GennetsepioApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioApproval represents a Approval event raised by the Gennetsepio contract.
type GennetsepioApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*GennetsepioApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioApprovalIterator{contract: _Gennetsepio.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GennetsepioApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioApproval)
				if err := _Gennetsepio.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) ParseApproval(log types.Log) (*GennetsepioApproval, error) {
	event := new(GennetsepioApproval)
	if err := _Gennetsepio.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Gennetsepio contract.
type GennetsepioApprovalForAllIterator struct {
	Event *GennetsepioApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GennetsepioApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioApprovalForAll)
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
		it.Event = new(GennetsepioApprovalForAll)
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
func (it *GennetsepioApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioApprovalForAll represents a ApprovalForAll event raised by the Gennetsepio contract.
type GennetsepioApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Gennetsepio *GennetsepioFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*GennetsepioApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioApprovalForAllIterator{contract: _Gennetsepio.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Gennetsepio *GennetsepioFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GennetsepioApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioApprovalForAll)
				if err := _Gennetsepio.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Gennetsepio *GennetsepioFilterer) ParseApprovalForAll(log types.Log) (*GennetsepioApprovalForAll, error) {
	event := new(GennetsepioApprovalForAll)
	if err := _Gennetsepio.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Gennetsepio contract.
type GennetsepioPausedIterator struct {
	Event *GennetsepioPaused // Event containing the contract specifics and raw log

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
func (it *GennetsepioPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioPaused)
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
		it.Event = new(GennetsepioPaused)
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
func (it *GennetsepioPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioPaused represents a Paused event raised by the Gennetsepio contract.
type GennetsepioPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gennetsepio *GennetsepioFilterer) FilterPaused(opts *bind.FilterOpts) (*GennetsepioPausedIterator, error) {

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GennetsepioPausedIterator{contract: _Gennetsepio.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gennetsepio *GennetsepioFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GennetsepioPaused) (event.Subscription, error) {

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioPaused)
				if err := _Gennetsepio.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gennetsepio *GennetsepioFilterer) ParsePaused(log types.Log) (*GennetsepioPaused, error) {
	event := new(GennetsepioPaused)
	if err := _Gennetsepio.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioReviewCreatedIterator is returned from FilterReviewCreated and is used to iterate over the raw logs and unpacked data for ReviewCreated events raised by the Gennetsepio contract.
type GennetsepioReviewCreatedIterator struct {
	Event *GennetsepioReviewCreated // Event containing the contract specifics and raw log

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
func (it *GennetsepioReviewCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioReviewCreated)
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
		it.Event = new(GennetsepioReviewCreated)
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
func (it *GennetsepioReviewCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioReviewCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioReviewCreated represents a ReviewCreated event raised by the Gennetsepio contract.
type GennetsepioReviewCreated struct {
	Receiver      common.Address
	TokenId       *big.Int
	Category      string
	DomainAddress string
	SiteURL       string
	SiteType      string
	SiteTag       string
	SiteSafety    string
	MetadataURI   string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReviewCreated is a free log retrieval operation binding the contract event 0xbf3aecac43badf8f9faced29e91bd3c9e82d11c848f43cf088e7c8d006a8a9fd.
//
// Solidity: event ReviewCreated(address indexed receiver, uint256 indexed tokenId, string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI)
func (_Gennetsepio *GennetsepioFilterer) FilterReviewCreated(opts *bind.FilterOpts, receiver []common.Address, tokenId []*big.Int) (*GennetsepioReviewCreatedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "ReviewCreated", receiverRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioReviewCreatedIterator{contract: _Gennetsepio.contract, event: "ReviewCreated", logs: logs, sub: sub}, nil
}

// WatchReviewCreated is a free log subscription operation binding the contract event 0xbf3aecac43badf8f9faced29e91bd3c9e82d11c848f43cf088e7c8d006a8a9fd.
//
// Solidity: event ReviewCreated(address indexed receiver, uint256 indexed tokenId, string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI)
func (_Gennetsepio *GennetsepioFilterer) WatchReviewCreated(opts *bind.WatchOpts, sink chan<- *GennetsepioReviewCreated, receiver []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "ReviewCreated", receiverRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioReviewCreated)
				if err := _Gennetsepio.contract.UnpackLog(event, "ReviewCreated", log); err != nil {
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

// ParseReviewCreated is a log parse operation binding the contract event 0xbf3aecac43badf8f9faced29e91bd3c9e82d11c848f43cf088e7c8d006a8a9fd.
//
// Solidity: event ReviewCreated(address indexed receiver, uint256 indexed tokenId, string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI)
func (_Gennetsepio *GennetsepioFilterer) ParseReviewCreated(log types.Log) (*GennetsepioReviewCreated, error) {
	event := new(GennetsepioReviewCreated)
	if err := _Gennetsepio.contract.UnpackLog(event, "ReviewCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioReviewDeletedIterator is returned from FilterReviewDeleted and is used to iterate over the raw logs and unpacked data for ReviewDeleted events raised by the Gennetsepio contract.
type GennetsepioReviewDeletedIterator struct {
	Event *GennetsepioReviewDeleted // Event containing the contract specifics and raw log

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
func (it *GennetsepioReviewDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioReviewDeleted)
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
		it.Event = new(GennetsepioReviewDeleted)
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
func (it *GennetsepioReviewDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioReviewDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioReviewDeleted represents a ReviewDeleted event raised by the Gennetsepio contract.
type GennetsepioReviewDeleted struct {
	OwnerOrApproved common.Address
	TokenId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReviewDeleted is a free log retrieval operation binding the contract event 0xdff6a326b030476d6cc240c281b441ec1be2bb84fd9c09fdeb176ba805cbd624.
//
// Solidity: event ReviewDeleted(address indexed ownerOrApproved, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) FilterReviewDeleted(opts *bind.FilterOpts, ownerOrApproved []common.Address, tokenId []*big.Int) (*GennetsepioReviewDeletedIterator, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "ReviewDeleted", ownerOrApprovedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioReviewDeletedIterator{contract: _Gennetsepio.contract, event: "ReviewDeleted", logs: logs, sub: sub}, nil
}

// WatchReviewDeleted is a free log subscription operation binding the contract event 0xdff6a326b030476d6cc240c281b441ec1be2bb84fd9c09fdeb176ba805cbd624.
//
// Solidity: event ReviewDeleted(address indexed ownerOrApproved, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) WatchReviewDeleted(opts *bind.WatchOpts, sink chan<- *GennetsepioReviewDeleted, ownerOrApproved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "ReviewDeleted", ownerOrApprovedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioReviewDeleted)
				if err := _Gennetsepio.contract.UnpackLog(event, "ReviewDeleted", log); err != nil {
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

// ParseReviewDeleted is a log parse operation binding the contract event 0xdff6a326b030476d6cc240c281b441ec1be2bb84fd9c09fdeb176ba805cbd624.
//
// Solidity: event ReviewDeleted(address indexed ownerOrApproved, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) ParseReviewDeleted(log types.Log) (*GennetsepioReviewDeleted, error) {
	event := new(GennetsepioReviewDeleted)
	if err := _Gennetsepio.contract.UnpackLog(event, "ReviewDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioReviewUpdatedIterator is returned from FilterReviewUpdated and is used to iterate over the raw logs and unpacked data for ReviewUpdated events raised by the Gennetsepio contract.
type GennetsepioReviewUpdatedIterator struct {
	Event *GennetsepioReviewUpdated // Event containing the contract specifics and raw log

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
func (it *GennetsepioReviewUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioReviewUpdated)
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
		it.Event = new(GennetsepioReviewUpdated)
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
func (it *GennetsepioReviewUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioReviewUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioReviewUpdated represents a ReviewUpdated event raised by the Gennetsepio contract.
type GennetsepioReviewUpdated struct {
	OwnerOrApproved common.Address
	TokenId         *big.Int
	OldInfoHash     string
	NewInfoHash     string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReviewUpdated is a free log retrieval operation binding the contract event 0x0bf37ab2e7f0b608d2e483599c9fa0d6b6cb52a7380f7233287cef777d4d7920.
//
// Solidity: event ReviewUpdated(address indexed ownerOrApproved, uint256 indexed tokenId, string oldInfoHash, string newInfoHash)
func (_Gennetsepio *GennetsepioFilterer) FilterReviewUpdated(opts *bind.FilterOpts, ownerOrApproved []common.Address, tokenId []*big.Int) (*GennetsepioReviewUpdatedIterator, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "ReviewUpdated", ownerOrApprovedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioReviewUpdatedIterator{contract: _Gennetsepio.contract, event: "ReviewUpdated", logs: logs, sub: sub}, nil
}

// WatchReviewUpdated is a free log subscription operation binding the contract event 0x0bf37ab2e7f0b608d2e483599c9fa0d6b6cb52a7380f7233287cef777d4d7920.
//
// Solidity: event ReviewUpdated(address indexed ownerOrApproved, uint256 indexed tokenId, string oldInfoHash, string newInfoHash)
func (_Gennetsepio *GennetsepioFilterer) WatchReviewUpdated(opts *bind.WatchOpts, sink chan<- *GennetsepioReviewUpdated, ownerOrApproved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "ReviewUpdated", ownerOrApprovedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioReviewUpdated)
				if err := _Gennetsepio.contract.UnpackLog(event, "ReviewUpdated", log); err != nil {
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

// ParseReviewUpdated is a log parse operation binding the contract event 0x0bf37ab2e7f0b608d2e483599c9fa0d6b6cb52a7380f7233287cef777d4d7920.
//
// Solidity: event ReviewUpdated(address indexed ownerOrApproved, uint256 indexed tokenId, string oldInfoHash, string newInfoHash)
func (_Gennetsepio *GennetsepioFilterer) ParseReviewUpdated(log types.Log) (*GennetsepioReviewUpdated, error) {
	event := new(GennetsepioReviewUpdated)
	if err := _Gennetsepio.contract.UnpackLog(event, "ReviewUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Gennetsepio contract.
type GennetsepioRoleAdminChangedIterator struct {
	Event *GennetsepioRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GennetsepioRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioRoleAdminChanged)
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
		it.Event = new(GennetsepioRoleAdminChanged)
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
func (it *GennetsepioRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioRoleAdminChanged represents a RoleAdminChanged event raised by the Gennetsepio contract.
type GennetsepioRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Gennetsepio *GennetsepioFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GennetsepioRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioRoleAdminChangedIterator{contract: _Gennetsepio.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Gennetsepio *GennetsepioFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GennetsepioRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioRoleAdminChanged)
				if err := _Gennetsepio.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Gennetsepio *GennetsepioFilterer) ParseRoleAdminChanged(log types.Log) (*GennetsepioRoleAdminChanged, error) {
	event := new(GennetsepioRoleAdminChanged)
	if err := _Gennetsepio.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Gennetsepio contract.
type GennetsepioRoleGrantedIterator struct {
	Event *GennetsepioRoleGranted // Event containing the contract specifics and raw log

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
func (it *GennetsepioRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioRoleGranted)
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
		it.Event = new(GennetsepioRoleGranted)
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
func (it *GennetsepioRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioRoleGranted represents a RoleGranted event raised by the Gennetsepio contract.
type GennetsepioRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gennetsepio *GennetsepioFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GennetsepioRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioRoleGrantedIterator{contract: _Gennetsepio.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gennetsepio *GennetsepioFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GennetsepioRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioRoleGranted)
				if err := _Gennetsepio.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gennetsepio *GennetsepioFilterer) ParseRoleGranted(log types.Log) (*GennetsepioRoleGranted, error) {
	event := new(GennetsepioRoleGranted)
	if err := _Gennetsepio.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Gennetsepio contract.
type GennetsepioRoleRevokedIterator struct {
	Event *GennetsepioRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GennetsepioRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioRoleRevoked)
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
		it.Event = new(GennetsepioRoleRevoked)
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
func (it *GennetsepioRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioRoleRevoked represents a RoleRevoked event raised by the Gennetsepio contract.
type GennetsepioRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gennetsepio *GennetsepioFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GennetsepioRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioRoleRevokedIterator{contract: _Gennetsepio.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gennetsepio *GennetsepioFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GennetsepioRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioRoleRevoked)
				if err := _Gennetsepio.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gennetsepio *GennetsepioFilterer) ParseRoleRevoked(log types.Log) (*GennetsepioRoleRevoked, error) {
	event := new(GennetsepioRoleRevoked)
	if err := _Gennetsepio.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Gennetsepio contract.
type GennetsepioTransferIterator struct {
	Event *GennetsepioTransfer // Event containing the contract specifics and raw log

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
func (it *GennetsepioTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioTransfer)
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
		it.Event = new(GennetsepioTransfer)
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
func (it *GennetsepioTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioTransfer represents a Transfer event raised by the Gennetsepio contract.
type GennetsepioTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*GennetsepioTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GennetsepioTransferIterator{contract: _Gennetsepio.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GennetsepioTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioTransfer)
				if err := _Gennetsepio.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Gennetsepio *GennetsepioFilterer) ParseTransfer(log types.Log) (*GennetsepioTransfer, error) {
	event := new(GennetsepioTransfer)
	if err := _Gennetsepio.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GennetsepioUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Gennetsepio contract.
type GennetsepioUnpausedIterator struct {
	Event *GennetsepioUnpaused // Event containing the contract specifics and raw log

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
func (it *GennetsepioUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GennetsepioUnpaused)
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
		it.Event = new(GennetsepioUnpaused)
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
func (it *GennetsepioUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GennetsepioUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GennetsepioUnpaused represents a Unpaused event raised by the Gennetsepio contract.
type GennetsepioUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gennetsepio *GennetsepioFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GennetsepioUnpausedIterator, error) {

	logs, sub, err := _Gennetsepio.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GennetsepioUnpausedIterator{contract: _Gennetsepio.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gennetsepio *GennetsepioFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GennetsepioUnpaused) (event.Subscription, error) {

	logs, sub, err := _Gennetsepio.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GennetsepioUnpaused)
				if err := _Gennetsepio.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gennetsepio *GennetsepioFilterer) ParseUnpaused(log types.Log) (*GennetsepioUnpaused, error) {
	event := new(GennetsepioUnpaused)
	if err := _Gennetsepio.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
