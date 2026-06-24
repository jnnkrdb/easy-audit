package apiV1Audits

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// receive the handler for the GET request to /api/v1/audits
func HandleList(auditsService *audits.AuditsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("api_v1_audits_list called")

		audits, err := auditsService.List(r.Context())
		if err != nil {
			slog.Error("failed to list audits", "error", err)
			http.Error(w, "failed to list audits", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(audits); err != nil {
			slog.Error("failed to encode audits", "error", err)
			http.Error(w, "failed to encode audits", http.StatusInternalServerError)
			return
		}
	}
}

// execute an http list request
func SendList(ctx context.Context, host string) (audits.AuditRows, error) {

	var auditRows = audits.AuditRows{}
	var url = fmt.Sprintf("%s%s", host, GetApiSubPath())

	// create a new HTTP GET request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return audits.AuditRows{}, fmt.Errorf("failed to create list request: %w", err)
	}

	// send the list request to the server and handle the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return audits.AuditRows{}, fmt.Errorf("failed to list audits: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return audits.AuditRows{}, fmt.Errorf("failed to list audits: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&auditRows); err != nil {
		return audits.AuditRows{}, fmt.Errorf("failed to decode audits: %w", err)
	}

	return auditRows, nil
}
