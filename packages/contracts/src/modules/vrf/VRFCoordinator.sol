// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import { IBaseWorld } from "@latticexyz/world/src/interfaces/IBaseWorld.sol";
import { IModule } from "@latticexyz/world/src/interfaces/IModule.sol";
import { WorldContext } from "@latticexyz/world/src/WorldContext.sol";

import { VRFRequests } from "./tables/VRFRequests.sol";
import { VRFCoordinatorSystem } from "./VRFCoordinatorSystem.sol";
import { NAMESPACE, MODULE_NAME, SYSTEM_NAME, TABLE_NAME } from "./constants.sol";

contract VRFCoordinator is IModule, WorldContext {
  VRFCoordinatorSystem public immutable vrfCoordinator = new VRFCoordinatorSystem();

  function getName() public pure returns (bytes16) {
    return MODULE_NAME;
  }

  function install(bytes memory) public {
    IBaseWorld world = IBaseWorld(_world());

    // Register table
    world.registerTable(NAMESPACE, TABLE_NAME, VRFRequests.getSchema(), VRFRequests.getKeySchema());

    // Set table's metadata
    (string memory tableName, string[] memory fieldNames) = VRFRequests.getMetadata();
    world.setMetadata(NAMESPACE, TABLE_NAME, tableName, fieldNames);

    // Register system
    world.registerSystem(NAMESPACE, SYSTEM_NAME, vrfCoordinator, true);

    // Register system's functions
    world.registerFunctionSelector(NAMESPACE, SYSTEM_NAME, "requestRandomWords", "(bytes32,uint16,uint32,uint32,bytes4)");
    world.registerFunctionSelector(NAMESPACE, SYSTEM_NAME, "fulfillRandomWords", "((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256),(uint64,uint32,uint32,address,bytes4))");
  }
}