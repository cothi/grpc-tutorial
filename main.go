package main

import (
	"log"
	"net"

	"github.com/cothi/proto_test/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1")
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()

	s := chat.ChatServiceServer{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to server: ", err)
	}

}
