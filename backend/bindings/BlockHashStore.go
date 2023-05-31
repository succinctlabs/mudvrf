// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// BlockHashStoreMetaData contains all meta data concerning the BlockHashStore contract.
var BlockHashStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_blockHeaders\",\"type\":\"bytes[]\"}],\"name\":\"storeBlockHashesViaMerkleProofs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_blockNumbers\",\"type\":\"uint256[]\"}],\"name\":\"storeBlockHashesViaOpCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BlockHashStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockHashStoreMetaData.ABI instead.
var BlockHashStoreABI = BlockHashStoreMetaData.ABI

// BlockHashStore is an auto generated Go binding around an Ethereum contract.
type BlockHashStore struct {
	BlockHashStoreCaller     // Read-only binding to the contract
	BlockHashStoreTransactor // Write-only binding to the contract
	BlockHashStoreFilterer   // Log filterer for contract events
}

// BlockHashStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockHashStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockHashStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockHashStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockHashStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockHashStoreSession struct {
	Contract     *BlockHashStore   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockHashStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockHashStoreCallerSession struct {
	Contract *BlockHashStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BlockHashStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockHashStoreTransactorSession struct {
	Contract     *BlockHashStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BlockHashStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockHashStoreRaw struct {
	Contract *BlockHashStore // Generic contract binding to access the raw methods on
}

// BlockHashStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockHashStoreCallerRaw struct {
	Contract *BlockHashStoreCaller // Generic read-only contract binding to access the raw methods on
}

// BlockHashStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockHashStoreTransactorRaw struct {
	Contract *BlockHashStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockHashStore creates a new instance of BlockHashStore, bound to a specific deployed contract.
func NewBlockHashStore(address common.Address, backend bind.ContractBackend) (*BlockHashStore, error) {
	contract, err := bindBlockHashStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockHashStore{BlockHashStoreCaller: BlockHashStoreCaller{contract: contract}, BlockHashStoreTransactor: BlockHashStoreTransactor{contract: contract}, BlockHashStoreFilterer: BlockHashStoreFilterer{contract: contract}}, nil
}

// NewBlockHashStoreCaller creates a new read-only instance of BlockHashStore, bound to a specific deployed contract.
func NewBlockHashStoreCaller(address common.Address, caller bind.ContractCaller) (*BlockHashStoreCaller, error) {
	contract, err := bindBlockHashStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashStoreCaller{contract: contract}, nil
}

// NewBlockHashStoreTransactor creates a new write-only instance of BlockHashStore, bound to a specific deployed contract.
func NewBlockHashStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockHashStoreTransactor, error) {
	contract, err := bindBlockHashStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockHashStoreTransactor{contract: contract}, nil
}

// NewBlockHashStoreFilterer creates a new log filterer instance of BlockHashStore, bound to a specific deployed contract.
func NewBlockHashStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockHashStoreFilterer, error) {
	contract, err := bindBlockHashStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockHashStoreFilterer{contract: contract}, nil
}

// bindBlockHashStore binds a generic wrapper to an already deployed contract.
func bindBlockHashStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockHashStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashStore *BlockHashStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashStore.Contract.BlockHashStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashStore *BlockHashStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashStore.Contract.BlockHashStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashStore *BlockHashStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashStore.Contract.BlockHashStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockHashStore *BlockHashStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockHashStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockHashStore *BlockHashStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockHashStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockHashStore *BlockHashStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockHashStore.Contract.contract.Transact(opts, method, params...)
}

// BlockHashes is a free data retrieval call binding the contract method 0x34cdf78d.
//
// Solidity: function blockHashes(uint256 ) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreCaller) BlockHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BlockHashStore.contract.Call(opts, &out, "blockHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BlockHashes is a free data retrieval call binding the contract method 0x34cdf78d.
//
// Solidity: function blockHashes(uint256 ) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreSession) BlockHashes(arg0 *big.Int) ([32]byte, error) {
	return _BlockHashStore.Contract.BlockHashes(&_BlockHashStore.CallOpts, arg0)
}

