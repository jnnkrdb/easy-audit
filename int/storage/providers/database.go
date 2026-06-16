package providers

import (
	"context"
	"log/slog"

	"github.com/jnnkrdb/easy-audit/api/v1/audits"
	"github.com/jnnkrdb/easy-audit/pkg/dbtx"
)

type Database struct {
	dbtx dbtx.DBTx
}

func NewDatabase(dbtx dbtx.DBTx) *Database {
	return &Database{
		dbtx: dbtx,
	}
}

// required interface functions
func (d *Database) List(ctx context.Context) ([]audits.AuditRow, error) {
	slog.WarnContext(ctx, "List() not implemented")
	return nil, nil
}

func (d *Database) Get(ctx context.Context, id string) (audits.AuditRow, bool, error) {
	slog.WarnContext(ctx, "Get() not implemented")
	return audits.AuditRow{}, false, nil
}

func (d *Database) Write(ctx context.Context, audit audits.AuditRow) (audits.AuditRow, error) {
	slog.WarnContext(ctx, "Write() not implemented")
	return audits.AuditRow{}, nil
}

func (d *Database) Delete(ctx context.Context, id string) error {
	slog.WarnContext(ctx, "Delete() not implemented")
	return nil
}
