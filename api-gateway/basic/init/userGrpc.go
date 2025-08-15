package init

import (
	"api-gateway/basic/config"
	zhc "api-gateway/basic/user-proto"
	"api-gateway/untils"
	"flag"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func InitUserGrpc() {
	client, err := untils.NewConsulClient("14.103.175.131:8500")
	if err != nil {
		panic("init consul failed:" + err.Error())
	}
	consul := untils.Consul{
		ID:      uuid.NewString(),
		Name:    "user-srv",
		Tags:    []string{"health", "user"},
		Address: "127.0.0.1",
		Port:    50052,
	}

	err = client.RegisterConsul(consul)
	if err != nil {
		panic("init consul failed:" + err.Error())
	}

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	config.UserSrv = zhc.NewUserClient(conn)

}
