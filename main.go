package main

import (
	"log"
	"os"
	"os/signal"

	//"daltondiaz/async-jobs/http"
	"daltondiaz/async-jobs/models"
	"daltondiaz/async-jobs/pkg"
)

func main() {
	pkg.Start()
    // pkg.AddNewJob(insert())
    //http.Start()
    listen()
}

func listen() {
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, os.Kill)
    <-sig
    log.Println("Finished job")
}

func insert() models.Job {
	var job models.Job
	job.Description = "Added job"
	job.Name = "add_job"
	job.Cron = "@every 2s"
	job.Args = "9"
	job.Enabled = true
    return job
}
