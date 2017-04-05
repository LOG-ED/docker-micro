package main

import (
	"log"
	"time"

	proto "github.com/LOG-ED/docker-micro/proto"
	"github.com/LOG-ED/docker-micro/server"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.subtask"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	proto.RegisterSubTaskHandler(service.Server(), new(server.SubTask))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
