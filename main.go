package main

import (
	"log"
	"time"

	"github.com/LOG-ED/docker-micro/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

func main() {
	service := micro.NewService(
		micro.Name("docker-micro.srv.task"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(handler.Task),
		),
	)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
