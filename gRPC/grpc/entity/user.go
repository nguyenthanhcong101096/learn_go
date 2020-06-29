package entity

import (
	"context"
	rpc "grpc/rpc"
	"log"
)

type Server struct{}

func (u *Server) UserLogin(ctx context.Context, in *rpc.Credentials) (*rpc.LoginResult, error) {
	log.Println("Client Login: ", in.Username, " - ", in.Password)

	return &rpc.LoginResult{
		Ok:   true,
		Data: &rpc.AccessToken{AccessToken: "123"},
	}, nil
}

func (u *Server) UserRegister(ctx context.Context, in *rpc.FromRegister) (*rpc.RegisterResult, error) {
	log.Println("Client Register: ", in.Email, "-", in.Passwotd)

	return &rpc.RegisterResult{
		Ok:   true,
		Data: &rpc.AccessToken{AccessToken: "456"},
	}, nil
}
