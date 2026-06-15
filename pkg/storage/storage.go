package storage

import "github.com/jnnkrdb/easy-audit/api/v1/audits"

type Storage interface {
	// List retrieves all audit logs from the storage backend. Returns an error if there is an issue with retrieval.
	List() ([]*audits.AuditRow, error)

	// Read retrieves an audit log by its ID. Returns an error if the log is not found.
	Read(id string) (*audits.AuditRow, error)

	// Write saves the given audit log to the storage backend.
	Write(log *audits.AuditRow) error

	// Delete removes an audit log by its ID. Returns an error if the log is not found or if there is an issue with deletion.
	Delete(id string) error
}
