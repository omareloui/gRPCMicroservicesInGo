package config

import (
	"log"
	"strconv"

	"github.com/omareloui/gRPCMicroservicesInGo/lib/go/config_helpers"
)

func GetEnv() string {
	return config_helpers.GetEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return config_helpers.GetEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := config_helpers.GetEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}
