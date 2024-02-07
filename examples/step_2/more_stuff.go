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

	step1Output := g.Environment["SOME_INT"]

	g.Slack.Info("C04H2AH6SAU", fmt.Sprintf("Step 2: output from step 1 = %s", step1Output))
}
