package env

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

/*
* Read all env vars except the github specific ones
 */
func ReadEnvironmentVariables() map[string]string {
	result := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.HasPrefix(pair[0], "GITHUB_") || strings.HasPrefix(pair[0], "RUNNER_") || strings.HasPrefix(pair[0], "INPUT_") {
			continue
		}
		result[pair[0]] = pair[1]
	}
	return result
}

// Opinions from here
// https://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18_10_02
func IsValidEnvironmentVariableName(name string) bool {

	r := regexp.MustCompile(`\A[a-zA-Z_]+[a-zA-Z0-9_]*\z`)
	return r.Match([]byte(name))
}

func Exists(name string) bool {
	_, exists := os.LookupEnv(name)

	return exists
}

// Update changes an environment variable
func Update(name, value string) error {

	if !IsValidEnvironmentVariableName(name) {
		return errors.New("Not a valid name: " + name)
	}

	if !Exists(name) {
		return errors.New("Variable does not exist: " + name)
	}
	return os.Setenv(name, value)
}

// Create make a new environment variable
func Create(name, value string) error {

	if !IsValidEnvironmentVariableName(name) {
		return errors.New("Not a valid name: " + name)
	}

	if Exists(name) {
		return errors.New("Variable already exists: " + name)
	}
	return os.Setenv(name, value)
}

// Read returns the value of the environment variable specified
func Read(name string) (string, error) {

	if !IsValidEnvironmentVariableName(name) {
		return "", errors.New("Not a valid name: " + name)
	}

	value, exists := os.LookupEnv(name)

	if !exists {
		return "", errors.New("No such environment variable: " + name)
	}

	return value, nil
}
