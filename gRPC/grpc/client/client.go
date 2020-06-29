package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	rpc "grpc/rpc"

	"google.golang.org/grpc"
)

func rescueError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

type GRPCClient struct {
	Client rpc.UserClient
}

func (cc GRPCClient) Login(rw http.ResponseWriter, r *http.Request) {
	message := &rpc.Credentials{Username: "congttl", Password: "password"}

	response, _ := cc.Client.UserLogin(context.Background(), message)

	rw.Write([]byte(`{"message": "200"}`))
	log.Println(response)
}

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	defer conn.Close()
	rescueError(err, "Connect to server fail")

	client := GRPCClient{Client: rpc.NewUserClient(conn)}

	router := mux.NewRouter()

	router.HandleFunc("/", client.Login).Methods(http.MethodGet)
	http.ListenAndServe(":8082", router)
}
