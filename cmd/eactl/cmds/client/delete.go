package client

import (
	"fmt"
	"net/http"

	"github.com/jnnkrdb/easy-audit/cmd/eactl/cfg"
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

		audit_id := args[0]

		api_v1_audits_url := fmt.Sprintf("%s://%s:%d/api/v1/audits/%s",
			getProtocol(),
			cfg.Address,
			cfg.Port,
			audit_id,
		)

		// create a new HTTP DELETE request
		req, err := http.NewRequest(http.MethodDelete, api_v1_audits_url, nil)
		if err != nil {
			return fmt.Errorf("failed to create delete request: %w", err)
		}

		// send the delete request to the server and handle the response
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to delete audits: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to delete audits: %s", resp.Status)
		}

		fmt.Println("Successfully deleted audit", audit_id)
		return nil
	},
}
