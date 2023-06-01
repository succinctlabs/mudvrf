// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {VRF} from "../VRF.sol";

interface IVRFCoordinator {
    event RequestRandomWords(
        bytes32 requestId,
        address sender,
        uint256 nonce,
        bytes32 oracleId,
        uint32 nbWords,
        uint16 requestConfirmations,
        uint32 callbackGasLimit,
        address callbackAddress,
        bytes4 callbackSelector
    );
    event FulfillRandomWords(bytes32 requestId);

    error InvalidRequestConfirmations();
    error InvalidCallbackGasLimit();
    error InvalidNumberOfWords();
    error InvalidOracleId();
    error InvalidCommitment();
    error InvalidRequestParameters();
    error FailedToFulfillRandomness();

    function requestRandomWords(
        bytes32 _oracleId,
        uint32 _nbWords,
        uint16 _requestConfirmations,
        uint32 _callbackGasLimit,
        address _callbackAddress,
        bytes4 _callbackSelector
    ) external returns (bytes32);

    function fulfillRandomWords(
        VRF.Proof memory _proof,
        VRF.Request memory _request
    ) external;
}