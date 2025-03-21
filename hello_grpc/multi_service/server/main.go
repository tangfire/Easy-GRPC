package main

import (
	"Easy-GRPC/hello_grpc/multi_service/multi_service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type VideoServer struct {
}

func (VideoServer) Look(ctx context.Context, request *multi_service.Request) (res *multi_service.Response, err error) {
	fmt.Println("video:", request)

	return &multi_service.Response{
		Name: "tangfire",
	}, nil
}

type OrderServer struct {
}

func (OrderServer) Buy(ctx context.Context, request *multi_service.Request) (res *multi_service.Response, err error) {
	fmt.Println("order:", request)

	return &multi_service.Response{
		Name: "tangfire",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	multi_service.RegisterVideoServiceServer(s, &VideoServer{})
	multi_service.RegisterOrderServiceServer(s, &OrderServer{})

	fmt.Println("server listen on :8080")
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}

}
