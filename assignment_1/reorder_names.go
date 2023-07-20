package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please input something.")
		return
	} else if len(args) <= 2 {
		fmt.Println("Not enough arguments, please input first name, last name and country code.")
		return
	} else if len(args) == 3 {
		firstName := args[0]
		lastName := args[1]
		countryCode := args[2]

		switch countryCode {
		case "VN":
			fmt.Printf("Output: %s %s", lastName, firstName)
		case "US":
			fmt.Printf("Output: %s %s", firstName, lastName)
		default:
			fmt.Println("Please input a supported country code (VN or US)")
		}

	} else {
		firstName := args[0]
		lastName := args[1]
		middleNames := strings.Join(args[2:len(args)-1], " ")
		countryCode := args[len(args)-1]

		switch countryCode {
		case "VN":
			fmt.Printf("Output: %s %s %s", lastName, middleNames, firstName)
		case "US":
			fmt.Printf("Output: %s %s %s", firstName, lastName, middleNames)
		default:
			fmt.Println("Please input a supported country code (VN or US)")
		}
	}
}
