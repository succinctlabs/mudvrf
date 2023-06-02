package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/akamensky/argparse"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	vrfJSONPath := parser.String("v", "vrfJsonPath", &argparse.Options{Required: true, Help: "VRF deployment json path"})
	rpcUrl := parser.String("r", "rpcUrl", &argparse.Options{Required: true, Help: "Ethereum RPC url"})
	privateKey := parser.String("p", "privateKey", &argparse.Options{Required: true, Help: "Private key of the prover"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	vrfJson := VRFDeploymentJSON{}
	vrfJsonFile, err := os.Open(*vrfJSONPath)
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

	ethclient, err := ethclient.Dial(*rpcUrl)
	if err != nil {
		panic(err)
	}
	vrfCoordinatorAddress := common.HexToAddress(vrfJson.VrfCoordinatorAddress)
	blockHashStoreAddress := common.HexToAddress(vrfJson.BlockHashStoreAddress)
	server := server.New(*privateKey, vrfCoordinatorAddress, blockHashStoreAddress, ethclient)
	server.Start()
}
