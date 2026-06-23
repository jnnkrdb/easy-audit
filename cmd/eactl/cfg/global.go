package cfg

import "github.com/spf13/cobra"

// Host configs
var (
	Host string
	Port int
)

// log configs
var (
	LogLevel   string
	LogFormat  string
	LogVerbose bool
)

func AddGlobalVars(c *cobra.Command) {

	// host configs
	c.Flags().StringVarP(&Host, "address", "h", "localhost",
		"the address on which the server should listen for incoming requests")
	c.Flags().IntVarP(&Port, "port", "p", 80,
		"the port on which the server should listen for incoming requests")

	// set the logging flags for the root command
	c.Flags().StringVarP(&LogLevel, "log-level", "", "error", "Set the log level to either debug, info, warn or error.")
	c.Flags().StringVarP(&LogFormat, "log-format", "", "text", "Set the log format to either text or json.")
	c.Flags().BoolVarP(&LogVerbose, "verbose", "v", false, "Prints the source of logs when set to true.")
}
