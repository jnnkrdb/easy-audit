package cfg

var (

	// easy-audit connection and configuration flags
	EA_ServerUrl string

	// logging
	LOG_Level   string //flag.String("log-level", "error", "Set the log level (debug, info, warn, error)")
	LOG_Verbose bool   //flag.Bool("verbose", false, "Prints the source of logs when set to true.")
	LOG_Format  string //flag.String("log-format", "text", "Set the log format (text, json)")
)
