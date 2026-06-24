package apiV1Audits

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// receive the handler for the PATCH request to /api/v1/audits
func HandlePatch(auditsService *audits.AuditsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slog.Debug("api_v1_audits_update called", "vars", vars)

		id, ok := vars["id"]
		if !ok {
			slog.Warn("id not provided in request")
			http.Error(w, "id not provided", http.StatusBadRequest)
			return
		}

		var audit audits.AuditRow
		if err := json.NewDecoder(r.Body).Decode(&audit); err != nil {
			slog.Error("failed to decode request body", "error", err)
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}

		res, err := auditsService.Update(r.Context(), id, audit)
		if err != nil {
			slog.Error("failed to update audit", "id", res.ID, "error", err)
			http.Error(w, "failed to update audit", http.StatusInternalServerError)
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
