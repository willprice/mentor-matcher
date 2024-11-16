package migrations

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbFilePath string, createFile bool) {
	if createFile {
		// Create the database file if it doesn't exist
		if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
			file, err := os.Create(dbFilePath)
			if err != nil {
				log.Fatalf("Failed to create database file: %v", err)
			}
			file.Close()
			log.Printf("Database file created at %s", dbFilePath)
		}
	}

	m, err := migrate.New(
		"file://migrations",     // Path to migrations folder
		"sqlite3://"+dbFilePath, // SQLite database URL
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	// Get the current version before running migrations
	versionBefore, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatalf("Failed to get current migration version: %v", err)
	}
	log.Printf("Current migration version: %d, dirty: %v", versionBefore, dirty)

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Get the current version after running migrations
	versionAfter, dirty, err := m.Version()
	if err != nil {
		log.Fatalf("Failed to get current migration version: %v", err)
	}
	log.Printf("Migrations applied successfully! Current migration version: %d, dirty: %v", versionAfter, dirty)
}
