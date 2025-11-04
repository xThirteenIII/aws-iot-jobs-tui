package config

import (
	"aws-iot-jobs-tui/where"
	"fmt"

	"github.com/spf13/viper"
)

func Setup() error {

	// Setup name and type of config files
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(where.Config())

	// Load config from file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file. Does the file exist?")
	}

	return nil
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetString returns the value associated with the key as an int64.
func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}
