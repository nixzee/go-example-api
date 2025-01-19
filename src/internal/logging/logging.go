package logging

import (
	"io"
	"log"
	"os"
)

var (
	Info         *log.Logger
	Warn         *log.Logger
	Error        *log.Logger
	Debug        *log.Logger
	DebugEnabled bool // Global debug flag
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.LUTC)
	Warn = log.New(os.Stderr, "WARN: ", log.Ldate|log.Ltime|log.LUTC)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.LUTC)

	// If DebugEnabled is true, initialize the Debug logger
	if DebugEnabled {
		Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.LUTC)
	} else {
		Debug = log.New(io.Discard, "", 0) // Discards log output when DebugEnabled is false
	}
}
