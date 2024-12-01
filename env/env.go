package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open .env file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		err := os.Setenv(parts[0], parts[1])
		if err != nil {
			return fmt.Errorf("could not set environment variable: %v", err)
		}
	}
	return nil
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Println("env not found key: ", key)
	}
	return value
}
