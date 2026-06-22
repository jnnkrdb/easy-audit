/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/jnnkrdb/easy-audit/cmd/cli/cfg"
	"github.com/jnnkrdb/easy-audit/cmd/cli/cmd/delete"
	"github.com/jnnkrdb/easy-audit/cmd/cli/cmd/get"
	"github.com/jnnkrdb/easy-audit/cmd/cli/cmd/list"
	"github.com/jnnkrdb/easy-audit/cmd/cli/cmd/patch"
	"github.com/jnnkrdb/easy-audit/cmd/cli/cmd/post"
	"github.com/jnnkrdb/easy-audit/int/logging"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eactl",
	Short: "A cli tool to interact with the easy-audit server",
	Long: `eactl is a command-line interface (CLI) tool designed to interact with the easy-audit server. 
	It provides various commands to manage and retrieve audit data from the server, 
	allowing users to efficiently monitor and analyze their audit logs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		logging.InitLogger(cfg.LOG_Level, cfg.LOG_Verbose, cfg.LOG_Format)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&cfg.EA_ServerUrl, "ea-server-address", "a", "http://localhost:80", "Set the URL of the easy-audit server to connect to.")
	rootCmd.PersistentFlags().StringVarP(&cfg.LOG_Level, "log-level", "", "error", "Set the log level to either debug, info, warn or error.")
	rootCmd.PersistentFlags().BoolVarP(&cfg.LOG_Verbose, "verbose", "v", false, "Prints the source of logs when set to true.")
	rootCmd.PersistentFlags().StringVarP(&cfg.LOG_Format, "log-format", "", "text", "Set the log format to either text or json.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(post.PostCmd)
	rootCmd.AddCommand(patch.PatchCmd)
	rootCmd.AddCommand(delete.DeleteCmd)
}
