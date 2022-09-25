package logs

import (
	"strings"

	envs "github.com/luisfernandomoraes/todo-list-golang/infraestructure/envs"
	"go.uber.org/zap"
)

var zapLog *zap.Logger

func init() {
	var err error

	value := envs.GetEnvironmentVariableByKey("ENVIRONMENT")
	if strings.ToUpper(value) == "PRODUCTION" {
		config := zap.NewProductionConfig()
		zapLog, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
	} else if strings.ToUpper(value) == "DEVELOPMENT" {
		config := zap.NewDevelopmentConfig()
		config.DisableStacktrace = true
		zapLog, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
	}

}
func Info(message string, fields ...zap.Field) {
	zapLog.Info(message, fields...)
}
func Debug(message string, fields ...zap.Field) {
	zapLog.Debug(message, fields...)
}
func Error(message string, fields ...zap.Field) {
	zapLog.Error(message, fields...)
}
func Fatal(message string, fields ...zap.Field) {
	zapLog.Fatal(message, fields...)
}
