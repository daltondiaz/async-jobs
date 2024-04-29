package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tursodatabase/go-libsql"
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

	env := os.Getenv("ENV")

	if strings.ToLower("local") == env {
		return getLocalConnection()
	}

	if strings.ToLower("replica") == env {
		return modeReplica()
	}

	return onlyRemoteConnection()
}

func modeReplica() *sql.DB {
	dbName := "local.db"
	primaryUrl := "libsql://[DATABASE].turso.io"
	authToken := "..."

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
		libsql.WithAuthToken(authToken),
        libsql.WithSyncInterval(time.Minute * 3),
	)
	if err != nil {
		fmt.Println("Error creating connector:", err)
		os.Exit(1)
	}
	defer connector.Close()

	db := sql.OpenDB(connector)
    return db
}


func onlyRemoteConnection() *sql.DB {
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

func getLocalConnection() *sql.DB {
	dbName := "file:./local.db"
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	ping(ctx, db)
	return db
}
