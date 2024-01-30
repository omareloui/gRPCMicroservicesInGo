package main

import (
	"log"

	"github.com/omareloui/gRPCMicroservicesInGo/services/order/config"
	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/adapters/db"
	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/adapters/grpc"
	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/adapters/payment"
	"github.com/omareloui/gRPCMicroservicesInGo/services/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database. Error: %v", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	port := config.GetApplicationPort()
	grpcAdapter := grpc.NewAdapter(application, port)
	log.Printf("Starting gRPC server on port %d...", port)
	grpcAdapter.Run()
}
