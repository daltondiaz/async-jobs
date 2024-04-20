package pkg

import (
	"daltondiaz/async-jobs/db"
	"daltondiaz/async-jobs/models"
	"fmt"
	"log"
	"log/slog"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

var cronJob *cron.Cron

func execution(job models.Job, c *cron.Cron) {
	slog.Info(fmt.Sprintf("JOB %d: ", job.Id), "job", job)
	id, _ := c.AddFunc(job.Cron, func() {
		lastestJob, _ := db.LoadJob(job.Id)
		slog.Info(fmt.Sprintf("LATEST_JOB %d: ", lastestJob.Id), "job", lastestJob.Executed)
		if lastestJob.Executed != models.EXECUTING {
            idExecution := time.Now().Unix()
			slog.Info(fmt.Sprintf("START_EXEC %d: ", idExecution), "job", lastestJob.Name)
			db.SetJobExecuted(lastestJob.Id, models.EXECUTING)
            // TODO change by args
			path := "/home/dalton/Dev/personal/async-jobs/test.php"
			cmd := exec.Command("php", path, lastestJob.Args)
			stdout, err := cmd.Output()
			if err != nil {
				log.Println(err.Error())
				slog.Error("ERROR_CRON", "error", err)
			}
            slog.Info(fmt.Sprintf("END_JOB %d: job %d ", idExecution, idExecution), "output", string(stdout))
			db.SetJobExecuted(lastestJob.Id, models.EXECUTED)
		}
	})
    job.CronId = int(id) 
    db.SetCronId(job.Id, int(id))
	slog.Info("CRON", "id", int(id))
}

// Start the crons to scheduler the jobs
func Start() {
	db.SetAllJobsToExecute()
	jobs, err := db.GetAvailableJobs()
	if err != nil {
		log.Println(err)
	}
	cronJob = cron.New()
	for _, job := range jobs {
		execution(job, cronJob)
	}
	cronJob.Start()
}

// Add new Job and create a new cron
func AddNewJob(job models.Job){
    job = db.InsertJob(job)
    execution(job, cronJob)
}

// Stop job
func StopJob(job models.Job){
}
