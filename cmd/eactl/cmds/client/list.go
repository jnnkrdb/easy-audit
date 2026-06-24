package client

import (
	"fmt"
	"os"

	"github.com/jnnkrdb/easy-audit/cmd/eactl/cfg"
	"github.com/jnnkrdb/easy-audit/int/http/handlers/apiV1Audits"
	"github.com/jnnkrdb/easy-audit/pkg/format"
	"github.com/spf13/cobra"
)

func init() {

	// add the default category flags to the list command
	AddCategoryFlags(ListCmd)
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List audits from the easy-audit server",
	Long:  `List audits from the easy-audit server`,
	RunE: func(cmd *cobra.Command, args []string) error {

		host := fmt.Sprintf("%s://%s:%d", getProtocol(), cfg.Address, cfg.Port)

		audits, err := apiV1Audits.SendList(cmd.Context(), host)
		if err != nil {
			return fmt.Errorf("failed to list audits: %w", err)
		}

		return format.WriteFormat(os.Stdout, format.FormatObject{Object: audits}, OutputFormat)
	},
}
