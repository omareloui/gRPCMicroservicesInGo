package config_helpers

import (
	"log"
	"os"
	"strconv"
)

func GetEnvironmentValue(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return v
}

func GetEnvironmentInt(key string) int {
	portStr := GetEnvironmentValue(key)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("%s: %s is invalid integer", key, portStr)
	}
	return port
}
