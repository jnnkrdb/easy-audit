package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// implementation of the audits store, this will be used
// by the API handlers to interact with the storage layer
var auditsStore audits.AuditsStore

// set the audits store for the API handlers,
// this should be called before starting the server
func SetAuditsStore(store audits.AuditsStore) {
	auditsStore = store
}

// API handlers for the audits endpoints

// Pathprefix: <...>/v1/audits
func apiv1Audits(ep *mux.Router) {
	_api := ep.PathPrefix("/v1/audits").Subrouter()

	_api.Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	_api.Methods("POST", "PUT").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	_api.Methods("DELETE").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
}
