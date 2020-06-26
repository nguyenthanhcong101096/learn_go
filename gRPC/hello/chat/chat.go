package chat

import (
	context "context"
	"io"
	"log"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, msg *String) (*String, error) {
	log.Println("From Client: ", msg.GetValue())
	serverSay := &String{Value: "Hello Client"}

	return serverSay, nil
}

func (s *Server) Channel(stream HelloService_ChannelServer) error {
	for {
		// Server nhận dữ liệu được gởi từ client
		// trong vòng lặp

		ags, err := stream.Recv()

		if err != nil {
			// Nếu gặp `io.EOF`, client stream sẽ đóng

			if err == io.EOF {
				return nil
			}

			return err
		}

		log.Println("Stream from client: ", ags.Value)

		serverSay := &String{Value: "Hello client"}
		// Dữ liệu trả về được gửi đến client
		// thông qua Stream và việc gửi nhận
		// dữ liệu stream hai chiều là hoàn toàn độc lập

		err = stream.Send(serverSay)

		if err != nil {
			return err
		}
	}
}
