package api

import (
	"log/slog"

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

	_apiv1.HandleFunc("/audits", api_v1_audits_list).Methods("GET")
	_apiv1.HandleFunc("/audits/{id}", api_v1_audits_get).Methods("GET")
	_apiv1.HandleFunc("/audits", api_v1_audits_write).Methods("POST", "PUT")
	_apiv1.HandleFunc("/audits/{id}", api_v1_audits_delete).Methods("DELETE")
}
