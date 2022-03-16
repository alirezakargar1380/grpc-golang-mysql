package user_log

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) AddMessage(ctx context.Context, message *Log) (*Log, error) {
	log.Printf("recieved message from body: ")
	fmt.Println(message)
	return nil, nil
}
