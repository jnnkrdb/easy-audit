package storage

import (
	"fmt"
	"log/slog"

	"github.com/jnnkrdb/easy-audit/api/v1/audits"
	"github.com/jnnkrdb/easy-audit/int/storage/providers"
)

// GetStorageProvider returns an instance of the AuditsStore interface
// based on the provided storage provider name and driver.
// It supports "memory" and "database" providers, and returns an error
// if an unknown provider is specified.
func GetStorageProvider(provider string, driver string) (audits.AuditsStore, error) {

	slog.Debug("initializing storage provider",
		"provider", provider,
		"driver", driver,
	)

	switch fmt.Sprintf("%s/%s", provider, driver) {
	case "memory/cache":
		return providers.NewCache(), nil
	case "database/sqlite3":
		return providers.NewDatabase(nil), nil
	default:
		return nil, fmt.Errorf("unknown storage provider: %s/%s", provider, driver)
	}
}
