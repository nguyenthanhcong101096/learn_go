## gRPC streaming
> RPC là lời gọi hàm từ xa, vì vậy các tham số hàm và giá trị trả về của mỗi cuộc gọi không thể quá lớn, nếu không thời gian phản hồi của mỗi lời gọi sẽ bị ảnh hưởng nghiêm trọng. Do đó, các lời gọi phương thức RPC truyền thống không phù hợp để tải lên và tải xuống trong trường hợp khối lượng dữ liệu lớn. Để khắc phục điểm này, framework gRPC cung cấp các chức năng stream cho phía server và client tương ứng.

Ta viết thêm phương thức channel hỗ trợ luồng hai chiều (Bidirect Streaming) trong HelloService:

```
File: project/hello/chat.proto

service HelloService {
    rpc Hello (String) returns (String);

    // nhận vào tham số một stream và trả về giá trị là một stream.
    rpc Channel (stream String) returns (stream String);
}
```

Tạo lại code để thấy định nghĩa mới được thêm vào phương thức kiểu channel trong interface:

```
type HelloServiceServer interface {
    Hello(context.Context, *String) (*String, error)

    // tham số kiểu HelloService_ChannelServer được sử dụng
    // để liên lạc hai chiều với client.
    Channel(HelloService_ChannelServer) error
}

type HelloServiceClient interface {
    Hello(ctx context.Context, in *String, opts ...grpc.CallOption) (
        *String, error,
    )

    // trả về giá trị trả về thuộc kiểu `HelloService_ChannelClient`
    // có thể được sử dụng để liên lạc hai chiều với server.
    Channel(ctx context.Context, opts ...grpc.CallOption) (
        HelloService_ChannelClient, error,
    )
}
```

HelloService_ChannelServer và HelloService_ChannelClient thuộc interface:

```
type HelloService_ChannelServer interface {
    Send(*String) error
    Recv() (*String, error)
    grpc.ServerStream
}

type HelloService_ChannelClient interface {
    Send(*String) error
    Recv() (*String, error)
    grpc.ClientStream
}
```

Có thể thấy các interface hỗ trợ server và client stream đều có định nghĩa phương thức Send và Recv cho giao tiếp hai chiều của dữ liệu streaming

Bây giờ ta có thể xây dựng các streaming service:

```
File: project/hello/chat/chat.go

func (p *Server) Channel(stream HelloService_ChannelServer) error {
    for {
        // Server nhận dữ liệu được gửi từ client
        // trong vòng lặp.
        args, err := stream.Recv()
        if err != nil {
            // Nếu gặp `io.EOF`, client stream sẽ đóng.
            if err == io.EOF {
                return nil
            }
            return err
        }

        reply := &String{Value: "hello:" + args.GetValue()}

        // Dữ liệu trả về được  gửi đến client
        // thông qua stream và việc gửi nhận
        // dữ liệu stream hai chiều là hoàn toàn độc lập
        err = stream.Send(reply)
        if err != nil {
            return err
        }
    }
}
```

Client cần gọi phương thức Channel để lấy đối tượng stream trả về:

```
File: project/hello/client.go

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

```
