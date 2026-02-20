package server

import (
	dotnet_service "go-service/rpc/dotnet-service"
	go_service "go-service/rpc/go-service"
)

type Service struct {
	go_service.UnimplementedGreeterServer
	DotnetClient dotnet_service.GreeterClient
}
