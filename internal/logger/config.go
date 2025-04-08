package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Env        string `env:"APP_ENV" envDefault:"development"`
	Level      string `env:"LOG_LEVEL" envDefault:"info"`
	Output     string `env:"LOG_OUTPUT" envDefault:"stdout"`
	FilePath   string `env:"LOG_FILE_PATH" envDefault:"/var/log/app.log"`
	AddSource  bool   `env:"LOG_ADD_SOURCE" envDefault:"false"`
	JSONFormat bool   `env:"LOG_JSON" envDefault:"false"`
}

func (c *Config) parseLevel() slog.Level {
	switch strings.ToLower(c.Level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (c *Config) getOutput() (io.Writer, error) {
	switch c.Output {
	case "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	case "file":
		return os.OpenFile(c.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	default:
		return os.Stdout, nil
	}
}
