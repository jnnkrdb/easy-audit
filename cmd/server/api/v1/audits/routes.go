package audits

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// LoadRoutes registers the API routes on the provided router.
//
// Pathprefix: /audits
func LoadRoutes(ep *mux.Router) {

	_api := ep.PathPrefix("/audits").Subrouter()

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
