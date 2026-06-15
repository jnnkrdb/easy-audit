package api

import (
	"github.com/gorilla/mux"
	v1 "github.com/jnnkrdb/easy-audit/cmd/server/api/v1"
)

// LoadRoutes registers the API routes on the provided router.
// It sets up a logging middleware to log details of each request and response,
// and then loads the specific routes for the key-value store API endpoints.
//
// Pathprefix: /api
func LoadRoutes(ep *mux.Router) {

	_api := ep.PathPrefix("/api/").Subrouter()

	v1.LoadRoutes(_api)
}
