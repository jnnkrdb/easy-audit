package server

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/cmd/server/api"
	"github.com/jnnkrdb/easy-audit/cmd/server/health"
	"github.com/jnnkrdb/easy-audit/int/logging"
)

var (
	// http endpoint
	mx *mux.Router = mux.NewRouter()
)

const (
	serverPortHttp = 80
)

// start the server
func main() {

	flag.Parse()

	logging.InitLogger()

	slog.Debug("add additional handlers to http backend")

	// register routes for health checks
	health.LoadRoutes(mx)

	// register routes for audits
	api.LoadRoutes(mx)

	slog.Info("starting http server", "port", serverPortHttp)
	if err := (&http.Server{
		Handler:      mx,
		Addr:         fmt.Sprintf(":%d", serverPortHttp),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe(); err != nil {
		slog.Error("error keeping http server alive", "error", err)
	}
}
