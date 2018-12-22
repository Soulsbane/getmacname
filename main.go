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
	}

	fmt.Println(resp.String())
}

func main() {
	app := cli.NewApp()
	app.Name = "getmacname"
	app.Usage = "HELP"

	err := app.Run(os.Args)

	if err != nil {
		getMacAddressName("FC:FB:FB:01:FA:21")
	}
}
