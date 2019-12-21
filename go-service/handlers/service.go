package handlers

import contracts "go-service/contracts/greeter-go"

type Service struct {
	contracts.UnimplementedGreeterServer
}
