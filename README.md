# MUDVRF

**A secure, fast, and easy-to-use random number generator (RNG) for on-chain games built using MUD.**

![mudvrf](https://github.com/jtguibas/mudvrf/assets/25734765/09db1c47-3053-47e3-868e-2f2240dbb8aa)

Accessing secure randomness on apps built on top of [MUD](https://mud.dev/) can be challenging as you either have to roll your own solution or use on-chain pseudorandomness. To solve this proble, the MUDVRF module can get you setup with randomness generated from a VRF in your application within minutes.

*This project was built during the Autonomous Worlds Hackathon hosted by EthGlobal in 2023*.

## Get Started

Begin the MUD development server.
```sh
git clone https://github.com/succinctlabs/mudvrf.git
cd mudvrf
pnpm install
pnpm run dev
```
Once the development server is ready, in another terminal, run the typescript mock relayer. This relayer will not use the VRF and instead use randomness available from your operating system (to use the real VRF, set `USE_MOCK=false` in `packages/example-contracts/.env` but you will need to install Go 1.20+).

**Open `localhost:3000` in your browser and play some BlackJack!**

## Installing MUDVRF

Install the MUDVRF dependencies into the package where your MUD contracts live.

```sh
pnpm add @succinctlabs/mudvrf-contracts
```

Deploy the MUDVRF contracts within your post deploy script (i.e., `PostDeploy.s.sol`). View this [script](https://github.com/succinctlabs/mudvrf/blob/main/packages/example-contracts/script/PostDeploy.s.sol) as a reference.
```solidity
function run(address worldAddress) external {
    // Load the private key from the `PRIVATE_KEY` environment variable (in .env)
    uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
    bool useMock = vm.envBool("USE_MOCK");

    // Deploy MUDVRF contracts
    vm.startBroadcast(deployerPrivateKey);
    address blockHashStore = address(new BlockHashStore());
    address coordinator;
    if (useMock) {
        coordinator = address(new MockVRFCoordinator(blockHashStore));
        console.log("-----MOCK COORDINATOR ADDRESS-----");
    } else {
        console.log("-----COORDINATOR ADDRESS-----");
        coordinator = address(new VRFCoordinator(blockHashStore));
    }
    console.logAddress(coordinator);
    IVRFCoordinatorSystem(worldAddress).mudvrf_VRFCoordinatorSy_setCoordinator(coordinator);
    vm.stopBroadcast();

    string memory obj1 = "vrfCoordinatorDeployment";
    string memory finalJson = vm.serializeAddress(obj1, "vrfCoordinatorAddress", coordinator);
    finalJson = vm.serializeAddress(obj1, "blockHashStoreAddress", blockHashStore);
    vm.writeJson(finalJson, "./vrf.json");
}
```

Import and use the `VRFCoordinator` inside systems within your MUD project to request randomness. View an example [here](https://github.com/succinctlabs/mudvrf/blob/main/packages/example-contracts/src/systems/BlackJackSystem.sol#L53-L55).
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

Run your MUD development server and also run a VRF prover. You must install Go 1.18+ to generate the proofs if `USE_MOCK=false`.
```sh
pnpm run dev

# If USE_MOCK=false, run the following commands in a seperate terminal
cd packages/prover
pnpm run dev

# If USE_MOCK=true, run the following commands in a seperate terminal
cd packages/mock-prover
pnpm run dev
```

You will see an output like this.
```
World Address: 0x5FbDB2315678afecb367f032d93F642f64180aa3
VRFCoordinator Address: 0x7a2088a1bFc9d81c55368AE168C2C02570cB814F
BlockHashStore Address: 0x4A679253410272dd5232B3Ff7cF5dbB88f295319

Starting VRFRequestWatcher...
Oracle: c1ffd3cfee2d9e5cd67643f8f39fd6e51aad88f6f4ce6ab8827279cfffb92266
64

Listening for RequestRandomWords events...
```

## Repo Structure

```
.
├── ...
├── packages                
│   ├── contracts           # Core implementation of MUD module and VRF
│   ├── prover              # Go implementation of VRF prover
│   ├── mock-prover         # Typescript implementation of mock VRF prover
│   ├── example-contracts   # Example usage of module in a MUD project (solidity)
│   └── example-client      # Example usage of module in a MUD project (react)
└── ...
```

## Security 

This code has not yet been audited, and should not be used in any production systems.

The randomness generated relies on a 1/N honesty assumption between the underlying chain operators (i.e., validators or sequencers) and the party holding the secret key corresponding to the VRF (i.e, the prover). This is the same security model used by Chainlink in production for their VRF. In the future, an MPC protocol over the VRF can be used to seamlessly add more parties to the security model.

Our verifiable random function (VRF) implementation is based on the ECVFR (Elliptic Curve Verifiable Random Function) specification, which is an industry-standard approach to constructing VRFs based on elliptic curves. The ECVRF spec is an IETF (Internet Engineering Task Force) standard, and has undergone extensive review to ensure its security and reliability.

That said, the contracts are not fully audited, so please becareful when using this software in production. The only contracts that have been fully audited is the VRF verifier contract, which is a fork of Chainlink's ECVRF implementation.