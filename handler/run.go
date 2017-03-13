package handler

import (
	"log"

	proto "github.com/LOG-ED/docker-micro/proto"

	"golang.org/x/net/context"
)

type Task struct{}

// Run handler set the task as completed if a POST request occured
func (u *Task) Run(ctx context.Context, req *proto.RunRequest, rsp *proto.RunResponse) error {
	log.Printf("Received Task.Run with method: %s", req.Method)

	if req.Method == "POST" {
		rsp.StatusCode = proto.RunResponse_COMPLETED
	} else {
		rsp.StatusCode = proto.RunResponse_FAILED
	}
	return nil
}
