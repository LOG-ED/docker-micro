package handler

import (
	"log"

	proto "github.com/log-ed/docker-micro/proto"

	"golang.org/x/net/context"
)

type Task struct{}

// Run handler set the task as completed
func (u *Task) Run(ctx context.Context, req *proto.RunRequest, rsp *proto.RunResponse) error {
	log.Print("Received Task.Run request")

	rsp.Status = proto.RunResponse_COMPLETED
	return nil
}
