package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	pb "github.com/Henrod/message-board/protogen"
	"google.golang.org/grpc"
)

func startHTTP() error {
	server := runtime.NewServeMux()

	err := pb.RegisterBoardAPIHandlerFromEndpoint(
		context.Background(),
		server,
		"localhost:8888",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		return err
	}

	log.Print("starting http at localhost:8889")
	return http.ListenAndServe("localhost:8889", server)
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterBoardAPIServer(server, NewBoardAPI())

	go func() {
		err = startHTTP()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Print("starting grpc at localhost:8888")
	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
