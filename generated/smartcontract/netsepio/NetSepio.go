// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package netsepio

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

// NetsepioMetaData contains all meta data concerning the Netsepio contract.
var NetsepioMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domainAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"siteURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"siteType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"siteTag\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"siteSafety\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"ReviewCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerOrApproved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ReviewDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerOrApproved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldInfoHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newInfoHash\",\"type\":\"string\"}],\"name\":\"ReviewUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETSEPIO_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETSEPIO_MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETSEPIO_VOTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Reviews\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteTag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteSafety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"infoHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteTag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteSafety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"createReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domainAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteTag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"siteSafety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"delegateReviewCreation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"deleteReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"readMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newInfoHash\",\"type\":\"string\"}],\"name\":\"updateReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NetsepioABI is the input ABI used to generate the binding from.
// Deprecated: Use NetsepioMetaData.ABI instead.
var NetsepioABI = NetsepioMetaData.ABI

// Netsepio is an auto generated Go binding around an Ethereum contract.
type Netsepio struct {
	NetsepioCaller     // Read-only binding to the contract
	NetsepioTransactor // Write-only binding to the contract
	NetsepioFilterer   // Log filterer for contract events
}

// NetsepioCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetsepioCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetsepioTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetsepioTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetsepioFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetsepioFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetsepioSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetsepioSession struct {
	Contract     *Netsepio         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetsepioCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetsepioCallerSession struct {
	Contract *NetsepioCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NetsepioTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetsepioTransactorSession struct {
	Contract     *NetsepioTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NetsepioRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetsepioRaw struct {
	Contract *Netsepio // Generic contract binding to access the raw methods on
}

// NetsepioCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetsepioCallerRaw struct {
	Contract *NetsepioCaller // Generic read-only contract binding to access the raw methods on
}

// NetsepioTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetsepioTransactorRaw struct {
	Contract *NetsepioTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetsepio creates a new instance of Netsepio, bound to a specific deployed contract.
func NewNetsepio(address common.Address, backend bind.ContractBackend) (*Netsepio, error) {
	contract, err := bindNetsepio(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Netsepio{NetsepioCaller: NetsepioCaller{contract: contract}, NetsepioTransactor: NetsepioTransactor{contract: contract}, NetsepioFilterer: NetsepioFilterer{contract: contract}}, nil
}

// NewNetsepioCaller creates a new read-only instance of Netsepio, bound to a specific deployed contract.
func NewNetsepioCaller(address common.Address, caller bind.ContractCaller) (*NetsepioCaller, error) {
	contract, err := bindNetsepio(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetsepioCaller{contract: contract}, nil
}

// NewNetsepioTransactor creates a new write-only instance of Netsepio, bound to a specific deployed contract.
func NewNetsepioTransactor(address common.Address, transactor bind.ContractTransactor) (*NetsepioTransactor, error) {
	contract, err := bindNetsepio(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetsepioTransactor{contract: contract}, nil
}

// NewNetsepioFilterer creates a new log filterer instance of Netsepio, bound to a specific deployed contract.
func NewNetsepioFilterer(address common.Address, filterer bind.ContractFilterer) (*NetsepioFilterer, error) {
	contract, err := bindNetsepio(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetsepioFilterer{contract: contract}, nil
}

// bindNetsepio binds a generic wrapper to an already deployed contract.
func bindNetsepio(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NetsepioABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Netsepio *NetsepioRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Netsepio.Contract.NetsepioCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Netsepio *NetsepioRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Netsepio.Contract.NetsepioTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Netsepio *NetsepioRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Netsepio.Contract.NetsepioTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Netsepio *NetsepioCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Netsepio.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Netsepio *NetsepioTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Netsepio.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Netsepio *NetsepioTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Netsepio.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Netsepio.Contract.DEFAULTADMINROLE(&_Netsepio.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Netsepio.Contract.DEFAULTADMINROLE(&_Netsepio.CallOpts)
}

// NETSEPIOADMINROLE is a free data retrieval call binding the contract method 0x2297017f.
//
// Solidity: function NETSEPIO_ADMIN_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioCaller) NETSEPIOADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "NETSEPIO_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NETSEPIOADMINROLE is a free data retrieval call binding the contract method 0x2297017f.
//
// Solidity: function NETSEPIO_ADMIN_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioSession) NETSEPIOADMINROLE() ([32]byte, error) {
	return _Netsepio.Contract.NETSEPIOADMINROLE(&_Netsepio.CallOpts)
}

// NETSEPIOADMINROLE is a free data retrieval call binding the contract method 0x2297017f.
//
// Solidity: function NETSEPIO_ADMIN_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioCallerSession) NETSEPIOADMINROLE() ([32]byte, error) {
	return _Netsepio.Contract.NETSEPIOADMINROLE(&_Netsepio.CallOpts)
}

// NETSEPIOMODERATORROLE is a free data retrieval call binding the contract method 0x6b1f3b5e.
//
// Solidity: function NETSEPIO_MODERATOR_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioCaller) NETSEPIOMODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "NETSEPIO_MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NETSEPIOMODERATORROLE is a free data retrieval call binding the contract method 0x6b1f3b5e.
//
// Solidity: function NETSEPIO_MODERATOR_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioSession) NETSEPIOMODERATORROLE() ([32]byte, error) {
	return _Netsepio.Contract.NETSEPIOMODERATORROLE(&_Netsepio.CallOpts)
}

// NETSEPIOMODERATORROLE is a free data retrieval call binding the contract method 0x6b1f3b5e.
//
// Solidity: function NETSEPIO_MODERATOR_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioCallerSession) NETSEPIOMODERATORROLE() ([32]byte, error) {
	return _Netsepio.Contract.NETSEPIOMODERATORROLE(&_Netsepio.CallOpts)
}

// NETSEPIOVOTERROLE is a free data retrieval call binding the contract method 0x30879cb7.
//
// Solidity: function NETSEPIO_VOTER_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioCaller) NETSEPIOVOTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "NETSEPIO_VOTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NETSEPIOVOTERROLE is a free data retrieval call binding the contract method 0x30879cb7.
//
// Solidity: function NETSEPIO_VOTER_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioSession) NETSEPIOVOTERROLE() ([32]byte, error) {
	return _Netsepio.Contract.NETSEPIOVOTERROLE(&_Netsepio.CallOpts)
}

// NETSEPIOVOTERROLE is a free data retrieval call binding the contract method 0x30879cb7.
//
// Solidity: function NETSEPIO_VOTER_ROLE() view returns(bytes32)
func (_Netsepio *NetsepioCallerSession) NETSEPIOVOTERROLE() ([32]byte, error) {
	return _Netsepio.Contract.NETSEPIOVOTERROLE(&_Netsepio.CallOpts)
}

// Reviews is a free data retrieval call binding the contract method 0x113ea69f.
//
// Solidity: function Reviews(uint256 ) view returns(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string infoHash)
func (_Netsepio *NetsepioCaller) Reviews(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Category      string
	DomainAddress string
	SiteURL       string
	SiteType      string
	SiteTag       string
	SiteSafety    string
	InfoHash      string
}, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "Reviews", arg0)

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
func (_Netsepio *NetsepioSession) Reviews(arg0 *big.Int) (struct {
	Category      string
	DomainAddress string
	SiteURL       string
	SiteType      string
	SiteTag       string
	SiteSafety    string
	InfoHash      string
}, error) {
	return _Netsepio.Contract.Reviews(&_Netsepio.CallOpts, arg0)
}

