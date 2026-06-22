package get

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get audits from the easy-audit server",
	Long:  `Get audits from the easy-audit server`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement the logic to get audits from the easy-audit server
		slog.Debug("get audits not implemented yet...")
		slog.Info("get audits not implemented yet...")
		slog.Warn("get audits not implemented yet...")
		slog.Error("get audits not implemented yet...", "error", fmt.Errorf("get audits not implemented yet"))
	},
}
