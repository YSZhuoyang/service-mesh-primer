package server

import (
	"context"

	"go-service/rpc"
)

// SayHello - Scrape Bing news meta data
func (*Service) SayHello(ctx context.Context, req *rpc.HelloRequest) (*rpc.HelloReply, error) {
	return &rpc.HelloReply{Msg: "Reply from go service: " + req.GetMsg()}, nil
}
