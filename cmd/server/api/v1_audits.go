package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// implementation of the audits store, this will be used
// by the API handlers to interact with the storage layer
var auditsService *audits.AuditsService

// http functions

// list all audits
func api_v1_audits_list(w http.ResponseWriter, r *http.Request) {
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

// get a single audit by id
func api_v1_audits_get(w http.ResponseWriter, r *http.Request) {
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

// create a new audit
func api_v1_audits_create(w http.ResponseWriter, r *http.Request) {
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

// update an existing audit
func api_v1_audits_update(w http.ResponseWriter, r *http.Request) {
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

// delete an audit by id
func api_v1_audits_delete(w http.ResponseWriter, r *http.Request) {
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
