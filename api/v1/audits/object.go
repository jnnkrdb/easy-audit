package audits

import (
	"fmt"
	"time"
)

// used to return a list of audits in the API response
type AuditRows struct {
	Items []AuditRow `json:"items"`
}

// used to return a single audit in the API response
type AuditRow struct {

	// used to identify a specific audit row in the database
	ID string `json:"id"`

	// data fields for the audit row
	Timestamp   string `json:"timestamp"`
	Action      string `json:"action"`
	User        string `json:"user"`
	Resource    string `json:"resource"`
	Result      string `json:"result"`
	FurtherInfo string `json:"further_info"`

	// timestamps for when the audit row was created and last updated
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate checks if the AuditRow has all required fields and if the action is valid. It returns an error if any validation fails.
func (a AuditRow) Validate() error {

	if a.Timestamp == "" {
		return fmt.Errorf("timestamp is required")
	}

	if a.Action == "" {
		return fmt.Errorf("action is required")
	}

	if a.User == "" {
		return fmt.Errorf("user is required")
	}

	if a.Resource == "" {
		return fmt.Errorf("resource is required")
	}

	if a.Result == "" {
		return fmt.Errorf("result is required")
	}

	return nil
}
