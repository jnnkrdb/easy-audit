package providers

import (
	"sync"

	"github.com/jnnkrdb/easy-audit/api/v1/audits"
)

type Cache struct {
	db map[string]*audits.AuditRow
	mx sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		db: make(map[string]*audits.AuditRow),
		mx: sync.Mutex{},
	}
}

// required interface functions

func (c *Cache) List() ([]*audits.AuditRow, error) {
	c.mx.Lock()
	defer c.mx.Unlock()
	var res []*audits.AuditRow = make([]*audits.AuditRow, 0, len(c.db))
	for _, v := range c.db {
		res = append(res, v)
	}
	return res, nil
}

func (c *Cache) Read(id string) (*audits.AuditRow, error) {
	c.mx.Lock()
	defer c.mx.Unlock()
	return c.db[id], nil
}

func (c *Cache) Write(log *audits.AuditRow) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.db[log.ID] = log
	return nil
}

func (c *Cache) Delete(id string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.db, id)
	return nil
}
