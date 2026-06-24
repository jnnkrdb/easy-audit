package audits

import (
	"fmt"
	"strings"
)

// used to return a list of audits in the API response
type AuditRows struct {
	Items []AuditRow `json:"items" yaml:"items"`
}

// print as text
func (ar AuditRows) GoString() string {
	var result strings.Builder
	for _, audit := range ar.Items {
		fmt.Fprintf(&result, "%s\n", audit.GoString())
	}
	return result.String()
}
