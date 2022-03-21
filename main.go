package main

import (
	"net"

	"github.com/alirezakargar1380/grpc_mysql/app/server"
	"github.com/alirezakargar1380/grpc_mysql/user_log"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	s := server.Serverr{}
	// user_log.RegisterChatServiceServer(grpcServer, &s)
	user_log.RegisterGuploadServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
