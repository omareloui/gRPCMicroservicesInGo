package main

import (
	"log"

	"github.com/omareloui/gRPCMicroservicesInGo/services/payment/config"
	"github.com/omareloui/gRPCMicroservicesInGo/services/payment/internal/adapters/db"
	"github.com/omareloui/gRPCMicroservicesInGo/services/payment/internal/adapters/grpc"
	"github.com/omareloui/gRPCMicroservicesInGo/services/payment/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
