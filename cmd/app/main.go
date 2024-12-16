package main

import (
	"log"
	"slices"
	"time"
	"watchdog-go/data"
	engines "watchdog-go/log_engines"

	"github.com/shirou/gopsutil/v4/process"
)

type App struct {
	settings Settings
	logger   data.Logger
}

func main() {
	settings, err := ParseSettings("settings.toml")
	if err != nil {
		log.Fatalf("error parsing settings.toml file: %v", err)
	}

	logger, err := data.NewLogger(settings.LogFile, &engines.CsvLoggerEngine{})
	if err != nil {
		log.Fatalf("error creating logger instance: %v", err)
	}
	defer logger.Close()

	app := App{
		settings: settings,
		logger:   logger,
	}

	log.Println("Starting watchdog ...")

	caughtProcesses := map[int32]*data.ProcessItem{}

	for {
		// loop over saved processes and log if ended
		for _, savedProc := range caughtProcesses {
			exists, err := process.PidExists(savedProc.Pid)
			if err != nil {
				log.Fatalf("error getting process by PID: %v", err)
			}

			if !exists {
				err = app.logger.Log(savedProc.GetLogItem("end"))
				if err != nil {
					log.Fatalf("error saving log to log file: %v", err)
				}

				delete(caughtProcesses, savedProc.Pid)
			}
		}

		procs, err := process.Processes()
		if err != nil {
			log.Fatalf("error fetching processes: %v", err)
		}

		for _, proc := range procs {
			name, err := proc.Name()
			if err != nil {
				continue
			}

			// only loop processes from settings
			if !slices.Contains(app.settings.Processes, name) {
				continue
			}

			// check if already in caught processes
			if _, ok := caughtProcesses[proc.Pid]; ok {
				continue
			}

			prItem := data.NewProcessItem(proc)

			// add NEW proc to caught processes
			caughtProcesses[proc.Pid] = &prItem

			err = app.logger.Log(prItem.GetLogItem("start"))
			if err != nil {
				log.Fatalf("error saving log to log file: %v", err)
			}
		}

		time.Sleep(time.Duration(app.settings.RefreshSeconds) * time.Second)
	}

}
