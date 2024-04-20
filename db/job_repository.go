package db

import (
	"daltondiaz/async-jobs/models"
	"database/sql"
	"log/slog"
)

func CreateTableJobs(db *sql.DB) sql.Result {
	res, err := db.Exec("DROP TABLE IF EXISTS job; CREATE TABLE IF NOT EXISTS job (id INTEGER PRIMARY KEY AUTOINCREMENT, description varchar(50), name varchar(50), cron varchar (15), enabled boolean default false, executed int default 0, args varchar(150), id_cron INTEGER ) ")
	if err != nil {
		slog.Error("Error to create table Job", "msg", err)
	}
	return res
}

// Insert new Job in database
func InsertJob(job models.Job) (models.Job, error) {
	conn := GetConnection()
	insertJob := "INSERT INTO job(description, name, cron, enabled, executed, args) VALUES (:desc, :name, :cron, :enabled, :exec, :args)"
	result, err := conn.Exec(insertJob, job.Description, job.Name, job.Cron, job.Enabled, job.Executed, job.Args)
    defer conn.Close()
	if err != nil {
		slog.Error("Error to inser new Job", "msg", err)
        return job, err
	}
	id, err := result.LastInsertId()
	job.Id = id
	return job, nil
}

// Get all Jobs enabled that means are available
func GetAvailableJobs() ([]models.Job, error) {
	conn := GetConnection()
	query := "SELECT id, description, name, cron, enabled, executed, args FROM job WHERE enabled = true"
	rows, err := conn.Query(query)
	if err != nil {
		slog.Error("Error to get available jobs", "error", err)
	}
	defer rows.Close()
	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.Id, &job.Description, &job.Name, &job.Cron, &job.Enabled,
			&job.Executed, &job.Args); err != nil {
			slog.Error("Error to scan jobs", "error", err)
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

// Set the execution of Job
func SetJobExecuted(id int64, exec int) {
	conn := GetConnection()
	upd := "UPDATE job SET executed = :exec WHERE id = :id"
	result, err := conn.Exec(upd, exec, id)
    defer conn.Close()
	if err != nil {
		slog.Error("UPDATE", "msg", err)
	}
	_, err = result.LastInsertId()
	if err != nil {
		slog.Error("UPDATE", "msg", err)
	}
    slog.Info("SET_JOB_EXECUTED","value", exec)
}

// Load the job, important to know with the Job is executing or executed
// or to know the state
func LoadJob(id int64) (models.Job, error) {
	conn := GetConnection()
	query := "SELECT id, description, name, cron, enabled, executed, args, id_cron FROM job WHERE enabled = true and id = :id"
	row := conn.QueryRow(query, id)
    defer conn.Close()
	var job models.Job
	if err := row.Scan(&job.Id, &job.Description, &job.Name, &job.Cron, &job.Enabled,
		&job.Executed, &job.Args, &job.CronId); err != nil {
		slog.Error("Error to scan jobs", "error", err)
		return job, err
	}
	return job, nil
}

// Set all jobs to be executed, important to use when the application is started
func SetAllJobsToExecute() {
	conn := GetConnection()
	upd := "UPDATE job SET executed = 0 WHERE enabled = true"
	result, err := conn.Exec(upd)
    defer conn.Close()
	if err != nil {
		slog.Error("UPDATE", "msg", err)
	}
	_, err = result.LastInsertId()
	if err != nil {
		slog.Error("UPDATE", "msg", err)
	}
    slog.Info("SET_ALL_JOBS_EXECUTED","value", exec)
}

// Set the value of Cron id on Job, help us in the future to stop the cron of the on job
func SetCronId(id int64, cronId int) {
	conn := GetConnection()
    upd := "UPDATE job SET id_cron = :cronId WHERE id = :id"
	result, err := conn.Exec(upd, cronId, id)
    defer conn.Close()
	if err != nil {
		slog.Error("UPDATE", "msg", err)
	}
	_, err = result.LastInsertId()
	if err != nil {
		slog.Error("UPDATE", "msg", err)
	}
    slog.Info("SET_JOB_CRON_ID","value", cronId)
}

// Enabled Job  
func SetEnabledJob(id int64, enabled bool) error{
	conn := GetConnection()
    upd := "UPDATE job SET enabled = :enabled WHERE id = :id"
	result, err := conn.Exec(upd, enabled, id)
    defer conn.Close()
	if err != nil {
		slog.Error("UPDATE", "msg", err)
        return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		slog.Error("UPDATE", "msg", err)
        return err
	}
    slog.Info("SET_JOB_ENABLED","value", enabled)
    return nil
}
