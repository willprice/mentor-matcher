package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/willprice/mentor-matcher/platform/authenticator"
	"github.com/willprice/mentor-matcher/platform/migrations"
	"github.com/willprice/mentor-matcher/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load .env file: %v", err)
	}

	dbFilePath := os.Getenv("SQLITE_FILE_PATH")
	if dbFilePath == "" {
		log.Fatal("SQLITE_FILE_PATH environment variable is not set")
	}

	createFile := true
	if createFileEnv := os.Getenv("SQLITE_CREATE_FILE"); createFileEnv != "" {
		var err error
		createFile, err = strconv.ParseBool(createFileEnv)
		if err != nil {
			log.Fatalf("Invalid value for SQLITE_CREATE_FILE: %v", err)
		}
	}

	migrations.RunMigrations(dbFilePath, createFile)
	log.Println("Migration complete. Starting the application...")

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:8080/")
	if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
