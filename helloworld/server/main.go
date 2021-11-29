package main

import (
	"fmt"
	pb "go_gRPC/helloworld/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

const (
	//gRPC服务地址
	Address = "127.0.0.1:5050"
)

//定义一个helloServer并实现约定的接口
type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "hello " + in.Name + "."
	return resp, nil
}

func main() {
	// 监听本地端口
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}

	//实现gRPC Server
	s := grpc.NewServer()

	//注册helloServer为客户端提供服务
	pb.RegisterHelloServer(s, &helloService{}) //内部调用了s.RegisterServer()
	fmt.Println("Listen on" + Address)

	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("Start service failed:%v", err)
		return
	}
}
