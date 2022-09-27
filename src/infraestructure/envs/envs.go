package environmentvariables

import (
	"log"
	"os"
)

func GetEnvironmentVariableByKey(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		log.Fatalf("empty value in environment variable: %s", key)
	}
	return
}

func CheckEnvironmentVariables() {
	// Postgres Database
	GetEnvironmentVariableByKey("DB_HOST")
	GetEnvironmentVariableByKey("DB_PORT")
	GetEnvironmentVariableByKey("DB_USER")
	GetEnvironmentVariableByKey("DB_NAME")
	GetEnvironmentVariableByKey("DB_PASSWORD")
}
