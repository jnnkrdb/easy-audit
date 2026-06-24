package apiV1Audits

import (
	"context"
	"fmt"
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

// execute an http delete request
func SendDelete(ctx context.Context, host string, id string) error {

	var url = fmt.Sprintf("%s%s/%s", host, GetApiSubPath(), id)

	// create a new HTTP DELETE request
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	// send the delete request to the server and handle the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete audits: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete audits: %s", resp.Status)
	}

	return nil
}
