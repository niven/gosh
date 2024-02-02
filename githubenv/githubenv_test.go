package githubenv_test

import (
	"testing"

	"github.com/niven/gosh/env"
	"github.com/niven/gosh/githubenv"
)

func setDefaultEnvironmentVariables() {
	env.Create("CI", "true")
	env.Create("GITHUB_ACTION", "example-action")
	env.Create("GITHUB_ACTIONS", "false")
	env.Create("GITHUB_WORKFLOW_REF", "../example-github-action.yaml@refs/heads/main")

}

func TestGetGithubEnv(t *testing.T) {

	setDefaultEnvironmentVariables()

	_, err := githubenv.GetDefaultEnvironmentVariables()

	if err != nil {
		t.Fatalf(`env.GetDefaultEnvironmentVariables(): %v`, err)
	}

}
