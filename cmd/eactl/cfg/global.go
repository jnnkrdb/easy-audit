package cfg

import "github.com/spf13/cobra"

// Host configs
var (
	Address string
	Port    int
)

// log configs
var (
	LogLevel   string
	LogFormat  string
	LogVerbose bool
)

func AddGlobalVars(c *cobra.Command) {

	// host configs
	c.PersistentFlags().StringVarP(&Address, "address", "a", "localhost",
		`If used for serving, this translates to the address on which the 
    server should listen for incoming requests. If used for 
    client operations, this translates to the address of the 
    server to which the client should connect.`)
	c.PersistentFlags().IntVarP(&Port, "port", "p", 80,
		`If used for serving, this translates to the port on which the 
    server should listen for incoming requests. If used for 
    client operations, this translates to the port of the 
    server to which the client should connect.`)

	// set the logging flags for the root command
	c.PersistentFlags().StringVarP(&LogLevel, "log-level", "", "error", "Set the log level to either debug, info, warn or error.")
	c.PersistentFlags().StringVarP(&LogFormat, "log-format", "", "text", "Set the log format to either text or json.")
	c.PersistentFlags().BoolVarP(&LogVerbose, "verbose", "v", false, "Prints the source of logs when set to true.")
}
