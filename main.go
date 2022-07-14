package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ameyarao98/pong-bot/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	kovanEndpoint := os.Getenv("KOVAN_ENDPOINT")
	client, err := ethclient.Dial(kovanEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("block number %v\n", block.Number())

	ctr, err := contract.NewContract(common.HexToAddress(os.Getenv("PING_PONG_CONTRACT_ADDRESS")), client)
	if err != nil {
		log.Fatal(err)
	}

	channel := make(chan *contract.ContractPing)
	go func() {
		sub, err := ctr.WatchPing(&bind.WatchOpts{Context: context.Background(), Start: nil}, channel)
		if err != nil {
			log.Fatal(err)
		}
		defer sub.Unsubscribe()
	}()
	fmt.Println(<-channel)
}
