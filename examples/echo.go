package main

import (
	"fmt"

	"github.com/niven/gosh/githubenv"
)

func main() {

	defaults, err := githubenv.GetDefaultEnvironmentVariables()

	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Workflow: %s\n", defaults.WorkflowRef)

}
