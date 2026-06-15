package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/cmd/server/api"
	"github.com/jnnkrdb/easy-audit/cmd/server/health"
)

var (

	// set the startuptime just for first debugging purposes, will be removed later
	startupTime = time.Now()

	// http endpoint
	mx *mux.Router = mux.NewRouter()
)

const (
	serverPortHttp = 80
)

func Start() {

	go func() {
		for {
			slog.Info("Hello World! Alive since...", "seconds", int(time.Since(startupTime).Seconds()))
			time.Sleep(time.Second * 30)
		}
	}()

	slog.Debug("add additional handlers to http backend")

	// register routes for health checks
	health.LoadRoutes(mx)

	// register routes for audits
	api.LoadRoutes(mx)

	if err := (&http.Server{
		Handler:      mx,
		Addr:         fmt.Sprintf(":%d", serverPortHttp),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe(); err != nil {
		slog.Error("error keeping http server alive", "error", err)
	}
}
