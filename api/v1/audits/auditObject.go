package audits

import (
	"fmt"
	"slices"
	"strings"
)

type AuditRow struct {
	ID          string `json:"id"`
	Timestamp   int64  `json:"timestamp"`
	Action      string `json:"action"`
	User        string `json:"user"`
	Destination string `json:"destination"`
	Resource    string `json:"resource"`
	Outcome     string `json:"outcome"`
	FurtherInfo string `json:"further_info"`
}

func (a *AuditRow) Validate() error {
	if a.ID == "" {
		return fmt.Errorf("id is required")
	}

	if a.Timestamp == 0 {
		return fmt.Errorf("timestamp is required")
	}

	a.Action = strings.ToLower(a.Action)
	if !slices.Contains([]string{
		"read",
		"write",
		"delete",
	}, a.Action) {
		return fmt.Errorf("invalid action: %s", a.Action)
	}

	if a.User == "" {
		return fmt.Errorf("user is required")
	}

	return nil
}
