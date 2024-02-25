package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort       string `mapstructure:"SERVER_PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDbName   string `mapstructure:"POSTGRES_DB"`
}

// Setup all configs
func NewConfig() *Config {

	// Set the names of the configuration files and the paths to search for them.
	viper.SetConfigName("config")
	viper.AddConfigPath("./config") // Adjusted to look for app.yaml in the current directory

	// Enable Viper to read environment variables
	viper.AutomaticEnv()

	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	cfg := &Config{}
	// Unmarshal the config file into the Config struct
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return cfg
}
