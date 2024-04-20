package db

import (
	"context"
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
func GetConnection() *sql.DB {
    // TODO add Turso
	dbName := "file:./local.db"
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	ping(ctx, db)
	// CreateTableJobs(db)
	//log.Println(res.LastInsertId())
	// Insert(db)
	return db
}
