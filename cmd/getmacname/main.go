package main

import (
	"fmt"
	"regexp"

	"github.com/alexflint/go-arg"
	"github.com/levigross/grequests"
)

func getMacAddressName(address string) {
	resp, err := grequests.Get("https://api.macvendors.com/"+address, nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(resp.String())
}

func isValidAddressFormat(address string) bool {
	match, _ := regexp.MatchString("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$", address)
	return match
}

func main() {
	var args struct {
		Address string `arg:"positional, required"`
	}

	arg.MustParse(&args)

	if len(args.Address) > 0 {
		var isValidAddress = isValidAddressFormat(args.Address)

		if isValidAddress == true {
			getMacAddressName(args.Address)
		} else {
			fmt.Println("Invalid MAC Address!")
		}

	}
}
