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
	BlockNumber      uint64
	CallbackGasLimit uint32
	NbWords          uint32
	Sender           common.Address
	CallbackSelector [4]byte
}

// IWorldMetaData contains all meta data concerning the IWorld contract.
var IWorldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToFulfillRandomness\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"functionSelector\",\"type\":\"bytes4\"}],\"name\":\"FunctionSelectorExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"functionSelector\",\"type\":\"bytes4\"}],\"name\":\"FunctionSelectorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCallbackGasLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNumberOfWords\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOracleId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestParameters\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"}],\"name\":\"InvalidSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"}],\"name\":\"ModuleAlreadyInstalled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"}],\"name\":\"ResourceExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"}],\"name\":\"ResourceNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"StoreCore_DataIndexOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"StoreCore_InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"StoreCore_InvalidFieldNamesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoreCore_NotDynamicField\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoreCore_NotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tableId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"tableIdString\",\"type\":\"string\"}],\"name\":\"StoreCore_TableAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tableId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"tableIdString\",\"type\":\"string\"}],\"name\":\"StoreCore_TableNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"system\",\"type\":\"address\"}],\"name\":\"SystemExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"HelloWorld\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"}],\"name\":\"StoreDeleteRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StoreEphemeralRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StoreSetField\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StoreSetRecord\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"raffleId\",\"type\":\"uint32\"}],\"name\":\"buyTicket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes\",\"name\":\"funcSelectorAndArgs\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"}],\"name\":\"deleteRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"}],\"name\":\"deleteRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitEphemeralRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"emitEphemeralRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"raffleId\",\"type\":\"uint32\"}],\"name\":\"endRaffle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"endRaffleCallback\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nbWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"callbackSelector\",\"type\":\"bytes4\"}],\"internalType\":\"structVRFRequest\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"}],\"name\":\"getField\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"Schema\",\"name\":\"schema\",\"type\":\"bytes32\"}],\"name\":\"getFieldLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"Schema\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getFieldSlice\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"}],\"name\":\"getKeySchema\",\"outputs\":[{\"internalType\":\"Schema\",\"name\":\"schema\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"Schema\",\"name\":\"schema\",\"type\":\"bytes32\"}],\"name\":\"getRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"}],\"name\":\"getRecord\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"}],\"name\":\"getSchema\",\"outputs\":[{\"internalType\":\"Schema\",\"name\":\"schema\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"}],\"name\":\"grantAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIModule\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"installModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIModule\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"installRootModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStore\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"byteLengthToPop\",\"type\":\"uint256\"}],\"name\":\"popFromField\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"byteLengthToPop\",\"type\":\"uint256\"}],\"name\":\"popFromField\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"dataToPush\",\"type\":\"bytes\"}],\"name\":\"pushToField\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"dataToPush\",\"type\":\"bytes\"}],\"name\":\"pushToField\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"string\",\"name\":\"systemFunctionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"systemFunctionArguments\",\"type\":\"string\"}],\"name\":\"registerFunctionSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"registerHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"}],\"name\":\"registerNamespace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes4\",\"name\":\"worldFunctionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"systemFunctionSelector\",\"type\":\"bytes4\"}],\"name\":\"registerRootFunctionSelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"Schema\",\"name\":\"schema\",\"type\":\"bytes32\"},{\"internalType\":\"Schema\",\"name\":\"keySchema\",\"type\":\"bytes32\"}],\"name\":\"registerSchema\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"contractIStoreHook\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"registerStoreHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"contractSystem\",\"name\":\"system\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"publicAccess\",\"type\":\"bool\"}],\"name\":\"registerSystem\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceSelector\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"contractISystemHook\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"registerSystemHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"Schema\",\"name\":\"valueSchema\",\"type\":\"bytes32\"},{\"internalType\":\"Schema\",\"name\":\"keySchema\",\"type\":\"bytes32\"}],\"name\":\"registerTable\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceSelector\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"contractIStoreHook\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"registerTableHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_oracleId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nbWords\",\"type\":\"uint32\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackSelector\",\"type\":\"bytes4\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"}],\"name\":\"revokeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setField\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setField\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"fieldNames\",\"type\":\"string[]\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"tableName\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"fieldNames\",\"type\":\"string[]\"}],\"name\":\"setMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startNewRaffle\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"table\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startByteIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dataToSet\",\"type\":\"bytes\"}],\"name\":\"updateInField\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"namespace\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"name\",\"type\":\"bytes16\"},{\"internalType\":\"bytes32[]\",\"name\":\"key\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"schemaIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startByteIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dataToSet\",\"type\":\"bytes\"}],\"name\":\"updateInField\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IWorldABI is the input ABI used to generate the binding from.
// Deprecated: Use IWorldMetaData.ABI instead.
var IWorldABI = IWorldMetaData.ABI

// IWorld is an auto generated Go binding around an Ethereum contract.
type IWorld struct {
	IWorldCaller     // Read-only binding to the contract
	IWorldTransactor // Write-only binding to the contract
	IWorldFilterer   // Log filterer for contract events
}

// IWorldCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWorldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWorldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWorldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWorldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWorldSession struct {
	Contract     *IWorld           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWorldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWorldCallerSession struct {
	Contract *IWorldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWorldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWorldTransactorSession struct {
	Contract     *IWorldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWorldRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWorldRaw struct {
	Contract *IWorld // Generic contract binding to access the raw methods on
}

// IWorldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWorldCallerRaw struct {
	Contract *IWorldCaller // Generic read-only contract binding to access the raw methods on
}

// IWorldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWorldTransactorRaw struct {
	Contract *IWorldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWorld creates a new instance of IWorld, bound to a specific deployed contract.
func NewIWorld(address common.Address, backend bind.ContractBackend) (*IWorld, error) {
	contract, err := bindIWorld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWorld{IWorldCaller: IWorldCaller{contract: contract}, IWorldTransactor: IWorldTransactor{contract: contract}, IWorldFilterer: IWorldFilterer{contract: contract}}, nil
}

// NewIWorldCaller creates a new read-only instance of IWorld, bound to a specific deployed contract.
func NewIWorldCaller(address common.Address, caller bind.ContractCaller) (*IWorldCaller, error) {
	contract, err := bindIWorld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWorldCaller{contract: contract}, nil
}

// NewIWorldTransactor creates a new write-only instance of IWorld, bound to a specific deployed contract.
func NewIWorldTransactor(address common.Address, transactor bind.ContractTransactor) (*IWorldTransactor, error) {
	contract, err := bindIWorld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWorldTransactor{contract: contract}, nil
}

// NewIWorldFilterer creates a new log filterer instance of IWorld, bound to a specific deployed contract.
func NewIWorldFilterer(address common.Address, filterer bind.ContractFilterer) (*IWorldFilterer, error) {
	contract, err := bindIWorld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWorldFilterer{contract: contract}, nil
}

// bindIWorld binds a generic wrapper to an already deployed contract.
func bindIWorld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IWorldMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorld *IWorldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWorld.Contract.IWorldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorld *IWorldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorld.Contract.IWorldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorld *IWorldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorld.Contract.IWorldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWorld *IWorldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWorld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWorld *IWorldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWorld *IWorldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWorld.Contract.contract.Transact(opts, method, params...)
}

