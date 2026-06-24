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

	// add the default category flags to the patch command
	AddCategoryFlags(PatchCmd)
}

var PatchCmd = &cobra.Command{
	Use:   "patch <id>",
	Short: "Patch audits on the easy-audit server",
	Long:  `Patch audits on the easy-audit server`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		id := args[0]
		host := fmt.Sprintf("%s://%s:%d", getProtocol(), cfg.Address, cfg.Port)

		audit, err := apiV1Audits.SendGet(cmd.Context(), host, id)
		if err != nil {
			return fmt.Errorf("failed to get audit: %w", err)
		}

		audit.UpdateFrom(tempAudit)

		audit, err = apiV1Audits.SendPatch(cmd.Context(), host, id, audit)
		if err != nil {
			return fmt.Errorf("failed to patch audit: %w", err)
		}

		return format.WriteFormat(os.Stdout, format.FormatObject{Object: audit}, OutputFormat)
	},
}
