package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	kovanAPIKey := os.Getenv("KOVAN_API_KEY")
	fmt.Println("https://kovan.infura.io/v3/" + kovanAPIKey)
}
