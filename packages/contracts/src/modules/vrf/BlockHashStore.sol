// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {RLPReader} from "./libraries/RLPReader.sol";

contract BlockHashStore {
    mapping(uint256 => bytes32) public blockHashes;
    uint256 public latestBlockNumber;

    /// @notice Gets a block hash at a particular block number.
    /// @param _n The block number;
    function getBlockHash(uint256 _n) external view returns (bytes32) {
        bytes32 h = blockHashes[_n];
        require(h != 0, "block hash not available");
        return h;
    }

    /// @notice Stores block hashes for a range of blocks using the blockhash opcode.
    /// @param _blockNumbers The block numbers to store hashes for.
    function storeBlockHashesViaOpCode(uint256[] memory _blockNumbers) external {
        for (uint256 i = 0; i < _blockNumbers.length; i++) {
            bytes32 h = blockhash(_blockNumbers[i]);
            require(h != 0, "block hash not available");
            blockHashes[_blockNumbers[i]] = h;
        }
    }

    /// @notice Stores block hashes for a range of blocks using merkle proofs.
    /// @param _blockHeaders The block headers to store hashes for.
    function storeBlockHashesViaMerkleProofs(bytes[] memory _blockHeaders) external {
        for (uint256 i = 0; i < _blockHeaders.length; i++) {
            RLPReader.RLPItem[] memory blockHeader = RLPReader.readList(_blockHeaders[i]);

            uint256 blockNumber = RLPReader.readUint256(blockHeader[8]);
            bytes32 blockHash = keccak256(_blockHeaders[i]);
            require(blockHashes[blockNumber] == blockHash, "Block hash not proven");

            uint256 parentNumber = blockNumber - 1;
            bytes32 parentHash = RLPReader.readBytes32(blockHeader[0]);
            blockHashes[parentNumber] = parentHash;
        }
    }
}