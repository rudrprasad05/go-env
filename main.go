package main

import (
	"fmt"
	"log"

	"github.com/rudrprasad05/go-env/env"
)

func main() {
	err := env.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := env.GetEnv("DB_HOST")
	dbPort := env.GetEnv("DB_PORT")
	dbUser := env.GetEnv("DB_USER")
	dbPass := env.GetEnv("DB_PASS")

	fmt.Printf("Database Details:\n")
	fmt.Printf("Host: %s\n", dbHost)
	fmt.Printf("Port: %s\n", dbPort)
	fmt.Printf("User: %s\n", dbUser)
	fmt.Printf("Password: %s\n", dbPass)
}
