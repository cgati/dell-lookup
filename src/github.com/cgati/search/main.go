package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide at least one service tag.")
		fmt.Printf("Usage: %s 12345", args[0])
	} else {
		searchServiceTags(args[1:])
	}
}
