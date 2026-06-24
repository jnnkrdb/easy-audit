package apiV1Audits

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// receive the handler for the DELETE request to /api/v1/audits
func HandleDelete(auditsService *audits.AuditsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slog.Debug("api_v1_audits_delete called", "vars", vars)

		id, ok := vars["id"]
		if !ok {
			slog.Warn("id not provided in request")
			http.Error(w, "id not provided", http.StatusBadRequest)
			return
		}

		if err := auditsService.Delete(r.Context(), id); err != nil {
			slog.Error("failed to delete audit", "id", id, "error", err)
			http.Error(w, "failed to delete audit", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("deleted"))
	}
}

// execute an http get request against the given url and return the response body as a byte slice
