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
