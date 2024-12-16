package main

import (
	"fmt"
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

	fmt.Println(settings.LogFile)
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
		procs, err := process.Processes()
		if err != nil {
			log.Fatalf("error fetching processes: %v", err)
		}

		for _, proc := range procs {
			name, err := proc.Name()
			if err != nil {
				continue
			}

			if !slices.Contains(app.settings.Processes, name) {
				continue
			}

			if _, ok := caughtProcesses[proc.Pid]; ok {
				continue
			}

			prItem := data.NewProcessItem(proc)

			caughtProcesses[proc.Pid] = &prItem

			err = app.logger.Log(prItem.GetLogItem("start"))
			if err != nil {
				log.Fatalf("error saving log to log file: %v", err)
			}
		}

		time.Sleep(time.Duration(app.settings.RefreshSeconds) * time.Second)
	}
}
