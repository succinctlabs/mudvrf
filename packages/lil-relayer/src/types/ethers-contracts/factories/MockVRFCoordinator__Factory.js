"use strict";
/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
Object.defineProperty(exports, "__esModule", { value: true });
exports.MockVRFCoordinator__factory = void 0;
var ethers_1 = require("ethers");
var _abi = [
    {
        inputs: [],
        stateMutability: "nonpayable",
        type: "constructor",
    },
    {
        inputs: [],
        name: "FailedToFulfillRandomness",
        type: "error",
    },
    {
        inputs: [],
        name: "InvalidCallbackGasLimit",
        type: "error",
    },
    {
        inputs: [],
        name: "InvalidNumberOfWords",
        type: "error",
    },
    {
        inputs: [],
        name: "InvalidOracleId",
        type: "error",
    },
    {
        inputs: [],
        name: "InvalidRequestConfirmations",
        type: "error",
    },
    {
        anonymous: false,
        inputs: [
            {
                indexed: true,
                internalType: "uint256",
                name: "nonce",
                type: "uint256",
            },
            {
                indexed: true,
                internalType: "bytes32",
                name: "requestId",
                type: "bytes32",
            },
            {
                indexed: false,
                internalType: "uint256[]",
                name: "words",
                type: "uint256[]",
            },
        ],
        name: "FulfilledRandomness",
        type: "event",
    },
    {
        anonymous: false,
        inputs: [
            {
                indexed: true,
                internalType: "uint256",
                name: "nonce",
                type: "uint256",
            },
            {
                indexed: true,
                internalType: "bytes32",
                name: "requestId",
                type: "bytes32",
            },
            {
                indexed: false,
                internalType: "address",
                name: "sender",
                type: "address",
            },
            {
                indexed: false,
                internalType: "uint256",
                name: "blockNumber",
                type: "uint256",
            },
            {
                indexed: false,
                internalType: "bytes32",
                name: "seed",
                type: "bytes32",
            },
            {
                indexed: false,
                internalType: "uint32",
                name: "callbackGasLimit",
                type: "uint32",
            },
            {
                indexed: false,
                internalType: "uint64",
                name: "nbWords",
                type: "uint64",
            },
            {
                indexed: false,
                internalType: "bytes4",
                name: "callbackSelector",
                type: "bytes4",
            },
        ],
        name: "RandomnessRequest",
        type: "event",
    },
    {
        inputs: [
            {
                components: [
                    {
                        internalType: "uint256[2]",
                        name: "pk",
                        type: "uint256[2]",
                    },
                    {
                        internalType: "uint256[2]",
                        name: "gamma",
                        type: "uint256[2]",
                    },
                    {
                        internalType: "uint256",
                        name: "c",
                        type: "uint256",
                    },
                    {
                        internalType: "uint256",
                        name: "s",
                        type: "uint256",
                    },
                    {
                        internalType: "uint256",
                        name: "seed",
                        type: "uint256",
                    },
                    {
                        internalType: "address",
                        name: "uWitness",
                        type: "address",
                    },
                    {
                        internalType: "uint256[2]",
                        name: "cGammaWitness",
                        type: "uint256[2]",
                    },
                    {
                        internalType: "uint256[2]",
                        name: "sHashWitness",
                        type: "uint256[2]",
                    },
                    {
                        internalType: "uint256",
                        name: "zInv",
                        type: "uint256",
                    },
                ],
                internalType: "struct VRF.Proof",
                name: "_proof",
                type: "tuple",
            },
            {
                components: [
                    {
                        internalType: "address",
                        name: "sender",
                        type: "address",
                    },
                    {
                        internalType: "uint256",
                        name: "nonce",
                        type: "uint256",
                    },
                    {
                        internalType: "uint64",
                        name: "blockNumber",
                        type: "uint64",
                    },
                    {
                        internalType: "bytes32",
                        name: "oracleId",
                        type: "bytes32",
                    },
                    {
                        internalType: "uint16",
                        name: "requestConfirmations",
                        type: "uint16",
                    },
                    {
                        internalType: "uint32",
                        name: "callbackGasLimit",
                        type: "uint32",
                    },
                    {
                        internalType: "uint32",
                        name: "nbWords",
                        type: "uint32",
                    },
                    {
                        internalType: "bytes4",
                        name: "callbackSelector",
                        type: "bytes4",
                    },
                ],
                internalType: "struct VRF.Request",
                name: "_request",
                type: "tuple",
            },
        ],
        name: "fulfillRandomWords",
        outputs: [],
        stateMutability: "nonpayable",
        type: "function",
    },
    {
        inputs: [],
        name: "nonce",
        outputs: [
            {
                internalType: "uint256",
                name: "",
                type: "uint256",
            },
        ],
        stateMutability: "view",
        type: "function",
    },
    {
        inputs: [
            {
                internalType: "bytes32",
                name: "",
                type: "bytes32",
            },
        ],
        name: "oracles",
        outputs: [
            {
                internalType: "address",
                name: "",
                type: "address",
            },
        ],
        stateMutability: "view",
        type: "function",
    },
    {
        inputs: [
            {
                internalType: "bytes32",
                name: "_oracleId",
                type: "bytes32",
            },
            {
                internalType: "uint16",
                name: "_requestConfirmations",
                type: "uint16",
            },
            {
                internalType: "uint32",
                name: "_callbackGasLimit",
                type: "uint32",
            },
            {
                internalType: "uint32",
                name: "_nbWords",
                type: "uint32",
            },
            {
                internalType: "bytes4",
                name: "_callbackSelector",
                type: "bytes4",
            },
        ],
        name: "requestRandomWords",
        outputs: [
            {
                internalType: "bytes32",
                name: "",
                type: "bytes32",
            },
        ],
        stateMutability: "nonpayable",
        type: "function",
    },
    {
        inputs: [
            {
                internalType: "bytes32",
                name: "",
                type: "bytes32",
            },
        ],
        name: "requests",
        outputs: [
            {
                internalType: "bytes32",
                name: "",
                type: "bytes32",
            },
        ],
        stateMutability: "view",
        type: "function",
    },
];
var MockVRFCoordinator__factory = exports.MockVRFCoordinator__factory = /** @class */ (function () {
    function MockVRFCoordinator__factory() {
    }
    MockVRFCoordinator__factory.createInterface = function () {
        return new ethers_1.utils.Interface(_abi);
    };
    MockVRFCoordinator__factory.connect = function (address, signerOrProvider) {
        return new ethers_1.Contract(address, _abi, signerOrProvider);
    };
    MockVRFCoordinator__factory.abi = _abi;
    return MockVRFCoordinator__factory;
}());
