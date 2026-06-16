package logging

import (
	"flag"
	"log/slog"
	"os"
)

var (
	// logging
	logLevel   = flag.String("log-level", "error", "Set the log level (debug, info, warn, error)")
	logVerbose = flag.Bool("verbose", false, "Prints the source of logs when set to true.")
	logFormat  = flag.String("log-format", "text", "Set the log format (text, json)")
)

// initializes the logger with default settings.
func InitLogger() {

	// configure logging
	var opts = &slog.HandlerOptions{
		AddSource: logVerbose != nil && *logVerbose,
		Level:     slog.LevelError, // default log level is error, can be overridden by cli flag --log-level
	}
	switch *logLevel {
	case "info":
		opts.Level = slog.LevelInfo
	case "warn":
		opts.Level = slog.LevelWarn
	case "debug":
		opts.Level = slog.LevelDebug
	}

	if *logFormat == "json" {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, opts)))
	} else {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))
	}

	slog.Debug("log initialized", "level", opts.Level, "add_source", opts.AddSource)
}
