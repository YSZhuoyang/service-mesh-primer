package server

import (
	"time"

	"go-service/rpc"
)

// GetLiveData - Push live data to client
func (*Service) GetLiveData(req *rpc.LiveDataRequest, srv rpc.Greeter_GetLiveDataServer) error {
	for i := 0; ; i++ {
		srv.Send(&rpc.LiveDataReply{
			Data: (int32)(i),
		})
		time.Sleep(time.Second)
	}

	return nil
}
