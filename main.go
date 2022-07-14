package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	kovanEndpoint := os.Getenv("KOVAN_ENDPOINT")
	cl, err := ethclient.Dial(kovanEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	block, err := cl.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("block number %v", block.Number())
	// channel := make(chan *contract.ContractPing)
	// go func() {
	// 	sub, err := ctr.WatchDeposited(&bind.WatchOpts{Context: context.Background(), Start: nil}, channel)
	// 	defer sub.Unsubscribe()
	// }()
	// event := <-channel
}
