package main

import (
	"context"
	pb "demo/grpc-example/server_stream/proto"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8989", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := pb.NewPushDataClient(conn)
	r, err := c.Call(context.TODO(), &pb.In{
		Token: "eooall",
	})
	if err != nil {
		panic(err)
	}
	for {
		data, err := r.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(data.GetData())
	}
}
