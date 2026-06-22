package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/jnnkrdb/easy-audit/int/logging"
)

var (

	// logging
	logLevel   = flag.String("log-level", "error", "Set the log level (debug, info, warn, error)")
	logVerbose = flag.Bool("verbose", false, "Prints the source of logs when set to true.")
	logFormat  = flag.String("log-format", "text", "Set the log format (text, json)")
)

// start the cli application
func main() {

	flag.Parse()

	logging.InitLogger(*logLevel, *logVerbose, *logFormat)

	slog.Error("cli not implemented yet...")
	os.Exit(1)
}
