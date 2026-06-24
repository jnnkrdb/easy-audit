package apiV1Audits

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// receive the handler for the POST request to /api/v1/audits
func HandlePost(auditsService *audits.AuditsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("api_v1_audits_create called")

		var audit audits.AuditRow
		if err := json.NewDecoder(r.Body).Decode(&audit); err != nil {
			slog.Error("failed to decode request body", "error", err)
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}

		res, err := auditsService.Create(r.Context(), audit)
		if err != nil {
			slog.Error("failed to create audit", "id", res.ID, "error", err)
			http.Error(w, "failed to create audit", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			slog.Error("failed to encode audit", "id", res.ID, "error", err)
			http.Error(w, "failed to encode audit", http.StatusInternalServerError)
			return
		}
	}
}

// execute an http post request
func SendPost(ctx context.Context, host string, audit audits.AuditRow) (audits.AuditRow, error) {

	var auditRow = audits.AuditRow{}
	var url = fmt.Sprintf("%s%s", host, GetApiSubPath())

	// create an io.Writer from the audit object to send as the request body
	body, err := json.Marshal(audit)
	if err != nil {
		return audits.AuditRow{}, fmt.Errorf("failed to marshal audit: %w", err)
	}

	// create a new HTTP POST request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return audits.AuditRow{}, fmt.Errorf("failed to create post request: %w", err)
	}

	// send the post request to the server and handle the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return audits.AuditRow{}, fmt.Errorf("failed to post audit: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return audits.AuditRow{}, fmt.Errorf("failed to post audit: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&auditRow); err != nil {
		return audits.AuditRow{}, fmt.Errorf("failed to decode audit: %w", err)
	}

	return auditRow, nil
}
