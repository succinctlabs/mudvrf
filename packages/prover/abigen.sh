rm -rf bindings
mkdir bindings
abigen --abi ../contracts/out/VRFCoordinator.sol/VRFCoordinator.abi.json --out ./bindings/VRFCoordinator.go --pkg bindings --type VRFCoordinator
abigen --abi ../contracts/out/BlockHashStore.sol/BlockhashStore.abi.json --out ./bindings/BlockHashStore.go --pkg bindings --type BlockHashStore