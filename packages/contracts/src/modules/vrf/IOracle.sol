// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

interface IOracle {
    function getBlockHash(uint256 blockNumber) external view returns (bytes32);

    function getBalance(uint256 blockNumber, address addr) external view returns (uint256);

    function getStorageSlot(uint256 blockNumber, address addr, bytes32 hashh)
        external
        view
        returns (bytes32);

    function getCodeHash(uint256 blockNumber, address addr) external view returns (bytes32);

    function getCode(uint256 blockNumber, address addr) external view returns (bytes memory);

    function getCallResult(uint256 blockNumber, address from, address to, bytes memory args)
        external
        view
        returns (bool success, bytes memory);

    function getTimestamp(uint256 blockNumber) external view returns (uint256);
}
