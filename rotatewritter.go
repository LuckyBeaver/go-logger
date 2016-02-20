package logger

import (
	"os"
	"sync"
	"io"
)

type RotateFile interface {
	io.Writer
	Rotate() error
}

type LogWriter struct {
	sync.Mutex
	filePath string
	file     *os.File
}

func NewLogWriter(filePath string) (*LogWriter, error) {
	logWriter := &LogWriter{filePath: filePath}
	err := logWriter.Rotate()
	return logWriter, err
}

func (w *LogWriter) Rotate() error {
	w.Lock()
	defer w.Unlock()
	if w.file != nil {
		w.file.Close()
	}
	var err error
	w.file, err = os.OpenFile(w.filePath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
	return err
}

func (w *LogWriter) Write(data []byte) (n int, err error) {
	w.Lock()
	defer w.Unlock()
	return w.file.Write(data)
}