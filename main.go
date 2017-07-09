package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "aliaslist"
	app.Usage = "show unused aliases."

	app.Action = func(c *cli.Context) error {
		if c.Args().Get(0) == "show" {
			out, err := exec.Command("sh", "-c", "source ~/.bashrc && alias").Output()

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(out))
			return nil
		}

		return nil
	}

	app.Run(os.Args)
}
