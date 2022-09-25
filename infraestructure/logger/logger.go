package logger

import (
	"os"
	"strings"

	envs "github.com/luisfernandomoraes/todo-list-golang/infraestructure/envs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	value := envs.GetEnvironmentVariableByKey("ENVIRONMENT")
	if strings.ToUpper(value) == "DEVELOPMENT" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	}
}

func GetLogger() *zerolog.Logger {
	return &logger
}
