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

	// add the default category flags to the post command
	AddCategoryFlags(PostCmd)
}

var PostCmd = &cobra.Command{
	Use:   "post",
	Short: "Post audits to the easy-audit server",
	Long:  `Post audits to the easy-audit server`,
	RunE: func(cmd *cobra.Command, args []string) error {

		host := fmt.Sprintf("%s://%s:%d", getProtocol(), cfg.Address, cfg.Port)

		audit, err := apiV1Audits.SendPost(cmd.Context(), host, tempAudit)
		if err != nil {
			return fmt.Errorf("failed to post audit: %w", err)
		}

		return format.WriteFormat(os.Stdout, format.FormatObject{Object: audit}, OutputFormat)
	},
}
