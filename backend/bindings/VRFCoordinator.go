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

// VRFProof is an auto generated low-level Go binding around an user-defined struct.
type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

// VRFRequest is an auto generated low-level Go binding around an user-defined struct.
type VRFRequest struct {
	Sender               common.Address
	Nonce                *big.Int
	OracleId             [32]byte
	NbWords              uint32
	RequestConfirmations uint16
	CallbackGasLimit     uint32
	CallbackSelector     [4]byte
}

// VRFCoordinatorMetaData contains all meta data concerning the VRFCoordinator contract.
var VRFCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageProofOracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedToFulfillRandomness\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCallbackGasLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNumberOfWords\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOracleId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestParameters\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"oracleId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nbWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"callbackSelector\",\"type\":\"bytes4\"}],\"name\":\"RequestRandomWords\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAXIMUM_CALLBACK_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_NB_WORDS\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"oracleId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nbWords\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes4\",\"name\":\"callbackSelector\",\"type\":\"bytes4\"}],\"internalType\":\"structVRF.Request\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_oracleId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nbWords\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackSelector\",\"type\":\"bytes4\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageProofOracle\",\"outputs\":[{\"internalType\":\"contractStorageProofOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VRFCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFCoordinatorMetaData.ABI instead.
var VRFCoordinatorABI = VRFCoordinatorMetaData.ABI

// VRFCoordinator is an auto generated Go binding around an Ethereum contract.
type VRFCoordinator struct {
	VRFCoordinatorCaller     // Read-only binding to the contract
	VRFCoordinatorTransactor // Write-only binding to the contract
	VRFCoordinatorFilterer   // Log filterer for contract events
}

// VRFCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFCoordinatorSession struct {
	Contract     *VRFCoordinator   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCoordinatorCallerSession struct {
	Contract *VRFCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VRFCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFCoordinatorTransactorSession struct {
	Contract     *VRFCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VRFCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFCoordinatorRaw struct {
	Contract *VRFCoordinator // Generic contract binding to access the raw methods on
}

// VRFCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCoordinatorCallerRaw struct {
	Contract *VRFCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// VRFCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFCoordinatorTransactorRaw struct {
	Contract *VRFCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFCoordinator creates a new instance of VRFCoordinator, bound to a specific deployed contract.
func NewVRFCoordinator(address common.Address, backend bind.ContractBackend) (*VRFCoordinator, error) {
	contract, err := bindVRFCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinator{VRFCoordinatorCaller: VRFCoordinatorCaller{contract: contract}, VRFCoordinatorTransactor: VRFCoordinatorTransactor{contract: contract}, VRFCoordinatorFilterer: VRFCoordinatorFilterer{contract: contract}}, nil
}

// NewVRFCoordinatorCaller creates a new read-only instance of VRFCoordinator, bound to a specific deployed contract.
func NewVRFCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorCaller, error) {
	contract, err := bindVRFCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorCaller{contract: contract}, nil
}

// NewVRFCoordinatorTransactor creates a new write-only instance of VRFCoordinator, bound to a specific deployed contract.
func NewVRFCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorTransactor, error) {
	contract, err := bindVRFCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorTransactor{contract: contract}, nil
}

// NewVRFCoordinatorFilterer creates a new log filterer instance of VRFCoordinator, bound to a specific deployed contract.
func NewVRFCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorFilterer, error) {
	contract, err := bindVRFCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorFilterer{contract: contract}, nil
}

// bindVRFCoordinator binds a generic wrapper to an already deployed contract.
func bindVRFCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinator *VRFCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinator.Contract.VRFCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinator *VRFCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.VRFCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinator *VRFCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.VRFCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinator *VRFCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinator *VRFCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinator *VRFCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.contract.Transact(opts, method, params...)
}

// MAXIMUMCALLBACKGASLIMIT is a free data retrieval call binding the contract method 0x3484b13b.
//
// Solidity: function MAXIMUM_CALLBACK_GAS_LIMIT() view returns(uint64)
func (_VRFCoordinator *VRFCoordinatorCaller) MAXIMUMCALLBACKGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "MAXIMUM_CALLBACK_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMCALLBACKGASLIMIT is a free data retrieval call binding the contract method 0x3484b13b.
//
// Solidity: function MAXIMUM_CALLBACK_GAS_LIMIT() view returns(uint64)
func (_VRFCoordinator *VRFCoordinatorSession) MAXIMUMCALLBACKGASLIMIT() (uint64, error) {
	return _VRFCoordinator.Contract.MAXIMUMCALLBACKGASLIMIT(&_VRFCoordinator.CallOpts)
}

