package main

import (
	"context"
	"log"

	"github.com/Shakhrik/Just-Golang/grpc_learn/proto"
)

type server struct{}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	log.Println(a, b)
	return nil, nil
}

func main() {

}