// GetField is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 table, bytes32[] key, uint8 schemaIndex) view returns(bytes data)
func (_IWorld *IWorldCaller) GetField(opts *bind.CallOpts, table [32]byte, key [][32]byte, schemaIndex uint8) ([]byte, error) {
	var out []interface{}
	err := _IWorld.contract.Call(opts, &out, "getField", table, key, schemaIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetField is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 table, bytes32[] key, uint8 schemaIndex) view returns(bytes data)
func (_IWorld *IWorldSession) GetField(table [32]byte, key [][32]byte, schemaIndex uint8) ([]byte, error) {
	return _IWorld.Contract.GetField(&_IWorld.CallOpts, table, key, schemaIndex)
}

// GetField is a free data retrieval call binding the contract method 0xd03edb8c.
//
// Solidity: function getField(bytes32 table, bytes32[] key, uint8 schemaIndex) view returns(bytes data)
func (_IWorld *IWorldCallerSession) GetField(table [32]byte, key [][32]byte, schemaIndex uint8) ([]byte, error) {
	return _IWorld.Contract.GetField(&_IWorld.CallOpts, table, key, schemaIndex)
}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes32 schema) view returns(uint256)
func (_IWorld *IWorldCaller) GetFieldLength(opts *bind.CallOpts, table [32]byte, key [][32]byte, schemaIndex uint8, schema [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IWorld.contract.Call(opts, &out, "getFieldLength", table, key, schemaIndex, schema)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes32 schema) view returns(uint256)
func (_IWorld *IWorldSession) GetFieldLength(table [32]byte, key [][32]byte, schemaIndex uint8, schema [32]byte) (*big.Int, error) {
	return _IWorld.Contract.GetFieldLength(&_IWorld.CallOpts, table, key, schemaIndex, schema)
}

// GetFieldLength is a free data retrieval call binding the contract method 0x9f1fcf0a.
//
// Solidity: function getFieldLength(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes32 schema) view returns(uint256)
func (_IWorld *IWorldCallerSession) GetFieldLength(table [32]byte, key [][32]byte, schemaIndex uint8, schema [32]byte) (*big.Int, error) {
	return _IWorld.Contract.GetFieldLength(&_IWorld.CallOpts, table, key, schemaIndex, schema)
}

// GetFieldSlice is a free data retrieval call binding the contract method 0xd3a26b06.
//
// Solidity: function getFieldSlice(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes32 schema, uint256 start, uint256 end) view returns(bytes data)
func (_IWorld *IWorldCaller) GetFieldSlice(opts *bind.CallOpts, table [32]byte, key [][32]byte, schemaIndex uint8, schema [32]byte, start *big.Int, end *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IWorld.contract.Call(opts, &out, "getFieldSlice", table, key, schemaIndex, schema, start, end)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetFieldSlice is a free data retrieval call binding the contract method 0xd3a26b06.
//
// Solidity: function getFieldSlice(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes32 schema, uint256 start, uint256 end) view returns(bytes data)
func (_IWorld *IWorldSession) GetFieldSlice(table [32]byte, key [][32]byte, schemaIndex uint8, schema [32]byte, start *big.Int, end *big.Int) ([]byte, error) {
	return _IWorld.Contract.GetFieldSlice(&_IWorld.CallOpts, table, key, schemaIndex, schema, start, end)
}

// GetFieldSlice is a free data retrieval call binding the contract method 0xd3a26b06.
//
// Solidity: function getFieldSlice(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes32 schema, uint256 start, uint256 end) view returns(bytes data)
func (_IWorld *IWorldCallerSession) GetFieldSlice(table [32]byte, key [][32]byte, schemaIndex uint8, schema [32]byte, start *big.Int, end *big.Int) ([]byte, error) {
	return _IWorld.Contract.GetFieldSlice(&_IWorld.CallOpts, table, key, schemaIndex, schema, start, end)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 table) view returns(bytes32 schema)
func (_IWorld *IWorldCaller) GetKeySchema(opts *bind.CallOpts, table [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IWorld.contract.Call(opts, &out, "getKeySchema", table)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 table) view returns(bytes32 schema)
func (_IWorld *IWorldSession) GetKeySchema(table [32]byte) ([32]byte, error) {
	return _IWorld.Contract.GetKeySchema(&_IWorld.CallOpts, table)
}

// GetKeySchema is a free data retrieval call binding the contract method 0xd4285dc2.
//
// Solidity: function getKeySchema(bytes32 table) view returns(bytes32 schema)
func (_IWorld *IWorldCallerSession) GetKeySchema(table [32]byte) ([32]byte, error) {
	return _IWorld.Contract.GetKeySchema(&_IWorld.CallOpts, table)
}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 table, bytes32[] key, bytes32 schema) view returns(bytes data)
func (_IWorld *IWorldCaller) GetRecord(opts *bind.CallOpts, table [32]byte, key [][32]byte, schema [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IWorld.contract.Call(opts, &out, "getRecord", table, key, schema)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 table, bytes32[] key, bytes32 schema) view returns(bytes data)
func (_IWorld *IWorldSession) GetRecord(table [32]byte, key [][32]byte, schema [32]byte) ([]byte, error) {
	return _IWorld.Contract.GetRecord(&_IWorld.CallOpts, table, key, schema)
}

// GetRecord is a free data retrieval call binding the contract method 0x419b58fd.
//
// Solidity: function getRecord(bytes32 table, bytes32[] key, bytes32 schema) view returns(bytes data)
func (_IWorld *IWorldCallerSession) GetRecord(table [32]byte, key [][32]byte, schema [32]byte) ([]byte, error) {
	return _IWorld.Contract.GetRecord(&_IWorld.CallOpts, table, key, schema)
}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 table, bytes32[] key) view returns(bytes data)
func (_IWorld *IWorldCaller) GetRecord0(opts *bind.CallOpts, table [32]byte, key [][32]byte) ([]byte, error) {
	var out []interface{}
	err := _IWorld.contract.Call(opts, &out, "getRecord0", table, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 table, bytes32[] key) view returns(bytes data)
func (_IWorld *IWorldSession) GetRecord0(table [32]byte, key [][32]byte) ([]byte, error) {
	return _IWorld.Contract.GetRecord0(&_IWorld.CallOpts, table, key)
}

// GetRecord0 is a free data retrieval call binding the contract method 0xcc49db7e.
//
// Solidity: function getRecord(bytes32 table, bytes32[] key) view returns(bytes data)
func (_IWorld *IWorldCallerSession) GetRecord0(table [32]byte, key [][32]byte) ([]byte, error) {
	return _IWorld.Contract.GetRecord0(&_IWorld.CallOpts, table, key)
}

// GetSchema is a free data retrieval call binding the contract method 0xa2ea7c6e.
//
// Solidity: function getSchema(bytes32 table) view returns(bytes32 schema)
func (_IWorld *IWorldCaller) GetSchema(opts *bind.CallOpts, table [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IWorld.contract.Call(opts, &out, "getSchema", table)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSchema is a free data retrieval call binding the contract method 0xa2ea7c6e.
//
// Solidity: function getSchema(bytes32 table) view returns(bytes32 schema)
func (_IWorld *IWorldSession) GetSchema(table [32]byte) ([32]byte, error) {
	return _IWorld.Contract.GetSchema(&_IWorld.CallOpts, table)
}

// GetSchema is a free data retrieval call binding the contract method 0xa2ea7c6e.
//
// Solidity: function getSchema(bytes32 table) view returns(bytes32 schema)
func (_IWorld *IWorldCallerSession) GetSchema(table [32]byte) ([32]byte, error) {
	return _IWorld.Contract.GetSchema(&_IWorld.CallOpts, table)
}

// IsStore is a free data retrieval call binding the contract method 0xa5c2f007.
//
// Solidity: function isStore() view returns()
func (_IWorld *IWorldCaller) IsStore(opts *bind.CallOpts) error {
	var out []interface{}
	err := _IWorld.contract.Call(opts, &out, "isStore")

	if err != nil {
		return err
	}

	return err

}

// IsStore is a free data retrieval call binding the contract method 0xa5c2f007.
//
// Solidity: function isStore() view returns()
func (_IWorld *IWorldSession) IsStore() error {
	return _IWorld.Contract.IsStore(&_IWorld.CallOpts)
}

// IsStore is a free data retrieval call binding the contract method 0xa5c2f007.
//
// Solidity: function isStore() view returns()
func (_IWorld *IWorldCallerSession) IsStore() error {
	return _IWorld.Contract.IsStore(&_IWorld.CallOpts)
}

// BuyTicket is a paid mutator transaction binding the contract method 0x35477070.
//
// Solidity: function buyTicket(uint32 raffleId) returns()
func (_IWorld *IWorldTransactor) BuyTicket(opts *bind.TransactOpts, raffleId uint32) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "buyTicket", raffleId)
}

// BuyTicket is a paid mutator transaction binding the contract method 0x35477070.
//
// Solidity: function buyTicket(uint32 raffleId) returns()
func (_IWorld *IWorldSession) BuyTicket(raffleId uint32) (*types.Transaction, error) {
	return _IWorld.Contract.BuyTicket(&_IWorld.TransactOpts, raffleId)
}

// BuyTicket is a paid mutator transaction binding the contract method 0x35477070.
//
// Solidity: function buyTicket(uint32 raffleId) returns()
func (_IWorld *IWorldTransactorSession) BuyTicket(raffleId uint32) (*types.Transaction, error) {
	return _IWorld.Contract.BuyTicket(&_IWorld.TransactOpts, raffleId)
}

// Call is a paid mutator transaction binding the contract method 0x832f05f4.
//
// Solidity: function call(bytes16 namespace, bytes16 name, bytes funcSelectorAndArgs) payable returns(bytes)
func (_IWorld *IWorldTransactor) Call(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, funcSelectorAndArgs []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "call", namespace, name, funcSelectorAndArgs)
}

// Call is a paid mutator transaction binding the contract method 0x832f05f4.
//
// Solidity: function call(bytes16 namespace, bytes16 name, bytes funcSelectorAndArgs) payable returns(bytes)
func (_IWorld *IWorldSession) Call(namespace [16]byte, name [16]byte, funcSelectorAndArgs []byte) (*types.Transaction, error) {
	return _IWorld.Contract.Call(&_IWorld.TransactOpts, namespace, name, funcSelectorAndArgs)
}

// Call is a paid mutator transaction binding the contract method 0x832f05f4.
//
// Solidity: function call(bytes16 namespace, bytes16 name, bytes funcSelectorAndArgs) payable returns(bytes)
func (_IWorld *IWorldTransactorSession) Call(namespace [16]byte, name [16]byte, funcSelectorAndArgs []byte) (*types.Transaction, error) {
	return _IWorld.Contract.Call(&_IWorld.TransactOpts, namespace, name, funcSelectorAndArgs)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 table, bytes32[] key) returns()
func (_IWorld *IWorldTransactor) DeleteRecord(opts *bind.TransactOpts, table [32]byte, key [][32]byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "deleteRecord", table, key)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 table, bytes32[] key) returns()
func (_IWorld *IWorldSession) DeleteRecord(table [32]byte, key [][32]byte) (*types.Transaction, error) {
	return _IWorld.Contract.DeleteRecord(&_IWorld.TransactOpts, table, key)
}

// DeleteRecord is a paid mutator transaction binding the contract method 0x505a181d.
//
// Solidity: function deleteRecord(bytes32 table, bytes32[] key) returns()
func (_IWorld *IWorldTransactorSession) DeleteRecord(table [32]byte, key [][32]byte) (*types.Transaction, error) {
	return _IWorld.Contract.DeleteRecord(&_IWorld.TransactOpts, table, key)
}

// DeleteRecord0 is a paid mutator transaction binding the contract method 0xe3f8cc6b.
//
// Solidity: function deleteRecord(bytes16 namespace, bytes16 name, bytes32[] key) returns()
func (_IWorld *IWorldTransactor) DeleteRecord0(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, key [][32]byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "deleteRecord0", namespace, name, key)
}

