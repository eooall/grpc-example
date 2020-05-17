package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	pb "grpc-example/grpc2http/proto"
	"net"
	"net/http"
)

type service struct {
}

func (s *service) Call(ctx context.Context, input *pb.UserInput) (*pb.UserOutput, error) {
	if input.GetName() == "eooall" && input.GetPwd() == "eooall" {
		return &pb.UserOutput{
			Code: 200,
			Msg:  "登录成功",
		}, nil
	}
	return &pb.UserOutput{
		Code: 0,
		Msg:  "登录失败",
	}, nil
}

func listenGRPC() {
	addr, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServiceServer(s, &service{})
	err = s.Serve(addr)
	if err != nil {
		panic(err)
	}
}

//api网关转发方式
func apiGateway() {
	go listenGRPC()
	//网关
	mux := runtime.NewServeMux()
	otps := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterLoginServiceHandlerFromEndpoint(context.TODO(), mux, ":9999", otps)
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":8088", mux)
	if err != nil {
		panic(err)
	}
}
func main() {
	//apiGateway()
	//使用gw自带的注册
	mux := runtime.NewServeMux()
	err := pb.RegisterLoginServiceHandlerServer(context.TODO(), mux, &service{})
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":8989", mux)
	if err != nil {
		panic(err)
	}
}
