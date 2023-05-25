// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";

import {VRF} from "./VRF.sol";
import {IVRFCoordinator} from "./interfaces/IVRFCoordinator.sol";

/// @title VRFCoordinator
/// @notice This contract handles requests and fulfillments of random words from a VRF.
contract VRFCoordinator is VRF, IVRFCoordinator {
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

    /// @notice The request nonce.
    uint256 public nonce = 0;

    /// @notice The storage proof oracle.
    StorageProofOracle public storageProofOracle;

    /// @notice The mapping of request ids to commitments to what is stored in the request.
    mapping(bytes32 => bytes32) public requests;

    /// @notice The mapping of oracle ids to oracle addresses.
    mapping(bytes32 => address) public oracles;

    event RequestRandomWords(uint256 indexed nonce, bytes32 indexed requestId, bytes32 seed, uint64 nbWords);
    event FulfillRandomWords(uint256 indexed nonce, bytes32 indexed requestId, uint256[] words);

    error InvalidRequestConfirmations();
    error InvalidCallbackGasLimit();
    error InvalidNumberOfWords();
    error InvalidOracleId();
    error InvalidCommitment();
    error InvalidRequestParameters();
    error FailedToFulfillRandomness();

    constructor(address _storageProofOracle) {
        oracles[ORACLE_ID] = ORACLE_ADDRESS;
        storageProofOracle = StorageProofOracle(_storageProofOracle);
    }

    /// @notice Requests random words from the VRF.
    /// @param _oracleId The address of the operator to get shares for.
    /// @param _requestConfirmations The number of blocks to wait before posting the VRF request.
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

        bytes32 seed = keccak256(abi.encode(msg.sender, nonce));
        bytes32 requestId = keccak256(abi.encode(_oracleId, seed));
        requests[requestId] = keccak256(
            abi.encode(
                requestId,
                msg.sender,
                block.number,
                _oracleId,
                _requestConfirmations,
                _callbackGasLimit,
                _nbWords,
                _callbackSelector
            )
        );
        nonce += 1;

        emit RequestRandomWords(nonce, requestId, seed, _nbWords);

        return requestId;
    }

    /// @notice Fulfills the request for random words.
    /// @param _proof The address of the operator to get shares for.
    /// @param _request The number of shares for the operator.
    function fulfillRandomWords(VRF.Proof memory _proof, VRF.Request memory _request) external {
        bytes32 oracleId = keccak256(abi.encode(_proof.pk));
        address oracle = oracles[oracleId];
        if (oracle == address(0)) {
            revert InvalidOracleId();
        }

        bytes32 requestId = keccak256(abi.encode(oracleId, _proof.seed));
        bytes32 commitment = requests[requestId];
        bytes32 expectedCommitment = keccak256(
            abi.encode(
                requestId,
                _request.sender,
                _request.blockNumber,
                _request.oracleId,
                _request.requestConfirmations,
                _request.callbackGasLimit,
                _request.nbWords,
                _request.callbackSelector
            )
        );
        if (commitment == bytes32(0)) {
            revert InvalidCommitment();
        } else if (commitment != expectedCommitment) {
            revert InvalidCommitment();
        }
        delete requests[requestId];

        bytes32 blockHash = blockhash(_request.blockNumber);
        if (blockHash == bytes32(0)) {
            blockHash = storageProofOracle.getBlockHash(_request.blockNumber);
            require(blockHash != bytes32(0), "Block hash is not proven.");
        }
        uint256 actualSeed = uint256(keccak256(abi.encodePacked(_proof.seed, blockHash)));

        uint256 randomness = VRF.randomValueFromVRFProof(_proof, actualSeed);
        uint256[] memory randomWords = new uint256[](_request.nbWords);
        for (uint256 i = 0; i < _request.nbWords; i++) {
            randomWords[i] = uint256(keccak256(abi.encode(randomness, i)));
        }

        bytes memory fulfillRandomnessCall = abi.encodeWithSelector(_request.callbackSelector, requestId, randomWords);
        (bool status,) = _request.sender.call(fulfillRandomnessCall);

        if (!status) {
            revert FailedToFulfillRandomness();
        }

        emit FulfillRandomWords(nonce, requestId, randomWords);
    }
}
