package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "user-srv/basic/init"
	__ "user-srv/basic/proto"
	"user-srv/handler/service"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterUserServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
