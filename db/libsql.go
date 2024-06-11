package db

import (
	"context"
	"daltondiaz/async-jobs/logs"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func exec(ctx context.Context, db *sql.DB, stmt string, args ...any) sql.Result {
	res, err := db.ExecContext(ctx, stmt, args)
	if err != nil {
		logs.ErrorLog.Printf("failed to execute statement %s: %s", stmt, err)
		os.Exit(1)
	}
	return res
}

func ping(ctx context.Context, db *sql.DB) {
	err := db.PingContext(ctx)
	if err != nil {
		logs.ErrorLog.Printf("failed to ping in database: %s", err)
		os.Exit(1)
	}
}
func GetConnection() *sql.DB {

	env := os.Getenv("ENV")

	if strings.ToLower("prod") != env {
        return getLocalConnection()
	}

	database := os.Getenv("TURSO_DATABASE_URL")
	token := os.Getenv("TURSO_AUTH_TOKEN")
	dbName := fmt.Sprintf("%s?authToken=%s", database, token)
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		logs.ErrorLog.Fatal(err)
		os.Exit(1)
	}
	return db
}

func getLocalConnection() *sql.DB {
    localDb, found :=  os.LookupEnv("LIBSQL_PATH")
    
    if !found {
        logs.ErrorLog.Fatal("You should define the property LIBSQL_PATH in .env")
    }
    dbName := fmt.Sprintf("file:%s", localDb)
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		logs.ErrorLog.Fatal(err)
	}
	ctx := context.Background()
	ping(ctx, db)
	return db
}
