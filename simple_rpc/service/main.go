package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-example/simple_rpc/proto"
	"net"
)

type loginServer struct{}

func (s *loginServer) Call(ctx context.Context, in *proto.User) (*proto.Data, error) {
	fmt.Println(in)
	return &proto.Data{
		Code: 200,
		Msg:  "请求成功",
	}, nil

}
func main() {
	s := &loginServer{}
	addr, err := net.Listen("tcp", ":18900")
	if err != nil {
		panic(err)
	}
	g := grpc.NewServer()
	proto.RegisterLoginServer(g, s)
	err = g.Serve(addr)
	if err != nil {
		panic(err)
	}
}
