package client

import (
	"fmt"

	"github.com/jnnkrdb/easy-audit/cmd/eactl/cfg"
	"github.com/jnnkrdb/easy-audit/int/http/handlers/apiV1Audits"
	"github.com/spf13/cobra"
)

func init() {

	// add the default category flags to the delete command
	AddCategoryFlags(DeleteCmd)
}

var DeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete audits from the easy-audit server",
	Long:  `Delete audits from the easy-audit server`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		id := args[0]
		host := fmt.Sprintf("%s://%s:%d", getProtocol(), cfg.Address, cfg.Port)

		if err := apiV1Audits.SendDelete(cmd.Context(), host, id); err != nil {
			return fmt.Errorf("failed to delete audits: %w", err)
		}

		fmt.Println("Successfully deleted audit", id)
		return nil
	},
}
