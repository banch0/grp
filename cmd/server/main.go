package main

import (
	"grp/pkg/adder"
	"grp/pkg/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main(){
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	api.RegsterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}