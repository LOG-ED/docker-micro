package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/micro/go-micro/metadata"

	"sourcegraph.com/sourcegraph/appdash"
	appdashot "sourcegraph.com/sourcegraph/appdash/opentracing"

	proto "github.com/LOG-ED/docker-micro/proto"
	opentracing "github.com/opentracing/opentracing-go"

	"github.com/gorilla/mux"
	micro "github.com/micro/go-micro"
	trace "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

const serviceName = "go.micro.api.task"

var cl proto.TaskClient

func main() {
	service := micro.NewService()
	service.Init()

	// Would it make sense to embed Appdash?
	addr := startAppdashServer(8700)
	tracer := appdashot.NewTracer(appdash.NewRemoteCollector(addr))

	opentracing.SetGlobalTracer(tracer)

	cw := trace.NewClientWrapper(tracer)

	// Set tracing map in context
	ctx := metadata.NewContext(context.Background(), map[string]string{})

	// wrap the generated client stub
	cl = proto.NewTaskClient(serviceName, cw(service.Client()))

	port := 9000
	bind := fmt.Sprintf(":%d", port)
	mux := mux.NewRouter()
	mux.Handle("/test", testHandler(ctx))
	log.Fatal(http.ListenAndServe(bind, mux))
}

func testHandler(c context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("trace %s request to %s", r.Method, r.URL.Path)
		md, _ := metadata.FromContext(c)
		sp, ctx := opentracing.StartSpanFromContext(c, "GET /test")
		defer sp.Finish()

		err := sp.Tracer().Inject(sp.Context(),
			opentracing.TextMap,
			opentracing.TextMapCarrier(md))
		if err != nil {
			log.Fatalf("%s: couldn't inject data (%v)", r.URL.Path, err)
		}

		rsp, err1 := cl.Run(ctx, &proto.RunRequest{
			Method: "GET",
		})
		if err1 != nil {
			sp.LogEventWithPayload("request error", err)
			return
		}

		w.Write([]byte("output: " + rsp.StatusCode.String()))
	})
}
