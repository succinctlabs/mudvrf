// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

/* Autogenerated file. Do not edit manually. */

interface IVRFCoordinatorSystem {
  function setCoordinator(address _coordinator) external;

  function requestRandomWords(
    bytes32 _oracleId,
    uint32 _nbWords,
    uint16 _requestConfirmations,
    uint32 _callbackGasLimit,
    bytes4 _callbackSelector
  ) external returns (bytes32);
}
