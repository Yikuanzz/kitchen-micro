package main

import (
	"log"
	"net"

	handler "github.com/yikuanzz/kitchen/services/orders/handler/orders"
	"github.com/yikuanzz/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	add string
}

func NewGRPCServer(add string) *gRPCServer {
	return &gRPCServer{add: add}
}

func (s *gRPCServer) Run() error {
	listener, err := net.Listen("tcp", s.add)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrderHandler(grpcServer, orderService)

	log.Println("Starting gRPC server on ", s.add)
	return grpcServer.Serve(listener)
}
