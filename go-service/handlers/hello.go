package handlers

import (
	"context"

	"go-service/contracts/greeter"
)

// SayHello - Scrape Bing news meta data
func (*Service) SayHello(ctx context.Context, req *contracts.HelloRequest) (*contracts.HelloReply, error) {
	return &contracts.HelloReply{Msg: "Reply to: " + req.GetMsg()}, nil
}
