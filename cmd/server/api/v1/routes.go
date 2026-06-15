package v1

import (
	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/cmd/server/api/v1/audits"
)

// LoadRoutes registers the API routes on the provided router.
// It sets up a logging middleware to log details of each request and response,
// and then loads the specific routes for the key-value store API endpoints.
//
// Pathprefix: /v1
func LoadRoutes(ep *mux.Router) {

	_api := ep.PathPrefix("/v1/").Subrouter()

	audits.LoadRoutes(_api)
}
