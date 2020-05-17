package main

import (
	"errors"
	"google.golang.org/grpc"
	pb "grpc-example/server_stream/proto"

	"net"
	"time"
)

type service struct{}

func (s *service) Call(in *pb.In, out pb.PushData_CallServer) error {
	if in.Token != "eooall" {
		return errors.New("token不合法")
	}
	for {
		o := &pb.Out{
			Data: time.Now().Format("2006-01-02 15:04:05"),
		}
		err := out.Send(o)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1) //暂停一秒发送消息
	}
	return nil
}
func main() {
	addr, err := net.Listen("tcp", ":8989")
	if err != nil {
		panic(err)
	}
	g := grpc.NewServer()
	pb.RegisterPushDataServer(g, &service{})
	err = g.Serve(addr)
	if err != nil {
		panic(err)
	}
}
