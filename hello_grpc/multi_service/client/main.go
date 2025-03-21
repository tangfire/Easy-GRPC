package main

import (
	"Easy-GRPC/hello_grpc/multi_service/multi_service"
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
	orderClient := multi_service.NewOrderServiceClient(conn)
	res, err := orderClient.Buy(context.Background(), &multi_service.Request{
		Name: "crazyFire",
	})

	fmt.Println(res, err)

	videoClient := multi_service.NewVideoServiceClient(conn)

	res, err = videoClient.Look(context.Background(), &multi_service.Request{
		Name: "fireShine",
	})

	fmt.Println(res, err)

}
