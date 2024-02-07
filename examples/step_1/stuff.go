package main

import (
	"fmt"

	"github.com/niven/gosh"
)

func main() {

	g, err := gosh.New()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("Default Env: %vs\n", g.Defaults)
	fmt.Printf("example-input-string-propagated: %v\n", g.Input["example-input-string-propagated"].Value)
	fmt.Printf("EXAMPLE_WORKFLOW_VAR = %s", g.Environment["EXAMPLE_WORKFLOW_VAR"])
	fmt.Printf("EXAMPLE_JOB_VAR = %s", g.Environment["EXAMPLE_WORKFLOW_VAR"])
	fmt.Printf("EXAMPLE_STEP_VAR = %s", g.Environment["EXAMPLE_WORKFLOW_VAR"])

	g.Output.Set("SOME_INT", 42)
	g.Output.Set("SOME_STRING", "Teenage Mutant Ninja Turtles")
	g.Output.Set("SOME_INPUT", g.Input["example-input-string-propagated"].Value)
	g.Output.Commit()

	g.Slack.Info("C04H2AH6SAU", "Hello from Go!")
}
