import { mudConfig } from "@latticexyz/world/register";

export default mudConfig({
  tables: {
    RaffleCounter: {
      keySchema: {},
      schema: "uint32",
    },
    RaffleTable: {
      keySchema: {
        raffleId: "uint32",
      },
      schema: {
        tickets: "address[]"
      }
    },
    VRFRequestNonce: {
      keySchema: {},
      schema: "uint256",
    },
    VRFRequestTable: {
      keySchema: {
        requestId: "bytes32"
      },
      schema: {
        commitment: "bytes32"
      }
    },
    RequestIdToRaffleId: {
      keySchema: {
        requestId: "bytes32"
      },
      schema: {
        raffleId: "uint32"
      }
    },
    VRFRequestTableV2: {
      keySchema: {
        requestId: "bytes32"
      },
      schema: { 
        msgSender: "address",
        nonce: "uint256",
        blockNumber: "uint256",
        callbackGasLimit: "uint32",
        nbWords: "uint32",
        callbackSelector: "bytes4"
      }
    },
    BlackJack: {
      keySchema: {
        userAddress: "address"
      },
      schema: {
        userCards: "uint256[]",
        dealerCards: "uint256[]",
        userWins: "uint256",
        userLosses: "uint256",
      }
    },
    RequestIdToBlackJackUser: {
      keySchema: {
        requestId: "bytes32"
      },
      schema: {
        user: "address"
      }
    }
  },
});
