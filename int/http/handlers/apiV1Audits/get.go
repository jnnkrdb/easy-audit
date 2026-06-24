package apiV1Audits

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// receive the handler for the GET request to /api/v1/audits/{id}
func HandleGet(auditsService *audits.AuditsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slog.Debug("api_v1_audits_get called", "vars", vars)

		id, ok := vars["id"]
		if !ok {
			slog.Warn("id not provided in request")
			http.Error(w, "id not provided", http.StatusBadRequest)
			return
		}

		// get the audit from the service
		audit, exists, err := auditsService.Get(r.Context(), id)

		// if there was an error getting the audit, return a 500 error
		if err != nil {
			slog.Error("failed to get audit", "id", id, "error", err)
			http.Error(w, "failed to get audit", http.StatusInternalServerError)
			return
		}

		// if the audit does not exist, return a 404 error
		if !exists {
			slog.Warn("audit not found", "id", id)
			http.Error(w, "audit not found", http.StatusNotFound)
			return
		}

		// send the response as JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(audit); err != nil {
			slog.Error("failed to encode audit", "id", id, "error", err)
			http.Error(w, "failed to encode audit", http.StatusInternalServerError)
			return
		}
	}
}

// execute an http get request
func SendGet(ctx context.Context, host string, id string) (audits.AuditRow, error) {

	var auditRow = audits.AuditRow{}
	var url = fmt.Sprintf("%s%s/%s", host, GetApiSubPath(), id)

	// create a new HTTP GET request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return audits.AuditRow{}, fmt.Errorf("failed to create get request: %w", err)
	}

	// send the get request to the server and handle the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return audits.AuditRow{}, fmt.Errorf("failed to get audit: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return audits.AuditRow{}, fmt.Errorf("failed to get audit: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&auditRow); err != nil {
		return audits.AuditRow{}, fmt.Errorf("failed to decode audit: %w", err)
	}

	return auditRow, nil
}
