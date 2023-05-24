// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import { System } from "@latticexyz/world/src/System.sol";

import { VRF } from "./VRF.sol";
import { VRFRequest } from "./Types.sol";
import { VRFRequests, VRFRequestsData } from "./tables/VRFRequests.sol";
import { VRFNonce } from "./tables/VRFNonce.sol";

/// @title VRFCoordinatorSystem
/// @notice This contract handles requests and fulfillments of random words from a VRF.
contract VRFCoordinatorSystem is System, VRF {
    /// @notice The oracle identifier used for validating the VRF proof came from the oracle.
    bytes32 public constant ORACLE_ID = 0xc0a6c424ac7157ae408398df7e5f4552091a69125d5dfcb7b8c2659029395bdf;

    /// @notice The oracle address used for validating that the VRF proof came from the oracle.
    address public constant ORACLE_ADDRESS = 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf;

    /// @notice The minimum number of request confirmatins.
    uint16 public constant MINIMUM_REQUEST_CONFIRMATIONS = 0;

    /// @notice The maximum callback gas limit.
    uint64 public constant MAXIMUM_CALLBACK_GAS_LIMIT = 10000000;

    /// @notice The maximum number of random words that can be provided in the callback.
    uint64 public constant MAXIMUM_NB_WORDS = 64;

    error InvalidRequestConfirmations();
    error InvalidCallbackGasLimit();
    error InvalidNumberOfWords();
    error InvalidOracleId();
    error InvalidCommitment();
    error InvalidRequestParameters();
    error FailedToFulfillRandomness();

    /// @notice Requests random words from the VRF.
    /// @param _oracleId The address of the operator to get shares for.
    /// @param _requestConfirmations The number of shares for the operator.
    /// @param _callbackGasLimit The maximum amount of gas the callback can use.
    /// @param _nbWords The number of random words to request.
    /// @param _callbackSelector The selector of the callback function.
    function requestRandomWords(
        bytes32 _oracleId,
        uint16 _requestConfirmations,
        uint32 _callbackGasLimit,
        uint32 _nbWords,
        bytes4 _callbackSelector
    ) external returns (bytes32) {
        if (_requestConfirmations < MINIMUM_REQUEST_CONFIRMATIONS) {
            revert InvalidRequestConfirmations();
        } else if (_callbackGasLimit > MAXIMUM_CALLBACK_GAS_LIMIT) {
            revert InvalidCallbackGasLimit();
        } else if (_nbWords > MAXIMUM_NB_WORDS) {
            revert InvalidNumberOfWords();
        }

        uint256 nonce = VRFNonce.get();
        VRFNonce.set(nonce + 1);

        bytes32 seed = keccak256(abi.encode(_msgSender(), nonce));
        bytes32 requestId = keccak256(abi.encode(_oracleId, seed));
        VRFRequests.set(
            requestId, _msgSender(), nonce, block.number, _callbackGasLimit, _nbWords, _callbackSelector
        );

        return requestId;
    }

    /// @notice Fulfills the request for random words.
    /// @param _proof The address of the operator to get shares for.
    /// @param _request The number of shares for the operator.
    function fulfillRandomWords(VRF.Proof memory _proof, VRFRequest memory _request) external {
        bytes32 oracleId = keccak256(abi.encode(_proof.pk));
        // TODO: check the oracleId is actually correct.

        bytes32 requestId = keccak256(abi.encode(oracleId, _proof.seed));
        VRFRequestsData memory request = VRFRequests.get(requestId);
        if (request.msgSender == address(0)) {
            revert InvalidRequestParameters();
        } else if (request.blockNumber != _request.blockNumber) {
            revert InvalidRequestParameters();
        } else if (request.callbackGasLimit != _request.callbackGasLimit) {
            revert InvalidRequestParameters();
        } else if (request.nbWords != _request.nbWords) {
            revert InvalidRequestParameters();
        } else if (request.callbackSelector != _request.callbackSelector) {
            revert InvalidRequestParameters();
        }
        VRFRequests.deleteRecord(requestId);

        uint256 randomness = VRF.randomValueFromVRFProof(_proof, _proof.seed);
        uint256[] memory randomWords = new uint256[](_request.nbWords);
        for (uint256 i = 0; i < _request.nbWords; i++) {
            randomWords[i] = uint256(keccak256(abi.encode(randomness, i)));
        }

        bytes memory callbackCall = abi.encodeWithSelector(_request.callbackSelector, requestId, randomWords);
        (bool status,) = _world().call(callbackCall);

        if (!status) {
            revert FailedToFulfillRandomness();
        }
    }
}
