package teeLog

import (
	"io"
	"io/ioutil"
	"os"

	logs "github.com/sirupsen/logrus"
)

var logger *teeLogger

type teeLogger struct {
	jsonLogger *logs.Entry
}

// initialize logger to print text to os.Stderr using the StdOutFormatter and
// the JSON output will just be written to ioutil.Discard until it's set using
// the SetJSONWriter function
func init() {
	logs.SetOutput(os.Stderr)
	logs.SetFormatter(&StdOutFormatter{})

	jsonLogger := &logs.Logger{
		Out:       ioutil.Discard,
		Formatter: &jsonFormatter{},
		Hooks:     make(logs.LevelHooks),
		Level:     logs.InfoLevel,
	}
	logger = &teeLogger{logs.NewEntry(jsonLogger)}
}

// SetJSONWriter initializes a new logger that tees the output to both os.Stderr
// as readable text and also to the provided jsonWtr as a JSON string
func SetJSONWriter(jsonWtr io.Writer) {
	jsonLogger := &logs.Logger{
		Out:       jsonWtr,
		Formatter: &jsonFormatter{},
		Hooks:     make(logs.LevelHooks),
		Level:     logs.InfoLevel,
	}
	logger = &teeLogger{logs.NewEntry(jsonLogger)}
}

// Infof only logs to os.Stderr
func Infof(msg string, args ...interface{}) {
	logs.Infof(msg, args...)
}

// Debugf only logs to os.Stderr
func Debugf(msg string, args ...interface{}) {
	logs.Debugf(msg, args...)
}

// Warnf only logs to os.Stderr
func Warnf(msg string, args ...interface{}) {
	logs.Warnf(msg, args...)
}

// IfErrWarnf only logs to os.Stderr if err != nil
func IfErrWarnf(err error, msg string, args ...interface{}) {
	if err != nil {
		Warnf(msg, args...)
	}
}

// Errorf logs to all
func Errorf(msg string, args ...interface{}) {
	logs.Errorf(msg, args...)
	logger.jsonLogger.Errorf(msg, args...)
}

// IfErrErrorf logs to all if err != nil
func IfErrErrorf(err error, msg string, args ...interface{}) {
	if err != nil {
		Errorf(msg, args...)
	}
}

// Fatalf logs to all
func Fatalf(msg string, args ...interface{}) {
	logs.Fatalf(msg, args...)
	logger.jsonLogger.Fatalf(msg, args...)
	os.Exit(1)
}

// IfErrFatalf logs to all if err != nil
func IfErrFatalf(err error, msg string, args ...interface{}) {
	if err != nil {
		Fatalf(msg, args...)
	}
}
