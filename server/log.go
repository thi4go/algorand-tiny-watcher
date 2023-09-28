package server

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

func initLogger() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: "2006-01-02 15:04:05",
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("[server ] | %v", i)
	}
	logger = log.Output(output)
}
