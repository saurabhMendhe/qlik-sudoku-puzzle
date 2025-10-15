package config

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	App AppConfig `env:", prefix=APP_"`
	Log LogConfig `env:", prefix=LOG_"`
}

type AppConfig struct {
	Environment string `env:"ENVIRONMENT, default=development"`
	Name        string `env:"NAME, default=qlik-sudoku-puzzle"`
	Version     string `env:"VERSION, default=1.0.0"`
}

type LogConfig struct {
	Level      string `env:"LEVEL, default=info"`
	Format     string `env:"FORMAT, default=console"`
	OutputPath string `env:"OUTPUT_PATH, default=stdout"`

	EnableCaller     bool `env:"ENABLE_CALLER, default=true"`
	EnableStacktrace bool `env:"ENABLE_STACKTRACE, default=false"`
}

func (c *Config) Validate() error {
	allowed := map[string]bool{
		"development": true,
		"test":        true,
		"production":  true,
	}
	if !allowed[c.App.Environment] {
		return fmt.Errorf("invalid environment: %s (allowed: development, test, production)", c.App.Environment)
	}

	return nil
}

func Load(ctx context.Context) (*Config, error) {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}

func (c Config) GetLoggable() (string, error) {
	configData, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return string(configData), nil
}

func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}

func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}

func (c *Config) IsTest() bool {
	return c.App.Environment == "test"
}
