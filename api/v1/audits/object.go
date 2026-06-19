package audits

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type AuditRow struct {
	ID          string `json:"id"`
	Timestamp   string `json:"timestamp"`
	Action      string `json:"action"`
	User        string `json:"user"`
	Resource    string `json:"resource"`
	Result      string `json:"result"`
	FurtherInfo string `json:"further_info"`
}

// Validate checks if the AuditRow has all required fields and if the action is valid. It returns an error if any validation fails.
func (a *AuditRow) Validate() error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}

	a.Timestamp = time.Now().Format(time.RFC3339Nano)

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

	if a.Resource == "" {
		return fmt.Errorf("resource is required")
	}

	if a.Result == "" {
		return fmt.Errorf("result is required")
	}

	return nil
}
