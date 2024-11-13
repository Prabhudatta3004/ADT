package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config struct holds the configuration data that we'll need throughout our application.
type Config struct {
	DataPath string
}

// LoadConfig reads configuration from a file or environment variables.
func LoadConfig() Config {
	// Configure Viper to read a config file named "config" (without extension)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv() // Enable reading from environment variables

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	return Config{
		DataPath: viper.GetString("data_path"), // Retrieve the value of "data_path"
	}
}
