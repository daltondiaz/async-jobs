package pkg

import (
	"daltondiaz/async-jobs/db"
	"daltondiaz/async-jobs/models"
	"errors"
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
		// The comment about its to see each execution of job
        slog.Info(fmt.Sprintf("LATEST_JOB %d: ", lastestJob.Id), "job", lastestJob.Executed)
		if lastestJob.Executed != models.EXECUTING {
			timeExec := time.Now().Unix()
			slog.Info(fmt.Sprintf("START_EXEC Job %d: ", lastestJob.Id), "job", lastestJob.Name)
			db.SetJobExecuted(lastestJob.Id, models.EXECUTING)
			cmd := exec.Command(lastestJob.Args.Cmd, lastestJob.Args.Path, lastestJob.Args.Args)
			stdout, err := cmd.Output()
			if err != nil {
				log.Println(err.Error())
				slog.Error("ERROR_CRON", "error", err)
			}
			slog.Info(fmt.Sprintf("END_JOB in %d s: job %d ", time.Now().Unix()-timeExec, job.Id), "output", string(stdout))
			db.SetJobExecuted(lastestJob.Id, models.EXECUTED)
		}
	})
	job.CronId = int(id)
	db.SetCronId(job.Id, int(id))
	slog.Info("CRON", "id", int(id))
}

// Start the crons to scheduler the jobs
func Start() {
	hasTableJob, err := db.CheckExistsJobTable("job")
	if err != nil {
		slog.Error("Error to check if main table job exists", "error", err)
	}
	if !hasTableJob {
		db.CreateTableJobs()
	}
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

// Add new or enabled Job and create a new cron
func AddJobNextExecution(job models.Job) {
	execution(job, cronJob)
}

// Stop job
func StopJob(job models.Job) {
	cronJob.Remove(cron.EntryID(job.CronId))
	slog.Info(fmt.Sprintf("JOB_STOP %d: ", job.Id), "cron id", job.CronId)
}

// Get status Job to know if is running or not
func StatusJob(id int64) (string, error) {
	job, err := db.LoadJob(id)

	if err != nil {
		return "", err
	}

	if job.Executed == models.EXECUTED {
		return "Executed", nil
	}
	if job.Executed == models.EXECUTING {
		return "Executing", nil
	}
	err = errors.New("Job Status not Found")
	return "", err
}

// Enabled the job loading its by Id and if enabled equals true add to a cron job
// in next exection or false stop execution of cron job
func EnabledJob(id int64, enabled bool) (string, error) {
	if enabled == true {
		err := db.SetEnabledJob(id, enabled)
		if err != nil {
			return "Error to set enabled value on Job", err
		}
		job, err := db.LoadJob(id)
		if err != nil {
			return "Error to Load Job", err
		}
		AddJobNextExecution(job)
		return "Job was Enabled with success", nil
	}

	job, err := db.LoadJob(id)
	if err != nil {
		return "Error to Load Job", err
	}
	StopJob(job)
	return "Job was Disabled with success", nil
}
