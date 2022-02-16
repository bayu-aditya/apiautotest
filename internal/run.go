package internal

import (
	"fmt"

	"github.com/bayu-aditya/apiautotest/internal/core"
	"github.com/bayu-aditya/apiautotest/internal/core/input"
	"github.com/bayu-aditya/apiautotest/internal/core/rest"
	"github.com/urfave/cli/v2"
)

func ActionRun() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		ctx := c.Context

		inputEngine := input.New()
		restEngine := rest.New()

		fmt.Println("running api test")
		fmt.Println("------------------------------------------------------------------------")

		inputFileLoc := c.String("input")
		fmt.Printf("Reading input file from: %s \n\n", inputFileLoc)

		inputFile, err := inputEngine.ReadYamlFromFs(inputFileLoc)
		if err != nil {
			return err
		}

		fmt.Printf("Project Name: %s \n", inputFile.ProjectName)

		// Looping for each case
		for _, httpCase := range inputFile.Tests.Http.Cases {
			fmt.Printf("Running Case: %s \n", httpCase.Name)

			// Looping for each job
			for _, job := range httpCase.Jobs {
				fmt.Printf("\tRunning Job: %s \n", job.Name)

				var restOutput *core.RestEngineGetOutput

				switch job.Request.Method {
				case "GET":
					restOutput, err = restEngine.Get(ctx, job.Request.Url, core.RestEngineGetInput{})
					if err != nil {
						return err
					}
				default:
					return fmt.Errorf("invalid method for case %s job %s", httpCase.Name, job.Name)
				}

				fmt.Printf("\t\tStatus Code: %d \n", restOutput.StatusCode)

				var warnings []string
				if job.Response.StatusCode != restOutput.StatusCode {
					warnings = append(warnings, fmt.Sprintf("Wrong Status Code, want %d got %d", job.Response.StatusCode, restOutput.StatusCode))
				}

				if len(warnings) > 0 {
					fmt.Printf("\t\tWARNING:\n")
					for _, warn := range warnings {
						fmt.Printf("\t\t- %s \n", warn)
					}
				} else {
					fmt.Printf("\t\tSUCCESS\n")
				}

				fmt.Printf("\n")
			}
		}

		return nil
	}
}
