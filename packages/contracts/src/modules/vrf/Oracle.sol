pragma solidity ^0.8.16;

import {StorageProof} from "src/libraries/StateProofHelper.sol";
import {RLPReader} from "@optimism-bedrock/rlp/RLPReader.sol";
import {Interpreter} from "src/evm/Interpreter.sol";
import {Types} from "src/evm/Instructions.sol";
import {IOracle} from "src/evm/IOracle.sol";

abstract contract BaseOracle is IOracle {
    mapping(uint256 => bytes32) public blockHashes;

    function getBlockHash(uint256 blockNumber) public view override returns (bytes32) {
        if (blockHashes[blockNumber] != 0) {
            return blockHashes[blockNumber];
        }
        revert("No block hash found");
    }
}

contract MockOracle is BaseOracle {
    function setBlockHash(uint256 blockNumber, bytes32 hashh) external {
        blockHashes[blockNumber] = hashh;
    }
}

interface IBlockhashStore {
    function getBlockhash(uint256 n) external view returns (bytes32);
}

contract StorageProofOracle is BaseOracle {
    address public immutable interpreter;
    address public immutable blockHashStore;

    mapping(uint256 => bytes32) public stateRoots;
    mapping(bytes32 => bytes32) public storageRoots;

    constructor(address interpreterAddr, address blockHashStoreAddr) {
        interpreter = interpreterAddr;
        blockHashStore = blockHashStoreAddr;
    }

    // @notice Given a block number, proves the block hash directly from the EVM's blockhash opcode.
    //         If the block hash is not available, it will be 0 so nothing happens.
    //         Note that the state root is not saved.
    function proveBlockHash(uint256 blockNumber) external {
        blockHashes[blockNumber] = blockhash(blockNumber);
    }

    // @notice Given a block number, imports the block hash from the BlockhashStore contract.
    //         Will revert if the block hash is not found.
    //         Note that the state root is not saved.
    function importBlockHash(uint256 blockNumber) external {
        blockHashes[blockNumber] = IBlockhashStore(blockHashStore).getBlockhash(blockNumber);
    }

    /// @notice Given block headers, stores the parentHash and the stateRoot of each header.
    function proveBlockHashesAndStateRoots(bytes[] memory blockHeaders) external {
        for (uint256 i = 0; i < blockHeaders.length; i++) {
            RLPReader.RLPItem[] memory blockHeader = RLPReader.readList(blockHeaders[i]);

            // Ensure block hash is valid
            uint256 blockNumber = RLPReader.readUint256(blockHeader[8]);
            bytes32 blockHash = keccak256(blockHeaders[i]);
            require(blockHashes[blockNumber] == blockHash, "Block hash not proven");

            // Store parent hash
            uint256 parentNumber = blockNumber - 1;
            bytes32 parentHash = RLPReader.readBytes32(blockHeader[0]);
            blockHashes[parentNumber] = parentHash;
        }
    }
}
