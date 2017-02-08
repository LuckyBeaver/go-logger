package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"
)

const (
	ERROR = 1
	WARNING = 2
	INFO = 3
	DEBUG = 4
	TRACE = 5
)
var LOG_LEVEL = map[int]string{1: `ERROR`, 2: `WARNING`, 3: `INFO`, 4: `DEBUG`, 5: `TRACE`}

type Logger struct {

	logger   *log.Logger
	writer   *io.Writer
	logLevel int32
	showDateTime int32
}

var logSingleton = New(os.Stdout, DEBUG)

func New(writer io.Writer, logLevel int) *Logger {
	logger := Logger{logLevel: int32(logLevel), writer: &writer, showDateTime: 1}
	logger.logger = log.New(writer, "", log.LstdFlags)
	return &logger
}

func (l *Logger) getLogLevel() int {
	return int(atomic.LoadInt32(&l.logLevel))
}

func (l *Logger) SetlogLevel(loglevel int) {
	atomic.StoreInt32(&l.logLevel, int32(loglevel))
}

func (l *Logger) SetFlags(flag int)  {
	l.logger.SetFlags(flag)
}

func (l *Logger) IsShowDateTime() bool {
	return atomic.LoadInt32(l.showDateTime) > 0
}

func (l *Logger) print(logLevel int, message ...interface{}) {
	if l.getLogLevel() < logLevel {
		return
	}
	l.logger.Print(LOG_LEVEL[logLevel], fmt.Sprintln(message...))
}

func (l *Logger) printf(logLevel int, format string, message interface{}) {
	if l.getLogLevel() < logLevel {
		return
	}
	l.logger.Println(LOG_LEVEL[logLevel], fmt.Sprintf(format, message...))
}


func (l *Logger) Trace(message ...interface{}) {
	l.print(TRACE, message)
}

func (l *Logger) Tracef(format string, message ...interface{}) {
	l.printf(TRACE, format, message)
}

func (l *Logger) Debug(message ...interface{}) {
	l.print(DEBUG, message)
}

func (l *Logger) Debugf(format string, message ...interface{}) {
	l.printf(DEBUG, format, message)
}

func (l *Logger) Info(message ...interface{}) {
	l.print(INFO, message)
}

func (l *Logger) Infof(format string, message ...interface{}) {
	l.printf(INFO, format, message)
}

func (l *Logger) Warning(message ...interface{}) {
	l.print(WARNING, message)
}

func (l *Logger) Warningf(format string, message ...interface{}) {
	l.printf(WARNING, format, message)
}

func (l *Logger) Error(message ...interface{}) {
	l.print(ERROR, message)
}

func (l *Logger) Errorf(format string, message ...interface{}) {
	l.printf(ERROR, format, message)
}



func SetLogger(l *Logger) {
	logSingleton = l
}

func SetLogLevel(loglevel int) {
	logSingleton.SetlogLevel(loglevel)
}

func Trace(message ...interface{}) {
	logSingleton.Trace(message)
}

func Tracef(format string, message ...interface{}) {
	logSingleton.Tracef(format, message)
}

func Debug(message ...interface{}) {
	logSingleton.Debug(message)
}

func Debugf(format string, message ...interface{}) {
	logSingleton.Debugf(format, message)
}

func Info(message ...interface{}) {
	logSingleton.Info(message)
}

func Infof(format string, message ...interface{}) {
	logSingleton.Infof(format, message)
}

func Warning(message ...interface{}) {
	logSingleton.Warning(message)
}

func Warningf(format string, message ...interface{}) {
	logSingleton.Warningf(format, message)
}

func Error(message ...interface{}) {
	logSingleton.Error(message)
}

func Errorf(format string, message ...interface{}) {
	logSingleton.Errorf(format, message)
}