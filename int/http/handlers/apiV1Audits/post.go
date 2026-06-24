package apiV1Audits

import (
	"encoding/json"
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

// execute an http get request against the given url and return the response body as a byte slice
