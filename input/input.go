package input

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/niven/gosh/env"
	"gopkg.in/yaml.v3"
)

type Input struct {
	Name         string
	Description  string
	Required     bool
	DefaultValue any
	ValueType    string
	Options      []string
	Value        any
}

// only read and unmarshall the YAML once
var yamldata map[string]interface{}

func inputFromYAMLMap(input string, data map[string]interface{}) (Input, error) {

	// the input data should be added to the env prefixed with "INPUT_"

	replacedInput := regexp.MustCompile("[^\\w]").ReplaceAll([]byte(strings.ToUpper(input)), []byte("_"))
	variable := fmt.Sprintf("INPUT_%s", replacedInput)
	fmt.Printf("Looking for %s\n", variable)
	value := data["default"]
	if env.Exists(variable) {
		envValue, err := env.Read(variable)
		if err != nil {
			return Input{}, err
		}
		value = envValue
	}
	return Input{
		Name:      input,
		ValueType: data["type"].(string),
		Value:     value,
	}, nil
}

func readYAML(workflowName string) error {

	// Since github doesn't actually pass the path of the current workflow, we have to use the name to
	// find the file
	matches, err := filepath.Glob("*.yaml")
	// fmt.Printf("Found %d yaml files\n", len(matches))
	var correctFile []byte
	for _, f := range matches {
		// fmt.Println("Checking: " + f)
		yamlFile, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		if bytes.Contains(yamlFile, []byte(workflowName)) {
			correctFile = yamlFile
			break
		}
	}

	if correctFile == nil {
		return errors.New(fmt.Sprintf("No workflow file found with name: %s", workflowName))
	}

	yamldata = make(map[string]interface{})
	err = yaml.Unmarshal(correctFile, yamldata)
	return err
}

// https://docs.github.com/en/enterprise-cloud@latest/actions/using-workflows/workflow-syntax-for-github-actions#onworkflow_dispatchinputsinput_idtype
/* Because inputs in a workflow are not automatically turned into INPUT_... env vars you should:

env:
  INPUT_FOO: ${{ inputs.foo }}

We should warn I guess if that is not the case. Quite annoying as inputs has all the metadata
*/
func Read(workflowName string) (map[string]Input, error) {
	err := readYAML(workflowName)
	if err != nil {
		return nil, err
	}
	result := make(map[string]Input)

	inputs := yamldata["on"].(map[string]interface{})["workflow_dispatch"].(map[string]interface{})["inputs"].(map[string]interface{})

	for k, v := range inputs {
		result[k], err = inputFromYAMLMap(k, v.(map[string]interface{}))
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func ReadEnvironmentVariables(workflowFilePath string) (map[string]string, error) {

	return nil, nil

}
