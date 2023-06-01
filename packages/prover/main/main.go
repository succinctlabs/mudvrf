package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/jtguibas/mudvrf/server"
)

type VRFDeploymentJSON struct {
	VrfCoordinatorAddress string `json:"vrfCoordinatorAddress"`
	BlockHashStoreAddress string `json:"blockHashStoreAddress"`
}

type WorldDeploymentJSON struct {
	Foundry struct {
		Address string `json:"address"`
	} `json:"31337"`
}

func main() {
	// Load WorldDeploymentJSON.
	worldJson := WorldDeploymentJSON{}
	worldJsonFile, err := os.Open("../example-contracts/worlds.json")
	if err != nil {
		panic(err)
	}
	worldsJsonBytes, err := ioutil.ReadAll(worldJsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(worldsJsonBytes, &worldJson)
	if err != nil {
		panic(err)
	}
	fmt.Println("World Address:", worldJson.Foundry.Address)

	// Load VRFDeploymentJSON.
	vrfJson := VRFDeploymentJSON{}
	vrfJsonFile, err := os.Open("../example-contracts/vrf.json")
	if err != nil {
		panic(err)
	}
	vrfJsonBytes, err := ioutil.ReadAll(vrfJsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(vrfJsonBytes, &vrfJson)
	if err != nil {
		panic(err)
	}
	fmt.Println("VRFCoordinator Address:", vrfJson.VrfCoordinatorAddress)
	fmt.Println("BlockHashStore Address:", vrfJson.BlockHashStoreAddress)

	// Load enviroment.
	godotenv.Load()
	rpcUrl := os.Getenv("RPC_URL")
	privateKey := os.Getenv("PRIVATE_KEY")

	// Connect to RPC and start server.
	ethclient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}
	vrfCoordinatorAddress := common.HexToAddress(vrfJson.VrfCoordinatorAddress)
	blockHashStoreAddress := common.HexToAddress(vrfJson.BlockHashStoreAddress)
	server := server.New(privateKey, vrfCoordinatorAddress, blockHashStoreAddress, ethclient)
	server.Start()
}
