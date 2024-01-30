package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/omareloui/gRPCMicroservicesInGo/proto/go/payment"
	"github.com/omareloui/gRPCMicroservicesInGo/services/payment/config"
	"github.com/omareloui/gRPCMicroservicesInGo/services/payment/internal/ports"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	payment.UnimplementedPaymentServiceServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	payment.RegisterPaymentServiceServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	log.Printf("starting payment service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port: %d. Error: %s", a.port, err)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
