package gosh

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/niven/gosh/env"
	"github.com/niven/gosh/github"
	"gopkg.in/yaml.v3"
)

type Gosh struct {
	Defaults    github.GithubEnv
	Environment map[string]string
	Input       map[string]Input
	Output      Output
}

func (g *Gosh) Info(message string) {
	fmt.Printf("%s", message)
}
func (g *Gosh) Error(message string) {
	fmt.Printf("::error::%s", message)
}

func New() (Gosh, error) {

	result := Gosh{}

	var err error

	result.Defaults, err = github.GetDefaultEnvironmentVariables()
	if err != nil {
		return Gosh{}, errors.New(fmt.Sprintf("Unable to read defaults: %v", err))
	}
	result.Environment = env.ReadEnvironmentVariables()
	result.Input, err = Read(result.Defaults.Workspace, result.Defaults.Workflow)
	if err != nil {
		return Gosh{}, errors.New(fmt.Sprintf("Unable to parse inputs: %v", err))
	}

	result.Output = Output{filePath: result.Defaults.OutputFilePath, data: make(map[string]any)}
	result.Output.Clear()

	return result, nil
}

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

func readYAML(workspace, workflowName string) error {

	// Since github doesn't actually pass the path of the current workflow, we have to use the name to
	// find the file
	glob := fmt.Sprintf("%s/.github/workflows/*.yaml", workspace)
	matches, err := filepath.Glob(glob)
	dir, _ := os.Getwd()
	fmt.Printf("Found %d yaml files in %s%s\n", len(matches), dir, glob)
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
		return errors.New(fmt.Sprintf("No workflow file found with name: %s, looked in %s", workflowName, glob))
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
func Read(workspace, workflowName string) (map[string]Input, error) {
	err := readYAML(workspace, workflowName)
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

type Output struct {
	filePath string
	data     map[string]any
}

func (o Output) Clear() {
	o.data = make(map[string]any)
}

func (o Output) Set(name string, value any) {
	o.data[name] = value
}

func (o Output) Commit() error {

	f, err := os.Create(o.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for k, v := range o.data {
		_, err := f.WriteString(fmt.Sprintf("%s=%v", k, v))
		if err != nil {
			return err
		}
	}

	return nil
}
