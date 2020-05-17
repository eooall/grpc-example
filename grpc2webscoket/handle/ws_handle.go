package handle

import (
	pb "demo/grpc-example/grpc2webscoket/prpto"
	"time"
)

type Ws struct {
}

func (h *Ws) Call(in *pb.In, s pb.Ws_CallServer) error {
	for {
		s.Send(&pb.Out{
			Message: "hello",
		})
		time.Sleep(time.Second * 2)
	}
	return nil
}
