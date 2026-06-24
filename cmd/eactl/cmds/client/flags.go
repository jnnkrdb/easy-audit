package client

import (
	"github.com/jnnkrdb/easy-audit/api/v1/audits"
	"github.com/spf13/cobra"
)

var (

	// determines whether to use insecure connections to the server or not
	Insecure bool

	// output format for the command
	OutputFormat string
)

func AddCategoryFlags(c *cobra.Command) {
	c.Flags().BoolVarP(&Insecure, "insecure", "k", false, "Allow insecure server connections when using SSL")
	c.Flags().StringVarP(&OutputFormat, "output", "o", "text", "Output format: json, yaml, or text")
}

// used for temporary storage of the audit row flags
var tempAudit = audits.AuditRow{}

// add the default apiv1audits object flags to the given command
func AddAPIV1AuditsFlags(c *cobra.Command) {
	c.Flags().StringVarP(&tempAudit.Timestamp, "audit-timestamp", "", "", "Timestamp of the audit")
	c.Flags().StringVarP(&tempAudit.Action, "audit-action", "", "", "Action of the audit")
	c.Flags().StringVarP(&tempAudit.User, "audit-user", "", "", "User of the audit")
	c.Flags().StringVarP(&tempAudit.Resource, "audit-resource", "", "", "Resource of the audit")
	c.Flags().StringVarP(&tempAudit.Result, "audit-result", "", "", "Result of the audit")
	c.Flags().StringVarP(&tempAudit.FurtherInfo, "audit-furtherinfo", "", "", "Further info of the audit")
}

func AddInputFlags(c *cobra.Command) {
	c.Flags().StringP("from-json", "", "", "Input from json string")
	c.Flags().StringP("from-jsonfile", "", "", "Input from json file")
}

// get the correct protocol based on the insecure flag
func getProtocol() string {
	if Insecure {
		return "http"
	} else {
		return "https"
	}
}
