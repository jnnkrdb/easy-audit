package logging

import (
	"log/slog"
	"os"
)

// initializes the logger with default settings.
func InitLogger(level string, verbose bool, format string) {

	// configure logging
	var opts = &slog.HandlerOptions{
		AddSource: verbose,
		Level:     slog.LevelError, // default log level is error, can be overridden by cli flag --log-level
	}
	switch level {
	case "info":
		opts.Level = slog.LevelInfo
	case "warn":
		opts.Level = slog.LevelWarn
	case "debug":
		opts.Level = slog.LevelDebug
	}

	if format == "json" {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, opts)))
	} else {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))
	}

	slog.Debug("log initialized", "level", opts.Level, "add_source", opts.AddSource)
}
