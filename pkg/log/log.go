package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// nolint: gomnd
func Init(logLevelString string, format string) {
	logLevel, err := logrus.ParseLevel(logLevelString)
	if err != nil {
		logrus.Error(err)
	}

	logrus.SetLevel(logLevel)

	prettyfier := func(f *runtime.Frame) (string, string) {
		programCounter := make([]uintptr, 10)
		runtime.Callers(9, programCounter)
		fun := runtime.FuncForPC(programCounter[0])
		name := fun.Name()

		idx := strings.LastIndex(name, "/")
		l := len(name)
		r := name[idx+1 : l]

		return fmt.Sprintf("%s()", r), ""
	}

	switch format {
	case "text":
		customFormatter := &logrus.TextFormatter{
			CallerPrettyfier: prettyfier,
			FullTimestamp:    true,
			TimestampFormat:  "2006-01-02 15:04:05.000",
		}

		logrus.SetFormatter(customFormatter)
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			CallerPrettyfier: prettyfier,
			TimestampFormat:  "2006-01-02 15:04:05.000",
		})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{
			CallerPrettyfier: prettyfier,
			TimestampFormat:  "2006-01-02 15:04:05.000",
		})
		logrus.Warnf("incorrect log format: %s", format)
	}

	logrus.SetReportCaller(true)
}

// L is an alias for the standard log.
// nolint: gochecknoglobals
var L = logrus.NewEntry(logrus.StandardLogger())

type (
	loggerKey struct{}
	Fields    logrus.Fields
)

// WithLogger returns a new context with the provided log. Use in
// combination with log.WithField(s) for great effect.
func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	e := logger.WithContext(ctx)
	return context.WithValue(ctx, loggerKey{}, e)
}

func AddField(ctx context.Context, key string, value interface{}) context.Context {
	return WithLogger(ctx, getLogger(ctx).WithField(key, value))
}

func WithFields(ctx context.Context, fields Fields) context.Context {
	return WithLogger(ctx, getLogger(ctx).WithFields(logrus.Fields(fields)))
}

// getLogger retrieves the current log from the context. If no log is
// available, the default log is returned.
func getLogger(ctx context.Context) *logrus.Entry {
	logger := ctx.Value(loggerKey{})

	if logger == nil {
		return L.WithContext(ctx)
	}

	entry, ok := logger.(*logrus.Entry)
	if !ok {
		return L.WithContext(ctx)
	}

	return entry
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	getLogger(ctx).Infof(format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	getLogger(ctx).Info(args...)
}

func InfoDuration(ctx context.Context, functionName string, t time.Time) {
	getLogger(ctx).Infof("%v took: %v", functionName, time.Since(t))
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	getLogger(ctx).Debugf(format, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	getLogger(ctx).Debug(args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	getLogger(ctx).Warnf(format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	getLogger(ctx).Warn(args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	getLogger(ctx).Errorf(format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	getLogger(ctx).Error(args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	getLogger(ctx).Fatalf(format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	getLogger(ctx).Fatal(args...)
}
