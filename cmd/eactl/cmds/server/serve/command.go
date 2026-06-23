package serve

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
	"github.com/jnnkrdb/easy-audit/cmd/eactl/cmds/server/serve/api"
	"github.com/jnnkrdb/easy-audit/cmd/eactl/cmds/server/serve/health"
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
	ServeCmd.Flags().StringVarP(&databaseDriver, "database-driver", "db-d", "sqlite3",
		"which database driver should be used for storage (e.g. sqlite3, postgres, mysql, etc.)")

	ServeCmd.Flags().StringVarP(&databaseDsn, "database-dsn", "db-dsn", "file:audits.db",
		`the data source name for the database connection, format depends on the driver (e.g. sqlite3: file:audits.db)`)
}

var ServeCmd = &cobra.Command{
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