// DeleteRecord0 is a paid mutator transaction binding the contract method 0xe3f8cc6b.
//
// Solidity: function deleteRecord(bytes16 namespace, bytes16 name, bytes32[] key) returns()
func (_IWorld *IWorldSession) DeleteRecord0(namespace [16]byte, name [16]byte, key [][32]byte) (*types.Transaction, error) {
	return _IWorld.Contract.DeleteRecord0(&_IWorld.TransactOpts, namespace, name, key)
}

// DeleteRecord0 is a paid mutator transaction binding the contract method 0xe3f8cc6b.
//
// Solidity: function deleteRecord(bytes16 namespace, bytes16 name, bytes32[] key) returns()
func (_IWorld *IWorldTransactorSession) DeleteRecord0(namespace [16]byte, name [16]byte, key [][32]byte) (*types.Transaction, error) {
	return _IWorld.Contract.DeleteRecord0(&_IWorld.TransactOpts, namespace, name, key)
}

// EmitEphemeralRecord is a paid mutator transaction binding the contract method 0xab7404bb.
//
// Solidity: function emitEphemeralRecord(bytes16 namespace, bytes16 name, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldTransactor) EmitEphemeralRecord(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "emitEphemeralRecord", namespace, name, key, data)
}

// EmitEphemeralRecord is a paid mutator transaction binding the contract method 0xab7404bb.
//
// Solidity: function emitEphemeralRecord(bytes16 namespace, bytes16 name, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldSession) EmitEphemeralRecord(namespace [16]byte, name [16]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.EmitEphemeralRecord(&_IWorld.TransactOpts, namespace, name, key, data)
}

// EmitEphemeralRecord is a paid mutator transaction binding the contract method 0xab7404bb.
//
// Solidity: function emitEphemeralRecord(bytes16 namespace, bytes16 name, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldTransactorSession) EmitEphemeralRecord(namespace [16]byte, name [16]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.EmitEphemeralRecord(&_IWorld.TransactOpts, namespace, name, key, data)
}

// EmitEphemeralRecord0 is a paid mutator transaction binding the contract method 0xbe83698c.
//
// Solidity: function emitEphemeralRecord(bytes32 table, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldTransactor) EmitEphemeralRecord0(opts *bind.TransactOpts, table [32]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "emitEphemeralRecord0", table, key, data)
}

// EmitEphemeralRecord0 is a paid mutator transaction binding the contract method 0xbe83698c.
//
// Solidity: function emitEphemeralRecord(bytes32 table, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldSession) EmitEphemeralRecord0(table [32]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.EmitEphemeralRecord0(&_IWorld.TransactOpts, table, key, data)
}

// EmitEphemeralRecord0 is a paid mutator transaction binding the contract method 0xbe83698c.
//
// Solidity: function emitEphemeralRecord(bytes32 table, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldTransactorSession) EmitEphemeralRecord0(table [32]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.EmitEphemeralRecord0(&_IWorld.TransactOpts, table, key, data)
}

// EndRaffle is a paid mutator transaction binding the contract method 0xcba8e862.
//
// Solidity: function endRaffle(uint32 raffleId) returns()
func (_IWorld *IWorldTransactor) EndRaffle(opts *bind.TransactOpts, raffleId uint32) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "endRaffle", raffleId)
}

// EndRaffle is a paid mutator transaction binding the contract method 0xcba8e862.
//
// Solidity: function endRaffle(uint32 raffleId) returns()
func (_IWorld *IWorldSession) EndRaffle(raffleId uint32) (*types.Transaction, error) {
	return _IWorld.Contract.EndRaffle(&_IWorld.TransactOpts, raffleId)
}

// EndRaffle is a paid mutator transaction binding the contract method 0xcba8e862.
//
// Solidity: function endRaffle(uint32 raffleId) returns()
func (_IWorld *IWorldTransactorSession) EndRaffle(raffleId uint32) (*types.Transaction, error) {
	return _IWorld.Contract.EndRaffle(&_IWorld.TransactOpts, raffleId)
}

// EndRaffleCallback is a paid mutator transaction binding the contract method 0x8c17d438.
//
// Solidity: function endRaffleCallback(bytes32 requestId, uint256[] randomWords) returns(uint256)
func (_IWorld *IWorldTransactor) EndRaffleCallback(opts *bind.TransactOpts, requestId [32]byte, randomWords []*big.Int) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "endRaffleCallback", requestId, randomWords)
}

// EndRaffleCallback is a paid mutator transaction binding the contract method 0x8c17d438.
//
// Solidity: function endRaffleCallback(bytes32 requestId, uint256[] randomWords) returns(uint256)
func (_IWorld *IWorldSession) EndRaffleCallback(requestId [32]byte, randomWords []*big.Int) (*types.Transaction, error) {
	return _IWorld.Contract.EndRaffleCallback(&_IWorld.TransactOpts, requestId, randomWords)
}

// EndRaffleCallback is a paid mutator transaction binding the contract method 0x8c17d438.
//
// Solidity: function endRaffleCallback(bytes32 requestId, uint256[] randomWords) returns(uint256)
func (_IWorld *IWorldTransactorSession) EndRaffleCallback(requestId [32]byte, randomWords []*big.Int) (*types.Transaction, error) {
	return _IWorld.Contract.EndRaffleCallback(&_IWorld.TransactOpts, requestId, randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x79930123.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (uint64,uint32,uint32,address,bytes4) _request) returns()
func (_IWorld *IWorldTransactor) FulfillRandomWords(opts *bind.TransactOpts, _proof VRFProof, _request VRFRequest) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "fulfillRandomWords", _proof, _request)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x79930123.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (uint64,uint32,uint32,address,bytes4) _request) returns()
func (_IWorld *IWorldSession) FulfillRandomWords(_proof VRFProof, _request VRFRequest) (*types.Transaction, error) {
	return _IWorld.Contract.FulfillRandomWords(&_IWorld.TransactOpts, _proof, _request)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x79930123.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (uint64,uint32,uint32,address,bytes4) _request) returns()
