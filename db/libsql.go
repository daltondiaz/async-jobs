package db

import (
	"context"
	"daltondiaz/async-jobs/models"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func exec(ctx context.Context, db *sql.DB, stmt string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, stmt, args)
	if err != nil {
		log.Printf("failed to execute statement %s: %s", stmt, err)
		os.Exit(1)
	}
	return res
}

func ping(ctx context.Context, db *sql.DB) {
	err := db.PingContext(ctx)
	if err != nil {
		log.Printf("failed to ping in database: %s", err)
		os.Exit(1)
	}
}

func Insert(db *sql.DB){
	var job models.Job
	job.Description = "My First Job"
    job.Name = "first_job"
    job.Cron = "@every 5s"
    job.Args = "10"
	job.Enabled = true
	result := InsertJob(db, job)
	log.Println(result.Id)
}

func GetConnection() *sql.DB {
    dbName := "file:./local.db"
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	ping(ctx, db)
	//res := CreateTableJobs(db)
	//log.Println(res.LastInsertId())
   // Insert(db)
    return db
}

func CloseConnection(db *sql.DB){
    defer db.Close()
}
