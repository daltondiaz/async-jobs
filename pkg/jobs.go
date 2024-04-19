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
	slog.Info("CRON", "id", id)
}

func Run() {
	db.SetAllJobsToExecute()
	// run jobs according to cron
	jobs, err := db.GetAvailableJobs()
	if err != nil {
		log.Println(err)
	}
	c := cron.New()
	for _, job := range jobs {
		execution(job, c)
	}
	c.Start()
}
