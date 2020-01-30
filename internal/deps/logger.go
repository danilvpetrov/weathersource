package deps

import (
	"log"
	"os"
)

// ProvideLogger provides log.Logger.
func ProvideLogger() *log.Logger {
	return log.New(
		os.Stdout,
		"",
		log.LstdFlags,
	)
}
