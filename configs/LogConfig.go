package configs

import (
	"fcm-sender/helper/env"
	"fcm-sender/helper/zlog"
	"sync"
)

var (
	loggerConfig     Logger
	loggerConfigOnce sync.Once
)

type Logger struct {
	Level    zlog.LoggerLevel
	LogStyle string
}

func DefaultLoggerConfigFromEnv() Logger {

	logLevel := DEBUG_LEVEL
	if env.GetEnvEnvironment() == "development" {
		logLevel = DEV_DEBUG_LEVEL
	}

	loggerConfigOnce.Do(func() {
		loggerConfig = Logger{
			Level:    zlog.GetLogLevelFromString(env.GetEnv("DEBUG_LEVEL", logLevel)),
			LogStyle: env.GetEnv("LOG_STYLE", LOG_STYLE),
		}
	})
	return loggerConfig
}
