package main

import (
	"context"
	"fmt"

	proto "github.com/LOG-ED/docker-micro/proto"
	"github.com/micro/go-grpc"
)

const serviceName = "go.micro.api.task"

var cl proto.TaskClient

func init() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl = proto.NewTaskClient(serviceName, service.Client())
}

func ExampleGetSuccessfulResponseIfRequestMethodIsGet() {

	rsp, err := cl.Run(context.Background(), &proto.RunRequest{
		Method: "GET",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
	//Output: statusCode:OK
}

func ExampleGetCreatedResponseIfRequestMethodIsPost() {

	rsp, err := cl.Run(context.Background(), &proto.RunRequest{
		Method: "POST",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
	//Output: statusCode:CREATED
}
