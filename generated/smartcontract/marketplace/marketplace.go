// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creatify

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

// CreatifyMetaData contains all meta data concerning the Creatify contract.
var CreatifyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_platformFee\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metaDataURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"forSale\",\"type\":\"bool\"}],\"name\":\"MarketItemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"MarketItemRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"MarketItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKETPLACE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"newPlatformFee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"newPayoutAddress\",\"type\":\"address\"}],\"name\":\"changeFeeAndPayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createMarketItem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"createMarketSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idToMarketItem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forSale\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeBasisPoint\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"removeFromSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreatifyABI is the input ABI used to generate the binding from.
// Deprecated: Use CreatifyMetaData.ABI instead.
var CreatifyABI = CreatifyMetaData.ABI

// Creatify is an auto generated Go binding around an Ethereum contract.
type Creatify struct {
	CreatifyCaller     // Read-only binding to the contract
	CreatifyTransactor // Write-only binding to the contract
	CreatifyFilterer   // Log filterer for contract events
}

// CreatifyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreatifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatifyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreatifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatifyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreatifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreatifySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreatifySession struct {
	Contract     *Creatify         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreatifyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreatifyCallerSession struct {
	Contract *CreatifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CreatifyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreatifyTransactorSession struct {
	Contract     *CreatifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CreatifyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreatifyRaw struct {
	Contract *Creatify // Generic contract binding to access the raw methods on
}

// CreatifyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreatifyCallerRaw struct {
	Contract *CreatifyCaller // Generic read-only contract binding to access the raw methods on
}

// CreatifyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreatifyTransactorRaw struct {
	Contract *CreatifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreatify creates a new instance of Creatify, bound to a specific deployed contract.
func NewCreatify(address common.Address, backend bind.ContractBackend) (*Creatify, error) {
	contract, err := bindCreatify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Creatify{CreatifyCaller: CreatifyCaller{contract: contract}, CreatifyTransactor: CreatifyTransactor{contract: contract}, CreatifyFilterer: CreatifyFilterer{contract: contract}}, nil
}

// NewCreatifyCaller creates a new read-only instance of Creatify, bound to a specific deployed contract.
func NewCreatifyCaller(address common.Address, caller bind.ContractCaller) (*CreatifyCaller, error) {
	contract, err := bindCreatify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreatifyCaller{contract: contract}, nil
}

// NewCreatifyTransactor creates a new write-only instance of Creatify, bound to a specific deployed contract.
func NewCreatifyTransactor(address common.Address, transactor bind.ContractTransactor) (*CreatifyTransactor, error) {
	contract, err := bindCreatify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreatifyTransactor{contract: contract}, nil
}

// NewCreatifyFilterer creates a new log filterer instance of Creatify, bound to a specific deployed contract.
func NewCreatifyFilterer(address common.Address, filterer bind.ContractFilterer) (*CreatifyFilterer, error) {
	contract, err := bindCreatify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreatifyFilterer{contract: contract}, nil
}

// bindCreatify binds a generic wrapper to an already deployed contract.
func bindCreatify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreatifyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Creatify *CreatifyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Creatify.Contract.CreatifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Creatify *CreatifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creatify.Contract.CreatifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Creatify *CreatifyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Creatify.Contract.CreatifyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Creatify *CreatifyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Creatify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Creatify *CreatifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Creatify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Creatify *CreatifyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Creatify.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Creatify *CreatifyCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Creatify *CreatifySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Creatify.Contract.DEFAULTADMINROLE(&_Creatify.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Creatify *CreatifyCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Creatify.Contract.DEFAULTADMINROLE(&_Creatify.CallOpts)
}

// MARKETPLACEADMINROLE is a free data retrieval call binding the contract method 0xb2a4eea0.
//
// Solidity: function MARKETPLACE_ADMIN_ROLE() view returns(bytes32)
func (_Creatify *CreatifyCaller) MARKETPLACEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "MARKETPLACE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETPLACEADMINROLE is a free data retrieval call binding the contract method 0xb2a4eea0.
//
// Solidity: function MARKETPLACE_ADMIN_ROLE() view returns(bytes32)
func (_Creatify *CreatifySession) MARKETPLACEADMINROLE() ([32]byte, error) {
	return _Creatify.Contract.MARKETPLACEADMINROLE(&_Creatify.CallOpts)
}

// MARKETPLACEADMINROLE is a free data retrieval call binding the contract method 0xb2a4eea0.
//
// Solidity: function MARKETPLACE_ADMIN_ROLE() view returns(bytes32)
func (_Creatify *CreatifyCallerSession) MARKETPLACEADMINROLE() ([32]byte, error) {
	return _Creatify.Contract.MARKETPLACEADMINROLE(&_Creatify.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Creatify *CreatifyCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Creatify *CreatifySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Creatify.Contract.GetRoleAdmin(&_Creatify.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Creatify *CreatifyCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Creatify.Contract.GetRoleAdmin(&_Creatify.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Creatify *CreatifyCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Creatify *CreatifySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Creatify.Contract.GetRoleMember(&_Creatify.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Creatify *CreatifyCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Creatify.Contract.GetRoleMember(&_Creatify.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Creatify *CreatifyCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Creatify *CreatifySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Creatify.Contract.GetRoleMemberCount(&_Creatify.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Creatify *CreatifyCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Creatify.Contract.GetRoleMemberCount(&_Creatify.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Creatify *CreatifyCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Creatify *CreatifySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Creatify.Contract.HasRole(&_Creatify.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Creatify *CreatifyCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Creatify.Contract.HasRole(&_Creatify.CallOpts, role, account)
}

// IdToMarketItem is a free data retrieval call binding the contract method 0xe61a70c0.
//
// Solidity: function idToMarketItem(uint256 ) view returns(uint256 itemId, address nftContract, uint256 tokenId, address seller, address owner, uint256 price, bool forSale, bool deleted)
func (_Creatify *CreatifyCaller) IdToMarketItem(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	ForSale     bool
	Deleted     bool
}, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "idToMarketItem", arg0)

	outstruct := new(struct {
		ItemId      *big.Int
		NftContract common.Address
		TokenId     *big.Int
		Seller      common.Address
		Owner       common.Address
		Price       *big.Int
		ForSale     bool
		Deleted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ItemId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ForSale = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Deleted = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// IdToMarketItem is a free data retrieval call binding the contract method 0xe61a70c0.
//
// Solidity: function idToMarketItem(uint256 ) view returns(uint256 itemId, address nftContract, uint256 tokenId, address seller, address owner, uint256 price, bool forSale, bool deleted)
func (_Creatify *CreatifySession) IdToMarketItem(arg0 *big.Int) (struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	ForSale     bool
	Deleted     bool
}, error) {
	return _Creatify.Contract.IdToMarketItem(&_Creatify.CallOpts, arg0)
}

// IdToMarketItem is a free data retrieval call binding the contract method 0xe61a70c0.
//
// Solidity: function idToMarketItem(uint256 ) view returns(uint256 itemId, address nftContract, uint256 tokenId, address seller, address owner, uint256 price, bool forSale, bool deleted)
func (_Creatify *CreatifyCallerSession) IdToMarketItem(arg0 *big.Int) (struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	ForSale     bool
	Deleted     bool
}, error) {
	return _Creatify.Contract.IdToMarketItem(&_Creatify.CallOpts, arg0)
}

// PayoutAddress is a free data retrieval call binding the contract method 0x5b8d02d7.
//
// Solidity: function payoutAddress() view returns(address)
func (_Creatify *CreatifyCaller) PayoutAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "payoutAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PayoutAddress is a free data retrieval call binding the contract method 0x5b8d02d7.
//
// Solidity: function payoutAddress() view returns(address)
func (_Creatify *CreatifySession) PayoutAddress() (common.Address, error) {
	return _Creatify.Contract.PayoutAddress(&_Creatify.CallOpts)
}

// PayoutAddress is a free data retrieval call binding the contract method 0x5b8d02d7.
//
// Solidity: function payoutAddress() view returns(address)
func (_Creatify *CreatifyCallerSession) PayoutAddress() (common.Address, error) {
	return _Creatify.Contract.PayoutAddress(&_Creatify.CallOpts)
}

// PlatformFeeBasisPoint is a free data retrieval call binding the contract method 0x5a9cd033.
//
// Solidity: function platformFeeBasisPoint() view returns(uint96)
func (_Creatify *CreatifyCaller) PlatformFeeBasisPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "platformFeeBasisPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeeBasisPoint is a free data retrieval call binding the contract method 0x5a9cd033.
//
// Solidity: function platformFeeBasisPoint() view returns(uint96)
func (_Creatify *CreatifySession) PlatformFeeBasisPoint() (*big.Int, error) {
	return _Creatify.Contract.PlatformFeeBasisPoint(&_Creatify.CallOpts)
}

// PlatformFeeBasisPoint is a free data retrieval call binding the contract method 0x5a9cd033.
//
// Solidity: function platformFeeBasisPoint() view returns(uint96)
func (_Creatify *CreatifyCallerSession) PlatformFeeBasisPoint() (*big.Int, error) {
	return _Creatify.Contract.PlatformFeeBasisPoint(&_Creatify.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Creatify *CreatifyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Creatify.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Creatify *CreatifySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Creatify.Contract.SupportsInterface(&_Creatify.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Creatify *CreatifyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Creatify.Contract.SupportsInterface(&_Creatify.CallOpts, interfaceId)
}

// ChangeFeeAndPayoutAddress is a paid mutator transaction binding the contract method 0x3738b55a.
//
// Solidity: function changeFeeAndPayoutAddress(uint96 newPlatformFee, address newPayoutAddress) returns()
func (_Creatify *CreatifyTransactor) ChangeFeeAndPayoutAddress(opts *bind.TransactOpts, newPlatformFee *big.Int, newPayoutAddress common.Address) (*types.Transaction, error) {
	return _Creatify.contract.Transact(opts, "changeFeeAndPayoutAddress", newPlatformFee, newPayoutAddress)
}

// ChangeFeeAndPayoutAddress is a paid mutator transaction binding the contract method 0x3738b55a.
//
// Solidity: function changeFeeAndPayoutAddress(uint96 newPlatformFee, address newPayoutAddress) returns()
func (_Creatify *CreatifySession) ChangeFeeAndPayoutAddress(newPlatformFee *big.Int, newPayoutAddress common.Address) (*types.Transaction, error) {
	return _Creatify.Contract.ChangeFeeAndPayoutAddress(&_Creatify.TransactOpts, newPlatformFee, newPayoutAddress)
}

// ChangeFeeAndPayoutAddress is a paid mutator transaction binding the contract method 0x3738b55a.
//
// Solidity: function changeFeeAndPayoutAddress(uint96 newPlatformFee, address newPayoutAddress) returns()
func (_Creatify *CreatifyTransactorSession) ChangeFeeAndPayoutAddress(newPlatformFee *big.Int, newPayoutAddress common.Address) (*types.Transaction, error) {
	return _Creatify.Contract.ChangeFeeAndPayoutAddress(&_Creatify.TransactOpts, newPlatformFee, newPayoutAddress)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_Creatify *CreatifyTransactor) CreateMarketItem(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Creatify.contract.Transact(opts, "createMarketItem", nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_Creatify *CreatifySession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Creatify.Contract.CreateMarketItem(&_Creatify.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_Creatify *CreatifyTransactorSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Creatify.Contract.CreateMarketItem(&_Creatify.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Creatify *CreatifyTransactor) CreateMarketSale(opts *bind.TransactOpts, nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Creatify.contract.Transact(opts, "createMarketSale", nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Creatify *CreatifySession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Creatify.Contract.CreateMarketSale(&_Creatify.TransactOpts, nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_Creatify *CreatifyTransactorSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _Creatify.Contract.CreateMarketSale(&_Creatify.TransactOpts, nftContract, itemId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Creatify *CreatifyTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Creatify *CreatifySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.Contract.GrantRole(&_Creatify.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Creatify *CreatifyTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.Contract.GrantRole(&_Creatify.TransactOpts, role, account)
}

// RemoveFromSale is a paid mutator transaction binding the contract method 0x1361a3b6.
//
// Solidity: function removeFromSale(uint256 itemId) returns()
func (_Creatify *CreatifyTransactor) RemoveFromSale(opts *bind.TransactOpts, itemId *big.Int) (*types.Transaction, error) {
	return _Creatify.contract.Transact(opts, "removeFromSale", itemId)
}

// RemoveFromSale is a paid mutator transaction binding the contract method 0x1361a3b6.
//
// Solidity: function removeFromSale(uint256 itemId) returns()
func (_Creatify *CreatifySession) RemoveFromSale(itemId *big.Int) (*types.Transaction, error) {
	return _Creatify.Contract.RemoveFromSale(&_Creatify.TransactOpts, itemId)
}

// RemoveFromSale is a paid mutator transaction binding the contract method 0x1361a3b6.
//
// Solidity: function removeFromSale(uint256 itemId) returns()
func (_Creatify *CreatifyTransactorSession) RemoveFromSale(itemId *big.Int) (*types.Transaction, error) {
	return _Creatify.Contract.RemoveFromSale(&_Creatify.TransactOpts, itemId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Creatify *CreatifyTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Creatify *CreatifySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.Contract.RenounceRole(&_Creatify.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Creatify *CreatifyTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.Contract.RenounceRole(&_Creatify.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Creatify *CreatifyTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Creatify *CreatifySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.Contract.RevokeRole(&_Creatify.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Creatify *CreatifyTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Creatify.Contract.RevokeRole(&_Creatify.TransactOpts, role, account)
}

// CreatifyMarketItemCreatedIterator is returned from FilterMarketItemCreated and is used to iterate over the raw logs and unpacked data for MarketItemCreated events raised by the Creatify contract.
type CreatifyMarketItemCreatedIterator struct {
	Event *CreatifyMarketItemCreated // Event containing the contract specifics and raw log

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
func (it *CreatifyMarketItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatifyMarketItemCreated)
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
		it.Event = new(CreatifyMarketItemCreated)
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
func (it *CreatifyMarketItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatifyMarketItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatifyMarketItemCreated represents a MarketItemCreated event raised by the Creatify contract.
type CreatifyMarketItemCreated struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	MetaDataURI string
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	ForSale     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketItemCreated is a free log retrieval operation binding the contract event 0xd39071221960ab18bca34a3fb540f0da19655735105a97ecd89dc2482dc4f857.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, string metaDataURI, address seller, address owner, uint256 price, bool forSale)
func (_Creatify *CreatifyFilterer) FilterMarketItemCreated(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*CreatifyMarketItemCreatedIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Creatify.contract.FilterLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CreatifyMarketItemCreatedIterator{contract: _Creatify.contract, event: "MarketItemCreated", logs: logs, sub: sub}, nil
}

// WatchMarketItemCreated is a free log subscription operation binding the contract event 0xd39071221960ab18bca34a3fb540f0da19655735105a97ecd89dc2482dc4f857.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, string metaDataURI, address seller, address owner, uint256 price, bool forSale)
func (_Creatify *CreatifyFilterer) WatchMarketItemCreated(opts *bind.WatchOpts, sink chan<- *CreatifyMarketItemCreated, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Creatify.contract.WatchLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatifyMarketItemCreated)
				if err := _Creatify.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
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

// ParseMarketItemCreated is a log parse operation binding the contract event 0xd39071221960ab18bca34a3fb540f0da19655735105a97ecd89dc2482dc4f857.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, string metaDataURI, address seller, address owner, uint256 price, bool forSale)
func (_Creatify *CreatifyFilterer) ParseMarketItemCreated(log types.Log) (*CreatifyMarketItemCreated, error) {
	event := new(CreatifyMarketItemCreated)
	if err := _Creatify.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatifyMarketItemRemovedIterator is returned from FilterMarketItemRemoved and is used to iterate over the raw logs and unpacked data for MarketItemRemoved events raised by the Creatify contract.
type CreatifyMarketItemRemovedIterator struct {
	Event *CreatifyMarketItemRemoved // Event containing the contract specifics and raw log

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
func (it *CreatifyMarketItemRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatifyMarketItemRemoved)
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
		it.Event = new(CreatifyMarketItemRemoved)
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
func (it *CreatifyMarketItemRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatifyMarketItemRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatifyMarketItemRemoved represents a MarketItemRemoved event raised by the Creatify contract.
type CreatifyMarketItemRemoved struct {
	ItemId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketItemRemoved is a free log retrieval operation binding the contract event 0xd371e668750cb458fa9a55e99ade07ce913d63ab733d6e30fe303723e106cf96.
//
// Solidity: event MarketItemRemoved(uint256 itemId)
func (_Creatify *CreatifyFilterer) FilterMarketItemRemoved(opts *bind.FilterOpts) (*CreatifyMarketItemRemovedIterator, error) {

	logs, sub, err := _Creatify.contract.FilterLogs(opts, "MarketItemRemoved")
	if err != nil {
		return nil, err
	}
	return &CreatifyMarketItemRemovedIterator{contract: _Creatify.contract, event: "MarketItemRemoved", logs: logs, sub: sub}, nil
}

// WatchMarketItemRemoved is a free log subscription operation binding the contract event 0xd371e668750cb458fa9a55e99ade07ce913d63ab733d6e30fe303723e106cf96.
//
// Solidity: event MarketItemRemoved(uint256 itemId)
func (_Creatify *CreatifyFilterer) WatchMarketItemRemoved(opts *bind.WatchOpts, sink chan<- *CreatifyMarketItemRemoved) (event.Subscription, error) {

	logs, sub, err := _Creatify.contract.WatchLogs(opts, "MarketItemRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatifyMarketItemRemoved)
				if err := _Creatify.contract.UnpackLog(event, "MarketItemRemoved", log); err != nil {
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

// ParseMarketItemRemoved is a log parse operation binding the contract event 0xd371e668750cb458fa9a55e99ade07ce913d63ab733d6e30fe303723e106cf96.
//
// Solidity: event MarketItemRemoved(uint256 itemId)
func (_Creatify *CreatifyFilterer) ParseMarketItemRemoved(log types.Log) (*CreatifyMarketItemRemoved, error) {
	event := new(CreatifyMarketItemRemoved)
	if err := _Creatify.contract.UnpackLog(event, "MarketItemRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatifyMarketItemSoldIterator is returned from FilterMarketItemSold and is used to iterate over the raw logs and unpacked data for MarketItemSold events raised by the Creatify contract.
type CreatifyMarketItemSoldIterator struct {
	Event *CreatifyMarketItemSold // Event containing the contract specifics and raw log

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
func (it *CreatifyMarketItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatifyMarketItemSold)
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
		it.Event = new(CreatifyMarketItemSold)
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
func (it *CreatifyMarketItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatifyMarketItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatifyMarketItemSold represents a MarketItemSold event raised by the Creatify contract.
type CreatifyMarketItemSold struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Buyer       common.Address
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketItemSold is a free log retrieval operation binding the contract event 0x9789d8d6748e7f3e6c12fe6b244e2765b7e805f6b4b65a2474cad0ca8e788408.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address buyer, uint256 price)
func (_Creatify *CreatifyFilterer) FilterMarketItemSold(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*CreatifyMarketItemSoldIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Creatify.contract.FilterLogs(opts, "MarketItemSold", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CreatifyMarketItemSoldIterator{contract: _Creatify.contract, event: "MarketItemSold", logs: logs, sub: sub}, nil
}

// WatchMarketItemSold is a free log subscription operation binding the contract event 0x9789d8d6748e7f3e6c12fe6b244e2765b7e805f6b4b65a2474cad0ca8e788408.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address buyer, uint256 price)
func (_Creatify *CreatifyFilterer) WatchMarketItemSold(opts *bind.WatchOpts, sink chan<- *CreatifyMarketItemSold, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Creatify.contract.WatchLogs(opts, "MarketItemSold", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatifyMarketItemSold)
				if err := _Creatify.contract.UnpackLog(event, "MarketItemSold", log); err != nil {
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

// ParseMarketItemSold is a log parse operation binding the contract event 0x9789d8d6748e7f3e6c12fe6b244e2765b7e805f6b4b65a2474cad0ca8e788408.
//
// Solidity: event MarketItemSold(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address buyer, uint256 price)
func (_Creatify *CreatifyFilterer) ParseMarketItemSold(log types.Log) (*CreatifyMarketItemSold, error) {
	event := new(CreatifyMarketItemSold)
	if err := _Creatify.contract.UnpackLog(event, "MarketItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatifyRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Creatify contract.
type CreatifyRoleAdminChangedIterator struct {
	Event *CreatifyRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CreatifyRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatifyRoleAdminChanged)
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
		it.Event = new(CreatifyRoleAdminChanged)
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
func (it *CreatifyRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatifyRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatifyRoleAdminChanged represents a RoleAdminChanged event raised by the Creatify contract.
type CreatifyRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Creatify *CreatifyFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CreatifyRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Creatify.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CreatifyRoleAdminChangedIterator{contract: _Creatify.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Creatify *CreatifyFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CreatifyRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Creatify.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatifyRoleAdminChanged)
				if err := _Creatify.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Creatify *CreatifyFilterer) ParseRoleAdminChanged(log types.Log) (*CreatifyRoleAdminChanged, error) {
	event := new(CreatifyRoleAdminChanged)
	if err := _Creatify.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatifyRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Creatify contract.
type CreatifyRoleGrantedIterator struct {
	Event *CreatifyRoleGranted // Event containing the contract specifics and raw log

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
func (it *CreatifyRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatifyRoleGranted)
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
		it.Event = new(CreatifyRoleGranted)
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
func (it *CreatifyRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatifyRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatifyRoleGranted represents a RoleGranted event raised by the Creatify contract.
type CreatifyRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Creatify *CreatifyFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CreatifyRoleGrantedIterator, error) {

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

	logs, sub, err := _Creatify.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatifyRoleGrantedIterator{contract: _Creatify.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Creatify *CreatifyFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CreatifyRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Creatify.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatifyRoleGranted)
				if err := _Creatify.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Creatify *CreatifyFilterer) ParseRoleGranted(log types.Log) (*CreatifyRoleGranted, error) {
	event := new(CreatifyRoleGranted)
	if err := _Creatify.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreatifyRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Creatify contract.
type CreatifyRoleRevokedIterator struct {
	Event *CreatifyRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CreatifyRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreatifyRoleRevoked)
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
		it.Event = new(CreatifyRoleRevoked)
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
func (it *CreatifyRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreatifyRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreatifyRoleRevoked represents a RoleRevoked event raised by the Creatify contract.
type CreatifyRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Creatify *CreatifyFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CreatifyRoleRevokedIterator, error) {

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

	logs, sub, err := _Creatify.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreatifyRoleRevokedIterator{contract: _Creatify.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Creatify *CreatifyFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CreatifyRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Creatify.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreatifyRoleRevoked)
				if err := _Creatify.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Creatify *CreatifyFilterer) ParseRoleRevoked(log types.Log) (*CreatifyRoleRevoked, error) {
	event := new(CreatifyRoleRevoked)
	if err := _Creatify.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
