package main

import (
	"log"
	"watchdog-go/data"
	engines "watchdog-go/log_engines"

	"github.com/shirou/gopsutil/v4/process"
)

func getSettings() Settings {
	settings, err := ParseSettings("settings.toml")
	if err != nil {
		log.Fatalf("error parsing settings.toml file: %v", err)
	}

	return settings
}

func getLogger(logFilepath string) *data.Logger {
	logger, err := data.NewLogger(logFilepath, &engines.CsvLoggerEngine{})
	if err != nil {
		log.Fatalf("error creating logger instance: %v", err)
	}

	return logger
}

func getRunningProcesses() []*process.Process {
	procs, err := process.Processes()
	if err != nil {
		log.Fatalf("error fetching processes: %v", err)
	}

	return procs
}

func processExists(pid int32) bool {
	exists, err := process.PidExists(pid)
	if err != nil {
		log.Fatalf("error getting process by PID: %v", err)
	}

	return exists
}

func logEventItem(logger *data.Logger, procItem data.ProcessItem, event string) {
	err := logger.Log(procItem.GetLogItem(event))
	if err != nil {
		log.Fatalf("error saving log to log file: %v", err)
	}
}
