package handler

import (
	"log"

	proto "github.com/LOG-ED/docker-micro/proto"

	"golang.org/x/net/context"
)

type Task struct{}

// Run handler set the task accordingly to the request method
func (u *Task) Run(ctx context.Context, req *proto.RunRequest, rsp *proto.RunResponse) error {
	log.Printf("Received Task.Run with method: %s", req.Method)

	switch req.Method {
	case "GET":
		rsp.StatusCode = proto.RunResponse_OK
	case "POST":
		rsp.StatusCode = proto.RunResponse_CREATED
	default:
		rsp.StatusCode = proto.RunResponse_FAILED
	}

	return nil
}
