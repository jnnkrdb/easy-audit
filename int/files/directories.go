package files

import (
	"fmt"
	"os"
)

const (
	// Directories
	ConfigDir     = "/opt/easy-audit/config"
	DataDir       = "/opt/easy-audit/data"
	EtcDir        = "/etc/easy-audit"
	UserConfigDir = "~/.easy-audit"
)

// CreateDir creates a directory if it does not exist
func CreateDir(dir string, permission os.FileMode) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, permission); err != nil {
				return fmt.Errorf("failed to create %s directory: %w", dir, err)
			}
		} else {
			return fmt.Errorf("failed to stat %s directory: %w", dir, err)
		}
	}

	return nil
}
