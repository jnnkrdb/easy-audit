package audits

import "context"

type AuditsService struct {
	store AuditsStore
}

func NewAuditsService(store AuditsStore) *AuditsService {
	return &AuditsService{
		store: store,
	}
}

func (s *AuditsService) List(ctx context.Context) (AuditRows, error) {
	return s.store.List(ctx)
}

func (s *AuditsService) Get(ctx context.Context, id string) (AuditRow, bool, error) {
	return s.store.Get(ctx, id)
}

func (s *AuditsService) Create(ctx context.Context, audit AuditRow) (AuditRow, error) {
	if err := audit.Validate(); err != nil {
		return AuditRow{}, err
	}
	return s.store.Create(ctx, audit)
}

func (s *AuditsService) Update(ctx context.Context, id string, audit AuditRow) (AuditRow, error) {
	if err := audit.Validate(); err != nil {
		return AuditRow{}, err
	}
	return s.store.Update(ctx, id, audit)
}

func (s *AuditsService) Delete(ctx context.Context, id string) error {
	return s.store.Delete(ctx, id)
}
