package internal

import (
	"fmt"

	"github.com/bayu-aditya/apiautotest/internal/core/input"
	"github.com/urfave/cli/v2"
)

func ActionRun() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		inputEngine := input.New()

		fmt.Println("running api test")
		fmt.Println("------------------------------------------------------------------------")

		inputFileLoc := c.String("input")
		fmt.Printf("Reading input file from: %s \n\n", inputFileLoc)
		inputFile, err := inputEngine.ReadYamlFromFs(inputFileLoc)
		if err != nil {
			return err
		}
		fmt.Printf("Project Name: %s \n", inputFile.ProjectName)
		fmt.Printf("HTTP Cases Num: %d \n", len(inputFile.Tests.Http.Cases))

		return nil
	}
}
