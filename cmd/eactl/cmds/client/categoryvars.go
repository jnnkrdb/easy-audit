package client

import "github.com/spf13/cobra"

var (
	Insecure bool
)

func AddCategoryFlags(c *cobra.Command) {
	c.Flags().BoolVarP(&Insecure, "insecure", "k", false, "Allow insecure server connections when using SSL")
}

// get the correct protocol based on the insecure flag
func getProtocol() string {
	if Insecure {
		return "http"
	} else {
		return "https"
	}
}
