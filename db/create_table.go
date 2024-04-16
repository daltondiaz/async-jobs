package db

import (
	"daltondiaz/async-jobs/models"
	"database/sql"
	"log/slog"
)

func CreateTableJobs(db *sql.DB) sql.Result {
	res, err := db.Exec("DROP TABLE IF EXISTS job; CREATE TABLE IF NOT EXISTS job (id INTEGER PRIMARY KEY AUTOINCREMENT, description varchar(50), name varchar(50), cron varchar (15), enabled boolean default false, executed int default 0, args varchar(150) ) ")
	if err != nil {
		slog.Error("Error to create table Job", "msg", err)
	}
	return res
}

func InsertJob(db *sql.DB, job models.Job) models.Job {
	insertJob := "INSERT INTO job(description, name, cron, enabled, executed, args) VALUES (:desc, :name, :cron, :enabled, :exec, :args)"
	result, err := db.Exec(insertJob, job.Description, job.Name, job.Cron, job.Enabled, job.Executed, job.Args)
	if err != nil {
		slog.Error("Error to inser new Job", "msg", err)
	}
	id, err := result.LastInsertId()
	job.Id = id
	return job
}

func GetAvailableJobs(db *sql.DB) ([]models.Job, error) {

	query := "SELECT id, description, name, cron, enabled, executed, args FROM job WHERE enabled = true"
	rows, err := db.Query(query)
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
