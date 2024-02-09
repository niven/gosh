package github

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/niven/gosh/env"
)

// See : https://docs.github.com/en/actions/learn-github-actions/variables#default-environment-variables
type GithubEnv struct {
	// Always set to true.
	CI bool `env:"CI"`

	// 	The name of the action currently running, or the id of a step. For example, for an action, __repo-owner_name-of-action-repo.
	Action string `env:"GITHUB_ACTION"`

	// Always set to true when GitHub Actions is running the workflow. You can use this variable to differentiate when tests are being run locally or by GitHub Actions.
	Actions bool `env:"GITHUB_ACTIONS"`

	// The path on the runner to the file that sets the current step's outputs from workflow commands. This file is unique to the current step and changes for each step in a job.
	OutputFilePath string `env:"GITHUB_OUTPUT"`

	// This should be a path to the file that is the current workflow, but this is not set (documentation is lying)
	// WorkflowRef string `env:"GITHUB_WORKFLOW_REF"`

	// Name of the current workflow
	Workflow string `env:"GITHUB_WORKFLOW"`

	// The working directory
	Workspace string `env:"GITHUB_WORKSPACE"`
}

func GetDefaultEnvironmentVariables() (GithubEnv, error) {

	result := GithubEnv{}

	instance := reflect.ValueOf(&result)
	typ := reflect.TypeOf(result)

	for i := 0; i < typ.NumField(); i++ {

		name := typ.Field(i).Tag.Get("env")
		value, err := env.Read(name)
		if err != nil {
			return result, errors.New(fmt.Sprintf("Error reading default variables: %v", err))
		}
		switch tp := typ.Field(i).Type.Kind(); tp {
		case reflect.String:
			instance.Elem().Field(i).SetString(value)
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				return result, errors.New(fmt.Sprintf("Unable to convert value '%s' to boolean", value))
			}
			instance.Elem().Field(i).SetBool(boolValue)
		default:
			return result, errors.New(fmt.Sprintf("Don't know how to set variable of type: %v", tp))
		}
	}

	return result, nil
}
