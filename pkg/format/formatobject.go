package format

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// a generic object that can be formatted as text, json, or yaml
//
// requires that the object has json and yaml tags for the fields to
// be formatted correctly
type FormatObject struct {
	// the object to be formatted
	Object fmt.GoStringer
}

// format any object as a string
func (fo FormatObject) ToText() string {
	return fo.Object.GoString()
}

// format any object with json tags into json bytes
func (fo FormatObject) ToJSON() ([]byte, error) {
	return json.MarshalIndent(fo.Object, "", "  ")
}

// format any object with yaml tags into yaml bytes
func (fo FormatObject) ToYAML() ([]byte, error) {
	return yaml.Marshal(fo.Object)
}
