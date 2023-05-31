rm -rf bindings
mkdir bindings
abigen --abi ../packages/contracts/out/VRFCoordinator.sol/VRFCoordinator.abi.json --out ../backend/bindings/VRFCoordinator.go --pkg bindings --type VRFCoordinator
abigen --abi ../packages/contracts/out/BlockHashStore.sol/BlockhashStore.abi.json --out ../backend/bindings/BlockHashStore.go --pkg bindings --type BlockHashStore