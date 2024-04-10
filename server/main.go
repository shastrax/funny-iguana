package main

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/shastrax/funny-iguana/gen/iguana/v1"
	"github.com/shastrax/funny-iguana/gen/iguana/v1/iguanav1connect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/bufbuild/connect-go"
)

//v1 "github.com/shastrax/funny-iguana/gen/iguana/v1"
//"github.com/shastrax/funny-iguana/gen/iguana/v1/iguanav1connect"

const address = "localhost:8080"

func main() {
	path, handler := iguanav1connect.NewIguanaServiceHandler(&iguanaServiceServer{})

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

type iguanaServiceServer struct {
	iguanav1connect.UnimplementedIguanaServiceHandler
}

func (ss *iguanaServiceServer) Ping(ctx context.Context, req *connect.Request[v1.PingRequest]) (*connect.Response[petv1.PingResponse], error) {
	fmt.Println("ping ping ping")
	//name := req.Msg.GetName()
	//petType := req.Msg.GetPetType()
	//log.Printf("Got a request to create a %v named %s", petType, name)

	return connect.NewResponse(&v1.PingResponse{}), nil
}
