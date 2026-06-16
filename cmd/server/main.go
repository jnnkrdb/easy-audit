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
	"github.com/jnnkrdb/easy-audit/int/storage"
)

var (
	// http endpoint
	mx *mux.Router = mux.NewRouter()

	storageProvider = flag.String("storage-provider", "memory", "the storage provider to use for audits, options are: memory, database")
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

	store, err := storage.GetStorageProvider(*storageProvider, "")
	if err != nil {
		slog.Error("error initializing storage provider", "error", err)
		return
	}

	// register routes for audits
	api.LoadRoutes(mx, store)

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
