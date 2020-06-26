package main

import (
	"context"
	"log"
	"time"

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

	// Client cần gọi phương thức Channel để lấy đối tượng stream trả về:
	stream, err := client.Channel(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	// Ở phía client ta thêm vào các thao tác gửi và nhận trong các Goroutine riêng biệt.
	// Trước hết là để gửi dữ liệu tới server:
	go func() {
		for {
			if err := stream.Send(&chat.String{Value: "Hello server"}); err != nil {
				log.Fatal(err)
			}
		}
	}()

	// Nhận dữ liệu
	go func() {
		for {
			reply, err := stream.Recv()

			if err != nil {
				log.Fatal(err)
			}

			log.Println("Stream server :", reply.Value)
		}
	}()

	time.Sleep(time.Second * 1)
}
