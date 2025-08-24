package logger

import (
	"log/slog"
	"os"

	"github.com/ksusonic/nitask/pkg/config"
)

func NewLogger(cfg config.LoggerConfig) *slog.Logger {
	w := os.Stdout
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     parseLevel(cfg.Level),
	}

	var h slog.Handler
	switch cfg.Format {
	case "json":
		h = slog.NewJSONHandler(w, options)
	default:
		h = slog.NewTextHandler(w, options)
	}

	return slog.New(h)
}

func parseLevel(lvl string) slog.Level {
	switch lvl {
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