func (_IWorld *IWorldTransactorSession) FulfillRandomWords(_proof VRFProof, _request VRFRequest) (*types.Transaction, error) {
	return _IWorld.Contract.FulfillRandomWords(&_IWorld.TransactOpts, _proof, _request)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xf227e653.
//
// Solidity: function grantAccess(bytes16 namespace, bytes16 name, address grantee) returns()
func (_IWorld *IWorldTransactor) GrantAccess(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, grantee common.Address) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "grantAccess", namespace, name, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xf227e653.
//
// Solidity: function grantAccess(bytes16 namespace, bytes16 name, address grantee) returns()
func (_IWorld *IWorldSession) GrantAccess(namespace [16]byte, name [16]byte, grantee common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.GrantAccess(&_IWorld.TransactOpts, namespace, name, grantee)
}

// GrantAccess is a paid mutator transaction binding the contract method 0xf227e653.
//
// Solidity: function grantAccess(bytes16 namespace, bytes16 name, address grantee) returns()
func (_IWorld *IWorldTransactorSession) GrantAccess(namespace [16]byte, name [16]byte, grantee common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.GrantAccess(&_IWorld.TransactOpts, namespace, name, grantee)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes args) returns()
func (_IWorld *IWorldTransactor) InstallModule(opts *bind.TransactOpts, module common.Address, args []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "installModule", module, args)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes args) returns()
func (_IWorld *IWorldSession) InstallModule(module common.Address, args []byte) (*types.Transaction, error) {
	return _IWorld.Contract.InstallModule(&_IWorld.TransactOpts, module, args)
}

// InstallModule is a paid mutator transaction binding the contract method 0x8da798da.
//
// Solidity: function installModule(address module, bytes args) returns()
func (_IWorld *IWorldTransactorSession) InstallModule(module common.Address, args []byte) (*types.Transaction, error) {
	return _IWorld.Contract.InstallModule(&_IWorld.TransactOpts, module, args)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes args) returns()
func (_IWorld *IWorldTransactor) InstallRootModule(opts *bind.TransactOpts, module common.Address, args []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "installRootModule", module, args)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes args) returns()
func (_IWorld *IWorldSession) InstallRootModule(module common.Address, args []byte) (*types.Transaction, error) {
	return _IWorld.Contract.InstallRootModule(&_IWorld.TransactOpts, module, args)
}

// InstallRootModule is a paid mutator transaction binding the contract method 0xaf068c9e.
//
// Solidity: function installRootModule(address module, bytes args) returns()
func (_IWorld *IWorldTransactorSession) InstallRootModule(module common.Address, args []byte) (*types.Transaction, error) {
	return _IWorld.Contract.InstallRootModule(&_IWorld.TransactOpts, module, args)
}

// PopFromField is a paid mutator transaction binding the contract method 0x77841557.
//
// Solidity: function popFromField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, uint256 byteLengthToPop) returns()
func (_IWorld *IWorldTransactor) PopFromField(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "popFromField", namespace, name, key, schemaIndex, byteLengthToPop)
}

// PopFromField is a paid mutator transaction binding the contract method 0x77841557.
//
// Solidity: function popFromField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, uint256 byteLengthToPop) returns()
func (_IWorld *IWorldSession) PopFromField(namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _IWorld.Contract.PopFromField(&_IWorld.TransactOpts, namespace, name, key, schemaIndex, byteLengthToPop)
}

// PopFromField is a paid mutator transaction binding the contract method 0x77841557.
//
// Solidity: function popFromField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, uint256 byteLengthToPop) returns()
func (_IWorld *IWorldTransactorSession) PopFromField(namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _IWorld.Contract.PopFromField(&_IWorld.TransactOpts, namespace, name, key, schemaIndex, byteLengthToPop)
}

// PopFromField0 is a paid mutator transaction binding the contract method 0x8c1f9a54.
//
// Solidity: function popFromField(bytes32 table, bytes32[] key, uint8 schemaIndex, uint256 byteLengthToPop) returns()
func (_IWorld *IWorldTransactor) PopFromField0(opts *bind.TransactOpts, table [32]byte, key [][32]byte, schemaIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "popFromField0", table, key, schemaIndex, byteLengthToPop)
}

// PopFromField0 is a paid mutator transaction binding the contract method 0x8c1f9a54.
//
// Solidity: function popFromField(bytes32 table, bytes32[] key, uint8 schemaIndex, uint256 byteLengthToPop) returns()
func (_IWorld *IWorldSession) PopFromField0(table [32]byte, key [][32]byte, schemaIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _IWorld.Contract.PopFromField0(&_IWorld.TransactOpts, table, key, schemaIndex, byteLengthToPop)
}

// PopFromField0 is a paid mutator transaction binding the contract method 0x8c1f9a54.
//
// Solidity: function popFromField(bytes32 table, bytes32[] key, uint8 schemaIndex, uint256 byteLengthToPop) returns()
func (_IWorld *IWorldTransactorSession) PopFromField0(table [32]byte, key [][32]byte, schemaIndex uint8, byteLengthToPop *big.Int) (*types.Transaction, error) {
	return _IWorld.Contract.PopFromField0(&_IWorld.TransactOpts, table, key, schemaIndex, byteLengthToPop)
}

// PushToField is a paid mutator transaction binding the contract method 0x0c1a6000.
//
// Solidity: function pushToField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes dataToPush) returns()
func (_IWorld *IWorldTransactor) PushToField(opts *bind.TransactOpts, table [32]byte, key [][32]byte, schemaIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "pushToField", table, key, schemaIndex, dataToPush)
}

// PushToField is a paid mutator transaction binding the contract method 0x0c1a6000.
//
// Solidity: function pushToField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes dataToPush) returns()
func (_IWorld *IWorldSession) PushToField(table [32]byte, key [][32]byte, schemaIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _IWorld.Contract.PushToField(&_IWorld.TransactOpts, table, key, schemaIndex, dataToPush)
}

// PushToField is a paid mutator transaction binding the contract method 0x0c1a6000.
//
// Solidity: function pushToField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes dataToPush) returns()
func (_IWorld *IWorldTransactorSession) PushToField(table [32]byte, key [][32]byte, schemaIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _IWorld.Contract.PushToField(&_IWorld.TransactOpts, table, key, schemaIndex, dataToPush)
}

// PushToField0 is a paid mutator transaction binding the contract method 0x6e627cb8.
//
// Solidity: function pushToField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, bytes dataToPush) returns()
func (_IWorld *IWorldTransactor) PushToField0(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "pushToField0", namespace, name, key, schemaIndex, dataToPush)
}

// PushToField0 is a paid mutator transaction binding the contract method 0x6e627cb8.
//
// Solidity: function pushToField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, bytes dataToPush) returns()
func (_IWorld *IWorldSession) PushToField0(namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _IWorld.Contract.PushToField0(&_IWorld.TransactOpts, namespace, name, key, schemaIndex, dataToPush)
}

// PushToField0 is a paid mutator transaction binding the contract method 0x6e627cb8.
//
// Solidity: function pushToField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, bytes dataToPush) returns()
func (_IWorld *IWorldTransactorSession) PushToField0(namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, dataToPush []byte) (*types.Transaction, error) {
	return _IWorld.Contract.PushToField0(&_IWorld.TransactOpts, namespace, name, key, schemaIndex, dataToPush)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0xca4c1f5e.
//
// Solidity: function registerFunctionSelector(bytes16 namespace, bytes16 name, string systemFunctionName, string systemFunctionArguments) returns(bytes4 worldFunctionSelector)
func (_IWorld *IWorldTransactor) RegisterFunctionSelector(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, systemFunctionName string, systemFunctionArguments string) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerFunctionSelector", namespace, name, systemFunctionName, systemFunctionArguments)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0xca4c1f5e.
//
// Solidity: function registerFunctionSelector(bytes16 namespace, bytes16 name, string systemFunctionName, string systemFunctionArguments) returns(bytes4 worldFunctionSelector)
func (_IWorld *IWorldSession) RegisterFunctionSelector(namespace [16]byte, name [16]byte, systemFunctionName string, systemFunctionArguments string) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterFunctionSelector(&_IWorld.TransactOpts, namespace, name, systemFunctionName, systemFunctionArguments)
}

// RegisterFunctionSelector is a paid mutator transaction binding the contract method 0xca4c1f5e.
//
// Solidity: function registerFunctionSelector(bytes16 namespace, bytes16 name, string systemFunctionName, string systemFunctionArguments) returns(bytes4 worldFunctionSelector)
func (_IWorld *IWorldTransactorSession) RegisterFunctionSelector(namespace [16]byte, name [16]byte, systemFunctionName string, systemFunctionArguments string) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterFunctionSelector(&_IWorld.TransactOpts, namespace, name, systemFunctionName, systemFunctionArguments)
}

// RegisterHook is a paid mutator transaction binding the contract method 0x7a58dce4.
//
// Solidity: function registerHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldTransactor) RegisterHook(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerHook", namespace, name, hook)
}

// RegisterHook is a paid mutator transaction binding the contract method 0x7a58dce4.
//
// Solidity: function registerHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldSession) RegisterHook(namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterHook(&_IWorld.TransactOpts, namespace, name, hook)
}

