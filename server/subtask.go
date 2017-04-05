package server

import (
	"log"

	proto "github.com/LOG-ED/docker-micro/proto"

	"golang.org/x/net/context"
)

type SubTask struct{}

// Run handler set the subtask exitcode accordingly to the operation
func (t *SubTask) Save(ctx context.Context, req *proto.SaveRequest, rsp *proto.SaveResponse) error {
	log.Printf("Received SubTask.Save with operation: %s", req.Operation)

	rsp.ExitCode = "0"

	return nil
}
