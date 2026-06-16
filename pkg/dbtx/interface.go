package dbtx

import (
	"context"
	"database/sql"
)

// this interface is used to satisfy both *sql.DB and *sql.Tx,
// so that we can use the same code for both transactions and
// non-transactional queries
type DBTx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
