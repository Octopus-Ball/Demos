package main

import (
	"fmt"
	"grpcDemo/proto"
	"net"

	"google.golang.org/grpc"
	// "unsafe"
)

func main() {
	addr := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Println("监听端口")

	// 实例化gRPC
	s := grpc.NewServer()
	// 将服务注册到gRPC
	proto.RegisterUserInfoServiceServer(s, new(proto.UnimplementedUserInfoServiceServer))
	//启动服务端
	s.Serve(listener)
}
