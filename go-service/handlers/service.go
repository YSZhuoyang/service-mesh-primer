package handlers

import "go-service/contracts/greeter"

type Service struct {
	contracts.UnimplementedGreeterServer
}
