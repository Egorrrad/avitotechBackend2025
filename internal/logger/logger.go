package logger

import (
	"log/slog"
)

type Logger struct {
	*slog.Logger
	config *Config
}

func New(cfg *Config) (*Logger, error) {
	output, err := cfg.getOutput()
	if err != nil {
		return nil, err
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level:     cfg.parseLevel(),
		AddSource: cfg.AddSource,
	}

	if cfg.JSONFormat || cfg.Env == "production" {
		handler = slog.NewJSONHandler(output, opts)
	} else {
		handler = slog.NewTextHandler(output, opts)
	}

	return &Logger{
		Logger: slog.New(handler),
		config: cfg,
	}, nil
}
