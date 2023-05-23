// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";
import {VRFRequestTableV2, VRFRequestNonce, VRFRequestTableV2Data} from "../codegen/Tables.sol";
import {VRF, VRFRequest} from "./VRF.sol";

contract VRFSystem is System, VRF {
    bytes32 public constant ORACLE_ID = 0xc0a6c424ac7157ae408398df7e5f4552091a69125d5dfcb7b8c2659029395bdf;
    address public constant ORACLE_ADDRESS = 0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf;
    uint16 public constant MINIMUM_REQUEST_CONFIRMATIONS = 0;
    uint64 public constant MAXIMUM_CALLBACK_GAS_LIMIT = 10000000;
    uint64 public constant MAXIMUM_NB_WORDS = 64;

    event RandomnessRequest(uint256 indexed nonce, bytes32 indexed requestId, bytes32 seed, uint64 nbWords);
    event FulfilledRandomness(bytes32 indexed requestId, uint256[] words);

    error InvalidRequestConfirmations();
    error InvalidCallbackGasLimit();
    error InvalidNumberOfWords();
    error InvalidOracleId();
    error InvalidCommitment();
    error InvalidRequestParameters();
    error FailedToFulfillRandomness();

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

        uint256 nonce = VRFRequestNonce.get();
        VRFRequestNonce.set(nonce + 1);

        bytes32 seed = keccak256(abi.encode(_msgSender(), nonce));
        bytes32 requestId = keccak256(abi.encode(_oracleId, seed));
        VRFRequestTableV2.set(
            requestId,
            _msgSender(),
            nonce,
            block.number,
            _callbackGasLimit,
            _nbWords,
            _callbackSelector
        );

        return requestId;
    }

    function fulfillRandomWords(VRF.Proof memory _proof, VRFRequest memory _request) external {
        bytes32 oracleId = keccak256(abi.encode(_proof.pk));
        // TODO: check the oracleId is actually correct.

        bytes32 requestId = keccak256(abi.encode(oracleId, _proof.seed));
        VRFRequestTableV2Data memory request = VRFRequestTableV2.get(requestId);
        // if (commitment == bytes32(0)) {
        //     revert InvalidCommitment();
        // } else if (
        //     commitment
        //         != keccak256(
        //             abi.encode(
        //                 requestId, _request.blockNumber, _request.callbackGasLimit, _request.nbWords, _request.sender
        //             )
        //         )
        // ) {
        //     revert InvalidRequestParameters();
        // }
        // VRFRequestTable.deleteRecord(requestId);

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

        emit FulfilledRandomness(requestId, randomWords);
    }
}
