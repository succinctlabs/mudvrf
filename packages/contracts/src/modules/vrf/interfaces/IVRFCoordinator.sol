// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {VRF} from "../VRF.sol";

interface IVRFCoordinator {
    function requestRandomWords(
        bytes32 _oracleId,
        uint16 _requestConfirmations,
        uint32 _callbackGasLimit,
        uint32 _nbWords,
        bytes4 _callbackSelector
    ) external returns (bytes32);

    function fulfillRandomWords(
        VRF.Proof memory _proof,
        VRF.Request memory _request
    ) external;
}