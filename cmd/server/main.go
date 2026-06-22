package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
	"github.com/jnnkrdb/easy-audit/cmd/server/api"
	"github.com/jnnkrdb/easy-audit/cmd/server/health"
	"github.com/jnnkrdb/easy-audit/int/logging"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// http endpoint
	mx *mux.Router = mux.NewRouter()

	databaseDriver = flag.String("database-driver", "sqlite3", "which database driver should be used for storage (e.g. sqlite3, postgres, mysql, etc.)")
	databaseDsn    = flag.String("database-dsn", "file:/opt/easy-audit/data/audits.db", "the data source name (DSN) for the database connection, format depends on the driver (e.g. for sqlite3: file:audits.db?cache=shared&mode=rwc, for postgres: user=postgres password=postgres dbname=audits sslmode=disable)")
)

const (
	serverPortHttp = 80
)

// start the server
func main() {

	flag.Parse()

	logging.InitLogger()

	slog.Info("connecting to database", "driver", *databaseDriver, "dsn", *databaseDsn)

	db, err := sql.Open(*databaseDriver, *databaseDsn)
	if err != nil {
		slog.Error("failed to connect to database", "driver", *databaseDriver, "dsn", *databaseDsn, "error", err)
		return
	}
	defer db.Close()

	store, err := audits.NewSQLTx(db)
	if err != nil {
		slog.Error("failed to initialize audits store", "error", err)
		return
	}

	// register routes for audits
	api.LoadRoutes(mx, store)

	// register routes for health checks
	health.LoadRoutes(mx)

	slog.Info("starting http server", "port", serverPortHttp)
	if err := (&http.Server{
		Handler:      mx,
		Addr:         fmt.Sprintf(":%d", serverPortHttp),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}).ListenAndServe(); err != nil {
		slog.Error("error keeping http server alive", "error", err)
		return
	}
}
