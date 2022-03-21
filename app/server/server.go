package server

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alirezakargar1380/grpc_mysql/app/model"
	"github.com/alirezakargar1380/grpc_mysql/user_log"
	"golang.org/x/net/context"
)

type Serverr struct {
}

func (s *Serverr) AddMessage(ctx context.Context, message *user_log.Log) (*user_log.Log, error) {
	data := model.Cpu{
		Name:  message.Body.Name,
		Brand: message.Body.Brand,
	}
	model.CreateBook(&data)
	log.Printf("recieved message from body: ")

	var asds user_log.Log_Res
	asds.Name = "server"
	return &user_log.Log{
		Body: &asds,
	}, nil
}

func (s *Serverr) MethodName(stream user_log.GuploadService_MethodNameServer) error {
	fmt.Println("fuck you im revibe")
	as, err := stream.Recv()
	if err != nil {
		panic(err)
	}
	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile("myfile.png", as.Content, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
	return nil
}
