package main

import (
	"Easy-GRPC/stream_grpc/proto"
	"bufio"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

type ClientStream struct {
}

func (ClientStream) UploadFile(stream proto.ClientStream_UploadFileServer) error {

	file, err := os.OpenFile("static/x.txt", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.Flush()
	var index int
	for {
		index++
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}

		writer.Write(response.Content)
		fmt.Printf("第%d次", index)

	}

	for i := 0; i < 10; i++ {
		response, err := stream.Recv()
		fmt.Println(response, err)
	}
	stream.SendAndClose(&proto.Response{Text: "完毕了"})
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	proto.RegisterClientStreamServer(server, &ClientStream{})

	server.Serve(listen)
}
