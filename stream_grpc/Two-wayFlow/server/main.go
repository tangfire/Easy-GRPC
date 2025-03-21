package main

import (
	"Easy-GRPC/stream_grpc/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type BothStream struct {
}

func (BothStream) Chat(stream proto.BothStream_ChatServer) error {

	for i := 0; i < 10; i++ {
		request, _ := stream.Recv()
		fmt.Println(request)
		stream.Send(&proto.Response{
			Text: "你好",
		})
	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterBothStreamServer(server, &BothStream{})

	server.Serve(listen)
}
