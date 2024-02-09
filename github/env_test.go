package github

import (
	"testing"

	"github.com/niven/gosh/env"
)

func setDefaultEnvironmentVariables() {
	env.Create("CI", "true")
	env.Create("GITHUB_ACTION", "example-action")
	env.Create("GITHUB_ACTIONS", "false")
}

func TestGetGithubEnv(t *testing.T) {

	setDefaultEnvironmentVariables()

	_, err := GetDefaultEnvironmentVariables()

	if err != nil {
		t.Fatalf(`env.GetDefaultEnvironmentVariables(): %v`, err)
	}

}
