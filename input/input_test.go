package input_test

import (
	"testing"

	"github.com/niven/gosh/input"
	"github.com/stretchr/testify/assert"
)

func TestReadInputs(t *testing.T) {

	_, err := input.Read("no-such-file.yaml")
	assert.Error(t, err)

	inputs, err := input.Read("../examples/example-github-action.yaml")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(inputs))

	expected := map[string]input.Input{
		"example-input-string-propagated": {
			Name:      "example-input-string-propagated",
			ValueType: "string",
			Value:     "Hello, Gosh!",
		},
		"example-input-string-not-propagated": {
			Name:      "example-input-string-not-propagated",
			ValueType: "string",
			Value:     "Gosh, Hello!",
		},
	}
	assert.EqualValues(t, expected, inputs)
}

// func TestReadEnvironmentVariables(t *testing.T) {

// }
