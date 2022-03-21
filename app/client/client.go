package main

import (
	"fmt"
	"io"
	"os"

	"github.com/alirezakargar1380/grpc_mysql/user_log"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type ClientGRPC struct {
	chunkSize int
}

func main() {
	var (
		writing bool
		buf     []byte
		// c       *ClientGRPC
	)
	writing = true

	fmt.Println("running message sending")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	file, err := os.Open(`D:\users\Downloads\flecha.png`)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// fmt.Println(file)
	// return

	// var ctx *context.Context
	cc := user_log.NewGuploadServiceClient(conn)
	stream, err := cc.MethodName(context.Background())
	if err != nil {
		panic(err)
	}
	defer stream.CloseAndRecv()

	buf = make([]byte, 164868)
	for writing {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}
			err = errors.Wrapf(err,
				"errored while copying from file to buf")
			return
		}

		stream.Send(&user_log.Chunk{
			Content: buf[:n],
		})
	}

	status, err := stream.CloseAndRecv()
	if err != nil {
		err = errors.Wrapf(err,
			"failed to receive upstream status response")
		return
	}

	if status.Code != user_log.UploadStatusCode_Ok {
		err = errors.Errorf(
			"upload failed - msg: %s",
			status.Message)
		return
	}

	return
	// c := user_log.NewChatServiceClient(conn)
	// message := user_log.Log{
	// 	Body: &user_log.Log_Res{
	// 		Name:  "dde",
	// 		Brand: "sdf",
	// 	},
	// }

	// response, err := c.AddMessage(context.Background(), &message)

	// if err != nil {
	// 	panic(err)
	// }

	// log.Printf("response from server: %s", response)
}
