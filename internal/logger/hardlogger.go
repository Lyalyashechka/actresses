package logger

import (
	"context"
	"os"
	"runtime"
	"strings"
	"time"

	context_lib "github.com/Lyalyashechka/actresses/internal/context"
	"github.com/sirupsen/logrus"
)

type hardLogger struct {
	*logrus.Entry
}

func NewLogrusLogger(cfg Config) (Logger, error) {
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(configLevels[cfg.Level])
	logrusLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	l := logrusLogger.WithField("app", cfg.App)

	hostname, err := os.Hostname()
	if err == nil {
		l = l.WithField("host", hostname)
	}

	return hardLogger{Entry: l}, nil
}

func (l hardLogger) WithCtx(ctx context.Context) Logger {
	return l.WithFields(context_lib.GetCtxFields(ctx))
}

func (l hardLogger) WithCaller(ctx context.Context) context.Context {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		funcInfo := strings.Split(details.Name(), ".")

		ctx = context_lib.AddCtxFields(ctx, context_lib.Fields{
			"method": funcInfo[len(funcInfo)-1],
		})
	}

	return ctx
}

func (l hardLogger) WithFields(fields context_lib.Fields) Logger {
	if fields == nil {
		return l
	}

	return hardLogger{Entry: l.Entry.WithFields(logrus.Fields(fields))}
}

func (l hardLogger) WithField(key string, val interface{}) Logger {
	return hardLogger{Entry: l.Entry.WithField(key, val)}
}

func (l hardLogger) WithError(err error) Logger {
	return hardLogger{Entry: l.Entry.WithError(err)}
}

func (l hardLogger) Debug(args ...interface{}) { l.Debugln(args...) }
func (l hardLogger) Info(args ...interface{})  { l.Infoln(args...) }
func (l hardLogger) Warn(args ...interface{})  { l.Warnln(args...) }
func (l hardLogger) Error(args ...interface{}) { l.Errorln(args...) }
func (l hardLogger) Fatal(args ...interface{}) { l.Fatalln(args...) }
func (l hardLogger) Print(args ...interface{}) { l.Println(args...) }

func (l hardLogger) Log(level Level, args ...interface{}) { l.Logln(configLevels[level], args...) }

func (l hardLogger) Logf(level Level, format string, args ...interface{}) {
	l.Entry.Logf(configLevels[level], format, args...)
}
