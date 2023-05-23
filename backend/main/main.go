package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/jtguibas/mudvrf/server"
)

func main() {
	godotenv.Load()
	// load from env
	privateKey := os.Getenv("PRIVATE_KEY")
	rawAddr := os.Getenv("CONTRACT_ADDRESS")
	rpcUrl := os.Getenv("RPC_URL")

	ethclient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(rawAddr)

	server := server.New(privateKey, contractAddress, ethclient)
	server.Start()
}
