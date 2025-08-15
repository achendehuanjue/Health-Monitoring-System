package init

import (
	"api-gateway/basic/config"
	"api-gateway/untils"
	"flag"
	"github.com/google/uuid"
	"log"

	__ "api-gateway/basic/device-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGrpc() {
	client, err := untils.NewConsulClient("14.103.175.131:8500")
	if err != nil {
		panic("init consul failed:" + err.Error())
	}
	consul := untils.Consul{
		ID:      uuid.NewString(),
		Name:    "health-srv",
		Tags:    []string{"health", "monitor"},
		Address: "127.0.0.1",
		Port:    50051,
	}

	err = client.RegisterConsul(consul)
	if err != nil {
		panic("init consul failed:" + err.Error())
	}

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	config.DeviceSrv = __.NewDeviceSrvClient(conn)

}