// Reviews is a free data retrieval call binding the contract method 0x113ea69f.
//
// Solidity: function Reviews(uint256 ) view returns(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string infoHash)
func (_Netsepio *NetsepioCallerSession) Reviews(arg0 *big.Int) (struct {
	Category      string
	DomainAddress string
	SiteURL       string
	SiteType      string
	SiteTag       string
	SiteSafety    string
	InfoHash      string
}, error) {
	return _Netsepio.Contract.Reviews(&_Netsepio.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Netsepio *NetsepioCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Netsepio *NetsepioSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Netsepio.Contract.BalanceOf(&_Netsepio.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Netsepio *NetsepioCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Netsepio.Contract.BalanceOf(&_Netsepio.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Netsepio *NetsepioCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Netsepio *NetsepioSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Netsepio.Contract.GetApproved(&_Netsepio.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Netsepio *NetsepioCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Netsepio.Contract.GetApproved(&_Netsepio.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Netsepio *NetsepioCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Netsepio *NetsepioSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Netsepio.Contract.GetRoleAdmin(&_Netsepio.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Netsepio *NetsepioCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Netsepio.Contract.GetRoleAdmin(&_Netsepio.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Netsepio *NetsepioCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Netsepio *NetsepioSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Netsepio.Contract.GetRoleMember(&_Netsepio.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Netsepio *NetsepioCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Netsepio.Contract.GetRoleMember(&_Netsepio.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Netsepio *NetsepioCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Netsepio *NetsepioSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Netsepio.Contract.GetRoleMemberCount(&_Netsepio.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Netsepio *NetsepioCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Netsepio.Contract.GetRoleMemberCount(&_Netsepio.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Netsepio *NetsepioCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Netsepio *NetsepioSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Netsepio.Contract.HasRole(&_Netsepio.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Netsepio *NetsepioCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Netsepio.Contract.HasRole(&_Netsepio.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Netsepio *NetsepioCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Netsepio *NetsepioSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Netsepio.Contract.IsApprovedForAll(&_Netsepio.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Netsepio *NetsepioCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Netsepio.Contract.IsApprovedForAll(&_Netsepio.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Netsepio *NetsepioCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Netsepio *NetsepioSession) Name() (string, error) {
	return _Netsepio.Contract.Name(&_Netsepio.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Netsepio *NetsepioCallerSession) Name() (string, error) {
	return _Netsepio.Contract.Name(&_Netsepio.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Netsepio *NetsepioCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Netsepio *NetsepioSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Netsepio.Contract.OwnerOf(&_Netsepio.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Netsepio *NetsepioCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Netsepio.Contract.OwnerOf(&_Netsepio.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Netsepio *NetsepioCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Netsepio *NetsepioSession) Paused() (bool, error) {
	return _Netsepio.Contract.Paused(&_Netsepio.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Netsepio *NetsepioCallerSession) Paused() (bool, error) {
	return _Netsepio.Contract.Paused(&_Netsepio.CallOpts)
}

// ReadMetadata is a free data retrieval call binding the contract method 0x48960a4f.
//
// Solidity: function readMetadata(uint256 tokenId) view returns(string)
func (_Netsepio *NetsepioCaller) ReadMetadata(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "readMetadata", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadMetadata is a free data retrieval call binding the contract method 0x48960a4f.
//
// Solidity: function readMetadata(uint256 tokenId) view returns(string)
func (_Netsepio *NetsepioSession) ReadMetadata(tokenId *big.Int) (string, error) {
	return _Netsepio.Contract.ReadMetadata(&_Netsepio.CallOpts, tokenId)
}

// ReadMetadata is a free data retrieval call binding the contract method 0x48960a4f.
//
// Solidity: function readMetadata(uint256 tokenId) view returns(string)
func (_Netsepio *NetsepioCallerSession) ReadMetadata(tokenId *big.Int) (string, error) {
	return _Netsepio.Contract.ReadMetadata(&_Netsepio.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Netsepio *NetsepioCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Netsepio *NetsepioSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Netsepio.Contract.SupportsInterface(&_Netsepio.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Netsepio *NetsepioCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Netsepio.Contract.SupportsInterface(&_Netsepio.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Netsepio *NetsepioCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Netsepio *NetsepioSession) Symbol() (string, error) {
	return _Netsepio.Contract.Symbol(&_Netsepio.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Netsepio *NetsepioCallerSession) Symbol() (string, error) {
	return _Netsepio.Contract.Symbol(&_Netsepio.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Netsepio *NetsepioCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Netsepio *NetsepioSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Netsepio.Contract.TokenByIndex(&_Netsepio.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Netsepio *NetsepioCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Netsepio.Contract.TokenByIndex(&_Netsepio.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Netsepio *NetsepioCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Netsepio *NetsepioSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Netsepio.Contract.TokenOfOwnerByIndex(&_Netsepio.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Netsepio *NetsepioCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Netsepio.Contract.TokenOfOwnerByIndex(&_Netsepio.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Netsepio *NetsepioCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Netsepio *NetsepioSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Netsepio.Contract.TokenURI(&_Netsepio.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Netsepio *NetsepioCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Netsepio.Contract.TokenURI(&_Netsepio.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Netsepio *NetsepioCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Netsepio.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Netsepio *NetsepioSession) TotalSupply() (*big.Int, error) {
	return _Netsepio.Contract.TotalSupply(&_Netsepio.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Netsepio *NetsepioCallerSession) TotalSupply() (*big.Int, error) {
	return _Netsepio.Contract.TotalSupply(&_Netsepio.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.Contract.Approve(&_Netsepio.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.Contract.Approve(&_Netsepio.TransactOpts, to, tokenId)
}

// CreateReview is a paid mutator transaction binding the contract method 0x6297fac1.
//
// Solidity: function createReview(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI) returns()
func (_Netsepio *NetsepioTransactor) CreateReview(opts *bind.TransactOpts, category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "createReview", category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI)
}

// CreateReview is a paid mutator transaction binding the contract method 0x6297fac1.
//
// Solidity: function createReview(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI) returns()
func (_Netsepio *NetsepioSession) CreateReview(category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string) (*types.Transaction, error) {
	return _Netsepio.Contract.CreateReview(&_Netsepio.TransactOpts, category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI)
}

// CreateReview is a paid mutator transaction binding the contract method 0x6297fac1.
//
// Solidity: function createReview(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI) returns()
func (_Netsepio *NetsepioTransactorSession) CreateReview(category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string) (*types.Transaction, error) {
	return _Netsepio.Contract.CreateReview(&_Netsepio.TransactOpts, category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI)
}

// DelegateReviewCreation is a paid mutator transaction binding the contract method 0xb1c58bba.
//
// Solidity: function delegateReviewCreation(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI, address voter) returns()
func (_Netsepio *NetsepioTransactor) DelegateReviewCreation(opts *bind.TransactOpts, category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string, voter common.Address) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "delegateReviewCreation", category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI, voter)
}

// DelegateReviewCreation is a paid mutator transaction binding the contract method 0xb1c58bba.
//
// Solidity: function delegateReviewCreation(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI, address voter) returns()
func (_Netsepio *NetsepioSession) DelegateReviewCreation(category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string, voter common.Address) (*types.Transaction, error) {
	return _Netsepio.Contract.DelegateReviewCreation(&_Netsepio.TransactOpts, category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI, voter)
}

// DelegateReviewCreation is a paid mutator transaction binding the contract method 0xb1c58bba.
//
// Solidity: function delegateReviewCreation(string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI, address voter) returns()
func (_Netsepio *NetsepioTransactorSession) DelegateReviewCreation(category string, domainAddress string, siteURL string, siteType string, siteTag string, siteSafety string, metadataURI string, voter common.Address) (*types.Transaction, error) {
	return _Netsepio.Contract.DelegateReviewCreation(&_Netsepio.TransactOpts, category, domainAddress, siteURL, siteType, siteTag, siteSafety, metadataURI, voter)
}

// DeleteReview is a paid mutator transaction binding the contract method 0xd71e4890.
//
// Solidity: function deleteReview(uint256 tokenId) returns()
func (_Netsepio *NetsepioTransactor) DeleteReview(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "deleteReview", tokenId)
}

// DeleteReview is a paid mutator transaction binding the contract method 0xd71e4890.
//
// Solidity: function deleteReview(uint256 tokenId) returns()
func (_Netsepio *NetsepioSession) DeleteReview(tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.Contract.DeleteReview(&_Netsepio.TransactOpts, tokenId)
}

// DeleteReview is a paid mutator transaction binding the contract method 0xd71e4890.
//
// Solidity: function deleteReview(uint256 tokenId) returns()
func (_Netsepio *NetsepioTransactorSession) DeleteReview(tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.Contract.DeleteReview(&_Netsepio.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.Contract.GrantRole(&_Netsepio.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.Contract.GrantRole(&_Netsepio.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Netsepio *NetsepioTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Netsepio *NetsepioSession) Pause() (*types.Transaction, error) {
	return _Netsepio.Contract.Pause(&_Netsepio.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Netsepio *NetsepioTransactorSession) Pause() (*types.Transaction, error) {
	return _Netsepio.Contract.Pause(&_Netsepio.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.Contract.RenounceRole(&_Netsepio.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.Contract.RenounceRole(&_Netsepio.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.Contract.RevokeRole(&_Netsepio.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Netsepio *NetsepioTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Netsepio.Contract.RevokeRole(&_Netsepio.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.Contract.SafeTransferFrom(&_Netsepio.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.Contract.SafeTransferFrom(&_Netsepio.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Netsepio *NetsepioTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Netsepio *NetsepioSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Netsepio.Contract.SafeTransferFrom0(&_Netsepio.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Netsepio *NetsepioTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Netsepio.Contract.SafeTransferFrom0(&_Netsepio.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Netsepio *NetsepioTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Netsepio *NetsepioSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Netsepio.Contract.SetApprovalForAll(&_Netsepio.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Netsepio *NetsepioTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Netsepio.Contract.SetApprovalForAll(&_Netsepio.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.Contract.TransferFrom(&_Netsepio.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Netsepio *NetsepioTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Netsepio.Contract.TransferFrom(&_Netsepio.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Netsepio *NetsepioTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Netsepio *NetsepioSession) Unpause() (*types.Transaction, error) {
	return _Netsepio.Contract.Unpause(&_Netsepio.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Netsepio *NetsepioTransactorSession) Unpause() (*types.Transaction, error) {
	return _Netsepio.Contract.Unpause(&_Netsepio.TransactOpts)
}

// UpdateReview is a paid mutator transaction binding the contract method 0x66b58c03.
//
// Solidity: function updateReview(uint256 tokenId, string newInfoHash) returns()
func (_Netsepio *NetsepioTransactor) UpdateReview(opts *bind.TransactOpts, tokenId *big.Int, newInfoHash string) (*types.Transaction, error) {
	return _Netsepio.contract.Transact(opts, "updateReview", tokenId, newInfoHash)
}

// UpdateReview is a paid mutator transaction binding the contract method 0x66b58c03.
//
// Solidity: function updateReview(uint256 tokenId, string newInfoHash) returns()
func (_Netsepio *NetsepioSession) UpdateReview(tokenId *big.Int, newInfoHash string) (*types.Transaction, error) {
	return _Netsepio.Contract.UpdateReview(&_Netsepio.TransactOpts, tokenId, newInfoHash)
}

// UpdateReview is a paid mutator transaction binding the contract method 0x66b58c03.
//
// Solidity: function updateReview(uint256 tokenId, string newInfoHash) returns()
func (_Netsepio *NetsepioTransactorSession) UpdateReview(tokenId *big.Int, newInfoHash string) (*types.Transaction, error) {
	return _Netsepio.Contract.UpdateReview(&_Netsepio.TransactOpts, tokenId, newInfoHash)
}

// NetsepioApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Netsepio contract.
type NetsepioApprovalIterator struct {
	Event *NetsepioApproval // Event containing the contract specifics and raw log

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
func (it *NetsepioApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioApproval)
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
		it.Event = new(NetsepioApproval)
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
func (it *NetsepioApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioApproval represents a Approval event raised by the Netsepio contract.
type NetsepioApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Netsepio *NetsepioFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NetsepioApprovalIterator, error) {

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

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioApprovalIterator{contract: _Netsepio.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Netsepio *NetsepioFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NetsepioApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioApproval)
				if err := _Netsepio.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseApproval(log types.Log) (*NetsepioApproval, error) {
	event := new(NetsepioApproval)
	if err := _Netsepio.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Netsepio contract.
type NetsepioApprovalForAllIterator struct {
	Event *NetsepioApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NetsepioApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioApprovalForAll)
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
		it.Event = new(NetsepioApprovalForAll)
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
func (it *NetsepioApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioApprovalForAll represents a ApprovalForAll event raised by the Netsepio contract.
type NetsepioApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Netsepio *NetsepioFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NetsepioApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioApprovalForAllIterator{contract: _Netsepio.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Netsepio *NetsepioFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NetsepioApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioApprovalForAll)
				if err := _Netsepio.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseApprovalForAll(log types.Log) (*NetsepioApprovalForAll, error) {
	event := new(NetsepioApprovalForAll)
	if err := _Netsepio.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Netsepio contract.
type NetsepioPausedIterator struct {
	Event *NetsepioPaused // Event containing the contract specifics and raw log

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
func (it *NetsepioPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioPaused)
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
		it.Event = new(NetsepioPaused)
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
func (it *NetsepioPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioPaused represents a Paused event raised by the Netsepio contract.
type NetsepioPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Netsepio *NetsepioFilterer) FilterPaused(opts *bind.FilterOpts) (*NetsepioPausedIterator, error) {

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NetsepioPausedIterator{contract: _Netsepio.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Netsepio *NetsepioFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NetsepioPaused) (event.Subscription, error) {

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioPaused)
				if err := _Netsepio.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParsePaused(log types.Log) (*NetsepioPaused, error) {
	event := new(NetsepioPaused)
	if err := _Netsepio.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioReviewCreatedIterator is returned from FilterReviewCreated and is used to iterate over the raw logs and unpacked data for ReviewCreated events raised by the Netsepio contract.
type NetsepioReviewCreatedIterator struct {
	Event *NetsepioReviewCreated // Event containing the contract specifics and raw log

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
func (it *NetsepioReviewCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioReviewCreated)
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
		it.Event = new(NetsepioReviewCreated)
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
func (it *NetsepioReviewCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioReviewCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioReviewCreated represents a ReviewCreated event raised by the Netsepio contract.
type NetsepioReviewCreated struct {
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
func (_Netsepio *NetsepioFilterer) FilterReviewCreated(opts *bind.FilterOpts, receiver []common.Address, tokenId []*big.Int) (*NetsepioReviewCreatedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "ReviewCreated", receiverRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioReviewCreatedIterator{contract: _Netsepio.contract, event: "ReviewCreated", logs: logs, sub: sub}, nil
}

// WatchReviewCreated is a free log subscription operation binding the contract event 0xbf3aecac43badf8f9faced29e91bd3c9e82d11c848f43cf088e7c8d006a8a9fd.
//
// Solidity: event ReviewCreated(address indexed receiver, uint256 indexed tokenId, string category, string domainAddress, string siteURL, string siteType, string siteTag, string siteSafety, string metadataURI)
func (_Netsepio *NetsepioFilterer) WatchReviewCreated(opts *bind.WatchOpts, sink chan<- *NetsepioReviewCreated, receiver []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "ReviewCreated", receiverRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioReviewCreated)
				if err := _Netsepio.contract.UnpackLog(event, "ReviewCreated", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseReviewCreated(log types.Log) (*NetsepioReviewCreated, error) {
	event := new(NetsepioReviewCreated)
	if err := _Netsepio.contract.UnpackLog(event, "ReviewCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioReviewDeletedIterator is returned from FilterReviewDeleted and is used to iterate over the raw logs and unpacked data for ReviewDeleted events raised by the Netsepio contract.
type NetsepioReviewDeletedIterator struct {
	Event *NetsepioReviewDeleted // Event containing the contract specifics and raw log

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
func (it *NetsepioReviewDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioReviewDeleted)
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
		it.Event = new(NetsepioReviewDeleted)
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
func (it *NetsepioReviewDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioReviewDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioReviewDeleted represents a ReviewDeleted event raised by the Netsepio contract.
type NetsepioReviewDeleted struct {
	OwnerOrApproved common.Address
	TokenId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReviewDeleted is a free log retrieval operation binding the contract event 0xdff6a326b030476d6cc240c281b441ec1be2bb84fd9c09fdeb176ba805cbd624.
//
// Solidity: event ReviewDeleted(address indexed ownerOrApproved, uint256 indexed tokenId)
func (_Netsepio *NetsepioFilterer) FilterReviewDeleted(opts *bind.FilterOpts, ownerOrApproved []common.Address, tokenId []*big.Int) (*NetsepioReviewDeletedIterator, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "ReviewDeleted", ownerOrApprovedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioReviewDeletedIterator{contract: _Netsepio.contract, event: "ReviewDeleted", logs: logs, sub: sub}, nil
}

// WatchReviewDeleted is a free log subscription operation binding the contract event 0xdff6a326b030476d6cc240c281b441ec1be2bb84fd9c09fdeb176ba805cbd624.
//
// Solidity: event ReviewDeleted(address indexed ownerOrApproved, uint256 indexed tokenId)
func (_Netsepio *NetsepioFilterer) WatchReviewDeleted(opts *bind.WatchOpts, sink chan<- *NetsepioReviewDeleted, ownerOrApproved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "ReviewDeleted", ownerOrApprovedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioReviewDeleted)
				if err := _Netsepio.contract.UnpackLog(event, "ReviewDeleted", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseReviewDeleted(log types.Log) (*NetsepioReviewDeleted, error) {
	event := new(NetsepioReviewDeleted)
	if err := _Netsepio.contract.UnpackLog(event, "ReviewDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioReviewUpdatedIterator is returned from FilterReviewUpdated and is used to iterate over the raw logs and unpacked data for ReviewUpdated events raised by the Netsepio contract.
type NetsepioReviewUpdatedIterator struct {
	Event *NetsepioReviewUpdated // Event containing the contract specifics and raw log

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
func (it *NetsepioReviewUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioReviewUpdated)
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
		it.Event = new(NetsepioReviewUpdated)
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
func (it *NetsepioReviewUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioReviewUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioReviewUpdated represents a ReviewUpdated event raised by the Netsepio contract.
type NetsepioReviewUpdated struct {
	OwnerOrApproved common.Address
	TokenId         *big.Int
	OldInfoHash     string
	NewInfoHash     string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReviewUpdated is a free log retrieval operation binding the contract event 0x0bf37ab2e7f0b608d2e483599c9fa0d6b6cb52a7380f7233287cef777d4d7920.
//
// Solidity: event ReviewUpdated(address indexed ownerOrApproved, uint256 indexed tokenId, string oldInfoHash, string newInfoHash)
func (_Netsepio *NetsepioFilterer) FilterReviewUpdated(opts *bind.FilterOpts, ownerOrApproved []common.Address, tokenId []*big.Int) (*NetsepioReviewUpdatedIterator, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "ReviewUpdated", ownerOrApprovedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioReviewUpdatedIterator{contract: _Netsepio.contract, event: "ReviewUpdated", logs: logs, sub: sub}, nil
}

// WatchReviewUpdated is a free log subscription operation binding the contract event 0x0bf37ab2e7f0b608d2e483599c9fa0d6b6cb52a7380f7233287cef777d4d7920.
//
// Solidity: event ReviewUpdated(address indexed ownerOrApproved, uint256 indexed tokenId, string oldInfoHash, string newInfoHash)
func (_Netsepio *NetsepioFilterer) WatchReviewUpdated(opts *bind.WatchOpts, sink chan<- *NetsepioReviewUpdated, ownerOrApproved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerOrApprovedRule []interface{}
	for _, ownerOrApprovedItem := range ownerOrApproved {
		ownerOrApprovedRule = append(ownerOrApprovedRule, ownerOrApprovedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "ReviewUpdated", ownerOrApprovedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioReviewUpdated)
				if err := _Netsepio.contract.UnpackLog(event, "ReviewUpdated", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseReviewUpdated(log types.Log) (*NetsepioReviewUpdated, error) {
	event := new(NetsepioReviewUpdated)
	if err := _Netsepio.contract.UnpackLog(event, "ReviewUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Netsepio contract.
type NetsepioRoleAdminChangedIterator struct {
	Event *NetsepioRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NetsepioRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioRoleAdminChanged)
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
		it.Event = new(NetsepioRoleAdminChanged)
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
func (it *NetsepioRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioRoleAdminChanged represents a RoleAdminChanged event raised by the Netsepio contract.
type NetsepioRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Netsepio *NetsepioFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NetsepioRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioRoleAdminChangedIterator{contract: _Netsepio.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Netsepio *NetsepioFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NetsepioRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioRoleAdminChanged)
				if err := _Netsepio.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseRoleAdminChanged(log types.Log) (*NetsepioRoleAdminChanged, error) {
	event := new(NetsepioRoleAdminChanged)
	if err := _Netsepio.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Netsepio contract.
type NetsepioRoleGrantedIterator struct {
	Event *NetsepioRoleGranted // Event containing the contract specifics and raw log

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
func (it *NetsepioRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioRoleGranted)
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
		it.Event = new(NetsepioRoleGranted)
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
func (it *NetsepioRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioRoleGranted represents a RoleGranted event raised by the Netsepio contract.
type NetsepioRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Netsepio *NetsepioFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NetsepioRoleGrantedIterator, error) {

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

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioRoleGrantedIterator{contract: _Netsepio.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Netsepio *NetsepioFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NetsepioRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioRoleGranted)
				if err := _Netsepio.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseRoleGranted(log types.Log) (*NetsepioRoleGranted, error) {
	event := new(NetsepioRoleGranted)
	if err := _Netsepio.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Netsepio contract.
type NetsepioRoleRevokedIterator struct {
	Event *NetsepioRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NetsepioRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioRoleRevoked)
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
		it.Event = new(NetsepioRoleRevoked)
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
func (it *NetsepioRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioRoleRevoked represents a RoleRevoked event raised by the Netsepio contract.
type NetsepioRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Netsepio *NetsepioFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NetsepioRoleRevokedIterator, error) {

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

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioRoleRevokedIterator{contract: _Netsepio.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Netsepio *NetsepioFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NetsepioRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioRoleRevoked)
				if err := _Netsepio.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseRoleRevoked(log types.Log) (*NetsepioRoleRevoked, error) {
	event := new(NetsepioRoleRevoked)
	if err := _Netsepio.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Netsepio contract.
type NetsepioTransferIterator struct {
	Event *NetsepioTransfer // Event containing the contract specifics and raw log

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
func (it *NetsepioTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioTransfer)
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
		it.Event = new(NetsepioTransfer)
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
func (it *NetsepioTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioTransfer represents a Transfer event raised by the Netsepio contract.
type NetsepioTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Netsepio *NetsepioFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NetsepioTransferIterator, error) {

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

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NetsepioTransferIterator{contract: _Netsepio.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Netsepio *NetsepioFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NetsepioTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioTransfer)
				if err := _Netsepio.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseTransfer(log types.Log) (*NetsepioTransfer, error) {
	event := new(NetsepioTransfer)
	if err := _Netsepio.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetsepioUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Netsepio contract.
type NetsepioUnpausedIterator struct {
	Event *NetsepioUnpaused // Event containing the contract specifics and raw log

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
func (it *NetsepioUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetsepioUnpaused)
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
		it.Event = new(NetsepioUnpaused)
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
func (it *NetsepioUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetsepioUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetsepioUnpaused represents a Unpaused event raised by the Netsepio contract.
type NetsepioUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Netsepio *NetsepioFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NetsepioUnpausedIterator, error) {

	logs, sub, err := _Netsepio.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NetsepioUnpausedIterator{contract: _Netsepio.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Netsepio *NetsepioFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NetsepioUnpaused) (event.Subscription, error) {

	logs, sub, err := _Netsepio.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetsepioUnpaused)
				if err := _Netsepio.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Netsepio *NetsepioFilterer) ParseUnpaused(log types.Log) (*NetsepioUnpaused, error) {
	event := new(NetsepioUnpaused)
	if err := _Netsepio.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
