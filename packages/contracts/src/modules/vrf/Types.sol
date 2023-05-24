// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

struct VRFRequest {
    uint64 blockNumber;
    uint32 callbackGasLimit;
    uint32 nbWords;
    address sender;
    bytes4 callbackSelector;
}