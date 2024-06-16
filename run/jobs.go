package run

import (
	"daltondiaz/async-jobs/db"
	"daltondiaz/async-jobs/logs"
	"daltondiaz/async-jobs/models"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

var cronJob *cron.Cron

// Schedule cron and execution
func execution(job models.Job, c *cron.Cron) {
	logs.JobLog.Println("JOB ", job.Id, "CRON", job.Cron)
	id, err := c.AddFunc(job.Cron, func() {
		lastestJob, _ := db.LoadJob(job.Id)
		// The comment about its to see each execution of job
		// logs.JobLog.Println("LATEST_JOB ", lastestJob.Id, "job", lastestJob.Executed, "execution", lastestJob.Executed)
		if lastestJob.Executed != models.EXECUTING {
			timeExec := time.Now().Unix()
			logs.JobLog.Println("START_EXEC job id:", lastestJob.Id, "name:", lastestJob.Name)
			db.SetJobExecuted(lastestJob.Id, models.EXECUTING)
			args := []string{lastestJob.Args.Path}
			args = append(args, lastestJob.Args.Args...)
			cmd := exec.Command(lastestJob.Args.Cmd, args...)
			stdout, err := cmd.Output()
			if err != nil {
				logs.ErrorLog.Println("ERROR_CRON job id:", lastestJob.Id, "error:", err)
				logs.JobLog.Println("ERROR_CRON job id:", lastestJob.Id, "error:", err)
			}
			logs.JobLog.Println(fmt.Sprintf("END_JOB in %d s: job %d ", time.Now().Unix()-timeExec, job.Id), "output", string(stdout))
			db.SetJobExecuted(lastestJob.Id, models.EXECUTED)
		}
	})

	if err != nil {
		logs.ErrorLog.Println("ERROR_CRON job ", job.Id, "msg", err)
	}

	job.CronId = int(id)
	db.SetCronId(job.Id, int(id))
	logs.JobLog.Println("JOB_CREATED", "id", int(id), "name", job.Name, "description", job.Description, "cron", job.Cron)
}

// Start the application
func Start() {
	hasTableJob, err := db.CheckExistsJobTable("job")
	if err != nil {
		logs.ErrorLog.Println("Error to check if main table job exists", "error", err)
	}
	if !hasTableJob {
		db.CreateTableJobs()
	}
	db.SetAllJobsToExecute()
	jobs, err := db.GetAvailableJobs()
	if err != nil {
		logs.ErrorLog.Println("START_APP", err)
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
	logs.JobLog.Println("JOB_STOPPED", job.Id, "cron id", job.CronId)
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
