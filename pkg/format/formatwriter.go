package format

import (
	"fmt"
	"io"
)

// inteface to mark the required functions for any object that
// can be printed in text, json, or yaml format
type FormatWriter interface {
	ToText() string
	ToJSON() ([]byte, error)
	ToYAML() ([]byte, error)
}

// print to console in the requested format (json, yaml, or text)
func WriteFormat(writer io.Writer, obj FormatWriter, format string) error {

	var output string = ""

	switch format {
	case "json":
		jsonOutput, err := obj.ToJSON()
		if err != nil {
			return fmt.Errorf("failed to convert incoming object to JSON: %w", err)
		}
		output = string(jsonOutput)

	case "yaml":
		yamlOutput, err := obj.ToYAML()
		if err != nil {
			return fmt.Errorf("failed to convert incoming object to YAML: %w", err)
		}
		output = string(yamlOutput)

	case "text":
		output = obj.ToText()

	default:
		return fmt.Errorf("invalid output format: %s", format)
	}
	_, err := fmt.Fprint(writer, output)
	return err
}
