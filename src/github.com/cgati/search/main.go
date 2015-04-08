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
		content, err := searchServiceTags(args[1:])
		if err != nil {
			fmt.Printf("Sorry, an error occurred.")
		} else {
			if len(args) == 2 {
				getWarrantyInformation(content, false)
			} else {
				getWarrantyInformation(content, true)
			}
		}
	}
}
