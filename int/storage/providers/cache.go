package providers

import (
	"context"
	"sync"

	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

// default implementation of the AuditsStore interface,
// using an in-memory map as the storage backend.
// This is not suitable for production use, but
// can be used for testing and development purposes.
type Cache struct {
	db map[string]audits.AuditRow
	mx sync.Mutex
}

// NewCache creates a new Cache instance with an empty
// database and a mutex for synchronization.
func NewCache() *Cache {
	return &Cache{
		db: make(map[string]audits.AuditRow),
		mx: sync.Mutex{},
	}
}

// required interface functions

func (c *Cache) List(ctx context.Context) ([]audits.AuditRow, error) {
	c.mx.Lock()
	defer c.mx.Unlock()
	var res []audits.AuditRow = make([]audits.AuditRow, 0, len(c.db))
	for _, v := range c.db {
		res = append(res, v)
	}
	return res, nil
}

func (c *Cache) Read(ctx context.Context, id string) (audits.AuditRow, error) {
	c.mx.Lock()
	defer c.mx.Unlock()
	audit, exists := c.db[id]
	if !exists {
		return audits.AuditRow{}, nil
	}
	return audit, nil
}

func (c *Cache) Write(ctx context.Context, audit audits.AuditRow) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.db[audit.ID] = audit
	return nil
}

func (c *Cache) Delete(ctx context.Context, id string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.db, id)
	return nil
}
