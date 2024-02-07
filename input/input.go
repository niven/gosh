package input

import (
	"fmt"
	"os"
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

func readYAML(file string) error {

	if yamldata != nil {
		return nil
	}

	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	yamldata = make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, yamldata)
	if err != nil {
		return err
	}
	return nil
}

// https://docs.github.com/en/enterprise-cloud@latest/actions/using-workflows/workflow-syntax-for-github-actions#onworkflow_dispatchinputsinput_idtype
/* Because inputs in a workflow are not automatically turned into INPUT_... env vars you should:

env:
  INPUT_FOO: ${{ inputs.foo }}

We should warn I guess if that is not the case. Quite annoying as inputs has all the metadata
*/
func Read(workflowFilePath string) (map[string]Input, error) {
	err := readYAML(workflowFilePath)
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
