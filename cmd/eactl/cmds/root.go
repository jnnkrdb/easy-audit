package cmds

import (
	"github.com/jnnkrdb/easy-audit/cmd/eactl/cfg"
	"github.com/jnnkrdb/easy-audit/cmd/eactl/cmds/client"
	"github.com/jnnkrdb/easy-audit/cmd/eactl/cmds/server"
	"github.com/jnnkrdb/easy-audit/pkg/logging"
	"github.com/spf13/cobra"
)

var (

	// logging

	logLevel   string
	logFormat  string
	logVerbose bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "eactl",
	Short: "A cli tool to manage audit data",
	Long: `eactl is a command-line interface (CLI) tool designed to manage audit data. 
	It provides various commands to manage and retrieve audit data from the server, 
	allowing users to efficiently monitor and analyze their audit logs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		logging.InitLogger(cfg.LogLevel, cfg.LogFormat, cfg.LogVerbose)

		return nil
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	cfg.AddGlobalVars(RootCmd)

	// server operation
	RootCmd.AddCommand(server.ServeCmd)

	// client operations
	RootCmd.AddCommand(client.ListCmd)
	RootCmd.AddCommand(client.GetCmd)
	RootCmd.AddCommand(client.PostCmd)
	RootCmd.AddCommand(client.PatchCmd)
	RootCmd.AddCommand(client.DeleteCmd)
}
