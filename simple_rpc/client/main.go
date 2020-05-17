package main

import (
	"context"
	"demo/grpc-example/simple_rpc/proto"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	grpcClient()
}
func grpcClient() {
	conn, err := grpc.Dial("127.0.0.1:18900", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := proto.NewLoginClient(conn)
	arg := &proto.User{
		Name: "eooall",
		Pwd:  "eoall",
	}
	if result, err := c.Call(context.TODO(), arg); err == nil {
		fmt.Println(result.GetCode())
		fmt.Println(result.GetMsg())
	} else {
		fmt.Println(err)
	}
}
