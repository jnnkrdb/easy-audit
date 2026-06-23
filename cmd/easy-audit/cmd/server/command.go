package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
	"github.com/jnnkrdb/easy-audit/cmd/easy-audit/cmd/server/api"
	"github.com/jnnkrdb/easy-audit/cmd/easy-audit/cmd/server/health"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// http endpoint
	mx *mux.Router = mux.NewRouter()

	// host configs
	serverAddress string
	serverPort    int

	// db configs
	databaseDriver string
	databaseDsn    string
)

func init() {

	// datasource configs
	ServerCmd.Flags().StringVarP(
		&databaseDriver,
		"database-driver",
		"db-d",
		"sqlite3",
		"which database driver should be used for storage (e.g. sqlite3, postgres, mysql, etc.)")
	ServerCmd.Flags().StringVarP(
		&databaseDsn,
		"database-dsn",
		"db-dsn",
		"file:/opt/easy-audit/data/audits.db",
		`the data source name (DSN) for the database connection, format depends on the driver 
		(e.g. for sqlite3: file:audits.db?cache=shared&mode=rwc, 
		for postgres: user=postgres password=postgres dbname=audits sslmode=disable)`)

	// host configs
	ServerCmd.Flags().StringVarP(
		&serverAddress,
		"server-address",
		"h",
		"0.0.0.0",
		"the address on which the server should listen for incoming requests")
	ServerCmd.Flags().IntVarP(
		&serverPort,
		"server-port",
		"p",
		80,
		"the port on which the server should listen for incoming requests")
}

var ServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the easy-audit server",
	Long:  `Start the easy-audit server with the specified args, to manage and retrieve audit data.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		slog.Info("connecting to database", "driver", databaseDriver, "dsn", databaseDsn)

		db, err := sql.Open(databaseDriver, databaseDsn)
		if err != nil {
			slog.Error("failed to connect to database", "driver", databaseDriver, "dsn", databaseDsn, "error", err)
			return err
		}
		defer db.Close()

		store, err := audits.NewSQLTx(db)
		if err != nil {
			slog.Error("failed to initialize audits store", "error", err)
			return err
		}

		// register routes for audits
		api.LoadRoutes(mx, store)

		// register routes for health checks
		health.LoadRoutes(mx)

		slog.Info("starting http server", "addr", serverAddress, "port", serverPort)

		if err := (&http.Server{
			Handler:      mx,
			Addr:         fmt.Sprintf("%s:%d", serverAddress, serverPort),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}).ListenAndServe(); err != nil {
			slog.Error("error keeping http server alive", "error", err)
			return err
		}
		return nil
	},
}
