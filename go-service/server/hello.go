package server

import (
	"context"

	dotnet_service "go-service/rpc/dotnet-service"
	go_service "go-service/rpc/go-service"
)

// SayHello - Scrape Bing news meta data
func (s *Service) SayHello(ctx context.Context, req *go_service.HelloRequest) (*go_service.HelloReply, error) {
	resp, err := s.DotnetClient.SayHello(ctx, &dotnet_service.HelloRequest{
		Msg: req.GetMsg(),
	})
	if err != nil {
		return nil, err
	}

	return &go_service.HelloReply{Msg: "Reply from go service calling dotnet service:  " + resp.GetMsg()}, nil
}
