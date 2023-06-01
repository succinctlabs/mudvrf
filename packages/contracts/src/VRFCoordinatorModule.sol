// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {IBaseWorld} from "@latticexyz/world/src/interfaces/IBaseWorld.sol";
import {IModule} from "@latticexyz/world/src/interfaces/IModule.sol";
import {WorldContext} from "@latticexyz/world/src/WorldContext.sol";
import { getTargetTableSelector } from "@latticexyz/world/src/modules/utils/getTargetTableSelector.sol";

import {VRFCoordinator} from "./VRFCoordinator.sol";
import {VRFCoordinatorSystem} from "./VRFCoordinatorSystem.sol";
import {VRFCoordinatorAddress} from "./tables/VRFCoordinatorAddress.sol";
import {NAMESPACE, MODULE_NAME, SYSTEM_NAME, TABLE_NAME} from "./constants.sol";

contract VRFCoordinatorModule is IModule, WorldContext {
    VRFCoordinatorSystem public immutable vrfCoordinatorSystem = new VRFCoordinatorSystem();

    function getName() public pure returns (bytes16) {
        return MODULE_NAME;
    }

    function install(bytes memory) public {
        IBaseWorld world = IBaseWorld(_world());

        // Register system
        world.registerSystem(NAMESPACE, SYSTEM_NAME, vrfCoordinatorSystem, true);

        // Register tables
        bytes32 resourceSelector = world.registerTable(
            NAMESPACE,
            TABLE_NAME,
            VRFCoordinatorAddress.getSchema(),
            VRFCoordinatorAddress.getKeySchema()
        );

        // Register metadata for tables
        (string memory tableName, string[] memory fieldNames) = VRFCoordinatorAddress.getMetadata();
        IBaseWorld(_world()).setMetadata(
            NAMESPACE,
            TABLE_NAME,
            tableName,
            fieldNames
        );

        // Register system functions
        world.registerFunctionSelector(
            NAMESPACE, SYSTEM_NAME, "setCoordinator", "(address)"
        );
        world.registerFunctionSelector(
            NAMESPACE, SYSTEM_NAME, "requestRandomWords", "(bytes32,uint32,uint16,uint32,bytes4)"
        );
    }
}
