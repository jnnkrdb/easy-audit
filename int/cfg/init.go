package cfg

import (
	"github.com/jnnkrdb/easy-audit/int/files"
	"github.com/spf13/viper"
)

func InitConfig() error {

	viper.SetEnvPrefix("ea")
	viper.SetConfigName("ea")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("/opt/easy-audit/config")
	viper.AddConfigPath("/etc/easy-audit")
	viper.AddConfigPath("$HOME/.easy-audit")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv() // read in environment variables that match

	return nil
}

// saves the current configuration to the user config directory
func SaveCurrentConfig() error {
	return viper.SafeWriteConfigAs(files.UserConfigDir)
}

func init() {
	// set default values for configuration
	viper.SetDefault("ea.server.host", "localhost")
	viper.SetDefault("ea.server.port", 80)

	viper.SetDefault("log.level", "error")
	viper.SetDefault("log.verbose", false)
	viper.SetDefault("log.format", "text")
}
