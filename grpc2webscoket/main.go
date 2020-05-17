package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpc-example/grpc2webscoket/handle"
	proto "grpc-example/grpc2webscoket/prpto"
	"net"

	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"net/http"
)

func listenGRPC() {
	addr, err := net.Listen("tcp", ":8989")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterWsServer(s, &handle.Ws{})
	err = s.Serve(addr)
	if err != nil {
		panic(err)
	}
}

func main() {
	go listenGRPC()
	//使用webscoket代理
	mux := runtime.NewServeMux()
	//忽略安全选项
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisterWsHandlerFromEndpoint(context.TODO(), mux, ":8989", opts)
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":8990", wsproxy.WebsocketProxy(mux))
	if err != nil {
		panic(err)
	}
}
