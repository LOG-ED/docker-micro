package server

import (
	"log"

	proto "github.com/LOG-ED/docker-micro/proto"
	micro "github.com/micro/go-micro"

	"golang.org/x/net/context"
)

type Task struct{}

// Run handler set the task accordingly to the request method
func (t *Task) Run(ctx context.Context, req *proto.RunRequest, rsp *proto.RunResponse) error {
	log.Printf("Received Task.Run with method: %s", req.Method)

	service := micro.NewService()
	service.Init()

	// use the generated client stub to save a subtask
	cl := proto.NewSubTaskClient("go.micro.api.subtask", service.Client())
	_, err := cl.Save(context.Background(), &proto.SaveRequest{
		Operation: "TEST",
	})
	if err != nil {
		log.Println(err)
	}

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
