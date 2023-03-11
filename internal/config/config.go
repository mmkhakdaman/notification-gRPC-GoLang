package config

import (
	"github.com/spf13/viper"
)

// Config is the application configuration struct
type Config struct {
	Port string `mapstructure:"PORT"`
}

// LoadConfig loads the application configuration from environment variables
func LoadConfig() (*Config, error) {
	viper.SetDefault("PORT", "50051")

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
