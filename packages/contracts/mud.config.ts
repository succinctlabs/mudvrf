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
    }
  },
});
