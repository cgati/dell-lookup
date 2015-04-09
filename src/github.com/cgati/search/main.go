package main

import (
	"fmt"
	"os"
	"regexp"
)

func printWarranty(assets []Asset) {
	for _, asset := range assets {
		fmt.Println(asset.MachineDescription + " -- " + asset.ServiceTag)
		for _, warranty := range asset.Warranties {
			fmt.Println(warranty.ServiceLevelDescription)
			fmt.Println(warranty.StartDate + " -- " + warranty.EndDate + "\n")
		}
		fmt.Printf("\n\n")
	}
}

func main() {
	args := os.Args
	r := regexp.MustCompile("^[0-9a-zA-Z]{7}$")

	if len(args) < 2 {
		fmt.Println("Please provide at least one service tag.")
		fmt.Printf("Usage: %s 1A2B3C4", args[0])
		return
	}
	for _, arg := range args[1:] {
		if r.MatchString(arg) == false {
			fmt.Printf("One or more of your service tags was invalid. Please try again.")
			return
		}
	}
	content, err := searchServiceTags(args[1:])
	if err != nil {
		fmt.Printf("Sorry, an error occurred.")
		return
	}

	dellAssets := []Asset{}
	if len(args) == 2 {
		dellAssets, err = getWarrantyInformation(content, false)
		if err != nil {
			fmt.Println("there was an error processing your service tag")
			return
		}
	} else {
		dellAssets, err = getWarrantyInformation(content, true)
		if err != nil {
			fmt.Println("there was an error processing your service tags")
			return
		}
	}

	printWarranty(dellAssets)
}
