import { mudConfig } from "@latticexyz/world/register";

export default mudConfig({
  codegenDirectory: "",
  tables: {
    BlackJack: {
      keySchema: {
        userAddress: "address"
      },
      schema: {
        userWins: "uint256",
        userLosses: "uint256",
        gameEnded: "bool",
        userWon: "bool",
        userCards: "uint256[]",
        dealerCards: "uint256[]",
      }
    },
    RequestIdToBlackJackUser: {
      keySchema: {
        requestId: "bytes32"
      },
      schema: {
        user: "address"
      }
    },
    VRFRequests: {
      directory: "modules/vrf/tables",
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
    VRFNonce: {
      directory: "modules/vrf/tables",
      keySchema: {},
      schema: "uint256", 
    }
  },
  modules: [
    {
      name: "VRFCoordinator",
      root: true,
      args: [],
    },
  ],
});
