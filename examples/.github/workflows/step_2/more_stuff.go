package main

import (
	"fmt"

	"github.com/niven/gosh"
	"github.com/niven/gosh/github"
	"github.com/niven/gosh/integrations"
	"github.com/slack-go/slack"
)

func main() {

	g, err := gosh.New()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	step1Output := g.Environment["SOME_INT"]

	slack = integrations.Slackbot(g.Environment["SLACK_BOT_TOKEN"])

	slack.Info("C04H2AH6SAU", fmt.Sprintf("Step 2: output from step 1 = %s", step1Output))

	fmt.Println("Repo secrets: " + github.GetRepositorySecretList(fmt.Sprintf(github.GithubApiListRespositorySecrets, "", ""), ""))
}
