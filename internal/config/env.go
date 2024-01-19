package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

// Config represents the configuration structure
type Config struct {
	PGHost     string `envconfig:"PG_HOST" validate:"required"`
	PGPort     int    `envconfig:"PG_PORT" default:"5432"`
	PGUser     string `envconfig:"PG_USER" validate:"required"`
	PGDB       string `envconfig:"PG_DB" validate:"required"`
	PGPassword string `envconfig:"PG_PASSWORD" validate:"required"`
}

// NewConfig creates a new Config instance
func NewConfig() *Config {
	return &Config{}
}

// Load loads the configuration from environment variables and validates it
func (c *Config) Load() (*Config, error) {
	// if err := loadEnv(); err != nil {
	// 	return nil, fmt.Errorf("failed to load environment variables: %v", err)
	// }

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}

	// Validate the configuration
	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// // loadEnv loads variables from a .env file
// func loadEnv() error {
// 	if err := godotenv.Load(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// validateConfig validates the configuration using the validator package
func validateConfig(cfg *Config) error {
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}
	return nil
}
