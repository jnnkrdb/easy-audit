package api

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// LoadRoutes registers the API routes on the provided router.
// It sets up a logging middleware to log details of each request and response,
// and then loads the specific routes for the key-value store API endpoints.
//
// Pathprefix: <...>/api/v1
func LoadRoutes(ep *mux.Router, store audits.AuditsStore) {

	slog.Info("setting audits store for API handlers")
	auditsService = audits.NewAuditsService(store)

	slog.Info("adding API routes for audits endpoints")
	_apiv1 := ep.PathPrefix("/api/v1/").Subrouter()

	_apiv1.HandleFunc("/audits", api_v1_audits_list).Methods(http.MethodGet)
	_apiv1.HandleFunc("/audits/{id}", api_v1_audits_get).Methods(http.MethodGet)
	_apiv1.HandleFunc("/audits", api_v1_audits_create).Methods(http.MethodPost)
	_apiv1.HandleFunc("/audits/{id}", api_v1_audits_update).Methods(http.MethodPut, http.MethodPatch)
	_apiv1.HandleFunc("/audits/{id}", api_v1_audits_delete).Methods(http.MethodDelete)
}
