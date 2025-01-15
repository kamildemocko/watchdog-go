package data

import (
	"errors"
	"log"
	"os"
	"path"
)

type ILog interface {
	AppendToFile(file *os.File, logItem *LogItem) error
	AppendToEmptyFile(file *os.File, data *LogItem) error
}

type Logger struct {
	Path      string
	Engine    ILog
	file      *os.File
	emptyFile bool
}

func NewLogger(filepath string, engine ILog) (*Logger, error) {
	root := path.Dir(filepath)
	err := os.MkdirAll(root, 0644)
	if err != nil {
		return &Logger{}, err
	}

	emptyFile := false
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		emptyFile = true
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return &Logger{}, err
	}

	return &Logger{
			filepath,
			engine,
			file,
			emptyFile,
		},
		nil
}

func (li *Logger) Close() {
	li.file.Close()
}

func (li *Logger) Log(logItem LogItem) error {
	log.Println("logging PID ", logItem.Pid)

	var err error
	if li.emptyFile {
		err = li.Engine.AppendToEmptyFile(li.file, &logItem)
		li.emptyFile = false
	} else {
		err = li.Engine.AppendToFile(li.file, &logItem)
	}
	if err != nil {
		return err
	}

	return nil
}
