package main

import (
	pb "go_gRPC/helloworld/protocol"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:5050"
)

func main() {
	//连接gRPC服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Connect server failed: %s", err)
		return
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	//调用方法
	r, err := c.SayHello(context.Background(), 	&pb.HelloRequest{Name: "gRPC(hello world)"})
	if err != nil {
		fmt.Printf("Invoke server failed: %s", err)
	}
	fmt.Printf("Invoke successful: %+v", r.Message)
}