// RegisterHook is a paid mutator transaction binding the contract method 0x7a58dce4.
//
// Solidity: function registerHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldTransactorSession) RegisterHook(namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterHook(&_IWorld.TransactOpts, namespace, name, hook)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xa886545e.
//
// Solidity: function registerNamespace(bytes16 namespace) returns()
func (_IWorld *IWorldTransactor) RegisterNamespace(opts *bind.TransactOpts, namespace [16]byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerNamespace", namespace)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xa886545e.
//
// Solidity: function registerNamespace(bytes16 namespace) returns()
func (_IWorld *IWorldSession) RegisterNamespace(namespace [16]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterNamespace(&_IWorld.TransactOpts, namespace)
}

// RegisterNamespace is a paid mutator transaction binding the contract method 0xa886545e.
//
// Solidity: function registerNamespace(bytes16 namespace) returns()
func (_IWorld *IWorldTransactorSession) RegisterNamespace(namespace [16]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterNamespace(&_IWorld.TransactOpts, namespace)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x56f55a97.
//
// Solidity: function registerRootFunctionSelector(bytes16 namespace, bytes16 name, bytes4 worldFunctionSelector, bytes4 systemFunctionSelector) returns(bytes4)
func (_IWorld *IWorldTransactor) RegisterRootFunctionSelector(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, worldFunctionSelector [4]byte, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerRootFunctionSelector", namespace, name, worldFunctionSelector, systemFunctionSelector)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x56f55a97.
//
// Solidity: function registerRootFunctionSelector(bytes16 namespace, bytes16 name, bytes4 worldFunctionSelector, bytes4 systemFunctionSelector) returns(bytes4)
func (_IWorld *IWorldSession) RegisterRootFunctionSelector(namespace [16]byte, name [16]byte, worldFunctionSelector [4]byte, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterRootFunctionSelector(&_IWorld.TransactOpts, namespace, name, worldFunctionSelector, systemFunctionSelector)
}

// RegisterRootFunctionSelector is a paid mutator transaction binding the contract method 0x56f55a97.
//
// Solidity: function registerRootFunctionSelector(bytes16 namespace, bytes16 name, bytes4 worldFunctionSelector, bytes4 systemFunctionSelector) returns(bytes4)
func (_IWorld *IWorldTransactorSession) RegisterRootFunctionSelector(namespace [16]byte, name [16]byte, worldFunctionSelector [4]byte, systemFunctionSelector [4]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterRootFunctionSelector(&_IWorld.TransactOpts, namespace, name, worldFunctionSelector, systemFunctionSelector)
}

// RegisterSchema is a paid mutator transaction binding the contract method 0xd5933686.
//
// Solidity: function registerSchema(bytes32 table, bytes32 schema, bytes32 keySchema) returns()
func (_IWorld *IWorldTransactor) RegisterSchema(opts *bind.TransactOpts, table [32]byte, schema [32]byte, keySchema [32]byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerSchema", table, schema, keySchema)
}

// RegisterSchema is a paid mutator transaction binding the contract method 0xd5933686.
//
// Solidity: function registerSchema(bytes32 table, bytes32 schema, bytes32 keySchema) returns()
func (_IWorld *IWorldSession) RegisterSchema(table [32]byte, schema [32]byte, keySchema [32]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterSchema(&_IWorld.TransactOpts, table, schema, keySchema)
}

// RegisterSchema is a paid mutator transaction binding the contract method 0xd5933686.
//
// Solidity: function registerSchema(bytes32 table, bytes32 schema, bytes32 keySchema) returns()
func (_IWorld *IWorldTransactorSession) RegisterSchema(table [32]byte, schema [32]byte, keySchema [32]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterSchema(&_IWorld.TransactOpts, table, schema, keySchema)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x6e81fd71.
//
// Solidity: function registerStoreHook(bytes32 table, address hook) returns()
func (_IWorld *IWorldTransactor) RegisterStoreHook(opts *bind.TransactOpts, table [32]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerStoreHook", table, hook)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x6e81fd71.
//
// Solidity: function registerStoreHook(bytes32 table, address hook) returns()
func (_IWorld *IWorldSession) RegisterStoreHook(table [32]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterStoreHook(&_IWorld.TransactOpts, table, hook)
}

// RegisterStoreHook is a paid mutator transaction binding the contract method 0x6e81fd71.
//
// Solidity: function registerStoreHook(bytes32 table, address hook) returns()
func (_IWorld *IWorldTransactorSession) RegisterStoreHook(table [32]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterStoreHook(&_IWorld.TransactOpts, table, hook)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x4c9e9f11.
//
// Solidity: function registerSystem(bytes16 namespace, bytes16 name, address system, bool publicAccess) returns(bytes32 resourceSelector)
func (_IWorld *IWorldTransactor) RegisterSystem(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerSystem", namespace, name, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x4c9e9f11.
//
// Solidity: function registerSystem(bytes16 namespace, bytes16 name, address system, bool publicAccess) returns(bytes32 resourceSelector)
func (_IWorld *IWorldSession) RegisterSystem(namespace [16]byte, name [16]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterSystem(&_IWorld.TransactOpts, namespace, name, system, publicAccess)
}

// RegisterSystem is a paid mutator transaction binding the contract method 0x4c9e9f11.
//
// Solidity: function registerSystem(bytes16 namespace, bytes16 name, address system, bool publicAccess) returns(bytes32 resourceSelector)
func (_IWorld *IWorldTransactorSession) RegisterSystem(namespace [16]byte, name [16]byte, system common.Address, publicAccess bool) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterSystem(&_IWorld.TransactOpts, namespace, name, system, publicAccess)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0x3ca90f22.
//
// Solidity: function registerSystemHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldTransactor) RegisterSystemHook(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerSystemHook", namespace, name, hook)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0x3ca90f22.
//
// Solidity: function registerSystemHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldSession) RegisterSystemHook(namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterSystemHook(&_IWorld.TransactOpts, namespace, name, hook)
}

// RegisterSystemHook is a paid mutator transaction binding the contract method 0x3ca90f22.
//
// Solidity: function registerSystemHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldTransactorSession) RegisterSystemHook(namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterSystemHook(&_IWorld.TransactOpts, namespace, name, hook)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x31b99d89.
//
// Solidity: function registerTable(bytes16 namespace, bytes16 name, bytes32 valueSchema, bytes32 keySchema) returns(bytes32 resourceSelector)
func (_IWorld *IWorldTransactor) RegisterTable(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, valueSchema [32]byte, keySchema [32]byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerTable", namespace, name, valueSchema, keySchema)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x31b99d89.
//
// Solidity: function registerTable(bytes16 namespace, bytes16 name, bytes32 valueSchema, bytes32 keySchema) returns(bytes32 resourceSelector)
func (_IWorld *IWorldSession) RegisterTable(namespace [16]byte, name [16]byte, valueSchema [32]byte, keySchema [32]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterTable(&_IWorld.TransactOpts, namespace, name, valueSchema, keySchema)
}

// RegisterTable is a paid mutator transaction binding the contract method 0x31b99d89.
//
// Solidity: function registerTable(bytes16 namespace, bytes16 name, bytes32 valueSchema, bytes32 keySchema) returns(bytes32 resourceSelector)
func (_IWorld *IWorldTransactorSession) RegisterTable(namespace [16]byte, name [16]byte, valueSchema [32]byte, keySchema [32]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterTable(&_IWorld.TransactOpts, namespace, name, valueSchema, keySchema)
}

// RegisterTableHook is a paid mutator transaction binding the contract method 0x8ae710e9.
//
// Solidity: function registerTableHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldTransactor) RegisterTableHook(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "registerTableHook", namespace, name, hook)
}

// RegisterTableHook is a paid mutator transaction binding the contract method 0x8ae710e9.
//
// Solidity: function registerTableHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldSession) RegisterTableHook(namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterTableHook(&_IWorld.TransactOpts, namespace, name, hook)
}

// RegisterTableHook is a paid mutator transaction binding the contract method 0x8ae710e9.
//
// Solidity: function registerTableHook(bytes16 namespace, bytes16 name, address hook) returns()
func (_IWorld *IWorldTransactorSession) RegisterTableHook(namespace [16]byte, name [16]byte, hook common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RegisterTableHook(&_IWorld.TransactOpts, namespace, name, hook)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5ed71610.
//
// Solidity: function requestRandomWords(bytes32 _oracleId, uint16 _requestConfirmations, uint32 _callbackGasLimit, uint32 _nbWords, bytes4 _callbackSelector) returns(bytes32)
func (_IWorld *IWorldTransactor) RequestRandomWords(opts *bind.TransactOpts, _oracleId [32]byte, _requestConfirmations uint16, _callbackGasLimit uint32, _nbWords uint32, _callbackSelector [4]byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "requestRandomWords", _oracleId, _requestConfirmations, _callbackGasLimit, _nbWords, _callbackSelector)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5ed71610.
//
// Solidity: function requestRandomWords(bytes32 _oracleId, uint16 _requestConfirmations, uint32 _callbackGasLimit, uint32 _nbWords, bytes4 _callbackSelector) returns(bytes32)
func (_IWorld *IWorldSession) RequestRandomWords(_oracleId [32]byte, _requestConfirmations uint16, _callbackGasLimit uint32, _nbWords uint32, _callbackSelector [4]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RequestRandomWords(&_IWorld.TransactOpts, _oracleId, _requestConfirmations, _callbackGasLimit, _nbWords, _callbackSelector)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5ed71610.
//
// Solidity: function requestRandomWords(bytes32 _oracleId, uint16 _requestConfirmations, uint32 _callbackGasLimit, uint32 _nbWords, bytes4 _callbackSelector) returns(bytes32)
func (_IWorld *IWorldTransactorSession) RequestRandomWords(_oracleId [32]byte, _requestConfirmations uint16, _callbackGasLimit uint32, _nbWords uint32, _callbackSelector [4]byte) (*types.Transaction, error) {
	return _IWorld.Contract.RequestRandomWords(&_IWorld.TransactOpts, _oracleId, _requestConfirmations, _callbackGasLimit, _nbWords, _callbackSelector)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x1e42269e.
//
// Solidity: function revokeAccess(bytes16 namespace, bytes16 name, address grantee) returns()
func (_IWorld *IWorldTransactor) RevokeAccess(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, grantee common.Address) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "revokeAccess", namespace, name, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x1e42269e.
//
// Solidity: function revokeAccess(bytes16 namespace, bytes16 name, address grantee) returns()
func (_IWorld *IWorldSession) RevokeAccess(namespace [16]byte, name [16]byte, grantee common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RevokeAccess(&_IWorld.TransactOpts, namespace, name, grantee)
}

// RevokeAccess is a paid mutator transaction binding the contract method 0x1e42269e.
//
// Solidity: function revokeAccess(bytes16 namespace, bytes16 name, address grantee) returns()
func (_IWorld *IWorldTransactorSession) RevokeAccess(namespace [16]byte, name [16]byte, grantee common.Address) (*types.Transaction, error) {
	return _IWorld.Contract.RevokeAccess(&_IWorld.TransactOpts, namespace, name, grantee)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes data) returns()
func (_IWorld *IWorldTransactor) SetField(opts *bind.TransactOpts, table [32]byte, key [][32]byte, schemaIndex uint8, data []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "setField", table, key, schemaIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes data) returns()
func (_IWorld *IWorldSession) SetField(table [32]byte, key [][32]byte, schemaIndex uint8, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.SetField(&_IWorld.TransactOpts, table, key, schemaIndex, data)
}

// SetField is a paid mutator transaction binding the contract method 0x114a7266.
//
// Solidity: function setField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes data) returns()
func (_IWorld *IWorldTransactorSession) SetField(table [32]byte, key [][32]byte, schemaIndex uint8, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.SetField(&_IWorld.TransactOpts, table, key, schemaIndex, data)
}

// SetField0 is a paid mutator transaction binding the contract method 0x9146e7c8.
//
// Solidity: function setField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, bytes data) returns()
func (_IWorld *IWorldTransactor) SetField0(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, data []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "setField0", namespace, name, key, schemaIndex, data)
}

// SetField0 is a paid mutator transaction binding the contract method 0x9146e7c8.
//
// Solidity: function setField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, bytes data) returns()
func (_IWorld *IWorldSession) SetField0(namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.SetField0(&_IWorld.TransactOpts, namespace, name, key, schemaIndex, data)
}

// SetField0 is a paid mutator transaction binding the contract method 0x9146e7c8.
//
// Solidity: function setField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, bytes data) returns()
func (_IWorld *IWorldTransactorSession) SetField0(namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.SetField0(&_IWorld.TransactOpts, namespace, name, key, schemaIndex, data)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x096fcaa9.
//
// Solidity: function setMetadata(bytes16 namespace, bytes16 name, string tableName, string[] fieldNames) returns()
func (_IWorld *IWorldTransactor) SetMetadata(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, tableName string, fieldNames []string) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "setMetadata", namespace, name, tableName, fieldNames)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x096fcaa9.
//
// Solidity: function setMetadata(bytes16 namespace, bytes16 name, string tableName, string[] fieldNames) returns()
func (_IWorld *IWorldSession) SetMetadata(namespace [16]byte, name [16]byte, tableName string, fieldNames []string) (*types.Transaction, error) {
	return _IWorld.Contract.SetMetadata(&_IWorld.TransactOpts, namespace, name, tableName, fieldNames)
}

// SetMetadata is a paid mutator transaction binding the contract method 0x096fcaa9.
//
// Solidity: function setMetadata(bytes16 namespace, bytes16 name, string tableName, string[] fieldNames) returns()
func (_IWorld *IWorldTransactorSession) SetMetadata(namespace [16]byte, name [16]byte, tableName string, fieldNames []string) (*types.Transaction, error) {
	return _IWorld.Contract.SetMetadata(&_IWorld.TransactOpts, namespace, name, tableName, fieldNames)
}

// SetMetadata0 is a paid mutator transaction binding the contract method 0xfe86a89f.
//
// Solidity: function setMetadata(bytes32 table, string tableName, string[] fieldNames) returns()
func (_IWorld *IWorldTransactor) SetMetadata0(opts *bind.TransactOpts, table [32]byte, tableName string, fieldNames []string) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "setMetadata0", table, tableName, fieldNames)
}

// SetMetadata0 is a paid mutator transaction binding the contract method 0xfe86a89f.
//
// Solidity: function setMetadata(bytes32 table, string tableName, string[] fieldNames) returns()
func (_IWorld *IWorldSession) SetMetadata0(table [32]byte, tableName string, fieldNames []string) (*types.Transaction, error) {
	return _IWorld.Contract.SetMetadata0(&_IWorld.TransactOpts, table, tableName, fieldNames)
}

// SetMetadata0 is a paid mutator transaction binding the contract method 0xfe86a89f.
//
// Solidity: function setMetadata(bytes32 table, string tableName, string[] fieldNames) returns()
func (_IWorld *IWorldTransactorSession) SetMetadata0(table [32]byte, tableName string, fieldNames []string) (*types.Transaction, error) {
	return _IWorld.Contract.SetMetadata0(&_IWorld.TransactOpts, table, tableName, fieldNames)
}

// SetRecord is a paid mutator transaction binding the contract method 0x11c73ffb.
//
// Solidity: function setRecord(bytes16 namespace, bytes16 name, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldTransactor) SetRecord(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "setRecord", namespace, name, key, data)
}

// SetRecord is a paid mutator transaction binding the contract method 0x11c73ffb.
//
// Solidity: function setRecord(bytes16 namespace, bytes16 name, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldSession) SetRecord(namespace [16]byte, name [16]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.SetRecord(&_IWorld.TransactOpts, namespace, name, key, data)
}

// SetRecord is a paid mutator transaction binding the contract method 0x11c73ffb.
//
// Solidity: function setRecord(bytes16 namespace, bytes16 name, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldTransactorSession) SetRecord(namespace [16]byte, name [16]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.SetRecord(&_IWorld.TransactOpts, namespace, name, key, data)
}

// SetRecord0 is a paid mutator transaction binding the contract method 0x8ed41f2f.
//
// Solidity: function setRecord(bytes32 table, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldTransactor) SetRecord0(opts *bind.TransactOpts, table [32]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "setRecord0", table, key, data)
}

// SetRecord0 is a paid mutator transaction binding the contract method 0x8ed41f2f.
//
// Solidity: function setRecord(bytes32 table, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldSession) SetRecord0(table [32]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.SetRecord0(&_IWorld.TransactOpts, table, key, data)
}

