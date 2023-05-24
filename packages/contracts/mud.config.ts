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
  },
  modules: [
    {
      name: "VRFCoordinatorModule",
      root: true,
      args: [],
    },
  ],
});
