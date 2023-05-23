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

type WorldsJson struct {
	Foundry struct {
		Address string `json:"address"`
	} `json:"31337"`
}

func main() {

	// echo PWD
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	// load json from ../packages/worlds.json
	worldsJson := WorldsJson{}
	jsonFile, err := os.Open("../packages/contracts/worlds.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(byteValue))
	json.Unmarshal(byteValue, &worldsJson)
	fmt.Printf("%+v\n", worldsJson)

	godotenv.Load()
	// load from env
	privateKey := os.Getenv("PRIVATE_KEY")
	// rawAddr := os.Getenv("CONTRACT_ADDRESS")
	rpcUrl := os.Getenv("RPC_URL")
	rawAddr := worldsJson.Foundry.Address

	ethclient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(rawAddr)

	server := server.New(privateKey, contractAddress, ethclient)
	server.Start()
}
