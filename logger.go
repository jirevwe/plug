package plug

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	RegisterModule(Logger{})
}

const ID ModuleID = "core.logger"

type Logger struct {
	logger *logrus.Logger
	entry  *logrus.Entry
}

func (l *Logger) Emit(value any) error {
	l.Info(value)
	return nil
}

func (l *Logger) Validate() error {
	if l.logger == nil {
		return ErrModuleValidation(ID, "logger")
	}

	if l.entry == nil {
		return ErrModuleValidation(ID, "entry")
	}

	return nil
}

func (l *Logger) Load(ctx Context) error {
	log := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
		Level:        logrus.DebugLevel,
		ReportCaller: false,
	}
	l.logger = log
	l.entry = logrus.NewEntry(log)

	return nil
}

func (Logger) ModuleInfo() ModuleInfo {
	return ModuleInfo{
		ID: "core.logger",
		New: func() Module {
			return new(Logger)
		},
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.entry.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.entry.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.entry.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.entry.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) Errorln(args ...interface{}) {
	l.entry.Errorln(args...)
}

func (l *Logger) WithFields(f Fields) *logrus.Entry {
	return l.entry.WithFields(f)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.entry.Printf(format, args...)
}

func (l *Logger) Println(format string, args ...interface{}) {
	l.entry.Printf(format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatal(fmt.Sprintf(format, args...))
}

func (l *Logger) WithError(err error) *logrus.Entry {
	return l.entry.WithError(err)
}

func (l *Logger) WithLogger() *logrus.Logger {
	return l.logger
}

// SetLevel sets the logger level.
// It panics if v is less than DebugLevel or greater than FatalLevel.
func (l *Logger) SetLevel(v Level) {
	lvl, err := v.ToLogrusLevel()
	if err != nil {
		panic(err)
	}

	l.logger.SetLevel(lvl)
}

// SetPrefix sets logger fields
func (l *Logger) SetPrefix(value interface{}) {
	l.entry = l.entry.WithField("source", value)
}

// Level represents a log level.
type Level int32
type Fields = logrus.Fields

const (
	// FatalLevel is used for undesired and unexpected events that
	// the program cannot recover from.
	FatalLevel Level = iota

	// ErrorLevel is used for undesired and unexpected events that
	// the program can recover from.
	ErrorLevel

	// WarnLevel is used for undesired but relatively expected events,
	// which may indicate a problem.
	WarnLevel

	// InfoLevel is used for general informational log messages.
	InfoLevel

	// DebugLevel is the lowest level of logging.
	// Debug logs are intended for debugging and development purposes.
	DebugLevel
)

// String is part of the fmt.Stringer interface.
//
// Used for testing and debugging purposes.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	default:
		return "unknown"
	}
}

func (l Level) ToLogrusLevel() (logrus.Level, error) {
	switch l {
	case DebugLevel:
		return logrus.DebugLevel, nil
	case InfoLevel:
		return logrus.InfoLevel, nil
	case WarnLevel:
		return logrus.WarnLevel, nil
	case ErrorLevel:
		return logrus.ErrorLevel, nil
	case FatalLevel:
		return logrus.FatalLevel, nil
	default:
		return 0, fmt.Errorf("not a valid log Level: %q", l)
	}
}

var (
	_ Loader    = (*Logger)(nil)
	_ Validator = (*Logger)(nil)
	_ Emitter   = (*Logger)(nil)
)
