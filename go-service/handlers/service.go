package handlers

import contracts "go-service/contracts/greeter-go/greeter_go"

type Service struct {
	contracts.UnimplementedGreeterServer
}
