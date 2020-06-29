package main

import (
	"log"
	"net"

	entity "grpc/entity"
	rpc "grpc/rpc"

	"google.golang.org/grpc"
)

func rescueError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func main() {
	grpcServer := grpc.NewServer()

	listen, err := net.Listen("tcp", "localhost:1234")

	rescueError(err, "Cannot listen port 1234: ")

	rpc.RegisterUserServer(grpcServer, &entity.Server{})

	grpcServer.Serve(listen)
}
