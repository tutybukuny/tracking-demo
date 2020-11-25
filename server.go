package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"tracking-demo/models"
)

func main() {
	log.Print("hello")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := models.NewServer()
	grpcServer := grpc.NewServer()

	models.RegisterTrackingServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}