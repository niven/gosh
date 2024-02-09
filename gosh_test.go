package gosh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputs(t *testing.T) {

	_, err := Read(".", "No Such Workflow")
	assert.Error(t, err)

	inputs, err := Read("examples", "Example Go Workflow")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(inputs))

	expected := map[string]Input{
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
