// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface IOracle {
    function getBlockHash(uint256 blockNumber) external view returns (bytes32);
}
