package cmd

import (
	"log/slog"
	"os"

	"github.com/jnnkrdb/easy-audit/int/cfg"
	"github.com/jnnkrdb/easy-audit/int/files"
	"github.com/jnnkrdb/easy-audit/pkg/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "easy-audit",
	Short: "A server application to manage and retrieve audit data",
	Long:  `easy-audit is a server application designed to manage and retrieve audit data.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		logging.InitLogger(cfg.LOG_Level, cfg.LOG_Format, cfg.LOG_Verbose)

		// creates the local configuration directory if it does not exist
		if err := files.CreateDir(files.UserConfigDir, 0700); err != nil {
			slog.Debug("failed to create config directory", "error", err)
			return err
		}

		//reads the configuration from the local configuration directory
		viper.SetConfigFile(files.UserConfigDir + "/ea.json")
		if err := viper.ReadInConfig(); err != nil {
			slog.Debug("failed to read config file", "error", err)
		}

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
	rootCmd.PersistentFlags().StringVarP(&cfg.EA_ServerUrl, "ea-url", "H", "http://localhost:80", "Set the URL of the easy-audit server to connect to.")
	rootCmd.PersistentFlags().StringVarP(&cfg.LOG_Level, "log-level", "", "error", "Set the log level to either debug, info, warn or error.")
	rootCmd.PersistentFlags().BoolVarP(&cfg.LOG_Verbose, "verbose", "v", false, "Prints the source of logs when set to true.")
	rootCmd.PersistentFlags().StringVarP(&cfg.LOG_Format, "log-format", "", "text", "Set the log format to either text or json.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
