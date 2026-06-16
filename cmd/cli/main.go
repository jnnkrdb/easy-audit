package cli

import (
	"flag"
	"log/slog"
	"os"

	"github.com/jnnkrdb/easy-audit/int/logging"
)

// start the cli application
func main() {

	flag.Parse()

	logging.InitLogger()

	slog.Error("cli not implemented yet...")
	os.Exit(1)
}
