package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bayu-aditya/apiautotest/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Api Auto Testing",
		Usage: "Make an api integration test become auto and easier",

		Action: func(c *cli.Context) error {
			fmt.Println("init program")
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run api test from input file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Usage:    "yaml input file",
						Required: true,
					},
				},
				Action: internal.ActionRun(),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
