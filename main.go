package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "aliaslist"
	app.Usage = "show unused aliases"

	app.Run(os.Args)
}