// SetRecord0 is a paid mutator transaction binding the contract method 0x8ed41f2f.
//
// Solidity: function setRecord(bytes32 table, bytes32[] key, bytes data) returns()
func (_IWorld *IWorldTransactorSession) SetRecord0(table [32]byte, key [][32]byte, data []byte) (*types.Transaction, error) {
	return _IWorld.Contract.SetRecord0(&_IWorld.TransactOpts, table, key, data)
}

// StartNewRaffle is a paid mutator transaction binding the contract method 0xa91f3e56.
//
// Solidity: function startNewRaffle() returns(uint32)
func (_IWorld *IWorldTransactor) StartNewRaffle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "startNewRaffle")
}

// StartNewRaffle is a paid mutator transaction binding the contract method 0xa91f3e56.
//
// Solidity: function startNewRaffle() returns(uint32)
func (_IWorld *IWorldSession) StartNewRaffle() (*types.Transaction, error) {
	return _IWorld.Contract.StartNewRaffle(&_IWorld.TransactOpts)
}

// StartNewRaffle is a paid mutator transaction binding the contract method 0xa91f3e56.
//
// Solidity: function startNewRaffle() returns(uint32)
func (_IWorld *IWorldTransactorSession) StartNewRaffle() (*types.Transaction, error) {
	return _IWorld.Contract.StartNewRaffle(&_IWorld.TransactOpts)
}

