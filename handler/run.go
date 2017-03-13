package handler

import (
	"log"

	proto "github.com/LOG-ED/docker-micro/proto"

	"golang.org/x/net/context"
)

type Task struct{}

// Run handler set the task as completed if a command is given
func (u *Task) Run(ctx context.Context, req *proto.RunRequest, rsp *proto.RunResponse) error {
	log.Print("Received Task.Run request")

	if req.Command == "" {
		rsp.Status = proto.RunResponse_FAILED
	} else {
		rsp.Status = proto.RunResponse_COMPLETED
	}
	return nil
}
