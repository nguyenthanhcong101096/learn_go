## gRPC là gì ?
- **gRPC** là một RPC platform được phát triển bởi Google nhằm tối ưu hoá và tăng tốc việc giao tiếp giữa các service với nhau trong kiến trúc microservice.

- **gRPC** dùng Protocal Buffer giảm kích thước request và response data, RPC để đơn giản hoá trong việc tạo ra các giao tiếp giữa các service với nhau, HTTP/2 để tăng tốc gửi/nhận HTTP request.

### RPC
- **RPC** là từ viết tắc của `Remote Procedure Call`, nó được xây dựng với ý tưởng là đơn giản hoá việc giao tiếp giữa những service với nhau, thay vì những service giao tiếp với nhau theo kiểu `RESTful API` thì giờ đơn giản là gọi hàm như những object nói chuyện với nhau thôi, còn việc phân tán các service là chuyện của tương lai không dính liếu đến việc code.

### Protocal Buffer
- **Protocal Buffer** là một ngôn ngữ trung lập để `serializing structured data` sử dụng cho việc giao tiếp giữa các service với nhau. 
- **Protocal Buffer** được tạo ra với ý tưởng là làm nhỏ kích thước data truyền đi trong giao tiếp và chỉ cần định nghĩa một lần và sử dụng cho các service với các ngôn ngữ lập trình khác nhau.

example

```
syntax = "proto2";

package tutorial;

message Person {
  required string name = 1;
  required int32 id = 2;
  optional string email = 3;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    required string number = 1;
    optional PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;
}

message AddressBook {
  repeated Person people = 1;
}
```

### HTTP/2
- **HTTP/2** là một phiên bản nâng cấp của `HTTP/1.1, HTTP/2` sinh với với mục đích cải thiện tốc độ giao tiếp giữa `client/server` trên nền tảng Web.

#### Một vài điểm hay của HTTP/2:
- **Request multiplexing:** HTTP/2 có thể gửi cùng lúc nhiều request đến 1 TCP connection và kết quả được trả về bất đồng bộ với nhau.

![](https://github.com/thenguyenit/blogs/raw/master/images/http2-multiplexing.png)

- **Header compression:** như bạn biết mỗi request của HTTP sẽ mang rất nhiều data header đi và đến cho dù nó giống nhau từ request thứ 2 trở đi, HTTP/2 tối ưu chổ này tí xíu, HTTP/2 sẽ loại bỏ những data header dư thừa ở những lần request thứ 2 trở đi và nén chúng lại trước khi gửi đi.

![](https://github.com/thenguyenit/blogs/raw/master/images/http2-header-compressing.png)

- **Binary protocol:** Browser sẽ convert text sang binary trước khi gửi qua đường network.

![](https://github.com/thenguyenit/blogs/raw/master/images/http2-binary.png)

- **HTTP/2 Server Push:** Thêm một cách để tối ưu tốc độ loading của website, thay vì phải có request từ client thì server mới trả resource về, HTTP/2 sẽ đẩy resource về cho client luôn mà không cần client gửi request.

![](https://github.com/thenguyenit/blogs/raw/master/images/http2-push.png)

## gRPC
- Vậy những ưu điểm rất lớn của **RPC, Protocal Buffer, HTTP/2** sẽ gói trong **gRPC**, giờ là một ví dự đơn giản

- Để compile file Proto, chúng ta cần phải cài một số tool và package cần thiết sau.
	- Tải `proto compiler binary` ở đây: [Protocolbuffers/Protobuf](https://github.com/protocolbuffers/protobuf/releases)
	- Giải nén đến thư mục bất kì, và thêm đường dẫn PATH đến zshrc 
	```
	export PATH="/Users/congnt/Downloads/protoc/bin:$PATH"
	và run: export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
	```
- Cài đặt code generator plugin của golang cho Proto compiler.
	- `go get -u github.com/golang/protobuf/protoc-gen-go `
- Cú pháp -> [Here](https://developers.google.com/protocol-buffers/docs/proto3)

## Làm quen với gRPC
- Từ quan điểm của `Protobuf, gRPC` không gì khác hơn là một trình tạo code cho `interface service`.

Tạo file hello.proto và định nghĩa interface HelloService:

```
File: /project/chat/chat.proto

syntax = "proto3";

package main;

message String {
  string value = 1;
}

service HelloService {
  rpc Hello (String) returns (String) {}
}
```

Tạo gRPC code sử dụng hàm dựng sẵn trong gRPC plugin từ `protoc-gen-go`

```
protoc --go_out=plugins=grpc:<folder> <file>.proto
protoc --go_out=plugins=grpc:chat hello.proto
```

gRPC plugin tạo ra các interface khác nhau cho server và client trong `folder/file.pb.go`


```
File: /project/chat/chat.pb.go

type HelloServiceServer interface {
    Hello(context.Context, *String) (*String, error)
}

type HelloServiceClient interface {
    Hello(context.Context, *String, ...grpc.CallOption) (*String, error)
}
```

> gRPC cung cấp hỗ trợ context cho mỗi lệnh gọi phương thức thông qua tham số context.Context. Khi client gọi phương thức, nó có thể cung cấp thông tin context bổ sung thông qua các tham số tùy chọn của kiểu grpc.CallOption.

Chúng ta sử dụng struct Server để thực hiện service HelloService dựa trên interface HelloSercieServer:


```
File: /project/chat/chat.go

type Server struct{}

func (s *Server) Hello(ctx context.Context, msg *String) (*String, error) {
	log.Println("From Client: ", msg.GetValue())
	serverSay := &String{Value: "Hello Client"}

	return serverSay, nil
}
```

Quá trình khởi động của gRPC service tương tự như quá trình khởi động RPC service của thư viện chuẩn:

```
File: /project/chat/server.go


func main() {
    // khởi tạo một đối tượng gRPC service
    grpcServer := grpc.NewServer()

    // đăng ký service với grpcServer (của gRPC plugin)
    RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

    // cung cấp gRPC service trên port `1234`
    lis, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal(err)
    }
    grpcServer.Serve(lis)
}
```

Tiếp theo bạn đã có thể kết nối tới gRPC service từ client:

```
File: /project/chat/client.go

func main() {
    // thiết lập kết nối với gRPC service
    conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // xây dựng đối tượng `HelloServiceClient` dựa trên kết nối đã thiết lập
    client := NewHelloServiceClient(conn)
    reply, err := client.Hello(context.Background(), &String{Value: "hello"})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(reply.GetValue())
}
```

> Có một sự khác biệt giữa gRPC và framework RPC của thư viện chuẩn: gRPC không hỗ trợ gọi asynchronous. Tuy nhiên, ta có thể chia sẻ kết nối HTTP/2 trên nhiều Goroutines, vì vậy có thể mô phỏng các lời gọi bất đồng bộ bằng cách block các lời gọi trong Goroutine khác.
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
