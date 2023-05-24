// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";

import {VRF} from "./VRF.sol";
import {VRFCoordinator} from "./VRFCoordinator.sol";
import {VRFRequests, VRFRequestsData} from "./tables/VRFRequests.sol";
import {VRFNonce} from "./tables/VRFNonce.sol";

/// @title VRFCoordinatorSystem
/// @notice This contract handles requests and fulfillments of random words from a VRF.
contract VRFCoordinatorSystem is System, VRF {
    /// @notice Requests random words from the VRF.
    /// @param _coordinator The address of the VRF coordinator.
    /// @param _oracleId The address of the operator to get shares for.
    /// @param _requestConfirmations The number of shares for the operator.
    /// @param _callbackGasLimit The maximum amount of gas the callback can use.
    /// @param _nbWords The number of random words to request.
    /// @param _callbackSelector The selector of the callback function.
    function requestRandomWords(
        address _coordinator,
        bytes32 _oracleId,
        uint16 _requestConfirmations,
        uint32 _callbackGasLimit,
        uint32 _nbWords,
        bytes4 _callbackSelector
    ) external returns (bytes32) {
        VRFCoordinator coordinator = VRFCoordinator(_coordinator);
        bytes32 requestId = coordinator.requestRandomWords(
            _oracleId,
            _requestConfirmations,
            _callbackGasLimit,
            _nbWords,
            _callbackSelector
        );
        return requestId;
    }
}
