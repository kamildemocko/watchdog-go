package main

import (
	"log"
	"slices"
	"time"
	"watchdog-go/data"
)

type App struct {
	settings Settings
	logger   data.Logger
}

func main() {
	app := App{}
	app.settings = getSettings()
	app.logger = getLogger(app.settings.LogFile)
	defer app.logger.Close()

	log.Println("Starting watchdog ...")

	runningProcesses := map[int32]*data.ProcessItem{}

	for {
		// loop over saved processes and log if ended
		for _, savedProc := range runningProcesses {
			if exists := processExists(savedProc.Pid); exists {
				continue
			}

			logEventItem(app.logger, *savedProc, "end")

			delete(runningProcesses, savedProc.Pid)
		}

		// loop over refreshed processes and log if started
		procs := getRunningProcesses()

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
			if _, ok := runningProcesses[proc.Pid]; ok {
				continue
			}

			prItem := data.NewProcessItem(proc)

			runningProcesses[proc.Pid] = &prItem

			logEventItem(app.logger, prItem, "start")
		}

		time.Sleep(time.Duration(app.settings.RefreshSeconds) * time.Second)
	}

}
