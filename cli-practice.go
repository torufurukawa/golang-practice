package main

// https://github.com/urfave/cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var language string

	app := cli.NewApp()
	app.Name = "show"
	app.Usage = "show chinco"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}

		if language == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	app.Run(os.Args)
}
