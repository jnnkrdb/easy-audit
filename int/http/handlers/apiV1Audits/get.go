package apiV1Audits

import (
	"encoding/json"
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

// execute an http get request against the given url and return the response body as a byte slice
