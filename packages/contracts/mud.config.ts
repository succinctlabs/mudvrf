import { mudConfig } from "@latticexyz/world/register";

export default mudConfig({
  codegenDirectory: "",
  namespace: "mudvrf",
  tables: {
    VRFCoordinatorAddress: {
      keySchema: {},
      schema: {
        vrfCoordinatorAddress: "address"
      }
    }
  },
  modules: [
    {
      name: "VRFCoordinatorModule",
      root: true,
      args: [],
    },
  ],
});
