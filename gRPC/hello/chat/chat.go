package chat

import (
	context "context"
	"log"
)

type Server struct{}

func (s *Server) Hello(ctx context.Context, msg *String) (*String, error) {
	log.Println("From Client: ", msg.GetValue())
	serverSay := &String{Value: "Hello Client"}

	return serverSay, nil
}
