package main

import (
	"grp/pkg/adder"
	"grp/pkg/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

// type Endpoint func(ctx context.Context, request
// 	interface{})
// 	(response interface{}, err error)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