// UpdateInField is a paid mutator transaction binding the contract method 0x776d4047.
//
// Solidity: function updateInField(bytes32 table, bytes32[] key, uint8 schemaIndex, uint256 startByteIndex, bytes dataToSet) returns()
func (_IWorld *IWorldTransactor) UpdateInField(opts *bind.TransactOpts, table [32]byte, key [][32]byte, schemaIndex uint8, startByteIndex *big.Int, dataToSet []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "updateInField", table, key, schemaIndex, startByteIndex, dataToSet)
}

// UpdateInField is a paid mutator transaction binding the contract method 0x776d4047.
//
// Solidity: function updateInField(bytes32 table, bytes32[] key, uint8 schemaIndex, uint256 startByteIndex, bytes dataToSet) returns()
func (_IWorld *IWorldSession) UpdateInField(table [32]byte, key [][32]byte, schemaIndex uint8, startByteIndex *big.Int, dataToSet []byte) (*types.Transaction, error) {
	return _IWorld.Contract.UpdateInField(&_IWorld.TransactOpts, table, key, schemaIndex, startByteIndex, dataToSet)
}

// UpdateInField is a paid mutator transaction binding the contract method 0x776d4047.
//
// Solidity: function updateInField(bytes32 table, bytes32[] key, uint8 schemaIndex, uint256 startByteIndex, bytes dataToSet) returns()
func (_IWorld *IWorldTransactorSession) UpdateInField(table [32]byte, key [][32]byte, schemaIndex uint8, startByteIndex *big.Int, dataToSet []byte) (*types.Transaction, error) {
	return _IWorld.Contract.UpdateInField(&_IWorld.TransactOpts, table, key, schemaIndex, startByteIndex, dataToSet)
}

// UpdateInField0 is a paid mutator transaction binding the contract method 0xd491953a.
//
// Solidity: function updateInField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, uint256 startByteIndex, bytes dataToSet) returns()
func (_IWorld *IWorldTransactor) UpdateInField0(opts *bind.TransactOpts, namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, startByteIndex *big.Int, dataToSet []byte) (*types.Transaction, error) {
	return _IWorld.contract.Transact(opts, "updateInField0", namespace, name, key, schemaIndex, startByteIndex, dataToSet)
}

// UpdateInField0 is a paid mutator transaction binding the contract method 0xd491953a.
//
// Solidity: function updateInField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, uint256 startByteIndex, bytes dataToSet) returns()
func (_IWorld *IWorldSession) UpdateInField0(namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, startByteIndex *big.Int, dataToSet []byte) (*types.Transaction, error) {
	return _IWorld.Contract.UpdateInField0(&_IWorld.TransactOpts, namespace, name, key, schemaIndex, startByteIndex, dataToSet)
}

// UpdateInField0 is a paid mutator transaction binding the contract method 0xd491953a.
//
// Solidity: function updateInField(bytes16 namespace, bytes16 name, bytes32[] key, uint8 schemaIndex, uint256 startByteIndex, bytes dataToSet) returns()
func (_IWorld *IWorldTransactorSession) UpdateInField0(namespace [16]byte, name [16]byte, key [][32]byte, schemaIndex uint8, startByteIndex *big.Int, dataToSet []byte) (*types.Transaction, error) {
	return _IWorld.Contract.UpdateInField0(&_IWorld.TransactOpts, namespace, name, key, schemaIndex, startByteIndex, dataToSet)
}

