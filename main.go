package main

import (
	"os"
	"os/signal"

	"daltondiaz/async-jobs/conf"
	"daltondiaz/async-jobs/http"
	"daltondiaz/async-jobs/run"
	"daltondiaz/async-jobs/logs"
)

func main() {
    conf.LoadEnv()
	run.Start()
    http.Start()
    listen()
}

// hack to only finish when stop the terminal
// this help a lot to test before of http running
func listen() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	logs.JobLog.Println("Application stopped")
}
