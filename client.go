package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/alirezakargar1380/grpc_mysql/user_log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := user_log.NewChatServiceClient(conn)

	var b user_log.Log_Res
	b.Name = "hello"

	message := user_log.Log{
		Res: &b,
	}

	response, err := c.AddMessage(context.Background(), &message)

	if err != nil {
		panic(err)
	}

	log.Printf("response from server: %s", response)
}
