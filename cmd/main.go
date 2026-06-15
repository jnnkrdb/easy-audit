package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/jnnkrdb/easy-audit/cmd/server"
)

var (

	// start in server mode if the flag is set, otherwise do nothing
	serverMode = flag.Bool("server", false, "Start in server mode")

	// logging
	logLevel   = flag.String("log-level", "error", "Set the log level (debug, info, warn, error)")
	logVerbose = flag.Bool("v", false, "Prints the source of logs when set to true.")
	logFormat  = flag.String("log-format", "text", "Set the log format (text, json)")
)

func main() {

	flag.Parse()

	// configure logging
	var opts = &slog.HandlerOptions{
		AddSource: logVerbose != nil && *logVerbose,
		Level:     slog.LevelError, // default log level is error, can be overridden by env var LOG_LEVEL
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

	// if set to server mode, start in server mode by default
	if *serverMode {
		// Start in server mode
		server.Start()
	}
}