// MAXIMUMCALLBACKGASLIMIT is a free data retrieval call binding the contract method 0x3484b13b.
//
// Solidity: function MAXIMUM_CALLBACK_GAS_LIMIT() view returns(uint64)
func (_VRFCoordinator *VRFCoordinatorCallerSession) MAXIMUMCALLBACKGASLIMIT() (uint64, error) {
	return _VRFCoordinator.Contract.MAXIMUMCALLBACKGASLIMIT(&_VRFCoordinator.CallOpts)
}

// MAXIMUMNBWORDS is a free data retrieval call binding the contract method 0x23fe886c.
//
// Solidity: function MAXIMUM_NB_WORDS() view returns(uint64)
func (_VRFCoordinator *VRFCoordinatorCaller) MAXIMUMNBWORDS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "MAXIMUM_NB_WORDS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXIMUMNBWORDS is a free data retrieval call binding the contract method 0x23fe886c.
//
// Solidity: function MAXIMUM_NB_WORDS() view returns(uint64)
func (_VRFCoordinator *VRFCoordinatorSession) MAXIMUMNBWORDS() (uint64, error) {
	return _VRFCoordinator.Contract.MAXIMUMNBWORDS(&_VRFCoordinator.CallOpts)
}

// MAXIMUMNBWORDS is a free data retrieval call binding the contract method 0x23fe886c.
//
// Solidity: function MAXIMUM_NB_WORDS() view returns(uint64)
func (_VRFCoordinator *VRFCoordinatorCallerSession) MAXIMUMNBWORDS() (uint64, error) {
	return _VRFCoordinator.Contract.MAXIMUMNBWORDS(&_VRFCoordinator.CallOpts)
}

// MINIMUMREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x88dfbadd.
//
// Solidity: function MINIMUM_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_VRFCoordinator *VRFCoordinatorCaller) MINIMUMREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "MINIMUM_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MINIMUMREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x88dfbadd.
//
// Solidity: function MINIMUM_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_VRFCoordinator *VRFCoordinatorSession) MINIMUMREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinator.Contract.MINIMUMREQUESTCONFIRMATIONS(&_VRFCoordinator.CallOpts)
}

// MINIMUMREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x88dfbadd.
//
// Solidity: function MINIMUM_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_VRFCoordinator *VRFCoordinatorCallerSession) MINIMUMREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinator.Contract.MINIMUMREQUESTCONFIRMATIONS(&_VRFCoordinator.CallOpts)
}

// ORACLEADDRESS is a free data retrieval call binding the contract method 0xfcfc430c.
//
// Solidity: function ORACLE_ADDRESS() view returns(address)
func (_VRFCoordinator *VRFCoordinatorCaller) ORACLEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "ORACLE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLEADDRESS is a free data retrieval call binding the contract method 0xfcfc430c.
//
// Solidity: function ORACLE_ADDRESS() view returns(address)
func (_VRFCoordinator *VRFCoordinatorSession) ORACLEADDRESS() (common.Address, error) {
	return _VRFCoordinator.Contract.ORACLEADDRESS(&_VRFCoordinator.CallOpts)
}

// ORACLEADDRESS is a free data retrieval call binding the contract method 0xfcfc430c.
//
// Solidity: function ORACLE_ADDRESS() view returns(address)
func (_VRFCoordinator *VRFCoordinatorCallerSession) ORACLEADDRESS() (common.Address, error) {
	return _VRFCoordinator.Contract.ORACLEADDRESS(&_VRFCoordinator.CallOpts)
}

// ORACLEID is a free data retrieval call binding the contract method 0x7a2ae11d.
//
// Solidity: function ORACLE_ID() view returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorCaller) ORACLEID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "ORACLE_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEID is a free data retrieval call binding the contract method 0x7a2ae11d.
//
// Solidity: function ORACLE_ID() view returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorSession) ORACLEID() ([32]byte, error) {
	return _VRFCoordinator.Contract.ORACLEID(&_VRFCoordinator.CallOpts)
}

