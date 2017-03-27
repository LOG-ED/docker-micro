package main

import (
	"context"
	"fmt"

	proto "github.com/LOG-ED/docker-micro/proto"
	"github.com/micro/go-grpc"
)

func ExampleGetSuccessfulResponseIfRequestMethodIsGet() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl := proto.NewTaskClient("go.micro.srv.task", service.Client())

	rsp, err := cl.Run(context.Background(), &proto.RunRequest{
		Method: "GET",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
	//Output: {StatusCode:OK}
}

func ExampleGetCreatedResponseIfRequestMethodIsPost() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl := proto.NewTaskClient("go.micro.srv.task", service.Client())

	rsp, err := cl.Run(context.Background(), &proto.RunRequest{
		Method: "POST",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
	//Output: {StatusCode:CREATED}
}

func ExampleGetFailedResponseIfRequestMethodIsDelete() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl := proto.NewTaskClient("go.micro.srv.task", service.Client())

	rsp, err := cl.Run(context.Background(), &proto.RunRequest{
		Method: "DELETE",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
	//Output: {StatusCode:FAILED}
}
