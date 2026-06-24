package apiV1Audits

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// api defining constants for the apiV1Audits package
const (
	APIContext = "api"
	ApiVersion = "v1"
	ApiObjects = "audits"
)

// return the used subpath for this api handling package
func GetApiSubPath() string {
	return fmt.Sprintf("/%s/%s/%s", APIContext, ApiVersion, ApiObjects)
}

// LoadRoutes registers the API routes on the provided router.
// It sets up a logging middleware to log details of each request and response,
// and then loads the specific routes for the key-value store API endpoints.
//
// Pathprefix: /api/v1/audits
func LoadRoutes(router *mux.Router, store audits.AuditsStore) {

	slog.Info("setting audits store for API handlers")
	auditsService := audits.NewAuditsService(store)

	slog.Info("adding API routes for audits endpoints")

	router.HandleFunc(GetApiSubPath(), HandleList(auditsService)).Methods(http.MethodGet)
	router.HandleFunc(GetApiSubPath()+"/{id}", HandleGet(auditsService)).Methods(http.MethodGet)
	router.HandleFunc(GetApiSubPath(), HandlePost(auditsService)).Methods(http.MethodPost)
	router.HandleFunc(GetApiSubPath()+"/{id}", HandlePatch(auditsService)).Methods(http.MethodPut, http.MethodPatch)
	router.HandleFunc(GetApiSubPath()+"/{id}", HandleDelete(auditsService)).Methods(http.MethodDelete)
}
