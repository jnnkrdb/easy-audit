package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

	AddHandlers()

	if err := (&http.Server{
		Handler:      mx,
		Addr:         fmt.Sprintf(":%d", serverPortHttp),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe(); err != nil {
		slog.Error("error keeping http server alive", "error", err)
	}
}

// add other handlers to the http mux
func AddHandlers() {

	// add health checks to server
	mx.HandleFunc("/health/livez", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	}).Methods("GET")

	mx.HandleFunc("/health/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	}).Methods("GET")
}
