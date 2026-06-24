package audits

import (
	"fmt"
	"time"
)

// used to return a single audit in the API response
type AuditRow struct {

	// used to identify a specific audit row in the database
	ID string `json:"id" yaml:"id"`

	// data fields for the audit row
	Timestamp   string `json:"timestamp" yaml:"timestamp"`
	Action      string `json:"action" yaml:"action"`
	User        string `json:"user" yaml:"user"`
	Resource    string `json:"resource" yaml:"resource"`
	Result      string `json:"result" yaml:"result"`
	FurtherInfo string `json:"further_info" yaml:"furtherInfo"`

	// timestamps for when the audit row was created and last updated
	CreatedAt time.Time `json:"created_at" yaml:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" yaml:"updatedAt"`
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

func (a *AuditRow) UpdateFrom(t AuditRow) {
	if t.Timestamp != "" && t.Timestamp != a.Timestamp {
		a.Timestamp = t.Timestamp
	}
	if t.Action != "" && t.Action != a.Action {
		a.Action = t.Action
	}
	if t.User != "" && t.User != a.User {
		a.User = t.User
	}
	if t.Resource != "" && t.Resource != a.Resource {
		a.Resource = t.Resource
	}
	if t.Result != "" && t.Result != a.Result {
		a.Result = t.Result
	}
	if t.FurtherInfo != "" && t.FurtherInfo != a.FurtherInfo {
		a.FurtherInfo = t.FurtherInfo
	}
}

// print a string of the object
func (a AuditRow) GoString() string {
	return fmt.Sprintf("%#v", a)
	/*
		fmt.Sprintf("AuditRow [%s]\n\tID: \"%s\"\n\tTimestamp: \"%s\"\n\tAction: \"%s\"\n\tUser: \"%s\"\n\tResource: \"%s\"\n\tResult: \"%s\"\n\tFurtherInfo: \"%s\"\n\tCreatedAt: \"%s\"\n\tUpdatedAt: \"%s\"",
			a.ID,
			a.ID,
			a.Timestamp,
			a.Action,
			a.User,
			a.Resource,
			a.Result,
			a.FurtherInfo,
			a.CreatedAt.Format(time.RFC3339),
			a.UpdatedAt.Format(time.RFC3339),
		)
	*/
}
