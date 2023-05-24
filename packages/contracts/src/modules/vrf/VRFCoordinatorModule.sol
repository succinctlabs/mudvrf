// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {IBaseWorld} from "@latticexyz/world/src/interfaces/IBaseWorld.sol";
import {IModule} from "@latticexyz/world/src/interfaces/IModule.sol";
import {WorldContext} from "@latticexyz/world/src/WorldContext.sol";

import {VRFRequests} from "./tables/VRFRequests.sol";
import {VRFCoordinatorSystem} from "./VRFCoordinatorSystem.sol";
import {NAMESPACE, MODULE_NAME, SYSTEM_NAME, TABLE_NAME} from "./constants.sol";

contract VRFCoordinatorModule is IModule, WorldContext {
    VRFCoordinatorSystem public immutable vrfCoordinator = new VRFCoordinatorSystem();

    function getName() public pure returns (bytes16) {
        return MODULE_NAME;
    }

    function install(bytes memory) public {
        IBaseWorld world = IBaseWorld(_world());

        // Register system
        world.registerSystem(NAMESPACE, SYSTEM_NAME, vrfCoordinator, true);

        // Register system's functions
        world.registerFunctionSelector(
            NAMESPACE, SYSTEM_NAME, "requestRandomWords", "(address,bytes32,uint16,uint32,uint32,bytes4)"
        );
    }
}
