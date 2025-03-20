package main

import (
	"Easy-GRPC/hello_grpc/hello_grpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type HelloService struct {
}

func (HelloService) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (res *hello_grpc.HelloResponse, err error) {
	fmt.Println(request)
	return &hello_grpc.HelloResponse{
		Name:    "tangfire",
		Message: "ok",
	}, nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	server := HelloService{}
	// 将server结构体注册为gRPC服务。
	hello_grpc.RegisterHelloServiceServer(s, &server)
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
