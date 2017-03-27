package main

import (
	"log"
	"time"

	proto "github.com/LOG-ED/docker-micro/proto"
	"github.com/LOG-ED/docker-micro/server"
	grpc "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.api.task"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	proto.RegisterTaskHandler(service.Server(), new(server.Task))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}