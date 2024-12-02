# Go Environment Loader

This is a simple Go package for loading environment variables from a `.env` file. It provides two main functions:

1. `LoadEnv(filename string) error`: Loads environment variables from a `.env` file.
2. `GetEnv(key string) string`: Retrieves the value of an environment variable.

## Requirements

- Go 1.16 or higher
- A `.env` file containing key-value pairs for environment variables.

## Installation

To use the package in your project, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/yourrepository.git
   ```

2. Install dependencies (if applicable) using Go modules:

   ```bash
   go mod tidy
   ```

## Usage

### Create a `.env` File

Create a `.env` file in the root directory of your project. It should contain environment variable definitions in the form of `KEY=VALUE`. Here's an example:

```ini
# .env file
DB_HOST=localhost
DB_PORT=5432
DB_USER=myuser
DB_PASS=mypassword
```

## Example Program

```go
package main

import (
	"fmt"
	"log"

	"github.com/rudrprasad05/go-env/env"
)

func main() {
	// Load environment variables from the .env file
	err := env.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve environment variables
	dbHost := env.GetEnv("DB_HOST")
	dbPort := env.GetEnv("DB_PORT")
	dbUser := env.GetEnv("DB_USER")
	dbPass := env.GetEnv("DB_PASS")

	// Print out the loaded environment variables
	fmt.Printf("Database Details:\n")
	fmt.Printf("Host: %s\n", dbHost)
	fmt.Printf("Port: %s\n", dbPort)
	fmt.Printf("User: %s\n", dbUser)
	fmt.Printf("Password: %s\n", dbPass)
}


```