// BlockHashes is a free data retrieval call binding the contract method 0x34cdf78d.
//
// Solidity: function blockHashes(uint256 ) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreCallerSession) BlockHashes(arg0 *big.Int) ([32]byte, error) {
	return _BlockHashStore.Contract.BlockHashes(&_BlockHashStore.CallOpts, arg0)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 _n) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreCaller) GetBlockHash(opts *bind.CallOpts, _n *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BlockHashStore.contract.Call(opts, &out, "getBlockHash", _n)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 _n) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreSession) GetBlockHash(_n *big.Int) ([32]byte, error) {
	return _BlockHashStore.Contract.GetBlockHash(&_BlockHashStore.CallOpts, _n)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 _n) view returns(bytes32)
func (_BlockHashStore *BlockHashStoreCallerSession) GetBlockHash(_n *big.Int) ([32]byte, error) {
	return _BlockHashStore.Contract.GetBlockHash(&_BlockHashStore.CallOpts, _n)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_BlockHashStore *BlockHashStoreCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockHashStore.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_BlockHashStore *BlockHashStoreSession) LatestBlockNumber() (*big.Int, error) {
	return _BlockHashStore.Contract.LatestBlockNumber(&_BlockHashStore.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_BlockHashStore *BlockHashStoreCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _BlockHashStore.Contract.LatestBlockNumber(&_BlockHashStore.CallOpts)
}

// StoreBlockHashesViaMerkleProofs is a paid mutator transaction binding the contract method 0x50b819da.
//
// Solidity: function storeBlockHashesViaMerkleProofs(bytes[] _blockHeaders) returns()
func (_BlockHashStore *BlockHashStoreTransactor) StoreBlockHashesViaMerkleProofs(opts *bind.TransactOpts, _blockHeaders [][]byte) (*types.Transaction, error) {
	return _BlockHashStore.contract.Transact(opts, "storeBlockHashesViaMerkleProofs", _blockHeaders)
}

// StoreBlockHashesViaMerkleProofs is a paid mutator transaction binding the contract method 0x50b819da.
//
// Solidity: function storeBlockHashesViaMerkleProofs(bytes[] _blockHeaders) returns()
func (_BlockHashStore *BlockHashStoreSession) StoreBlockHashesViaMerkleProofs(_blockHeaders [][]byte) (*types.Transaction, error) {
	return _BlockHashStore.Contract.StoreBlockHashesViaMerkleProofs(&_BlockHashStore.TransactOpts, _blockHeaders)
}

// StoreBlockHashesViaMerkleProofs is a paid mutator transaction binding the contract method 0x50b819da.
//
// Solidity: function storeBlockHashesViaMerkleProofs(bytes[] _blockHeaders) returns()
func (_BlockHashStore *BlockHashStoreTransactorSession) StoreBlockHashesViaMerkleProofs(_blockHeaders [][]byte) (*types.Transaction, error) {
	return _BlockHashStore.Contract.StoreBlockHashesViaMerkleProofs(&_BlockHashStore.TransactOpts, _blockHeaders)
}

// StoreBlockHashesViaOpCode is a paid mutator transaction binding the contract method 0x1aa78fce.
//
// Solidity: function storeBlockHashesViaOpCode(uint256[] _blockNumbers) returns()
func (_BlockHashStore *BlockHashStoreTransactor) StoreBlockHashesViaOpCode(opts *bind.TransactOpts, _blockNumbers []*big.Int) (*types.Transaction, error) {
	return _BlockHashStore.contract.Transact(opts, "storeBlockHashesViaOpCode", _blockNumbers)
}

// StoreBlockHashesViaOpCode is a paid mutator transaction binding the contract method 0x1aa78fce.
//
// Solidity: function storeBlockHashesViaOpCode(uint256[] _blockNumbers) returns()
func (_BlockHashStore *BlockHashStoreSession) StoreBlockHashesViaOpCode(_blockNumbers []*big.Int) (*types.Transaction, error) {
	return _BlockHashStore.Contract.StoreBlockHashesViaOpCode(&_BlockHashStore.TransactOpts, _blockNumbers)
}

// StoreBlockHashesViaOpCode is a paid mutator transaction binding the contract method 0x1aa78fce.
//
// Solidity: function storeBlockHashesViaOpCode(uint256[] _blockNumbers) returns()
func (_BlockHashStore *BlockHashStoreTransactorSession) StoreBlockHashesViaOpCode(_blockNumbers []*big.Int) (*types.Transaction, error) {
	return _BlockHashStore.Contract.StoreBlockHashesViaOpCode(&_BlockHashStore.TransactOpts, _blockNumbers)
}
