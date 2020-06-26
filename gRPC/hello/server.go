package main

import (
	"log"
	"net"

	chat "hello/chat"

	"google.golang.org/grpc"
)

func main() {
	// khởi tạo 1 đối tượng gRPC trong server
	grpcSerrver := grpc.NewServer()

	// đăng ký service với grpcServer (của gRPC plugin)
	chat.RegisterHelloServiceServer(grpcSerrver, &chat.Server{})

	// cung cấp grpc service trên port 1234
	listen, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Println("Cannot listen port 1234")
	}

	grpcSerrver.Serve(listen)
}
