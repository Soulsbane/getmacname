package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/alexflint/go-arg"
)

func getMacAddressName(address string) {
	resp, err := http.Get("https://api.macvendors.com/" + address)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(body))
}

func isValidAddressFormat(address string) bool {
	match, _ := regexp.MatchString("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$", address)
	return match
}

func main() {
	var args ProgramArgs

	arg.MustParse(&args)

	if len(args.Address) > 0 {
		var isValidAddress = isValidAddressFormat(args.Address)

		if isValidAddress {
			getMacAddressName(args.Address)
		} else {
			fmt.Println("Invalid MAC Address!")
		}

	}
}
