package main

import (
	"Easy-GRPC/stream_grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
	client := proto.NewBothStreamClient(conn)
	//stream, err := client.Fun(context.Background(), &proto.Request{
	//	Name: "tangfire_client",
	//})

	//for i := 0; i < 10; i++ {
	//	response, err := stream.Recv()
	//	fmt.Println(response, err)
	//}

	stream, err := client.Chat(context.Background())

	for i := 0; i < 10; i++ {
		stream.Send(&proto.Request{
			Name: fmt.Sprintf("第%d次", i),
		})
		response, err := stream.Recv()
		fmt.Println(response, err)
	}
}
