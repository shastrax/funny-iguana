package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"

	v1 "github.com/shastrax/funny-iguana/gen/iguana/v1"
	"github.com/shastrax/funny-iguana/gen/iguana/v1/iguanav1connect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// buf curl --schema proto --data '{"source":"bufcurl"}' http://localhost:8080/iguana.v1.IguanaService/Ping

type iguanaServiceServer struct {
	iguanav1connect.UnimplementedIguanaServiceHandler
}

func (ss *iguanaServiceServer) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[v1.PingResponse], error) {
	source := req.Msg.GetSource()
	fmt.Println("ping request source:", source)

	return connect.NewResponse(&v1.PingResponse{Status: "happy"}), nil
}

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	mux.Handle(iguanav1connect.NewIguanaServiceHandler(&iguanaServiceServer{}))
	err := http.ListenAndServe(address, h2c.NewHandler(mux, &http2.Server{}))
	log.Fatalf("listen failure: %v", err)
}
