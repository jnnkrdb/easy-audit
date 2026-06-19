package audits

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	// these constants can be used to identify the storage provider and driver in the configuration
	sql_fieldname_id           = "id"
	sql_fieldname_timestamp    = "timestamp"
	sql_fieldname_action       = "action"
	sql_fieldname_user         = "user"
	sql_fieldname_resource     = "resource"
	sql_fieldname_result       = "result"
	sql_fieldname_further_info = "further_info"

	sql_table_name = "audits_v1"
)

// this struct is a simple implementation of the AuditsStore interface using a SQL database
type SQLTx struct {
	db *sql.DB
}

// connect to a database or transaction and return an instance of the SQLTx struct, which is used
// to implement the AuditsStore interface
func NewSQLTx(db *sql.DB) (*SQLTx, error) {
	if db == nil {
		return nil, fmt.Errorf("db cannot be nil")
	}

	createTableQuery := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		%s TEXT PRIMARY KEY,
		%s TEXT NOT NULL,
		%s TEXT NOT NULL,
		%s TEXT NOT NULL,
		%s TEXT NOT NULL,
		%s TEXT NOT NULL,
		%s TEXT NOT NULL
	);
	`, sql_table_name,
		sql_fieldname_id,
		sql_fieldname_timestamp,
		sql_fieldname_action,
		sql_fieldname_user,
		sql_fieldname_resource,
		sql_fieldname_result,
		sql_fieldname_further_info,
	)

	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, fmt.Errorf("failed to create audits table: %w", err)
	}

	_, err := db.Exec("")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database and tables: %w", err)
	}

	return &SQLTx{
		db: db,
	}, nil
}

// required interface functions

// list the audit rows in the database and return them as a slice of AuditRow structs
func (s *SQLTx) List(ctx context.Context) ([]AuditRow, error) {
	var audits []AuditRow
	var query = fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s",
		sql_fieldname_id,
		sql_fieldname_timestamp,
		sql_fieldname_action,
		sql_fieldname_user,
		sql_fieldname_resource,
		sql_fieldname_result,
		sql_fieldname_further_info,
		sql_table_name,
	)

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list audit rows: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var audit AuditRow
		if err := rows.Scan(
			&audit.ID,
			&audit.Timestamp,
			&audit.Action,
			&audit.User,
			&audit.Resource,
			&audit.Result,
			&audit.FurtherInfo,
		); err != nil {
			return nil, fmt.Errorf("failed to scan audit row: %w", err)
		}
		audits = append(audits, audit)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over audit rows: %w", err)
	}

	return audits, nil
}

// get a single audit row by its ID and return it as an AuditRow struct, along with a boolean indicating if the row was found
func (s *SQLTx) Get(ctx context.Context, id string) (AuditRow, bool, error) {
	var audit AuditRow
	var query = fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = ?",
		sql_fieldname_id,
		sql_fieldname_timestamp,
		sql_fieldname_action,
		sql_fieldname_user,
		sql_fieldname_resource,
		sql_fieldname_result,
		sql_fieldname_further_info,
		sql_table_name,
		sql_fieldname_id,
	)

	row := s.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(
		&audit.ID,
		&audit.Timestamp,
		&audit.Action,
		&audit.User,
		&audit.Resource,
		&audit.Result,
		&audit.FurtherInfo,
	); err != nil {
		if err == sql.ErrNoRows {
			return AuditRow{}, false, nil
		}
		return AuditRow{}, false, fmt.Errorf("failed to scan audit row: %w", err)
	}
	return audit, true, nil
}

// create a new audit row in the database and return the created row as an AuditRow struct
func (s *SQLTx) Create(ctx context.Context, newAudit AuditRow) (AuditRow, error) {

	if err := newAudit.Validate(); err != nil {
		return AuditRow{}, fmt.Errorf("failed to validate new audit row: %w", err)
	}

	var query = fmt.Sprintf("INSERT INTO %s (%s, %s, %s, %s, %s, %s, %s) VALUES (?, ?, ?, ?, ?, ?, ?)",
		sql_table_name,
		sql_fieldname_id,
		sql_fieldname_timestamp,
		sql_fieldname_action,
		sql_fieldname_user,
		sql_fieldname_resource,
		sql_fieldname_result,
		sql_fieldname_further_info,
	)

	_, err := s.db.ExecContext(ctx, query,
		newAudit.ID,
		newAudit.Timestamp,
		newAudit.Action,
		newAudit.User,
		newAudit.Resource,
		newAudit.Result,
		newAudit.FurtherInfo,
	)
	if err != nil {
		return AuditRow{}, fmt.Errorf("failed to create new audit row: %w", err)
	}
	return newAudit, nil
}

// update an existing audit row in the database and return the updated row as an AuditRow struct
func (s *SQLTx) Update(ctx context.Context, id string, updatedAudit AuditRow) (AuditRow, error) {

	if err := updatedAudit.Validate(); err != nil {
		return AuditRow{}, fmt.Errorf("failed to validate updated audit row: %w", err)
	}

	var query = fmt.Sprintf("UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ?",
		sql_table_name,
		sql_fieldname_timestamp,
		sql_fieldname_action,
		sql_fieldname_user,
		sql_fieldname_resource,
		sql_fieldname_result,
		sql_fieldname_further_info,
		sql_fieldname_id,
	)

	_, err := s.db.ExecContext(ctx, query,
		updatedAudit.Timestamp,
		updatedAudit.Action,
		updatedAudit.User,
		updatedAudit.Resource,
		updatedAudit.Result,
		updatedAudit.FurtherInfo,
		id,
	)
	if err != nil {
		return AuditRow{}, fmt.Errorf("failed to update audit row: %w", err)
	}
	return updatedAudit, nil
}

// delete an audit row from the database by its ID
func (s *SQLTx) Delete(ctx context.Context, id string) error {
	var query = fmt.Sprintf("DELETE FROM %s WHERE %s = ?",
		sql_table_name,
		sql_fieldname_id,
	)

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete audit row: %w", err)
	}
	return nil
}
