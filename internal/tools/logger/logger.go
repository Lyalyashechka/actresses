package logger

import (
	"context"

	context_lib "github.com/Lyalyashechka/actresses/internal/tools/context"
	"github.com/sirupsen/logrus"
)

type Level string

const (
	LevelError Level = "error"
	LevelWarn  Level = "warn"
	LevelInfo  Level = "info"
	LevelDebug Level = "debug"
)

var configLevels = map[Level]logrus.Level{
	LevelError: logrus.ErrorLevel,
	LevelWarn:  logrus.WarnLevel,
	LevelInfo:  logrus.InfoLevel,
	LevelDebug: logrus.DebugLevel,
}

type Config struct {
	Level Level  `mapstructure:"level"`
	App   string `mapstructure:"app"`
}

type Logger interface {
	WithCtx(ctx context.Context) Logger
	WithField(key string, val interface{}) Logger
	WithFields(fields context_lib.Fields) Logger
	WithError(err error) Logger
	//Debug(args ...interface{})
	//Debugf(format string, args ...interface{})
	Info(args ...interface{})
	//Infof(format string, args ...interface{})
	//Warn(args ...interface{})
	//Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	//Fatal(args ...interface{})
	//Fatalf(format string, args ...interface{})
	//Print(args ...interface{})
	//Printf(format string, args ...interface{})
	//Log(level Level, args ...interface{})
	//Logf(level Level, format string, args ...interface{})
	WithCaller(ctx context.Context) context.Context
}
