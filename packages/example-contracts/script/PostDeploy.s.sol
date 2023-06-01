// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {IVRFCoordinatorSystem} from "@succinctlabs/mudvrf/src/world/IVRFCoordinatorSystem.sol";
import {IWorld} from "../src/world/IWorld.sol";
import {BlockHashStore} from "@succinctlabs/mudvrf/src/BlockHashStore.sol";
import {VRFCoordinator} from "@succinctlabs/mudvrf/src/VRFCoordinator.sol";
import {MockVRFCoordinator} from "@succinctlabs/mudvrf/src/mocks/MockVRFCoordinator.sol";

contract PostDeploy is Script {
    uint256 public constant ANVIL_CHAIN_ID = 31337; 

    function run(address worldAddress) external {
        // Load the private key from the `PRIVATE_KEY` environment variable (in .env)
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");

        // Deploy VRFCoordinator & Set Coordinator
        vm.startBroadcast(deployerPrivateKey);
        address blockHashStore = address(new BlockHashStore());
        address coordinator = address(new VRFCoordinator(blockHashStore));
        IVRFCoordinatorSystem(worldAddress).mudvrf_VRFCoordinatorSy_setCoordinator(
            coordinator
        );
        vm.stopBroadcast();

        string memory obj1 = "vrfCoordinatorDeployment";
        string memory finalJson = vm.serializeAddress(obj1, "vrfCoordinatorAddress", coordinator);
        finalJson = vm.serializeAddress(obj1, "blockHashStoreAddress", blockHashStore);
        vm.writeJson(finalJson, "./vrf.json");
    }
}