// ORACLEID is a free data retrieval call binding the contract method 0x7a2ae11d.
//
// Solidity: function ORACLE_ID() view returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorCallerSession) ORACLEID() ([32]byte, error) {
	return _VRFCoordinator.Contract.ORACLEID(&_VRFCoordinator.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_VRFCoordinator *VRFCoordinatorCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_VRFCoordinator *VRFCoordinatorSession) Nonce() (*big.Int, error) {
	return _VRFCoordinator.Contract.Nonce(&_VRFCoordinator.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_VRFCoordinator *VRFCoordinatorCallerSession) Nonce() (*big.Int, error) {
	return _VRFCoordinator.Contract.Nonce(&_VRFCoordinator.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0xa81a2677.
//
// Solidity: function oracles(bytes32 ) view returns(address)
func (_VRFCoordinator *VRFCoordinatorCaller) Oracles(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "oracles", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracles is a free data retrieval call binding the contract method 0xa81a2677.
//
// Solidity: function oracles(bytes32 ) view returns(address)
func (_VRFCoordinator *VRFCoordinatorSession) Oracles(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinator.Contract.Oracles(&_VRFCoordinator.CallOpts, arg0)
}

// Oracles is a free data retrieval call binding the contract method 0xa81a2677.
//
// Solidity: function oracles(bytes32 ) view returns(address)
func (_VRFCoordinator *VRFCoordinatorCallerSession) Oracles(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinator.Contract.Oracles(&_VRFCoordinator.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "requests", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorSession) Requests(arg0 [32]byte) ([32]byte, error) {
	return _VRFCoordinator.Contract.Requests(&_VRFCoordinator.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorCallerSession) Requests(arg0 [32]byte) ([32]byte, error) {
	return _VRFCoordinator.Contract.Requests(&_VRFCoordinator.CallOpts, arg0)
}

// StorageProofOracle is a free data retrieval call binding the contract method 0x67347b5f.
//
// Solidity: function storageProofOracle() view returns(address)
func (_VRFCoordinator *VRFCoordinatorCaller) StorageProofOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinator.contract.Call(opts, &out, "storageProofOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageProofOracle is a free data retrieval call binding the contract method 0x67347b5f.
//
// Solidity: function storageProofOracle() view returns(address)
func (_VRFCoordinator *VRFCoordinatorSession) StorageProofOracle() (common.Address, error) {
	return _VRFCoordinator.Contract.StorageProofOracle(&_VRFCoordinator.CallOpts)
}

// StorageProofOracle is a free data retrieval call binding the contract method 0x67347b5f.
//
// Solidity: function storageProofOracle() view returns(address)
func (_VRFCoordinator *VRFCoordinatorCallerSession) StorageProofOracle() (common.Address, error) {
	return _VRFCoordinator.Contract.StorageProofOracle(&_VRFCoordinator.CallOpts)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xc5cb1c8e.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (address,uint256,bytes32,uint32,uint16,uint32,bytes4) _request) returns()
func (_VRFCoordinator *VRFCoordinatorTransactor) FulfillRandomWords(opts *bind.TransactOpts, _proof VRFProof, _request VRFRequest) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "fulfillRandomWords", _proof, _request)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xc5cb1c8e.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (address,uint256,bytes32,uint32,uint16,uint32,bytes4) _request) returns()
func (_VRFCoordinator *VRFCoordinatorSession) FulfillRandomWords(_proof VRFProof, _request VRFRequest) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.FulfillRandomWords(&_VRFCoordinator.TransactOpts, _proof, _request)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xc5cb1c8e.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (address,uint256,bytes32,uint32,uint16,uint32,bytes4) _request) returns()
func (_VRFCoordinator *VRFCoordinatorTransactorSession) FulfillRandomWords(_proof VRFProof, _request VRFRequest) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.FulfillRandomWords(&_VRFCoordinator.TransactOpts, _proof, _request)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x21a5eaa3.
//
// Solidity: function requestRandomWords(bytes32 _oracleId, uint32 _nbWords, uint16 _requestConfirmations, uint32 _callbackGasLimit, bytes4 _callbackSelector) returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorTransactor) RequestRandomWords(opts *bind.TransactOpts, _oracleId [32]byte, _nbWords uint32, _requestConfirmations uint16, _callbackGasLimit uint32, _callbackSelector [4]byte) (*types.Transaction, error) {
	return _VRFCoordinator.contract.Transact(opts, "requestRandomWords", _oracleId, _nbWords, _requestConfirmations, _callbackGasLimit, _callbackSelector)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x21a5eaa3.
//
// Solidity: function requestRandomWords(bytes32 _oracleId, uint32 _nbWords, uint16 _requestConfirmations, uint32 _callbackGasLimit, bytes4 _callbackSelector) returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorSession) RequestRandomWords(_oracleId [32]byte, _nbWords uint32, _requestConfirmations uint16, _callbackGasLimit uint32, _callbackSelector [4]byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RequestRandomWords(&_VRFCoordinator.TransactOpts, _oracleId, _nbWords, _requestConfirmations, _callbackGasLimit, _callbackSelector)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x21a5eaa3.
//
// Solidity: function requestRandomWords(bytes32 _oracleId, uint32 _nbWords, uint16 _requestConfirmations, uint32 _callbackGasLimit, bytes4 _callbackSelector) returns(bytes32)
func (_VRFCoordinator *VRFCoordinatorTransactorSession) RequestRandomWords(_oracleId [32]byte, _nbWords uint32, _requestConfirmations uint16, _callbackGasLimit uint32, _callbackSelector [4]byte) (*types.Transaction, error) {
	return _VRFCoordinator.Contract.RequestRandomWords(&_VRFCoordinator.TransactOpts, _oracleId, _nbWords, _requestConfirmations, _callbackGasLimit, _callbackSelector)
}

// VRFCoordinatorRequestRandomWordsIterator is returned from FilterRequestRandomWords and is used to iterate over the raw logs and unpacked data for RequestRandomWords events raised by the VRFCoordinator contract.
type VRFCoordinatorRequestRandomWordsIterator struct {
	Event *VRFCoordinatorRequestRandomWords // Event containing the contract specifics and raw log

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
func (it *VRFCoordinatorRequestRandomWordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorRequestRandomWords)
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
		it.Event = new(VRFCoordinatorRequestRandomWords)
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
func (it *VRFCoordinatorRequestRandomWordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFCoordinatorRequestRandomWordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFCoordinatorRequestRandomWords represents a RequestRandomWords event raised by the VRFCoordinator contract.
type VRFCoordinatorRequestRandomWords struct {
	RequestId            [32]byte
	Sender               common.Address
	Nonce                *big.Int
	OracleId             [32]byte
	NbWords              uint32
	RequestConfirmations uint16
	CallbackGasLimit     uint32
	CallbackSelector     [4]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRequestRandomWords is a free log retrieval operation binding the contract event 0x62dbbfcf8ec894ad82eb4220e6771191f27422ba4349c3da86dbe46817403b2e.
//
// Solidity: event RequestRandomWords(bytes32 requestId, address sender, uint256 nonce, bytes32 oracleId, uint32 nbWords, uint16 requestConfirmations, uint32 callbackGasLimit, bytes4 callbackSelector)
func (_VRFCoordinator *VRFCoordinatorFilterer) FilterRequestRandomWords(opts *bind.FilterOpts) (*VRFCoordinatorRequestRandomWordsIterator, error) {

	logs, sub, err := _VRFCoordinator.contract.FilterLogs(opts, "RequestRandomWords")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorRequestRandomWordsIterator{contract: _VRFCoordinator.contract, event: "RequestRandomWords", logs: logs, sub: sub}, nil
}

// WatchRequestRandomWords is a free log subscription operation binding the contract event 0x62dbbfcf8ec894ad82eb4220e6771191f27422ba4349c3da86dbe46817403b2e.
//
// Solidity: event RequestRandomWords(bytes32 requestId, address sender, uint256 nonce, bytes32 oracleId, uint32 nbWords, uint16 requestConfirmations, uint32 callbackGasLimit, bytes4 callbackSelector)
func (_VRFCoordinator *VRFCoordinatorFilterer) WatchRequestRandomWords(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorRequestRandomWords) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinator.contract.WatchLogs(opts, "RequestRandomWords")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFCoordinatorRequestRandomWords)
				if err := _VRFCoordinator.contract.UnpackLog(event, "RequestRandomWords", log); err != nil {
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

// ParseRequestRandomWords is a log parse operation binding the contract event 0x62dbbfcf8ec894ad82eb4220e6771191f27422ba4349c3da86dbe46817403b2e.
//
// Solidity: event RequestRandomWords(bytes32 requestId, address sender, uint256 nonce, bytes32 oracleId, uint32 nbWords, uint16 requestConfirmations, uint32 callbackGasLimit, bytes4 callbackSelector)
func (_VRFCoordinator *VRFCoordinatorFilterer) ParseRequestRandomWords(log types.Log) (*VRFCoordinatorRequestRandomWords, error) {
	event := new(VRFCoordinatorRequestRandomWords)
	if err := _VRFCoordinator.contract.UnpackLog(event, "RequestRandomWords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
