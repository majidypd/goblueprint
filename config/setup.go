package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App app `mapstructure:"app"`
}

type app struct {
	Port string `mapstructure:"port"`
}

// Setup all configs
func NewConfig() *Config {
	// Set the names of the configuration files and the paths to search for them.
	viper.SetConfigName("app")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	cfg := &Config{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return cfg
}
