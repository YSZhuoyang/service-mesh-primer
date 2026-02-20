package server

import (
	"time"

	go_service "go-service/rpc/go-service"
)

// GetLiveData - Push live data to client
func (*Service) GetLiveData(req *go_service.LiveDataRequest, srv go_service.Greeter_GetLiveDataServer) error {
	for i := 0; ; i++ {
		srv.Send(&go_service.LiveDataReply{
			Data: (int32)(i),
		})
		time.Sleep(time.Second)
	}
}
