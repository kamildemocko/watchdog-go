package engines

import (
	"os"
	"watchdog-go/data"

	"github.com/gocarina/gocsv"
)

type CsvLoggerEngine struct{}

func (cl *CsvLoggerEngine) AppendToEmptyFile(file *os.File, logItem *data.LogItem) error {
	str, err := gocsv.MarshalString([]*data.LogItem{logItem})
	if err != nil {
		return err
	}

	_, err = file.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}

func (cl *CsvLoggerEngine) AppendToFile(file *os.File, logItem *data.LogItem) error {
	str, err := gocsv.MarshalStringWithoutHeaders([]*data.LogItem{logItem})
	if err != nil {
		return err
	}

	_, err = file.WriteString(str)
	if err != nil {
		return err
	}

	return nil
}
