package config

import (
	"github.com/omareloui/gRPCMicroservicesInGo/lib/go/config_helpers"
)

func GetEnv() string {
	return config_helpers.GetEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return config_helpers.GetEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	return config_helpers.GetEnvironmentInt("APPLICATION_PORT")
}

func GetPaymentServiceUrl() string {
	return config_helpers.GetEnvironmentValue("PAYMENT_SERVICE_URL")
}
