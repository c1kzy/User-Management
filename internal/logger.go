package internal

import (
	"os"
	"time"

	"github.com/phuslu/log"
)

func NewLogger() {
	log.DefaultLogger = log.Logger{
		Level:      log.InfoLevel,
		Caller:     0,
		TimeField:  "",
		TimeFormat: time.RFC850,
		Writer:     &log.ConsoleWriter{},
	}

	if log.IsTerminal(os.Stderr.Fd()) {
		log.DefaultLogger = log.Logger{
			TimeFormat: "15:04:05",
			Caller:     1,
			Writer: &log.ConsoleWriter{
				ColorOutput:    true,
				QuoteString:    true,
				EndWithMessage: true,
			},
		}
	}
}
