// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";

import {VRF} from "./VRF.sol";
import {IVRFCoordinator} from "./interfaces/IVRFCoordinator.sol";
import {BlockHashStore} from "./BlockHashStore.sol";

/// @title VRFCoordinator
/// @notice This contract handles requests and fulfillments of random words from a VRF.
contract VRFCoordinator is VRF, IVRFCoordinator {
    /// @notice The oracle identifier used for validating the VRF proof came from the oracle.
    bytes32 public constant ORACLE_ID = 0xc1ffd3cfee2d9e5cd67643f8f39fd6e51aad88f6f4ce6ab8827279cfffb92266;

    /// @notice The oracle address used for validating that the VRF proof came from the oracle.
    address public constant ORACLE_ADDRESS = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;

    /// @notice The minimum number of request confirmatins.
    uint16 public constant MINIMUM_REQUEST_CONFIRMATIONS = 0;

    /// @notice The maximum callback gas limit.
    uint64 public constant MAXIMUM_CALLBACK_GAS_LIMIT = 10000000;

    /// @notice The maximum number of random words that can be provided in the callback.
    uint64 public constant MAXIMUM_NB_WORDS = 64;

    /// @notice The request nonce.
    uint256 public nonce = 0;

    /// @notice The storage proof oracle.
    BlockHashStore public blockHashStore;

    /// @notice The mapping of request ids to commitments to what is stored in the request.
    mapping(bytes32 => bytes32) public requests;

    /// @notice The mapping of oracle ids to oracle addresses.
    mapping(bytes32 => address) public oracles;

    constructor(address _blockHashStore) {
        oracles[ORACLE_ID] = ORACLE_ADDRESS;
        blockHashStore = BlockHashStore(_blockHashStore);
    }

    /// @notice Requests random words from the VRF.
    /// @param _oracleId The address of the operator to get shares for.
    /// @param _requestConfirmations The number of blocks to wait before posting the VRF request.
    /// @param _callbackGasLimit The maximum amount of gas the callback can use.
    /// @param _nbWords The number of random words to request.
    /// @param _callbackSelector The selector of the callback function.
    function requestRandomWords(
        bytes32 _oracleId,
        uint32 _nbWords,
        uint16 _requestConfirmations,
        uint32 _callbackGasLimit,
        address _callbackAddress,
        bytes4 _callbackSelector
    ) external returns (bytes32) {
        if (_requestConfirmations < MINIMUM_REQUEST_CONFIRMATIONS) {
            revert InvalidRequestConfirmations();
        } else if (_callbackGasLimit > MAXIMUM_CALLBACK_GAS_LIMIT) {
            revert InvalidCallbackGasLimit();
        } else if (_nbWords > MAXIMUM_NB_WORDS) {
            revert InvalidNumberOfWords();
        }

        bytes32 seed = keccak256(abi.encode(_callbackAddress, nonce));
        bytes32 requestId = keccak256(abi.encode(_oracleId, seed));
        requests[requestId] = keccak256(
            abi.encode(
                requestId,
                msg.sender,
                nonce,
                _oracleId,
                _nbWords,
                _requestConfirmations,
                _callbackGasLimit,
                _callbackAddress,
                _callbackSelector
            )
        );

        emit RequestRandomWords(
            requestId,
            msg.sender,
            nonce,
            _oracleId,
            _nbWords,
            _requestConfirmations,
            _callbackGasLimit,
            _callbackAddress,
            _callbackSelector
        );

        nonce += 1;
        return requestId;
    }

    /// @notice Fulfills the request for random words.
    /// @param _proof The address of the operator to get shares for.
    /// @param _request The number of shares for the operator.
    function fulfillRandomWords(VRF.Proof memory _proof, VRF.Request memory _request) external {
        bytes32 oracleId = keccak256(abi.encode(_proof.pk));
        address oracle = oracles[oracleId];
        if (oracle == address(0)) {
            revert("Invalid oracle id");
        }

        bytes32 seed = keccak256(abi.encode(_request.callbackAddress, _request.nonce));
        bytes32 requestId = keccak256(abi.encode(oracleId, seed));
        bytes32 commitment = requests[requestId];
        bytes32 expectedCommitment = keccak256(
            abi.encode(
                requestId,
                _request.sender,
                _request.nonce,
                _request.oracleId,
                _request.nbWords,
                _request.requestConfirmations,
                _request.callbackGasLimit,
                _request.callbackAddress,
                _request.callbackSelector
            )
        );
        if (commitment == bytes32(0)) {
            revert("Invalid commitment 1");
        } else if (commitment != expectedCommitment) {
            revert("Invalid commitment 2");
        }
        delete requests[requestId];

        bytes32 blockHash = blockHashStore.getBlockHash(_request.blockNumber);
        uint256 actualSeed = uint256(keccak256(abi.encodePacked(seed, blockHash)));

        uint256 randomness = VRF.randomValueFromVRFProof(_proof, actualSeed);
        uint256[] memory randomWords = new uint256[](_request.nbWords);
        for (uint256 i = 0; i < _request.nbWords; i++) {
            randomWords[i] = uint256(keccak256(abi.encode(randomness, i)));
        }

        bytes memory fulfillRandomnessCall = abi.encodeWithSelector(_request.callbackSelector, requestId, randomWords);
        (bool status,) = _request.callbackAddress.call(fulfillRandomnessCall);

        if (!status) {
            revert("Failed to fulfill randomness");
        }

        emit FulfillRandomWords(requestId);
    }
}
