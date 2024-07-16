package zlog

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// GCP severity 기준 정의
const (
	DisableLevel   = iota
	EmergencyLevel //Emergency 일 경우 zero log 의 Panic Level 로  표시
	AlertLevel     //Alert 일 경우 zero log 의 Panic Level 로 표시
	CriticalLevel  //Critical 일 경우 zero log 의 Fatal Level 로 표시
	ErrorLevel
	WarningLevel
	NoticeLevel //Notice 일 경우 zero log 의 WARING Level 로 표시
	InfoLevel
	DebugLevel
	DefaultLevel //Default 일 경우 zero log 의 Debug Level 로 표시
	TraceLevel
)

type LoggerLevel int

type Logger struct {
	logger       *zerolog.Logger
	logLevel     LoggerLevel
	displayStyle string
}

var log *Logger

var LevelToString = [...]string{
	"DISABLE",
	"EMERGENCY",
	"ALERT",
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
	"DEFAULT",
	"TRACE",
}

func (l LoggerLevel) Name() string { return LevelToString[l] }

func GetLogLevelFromString(name string) LoggerLevel {
	level := DisableLevel
	for i, v := range LevelToString {
		if v == name {
			level = i
			break
		}
	}
	return LoggerLevel(level)
}

func GetLevel() LoggerLevel {
	return log.logLevel
}

func GetStyle() string {
	return log.displayStyle
}

func parsingLevel(setLevel LoggerLevel) zerolog.Level {
	level := zerolog.InfoLevel
	switch setLevel {
	case DisableLevel:
		level = zerolog.Disabled
	case EmergencyLevel:
		fallthrough
	case AlertLevel:
		level = zerolog.PanicLevel
	case CriticalLevel:
		level = zerolog.FatalLevel
	case ErrorLevel:
		level = zerolog.ErrorLevel
	case WarningLevel:
		fallthrough
	case NoticeLevel:
		level = zerolog.WarnLevel
	case InfoLevel:
		level = zerolog.InfoLevel
	case DebugLevel:
		fallthrough
	case DefaultLevel:
		level = zerolog.DebugLevel
	case TraceLevel:
		level = zerolog.TraceLevel
	default:
	}
	return level
}

func Init(setLevel LoggerLevel, setLogStyle string) {

	level := parsingLevel(setLevel)

	//GCP Level
	zerolog.LevelFieldName = "severity"
	zerolog.LevelPanicValue = "EMERGENCY"
	//"ALERT"
	zerolog.LevelFatalValue = "CRITICAL"
	zerolog.LevelErrorValue = "ERROR"
	zerolog.LevelWarnValue = "WARNING"
	//"NOTICE"
	zerolog.LevelInfoValue = "INFO"
	zerolog.LevelDebugValue = "DEBUG"
	//"DEFAULT"

	zerolog.LevelTraceValue = "TRACE"
	zerolog.SetGlobalLevel(level)

	var logger zerolog.Logger
	if setLogStyle == "JSON" {
		//json case
		logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	} else {
		//pretty case
		logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	log = &Logger{
		logger:       &logger,
		logLevel:     setLevel,
		displayStyle: setLogStyle,
	}
}

func ChangeLevel(setLevel LoggerLevel) {
	level := parsingLevel(setLevel)
	log.logLevel = setLevel
	zerolog.SetGlobalLevel(level)
}

func ChangeStyle(setLogStyle string) {
	var logger zerolog.Logger
	if setLogStyle == "JSON" {
		//json case
		logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	} else {
		//pretty case
		logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}
	log.logger = &logger
}

// Output duplicates the global log and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return log.logger.Output(w)
}

// With creates a child log with the field added to its context.
func With() zerolog.Context {
	return log.logger.With()
}

// Level creates a child log with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return log.logger.Level(level)
}

// Sample returns a log with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return log.logger.Sample(s)
}

// Hook returns a log with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return log.logger.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	return log.logger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	return log.logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return log.logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return log.logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return log.logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return log.logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return log.logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return log.logger.Panic()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	log.logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	log.logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// Emergency starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Emergency() *zerolog.Event {
	return log.logger.Panic()
}

// Critical starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Critical() *zerolog.Event {
	return log.logger.Fatal()
}

func Alert() *zerolog.Event {
	if log.logLevel >= AlertLevel {
		return log.logger.Log().Str("severity", "ALERT")
	} else {
		return log.logger.Panic()
	}
}

func Notice() *zerolog.Event {
	if log.logLevel >= NoticeLevel {
		return log.logger.Log().Str("severity", "NOTICE")
	} else {
		return log.logger.Info()
	}
}

func Default() *zerolog.Event {
	if log.logLevel >= DefaultLevel {
		return log.logger.Log().Str("severity", "DEFAULT")
	} else {
		return log.logger.Trace()
	}
}
