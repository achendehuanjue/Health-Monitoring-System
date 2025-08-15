package main

import (
	_ "device-srv/basic/init"
	__ "device-srv/basic/proto"
	"device-srv/handler/service"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterDeviceSrvServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
