package handlers

import (
	"context"

	contracts "go-service/contracts/greeter-go"
)

// SayHello - Scrape Bing news meta data
func (*Service) SayHello(ctx context.Context, req *contracts.HelloRequest) (*contracts.HelloReply, error) {
	return &contracts.HelloReply{Msg: "Reply from go service: " + req.GetMsg()}, nil
}
