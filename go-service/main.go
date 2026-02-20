package main

import (
	"fmt"
	"log"
	"net"

	dotnet_service "go-service/rpc/dotnet-service"
	go_service "go-service/rpc/go-service"
	"go-service/server"
	"go-service/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	port    int    = 8080
	network string = "tcp"
)

func main() {
	// Start GRPC server ...
	grpcServer := grpc.NewServer()

	// Initialize dotnet client
	conn, err := grpc.NewClient("mesh-gateway:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	util.CheckError(err)
	dotnetClient := dotnet_service.NewGreeterClient(conn)

	go_service.RegisterGreeterServer(grpcServer, &server.Service{
		DotnetClient: dotnetClient,
	})

	connection, err := net.Listen(network, fmt.Sprintf(":%d", port))
	util.CheckError(err)

	log.Printf("Start gRPC server listening at port: %d \n", port)

	reflection.Register(grpcServer)
	err = grpcServer.Serve(connection)
	util.CheckError(err)
}
