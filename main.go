package main

import (
	"fmt"
	"os"

	"github.com/chenzhijie/go-web3"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	kovanAPIKey := os.Getenv("KOVAN_API_KEY")
	web3, err := web3.NewWeb3("https://kovan.infura.io/v3/" + kovanAPIKey)
	if err != nil {
		panic(err)
	}
	blockNumber, err := web3.Eth.GetBlockNumber()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current block number: ", blockNumber)
}
