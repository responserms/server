package config

import (
	"github.com/responserms/server/internal/log"
	"github.com/responserms/server/pkg/config/schema"
	"github.com/responserms/spec"
)

func NewFromFile(path string) (*Config, *spec.Diagnostics) {
	s := spec.New(schema.Schema())
	ctx := schema.NewContext()

	if diags := s.FileGlob(path); diags.HasErrors() {
		return nil, diags
	}

	// validate the file
	if diags := s.Parse(ctx); diags.HasErrors() {
		return nil, diags
	}

	// empty config struct
	cfg := newEmptyConfig()

	if diags := s.Decode(ctx, cfg); diags.HasErrors() {
		return nil, diags
	}

	// apply defaults
	cfg = applyDefaults(cfg)

	return cfg, &spec.Diagnostics{}
}

// ApplyLogLevelFromString applies the appropriate log level from the string-variant
// of the level.
func ApplyLogLevelFromString(cfg *Config, level string) {
	switch level {
	case "trace":
		cfg.LogLevel = log.Trace
	case "debug":
		cfg.LogLevel = log.Debug
	case "info":
		cfg.LogLevel = log.Info
	case "warn":
		cfg.LogLevel = log.Warn
	case "error":
		cfg.LogLevel = log.Error
	case "off":
		cfg.LogLevel = log.Off
	default:
		cfg.LogLevel = log.Info
	}
}
