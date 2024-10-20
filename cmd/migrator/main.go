package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	start := time.Now()

	fmt.Println("created flags")

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations")
	flag.Parse()

	fmt.Println("parsed flags")

	if storagePath == "" {
		log.Fatal("storage-path is required")
	}

	if migrationsPath == "" {
		log.Fatal("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		log.Fatalf("unable to connect database, ended with error: %s", err)
	}

	fmt.Println("created struct")

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		log.Fatalf("unable to apply migrations, ended with error: %s", err)
	}

	fmt.Println("migrations successfully applied", time.Since(start))
}
