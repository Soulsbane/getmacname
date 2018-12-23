package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/levigross/grequests"
	"github.com/urfave/cli"
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
	app := cli.NewApp()
	app.Name = "getmacname"
	app.Usage = "HELP"

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			var address = c.Args().Get(0)
			var isValidAddress = isValidAddressFormat(address)

			if isValidAddress == true {
				getMacAddressName(address)
			} else {
				fmt.Println("Invalid MAC Address!")
			}

		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
	}
}
