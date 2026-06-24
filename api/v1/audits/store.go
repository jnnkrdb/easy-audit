package audits

import "context"

type AuditsStore interface {
	List(ctx context.Context) (AuditRows, error)
	Get(ctx context.Context, id string) (AuditRow, bool, error)
	Create(ctx context.Context, audit AuditRow) (AuditRow, error)
	Update(ctx context.Context, id string, audit AuditRow) (AuditRow, error)
	Delete(ctx context.Context, id string) error
}
