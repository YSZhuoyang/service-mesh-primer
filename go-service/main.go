package main

import (
	"fmt"
	"log"
	"net"

	"go-service/rpc"
	"go-service/server"
	"go-service/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port    int    = 8080
	network string = "tcp"
)

func main() {
	// Start GRPC server ...
	grpcServer := grpc.NewServer()
	rpc.RegisterGreeterServer(grpcServer, &server.Service{})

	connection, err := net.Listen(network, fmt.Sprintf(":%d", port))
	util.CheckError(err)

	log.Printf("Start gRPC server listening at port: %d \n", port)

	reflection.Register(grpcServer)
	err = grpcServer.Serve(connection)
	util.CheckError(err)
}
