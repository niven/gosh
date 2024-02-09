package main

import (
	"fmt"

	"github.com/niven/gosh"
	"github.com/niven/gosh/integrations"
	"github.com/slack-go/slack"
)

func main() {

	g, err := gosh.New()
	if err != nil {
		fmt.Printf("Error creating gosh: %v\n", err)
		return
	}

	fmt.Printf("Default Env: %v\n", g.Defaults)
	fmt.Printf("example-input-string-propagated: %v\n", g.Input["example-input-string-propagated"].Value)
	fmt.Printf("EXAMPLE_WORKFLOW_VAR = %s\n", g.Environment["EXAMPLE_WORKFLOW_VAR"])
	fmt.Printf("EXAMPLE_JOB_VAR = %s\n", g.Environment["EXAMPLE_JOB_VAR"])
	fmt.Printf("EXAMPLE_STEP_VAR = %s\n", g.Environment["EXAMPLE_STEP_VAR"])

	g.Output.Set("SOME_INT", 42)
	g.Output.Set("SOME_STRING", "Teenage Mutant Ninja Turtles")
	g.Output.Set("SOME_INPUT", g.Input["example-input-string-propagated"].Value)
	g.Output.Commit()

	slack = integrations.Slackbot(g.Environment["SLACK_BOT_TOKEN"])
	slack.Info("C04H2AH6SAU", "Hello from Go!")
}
