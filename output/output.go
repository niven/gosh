package output

import (
	"fmt"
	"os"
)

type Output struct {
	filePath string
	data     map[string]any
}

func New(filePath string) Output {
	result := Output{filePath: filePath, data: make(map[string]any)}
	result.data = map[string]any{}
	return result
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
