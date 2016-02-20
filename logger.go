package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

const (
	ERROR = 1
	WARNING = 2
	INFO = 3
	DEBUG = 4
)

type Logger struct {
	logger   *log.Logger
	writer   *io.Writer
	logLevel int32
}

var logSingleton = New(os.Stdout, DEBUG)

func New(writer io.Writer, logLevel int) *Logger {
	logger := Logger{logLevel: int32(logLevel), writer: &writer}
	logger.logger = log.New(writer, "", log.LstdFlags)
	return &logger
}

func (l *Logger) getLogLevel() int {
	return int(atomic.LoadInt32(&l.logLevel))
}

func (l *Logger) SetlogLevel(loglevel int) {
	atomic.StoreInt32(&l.logLevel, int32(loglevel))
}

func (l *Logger) Debug(message ...interface{}) {
	if l.getLogLevel() < DEBUG {
		return
	}
	l.logger.Print("[DEBUG] ", fmt.Sprintln(message...))
}

func (l *Logger) Info(message ...interface{}) {
	if l.getLogLevel() < INFO {
		return
	}
	l.logger.Print("[INFO] ", fmt.Sprintln(message...))
}

func (l *Logger) Infof(format string, message ...interface{}) {
	if l.getLogLevel() < INFO {
		return
	}
	l.logger.Println("[INFO] ", fmt.Sprintf(format, message...))
}

func (l *Logger) Warning(message ...interface{}) {
	if l.getLogLevel() < WARNING {
		return
	}
	l.logger.Print("[WARNING] ", fmt.Sprintln(message...))
}

func (l *Logger) Error(message ...interface{}) {
	if l.getLogLevel() < ERROR {
		return
	}
	l.logger.Print("[ERROR] ", fmt.Sprintln(message...))
}

func (l *Logger) Errorf(format string, message ...interface{}) {
	if l.getLogLevel() < ERROR {
		return
	}
	l.logger.Println("[ERROR] ", fmt.Sprintf(format, message...))
}



func SetLogger(l *Logger) {
	logSingleton = l
}

func SetLogLevel(loglevel int) {
	logSingleton.SetlogLevel(loglevel)
}

func Debug(message ...interface{}) {
	if logSingleton.logLevel < DEBUG {
		return
	}
	logSingleton.logger.Print("[DEBUG] ", fmt.Sprintln(message...))
}

func Debugf(format string, message ...interface{}) {
	if logSingleton.logLevel < DEBUG {
		return
	}
	logSingleton.logger.Print("[DEBUG] ", fmt.Sprintf(format, message...))
}

func Info(message ...interface{}) {
	if logSingleton.logLevel < INFO {
		return
	}
	logSingleton.logger.Print("[INFO] ", fmt.Sprintln(message...))
}

func Infof(format string, message ...interface{}) {
	if logSingleton.logLevel < INFO {
		return
	}
	logSingleton.logger.Println("[INFO] ", fmt.Sprintf(format, message...))
}

func Warning(message ...interface{}) {
	if logSingleton.logLevel < WARNING {
		return
	}
	logSingleton.logger.Print("[WARNING] ", fmt.Sprintln(message...))
}

func Warningf(format string, message ...interface{}) {
	if logSingleton.logLevel < WARNING {
		return
	}
	logSingleton.logger.Print("[WARNING] ", fmt.Sprintf(format, message...))
}

func Error(message ...interface{}) {
	if logSingleton.logLevel < ERROR {
		return
	}
	logSingleton.logger.Print("[ERROR] ", fmt.Sprintln(message...))
}

func Errorf(format string, message ...interface{}) {
	if logSingleton.logLevel < ERROR {
		return
	}
	logSingleton.logger.Println("[ERROR] ", fmt.Sprintf(format, message...))
}
