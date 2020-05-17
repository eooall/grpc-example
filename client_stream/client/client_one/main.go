package main

import (
	"context"
	pb "demo/grpc-example/client_stream/proto"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8989", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100000; i++ {
		input, err := pb.NewInputClient(conn).Call(context.TODO())
		if err != nil {
			panic(err)
		}
		err = input.Send(&pb.In{
			Input: "这是A服务器 :" + time.Now().Format("2006-01-02 15:04:05"),
		})
		if err != nil {
			fmt.Println("客户端流发送失败")
			panic(err)
		}
		//time.Sleep(time.Second * 1)
		out, err := input.CloseAndRecv()
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
	}
}
