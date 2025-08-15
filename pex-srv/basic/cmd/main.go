package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "pex-srv/basic/init"
	__ "pex-srv/basic/proto"
	"pex-srv/handler/service"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterEMQXSrvServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
