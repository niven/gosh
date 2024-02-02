package env_test

import (
	"testing"

	"github.com/niven/gosh/env"
)

func TestEnvironmentVariableNames(t *testing.T) {

	validNames := []string{"VALID", "_VALID", "valid", "_VaLiD45", "_VALID_NAME_WITH_MORE_UNDERSCORES"}

	for i, s := range validNames {
		if env.IsValidEnvironmentVariableName(s) == false {
			t.Fatalf("Valid name considered invalid: [%d]: %s", i, s)
		}
	}

	invalidNames := []string{"5NOTVALID", "NOTVALID\tTABS", "NOT   VALID  SPACES", ""}

	for i, s := range invalidNames {
		if env.IsValidEnvironmentVariableName(s) == true {
			t.Fatalf("Invalid name considered valid: [%d]: %s", i, s)
		}
	}

}

func TestCreate(t *testing.T) {
	name := "TEST_SOME_ENV_VAR"
	expectedValue := "SOME_VALUE"
	err := env.Create(name, expectedValue)

	if err != nil {
		t.Fatalf("env.Set() should not return an error: %v", err)
	}

	actualValue, err := env.Read(name)

	if err != nil {
		t.Fatalf("env.Create() should not return an error: %v", err)
	}
	if actualValue != expectedValue {
		t.Fatalf("env.Read(): expected %s, actual %s", expectedValue, actualValue)
	}

	err = env.Create(name, expectedValue)
	if err == nil || err.Error() != "Variable already exists: TEST_SOME_ENV_VAR" {
		t.Fatalf("env.Create() should error when the variable already exists '%v'", err.Error())
	}

}

// TestNonExistingEnvVar calls env.Read with a name that does not exist, checking
// for a valid return value.
func TestNonExistingEnvVar(t *testing.T) {
	_, err := env.Read("NO_SUCH_ENV_VAR")
	if err.Error() != "No such environment variable: NO_SUCH_ENV_VAR" {
		t.Fatalf(`env.Get("NO_SUCH_ENV_VAR) should return the correct error`)
	}
}
