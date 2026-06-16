package audits

import "context"

type AuditsStore interface {
	List(ctx context.Context) ([]AuditRow, error)
	Get(ctx context.Context, id string) (AuditRow, bool, error)
	Write(ctx context.Context, audit AuditRow) (AuditRow, error)
	Delete(ctx context.Context, id string) error
}
