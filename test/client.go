package main

import (
	"fmt"
	"log"

	proto "github.com/LOG-ED/docker-micro/proto"
	"github.com/micro/go-micro/client"

	"golang.org/x/net/context"
)

func main() {

	c := client.NewClient()

	request := &proto.RunRequest{
		Command: proto.RunRequest_NONE,
	}

	req := c.NewRequest("docker-micro.srv.task", "Task.Run", &request)

	rsp := &proto.RunResponse{}

	// Call service
	if err := c.Call(context.Background(), req, rsp); err != nil {
		log.Fatal(err)
	}

	if rsp.Status == proto.RunResponse_FAILED {
		log.Fatalf("Wrong Response with status: %s", rsp.Status)
	}

	fmt.Printf("Successful Response with status: %s", rsp.Status)
}
