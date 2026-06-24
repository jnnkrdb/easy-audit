package apiV1Audits

import (
	"encoding/json"
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

// execute an http get request against the given url and return the response body as a byte slice
