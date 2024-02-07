package github.com/niven/gosh

import (
	"errors"
	"fmt"

	"github.com/niven/gosh/githubenv"
	"github.com/niven/gosh/input"
	"github.com/niven/gosh/output"
	"github.com/niven/gosh/slackbot"
)

type Gosh struct {
	Defaults    githubenv.GithubEnv
	Environment map[string]string
	Input       map[string]input.Input
	Output      output.Output
	Slack       slackbot.SlackBot
}

func New() (Gosh, error) {

	result := Gosh{}

	var err error

	result.Defaults, err = githubenv.GetDefaultEnvironmentVariables()
	if err != nil {
		return Gosh{}, errors.New(fmt.Sprintf("Unable to read defaults: %v"))
	}
	fmt.Printf("defaults: %v", result.Defaults)
	// result.Environment = env.ReadEnvironmentVariables()
	// result.Input, err = input.Read(result.Defaults.WorkflowRef)

	// result.Output = output.New(result.Defaults.OutputFilePath)
	// result.Output.Clear()

	// result.Slack = slackbot.New(result.Environment["SLACK_BOT_TOKEN"])

	return result, nil
}
