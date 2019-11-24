package glog

import (
	"context"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var filePrefix = "file://"

// Logger is an interface of logging operations
type Logger interface {
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Panicf(format string, v ...interface{})

	Info(v ...interface{})
	Debug(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Panic(v ...interface{})

	Infoc(ctx context.Context, format string, v ...interface{})
	Debugc(ctx context.Context, format string, v ...interface{})
	Warnc(ctx context.Context, format string, v ...interface{})
	Errorc(ctx context.Context, format string, v ...interface{})
	Panicc(ctx context.Context, format string, v ...interface{})

	WithField(field string, value interface{}) Logger
	WithPrefix(prefix string) Logger
	Close() error
}

type closeFunc = func() error

//GLog is logger implementation the new implement by @khangpham
type GLog struct {
	logger *logrus.Entry
	writer io.WriteCloser
}

// New return a new logger
// It will lookup configuration from environment variables for initialization
// LOG_FORMAT can be text/json
// LOG_OUTPUT can be a file by setting value to file:///path/to/logfile
// if LOG_OUTPUT is different to a file, it will be redirect to os.Stdout.
func New() *GLog {
	l := &GLog{}

	logger := logrus.New()
	logger.SetFormatter(getFormater())
	logger.SetLevel(getLevel())

	out := getOutput()
	logger.SetOutput(out)

	l.writer = out
	l.logger = logrus.NewEntry(logger)
	return l
}

func getFormater() logrus.Formatter {
	var formatter logrus.Formatter
	formatter = &logrus.TextFormatter{
		TimestampFormat:           time.RFC1123,
		FullTimestamp:             true,
		EnvironmentOverrideColors: true,
	}
	if os.Getenv("LOG_FORMAT") == "json" {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: time.RFC1123,
		}
	}
	return formatter
}

func getLevel() logrus.Level {
	lvl, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		lvl = logrus.DebugLevel
	}
	return lvl
}

func getOutput() io.WriteCloser {
	out := os.Getenv("LOG_OUTPUT")
	if strings.HasPrefix(out, filePrefix) {
		name := out[len(filePrefix):]
		f, err := os.Create(name)
		if err != nil {
			log.Printf("log: failed to create log file: %s, err: %v\n", name, err)
		}
		return f
	}
	return os.Stdout
}

// Infof print info with format.
func (l *GLog) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

// Debugf print debug with format.
func (l *GLog) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

// Warnf print warning with format.
func (l *GLog) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

// Errorf print error with format.
func (l *GLog) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

// Panicf panic with format.
func (l *GLog) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}

// Info print log info leve
func (l *GLog) Info(v ...interface{}) {
	l.logger.Info(v...)
}

// Debug print debug with format.
func (l *GLog) Debug(v ...interface{}) {
	l.logger.Debug(v...)
}

// Warn print warning with format.
func (l *GLog) Warn(v ...interface{}) {
	l.logger.Warn(v...)
}

// Error print error with format.
func (l *GLog) Error(v ...interface{}) {
	l.logger.Error(v...)
}

// Panic panic with format.
func (l *GLog) Panic(v ...interface{}) {
	l.logger.Panic(v...)
}

// Infoc print info log with context
func (l *GLog) Infoc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Infof(format, v...)
}

// Debugc print debug with context
func (l *GLog) Debugc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Debugf(format, v...)
}

// Warnc print warning with context
func (l *GLog) Warnc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Warnf(format, v...)
}

//Errorc print error with context
func (l *GLog) Errorc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Errorf(format, v...)
}

// Panicc panic with context
func (l *GLog) Panicc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Panicf(format, v...)
}

func (l *GLog) withContext(ctx context.Context) Logger {
	if requestID := ctx.Value("request_id"); requestID != nil {
		return l.WithField("request_id", requestID)
	}
	return l
}

// WithField return a new logger with field
func (l *GLog) WithField(field string, value interface{}) Logger {
	nl := l.logger.WithField(field, value)
	return &GLog{
		logger: nl,
	}
}

// Close close the underlying writer
func (l *GLog) Close() error {
	if l.writer != nil {
		return l.writer.Close()
	}
	return nil
}

//WithPrefix return a logger has package name
func (l *GLog) WithPrefix(prefix string) Logger {
	nl := l.logger.WithField("package", prefix)
	return &GLog{
		logger: nl,
	}
}
