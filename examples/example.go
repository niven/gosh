package main

import (
	"fmt"

	"github.com/niven/gosh/githubaction"
	"github.com/niven/gosh/githubenv"
)

func main() {

	defaults, _ := githubenv.GetDefaultEnvironmentVariables()
	fmt.Printf("Default Env: %s\n", defaults.WorkflowRef)

	inputs, _ := githubaction.ReadInputs(defaults.WorkflowRef)
	fmt.Printf("example-input-string: %v\n", inputs["example-input-string"].Value)

}
