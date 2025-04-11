package logger

import (
	"avitotechBackend2025/configs"
	"log/slog"
)

type Logger struct {
	*slog.Logger
	config *configs.LoggerConf
}

func New(cfg *configs.LoggerConf) (*Logger, error) {
	output, err := cfg.GetOutput()
	if err != nil {
		return nil, err
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level:     cfg.ParseLevel(),
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
