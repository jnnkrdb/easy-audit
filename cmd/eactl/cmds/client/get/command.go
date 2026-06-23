package get

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jnnkrdb/easy-audit/api/v1/audits"
	"github.com/jnnkrdb/easy-audit/cmd/eactl/cfg"
	"github.com/jnnkrdb/easy-audit/cmd/eactl/cmds/client"
	"github.com/spf13/cobra"
)

func init() {

	// add the default category flags to the get command
	client.AddCategoryFlags(GetCmd)

	GetCmd.Flags().StringP("id", "i", "", "The ID of the audit to retrieve. If not provided, all audits will be retrieved.")
}

var GetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get audits from the easy-audit server",
	Long:  `Get audits from the easy-audit server`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		api_v1_audits_url := fmt.Sprintf("%s://%s:%d/api/v1/audits/%s", client.GetProtocol(), cfg.Host, cfg.Port, args[0])

		if resp, err := http.DefaultClient.Get(api_v1_audits_url); err != nil {
			return fmt.Errorf("failed to get audits: %w", err)
		} else {
			defer resp.Body.Close()
			// Process the response here

			slog.Debug("received response",
				"response_code", resp.StatusCode,
				"response_body_length", resp.ContentLength,
				"resp_headers", resp.Header,
			)
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("failed to get audits: %s", resp.Status)
			}

			var audits = []audits.AuditRow{}
			if err := json.NewDecoder(resp.Body).Decode(&audits); err != nil {
				return fmt.Errorf("failed to decode response: %w", err)
			}

			// Print the audits
			for _, audit := range audits {
				slog.Info("audit",
					"id", audit.ID,
					"timestamp", audit.Timestamp,
					"action", audit.Action,
					"user", audit.User,
					"resource", audit.Resource,
					"result", audit.Result,
					"furhter_info", audit.FurtherInfo,
				)
			}
		}

		return nil
	},
}
