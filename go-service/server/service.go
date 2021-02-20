package server

import "go-service/rpc"

type Service struct {
	rpc.UnimplementedGreeterServer
}
