package client

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jnnkrdb/easy-audit/api/v1/audits"
	"github.com/jnnkrdb/easy-audit/cmd/eactl/cfg"
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

		api_v1_audits_url := fmt.Sprintf("%s://%s:%d/api/v1/audits",
			getProtocol(),
			cfg.Address,
			cfg.Port,
		)

		// send get request to the server and handle the response
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
					"further_info", audit.FurtherInfo,
				)
			}
		}

		return nil
	},
}
