package main

import (
	"fmt"
	"os"

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
			getMacAddressName(address)
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
	}
}
