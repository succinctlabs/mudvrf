import { mudConfig } from "@latticexyz/world/register";

export default mudConfig({
  codegenDirectory: "",
  tables: {
    VRFCoordinatorAddress: {
      directory: "modules/vrf/tables", 
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
