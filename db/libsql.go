package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
    database := os.Getenv("TURSO_DATABASE_URL")
    token := os.Getenv("TURSO_AUTH_TOKEN")
    dbName := fmt.Sprintf("%s?authToken=%s", database, token)
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
        os.Exit(1)
	}
	return db
}
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
		log.Fatal("Error to load .env file on root of project")
	}

}

