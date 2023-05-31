rm -rf bindings
mkdir bindings
abigen --abi ../packages/contracts/out/VRFCoordinator.sol/VRFCoordinator.abi.json --out ../backend/bindings/VRFCoordinator.go --pkg bindings --type VRFCoordinator
