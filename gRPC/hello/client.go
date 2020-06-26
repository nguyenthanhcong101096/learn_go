package main

import (
	"context"
	"log"

	chat "hello/chat"

	"google.golang.org/grpc"
)

func main() {
	// Thiết lập kết nối với gRPC service
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())

	defer conn.Close()

	if err != nil {
		log.Println("Connect to server fail")
	}

	// Xây dựng đối tượng HelloServiceClient dựa trên kết nối đã thiết lập

	client := chat.NewHelloServiceClient(conn)
	message := chat.String{Value: "Hello Service"}
	respone, err := client.Hello(context.Background(), &message)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("From Server: ", respone.Value)
}
