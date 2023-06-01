# MUDVRF

**A secure, fast, and easy-to-use random number generator (RNG) for on-chain games built using MUD.**

![mudvrf](https://github.com/jtguibas/mudvrf/assets/25734765/09db1c47-3053-47e3-868e-2f2240dbb8aa)

Accessing secure randomness on apps built ontop of [MUD](https://mud.dev/) is difficult today. To solve this proble, the MUDVRF module can get you setup with randomness in your application within minutes.

*This project was built during the Autonomous Worlds Hackathon hosted by EthGlobal in 2023*.

## Get Started

Install the MUDVRF dependencies into the package where your MUD contracts live.

```sh
pnpm add @succinctlabs/mudvrf-contracts
```

Deploy the MUDVRF contracts within your post deploy script (i.e., `PostDeploy.s.sol`). View this [script]() as a reference.
```solidity
function run(address worldAddress) external {
    // Load the private key from the `PRIVATE_KEY` environment variable (in .env)
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");

    // Deploy VRFCoordinator and set Coordinator
    vm.startBroadcast(deployerPrivateKey);
    address blockHashStore = address(new BlockHashStore());
    address coordinator = address(new VRFCoordinator(blockHashStore));
    IVRFCoordinatorSystem(worldAddress).mudvrf_VRFCoordinatorSy_setCoordinator(
        coordinator
    );
    vm.stopBroadcast();

    // Write deployment addresses to JSON
    string memory obj1 = "vrfCoordinatorDeployment";
    string memory finalJson = vm.serializeAddress(obj1, "vrfCoordinatorAddress", coordinator);
    finalJson = vm.serializeAddress(obj1, "blockHashStoreAddress", blockHashStore);
    vm.writeJson(finalJson, "./vrf.json");
}
```

Import and use the `VRFCoordinator` inside systems within your MUD project to request randomness. View an example [here]().
```solidity
function dealCard() internal returns (bytes32) {
    IVRFCoordinator coordinator = IVRFCoordinatorSystem(_world());
    bytes32 requestId = coordinator.mudvrf_VRFCoordinatorSy_requestRandomWords(
        ORACLE_ID,
        NB_WORDS,
        REQUEST_CONFIRMATIONS,
        CALLBACK_GAS_LIMIT,
        selector
    );
    return requestId;
}

function handleDealCards(bytes32 requestId, uint256[] randomWords) {
    // Your logic here!
    ...
}
```

Run a VRF prover. You must install Go 1.18+ to generate the proofs.
```
cd packages/prover
pnpm run dev
```

## Mock Proving

If you would like to avoid spending time installing Go, you can instead deploy the `MockVRFCoordinator` with `MOCK_VRF_COORDINATOR=true` and use the `mock-prover` package with `cd packages/mock-prover; pnpm run dev`. 

## Repo Structure

```
.
├── ...
├── packages                
│   ├── contracts           # Core implementation of MUD module and VRF
│   ├── prover              # Go implementation of VRF prover
│   ├── mock-prover         # Example usage of module in MUD Contracts
│   ├── example-contracts   # Example usage of module in a MUD project (solidity)
│   └── example-client      # Example usage of module in a MUD project (react)
└── ...
```

## Security 

The randomness generated relies on a 1/N honesty assumption between the underlying chain operators (i.e., validators or sequencers) and the party holding the secret key corresponding to the VRF (i.e, the prover). This is the same security model used by Chainlink in production for their VRF. In the future, an MPC protocol over the VRF can be used to seamlessly add more parties to the security model.

Our verifiable random function (VRF) implementation is based on the ECVFR (Elliptic Curve Verifiable Random Function) specification, which is an industry-standard approach to constructing VRFs based on elliptic curves. The ECVRF spec is an IETF (Internet Engineering Task Force) standard, and has undergone extensive review to ensure its security and reliability.

That said, the contracts are not fully audited, so please becareful when using this software in production. The only contracts that have been fully audited is the VRF verifier contract, which is a fork of Chainlink's ECVRF implementation.