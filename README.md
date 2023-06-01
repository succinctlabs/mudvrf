# MUDVRF

**A secure, fast, and easy-to-use random number generator (RNG) for on-chain games built using MUD.**

![mudvrf](https://github.com/jtguibas/mudvrf/assets/25734765/09db1c47-3053-47e3-868e-2f2240dbb8aa)

Accessing secure randomness on apps built ontop of [MUD](https://mud.dev/) is difficult today. To solve this proble, the MUDVRF module can get you setup with randomness in your application within minutes.

*This project was built during the Autonomous Worlds Hackathon hosted by EthGlobal in 2023*.

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

## Security 

The randomness generated relies on a 1/N honesty assumption between the underlying chain operators (i.e., validators or sequencers) and the party holding the secret key corresponding to the VRF (i.e, the prover). This is the same security model used by Chainlink in production for their VRF. In the future, an MPC protocol over the VRF can be used to seamlessly add more parties to the security model.

Our verifiable random function (VRF) implementation is based on the ECVFR (Elliptic Curve Verifiable Random Function) specification, which is an industry-standard approach to constructing VRFs based on elliptic curves. The ECVRF spec is an IETF (Internet Engineering Task Force) standard, and has undergone extensive review to ensure its security and reliability.

That said, the contracts are not fully audited, so please becareful when using this software in production. The only contracts that have been fully audited is the VRF verifier contract, which is a fork of Chainlink's ECVRF implementation.