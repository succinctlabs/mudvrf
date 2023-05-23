// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

/* Autogenerated file. Do not edit manually. */

interface IBlackJackSystem {
  function startGame() external;

  function handleStartGame(bytes32 requestId, uint256[] memory randomWords) external;

  function dealUser() external;

  function handleDealUser(bytes32 requestId, uint256[] memory randomWords) external;

  function standUser() external;

  function handleStandUser(bytes32 requestId, uint256[] memory randomWords) external;
}
