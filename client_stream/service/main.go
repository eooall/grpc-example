package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-example/client_stream/proto"
	"io"
	"net"
)

type service struct {
}

var i = 0

func (s *service) Call(out pb.Input_CallServer) error {
	for {
		r, err := out.Recv()
		if err == io.EOF {
			return out.SendAndClose(&pb.Out{Code: 1, Msg: "接受成功"})
		}
		if err != nil {
			return err
		}
		i++
		fmt.Println("当前客户端传递: " + r.GetInput())
		fmt.Println(i)
	}
}
func main() {
	addr, err := net.Listen("tcp", ":8989")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterInputServer(s, &service{})
	s.Serve(addr)
}
