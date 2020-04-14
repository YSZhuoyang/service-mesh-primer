package handlers

import (
	contracts "go-service/contracts/greeter-go/greeter_go"
	"time"
)

// GetLiveData - Push live data to client
func (*Service) GetLiveData(req *contracts.LiveDataRequest, srv contracts.Greeter_GetLiveDataServer) error {
	for i := 0; ; i++ {
		srv.Send(&contracts.LiveDataReply{
			Data: (int32)(i),
		})
		time.Sleep(time.Second)
	}

	return nil
}