// IWorldHelloWorldIterator is returned from FilterHelloWorld and is used to iterate over the raw logs and unpacked data for HelloWorld events raised by the IWorld contract.
type IWorldHelloWorldIterator struct {
	Event *IWorldHelloWorld // Event containing the contract specifics and raw log

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
func (it *IWorldHelloWorldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorldHelloWorld)
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
		it.Event = new(IWorldHelloWorld)
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
func (it *IWorldHelloWorldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorldHelloWorldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorldHelloWorld represents a HelloWorld event raised by the IWorld contract.
type IWorldHelloWorld struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHelloWorld is a free log retrieval operation binding the contract event 0x7fffb7bdf7d16635144da549e9a4eedff43ed43d64e49e18d7e365f9e5521232.
//
// Solidity: event HelloWorld()
func (_IWorld *IWorldFilterer) FilterHelloWorld(opts *bind.FilterOpts) (*IWorldHelloWorldIterator, error) {

	logs, sub, err := _IWorld.contract.FilterLogs(opts, "HelloWorld")
	if err != nil {
		return nil, err
	}
	return &IWorldHelloWorldIterator{contract: _IWorld.contract, event: "HelloWorld", logs: logs, sub: sub}, nil
}

// WatchHelloWorld is a free log subscription operation binding the contract event 0x7fffb7bdf7d16635144da549e9a4eedff43ed43d64e49e18d7e365f9e5521232.
//
// Solidity: event HelloWorld()
func (_IWorld *IWorldFilterer) WatchHelloWorld(opts *bind.WatchOpts, sink chan<- *IWorldHelloWorld) (event.Subscription, error) {

	logs, sub, err := _IWorld.contract.WatchLogs(opts, "HelloWorld")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorldHelloWorld)
				if err := _IWorld.contract.UnpackLog(event, "HelloWorld", log); err != nil {
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

// ParseHelloWorld is a log parse operation binding the contract event 0x7fffb7bdf7d16635144da549e9a4eedff43ed43d64e49e18d7e365f9e5521232.
//
// Solidity: event HelloWorld()
func (_IWorld *IWorldFilterer) ParseHelloWorld(log types.Log) (*IWorldHelloWorld, error) {
	event := new(IWorldHelloWorld)
	if err := _IWorld.contract.UnpackLog(event, "HelloWorld", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorldStoreDeleteRecordIterator is returned from FilterStoreDeleteRecord and is used to iterate over the raw logs and unpacked data for StoreDeleteRecord events raised by the IWorld contract.
type IWorldStoreDeleteRecordIterator struct {
	Event *IWorldStoreDeleteRecord // Event containing the contract specifics and raw log

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
func (it *IWorldStoreDeleteRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorldStoreDeleteRecord)
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
		it.Event = new(IWorldStoreDeleteRecord)
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
func (it *IWorldStoreDeleteRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorldStoreDeleteRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorldStoreDeleteRecord represents a StoreDeleteRecord event raised by the IWorld contract.
type IWorldStoreDeleteRecord struct {
	Table [32]byte
	Key   [][32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStoreDeleteRecord is a free log retrieval operation binding the contract event 0x2cc8610b80ef19409ae51ecbdd9c137960fb22ae9ef2d817d36ec1b685d68ecd.
//
// Solidity: event StoreDeleteRecord(bytes32 table, bytes32[] key)
func (_IWorld *IWorldFilterer) FilterStoreDeleteRecord(opts *bind.FilterOpts) (*IWorldStoreDeleteRecordIterator, error) {

	logs, sub, err := _IWorld.contract.FilterLogs(opts, "StoreDeleteRecord")
	if err != nil {
		return nil, err
	}
	return &IWorldStoreDeleteRecordIterator{contract: _IWorld.contract, event: "StoreDeleteRecord", logs: logs, sub: sub}, nil
}

// WatchStoreDeleteRecord is a free log subscription operation binding the contract event 0x2cc8610b80ef19409ae51ecbdd9c137960fb22ae9ef2d817d36ec1b685d68ecd.
//
// Solidity: event StoreDeleteRecord(bytes32 table, bytes32[] key)
func (_IWorld *IWorldFilterer) WatchStoreDeleteRecord(opts *bind.WatchOpts, sink chan<- *IWorldStoreDeleteRecord) (event.Subscription, error) {

	logs, sub, err := _IWorld.contract.WatchLogs(opts, "StoreDeleteRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorldStoreDeleteRecord)
				if err := _IWorld.contract.UnpackLog(event, "StoreDeleteRecord", log); err != nil {
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

// ParseStoreDeleteRecord is a log parse operation binding the contract event 0x2cc8610b80ef19409ae51ecbdd9c137960fb22ae9ef2d817d36ec1b685d68ecd.
//
// Solidity: event StoreDeleteRecord(bytes32 table, bytes32[] key)
func (_IWorld *IWorldFilterer) ParseStoreDeleteRecord(log types.Log) (*IWorldStoreDeleteRecord, error) {
	event := new(IWorldStoreDeleteRecord)
	if err := _IWorld.contract.UnpackLog(event, "StoreDeleteRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorldStoreEphemeralRecordIterator is returned from FilterStoreEphemeralRecord and is used to iterate over the raw logs and unpacked data for StoreEphemeralRecord events raised by the IWorld contract.
type IWorldStoreEphemeralRecordIterator struct {
	Event *IWorldStoreEphemeralRecord // Event containing the contract specifics and raw log

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
func (it *IWorldStoreEphemeralRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorldStoreEphemeralRecord)
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
		it.Event = new(IWorldStoreEphemeralRecord)
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
func (it *IWorldStoreEphemeralRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorldStoreEphemeralRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorldStoreEphemeralRecord represents a StoreEphemeralRecord event raised by the IWorld contract.
type IWorldStoreEphemeralRecord struct {
	Table [32]byte
	Key   [][32]byte
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStoreEphemeralRecord is a free log retrieval operation binding the contract event 0x230ea578f076e518b96d8b1b4ca2e5dc4898fd02d11e43c9f067c22d1c391d3b.
//
// Solidity: event StoreEphemeralRecord(bytes32 table, bytes32[] key, bytes data)
func (_IWorld *IWorldFilterer) FilterStoreEphemeralRecord(opts *bind.FilterOpts) (*IWorldStoreEphemeralRecordIterator, error) {

	logs, sub, err := _IWorld.contract.FilterLogs(opts, "StoreEphemeralRecord")
	if err != nil {
		return nil, err
	}
	return &IWorldStoreEphemeralRecordIterator{contract: _IWorld.contract, event: "StoreEphemeralRecord", logs: logs, sub: sub}, nil
}

// WatchStoreEphemeralRecord is a free log subscription operation binding the contract event 0x230ea578f076e518b96d8b1b4ca2e5dc4898fd02d11e43c9f067c22d1c391d3b.
//
// Solidity: event StoreEphemeralRecord(bytes32 table, bytes32[] key, bytes data)
func (_IWorld *IWorldFilterer) WatchStoreEphemeralRecord(opts *bind.WatchOpts, sink chan<- *IWorldStoreEphemeralRecord) (event.Subscription, error) {

	logs, sub, err := _IWorld.contract.WatchLogs(opts, "StoreEphemeralRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorldStoreEphemeralRecord)
				if err := _IWorld.contract.UnpackLog(event, "StoreEphemeralRecord", log); err != nil {
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

// ParseStoreEphemeralRecord is a log parse operation binding the contract event 0x230ea578f076e518b96d8b1b4ca2e5dc4898fd02d11e43c9f067c22d1c391d3b.
//
// Solidity: event StoreEphemeralRecord(bytes32 table, bytes32[] key, bytes data)
func (_IWorld *IWorldFilterer) ParseStoreEphemeralRecord(log types.Log) (*IWorldStoreEphemeralRecord, error) {
	event := new(IWorldStoreEphemeralRecord)
	if err := _IWorld.contract.UnpackLog(event, "StoreEphemeralRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorldStoreSetFieldIterator is returned from FilterStoreSetField and is used to iterate over the raw logs and unpacked data for StoreSetField events raised by the IWorld contract.
type IWorldStoreSetFieldIterator struct {
	Event *IWorldStoreSetField // Event containing the contract specifics and raw log

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
func (it *IWorldStoreSetFieldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorldStoreSetField)
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
		it.Event = new(IWorldStoreSetField)
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
func (it *IWorldStoreSetFieldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorldStoreSetFieldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorldStoreSetField represents a StoreSetField event raised by the IWorld contract.
type IWorldStoreSetField struct {
	Table       [32]byte
	Key         [][32]byte
	SchemaIndex uint8
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStoreSetField is a free log retrieval operation binding the contract event 0xd01f9f1368f831528fc9fe6442366b2b7d957fbfff3bcf7c24d9ab5fe51f8c46.
//
// Solidity: event StoreSetField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes data)
func (_IWorld *IWorldFilterer) FilterStoreSetField(opts *bind.FilterOpts) (*IWorldStoreSetFieldIterator, error) {

	logs, sub, err := _IWorld.contract.FilterLogs(opts, "StoreSetField")
	if err != nil {
		return nil, err
	}
	return &IWorldStoreSetFieldIterator{contract: _IWorld.contract, event: "StoreSetField", logs: logs, sub: sub}, nil
}

// WatchStoreSetField is a free log subscription operation binding the contract event 0xd01f9f1368f831528fc9fe6442366b2b7d957fbfff3bcf7c24d9ab5fe51f8c46.
//
// Solidity: event StoreSetField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes data)
func (_IWorld *IWorldFilterer) WatchStoreSetField(opts *bind.WatchOpts, sink chan<- *IWorldStoreSetField) (event.Subscription, error) {

	logs, sub, err := _IWorld.contract.WatchLogs(opts, "StoreSetField")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorldStoreSetField)
				if err := _IWorld.contract.UnpackLog(event, "StoreSetField", log); err != nil {
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

// ParseStoreSetField is a log parse operation binding the contract event 0xd01f9f1368f831528fc9fe6442366b2b7d957fbfff3bcf7c24d9ab5fe51f8c46.
//
// Solidity: event StoreSetField(bytes32 table, bytes32[] key, uint8 schemaIndex, bytes data)
func (_IWorld *IWorldFilterer) ParseStoreSetField(log types.Log) (*IWorldStoreSetField, error) {
	event := new(IWorldStoreSetField)
	if err := _IWorld.contract.UnpackLog(event, "StoreSetField", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWorldStoreSetRecordIterator is returned from FilterStoreSetRecord and is used to iterate over the raw logs and unpacked data for StoreSetRecord events raised by the IWorld contract.
type IWorldStoreSetRecordIterator struct {
	Event *IWorldStoreSetRecord // Event containing the contract specifics and raw log

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
func (it *IWorldStoreSetRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWorldStoreSetRecord)
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
		it.Event = new(IWorldStoreSetRecord)
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
func (it *IWorldStoreSetRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWorldStoreSetRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWorldStoreSetRecord represents a StoreSetRecord event raised by the IWorld contract.
type IWorldStoreSetRecord struct {
	Table [32]byte
	Key   [][32]byte
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStoreSetRecord is a free log retrieval operation binding the contract event 0x912af873e852235aae78a1d25ae9bb28b616a67c36898c53a14fd8184504ee32.
//
// Solidity: event StoreSetRecord(bytes32 table, bytes32[] key, bytes data)
func (_IWorld *IWorldFilterer) FilterStoreSetRecord(opts *bind.FilterOpts) (*IWorldStoreSetRecordIterator, error) {

	logs, sub, err := _IWorld.contract.FilterLogs(opts, "StoreSetRecord")
	if err != nil {
		return nil, err
	}
	return &IWorldStoreSetRecordIterator{contract: _IWorld.contract, event: "StoreSetRecord", logs: logs, sub: sub}, nil
}

// WatchStoreSetRecord is a free log subscription operation binding the contract event 0x912af873e852235aae78a1d25ae9bb28b616a67c36898c53a14fd8184504ee32.
//
// Solidity: event StoreSetRecord(bytes32 table, bytes32[] key, bytes data)
func (_IWorld *IWorldFilterer) WatchStoreSetRecord(opts *bind.WatchOpts, sink chan<- *IWorldStoreSetRecord) (event.Subscription, error) {

	logs, sub, err := _IWorld.contract.WatchLogs(opts, "StoreSetRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWorldStoreSetRecord)
				if err := _IWorld.contract.UnpackLog(event, "StoreSetRecord", log); err != nil {
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

// ParseStoreSetRecord is a log parse operation binding the contract event 0x912af873e852235aae78a1d25ae9bb28b616a67c36898c53a14fd8184504ee32.
//
// Solidity: event StoreSetRecord(bytes32 table, bytes32[] key, bytes data)
func (_IWorld *IWorldFilterer) ParseStoreSetRecord(log types.Log) (*IWorldStoreSetRecord, error) {
	event := new(IWorldStoreSetRecord)
	if err := _IWorld.contract.UnpackLog(event, "StoreSetRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
