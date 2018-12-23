package main

import (
	"fmt"
	"os"
	"strings"

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

func main() {
	app := cli.NewApp()
	app.Name = "getmacname"
	app.Usage = "HELP"

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			var address = c.Args().Get(0)

			// INFO Could probably use a regex here to determine the validity of the passed MAC Address.
			if strings.Count(address, ":") == 5 {
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
