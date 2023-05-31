# MUDVRF

**A secure, fast, and easy-to-use VRF for on-chain games built using MUD.**

![mudvrf](https://github.com/jtguibas/mudvrf/assets/25734765/09db1c47-3053-47e3-868e-2f2240dbb8aa)

## Get Started

Install the MUDVRF dependencies into the package where your contracts live.

```sh
cd packages/contracts
pnpm add @succinctlabs/mudvrf-contracts
pnpm add @succinctlabs/mudvrf-relayer
```

Import and use the `VRFCoordinator` inside systems within your MUD project.

```solidity
import {System} from "@latticexyz/world/src/System.sol";

import {IVRFCoordinatorSystem} from "@succinctlabs/mudvrf/world/IVRFCoordinatorSystem.sol";

contract YourSystem is System {

    bytes32 constant ORACLE_ID = 0xc1ffd3cfee2d9e5cd67643f8f39fd6e51aad88f6f4ce6ab8827279cfffb92266;
    uint32 constant NB_WORDS = 1;
    uint16 constant REQUEST_CONFIRMATIONS = 1;
    uint32 constant CALLBACK_GAS_LIMIT = 100000;

    function dealCard() internal returns (bytes32) {
        bytes32 requestId = IVRFCoordinatorSystem(_world()).requestRandomWords(
            ORACLE_ID,
            NB_WORDS, 
            REQUEST_CONFIRMATIONS,
            CALLBACK_GAS_LIMIT,
            IYourSystem.handleDealCards.selector
        );
        return requestId;
    }

    function handleDealCards(bytes32 requestId, uint256[] randomWords) {
        // Your logic here!
        ...
    }
}
```
