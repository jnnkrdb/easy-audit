package client

import "github.com/spf13/cobra"

var (
	Insecure bool
)

func AddCategoryFlags(c *cobra.Command) {
	c.Flags().BoolVarP(&Insecure, "insecure", "k", false, "Allow insecure server connections when using SSL")
}

// add the default apiv1audits object flags to the given command
func AddAPIV1AuditsFlags(c *cobra.Command) {
	c.Flags().StringP("audit-id", "", "", "ID of the audit")
	c.Flags().StringP("audit-timestamp", "", "", "Timestamp of the audit")
	c.Flags().StringP("audit-action", "", "", "Action of the audit")
	c.Flags().StringP("audit-user", "", "", "User of the audit")
	c.Flags().StringP("audit-resource", "", "", "Resource of the audit")
	c.Flags().StringP("audit-result", "", "", "Result of the audit")
	c.Flags().StringP("audit-furtherinfo", "", "", "Further info of the audit")

	c.Flags().StringP("from-json", "", "", "Further info of the audit")
	c.Flags().StringP("from-jsonfile", "", "", "Further info of the audit")
}

// get the correct protocol based on the insecure flag
func getProtocol() string {
	if Insecure {
		return "http"
	} else {
		return "https"
	}
}
