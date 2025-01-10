package config

import (
	"errors"
)

type LogConfig struct {
	Format string `mapstructure:"format"`
	Level  string `mapstructure:"level"`
}

func (c *LogConfig) GetFormat() string {
	return c.Format
}

func (c *LogConfig) GetLevel() string {
	return c.Level
}

// Validate validates the Log configuration.
func (c *LogConfig) Validate() error {
	// Check if the format is one of the allowed values.
	validFormats := map[string]bool{
		"json": true,
		"text": true,
	}
	if !validFormats[c.Format] {
		return errors.New("format must be one of: json, text")
	}

	// Check if the level is one of the allowed values.
	validLevels := map[string]bool{
		"debug": true,
		"error": true,
		"info":  true,
		"warn":  true,
	}
	if !validLevels[c.Level] {
		return errors.New("level must be one of: debug, error, info, warn")
	}

	return nil
}

func DefaultLogConfig() LogConfig {
	return LogConfig{
		Format: "text",
		Level:  "info",
	}
}
