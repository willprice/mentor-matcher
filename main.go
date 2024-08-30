package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/willprice/mentor-matcher/platform/authenticator"
	"github.com/willprice/mentor-matcher/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load .env file: %v", err)
	}

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
