package main

import (
	"log"
	"time"

	"github.com/LOG-ED/docker-micro/handler"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.task"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			new(handler.Task),
		),
	)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
